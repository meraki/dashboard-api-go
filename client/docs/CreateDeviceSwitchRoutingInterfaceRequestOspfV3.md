# CreateDeviceSwitchRoutingInterfaceRequestOspfV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | The OSPFv3 area to which this interface should belong. Can be either &#39;disabled&#39; or the identifier of an           existing OSPFv3 area. Defaults to &#39;disabled&#39;. | [optional] 
**Cost** | Pointer to **int32** | The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority. | [optional] 
**IsPassiveEnabled** | Pointer to **bool** | When enabled, OSPFv3 will not run on the interface, but the subnet will still be advertised. | [optional] 

## Methods

### NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3

`func NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3() *CreateDeviceSwitchRoutingInterfaceRequestOspfV3`

NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3 instantiates a new CreateDeviceSwitchRoutingInterfaceRequestOspfV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3WithDefaults

`func NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3WithDefaults() *CreateDeviceSwitchRoutingInterfaceRequestOspfV3`

NewCreateDeviceSwitchRoutingInterfaceRequestOspfV3WithDefaults instantiates a new CreateDeviceSwitchRoutingInterfaceRequestOspfV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetIsPassiveEnabled

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetIsPassiveEnabled() bool`

GetIsPassiveEnabled returns the IsPassiveEnabled field if non-nil, zero value otherwise.

### GetIsPassiveEnabledOk

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) GetIsPassiveEnabledOk() (*bool, bool)`

GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassiveEnabled

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) SetIsPassiveEnabled(v bool)`

SetIsPassiveEnabled sets IsPassiveEnabled field to given value.

### HasIsPassiveEnabled

`func (o *CreateDeviceSwitchRoutingInterfaceRequestOspfV3) HasIsPassiveEnabled() bool`

HasIsPassiveEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


