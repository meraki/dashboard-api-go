# InlineObject10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]DevicesSerialCellularSimsSims**](DevicesSerialCellularSimsSims.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimFailover** | Pointer to [**DevicesSerialCellularSimsSimFailover**](DevicesSerialCellularSimsSimFailover.md) |  | [optional] 

## Methods

### NewInlineObject10

`func NewInlineObject10() *InlineObject10`

NewInlineObject10 instantiates a new InlineObject10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject10WithDefaults

`func NewInlineObject10WithDefaults() *InlineObject10`

NewInlineObject10WithDefaults instantiates a new InlineObject10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *InlineObject10) GetSims() []DevicesSerialCellularSimsSims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineObject10) GetSimsOk() (*[]DevicesSerialCellularSimsSims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineObject10) SetSims(v []DevicesSerialCellularSimsSims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineObject10) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimFailover

`func (o *InlineObject10) GetSimFailover() DevicesSerialCellularSimsSimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *InlineObject10) GetSimFailoverOk() (*DevicesSerialCellularSimsSimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *InlineObject10) SetSimFailover(v DevicesSerialCellularSimsSimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *InlineObject10) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


