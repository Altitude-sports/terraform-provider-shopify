# Shopify Terraform Provider

## Setup

From within the Shopify Admin Interface:

1. Click on "Apps" from the list on the left hand sidebar (or go to
   _yourstore_.myshopify.com/admin/apps)
2. Click on "Manage private apps" (or go to
   _yourstore_.myshopify.com/admin/apps/private)
3. Create a new private app
4. Configure the provider as follows:

```hcl
provider "shopify" {
  access_token = "<app password>"
  domain       = "<yourstore>.myshopify.com"
}
```

## Resources

### `shopify_webhook`

For reference, see [Shopify's Webhook
Documentation](https://help.shopify.com/en/api/reference/events/webhook).

**Note: Webhooks created via this resource are not visible within the Shopify
Admin GUI. To view webhook created this resource, an API request needs to be
made with the same app credentials passed into this provider.**

For example:

```bash
export SHOPIFY_ACCESS_TOKEN='your app password'
export STORE='your store'
curl "https://${STORE}.myshopify.com/admin/webhooks.json" \
  -H "Accept: application/json" \
  -H "X-Shopify-Access-Token: ${SHOPIFY_ACCESS_TOKEN}"
```

#### Example Usage

```hcl
resource "shopify_webhook" "example" {
  address = "https://shopify-webhook-handler.devaltitude.com"
  topic   = "orders/create"
  format  = "json"
}
```

#### Argument Reference

The following arguments are supported:

* topic (Required) - The event topic for which webhook messages should be sent.
  See the Shopify documentation for the full list of available topics.
* address (Required) - The full URL to send webhooks to
* format (Required) - The format to send webhook messages in. Can be either
  `json` or `xml`.

#### Attributes Reference

In addition to all arguments above, the following attributes are exported:

* topic - The event topic for which webhook messages are sent
* address - The full URL to send webhooks to
* format - The format to send webhook messages in. Will be one of `json` or
  `xml`.

## Building The Provider

```shell
make build
```
