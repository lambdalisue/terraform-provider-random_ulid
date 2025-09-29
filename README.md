# Terraform Random ULID Provider

[![CI](https://github.com/lambdalisue/terraform-provider-random_ulid/actions/workflows/ci.yml/badge.svg)](https://github.com/lambdalisue/terraform-provider-random_ulid/actions/workflows/ci.yml)
[![Release](https://github.com/lambdalisue/terraform-provider-random_ulid/actions/workflows/release.yml/badge.svg)](https://github.com/lambdalisue/terraform-provider-random_ulid/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lambdalisue/terraform-provider-random_ulid)](https://goreportcard.com/report/github.com/lambdalisue/terraform-provider-random_ulid)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Terraform provider for generating ULIDs (Universally Unique Lexicographically Sortable Identifiers).

## What is ULID?

ULID is a specification for generating 128-bit identifiers with the following properties:

- **Lexicographically sortable**: ULIDs can be sorted by their string representation
- **Timestamp-based**: Contains a millisecond precision timestamp
- **Cryptographically secure randomness**: 80 bits of randomness per millisecond
- **URL safe**: Uses Crockford's base32 for string encoding
- **Case insensitive**: No ambiguous characters

Example ULID: `01ARZ3NDEKTSV4RRFFQ69G5FAV`

## Installation

### Using the Provider

```hcl
terraform {
  required_providers {
    random_ulid = {
      source  = "lambdalisue/random_ulid"
      version = "~> 0.1.0"
    }
  }
}

provider "random_ulid" {}
```

### Building from Source

```bash
git clone https://github.com/lambdalisue/terraform-provider-random_ulid.git
cd terraform-provider-random_ulid
make install
```

## Usage

### Basic Example

```hcl
resource "random_ulid" "example" {}

output "ulid" {
  value = random_ulid.example.id
}
```

### With Prefix

```hcl
resource "random_ulid" "user_id" {
  prefix = "user_"
}

# Output: user_01ARZ3NDEKTSV4RRFFQ69G5FAV
```

### Using Keepers

Keepers allow you to force regeneration of the ULID when specific values change:

```hcl
resource "random_ulid" "deployment" {
  keepers = {
    version = var.app_version
    region  = var.deployment_region
  }
  prefix = "deploy_"
}
```

### Multiple ULIDs

```hcl
resource "random_ulid" "items" {
  count  = 5
  prefix = "item_"
}

output "item_ids" {
  value = random_ulid.items[*].id
}
```

## Resource Arguments

### `random_ulid`

#### Arguments

- `keepers` - (Optional) Arbitrary map of values that, when changed, will trigger recreation of the resource.
- `prefix` - (Optional) Arbitrary string to prefix the ULID with.

#### Attributes

- `id` - The generated ULID string (with prefix if specified).
- `timestamp` - The timestamp component of the ULID in milliseconds since Unix epoch.
- `randomness` - The randomness component of the ULID as a hexadecimal string.

## Import

Existing ULIDs can be imported:

```bash
terraform import random_ulid.example 01ARZ3NDEKTSV4RRFFQ69G5FAV
```

Or with a prefix:

```bash
terraform import random_ulid.example user_01ARZ3NDEKTSV4RRFFQ69G5FAV
```

## Development

### Requirements

- Go 1.19+
- Terraform 0.12+

### Building

```bash
make build
```

### Testing

```bash
make test
```

### Installing Locally

```bash
make install
```

### Debugging

For debugging the provider during development, you can use Terraform's provider development overrides. Create or update `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "lambdalisue/random_ulid" = "/path/to/your/built/provider"
  }
  direct {}
}
```

Then run your Terraform commands with `TF_LOG=DEBUG` for detailed logging.

## License

This project is licensed under the MIT License.

## CI/CD

This project uses simple GitHub Actions workflows:

- **CI**: Runs on every push and pull request
  - Code formatting check
  - `go vet` for potential issues
  - Unit tests
  - Build verification

- **Release**: Triggered by version tags (v*)
  - Automated cross-platform builds with GoReleaser
  - GitHub release with changelog

### Creating a Release

```bash
git tag v0.1.0
git push origin v0.1.0
```

GoReleaser will automatically create binaries for Linux, macOS, and Windows.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

1. Fork and clone the repository
2. Install Go 1.20 or later
3. Run `go mod download` to fetch dependencies
4. Make your changes
5. Run `make test` to ensure tests pass
6. Submit a pull request

All pull requests must pass CI checks before merging.