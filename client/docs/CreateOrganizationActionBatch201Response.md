# CreateOrganizationActionBatch201Response

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

### NewCreateOrganizationActionBatch201Response

`func NewCreateOrganizationActionBatch201Response(actions []CreateOrganizationActionBatch201ResponseActionsInner, ) *CreateOrganizationActionBatch201Response`

NewCreateOrganizationActionBatch201Response instantiates a new CreateOrganizationActionBatch201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationActionBatch201ResponseWithDefaults

`func NewCreateOrganizationActionBatch201ResponseWithDefaults() *CreateOrganizationActionBatch201Response`

NewCreateOrganizationActionBatch201ResponseWithDefaults instantiates a new CreateOrganizationActionBatch201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrganizationActionBatch201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrganizationActionBatch201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrganizationActionBatch201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrganizationActionBatch201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateOrganizationActionBatch201Response) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOrganizationActionBatch201Response) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOrganizationActionBatch201Response) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateOrganizationActionBatch201Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConfirmed

`func (o *CreateOrganizationActionBatch201Response) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *CreateOrganizationActionBatch201Response) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *CreateOrganizationActionBatch201Response) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *CreateOrganizationActionBatch201Response) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *CreateOrganizationActionBatch201Response) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *CreateOrganizationActionBatch201Response) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *CreateOrganizationActionBatch201Response) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *CreateOrganizationActionBatch201Response) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetStatus

`func (o *CreateOrganizationActionBatch201Response) GetStatus() CreateOrganizationActionBatch201ResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOrganizationActionBatch201Response) GetStatusOk() (*CreateOrganizationActionBatch201ResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOrganizationActionBatch201Response) SetStatus(v CreateOrganizationActionBatch201ResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateOrganizationActionBatch201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActions

`func (o *CreateOrganizationActionBatch201Response) GetActions() []CreateOrganizationActionBatch201ResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateOrganizationActionBatch201Response) GetActionsOk() (*[]CreateOrganizationActionBatch201ResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateOrganizationActionBatch201Response) SetActions(v []CreateOrganizationActionBatch201ResponseActionsInner)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


