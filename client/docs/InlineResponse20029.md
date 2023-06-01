# InlineResponse20029

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20029Products**](InlineResponse20029Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20029Stages**](InlineResponse20029Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20028Reasons**](InlineResponse20028Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20029

`func NewInlineResponse20029() *InlineResponse20029`

NewInlineResponse20029 instantiates a new InlineResponse20029 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029WithDefaults

`func NewInlineResponse20029WithDefaults() *InlineResponse20029`

NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20029) GetProducts() InlineResponse20029Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20029) GetProductsOk() (*InlineResponse20029Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20029) SetProducts(v InlineResponse20029Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20029) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20029) GetStages() []InlineResponse20029Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20029) GetStagesOk() (*[]InlineResponse20029Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20029) SetStages(v []InlineResponse20029Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20029) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20029) GetReasons() []InlineResponse20028Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20029) GetReasonsOk() (*[]InlineResponse20028Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20029) SetReasons(v []InlineResponse20028Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20029) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


