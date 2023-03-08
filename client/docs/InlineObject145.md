# InlineObject145

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the template | [optional] 
**Body** | Pointer to **string** | The liquid template used for the body of the webhook message. | [optional] 
**Headers** | Pointer to [**[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1**](NetworksNetworkIdWebhooksPayloadTemplatesHeaders1.md) | The liquid template used with the webhook headers. | [optional] 
**BodyFile** | Pointer to **string** | A file containing liquid template used for the body of the webhook message. | [optional] 
**HeadersFile** | Pointer to **string** | A file containing the liquid template used with the webhook headers. | [optional] 

## Methods

### NewInlineObject145

`func NewInlineObject145() *InlineObject145`

NewInlineObject145 instantiates a new InlineObject145 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject145WithDefaults

`func NewInlineObject145WithDefaults() *InlineObject145`

NewInlineObject145WithDefaults instantiates a new InlineObject145 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject145) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject145) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject145) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject145) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBody

`func (o *InlineObject145) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineObject145) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineObject145) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineObject145) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineObject145) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineObject145) GetHeadersOk() (*[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineObject145) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineObject145) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBodyFile

`func (o *InlineObject145) GetBodyFile() string`

GetBodyFile returns the BodyFile field if non-nil, zero value otherwise.

### GetBodyFileOk

`func (o *InlineObject145) GetBodyFileOk() (*string, bool)`

GetBodyFileOk returns a tuple with the BodyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyFile

`func (o *InlineObject145) SetBodyFile(v string)`

SetBodyFile sets BodyFile field to given value.

### HasBodyFile

`func (o *InlineObject145) HasBodyFile() bool`

HasBodyFile returns a boolean if a field has been set.

### GetHeadersFile

`func (o *InlineObject145) GetHeadersFile() string`

GetHeadersFile returns the HeadersFile field if non-nil, zero value otherwise.

### GetHeadersFileOk

`func (o *InlineObject145) GetHeadersFileOk() (*string, bool)`

GetHeadersFileOk returns a tuple with the HeadersFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersFile

`func (o *InlineObject145) SetHeadersFile(v string)`

SetHeadersFile sets HeadersFile field to given value.

### HasHeadersFile

`func (o *InlineObject145) HasHeadersFile() bool`

HasHeadersFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


