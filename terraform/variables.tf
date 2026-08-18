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
    makes this apply fail with Terraform's own resource-already-exists
    error - pick a name you know is free, or one meant to be unique
    per deploy.
  EOT
  type        = string
}

variable "source_dir" {
  description = "Local path to the repo checkout to zip and deploy as the function's source."
  type        = string
}

variable "source_bucket" {
  description = <<-EOT
    Name of an existing GCS bucket to stage the source zip in. Must
    already exist and be readable by Cloud Build - this config does not
    create or manage it, so re-running against the same bucket never
    conflicts.
  EOT
  type        = string
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
