resource "ulid_random" "api_keys" {
  count  = 3
  prefix = "apikey_"
}

output "api_key_ids" {
  value = ulid_random.api_keys[*].id
}