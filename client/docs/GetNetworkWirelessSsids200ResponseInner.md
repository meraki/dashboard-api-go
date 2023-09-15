# GetNetworkWirelessSsids200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Unique identifier of the SSID | [optional] 
**Name** | Pointer to **string** | The name of the SSID | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the SSID is enabled | [optional] 
**SplashPage** | Pointer to **string** | The type of splash page for the SSID | [optional] 
**SsidAdminAccessible** | Pointer to **bool** | SSID Administrator access status | [optional] 
**LocalAuth** | Pointer to **bool** | Extended local auth flag for Enterprise NAC | [optional] 
**AuthMode** | Pointer to **string** | The association control method for the SSID | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID | [optional] 
**WpaEncryptionMode** | Pointer to **string** | The types of WPA encryption | [optional] 
**RadiusServers** | Pointer to [**[]GetNetworkWirelessSsids200ResponseInnerRadiusServersInner**](GetNetworkWirelessSsids200ResponseInnerRadiusServersInner.md) | List of RADIUS 802.1X servers to be used for authentication | [optional] 
**RadiusAccountingServers** | Pointer to [**[]GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner**](GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner.md) | List of RADIUS accounting 802.1X servers to be used for authentication | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Whether or not RADIUS accounting is enabled | [optional] 
**RadiusEnabled** | Pointer to **bool** | Whether RADIUS authentication is enabled | [optional] 
**RadiusAttributeForGroupPolicies** | Pointer to **string** | RADIUS attribute used to look up group policies | [optional] 
**RadiusFailoverPolicy** | Pointer to **string** | Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable | [optional] 
**RadiusLoadBalancingPolicy** | Pointer to **string** | Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts | [optional] 
**IpAssignmentMode** | Pointer to **string** | The client IP assignment mode | [optional] 
**AdminSplashUrl** | Pointer to **string** | URL for the admin splash page | [optional] 
**SplashTimeout** | Pointer to **string** | Splash page timeout | [optional] 
**WalledGardenEnabled** | Pointer to **bool** | Allow users to access a configurable list of IP ranges prior to sign-on | [optional] 
**WalledGardenRanges** | Pointer to **[]string** | Domain names and IP address ranges available in Walled Garden mode | [optional] 
**MinBitrate** | Pointer to **int32** | The minimum bitrate in Mbps of this SSID in the default indoor RF profile | [optional] 
**BandSelection** | Pointer to **string** | The client-serving radio frequencies of this SSID in the default indoor RF profile | [optional] 
**PerClientBandwidthLimitUp** | Pointer to **int32** | The upload bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**PerClientBandwidthLimitDown** | Pointer to **int32** | The download bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**Visible** | Pointer to **bool** | Whether the SSID is advertised or hidden by the AP | [optional] 
**AvailableOnAllAps** | Pointer to **bool** | Whether all APs broadcast the SSID or if it&#39;s restricted to APs matching any availability tags | [optional] 
**AvailabilityTags** | Pointer to **[]string** | List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list | [optional] 
**PerSsidBandwidthLimitUp** | Pointer to **int32** | The total upload bandwidth limit in Kbps (0 represents no limit) | [optional] 
**PerSsidBandwidthLimitDown** | Pointer to **int32** | The total download bandwidth limit in Kbps (0 represents no limit) | [optional] 
**MandatoryDhcpEnabled** | Pointer to **bool** | Whether clients connecting to this SSID must use the IP address assigned by the DHCP server | [optional] 

## Methods

### NewGetNetworkWirelessSsids200ResponseInner

`func NewGetNetworkWirelessSsids200ResponseInner() *GetNetworkWirelessSsids200ResponseInner`

NewGetNetworkWirelessSsids200ResponseInner instantiates a new GetNetworkWirelessSsids200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsids200ResponseInnerWithDefaults

`func NewGetNetworkWirelessSsids200ResponseInnerWithDefaults() *GetNetworkWirelessSsids200ResponseInner`

