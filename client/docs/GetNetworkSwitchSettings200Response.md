# GetNetworkSwitchSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]GetNetworkSwitchSettings200ResponsePowerExceptionsInner**](GetNetworkSwitchSettings200ResponsePowerExceptionsInner.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 

## Methods

### NewGetNetworkSwitchSettings200Response

`func NewGetNetworkSwitchSettings200Response() *GetNetworkSwitchSettings200Response`

NewGetNetworkSwitchSettings200Response instantiates a new GetNetworkSwitchSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchSettings200ResponseWithDefaults

`func NewGetNetworkSwitchSettings200ResponseWithDefaults() *GetNetworkSwitchSettings200Response`

NewGetNetworkSwitchSettings200ResponseWithDefaults instantiates a new GetNetworkSwitchSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *GetNetworkSwitchSettings200Response) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetNetworkSwitchSettings200Response) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetNetworkSwitchSettings200Response) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetNetworkSwitchSettings200Response) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *GetNetworkSwitchSettings200Response) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *GetNetworkSwitchSettings200Response) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *GetNetworkSwitchSettings200Response) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *GetNetworkSwitchSettings200Response) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *GetNetworkSwitchSettings200Response) GetPowerExceptions() []GetNetworkSwitchSettings200ResponsePowerExceptionsInner`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *GetNetworkSwitchSettings200Response) GetPowerExceptionsOk() (*[]GetNetworkSwitchSettings200ResponsePowerExceptionsInner, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *GetNetworkSwitchSettings200Response) SetPowerExceptions(v []GetNetworkSwitchSettings200ResponsePowerExceptionsInner)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *GetNetworkSwitchSettings200Response) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


