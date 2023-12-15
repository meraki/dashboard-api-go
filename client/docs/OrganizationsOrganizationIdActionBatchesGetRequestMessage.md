# OrganizationsOrganizationIdActionBatchesGetRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId} | [optional] 
**OrganizationId** | Pointer to **string** | ID of the organization this action batch belongs to | [optional] 
**Confirmed** | Pointer to **bool** | Flag describing whether the action should be previewed before executing or not | [optional] 
**Synchronous** | Pointer to **bool** | Flag describing whether actions should run synchronously or asynchronously | [optional] 
**Status** | Pointer to [**CreateOrganizationActionBatch201ResponseStatus**](CreateOrganizationActionBatch201ResponseStatus.md) |  | [optional] 
**Actions** | [**[]CreateOrganizationActionBatch201ResponseActionsInner**](CreateOrganizationActionBatch201ResponseActionsInner.md) | A set of changes made as part of this action (&lt;a href&#x3D;&#39;https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/&#39;&gt;more details&lt;/a&gt;) | 

## Methods

### NewOrganizationsOrganizationIdActionBatchesGetRequestMessage

`func NewOrganizationsOrganizationIdActionBatchesGetRequestMessage(actions []CreateOrganizationActionBatch201ResponseActionsInner, ) *OrganizationsOrganizationIdActionBatchesGetRequestMessage`

NewOrganizationsOrganizationIdActionBatchesGetRequestMessage instantiates a new OrganizationsOrganizationIdActionBatchesGetRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdActionBatchesGetRequestMessageWithDefaults

`func NewOrganizationsOrganizationIdActionBatchesGetRequestMessageWithDefaults() *OrganizationsOrganizationIdActionBatchesGetRequestMessage`

NewOrganizationsOrganizationIdActionBatchesGetRequestMessageWithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesGetRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConfirmed

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetStatus() CreateOrganizationActionBatch201ResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetStatusOk() (*CreateOrganizationActionBatch201ResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) SetStatus(v CreateOrganizationActionBatch201ResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActions

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetActions() []CreateOrganizationActionBatch201ResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) GetActionsOk() (*[]CreateOrganizationActionBatch201ResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *OrganizationsOrganizationIdActionBatchesGetRequestMessage) SetActions(v []CreateOrganizationActionBatch201ResponseActionsInner)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