NewGetNetworkWirelessSsids200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessSsids200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GetNetworkWirelessSsids200ResponseInner) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetNetworkWirelessSsids200ResponseInner) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GetNetworkWirelessSsids200ResponseInner) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkWirelessSsids200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkWirelessSsids200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkWirelessSsids200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSplashPage

`func (o *GetNetworkWirelessSsids200ResponseInner) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *GetNetworkWirelessSsids200ResponseInner) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *GetNetworkWirelessSsids200ResponseInner) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetSsidAdminAccessible

`func (o *GetNetworkWirelessSsids200ResponseInner) GetSsidAdminAccessible() bool`

GetSsidAdminAccessible returns the SsidAdminAccessible field if non-nil, zero value otherwise.

### GetSsidAdminAccessibleOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetSsidAdminAccessibleOk() (*bool, bool)`

GetSsidAdminAccessibleOk returns a tuple with the SsidAdminAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidAdminAccessible

`func (o *GetNetworkWirelessSsids200ResponseInner) SetSsidAdminAccessible(v bool)`

SetSsidAdminAccessible sets SsidAdminAccessible field to given value.

### HasSsidAdminAccessible

`func (o *GetNetworkWirelessSsids200ResponseInner) HasSsidAdminAccessible() bool`

HasSsidAdminAccessible returns a boolean if a field has been set.

### GetLocalAuth

`func (o *GetNetworkWirelessSsids200ResponseInner) GetLocalAuth() bool`

GetLocalAuth returns the LocalAuth field if non-nil, zero value otherwise.

### GetLocalAuthOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetLocalAuthOk() (*bool, bool)`

GetLocalAuthOk returns a tuple with the LocalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuth

`func (o *GetNetworkWirelessSsids200ResponseInner) SetLocalAuth(v bool)`

SetLocalAuth sets LocalAuth field to given value.

### HasLocalAuth

`func (o *GetNetworkWirelessSsids200ResponseInner) HasLocalAuth() bool`

HasLocalAuth returns a boolean if a field has been set.

### GetAuthMode

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *GetNetworkWirelessSsids200ResponseInner) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *GetNetworkWirelessSsids200ResponseInner) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *GetNetworkWirelessSsids200ResponseInner) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *GetNetworkWirelessSsids200ResponseInner) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *GetNetworkWirelessSsids200ResponseInner) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *GetNetworkWirelessSsids200ResponseInner) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *GetNetworkWirelessSsids200ResponseInner) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *GetNetworkWirelessSsids200ResponseInner) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetRadiusServers

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusServers() []GetNetworkWirelessSsids200ResponseInnerRadiusServersInner`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusServersOk() (*[]GetNetworkWirelessSsids200ResponseInnerRadiusServersInner, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *GetNetworkWirelessSsids200ResponseInner) SetRadiusServers(v []GetNetworkWirelessSsids200ResponseInnerRadiusServersInner)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *GetNetworkWirelessSsids200ResponseInner) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusAccountingServers() []GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusAccountingServersOk() (*[]GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *GetNetworkWirelessSsids200ResponseInner) SetRadiusAccountingServers(v []GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *GetNetworkWirelessSsids200ResponseInner) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusEnabled() bool`

GetRadiusEnabled returns the RadiusEnabled field if non-nil, zero value otherwise.

### GetRadiusEnabledOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusEnabledOk() (*bool, bool)`

GetRadiusEnabledOk returns a tuple with the RadiusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) SetRadiusEnabled(v bool)`

SetRadiusEnabled sets RadiusEnabled field to given value.

### HasRadiusEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) HasRadiusEnabled() bool`

HasRadiusEnabled returns a boolean if a field has been set.

### GetRadiusAttributeForGroupPolicies

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusAttributeForGroupPolicies() string`

GetRadiusAttributeForGroupPolicies returns the RadiusAttributeForGroupPolicies field if non-nil, zero value otherwise.

### GetRadiusAttributeForGroupPoliciesOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusAttributeForGroupPoliciesOk() (*string, bool)`

