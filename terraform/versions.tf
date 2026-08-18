# No backend block on purpose: this config always uses Terraform's
# default local-file state, and the workflow that runs it never
# persists, caches, or uploads that file. Every apply starts from empty
# state - fire and forget. That means a name already in use in the
# target project (from a previous deploy, or anything else) surfaces as
# an idiomatic deploy error rather than being silently reconciled; see
# main.tf and the .github/workflows/deploy-gcf.yml workflow.

terraform {
  required_version = ">= 1.5.0"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 6.0.0, < 7.0.0"
    }
    null = {
      source  = "hashicorp/null"
      version = ">= 3.2.0, < 4.0.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}
