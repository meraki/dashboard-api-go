# InlineResponse20016Capabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | model number of appliance | [optional] 
**Bandwidths** | Pointer to [**[]InlineResponse20016Bandwidths**](InlineResponse20016Bandwidths.md) | array of uplink limits | [optional] 

## Methods

### NewInlineResponse20016Capabilities

`func NewInlineResponse20016Capabilities() *InlineResponse20016Capabilities`

NewInlineResponse20016Capabilities instantiates a new InlineResponse20016Capabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20016CapabilitiesWithDefaults

`func NewInlineResponse20016CapabilitiesWithDefaults() *InlineResponse20016Capabilities`

NewInlineResponse20016CapabilitiesWithDefaults instantiates a new InlineResponse20016Capabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *InlineResponse20016Capabilities) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse20016Capabilities) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse20016Capabilities) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse20016Capabilities) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetBandwidths

`func (o *InlineResponse20016Capabilities) GetBandwidths() []InlineResponse20016Bandwidths`

GetBandwidths returns the Bandwidths field if non-nil, zero value otherwise.

### GetBandwidthsOk

`func (o *InlineResponse20016Capabilities) GetBandwidthsOk() (*[]InlineResponse20016Bandwidths, bool)`

GetBandwidthsOk returns a tuple with the Bandwidths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidths

`func (o *InlineResponse20016Capabilities) SetBandwidths(v []InlineResponse20016Bandwidths)`

SetBandwidths sets Bandwidths field to given value.

### HasBandwidths

`func (o *InlineResponse20016Capabilities) HasBandwidths() bool`

HasBandwidths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


