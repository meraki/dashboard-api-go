# UpdateNetworkWebhooksHttpServerRequestPayloadTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadTemplateId** | Pointer to **string** | The ID of the payload template. Defaults to &#39;wpt_00001&#39; for the Meraki template. For Meraki-included templates: for the Webex (included) template use &#39;wpt_00002&#39;; for the Slack (included) template use &#39;wpt_00003&#39;; for the Microsoft Teams (included) template use &#39;wpt_00004&#39;; for the ServiceNow (included) template use &#39;wpt_00006&#39; | [optional] 

## Methods

### NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplate

`func NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplate() *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate`

NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplate instantiates a new UpdateNetworkWebhooksHttpServerRequestPayloadTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults

`func NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults() *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate`

NewUpdateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults instantiates a new UpdateNetworkWebhooksHttpServerRequestPayloadTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadTemplateId

`func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateId() string`

GetPayloadTemplateId returns the PayloadTemplateId field if non-nil, zero value otherwise.

### GetPayloadTemplateIdOk

`func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateIdOk() (*string, bool)`

GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateId

`func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) SetPayloadTemplateId(v string)`

SetPayloadTemplateId sets PayloadTemplateId field to given value.

### HasPayloadTemplateId

`func (o *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) HasPayloadTemplateId() bool`

HasPayloadTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


