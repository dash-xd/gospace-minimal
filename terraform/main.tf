# google_cloudfunctions2_function's build_config.source only accepts a
# GCS object reference, never a local path - there's no Terraform
# resource that deploys "whatever's on disk" directly. Since the goal
# here is deploying the checkout this workflow already made (not an
# artifact Terraform computes and manages, e.g. via archive_file +
# google_storage_bucket_object), this shells out to `gcloud functions
# deploy --source`, which stages and uploads the local source dir
# itself. It's still driven by `terraform apply`, just via a
# null_resource provisioner rather than a declarative resource.
#
# Because there's no persisted state, this can't tell "already
# deployed, update it" from "never deployed, create it" the way a
# normal `gcloud functions deploy` re-run would - deploy is inherently
# an upsert. So the conflict check below runs first and fails the
# provisioner (and so the apply) if the function already exists,
# preserving the create-only, error-on-conflict behavior this config
# is meant to have.
locals {
  deploy_flags = compact([
    var.service_account_email != "" ? "--service-account=${var.service_account_email}" : "",
    var.allow_unauthenticated ? "--allow-unauthenticated" : "--no-allow-unauthenticated",
  ])
}

resource "null_resource" "deploy" {
  triggers = {
    always_run = timestamp()
  }

  provisioner "local-exec" {
    interpreter = ["bash", "-c"]
    command     = <<-EOT
      set -euo pipefail

      if gcloud functions describe "${var.function_name}" --gen2 \
          --project="${var.project_id}" --region="${var.region}" \
          >/dev/null 2>&1; then
        echo "Error: function '${var.function_name}' already exists in ${var.project_id}/${var.region}. This config keeps no Terraform state between runs, so it only ever creates a function, never updates one - deploy under a name that's free, or delete the existing function first." >&2
        exit 1
      fi

      gcloud functions deploy "${var.function_name}" --gen2 --quiet \
        --project="${var.project_id}" \
        --region="${var.region}" \
        --runtime="${var.runtime}" \
        --entry-point="${var.entry_point}" \
        --source="${var.source_dir}" \
        --set-build-env-vars="GOOGLE_FUNCTION_SOURCE=${var.entry_point_dir}" \
        --trigger-http \
        --max-instances="${var.max_instance_count}" \
        --memory="${var.available_memory}" \
        --timeout="${var.timeout_seconds}" \
        --ingress-settings="${var.ingress_settings}" \
        ${join(" ", local.deploy_flags)}
    EOT
  }
}

# The provisioner above has no attributes of its own to expose, so the
# outputs come from reading back the function it just created.
data "google_cloudfunctions2_function" "this" {
  name       = var.function_name
  location   = var.region
  project    = var.project_id
  depends_on = [null_resource.deploy]
}
