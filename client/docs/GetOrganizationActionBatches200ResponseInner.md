# GetOrganizationActionBatches200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId} | [optional] 
**OrganizationId** | Pointer to **string** | ID of the organization this action batch belongs to | [optional] 
**Confirmed** | Pointer to **bool** | Flag describing whether the action should be previewed before executing or not | [optional] 
**Synchronous** | Pointer to **bool** | Flag describing whether actions should run synchronously or asynchronously | [optional] 
**Status** | Pointer to [**GetOrganizationActionBatches200ResponseInnerStatus**](GetOrganizationActionBatches200ResponseInnerStatus.md) |  | [optional] 
**Actions** | [**[]GetOrganizationActionBatches200ResponseInnerActionsInner**](GetOrganizationActionBatches200ResponseInnerActionsInner.md) | A set of changes made as part of this action (&lt;a href&#x3D;&#39;https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/&#39;&gt;more details&lt;/a&gt;) | 

## Methods

### NewGetOrganizationActionBatches200ResponseInner

`func NewGetOrganizationActionBatches200ResponseInner(actions []GetOrganizationActionBatches200ResponseInnerActionsInner, ) *GetOrganizationActionBatches200ResponseInner`

NewGetOrganizationActionBatches200ResponseInner instantiates a new GetOrganizationActionBatches200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationActionBatches200ResponseInnerWithDefaults

`func NewGetOrganizationActionBatches200ResponseInnerWithDefaults() *GetOrganizationActionBatches200ResponseInner`

NewGetOrganizationActionBatches200ResponseInnerWithDefaults instantiates a new GetOrganizationActionBatches200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationActionBatches200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationActionBatches200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationActionBatches200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationActionBatches200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetOrganizationActionBatches200ResponseInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetOrganizationActionBatches200ResponseInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetOrganizationActionBatches200ResponseInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetOrganizationActionBatches200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConfirmed

`func (o *GetOrganizationActionBatches200ResponseInner) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *GetOrganizationActionBatches200ResponseInner) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *GetOrganizationActionBatches200ResponseInner) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *GetOrganizationActionBatches200ResponseInner) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *GetOrganizationActionBatches200ResponseInner) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *GetOrganizationActionBatches200ResponseInner) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *GetOrganizationActionBatches200ResponseInner) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *GetOrganizationActionBatches200ResponseInner) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationActionBatches200ResponseInner) GetStatus() GetOrganizationActionBatches200ResponseInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationActionBatches200ResponseInner) GetStatusOk() (*GetOrganizationActionBatches200ResponseInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationActionBatches200ResponseInner) SetStatus(v GetOrganizationActionBatches200ResponseInnerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationActionBatches200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActions

`func (o *GetOrganizationActionBatches200ResponseInner) GetActions() []GetOrganizationActionBatches200ResponseInnerActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetOrganizationActionBatches200ResponseInner) GetActionsOk() (*[]GetOrganizationActionBatches200ResponseInnerActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetOrganizationActionBatches200ResponseInner) SetActions(v []GetOrganizationActionBatches200ResponseInnerActionsInner)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


