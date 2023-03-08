# InlineObject127

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default. | [optional] 
**HelloTimerInSeconds** | Pointer to **int32** | Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds. | [optional] 
**DeadTimerInSeconds** | Pointer to **int32** | Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535 | [optional] 
**Areas** | Pointer to [**[]NetworksNetworkIdSwitchRoutingOspfAreas**](NetworksNetworkIdSwitchRoutingOspfAreas.md) | OSPF areas | [optional] 
**V3** | Pointer to [**NetworksNetworkIdSwitchRoutingOspfV3**](NetworksNetworkIdSwitchRoutingOspfV3.md) |  | [optional] 
**Md5AuthenticationEnabled** | Pointer to **bool** | Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default. | [optional] 
**Md5AuthenticationKey** | Pointer to [**NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey**](NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey.md) |  | [optional] 

## Methods

### NewInlineObject127

`func NewInlineObject127() *InlineObject127`

NewInlineObject127 instantiates a new InlineObject127 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject127WithDefaults

`func NewInlineObject127WithDefaults() *InlineObject127`

NewInlineObject127WithDefaults instantiates a new InlineObject127 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject127) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject127) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject127) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject127) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHelloTimerInSeconds

`func (o *InlineObject127) GetHelloTimerInSeconds() int32`

GetHelloTimerInSeconds returns the HelloTimerInSeconds field if non-nil, zero value otherwise.

### GetHelloTimerInSecondsOk

`func (o *InlineObject127) GetHelloTimerInSecondsOk() (*int32, bool)`

GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTimerInSeconds

`func (o *InlineObject127) SetHelloTimerInSeconds(v int32)`

SetHelloTimerInSeconds sets HelloTimerInSeconds field to given value.

### HasHelloTimerInSeconds

`func (o *InlineObject127) HasHelloTimerInSeconds() bool`

HasHelloTimerInSeconds returns a boolean if a field has been set.

### GetDeadTimerInSeconds

`func (o *InlineObject127) GetDeadTimerInSeconds() int32`

GetDeadTimerInSeconds returns the DeadTimerInSeconds field if non-nil, zero value otherwise.

### GetDeadTimerInSecondsOk

`func (o *InlineObject127) GetDeadTimerInSecondsOk() (*int32, bool)`

GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadTimerInSeconds

`func (o *InlineObject127) SetDeadTimerInSeconds(v int32)`

SetDeadTimerInSeconds sets DeadTimerInSeconds field to given value.

### HasDeadTimerInSeconds

`func (o *InlineObject127) HasDeadTimerInSeconds() bool`

HasDeadTimerInSeconds returns a boolean if a field has been set.

### GetAreas

`func (o *InlineObject127) GetAreas() []NetworksNetworkIdSwitchRoutingOspfAreas`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *InlineObject127) GetAreasOk() (*[]NetworksNetworkIdSwitchRoutingOspfAreas, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *InlineObject127) SetAreas(v []NetworksNetworkIdSwitchRoutingOspfAreas)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *InlineObject127) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetV3

`func (o *InlineObject127) GetV3() NetworksNetworkIdSwitchRoutingOspfV3`

GetV3 returns the V3 field if non-nil, zero value otherwise.

### GetV3Ok

`func (o *InlineObject127) GetV3Ok() (*NetworksNetworkIdSwitchRoutingOspfV3, bool)`

GetV3Ok returns a tuple with the V3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3

`func (o *InlineObject127) SetV3(v NetworksNetworkIdSwitchRoutingOspfV3)`

SetV3 sets V3 field to given value.

### HasV3

`func (o *InlineObject127) HasV3() bool`

HasV3 returns a boolean if a field has been set.

### GetMd5AuthenticationEnabled

`func (o *InlineObject127) GetMd5AuthenticationEnabled() bool`

GetMd5AuthenticationEnabled returns the Md5AuthenticationEnabled field if non-nil, zero value otherwise.

### GetMd5AuthenticationEnabledOk

`func (o *InlineObject127) GetMd5AuthenticationEnabledOk() (*bool, bool)`

GetMd5AuthenticationEnabledOk returns a tuple with the Md5AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthenticationEnabled

`func (o *InlineObject127) SetMd5AuthenticationEnabled(v bool)`

SetMd5AuthenticationEnabled sets Md5AuthenticationEnabled field to given value.

### HasMd5AuthenticationEnabled

`func (o *InlineObject127) HasMd5AuthenticationEnabled() bool`

HasMd5AuthenticationEnabled returns a boolean if a field has been set.

### GetMd5AuthenticationKey

`func (o *InlineObject127) GetMd5AuthenticationKey() NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey`

GetMd5AuthenticationKey returns the Md5AuthenticationKey field if non-nil, zero value otherwise.

### GetMd5AuthenticationKeyOk

`func (o *InlineObject127) GetMd5AuthenticationKeyOk() (*NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey, bool)`

GetMd5AuthenticationKeyOk returns a tuple with the Md5AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthenticationKey

`func (o *InlineObject127) SetMd5AuthenticationKey(v NetworksNetworkIdSwitchRoutingOspfMd5AuthenticationKey)`

SetMd5AuthenticationKey sets Md5AuthenticationKey field to given value.

### HasMd5AuthenticationKey

`func (o *InlineObject127) HasMd5AuthenticationKey() bool`

HasMd5AuthenticationKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


