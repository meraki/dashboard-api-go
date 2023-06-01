# InlineObject144

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for easy reference to the HTTP server | [optional] 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki. | [optional] 
**PayloadTemplate** | Pointer to [**NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate**](NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate.md) |  | [optional] 

## Methods

### NewInlineObject144

`func NewInlineObject144() *InlineObject144`

NewInlineObject144 instantiates a new InlineObject144 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject144WithDefaults

`func NewInlineObject144WithDefaults() *InlineObject144`

NewInlineObject144WithDefaults instantiates a new InlineObject144 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject144) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject144) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject144) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject144) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedSecret

`func (o *InlineObject144) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *InlineObject144) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *InlineObject144) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *InlineObject144) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineObject144) GetPayloadTemplate() NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineObject144) GetPayloadTemplateOk() (*NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineObject144) SetPayloadTemplate(v NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineObject144) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


