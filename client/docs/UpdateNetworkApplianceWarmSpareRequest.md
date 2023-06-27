# UpdateNetworkApplianceWarmSpareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enable warm spare | 
**SpareSerial** | Pointer to **string** | Serial number of the warm spare appliance | [optional] 
**UplinkMode** | Pointer to **string** | Uplink mode, either virtual or public | [optional] 
**VirtualIp1** | Pointer to **string** | The WAN 1 shared IP | [optional] 
**VirtualIp2** | Pointer to **string** | The WAN 2 shared IP | [optional] 

## Methods

### NewUpdateNetworkApplianceWarmSpareRequest

`func NewUpdateNetworkApplianceWarmSpareRequest(enabled bool, ) *UpdateNetworkApplianceWarmSpareRequest`

NewUpdateNetworkApplianceWarmSpareRequest instantiates a new UpdateNetworkApplianceWarmSpareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceWarmSpareRequestWithDefaults

`func NewUpdateNetworkApplianceWarmSpareRequestWithDefaults() *UpdateNetworkApplianceWarmSpareRequest`

NewUpdateNetworkApplianceWarmSpareRequestWithDefaults instantiates a new UpdateNetworkApplianceWarmSpareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkApplianceWarmSpareRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSpareSerial

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetSpareSerial() string`

GetSpareSerial returns the SpareSerial field if non-nil, zero value otherwise.

### GetSpareSerialOk

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetSpareSerialOk() (*string, bool)`

GetSpareSerialOk returns a tuple with the SpareSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpareSerial

`func (o *UpdateNetworkApplianceWarmSpareRequest) SetSpareSerial(v string)`

SetSpareSerial sets SpareSerial field to given value.

### HasSpareSerial

`func (o *UpdateNetworkApplianceWarmSpareRequest) HasSpareSerial() bool`

HasSpareSerial returns a boolean if a field has been set.

### GetUplinkMode

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetUplinkMode() string`

GetUplinkMode returns the UplinkMode field if non-nil, zero value otherwise.

### GetUplinkModeOk

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetUplinkModeOk() (*string, bool)`

GetUplinkModeOk returns a tuple with the UplinkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMode

`func (o *UpdateNetworkApplianceWarmSpareRequest) SetUplinkMode(v string)`

SetUplinkMode sets UplinkMode field to given value.

### HasUplinkMode

`func (o *UpdateNetworkApplianceWarmSpareRequest) HasUplinkMode() bool`

HasUplinkMode returns a boolean if a field has been set.

### GetVirtualIp1

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp1() string`

GetVirtualIp1 returns the VirtualIp1 field if non-nil, zero value otherwise.

### GetVirtualIp1Ok

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp1Ok() (*string, bool)`

GetVirtualIp1Ok returns a tuple with the VirtualIp1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp1

`func (o *UpdateNetworkApplianceWarmSpareRequest) SetVirtualIp1(v string)`

SetVirtualIp1 sets VirtualIp1 field to given value.

### HasVirtualIp1

`func (o *UpdateNetworkApplianceWarmSpareRequest) HasVirtualIp1() bool`

HasVirtualIp1 returns a boolean if a field has been set.

### GetVirtualIp2

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp2() string`

GetVirtualIp2 returns the VirtualIp2 field if non-nil, zero value otherwise.

### GetVirtualIp2Ok

`func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp2Ok() (*string, bool)`

GetVirtualIp2Ok returns a tuple with the VirtualIp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp2

`func (o *UpdateNetworkApplianceWarmSpareRequest) SetVirtualIp2(v string)`

SetVirtualIp2 sets VirtualIp2 field to given value.

### HasVirtualIp2

`func (o *UpdateNetworkApplianceWarmSpareRequest) HasVirtualIp2() bool`

HasVirtualIp2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


