# InlineResponse20017

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The number of the SSID. | [optional] 
**Name** | Pointer to **string** | The name of the SSID. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the SSID is enabled. | [optional] 
**DefaultVlanId** | Pointer to **int32** | The VLAN ID of the VLAN associated to this SSID. | [optional] 
**AuthMode** | Pointer to **string** | The association control method for the SSID. | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdApplianceSsidsRadiusServers**](NetworksNetworkIdApplianceSsidsRadiusServers.md) | The RADIUS 802.1x servers to be used for authentication. | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID. | [optional] 
**WpaEncryptionMode** | Pointer to **string** | WPA encryption mode for the SSID. | [optional] 
**Visible** | Pointer to **bool** | Boolean indicating whether the MX should advertise or hide this SSID. | [optional] 

## Methods

### NewInlineResponse20017

`func NewInlineResponse20017() *InlineResponse20017`

NewInlineResponse20017 instantiates a new InlineResponse20017 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20017WithDefaults

`func NewInlineResponse20017WithDefaults() *InlineResponse20017`

NewInlineResponse20017WithDefaults instantiates a new InlineResponse20017 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse20017) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse20017) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse20017) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse20017) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20017) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20017) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20017) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20017) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20017) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20017) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20017) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20017) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *InlineResponse20017) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *InlineResponse20017) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *InlineResponse20017) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *InlineResponse20017) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetAuthMode

`func (o *InlineResponse20017) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *InlineResponse20017) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *InlineResponse20017) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *InlineResponse20017) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineResponse20017) GetRadiusServers() []NetworksNetworkIdApplianceSsidsRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineResponse20017) GetRadiusServersOk() (*[]NetworksNetworkIdApplianceSsidsRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineResponse20017) SetRadiusServers(v []NetworksNetworkIdApplianceSsidsRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineResponse20017) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *InlineResponse20017) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *InlineResponse20017) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *InlineResponse20017) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *InlineResponse20017) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *InlineResponse20017) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *InlineResponse20017) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *InlineResponse20017) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *InlineResponse20017) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetVisible

`func (o *InlineResponse20017) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InlineResponse20017) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InlineResponse20017) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InlineResponse20017) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


