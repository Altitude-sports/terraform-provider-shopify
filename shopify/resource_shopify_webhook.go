package shopify

import (
	"context"
	"log"
	"net/http"
	"strconv"

	shopify "github.com/Altitude-sports/terraform-provider-shopify/shopify/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceShopifyWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceShopifyWebhookCreate,
		ReadContext:   resourceShopifyWebhookRead,
		UpdateContext: resourceShopifyWebhookUpdate,
		DeleteContext: resourceShopifyWebhookDelete,

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
				Default:  "json",
				Optional: true,
			},
			"fields": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"metafield_namespaces": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"private_metafield_namespaces": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceShopifyWebhookCreate(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	fields := []string{}
	if v, ok := d.GetOk("fields"); ok {
		for _, val := range v.([]interface{}) {
			fields = append(fields, val.(string))
		}
	}

	metafieldNamespaces := []string{}
	if v, ok := d.GetOk("metafield_namespaces"); ok {
		for _, val := range v.([]interface{}) {
			metafieldNamespaces = append(metafieldNamespaces, val.(string))
		}
	}

	privateMetafieldNamespaces := []string{}
	if v, ok := d.GetOk("private_metafield_namespaces"); ok {
		for _, val := range v.([]interface{}) {
			privateMetafieldNamespaces = append(privateMetafieldNamespaces, val.(string))
		}
	}

	params := &shopify.WebhookInput{
		Topic:   d.Get("topic").(string),
		Address: d.Get("address").(string),

		Format:                     d.Get("format").(string),
		Fields:                     fields,
		MetafieldNamespaces:        metafieldNamespaces,
		PrivateMetafieldNamespaces: privateMetafieldNamespaces,
	}

	config := meta.(Config)
	client := config.NewClient()

	webhook, _, err := client.Webhooks.Create(params)
	if err != nil {
		return diag.Errorf("could not create Shopify webhook: %s", err)
	}

	d.SetId(strconv.FormatUint(webhook.Id, 10))

	_ = d.Set("topic", webhook.Topic)
	_ = d.Set("address", webhook.Address)

	_ = d.Set("format", webhook.Format)
	_ = d.Set("fields", webhook.Fields)
	_ = d.Set("metafield_namespaces", webhook.MetafieldNamespaces)
	_ = d.Set("private_metafield_namespaces", webhook.PrivateMetafieldNamespaces)

	return resourceShopifyWebhookRead(ctx, d, meta)
}

func resourceShopifyWebhookUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	fields := []string{}
	if v, ok := d.GetOk("fields"); ok {
		for _, val := range v.([]interface{}) {
			fields = append(fields, val.(string))
		}
	}

	metafieldNamespaces := []string{}
	if v, ok := d.GetOk("metafield_namespaces"); ok {
		for _, val := range v.([]interface{}) {
			metafieldNamespaces = append(metafieldNamespaces, val.(string))
		}
	}

	privateMetafieldNamespaces := []string{}
	if v, ok := d.GetOk("private_metafield_namespaces"); ok {
		for _, val := range v.([]interface{}) {
			privateMetafieldNamespaces = append(privateMetafieldNamespaces, val.(string))
		}
	}

	params := &shopify.WebhookInput{
		Topic:   d.Get("topic").(string),
		Address: d.Get("address").(string),

		Format:                     d.Get("format").(string),
		Fields:                     fields,
		MetafieldNamespaces:        metafieldNamespaces,
		PrivateMetafieldNamespaces: privateMetafieldNamespaces,
	}

	config := meta.(Config)
	client := config.NewClient()

	webhook, _, err := client.Webhooks.Update(d.Id(), params)
	if err != nil {
		return diag.Errorf("could not update Shopify webhook: %s", err)
	}

	_ = d.Set("topic", webhook.Topic)
	_ = d.Set("address", webhook.Address)

	_ = d.Set("format", webhook.Format)
	_ = d.Set("fields", webhook.Fields)
	_ = d.Set("metafield_namespaces", webhook.MetafieldNamespaces)
	_ = d.Set("private_metafield_namespaces", webhook.PrivateMetafieldNamespaces)

	return resourceShopifyWebhookRead(ctx, d, meta)
}

func resourceShopifyWebhookRead(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	config := meta.(Config)
	client := config.NewClient()

	webhook, resp, err := client.Webhooks.Read(d.Id())
	if resp.StatusCode == http.StatusNotFound {
		log.Printf("[DEBUG] Removing webhook ID='%s' from state because it no longer exists in Shopify", d.Id())
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.Errorf("could not retrieve Shopify webhook details: %s", err)
	}

	_ = d.Set("topic", webhook.Topic)
	_ = d.Set("address", webhook.Address)

	_ = d.Set("format", webhook.Format)
	_ = d.Set("fields", webhook.Fields)
	_ = d.Set("metafield_namespaces", webhook.MetafieldNamespaces)
	_ = d.Set("private_metafield_namespaces", webhook.PrivateMetafieldNamespaces)

	return nil
}

func resourceShopifyWebhookDelete(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	config := meta.(Config)
	client := config.NewClient()

	_, err := client.Webhooks.Delete(d.Id())
	if err != nil {
		return diag.Errorf("could not delete Shopify webhook: %s", err)
	}

	return nil
}
