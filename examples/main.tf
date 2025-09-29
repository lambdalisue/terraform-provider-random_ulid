terraform {
  required_providers {
    ulid = {
      source  = "lambdalisue/ulid"
      version = "~> 0.1.0"
    }
  }
}

# Simple ULID generation
resource "ulid_random" "example" {}

# ULID with prefix
resource "ulid_random" "with_prefix" {
  prefix = "user_"
}

# Using keepers to control regeneration
resource "ulid_random" "with_keepers" {
  keepers = {
    # New ULID will be generated when this value changes
    deployment_id = "v1.0.0"
  }
  prefix = "deploy_"
}

# Generate multiple ULIDs
resource "ulid_random" "multiple" {
  count  = 3
  prefix = "item_${count.index}_"
}

# Outputs
output "simple_ulid" {
  description = "Generated simple ULID"
  value       = ulid_random.example.id
}

output "ulid_with_prefix" {
  description = "ULID with prefix"
  value       = ulid_random.with_prefix.id
}

output "ulid_timestamp" {
  description = "ULID timestamp in milliseconds"
  value       = ulid_random.example.timestamp
}

output "multiple_ulids" {
  description = "Multiple ULIDs"
  value       = ulid_random.multiple[*].id
}