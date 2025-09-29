resource "ulid_random" "product_id" {
  prefix = "prod_"
}

resource "ulid_random" "order_id" {
  prefix = "ord_"
}

resource "aws_dynamodb_table_item" "product" {
  table_name = aws_dynamodb_table.products.name
  hash_key   = aws_dynamodb_table.products.hash_key

  item = jsonencode({
    id = {
      S = ulid_random.product_id.id
    }
    name = {
      S = "Example Product"
    }
    created_at = {
      N = tostring(ulid_random.product_id.timestamp)
    }
  })
}