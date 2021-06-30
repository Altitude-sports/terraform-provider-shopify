package main

import (
	"github.com/Altitude-sports/terraform-provider-shopify/shopify"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shopify.Provider,
	})
}
