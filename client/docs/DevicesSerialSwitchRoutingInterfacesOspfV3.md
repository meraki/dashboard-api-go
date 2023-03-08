# DevicesSerialSwitchRoutingInterfacesOspfV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | Area id | [optional] 
**Cost** | Pointer to **int32** | OSPF Cost | [optional] 
**IsPassiveEnabled** | Pointer to **bool** | Disable sending Hello packets on this interface&#39;s IPv6 area | [optional] 

## Methods

### NewDevicesSerialSwitchRoutingInterfacesOspfV3

`func NewDevicesSerialSwitchRoutingInterfacesOspfV3() *DevicesSerialSwitchRoutingInterfacesOspfV3`

NewDevicesSerialSwitchRoutingInterfacesOspfV3 instantiates a new DevicesSerialSwitchRoutingInterfacesOspfV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchRoutingInterfacesOspfV3WithDefaults

`func NewDevicesSerialSwitchRoutingInterfacesOspfV3WithDefaults() *DevicesSerialSwitchRoutingInterfacesOspfV3`

NewDevicesSerialSwitchRoutingInterfacesOspfV3WithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesOspfV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetIsPassiveEnabled

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) GetIsPassiveEnabled() bool`

GetIsPassiveEnabled returns the IsPassiveEnabled field if non-nil, zero value otherwise.

### GetIsPassiveEnabledOk

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) GetIsPassiveEnabledOk() (*bool, bool)`

GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassiveEnabled

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) SetIsPassiveEnabled(v bool)`

SetIsPassiveEnabled sets IsPassiveEnabled field to given value.

### HasIsPassiveEnabled

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV3) HasIsPassiveEnabled() bool`

HasIsPassiveEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


