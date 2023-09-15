# CreateNetworkWirelessEthernetPortsProfileRequestPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | AP port name | 
**Enabled** | Pointer to **bool** | AP port enabled | [optional] 
**Ssid** | Pointer to **int32** | AP port ssid number | [optional] 
**PskGroupId** | Pointer to **string** | AP port PSK Group ID | [optional] 

## Methods

### NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInner

`func NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInner(name string, ) *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner`

NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInner instantiates a new CreateNetworkWirelessEthernetPortsProfileRequestPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults

`func NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults() *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner`

NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults instantiates a new CreateNetworkWirelessEthernetPortsProfileRequestPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSsid

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetSsid() int32`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetSsidOk() (*int32, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetSsid(v int32)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetPskGroupId

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetPskGroupId() string`

GetPskGroupId returns the PskGroupId field if non-nil, zero value otherwise.

### GetPskGroupIdOk

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetPskGroupIdOk() (*string, bool)`

GetPskGroupIdOk returns a tuple with the PskGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskGroupId

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetPskGroupId(v string)`

SetPskGroupId sets PskGroupId field to given value.

### HasPskGroupId

`func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasPskGroupId() bool`

HasPskGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


