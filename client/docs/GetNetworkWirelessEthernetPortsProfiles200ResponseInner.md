# GetNetworkWirelessEthernetPortsProfiles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | AP port profile ID | [optional] 
**Name** | Pointer to **string** | AP port profile name | [optional] 
**IsDefault** | Pointer to **bool** | Is default profile | [optional] 
**Ports** | Pointer to [**[]GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner.md) | Ports config | [optional] 
**UsbPorts** | Pointer to [**[]GetNetworkWirelessEthernetPortsProfiles200ResponseInnerUsbPortsInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInnerUsbPortsInner.md) | Usb ports config | [optional] 

## Methods

### NewGetNetworkWirelessEthernetPortsProfiles200ResponseInner

`func NewGetNetworkWirelessEthernetPortsProfiles200ResponseInner() *GetNetworkWirelessEthernetPortsProfiles200ResponseInner`

NewGetNetworkWirelessEthernetPortsProfiles200ResponseInner instantiates a new GetNetworkWirelessEthernetPortsProfiles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessEthernetPortsProfiles200ResponseInnerWithDefaults

`func NewGetNetworkWirelessEthernetPortsProfiles200ResponseInnerWithDefaults() *GetNetworkWirelessEthernetPortsProfiles200ResponseInner`

NewGetNetworkWirelessEthernetPortsProfiles200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessEthernetPortsProfiles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetPorts

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetPorts() []GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetPortsOk() (*[]GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) SetPorts(v []GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetUsbPorts

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetUsbPorts() []GetNetworkWirelessEthernetPortsProfiles200ResponseInnerUsbPortsInner`

GetUsbPorts returns the UsbPorts field if non-nil, zero value otherwise.

### GetUsbPortsOk

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) GetUsbPortsOk() (*[]GetNetworkWirelessEthernetPortsProfiles200ResponseInnerUsbPortsInner, bool)`

GetUsbPortsOk returns a tuple with the UsbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPorts

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) SetUsbPorts(v []GetNetworkWirelessEthernetPortsProfiles200ResponseInnerUsbPortsInner)`

SetUsbPorts sets UsbPorts field to given value.

### HasUsbPorts

`func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInner) HasUsbPorts() bool`

HasUsbPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


