variable "app_version" {
  description = "Application version"
  type        = string
}

variable "environment" {
  description = "Deployment environment"
  type        = string
}

resource "ulid_random" "deployment_id" {
  keepers = {
    app_version = var.app_version
    environment = var.environment
  }
  prefix = "deploy_"
}

# The ULID will be regenerated whenever app_version or environment changes