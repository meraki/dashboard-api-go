# CreateNetworkWebhooksHttpServerRequestPayloadTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadTemplateId** | Pointer to **string** | The ID of the payload template. Defaults to &#39;wpt_00001&#39; for the Meraki template. For Meraki-included templates: for the Webex (included) template use &#39;wpt_00002&#39;; for the Slack (included) template use &#39;wpt_00003&#39;; for the Microsoft Teams (included) template use &#39;wpt_00004&#39;; for the ServiceNow (included) template use &#39;wpt_00006&#39; | [optional] 
**Name** | Pointer to **string** | The name of the payload template. | [optional] 

## Methods

### NewCreateNetworkWebhooksHttpServerRequestPayloadTemplate

`func NewCreateNetworkWebhooksHttpServerRequestPayloadTemplate() *CreateNetworkWebhooksHttpServerRequestPayloadTemplate`

NewCreateNetworkWebhooksHttpServerRequestPayloadTemplate instantiates a new CreateNetworkWebhooksHttpServerRequestPayloadTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults

`func NewCreateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults() *CreateNetworkWebhooksHttpServerRequestPayloadTemplate`

NewCreateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults instantiates a new CreateNetworkWebhooksHttpServerRequestPayloadTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadTemplateId

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateId() string`

GetPayloadTemplateId returns the PayloadTemplateId field if non-nil, zero value otherwise.

### GetPayloadTemplateIdOk

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateIdOk() (*string, bool)`

GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateId

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) SetPayloadTemplateId(v string)`

SetPayloadTemplateId sets PayloadTemplateId field to given value.

### HasPayloadTemplateId

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) HasPayloadTemplateId() bool`

HasPayloadTemplateId returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


