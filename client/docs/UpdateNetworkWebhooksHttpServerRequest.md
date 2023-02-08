# UpdateNetworkWebhooksHttpServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for easy reference to the HTTP server | [optional] 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki. | [optional] 
**PayloadTemplate** | Pointer to [**UpdateNetworkWebhooksHttpServerRequestPayloadTemplate**](UpdateNetworkWebhooksHttpServerRequestPayloadTemplate.md) |  | [optional] 

## Methods

### NewUpdateNetworkWebhooksHttpServerRequest

`func NewUpdateNetworkWebhooksHttpServerRequest() *UpdateNetworkWebhooksHttpServerRequest`

NewUpdateNetworkWebhooksHttpServerRequest instantiates a new UpdateNetworkWebhooksHttpServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWebhooksHttpServerRequestWithDefaults

`func NewUpdateNetworkWebhooksHttpServerRequestWithDefaults() *UpdateNetworkWebhooksHttpServerRequest`

NewUpdateNetworkWebhooksHttpServerRequestWithDefaults instantiates a new UpdateNetworkWebhooksHttpServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkWebhooksHttpServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkWebhooksHttpServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkWebhooksHttpServerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkWebhooksHttpServerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedSecret

`func (o *UpdateNetworkWebhooksHttpServerRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *UpdateNetworkWebhooksHttpServerRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *UpdateNetworkWebhooksHttpServerRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *UpdateNetworkWebhooksHttpServerRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *UpdateNetworkWebhooksHttpServerRequest) GetPayloadTemplate() UpdateNetworkWebhooksHttpServerRequestPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *UpdateNetworkWebhooksHttpServerRequest) GetPayloadTemplateOk() (*UpdateNetworkWebhooksHttpServerRequestPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *UpdateNetworkWebhooksHttpServerRequest) SetPayloadTemplate(v UpdateNetworkWebhooksHttpServerRequestPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *UpdateNetworkWebhooksHttpServerRequest) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


