package client

import (
	"fmt"

	"github.com/dghubble/sling"
)

type Client struct {
	sling    *sling.Sling
	Webhooks *WebhookService
}

func NewClient(shopifyDomain string, shopifyAccessToken string) *Client {
	baseUrl := fmt.Sprintf("https://%s.myshopify.com/", shopifyDomain)
	base := sling.New().Base(
		baseUrl,
	).Set(
		"Accept", "application/json",
	).Set(
		"Content-Type", "application/json",
	).Set(
		"X-Shopify-Access-Token", shopifyAccessToken,
	)

	return &Client{
		sling:    base,
		Webhooks: newWebhookService(base),
	}
}
