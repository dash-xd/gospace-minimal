data "archive_file" "source" {
  type        = "zip"
  source_dir  = var.source_dir
  output_path = "${path.module}/.build/${var.function_name}-source.zip"

  excludes = [
    ".git/**",
    ".github/**",
    "terraform/**",
    ".tmp/**",
    "localserve.log",
    "localserve.pid",
  ]
}

# Named after the zip's own content hash, so re-deploying identical
# source reuses the same object and re-deploying changed source never
# collides with what's already in the bucket - only the function
# resource itself is meant to fail on conflict.
resource "google_storage_bucket_object" "source" {
  name   = "gospace-minimal/${var.function_name}/${data.archive_file.source.output_md5}.zip"
  bucket = var.source_bucket
  source = data.archive_file.source.output_path
}

resource "google_cloudfunctions2_function" "this" {
  name     = var.function_name
  location = var.region

  build_config {
    runtime     = var.runtime
    entry_point = var.entry_point

    source {
      storage_source {
        bucket = var.source_bucket
        object = google_storage_bucket_object.source.name
      }
    }
  }

  service_config {
    max_instance_count             = var.max_instance_count
    available_memory               = var.available_memory
    timeout_seconds                = var.timeout_seconds
    service_account_email          = var.service_account_email != "" ? var.service_account_email : null
    ingress_settings               = var.ingress_settings
    all_traffic_on_latest_revision = true
  }
}

resource "google_cloud_run_v2_service_iam_member" "invoker" {
  count = var.allow_unauthenticated ? 1 : 0

  project  = var.project_id
  location = var.region
  name     = google_cloudfunctions2_function.this.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}
