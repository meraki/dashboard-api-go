# CreateNetworkWebhooksHttpServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for easy reference to the HTTP server | 
**Url** | **string** | The URL of the HTTP server. Once set, cannot be updated. | 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki. | [optional] 
**PayloadTemplate** | Pointer to [**CreateNetworkWebhooksHttpServerRequestPayloadTemplate**](CreateNetworkWebhooksHttpServerRequestPayloadTemplate.md) |  | [optional] 

## Methods

### NewCreateNetworkWebhooksHttpServerRequest

`func NewCreateNetworkWebhooksHttpServerRequest(name string, url string, ) *CreateNetworkWebhooksHttpServerRequest`

NewCreateNetworkWebhooksHttpServerRequest instantiates a new CreateNetworkWebhooksHttpServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWebhooksHttpServerRequestWithDefaults

`func NewCreateNetworkWebhooksHttpServerRequestWithDefaults() *CreateNetworkWebhooksHttpServerRequest`

NewCreateNetworkWebhooksHttpServerRequestWithDefaults instantiates a new CreateNetworkWebhooksHttpServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkWebhooksHttpServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkWebhooksHttpServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkWebhooksHttpServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CreateNetworkWebhooksHttpServerRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateNetworkWebhooksHttpServerRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateNetworkWebhooksHttpServerRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSharedSecret

`func (o *CreateNetworkWebhooksHttpServerRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *CreateNetworkWebhooksHttpServerRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *CreateNetworkWebhooksHttpServerRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *CreateNetworkWebhooksHttpServerRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *CreateNetworkWebhooksHttpServerRequest) GetPayloadTemplate() CreateNetworkWebhooksHttpServerRequestPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *CreateNetworkWebhooksHttpServerRequest) GetPayloadTemplateOk() (*CreateNetworkWebhooksHttpServerRequestPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *CreateNetworkWebhooksHttpServerRequest) SetPayloadTemplate(v CreateNetworkWebhooksHttpServerRequestPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *CreateNetworkWebhooksHttpServerRequest) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


