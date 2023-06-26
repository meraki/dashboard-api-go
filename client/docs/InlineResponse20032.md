# InlineResponse20032

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20032Products**](InlineResponse20032Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20032Stages**](InlineResponse20032Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20031Reasons**](InlineResponse20031Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20032

`func NewInlineResponse20032() *InlineResponse20032`

NewInlineResponse20032 instantiates a new InlineResponse20032 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20032WithDefaults

`func NewInlineResponse20032WithDefaults() *InlineResponse20032`

NewInlineResponse20032WithDefaults instantiates a new InlineResponse20032 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20032) GetProducts() InlineResponse20032Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20032) GetProductsOk() (*InlineResponse20032Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20032) SetProducts(v InlineResponse20032Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20032) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20032) GetStages() []InlineResponse20032Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20032) GetStagesOk() (*[]InlineResponse20032Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20032) SetStages(v []InlineResponse20032Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20032) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20032) GetReasons() []InlineResponse20031Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20032) GetReasonsOk() (*[]InlineResponse20031Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20032) SetReasons(v []InlineResponse20031Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20032) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


