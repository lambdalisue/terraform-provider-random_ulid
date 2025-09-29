resource "ulid_random" "user" {
  prefix = "user_"
}

resource "ulid_random" "session" {
  prefix = "sess_"
}

# Example outputs:
# user_01ARZ3NDEKTSV4RRFFQ69G5FAV
# sess_01ARZ3NDEKTSV4RRFFQ69G5FBX