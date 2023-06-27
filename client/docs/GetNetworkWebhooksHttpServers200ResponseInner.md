# GetNetworkWebhooksHttpServers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A Base64 encoded ID. | [optional] 
**Name** | Pointer to **string** | A name for easy reference to the HTTP server | [optional] 
**Url** | Pointer to **string** | The URL of the HTTP server. | [optional] 
**NetworkId** | Pointer to **string** | A Meraki network ID. | [optional] 
**PayloadTemplate** | Pointer to [**GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate**](GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate.md) |  | [optional] 

## Methods

### NewGetNetworkWebhooksHttpServers200ResponseInner

`func NewGetNetworkWebhooksHttpServers200ResponseInner() *GetNetworkWebhooksHttpServers200ResponseInner`

NewGetNetworkWebhooksHttpServers200ResponseInner instantiates a new GetNetworkWebhooksHttpServers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWebhooksHttpServers200ResponseInnerWithDefaults

`func NewGetNetworkWebhooksHttpServers200ResponseInnerWithDefaults() *GetNetworkWebhooksHttpServers200ResponseInner`

NewGetNetworkWebhooksHttpServers200ResponseInnerWithDefaults instantiates a new GetNetworkWebhooksHttpServers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetPayloadTemplate() GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) GetPayloadTemplateOk() (*GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) SetPayloadTemplate(v GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *GetNetworkWebhooksHttpServers200ResponseInner) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


