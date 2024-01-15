# GetOrganizationActionBatches200ResponseInnerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Flag describing whether all actions in the action batch have completed | [optional] 
**Failed** | Pointer to **bool** | Flag describing whether any actions in the action batch failed | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered when running actions in the action batch | [optional] 
**CreatedResources** | [**[]GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner**](GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner.md) | Resources created as a result of this action batch | 

## Methods

### NewGetOrganizationActionBatches200ResponseInnerStatus

`func NewGetOrganizationActionBatches200ResponseInnerStatus(createdResources []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner, ) *GetOrganizationActionBatches200ResponseInnerStatus`

NewGetOrganizationActionBatches200ResponseInnerStatus instantiates a new GetOrganizationActionBatches200ResponseInnerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationActionBatches200ResponseInnerStatusWithDefaults

`func NewGetOrganizationActionBatches200ResponseInnerStatusWithDefaults() *GetOrganizationActionBatches200ResponseInnerStatus`

NewGetOrganizationActionBatches200ResponseInnerStatusWithDefaults instantiates a new GetOrganizationActionBatches200ResponseInnerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetFailed

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetErrors

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedResources

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCreatedResources() []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner`

GetCreatedResources returns the CreatedResources field if non-nil, zero value otherwise.

### GetCreatedResourcesOk

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCreatedResourcesOk() (*[]GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner, bool)`

GetCreatedResourcesOk returns a tuple with the CreatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedResources

`func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetCreatedResources(v []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner)`

SetCreatedResources sets CreatedResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


