terraform {
  required_providers {
    ulid = {
      source  = "lambdalisue/ulid"
      version = "~> 0.1.0"
    }
  }
}

# Configure the ULID Provider
provider "ulid" {}

# Generate a basic ULID
resource "ulid_random" "example" {}

# Generate a ULID with a prefix
resource "ulid_random" "user_id" {
  prefix = "user_"
}

# Generate ULIDs with keepers for controlled regeneration
resource "ulid_random" "deployment" {
  keepers = {
    version = var.app_version
    region  = var.deployment_region
  }
  prefix = "deploy_"
}

# Output the generated ULID
output "example_ulid" {
  value       = ulid_random.example.id
  description = "A generated ULID"
}