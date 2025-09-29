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
						"ulid_random.test", "id", regexp.MustCompile("^[0-9A-Z]{26}$")),
					resource.TestCheckResourceAttrSet(
						"ulid_random.test", "timestamp"),
					resource.TestCheckResourceAttrSet(
						"ulid_random.test", "randomness"),
				),
			},
			{
				Config: testAccResourceRandomULIDWithPrefixConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"ulid_random.test_prefix", "id", regexp.MustCompile("^test_[0-9A-Z]{26}$")),
				),
			},
		},
	})
}

const testAccResourceRandomULIDConfig = `
resource "ulid_random" "test" {}
`

const testAccResourceRandomULIDWithPrefixConfig = `
resource "ulid_random" "test_prefix" {
  prefix = "test_"
}
`

func testAccPreCheck(t *testing.T) {
	// Provider pre-check can be extended here if needed
}

var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"ulid": func() (*schema.Provider, error) {
		return New("test")(), nil
	},
}
