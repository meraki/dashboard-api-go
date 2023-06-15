# InlineResponse2014Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Flag describing whether all actions in the action batch have completed | [optional] 
**Failed** | Pointer to **bool** | Flag describing whether any actions in the action batch failed | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered when running actions in the action batch | [optional] 
**CreatedResources** | [**[]InlineResponse2014StatusCreatedResources**](InlineResponse2014StatusCreatedResources.md) | Resources created as a result of this action batch | 

## Methods

### NewInlineResponse2014Status

`func NewInlineResponse2014Status(createdResources []InlineResponse2014StatusCreatedResources, ) *InlineResponse2014Status`

NewInlineResponse2014Status instantiates a new InlineResponse2014Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2014StatusWithDefaults

`func NewInlineResponse2014StatusWithDefaults() *InlineResponse2014Status`

NewInlineResponse2014StatusWithDefaults instantiates a new InlineResponse2014Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *InlineResponse2014Status) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse2014Status) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse2014Status) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse2014Status) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetFailed

`func (o *InlineResponse2014Status) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *InlineResponse2014Status) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *InlineResponse2014Status) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *InlineResponse2014Status) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse2014Status) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse2014Status) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse2014Status) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse2014Status) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedResources

`func (o *InlineResponse2014Status) GetCreatedResources() []InlineResponse2014StatusCreatedResources`

GetCreatedResources returns the CreatedResources field if non-nil, zero value otherwise.

### GetCreatedResourcesOk

`func (o *InlineResponse2014Status) GetCreatedResourcesOk() (*[]InlineResponse2014StatusCreatedResources, bool)`

GetCreatedResourcesOk returns a tuple with the CreatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedResources

`func (o *InlineResponse2014Status) SetCreatedResources(v []InlineResponse2014StatusCreatedResources)`

SetCreatedResources sets CreatedResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


