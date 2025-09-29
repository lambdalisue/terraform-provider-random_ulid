resource "ulid_random" "example" {}

output "generated_ulid" {
  value = ulid_random.example.id
}