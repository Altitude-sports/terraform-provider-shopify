package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/dghubble/sling"
)

type WebhookService struct {
	sling *sling.Sling
}

func newWebhookService(sling *sling.Sling) *WebhookService {
	return &WebhookService{
		sling: sling.New().Path("admin/").Path("webhooks/"),
	}
}

// -----------------------------------------------------------------------------
// Input
// -----------------------------------------------------------------------------

type WebhookInput struct {
	Topic   string `json:"topic"`
	Address string `json:"address"`

	Format                     string   `json:"format"`
	Fields                     []string `json:"fields"`
	MetafieldNamespaces        []string `json:"metafield_namespaces"`
	PrivateMetafieldNamespaces []string `json:"private_metafield_namespaces"`
}

type WebhookInputBody struct {
	Webhook WebhookInput `json:"webhook"`
}

// -----------------------------------------------------------------------------
// Responses
// -----------------------------------------------------------------------------

type Webhook struct {
	// Read-only fields
	Id         uint64 `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	ApiVersion string `json:"api_version"`

	WebhookInput
}

type WebhookResponse struct {
	Webhook Webhook `json:"webhook"`
}

type WebhookDeleteResponse struct{}

// -----------------------------------------------------------------------------
// Errors
// -----------------------------------------------------------------------------

type WebhookErrorMessage struct {
	WebhookMessage *string `json:"webhook,omitempty"`

	TopicMessage   *[]string `json:"topic,omitempty"`
	AddressMessage *[]string `json:"address,omitempty"`
	FormatMessage  *[]string `json:"format,omitempty"`

	FieldsMessage                     *[]string `json:"fields,omitempty"`
	MetafieldNamespacesMessage        *[]string `json:"metafield_namespaces,omitempty"`
	PrivateMetafieldNamespacesMessage *[]string `json:"private_metafield_namespaces,omitempty"`
}

type WebhookError struct {
	Errors WebhookErrorMessage `json:"errors"`
}

func (err WebhookError) Error() string {
	if err == (WebhookError{}) {
		return ""
	}

	webhookMessage := ""
	if err.Errors.WebhookMessage != nil {
		webhookMessage = *err.Errors.WebhookMessage
	}

	topicMessage := ""
	if err.Errors.TopicMessage != nil {
		topicMessage = fmt.Sprintf("topic: %v", *err.Errors.TopicMessage)
	}

	addressMessage := ""
	if err.Errors.AddressMessage != nil {
		addressMessage = fmt.Sprintf("address: %v", *err.Errors.AddressMessage)
	}

	formatMessage := ""
	if err.Errors.FormatMessage != nil {
		formatMessage = fmt.Sprintf("format: %v", *err.Errors.FormatMessage)
	}

	fieldsMessage := ""
	if err.Errors.FieldsMessage != nil {
		fieldsMessage = fmt.Sprintf("fields: %v", *err.Errors.FieldsMessage)
	}

	metafieldNamespacesMessage := ""
	if err.Errors.MetafieldNamespacesMessage != nil {
		metafieldNamespacesMessage = fmt.Sprintf(
			"metafield_namespaces: %v",
			*err.Errors.MetafieldNamespacesMessage,
		)
	}

	privateMetafieldNamespacesMessage := ""
	if err.Errors.PrivateMetafieldNamespacesMessage != nil {
		privateMetafieldNamespacesMessage = fmt.Sprintf(
			"private_metafield_namespaces: %v",
			*err.Errors.PrivateMetafieldNamespacesMessage,
		)
	}

	return strings.Join([]string{
		"Shopify:",
		webhookMessage,
		topicMessage,
		addressMessage,
		formatMessage,
		fieldsMessage,
		metafieldNamespacesMessage,
		privateMetafieldNamespacesMessage,
	}, " ")
}

// -----------------------------------------------------------------------------
// CRUD
// -----------------------------------------------------------------------------

func (service *WebhookService) Create(params *WebhookInput) (Webhook, *http.Response, error) {
	webhookResponse := new(WebhookResponse)
	genericError := new(json.RawMessage)
	payload := &WebhookInputBody{
		Webhook: *params,
	}
	request := service.sling.Post("").BodyJSON(payload)
	log.Printf("[DEBUG] request to Shopify: %s", request)

	resp, err := request.Receive(webhookResponse, genericError)
	return webhookResponse.Webhook, resp, relevantError(resp, err, genericError, new(WebhookError))
}

func (service *WebhookService) Read(webhookId string) (Webhook, *http.Response, error) {
	webhookResponse := new(WebhookResponse)
	genericError := new(json.RawMessage)
	path := url.PathEscape(webhookId)
	request := service.sling.Get(path)
	log.Printf("[DEBUG] request to Shopify: %s", request)

	resp, err := request.Receive(webhookResponse, genericError)
	return webhookResponse.Webhook, resp, relevantError(resp, err, genericError, new(WebhookError))
}

func (service *WebhookService) Update(webhookId string, params *WebhookInput) (Webhook, *http.Response, error) {
	webhookResponse := new(WebhookResponse)
	genericError := new(json.RawMessage)
	path := url.PathEscape(webhookId)
	payload := &WebhookInputBody{
		Webhook: *params,
	}
	request := service.sling.Put(path).BodyJSON(payload)
	log.Printf("[DEBUG] request to Shopify: %s", request)

	resp, err := request.Receive(webhookResponse, genericError)
	return webhookResponse.Webhook, resp, relevantError(resp, err, genericError, new(WebhookError))
}

func (service *WebhookService) Delete(webhookId string) (*http.Response, error) {
	webhookDeleteResponse := new(WebhookDeleteResponse)
	genericError := new(json.RawMessage)
	path := url.PathEscape(webhookId)
	request := service.sling.Delete(path)
	log.Printf("[DEBUG] request to Shopify: %s", request)

	resp, err := request.Receive(webhookDeleteResponse, genericError)
	return resp, relevantError(resp, err, genericError, new(WebhookError))
}
