# UpdateNetworkWebhooksPayloadTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the template | [optional] 
**Body** | Pointer to **string** | The liquid template used for the body of the webhook message. | [optional] 
**Headers** | Pointer to [**[]CreateNetworkWebhooksPayloadTemplateRequestHeadersInner**](CreateNetworkWebhooksPayloadTemplateRequestHeadersInner.md) | The liquid template used with the webhook headers. | [optional] 
**BodyFile** | Pointer to **string** | A file containing liquid template used for the body of the webhook message. | [optional] 
**HeadersFile** | Pointer to **string** | A file containing the liquid template used with the webhook headers. | [optional] 

## Methods

### NewUpdateNetworkWebhooksPayloadTemplateRequest

`func NewUpdateNetworkWebhooksPayloadTemplateRequest() *UpdateNetworkWebhooksPayloadTemplateRequest`

NewUpdateNetworkWebhooksPayloadTemplateRequest instantiates a new UpdateNetworkWebhooksPayloadTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWebhooksPayloadTemplateRequestWithDefaults

`func NewUpdateNetworkWebhooksPayloadTemplateRequestWithDefaults() *UpdateNetworkWebhooksPayloadTemplateRequest`

NewUpdateNetworkWebhooksPayloadTemplateRequestWithDefaults instantiates a new UpdateNetworkWebhooksPayloadTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBody

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeaders() []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeadersOk() (*[]CreateNetworkWebhooksPayloadTemplateRequestHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetHeaders(v []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBodyFile

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBodyFile() string`

GetBodyFile returns the BodyFile field if non-nil, zero value otherwise.

### GetBodyFileOk

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetBodyFileOk() (*string, bool)`

GetBodyFileOk returns a tuple with the BodyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyFile

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetBodyFile(v string)`

SetBodyFile sets BodyFile field to given value.

### HasBodyFile

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasBodyFile() bool`

HasBodyFile returns a boolean if a field has been set.

### GetHeadersFile

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeadersFile() string`

GetHeadersFile returns the HeadersFile field if non-nil, zero value otherwise.

### GetHeadersFileOk

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) GetHeadersFileOk() (*string, bool)`

GetHeadersFileOk returns a tuple with the HeadersFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersFile

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) SetHeadersFile(v string)`

SetHeadersFile sets HeadersFile field to given value.

### HasHeadersFile

`func (o *UpdateNetworkWebhooksPayloadTemplateRequest) HasHeadersFile() bool`

HasHeadersFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


