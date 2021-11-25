package shopify

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHOPIFY_DOMAIN", nil),
				Description: "Domain of the Shopify store",
			},
			"access_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHOPIFY_ACCESS_TOKEN", nil),
				Description: "Shopify access token",
			},
			"api_version": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHOPIFY_API_VERSION", nil),
				Description: "Shopify API version",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"shopify_webhook": resourceShopifyWebhook(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(
	ctx context.Context,
	d *schema.ResourceData,
) (config interface{}, diags diag.Diagnostics) {
	shopifyDomain := d.Get("domain").(string)
	shopifyAccessToken := d.Get("access_token").(string)
	shopifyApiVersion := d.Get("api_version").(string)
	if shopifyDomain == "" || shopifyAccessToken == "" {
		diags = diag.Errorf("Please specify both 'domain' and 'access_token'")
		return
	}

	config = Config{
		ShopifyDomain:      shopifyDomain,
		ShopifyAccessToken: shopifyAccessToken,
		ShopifyApiVersion:  shopifyApiVersion,
	}

	return
}
