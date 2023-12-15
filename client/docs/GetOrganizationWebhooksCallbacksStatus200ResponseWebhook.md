# GetOrganizationWebhooksCallbacksStatus200ResponseWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The webhook receiver URL where the callback will be sent | [optional] 
**HttpServer** | Pointer to [**GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer**](GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate**](GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate.md) |  | [optional] 
**SentAt** | Pointer to **time.Time** | The timestamp the callback was sent to the webhook receiver | [optional] 

## Methods

### NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhook

`func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhook() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook`

NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhook instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookWithDefaults

`func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookWithDefaults() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook`

NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookWithDefaults instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHttpServer

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetHttpServer() GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetHttpServerOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetHttpServer(v GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetPayloadTemplate() GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetPayloadTemplateOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetPayloadTemplate(v GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetSentAt

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


