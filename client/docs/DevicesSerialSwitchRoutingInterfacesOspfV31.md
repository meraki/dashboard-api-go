# DevicesSerialSwitchRoutingInterfacesOspfV31

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | The OSPFv3 area to which this interface should belong. Can be either &#39;disabled&#39; or the identifier of an           existing OSPFv3 area. Defaults to &#39;disabled&#39;. | [optional] 
**Cost** | Pointer to **int32** | The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority. | [optional] 
**IsPassiveEnabled** | Pointer to **bool** | When enabled, OSPFv3 will not run on the interface, but the subnet will still be advertised. | [optional] 

## Methods

### NewDevicesSerialSwitchRoutingInterfacesOspfV31

`func NewDevicesSerialSwitchRoutingInterfacesOspfV31() *DevicesSerialSwitchRoutingInterfacesOspfV31`

NewDevicesSerialSwitchRoutingInterfacesOspfV31 instantiates a new DevicesSerialSwitchRoutingInterfacesOspfV31 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchRoutingInterfacesOspfV31WithDefaults

`func NewDevicesSerialSwitchRoutingInterfacesOspfV31WithDefaults() *DevicesSerialSwitchRoutingInterfacesOspfV31`

NewDevicesSerialSwitchRoutingInterfacesOspfV31WithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesOspfV31 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetIsPassiveEnabled

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) GetIsPassiveEnabled() bool`

GetIsPassiveEnabled returns the IsPassiveEnabled field if non-nil, zero value otherwise.

### GetIsPassiveEnabledOk

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) GetIsPassiveEnabledOk() (*bool, bool)`

GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassiveEnabled

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) SetIsPassiveEnabled(v bool)`

SetIsPassiveEnabled sets IsPassiveEnabled field to given value.

### HasIsPassiveEnabled

`func (o *DevicesSerialSwitchRoutingInterfacesOspfV31) HasIsPassiveEnabled() bool`

HasIsPassiveEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


