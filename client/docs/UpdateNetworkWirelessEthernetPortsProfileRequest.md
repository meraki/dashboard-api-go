# UpdateNetworkWirelessEthernetPortsProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | AP port profile name | [optional] 
**Ports** | Pointer to [**[]UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner**](UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner.md) | AP ports configuration | [optional] 
**UsbPorts** | Pointer to [**[]CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner**](CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner.md) | AP usb ports configuration | [optional] 

## Methods

### NewUpdateNetworkWirelessEthernetPortsProfileRequest

`func NewUpdateNetworkWirelessEthernetPortsProfileRequest() *UpdateNetworkWirelessEthernetPortsProfileRequest`

NewUpdateNetworkWirelessEthernetPortsProfileRequest instantiates a new UpdateNetworkWirelessEthernetPortsProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessEthernetPortsProfileRequestWithDefaults

`func NewUpdateNetworkWirelessEthernetPortsProfileRequestWithDefaults() *UpdateNetworkWirelessEthernetPortsProfileRequest`

NewUpdateNetworkWirelessEthernetPortsProfileRequestWithDefaults instantiates a new UpdateNetworkWirelessEthernetPortsProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) GetPorts() []UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) GetPortsOk() (*[]UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) SetPorts(v []UpdateNetworkWirelessEthernetPortsProfileRequestPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetUsbPorts

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) GetUsbPorts() []CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner`

GetUsbPorts returns the UsbPorts field if non-nil, zero value otherwise.

### GetUsbPortsOk

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) GetUsbPortsOk() (*[]CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner, bool)`

GetUsbPortsOk returns a tuple with the UsbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPorts

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) SetUsbPorts(v []CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner)`

SetUsbPorts sets UsbPorts field to given value.

### HasUsbPorts

`func (o *UpdateNetworkWirelessEthernetPortsProfileRequest) HasUsbPorts() bool`

HasUsbPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


