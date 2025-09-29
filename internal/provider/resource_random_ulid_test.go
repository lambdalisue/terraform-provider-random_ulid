package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccResourceRandomULID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRandomULIDConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"random_ulid.test", "id", regexp.MustCompile("^[0-9A-Z]{26}$")),
					resource.TestCheckResourceAttrSet(
						"random_ulid.test", "timestamp"),
					resource.TestCheckResourceAttrSet(
						"random_ulid.test", "randomness"),
				),
			},
			{
				Config: testAccResourceRandomULIDWithPrefixConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"random_ulid.test_prefix", "id", regexp.MustCompile("^test_[0-9A-Z]{26}$")),
				),
			},
		},
	})
}

const testAccResourceRandomULIDConfig = `
resource "random_ulid" "test" {}
`

const testAccResourceRandomULIDWithPrefixConfig = `
resource "random_ulid" "test_prefix" {
  prefix = "test_"
}
`

func testAccPreCheck(t *testing.T) {
	// Provider pre-check can be extended here if needed
}

var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"random_ulid": func() (*schema.Provider, error) {
		return New("test")(), nil
	},
}
