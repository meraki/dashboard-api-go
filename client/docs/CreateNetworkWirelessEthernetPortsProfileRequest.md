# CreateNetworkWirelessEthernetPortsProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | AP port profile name | 
**Ports** | [**[]CreateNetworkWirelessEthernetPortsProfileRequestPortsInner**](CreateNetworkWirelessEthernetPortsProfileRequestPortsInner.md) | AP ports configuration | 
**UsbPorts** | Pointer to [**[]CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner**](CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner.md) | AP usb ports configuration | [optional] 

## Methods

### NewCreateNetworkWirelessEthernetPortsProfileRequest

`func NewCreateNetworkWirelessEthernetPortsProfileRequest(name string, ports []CreateNetworkWirelessEthernetPortsProfileRequestPortsInner, ) *CreateNetworkWirelessEthernetPortsProfileRequest`

NewCreateNetworkWirelessEthernetPortsProfileRequest instantiates a new CreateNetworkWirelessEthernetPortsProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessEthernetPortsProfileRequestWithDefaults

`func NewCreateNetworkWirelessEthernetPortsProfileRequestWithDefaults() *CreateNetworkWirelessEthernetPortsProfileRequest`

NewCreateNetworkWirelessEthernetPortsProfileRequestWithDefaults instantiates a new CreateNetworkWirelessEthernetPortsProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) GetPorts() []CreateNetworkWirelessEthernetPortsProfileRequestPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) GetPortsOk() (*[]CreateNetworkWirelessEthernetPortsProfileRequestPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) SetPorts(v []CreateNetworkWirelessEthernetPortsProfileRequestPortsInner)`

SetPorts sets Ports field to given value.


### GetUsbPorts

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) GetUsbPorts() []CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner`

GetUsbPorts returns the UsbPorts field if non-nil, zero value otherwise.

### GetUsbPortsOk

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) GetUsbPortsOk() (*[]CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner, bool)`

GetUsbPortsOk returns a tuple with the UsbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPorts

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) SetUsbPorts(v []CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner)`

SetUsbPorts sets UsbPorts field to given value.

### HasUsbPorts

`func (o *CreateNetworkWirelessEthernetPortsProfileRequest) HasUsbPorts() bool`

HasUsbPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


