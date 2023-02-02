# InlineResponse20026

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20026Products**](InlineResponse20026Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20026Stages**](InlineResponse20026Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20025Reasons**](InlineResponse20025Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20026

`func NewInlineResponse20026() *InlineResponse20026`

NewInlineResponse20026 instantiates a new InlineResponse20026 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20026WithDefaults

`func NewInlineResponse20026WithDefaults() *InlineResponse20026`

NewInlineResponse20026WithDefaults instantiates a new InlineResponse20026 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20026) GetProducts() InlineResponse20026Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20026) GetProductsOk() (*InlineResponse20026Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20026) SetProducts(v InlineResponse20026Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20026) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20026) GetStages() []InlineResponse20026Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20026) GetStagesOk() (*[]InlineResponse20026Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20026) SetStages(v []InlineResponse20026Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20026) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20026) GetReasons() []InlineResponse20025Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20026) GetReasonsOk() (*[]InlineResponse20025Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20026) SetReasons(v []InlineResponse20025Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20026) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


