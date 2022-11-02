# CreateNetworkHttpServersWebhookTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL where the test webhook will be sent | 
**SharedSecret** | Pointer to **string** | The shared secret the test webhook will send. Optional. Defaults to an empty string. | [optional] [default to ""]

## Methods

### NewCreateNetworkHttpServersWebhookTestRequest

`func NewCreateNetworkHttpServersWebhookTestRequest(url string, ) *CreateNetworkHttpServersWebhookTestRequest`

NewCreateNetworkHttpServersWebhookTestRequest instantiates a new CreateNetworkHttpServersWebhookTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkHttpServersWebhookTestRequestWithDefaults

`func NewCreateNetworkHttpServersWebhookTestRequestWithDefaults() *CreateNetworkHttpServersWebhookTestRequest`

NewCreateNetworkHttpServersWebhookTestRequestWithDefaults instantiates a new CreateNetworkHttpServersWebhookTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateNetworkHttpServersWebhookTestRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateNetworkHttpServersWebhookTestRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateNetworkHttpServersWebhookTestRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSharedSecret

`func (o *CreateNetworkHttpServersWebhookTestRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *CreateNetworkHttpServersWebhookTestRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *CreateNetworkHttpServersWebhookTestRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *CreateNetworkHttpServersWebhookTestRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


