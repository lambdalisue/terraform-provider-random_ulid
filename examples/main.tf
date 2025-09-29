terraform {
  required_providers {
    random_ulid = {
      source  = "lambdalisue/random_ulid"
      version = "~> 0.1.0"
    }
  }
}

# Simple ULID generation
resource "random_ulid" "example" {}

# ULID with prefix
resource "random_ulid" "with_prefix" {
  prefix = "user_"
}

# Using keepers to control regeneration
resource "random_ulid" "with_keepers" {
  keepers = {
    # New ULID will be generated when this value changes
    deployment_id = "v1.0.0"
  }
  prefix = "deploy_"
}

# Generate multiple ULIDs
resource "random_ulid" "multiple" {
  count  = 3
  prefix = "item_${count.index}_"
}

# Outputs
output "simple_ulid" {
  description = "Generated simple ULID"
  value       = random_ulid.example.id
}

output "ulid_with_prefix" {
  description = "ULID with prefix"
  value       = random_ulid.with_prefix.id
}

output "ulid_timestamp" {
  description = "ULID timestamp in milliseconds"
  value       = random_ulid.example.timestamp
}

output "multiple_ulids" {
  description = "Multiple ULIDs"
  value       = random_ulid.multiple[*].id
}