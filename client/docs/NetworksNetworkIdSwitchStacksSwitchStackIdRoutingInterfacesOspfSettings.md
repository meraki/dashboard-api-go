# NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | The OSPF area to which this interface should belong. Can be either &#39;disabled&#39; or the identifier of an existing OSPF area. Defaults to &#39;disabled&#39;. | [optional] 
**Cost** | Pointer to **int32** | The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority. | [optional] 
**IsPassiveEnabled** | Pointer to **bool** | When enabled, OSPF will not run on the interface, but the subnet will still be advertised. | [optional] 

## Methods

### NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings

`func NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings`

NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings instantiates a new NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettingsWithDefaults

`func NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettingsWithDefaults() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings`

NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettingsWithDefaults instantiates a new NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetIsPassiveEnabled

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) GetIsPassiveEnabled() bool`

GetIsPassiveEnabled returns the IsPassiveEnabled field if non-nil, zero value otherwise.

### GetIsPassiveEnabledOk

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) GetIsPassiveEnabledOk() (*bool, bool)`

GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassiveEnabled

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) SetIsPassiveEnabled(v bool)`

SetIsPassiveEnabled sets IsPassiveEnabled field to given value.

### HasIsPassiveEnabled

`func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) HasIsPassiveEnabled() bool`

HasIsPassiveEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


