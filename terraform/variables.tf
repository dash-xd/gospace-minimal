variable "project_id" {
  description = "GCP project to deploy the function into."
  type        = string
}

variable "region" {
  description = "GCP region for the function."
  type        = string
  default     = "us-central1"
}

variable "function_name" {
  description = <<-EOT
    Name of the Cloud Function (2nd gen) to create. Since this config
    keeps no state between runs, a name already in use in the project
    makes this apply fail (see main.tf) - pick a name you know is
    free, or one meant to be unique per deploy.
  EOT
  type        = string
}

variable "source_dir" {
  description = <<-EOT
    Local path to the already-checked-out repo to deploy as-is. gcloud
    stages and uploads it itself - this config does not zip or manage
    a GCS object for it.
  EOT
  type        = string
}

variable "entry_point_dir" {
  description = <<-EOT
    Path, relative to source_dir, of the package containing
    entry_point - passed to the buildpack via the
    GOOGLE_FUNCTION_SOURCE build environment variable, since
    entry_point lives under internal/ rather than at the module root.
  EOT
  type        = string
  default     = "internal/function"
}

variable "runtime" {
  description = "Cloud Functions Go runtime to build with."
  type        = string
  default     = "go122"
}

variable "entry_point" {
  description = "Exported function/handler name the buildpack invokes."
  type        = string
  default     = "Main"
}

variable "service_account_email" {
  description = "Runtime service account for the deployed function. Empty uses the project's default compute service account."
  type        = string
  default     = ""
}

variable "max_instance_count" {
  description = "Maximum concurrent instances."
  type        = number
  default     = 1
}

variable "available_memory" {
  description = "Memory allotted per instance, e.g. \"256M\"."
  type        = string
  default     = "256M"
}

variable "timeout_seconds" {
  description = "Per-request timeout."
  type        = number
  default     = 60
}

variable "ingress_settings" {
  description = "One of ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB."
  type        = string
  default     = "ALLOW_ALL"
}

variable "allow_unauthenticated" {
  description = "Grant roles/run.invoker to allUsers so the function can be called without an identity token."
  type        = bool
  default     = false
}
