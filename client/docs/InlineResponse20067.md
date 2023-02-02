# InlineResponse20067

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadTemplateId** | Pointer to **string** | Webhook payload template Id | [optional] 
**Type** | Pointer to **string** | The type of the payload template | [optional] 
**Name** | Pointer to **string** | The name of the payload template | [optional] 
**Headers** | Pointer to [**[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders**](NetworksNetworkIdWebhooksPayloadTemplatesHeaders.md) | The payload template headers, will be rendered as a key-value pair in the webhook. | [optional] 
**Body** | Pointer to **string** | The body of the payload template, in liquid template | [optional] 

## Methods

### NewInlineResponse20067

`func NewInlineResponse20067() *InlineResponse20067`

NewInlineResponse20067 instantiates a new InlineResponse20067 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20067WithDefaults

`func NewInlineResponse20067WithDefaults() *InlineResponse20067`

NewInlineResponse20067WithDefaults instantiates a new InlineResponse20067 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadTemplateId

`func (o *InlineResponse20067) GetPayloadTemplateId() string`

GetPayloadTemplateId returns the PayloadTemplateId field if non-nil, zero value otherwise.

### GetPayloadTemplateIdOk

`func (o *InlineResponse20067) GetPayloadTemplateIdOk() (*string, bool)`

GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateId

`func (o *InlineResponse20067) SetPayloadTemplateId(v string)`

SetPayloadTemplateId sets PayloadTemplateId field to given value.

### HasPayloadTemplateId

`func (o *InlineResponse20067) HasPayloadTemplateId() bool`

HasPayloadTemplateId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20067) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20067) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20067) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20067) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20067) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20067) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20067) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20067) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineResponse20067) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineResponse20067) GetHeadersOk() (*[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineResponse20067) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineResponse20067) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *InlineResponse20067) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineResponse20067) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineResponse20067) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineResponse20067) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