GetRadiusAttributeForGroupPoliciesOk returns a tuple with the RadiusAttributeForGroupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttributeForGroupPolicies

`func (o *GetNetworkWirelessSsids200ResponseInner) SetRadiusAttributeForGroupPolicies(v string)`

SetRadiusAttributeForGroupPolicies sets RadiusAttributeForGroupPolicies field to given value.

### HasRadiusAttributeForGroupPolicies

`func (o *GetNetworkWirelessSsids200ResponseInner) HasRadiusAttributeForGroupPolicies() bool`

HasRadiusAttributeForGroupPolicies returns a boolean if a field has been set.

### GetRadiusFailoverPolicy

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusFailoverPolicy() string`

GetRadiusFailoverPolicy returns the RadiusFailoverPolicy field if non-nil, zero value otherwise.

### GetRadiusFailoverPolicyOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusFailoverPolicyOk() (*string, bool)`

GetRadiusFailoverPolicyOk returns a tuple with the RadiusFailoverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusFailoverPolicy

`func (o *GetNetworkWirelessSsids200ResponseInner) SetRadiusFailoverPolicy(v string)`

SetRadiusFailoverPolicy sets RadiusFailoverPolicy field to given value.

### HasRadiusFailoverPolicy

`func (o *GetNetworkWirelessSsids200ResponseInner) HasRadiusFailoverPolicy() bool`

HasRadiusFailoverPolicy returns a boolean if a field has been set.

### GetRadiusLoadBalancingPolicy

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusLoadBalancingPolicy() string`

GetRadiusLoadBalancingPolicy returns the RadiusLoadBalancingPolicy field if non-nil, zero value otherwise.

### GetRadiusLoadBalancingPolicyOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetRadiusLoadBalancingPolicyOk() (*string, bool)`

GetRadiusLoadBalancingPolicyOk returns a tuple with the RadiusLoadBalancingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusLoadBalancingPolicy

`func (o *GetNetworkWirelessSsids200ResponseInner) SetRadiusLoadBalancingPolicy(v string)`

SetRadiusLoadBalancingPolicy sets RadiusLoadBalancingPolicy field to given value.

### HasRadiusLoadBalancingPolicy

`func (o *GetNetworkWirelessSsids200ResponseInner) HasRadiusLoadBalancingPolicy() bool`

HasRadiusLoadBalancingPolicy returns a boolean if a field has been set.

### GetIpAssignmentMode

`func (o *GetNetworkWirelessSsids200ResponseInner) GetIpAssignmentMode() string`

GetIpAssignmentMode returns the IpAssignmentMode field if non-nil, zero value otherwise.

### GetIpAssignmentModeOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetIpAssignmentModeOk() (*string, bool)`

GetIpAssignmentModeOk returns a tuple with the IpAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignmentMode

`func (o *GetNetworkWirelessSsids200ResponseInner) SetIpAssignmentMode(v string)`

SetIpAssignmentMode sets IpAssignmentMode field to given value.

### HasIpAssignmentMode

`func (o *GetNetworkWirelessSsids200ResponseInner) HasIpAssignmentMode() bool`

HasIpAssignmentMode returns a boolean if a field has been set.

### GetAdminSplashUrl

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAdminSplashUrl() string`

GetAdminSplashUrl returns the AdminSplashUrl field if non-nil, zero value otherwise.

### GetAdminSplashUrlOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAdminSplashUrlOk() (*string, bool)`

GetAdminSplashUrlOk returns a tuple with the AdminSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSplashUrl

`func (o *GetNetworkWirelessSsids200ResponseInner) SetAdminSplashUrl(v string)`

SetAdminSplashUrl sets AdminSplashUrl field to given value.

### HasAdminSplashUrl

`func (o *GetNetworkWirelessSsids200ResponseInner) HasAdminSplashUrl() bool`

HasAdminSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *GetNetworkWirelessSsids200ResponseInner) GetSplashTimeout() string`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetSplashTimeoutOk() (*string, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *GetNetworkWirelessSsids200ResponseInner) SetSplashTimeout(v string)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *GetNetworkWirelessSsids200ResponseInner) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetWalledGardenEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) GetWalledGardenEnabled() bool`

