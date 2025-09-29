// Package provider implements the ULID Terraform provider.
//
//go:generate tfplugindocs generate --provider-name ulid
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: map[string]*schema.Schema{},
			ResourcesMap: map[string]*schema.Resource{
				"ulid_random": resourceRandomULID(),
			},
			DataSourcesMap:       map[string]*schema.Resource{},
			ConfigureContextFunc: configure,
		}
	}
}

func configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return nil, nil
}
