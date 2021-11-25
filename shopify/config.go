package shopify

import (
	shopify "github.com/Altitude-sports/terraform-provider-shopify/shopify/internal/client"
)

type Config struct {
	ShopifyDomain      string
	ShopifyAccessToken string
	ShopifyApiVersion  string
}

func (c *Config) NewClient() *shopify.Client {
	return shopify.NewClient(c.ShopifyDomain, c.ShopifyAccessToken, c.ShopifyApiVersion)
}
