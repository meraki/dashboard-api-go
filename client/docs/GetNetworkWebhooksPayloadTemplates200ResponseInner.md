# GetNetworkWebhooksPayloadTemplates200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadTemplateId** | Pointer to **string** | Webhook payload template Id | [optional] 
**Type** | Pointer to **string** | The type of the payload template | [optional] 
**Name** | Pointer to **string** | The name of the payload template | [optional] 
**Headers** | Pointer to [**[]GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner**](GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner.md) | The payload template headers, will be rendered as a key-value pair in the webhook. | [optional] 
**Body** | Pointer to **string** | The body of the payload template, in liquid template | [optional] 

## Methods

### NewGetNetworkWebhooksPayloadTemplates200ResponseInner

`func NewGetNetworkWebhooksPayloadTemplates200ResponseInner() *GetNetworkWebhooksPayloadTemplates200ResponseInner`

NewGetNetworkWebhooksPayloadTemplates200ResponseInner instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWebhooksPayloadTemplates200ResponseInnerWithDefaults

`func NewGetNetworkWebhooksPayloadTemplates200ResponseInnerWithDefaults() *GetNetworkWebhooksPayloadTemplates200ResponseInner`

NewGetNetworkWebhooksPayloadTemplates200ResponseInnerWithDefaults instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadTemplateId

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetPayloadTemplateId() string`

GetPayloadTemplateId returns the PayloadTemplateId field if non-nil, zero value otherwise.

### GetPayloadTemplateIdOk

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetPayloadTemplateIdOk() (*string, bool)`

GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateId

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetPayloadTemplateId(v string)`

SetPayloadTemplateId sets PayloadTemplateId field to given value.

### HasPayloadTemplateId

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasPayloadTemplateId() bool`

HasPayloadTemplateId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHeaders

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetHeaders() []GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetHeadersOk() (*[]GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetHeaders(v []GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


