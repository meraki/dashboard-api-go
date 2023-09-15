# UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | AP port name | 
**Enabled** | Pointer to **bool** | AP port enabled | [optional] 
**Ssid** | Pointer to **int32** | AP port ssid number | [optional] 
**PskGroupId** | Pointer to **string** | AP port PSK Group number | [optional] 

## Methods

### NewUpdateNetworkWirelessEthernetPortsProfileRequestPortsInner

`func NewUpdateNetworkWirelessEthernetPortsProfileRequestPortsInner(name string, ) *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner`

NewUpdateNetworkWirelessEthernetPortsProfileRequestPortsInner instantiates a new UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults

`func NewUpdateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults() *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner`

NewUpdateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults instantiates a new UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSsid

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetSsid() int32`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetSsidOk() (*int32, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetSsid(v int32)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetPskGroupId

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetPskGroupId() string`

GetPskGroupId returns the PskGroupId field if non-nil, zero value otherwise.

### GetPskGroupIdOk

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetPskGroupIdOk() (*string, bool)`

GetPskGroupIdOk returns a tuple with the PskGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskGroupId

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetPskGroupId(v string)`

SetPskGroupId sets PskGroupId field to given value.

### HasPskGroupId

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasPskGroupId() bool`

HasPskGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