GetWalledGardenEnabled returns the WalledGardenEnabled field if non-nil, zero value otherwise.

### GetWalledGardenEnabledOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetWalledGardenEnabledOk() (*bool, bool)`

GetWalledGardenEnabledOk returns a tuple with the WalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) SetWalledGardenEnabled(v bool)`

SetWalledGardenEnabled sets WalledGardenEnabled field to given value.

### HasWalledGardenEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) HasWalledGardenEnabled() bool`

HasWalledGardenEnabled returns a boolean if a field has been set.

### GetWalledGardenRanges

`func (o *GetNetworkWirelessSsids200ResponseInner) GetWalledGardenRanges() []string`

GetWalledGardenRanges returns the WalledGardenRanges field if non-nil, zero value otherwise.

### GetWalledGardenRangesOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetWalledGardenRangesOk() (*[]string, bool)`

GetWalledGardenRangesOk returns a tuple with the WalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenRanges

`func (o *GetNetworkWirelessSsids200ResponseInner) SetWalledGardenRanges(v []string)`

SetWalledGardenRanges sets WalledGardenRanges field to given value.

### HasWalledGardenRanges

`func (o *GetNetworkWirelessSsids200ResponseInner) HasWalledGardenRanges() bool`

HasWalledGardenRanges returns a boolean if a field has been set.

### GetMinBitrate

`func (o *GetNetworkWirelessSsids200ResponseInner) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *GetNetworkWirelessSsids200ResponseInner) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *GetNetworkWirelessSsids200ResponseInner) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetBandSelection

`func (o *GetNetworkWirelessSsids200ResponseInner) GetBandSelection() string`

GetBandSelection returns the BandSelection field if non-nil, zero value otherwise.

### GetBandSelectionOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetBandSelectionOk() (*string, bool)`

GetBandSelectionOk returns a tuple with the BandSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelection

`func (o *GetNetworkWirelessSsids200ResponseInner) SetBandSelection(v string)`

SetBandSelection sets BandSelection field to given value.

### HasBandSelection

`func (o *GetNetworkWirelessSsids200ResponseInner) HasBandSelection() bool`

HasBandSelection returns a boolean if a field has been set.

### GetPerClientBandwidthLimitUp

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerClientBandwidthLimitUp() int32`

GetPerClientBandwidthLimitUp returns the PerClientBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitUpOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerClientBandwidthLimitUpOk() (*int32, bool)`

GetPerClientBandwidthLimitUpOk returns a tuple with the PerClientBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitUp

`func (o *GetNetworkWirelessSsids200ResponseInner) SetPerClientBandwidthLimitUp(v int32)`

SetPerClientBandwidthLimitUp sets PerClientBandwidthLimitUp field to given value.

### HasPerClientBandwidthLimitUp

`func (o *GetNetworkWirelessSsids200ResponseInner) HasPerClientBandwidthLimitUp() bool`

HasPerClientBandwidthLimitUp returns a boolean if a field has been set.

### GetPerClientBandwidthLimitDown

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerClientBandwidthLimitDown() int32`

GetPerClientBandwidthLimitDown returns the PerClientBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitDownOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerClientBandwidthLimitDownOk() (*int32, bool)`

GetPerClientBandwidthLimitDownOk returns a tuple with the PerClientBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitDown

`func (o *GetNetworkWirelessSsids200ResponseInner) SetPerClientBandwidthLimitDown(v int32)`

SetPerClientBandwidthLimitDown sets PerClientBandwidthLimitDown field to given value.

### HasPerClientBandwidthLimitDown

`func (o *GetNetworkWirelessSsids200ResponseInner) HasPerClientBandwidthLimitDown() bool`

HasPerClientBandwidthLimitDown returns a boolean if a field has been set.

### GetVisible

