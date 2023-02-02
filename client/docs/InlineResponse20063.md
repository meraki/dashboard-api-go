# InlineResponse20063

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]InlineResponse20063PowerExceptions**](InlineResponse20063PowerExceptions.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 

## Methods

### NewInlineResponse20063

`func NewInlineResponse20063() *InlineResponse20063`

NewInlineResponse20063 instantiates a new InlineResponse20063 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20063WithDefaults

`func NewInlineResponse20063WithDefaults() *InlineResponse20063`

NewInlineResponse20063WithDefaults instantiates a new InlineResponse20063 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineResponse20063) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20063) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20063) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20063) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *InlineResponse20063) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *InlineResponse20063) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *InlineResponse20063) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *InlineResponse20063) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *InlineResponse20063) GetPowerExceptions() []InlineResponse20063PowerExceptions`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *InlineResponse20063) GetPowerExceptionsOk() (*[]InlineResponse20063PowerExceptions, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *InlineResponse20063) SetPowerExceptions(v []InlineResponse20063PowerExceptions)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *InlineResponse20063) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


