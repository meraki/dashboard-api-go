# InlineObject12

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | FQDN, IPv4 or IPv6 address | 
**Count** | Pointer to **int32** | Count parameter to pass to ping. [1..5], default 5 | [optional] 

## Methods

### NewInlineObject12

`func NewInlineObject12(target string, ) *InlineObject12`

NewInlineObject12 instantiates a new InlineObject12 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject12WithDefaults

`func NewInlineObject12WithDefaults() *InlineObject12`

NewInlineObject12WithDefaults instantiates a new InlineObject12 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *InlineObject12) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineObject12) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineObject12) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetCount

`func (o *InlineObject12) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineObject12) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineObject12) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineObject12) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


