# CreateOrganizationActionBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | Pointer to **bool** | Set to true for immediate execution. Set to false if the action should be previewed before executing. This property cannot be unset once it is true. Defaults to false. | [optional] 
**Synchronous** | Pointer to **bool** | Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. Defaults to false. | [optional] 
**Actions** | [**[]CreateOrganizationActionBatchRequestActionsInner**](CreateOrganizationActionBatchRequestActionsInner.md) | A set of changes to make as part of this action (&lt;a href&#x3D;&#39;https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/&#39;&gt;more details&lt;/a&gt;) | 

## Methods

### NewCreateOrganizationActionBatchRequest

`func NewCreateOrganizationActionBatchRequest(actions []CreateOrganizationActionBatchRequestActionsInner, ) *CreateOrganizationActionBatchRequest`

NewCreateOrganizationActionBatchRequest instantiates a new CreateOrganizationActionBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationActionBatchRequestWithDefaults

`func NewCreateOrganizationActionBatchRequestWithDefaults() *CreateOrganizationActionBatchRequest`

NewCreateOrganizationActionBatchRequestWithDefaults instantiates a new CreateOrganizationActionBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *CreateOrganizationActionBatchRequest) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *CreateOrganizationActionBatchRequest) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *CreateOrganizationActionBatchRequest) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *CreateOrganizationActionBatchRequest) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *CreateOrganizationActionBatchRequest) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *CreateOrganizationActionBatchRequest) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *CreateOrganizationActionBatchRequest) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *CreateOrganizationActionBatchRequest) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetActions

`func (o *CreateOrganizationActionBatchRequest) GetActions() []CreateOrganizationActionBatchRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateOrganizationActionBatchRequest) GetActionsOk() (*[]CreateOrganizationActionBatchRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateOrganizationActionBatchRequest) SetActions(v []CreateOrganizationActionBatchRequestActionsInner)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


