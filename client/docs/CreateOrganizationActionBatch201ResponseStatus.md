# CreateOrganizationActionBatch201ResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Flag describing whether all actions in the action batch have completed | [optional] 
**Failed** | Pointer to **bool** | Flag describing whether any actions in the action batch failed | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered when running actions in the action batch | [optional] 
**CreatedResources** | [**[]CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner**](CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner.md) | Resources created as a result of this action batch | 

## Methods

### NewCreateOrganizationActionBatch201ResponseStatus

`func NewCreateOrganizationActionBatch201ResponseStatus(createdResources []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner, ) *CreateOrganizationActionBatch201ResponseStatus`

NewCreateOrganizationActionBatch201ResponseStatus instantiates a new CreateOrganizationActionBatch201ResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationActionBatch201ResponseStatusWithDefaults

`func NewCreateOrganizationActionBatch201ResponseStatusWithDefaults() *CreateOrganizationActionBatch201ResponseStatus`

NewCreateOrganizationActionBatch201ResponseStatusWithDefaults instantiates a new CreateOrganizationActionBatch201ResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *CreateOrganizationActionBatch201ResponseStatus) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *CreateOrganizationActionBatch201ResponseStatus) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetFailed

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CreateOrganizationActionBatch201ResponseStatus) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CreateOrganizationActionBatch201ResponseStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetErrors

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateOrganizationActionBatch201ResponseStatus) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateOrganizationActionBatch201ResponseStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedResources

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetCreatedResources() []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner`

GetCreatedResources returns the CreatedResources field if non-nil, zero value otherwise.

### GetCreatedResourcesOk

`func (o *CreateOrganizationActionBatch201ResponseStatus) GetCreatedResourcesOk() (*[]CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner, bool)`

GetCreatedResourcesOk returns a tuple with the CreatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedResources

`func (o *CreateOrganizationActionBatch201ResponseStatus) SetCreatedResources(v []CreateOrganizationActionBatch201ResponseStatusCreatedResourcesInner)`

SetCreatedResources sets CreatedResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