`func (o *GetNetworkWirelessSsids200ResponseInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GetNetworkWirelessSsids200ResponseInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *GetNetworkWirelessSsids200ResponseInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetAvailableOnAllAps

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAvailableOnAllAps() bool`

GetAvailableOnAllAps returns the AvailableOnAllAps field if non-nil, zero value otherwise.

### GetAvailableOnAllApsOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAvailableOnAllApsOk() (*bool, bool)`

GetAvailableOnAllApsOk returns a tuple with the AvailableOnAllAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOnAllAps

`func (o *GetNetworkWirelessSsids200ResponseInner) SetAvailableOnAllAps(v bool)`

SetAvailableOnAllAps sets AvailableOnAllAps field to given value.

### HasAvailableOnAllAps

`func (o *GetNetworkWirelessSsids200ResponseInner) HasAvailableOnAllAps() bool`

HasAvailableOnAllAps returns a boolean if a field has been set.

### GetAvailabilityTags

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAvailabilityTags() []string`

GetAvailabilityTags returns the AvailabilityTags field if non-nil, zero value otherwise.

### GetAvailabilityTagsOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetAvailabilityTagsOk() (*[]string, bool)`

GetAvailabilityTagsOk returns a tuple with the AvailabilityTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityTags

`func (o *GetNetworkWirelessSsids200ResponseInner) SetAvailabilityTags(v []string)`

SetAvailabilityTags sets AvailabilityTags field to given value.

### HasAvailabilityTags

`func (o *GetNetworkWirelessSsids200ResponseInner) HasAvailabilityTags() bool`

HasAvailabilityTags returns a boolean if a field has been set.

### GetPerSsidBandwidthLimitUp

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerSsidBandwidthLimitUp() int32`

GetPerSsidBandwidthLimitUp returns the PerSsidBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitUpOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerSsidBandwidthLimitUpOk() (*int32, bool)`

GetPerSsidBandwidthLimitUpOk returns a tuple with the PerSsidBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitUp

`func (o *GetNetworkWirelessSsids200ResponseInner) SetPerSsidBandwidthLimitUp(v int32)`

SetPerSsidBandwidthLimitUp sets PerSsidBandwidthLimitUp field to given value.

### HasPerSsidBandwidthLimitUp

`func (o *GetNetworkWirelessSsids200ResponseInner) HasPerSsidBandwidthLimitUp() bool`

HasPerSsidBandwidthLimitUp returns a boolean if a field has been set.

### GetPerSsidBandwidthLimitDown

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerSsidBandwidthLimitDown() int32`

GetPerSsidBandwidthLimitDown returns the PerSsidBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitDownOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetPerSsidBandwidthLimitDownOk() (*int32, bool)`

GetPerSsidBandwidthLimitDownOk returns a tuple with the PerSsidBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitDown

`func (o *GetNetworkWirelessSsids200ResponseInner) SetPerSsidBandwidthLimitDown(v int32)`

SetPerSsidBandwidthLimitDown sets PerSsidBandwidthLimitDown field to given value.

### HasPerSsidBandwidthLimitDown

`func (o *GetNetworkWirelessSsids200ResponseInner) HasPerSsidBandwidthLimitDown() bool`

HasPerSsidBandwidthLimitDown returns a boolean if a field has been set.

### GetMandatoryDhcpEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) GetMandatoryDhcpEnabled() bool`

GetMandatoryDhcpEnabled returns the MandatoryDhcpEnabled field if non-nil, zero value otherwise.

### GetMandatoryDhcpEnabledOk

`func (o *GetNetworkWirelessSsids200ResponseInner) GetMandatoryDhcpEnabledOk() (*bool, bool)`

GetMandatoryDhcpEnabledOk returns a tuple with the MandatoryDhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcpEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) SetMandatoryDhcpEnabled(v bool)`

SetMandatoryDhcpEnabled sets MandatoryDhcpEnabled field to given value.

### HasMandatoryDhcpEnabled

`func (o *GetNetworkWirelessSsids200ResponseInner) HasMandatoryDhcpEnabled() bool`

HasMandatoryDhcpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


