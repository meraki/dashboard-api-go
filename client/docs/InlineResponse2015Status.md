# InlineResponse2015Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Flag describing whether all actions in the action batch have completed | [optional] 
**Failed** | Pointer to **bool** | Flag describing whether any actions in the action batch failed | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered when running actions in the action batch | [optional] 
**CreatedResources** | [**[]InlineResponse2015StatusCreatedResources**](InlineResponse2015StatusCreatedResources.md) | Resources created as a result of this action batch | 

## Methods

### NewInlineResponse2015Status

`func NewInlineResponse2015Status(createdResources []InlineResponse2015StatusCreatedResources, ) *InlineResponse2015Status`

NewInlineResponse2015Status instantiates a new InlineResponse2015Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2015StatusWithDefaults

`func NewInlineResponse2015StatusWithDefaults() *InlineResponse2015Status`

NewInlineResponse2015StatusWithDefaults instantiates a new InlineResponse2015Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *InlineResponse2015Status) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse2015Status) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse2015Status) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse2015Status) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetFailed

`func (o *InlineResponse2015Status) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *InlineResponse2015Status) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *InlineResponse2015Status) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *InlineResponse2015Status) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse2015Status) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse2015Status) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse2015Status) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse2015Status) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedResources

`func (o *InlineResponse2015Status) GetCreatedResources() []InlineResponse2015StatusCreatedResources`

GetCreatedResources returns the CreatedResources field if non-nil, zero value otherwise.

### GetCreatedResourcesOk

`func (o *InlineResponse2015Status) GetCreatedResourcesOk() (*[]InlineResponse2015StatusCreatedResources, bool)`

GetCreatedResourcesOk returns a tuple with the CreatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedResources

`func (o *InlineResponse2015Status) SetCreatedResources(v []InlineResponse2015StatusCreatedResources)`

SetCreatedResources sets CreatedResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


