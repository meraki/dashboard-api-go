# CreateNetworkWebhooksWebhookTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL where the test webhook will be sent | 
**SharedSecret** | Pointer to **string** | The shared secret the test webhook will send. Optional. Defaults to an empty string. | [optional] [default to ""]
**PayloadTemplateId** | Pointer to **string** | The ID of the payload template of the test webhook. Defaults to the HTTP server&#39;s template ID if one exists for the given URL, or Generic template ID otherwise | [optional] 
**PayloadTemplateName** | Pointer to **string** | The name of the payload template. | [optional] 
**AlertTypeId** | Pointer to **string** | The type of alert which the test webhook will send. Optional. Defaults to power_supply_down. | [optional] [default to "power_supply_down"]

## Methods

### NewCreateNetworkWebhooksWebhookTestRequest

`func NewCreateNetworkWebhooksWebhookTestRequest(url string, ) *CreateNetworkWebhooksWebhookTestRequest`

NewCreateNetworkWebhooksWebhookTestRequest instantiates a new CreateNetworkWebhooksWebhookTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWebhooksWebhookTestRequestWithDefaults

`func NewCreateNetworkWebhooksWebhookTestRequestWithDefaults() *CreateNetworkWebhooksWebhookTestRequest`

NewCreateNetworkWebhooksWebhookTestRequestWithDefaults instantiates a new CreateNetworkWebhooksWebhookTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateNetworkWebhooksWebhookTestRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSharedSecret

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *CreateNetworkWebhooksWebhookTestRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *CreateNetworkWebhooksWebhookTestRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPayloadTemplateId

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateId() string`

GetPayloadTemplateId returns the PayloadTemplateId field if non-nil, zero value otherwise.

### GetPayloadTemplateIdOk

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateIdOk() (*string, bool)`

GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateId

`func (o *CreateNetworkWebhooksWebhookTestRequest) SetPayloadTemplateId(v string)`

SetPayloadTemplateId sets PayloadTemplateId field to given value.

### HasPayloadTemplateId

`func (o *CreateNetworkWebhooksWebhookTestRequest) HasPayloadTemplateId() bool`

HasPayloadTemplateId returns a boolean if a field has been set.

### GetPayloadTemplateName

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateName() string`

GetPayloadTemplateName returns the PayloadTemplateName field if non-nil, zero value otherwise.

### GetPayloadTemplateNameOk

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateNameOk() (*string, bool)`

GetPayloadTemplateNameOk returns a tuple with the PayloadTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateName

`func (o *CreateNetworkWebhooksWebhookTestRequest) SetPayloadTemplateName(v string)`

SetPayloadTemplateName sets PayloadTemplateName field to given value.

### HasPayloadTemplateName

`func (o *CreateNetworkWebhooksWebhookTestRequest) HasPayloadTemplateName() bool`

HasPayloadTemplateName returns a boolean if a field has been set.

### GetAlertTypeId

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetAlertTypeId() string`

GetAlertTypeId returns the AlertTypeId field if non-nil, zero value otherwise.

### GetAlertTypeIdOk

`func (o *CreateNetworkWebhooksWebhookTestRequest) GetAlertTypeIdOk() (*string, bool)`

GetAlertTypeIdOk returns a tuple with the AlertTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeId

`func (o *CreateNetworkWebhooksWebhookTestRequest) SetAlertTypeId(v string)`

SetAlertTypeId sets AlertTypeId field to given value.

### HasAlertTypeId

`func (o *CreateNetworkWebhooksWebhookTestRequest) HasAlertTypeId() bool`

HasAlertTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


