package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/lambdalisue/terraform-provider-ulid/internal/provider"
)

// These variables are set by goreleaser during build
var (
	version string = "dev"
)

func main() {
	opts := &plugin.ServeOpts{
		ProviderFunc: provider.New(version),
	}

	plugin.Serve(opts)
}
