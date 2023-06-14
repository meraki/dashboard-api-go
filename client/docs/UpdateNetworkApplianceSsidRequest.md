# UpdateNetworkApplianceSsidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the SSID. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the SSID is enabled. | [optional] 
**DefaultVlanId** | Pointer to **int32** | The VLAN ID of the VLAN associated to this SSID. This parameter is only valid if the network is in routed mode. | [optional] 
**AuthMode** | Pointer to **string** | The association control method for the SSID (&#39;open&#39;, &#39;psk&#39;, &#39;8021x-meraki&#39; or &#39;8021x-radius&#39;). | [optional] 
**Psk** | Pointer to **string** | The passkey for the SSID. This param is only valid if the authMode is &#39;psk&#39;. | [optional] 
**RadiusServers** | Pointer to [**[]UpdateNetworkApplianceSsidRequestRadiusServersInner**](UpdateNetworkApplianceSsidRequestRadiusServersInner.md) | The RADIUS 802.1x servers to be used for authentication. This param is only valid if the authMode is &#39;8021x-radius&#39;. | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID (&#39;wep&#39; or &#39;wpa&#39;). This param is only valid if the authMode is &#39;psk&#39;. | [optional] 
**WpaEncryptionMode** | Pointer to **string** | The types of WPA encryption. (&#39;WPA1 and WPA2&#39;, &#39;WPA2 only&#39;, &#39;WPA3 Transition Mode&#39; or &#39;WPA3 only&#39;). This param is only valid if (1) the authMode is &#39;psk&#39; &amp; the encryptionMode is &#39;wpa&#39; OR (2) the authMode is &#39;8021x-meraki&#39; OR (3) the authMode is &#39;8021x-radius&#39; | [optional] 
**Visible** | Pointer to **bool** | Boolean indicating whether the MX should advertise or hide this SSID. | [optional] 
**DhcpEnforcedDeauthentication** | Pointer to [**UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication**](UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication.md) |  | [optional] 

## Methods

### NewUpdateNetworkApplianceSsidRequest

`func NewUpdateNetworkApplianceSsidRequest() *UpdateNetworkApplianceSsidRequest`

NewUpdateNetworkApplianceSsidRequest instantiates a new UpdateNetworkApplianceSsidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceSsidRequestWithDefaults

`func NewUpdateNetworkApplianceSsidRequestWithDefaults() *UpdateNetworkApplianceSsidRequest`

NewUpdateNetworkApplianceSsidRequestWithDefaults instantiates a new UpdateNetworkApplianceSsidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkApplianceSsidRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkApplianceSsidRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkApplianceSsidRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkApplianceSsidRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateNetworkApplianceSsidRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkApplianceSsidRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkApplianceSsidRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkApplianceSsidRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *UpdateNetworkApplianceSsidRequest) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *UpdateNetworkApplianceSsidRequest) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *UpdateNetworkApplianceSsidRequest) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *UpdateNetworkApplianceSsidRequest) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetAuthMode

`func (o *UpdateNetworkApplianceSsidRequest) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *UpdateNetworkApplianceSsidRequest) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *UpdateNetworkApplianceSsidRequest) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *UpdateNetworkApplianceSsidRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetPsk

`func (o *UpdateNetworkApplianceSsidRequest) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *UpdateNetworkApplianceSsidRequest) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *UpdateNetworkApplianceSsidRequest) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *UpdateNetworkApplianceSsidRequest) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetRadiusServers

`func (o *UpdateNetworkApplianceSsidRequest) GetRadiusServers() []UpdateNetworkApplianceSsidRequestRadiusServersInner`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *UpdateNetworkApplianceSsidRequest) GetRadiusServersOk() (*[]UpdateNetworkApplianceSsidRequestRadiusServersInner, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *UpdateNetworkApplianceSsidRequest) SetRadiusServers(v []UpdateNetworkApplianceSsidRequestRadiusServersInner)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *UpdateNetworkApplianceSsidRequest) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *UpdateNetworkApplianceSsidRequest) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *UpdateNetworkApplianceSsidRequest) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *UpdateNetworkApplianceSsidRequest) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *UpdateNetworkApplianceSsidRequest) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *UpdateNetworkApplianceSsidRequest) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *UpdateNetworkApplianceSsidRequest) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *UpdateNetworkApplianceSsidRequest) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *UpdateNetworkApplianceSsidRequest) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetVisible

`func (o *UpdateNetworkApplianceSsidRequest) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *UpdateNetworkApplianceSsidRequest) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *UpdateNetworkApplianceSsidRequest) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *UpdateNetworkApplianceSsidRequest) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetDhcpEnforcedDeauthentication

`func (o *UpdateNetworkApplianceSsidRequest) GetDhcpEnforcedDeauthentication() UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication`

GetDhcpEnforcedDeauthentication returns the DhcpEnforcedDeauthentication field if non-nil, zero value otherwise.

### GetDhcpEnforcedDeauthenticationOk

`func (o *UpdateNetworkApplianceSsidRequest) GetDhcpEnforcedDeauthenticationOk() (*UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication, bool)`

GetDhcpEnforcedDeauthenticationOk returns a tuple with the DhcpEnforcedDeauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnforcedDeauthentication

`func (o *UpdateNetworkApplianceSsidRequest) SetDhcpEnforcedDeauthentication(v UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication)`

SetDhcpEnforcedDeauthentication sets DhcpEnforcedDeauthentication field to given value.

### HasDhcpEnforcedDeauthentication

`func (o *UpdateNetworkApplianceSsidRequest) HasDhcpEnforcedDeauthentication() bool`

HasDhcpEnforcedDeauthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


