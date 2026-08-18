output "url" {
  description = "HTTPS URL of the deployed function."
  value       = google_cloudfunctions2_function.this.url
}

output "function_name" {
  description = "Name of the deployed function, as created."
  value       = google_cloudfunctions2_function.this.name
}
