# CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | The OSPF area to which this interface should belong. Can be either &#39;disabled&#39; or the identifier of an existing OSPF area. Defaults to &#39;disabled&#39;. | [optional] 
**Cost** | Pointer to **int32** | The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority. | [optional] 
**IsPassiveEnabled** | Pointer to **bool** | When enabled, OSPF will not run on the interface, but the subnet will still be advertised. | [optional] 

## Methods

### NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings

`func NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings() *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings`

NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings instantiates a new CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettingsWithDefaults

`func NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettingsWithDefaults() *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings`

NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettingsWithDefaults instantiates a new CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetIsPassiveEnabled

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetIsPassiveEnabled() bool`

GetIsPassiveEnabled returns the IsPassiveEnabled field if non-nil, zero value otherwise.

### GetIsPassiveEnabledOk

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetIsPassiveEnabledOk() (*bool, bool)`

GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassiveEnabled

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) SetIsPassiveEnabled(v bool)`

SetIsPassiveEnabled sets IsPassiveEnabled field to given value.

### HasIsPassiveEnabled

`func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) HasIsPassiveEnabled() bool`

HasIsPassiveEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


