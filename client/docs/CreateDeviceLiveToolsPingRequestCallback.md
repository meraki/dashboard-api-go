# CreateDeviceLiveToolsPingRequestCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The callback URL for the webhook target. If using this field, please also specify a sharedSecret. | [optional] 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url. | [optional] 
**HttpServer** | Pointer to [**CreateDeviceLiveToolsPingRequestCallbackHttpServer**](CreateDeviceLiveToolsPingRequestCallbackHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate**](CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate.md) |  | [optional] 

## Methods

### NewCreateDeviceLiveToolsPingRequestCallback

`func NewCreateDeviceLiveToolsPingRequestCallback() *CreateDeviceLiveToolsPingRequestCallback`

NewCreateDeviceLiveToolsPingRequestCallback instantiates a new CreateDeviceLiveToolsPingRequestCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceLiveToolsPingRequestCallbackWithDefaults

`func NewCreateDeviceLiveToolsPingRequestCallbackWithDefaults() *CreateDeviceLiveToolsPingRequestCallback`

NewCreateDeviceLiveToolsPingRequestCallbackWithDefaults instantiates a new CreateDeviceLiveToolsPingRequestCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateDeviceLiveToolsPingRequestCallback) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateDeviceLiveToolsPingRequestCallback) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSharedSecret

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *CreateDeviceLiveToolsPingRequestCallback) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *CreateDeviceLiveToolsPingRequestCallback) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetHttpServer

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetHttpServer() CreateDeviceLiveToolsPingRequestCallbackHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetHttpServerOk() (*CreateDeviceLiveToolsPingRequestCallbackHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *CreateDeviceLiveToolsPingRequestCallback) SetHttpServer(v CreateDeviceLiveToolsPingRequestCallbackHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *CreateDeviceLiveToolsPingRequestCallback) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetPayloadTemplate() CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *CreateDeviceLiveToolsPingRequestCallback) GetPayloadTemplateOk() (*CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *CreateDeviceLiveToolsPingRequestCallback) SetPayloadTemplate(v CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *CreateDeviceLiveToolsPingRequestCallback) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


