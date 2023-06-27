# GetNetworkApplianceSsids200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The number of the SSID. | [optional] 
**Name** | Pointer to **string** | The name of the SSID. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the SSID is enabled. | [optional] 
**DefaultVlanId** | Pointer to **int32** | The VLAN ID of the VLAN associated to this SSID. | [optional] 
**AuthMode** | Pointer to **string** | The association control method for the SSID. | [optional] 
**RadiusServers** | Pointer to [**[]GetNetworkApplianceSsids200ResponseInnerRadiusServersInner**](GetNetworkApplianceSsids200ResponseInnerRadiusServersInner.md) | The RADIUS 802.1x servers to be used for authentication. | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID. | [optional] 
**WpaEncryptionMode** | Pointer to **string** | WPA encryption mode for the SSID. | [optional] 
**Visible** | Pointer to **bool** | Boolean indicating whether the MX should advertise or hide this SSID. | [optional] 

## Methods

### NewGetNetworkApplianceSsids200ResponseInner

`func NewGetNetworkApplianceSsids200ResponseInner() *GetNetworkApplianceSsids200ResponseInner`

NewGetNetworkApplianceSsids200ResponseInner instantiates a new GetNetworkApplianceSsids200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceSsids200ResponseInnerWithDefaults

`func NewGetNetworkApplianceSsids200ResponseInnerWithDefaults() *GetNetworkApplianceSsids200ResponseInner`

NewGetNetworkApplianceSsids200ResponseInnerWithDefaults instantiates a new GetNetworkApplianceSsids200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GetNetworkApplianceSsids200ResponseInner) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetNetworkApplianceSsids200ResponseInner) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GetNetworkApplianceSsids200ResponseInner) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkApplianceSsids200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkApplianceSsids200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkApplianceSsids200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkApplianceSsids200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkApplianceSsids200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkApplianceSsids200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *GetNetworkApplianceSsids200ResponseInner) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *GetNetworkApplianceSsids200ResponseInner) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *GetNetworkApplianceSsids200ResponseInner) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetAuthMode

`func (o *GetNetworkApplianceSsids200ResponseInner) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *GetNetworkApplianceSsids200ResponseInner) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *GetNetworkApplianceSsids200ResponseInner) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetRadiusServers

`func (o *GetNetworkApplianceSsids200ResponseInner) GetRadiusServers() []GetNetworkApplianceSsids200ResponseInnerRadiusServersInner`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetRadiusServersOk() (*[]GetNetworkApplianceSsids200ResponseInnerRadiusServersInner, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *GetNetworkApplianceSsids200ResponseInner) SetRadiusServers(v []GetNetworkApplianceSsids200ResponseInnerRadiusServersInner)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *GetNetworkApplianceSsids200ResponseInner) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *GetNetworkApplianceSsids200ResponseInner) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *GetNetworkApplianceSsids200ResponseInner) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *GetNetworkApplianceSsids200ResponseInner) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *GetNetworkApplianceSsids200ResponseInner) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *GetNetworkApplianceSsids200ResponseInner) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *GetNetworkApplianceSsids200ResponseInner) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetVisible

`func (o *GetNetworkApplianceSsids200ResponseInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GetNetworkApplianceSsids200ResponseInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GetNetworkApplianceSsids200ResponseInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *GetNetworkApplianceSsids200ResponseInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


