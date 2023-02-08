# GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | Area id | [optional] 
**Cost** | Pointer to **int32** | OSPF Cost | [optional] 
**IsPassiveEnabled** | Pointer to **bool** | Disable sending Hello packets on this interface&#39;s IPv4 area | [optional] 

## Methods

### NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings

`func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings`

NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettingsWithDefaults

`func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettingsWithDefaults() *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings`

NewGetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettingsWithDefaults instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetIsPassiveEnabled

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetIsPassiveEnabled() bool`

GetIsPassiveEnabled returns the IsPassiveEnabled field if non-nil, zero value otherwise.

### GetIsPassiveEnabledOk

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) GetIsPassiveEnabledOk() (*bool, bool)`

GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassiveEnabled

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) SetIsPassiveEnabled(v bool)`

SetIsPassiveEnabled sets IsPassiveEnabled field to given value.

### HasIsPassiveEnabled

`func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) HasIsPassiveEnabled() bool`

HasIsPassiveEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


