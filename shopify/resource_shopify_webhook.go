package shopify

import (
	shopify "./shopify-go"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceShopifyWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceShopifyWebhookCreate,
		Read:   resourceShopifyWebhookRead,
		Update: resourceShopifyWebhookUpdate,
		Delete: resourceShopifyWebhookDelete,

		Schema: map[string]*schema.Schema{
			"topic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"format": {
				Type:     schema.TypeString,
				Required: true,
			},
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceShopifyWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*shopify.Client)

	params := &shopify.WebhookInput{
		Topic:   d.Get("topic").(string),
		Address: d.Get("address").(string),
		Format:  d.Get("format").(string),
	}

	webhook, _, err := client.Webhooks.Create(params)
	if err != nil {
		return fmt.Errorf("Error creating Shopify webhook: %s", err)
	}

	d.SetId(fmt.Sprintf("%d", webhook.Id))
	d.Set("topic", webhook.Topic)
	d.Set("address", webhook.Address)
	d.Set("format", webhook.Format)

	return nil
}

func resourceShopifyWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*shopify.Client)

	params := &shopify.WebhookInput{
		Topic:   d.Get("topic").(string),
		Address: d.Get("address").(string),
		Format:  d.Get("format").(string),
	}

	webhook, _, err := client.Webhooks.Update(d.Id(), params)
	if err != nil {
		return fmt.Errorf("Error updating Shopify webhook: %s", err)
	}

	d.Set("topic", webhook.Topic)
	d.Set("address", webhook.Address)
	d.Set("format", webhook.Format)

	return nil
}

func resourceShopifyWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*shopify.Client)

	webhook, _, err := client.Webhooks.Read(d.Id())
	if err != nil {
		return fmt.Errorf("Error reading Shopify webhook: %s", err)
	}

	d.Set("topic", webhook.Topic)
	d.Set("address", webhook.Address)
	d.Set("format", webhook.Format)

	return nil
}

func resourceShopifyWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*shopify.Client)

	_, err := client.Webhooks.Delete(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Shopify webhook: %s", err)
	}

	return nil
}
