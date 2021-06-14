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
) (client interface{}, diags diag.Diagnostics) {
	shopifyDomain := d.Get("domain").(string)
	shopifyAccessToken := d.Get("access_token").(string)
	if shopifyDomain == "" || shopifyAccessToken == "" {
		diags = diag.Errorf("Please specify both 'domain' and 'access_token'")
		return
	}

	config := Config{
		ShopifyDomain:      shopifyDomain,
		ShopifyAccessToken: shopifyAccessToken,
	}
	client = config.NewClient()
	return
}
