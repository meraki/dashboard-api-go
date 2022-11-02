# GetNetworkSwitchAccessPolicies200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the access policy | [optional] 
**RadiusServers** | Pointer to [**[]GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner**](GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner.md) | List of RADIUS servers to require connecting devices to authenticate against before granting network access | [optional] 
**Radius** | Pointer to [**GetNetworkSwitchAccessPolicies200ResponseInnerRadius**](GetNetworkSwitchAccessPolicies200ResponseInnerRadius.md) |  | [optional] 
**RadiusTestingEnabled** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusCoaSupportEnabled** | Pointer to **bool** | Change of authentication for RADIUS re-authentication and disconnection | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients | [optional] 
**RadiusAccountingServers** | Pointer to [**[]GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner**](GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner.md) | List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access | [optional] 
**RadiusGroupAttribute** | Pointer to **string** | Acceptable values are &#x60;\&quot;\&quot;&#x60; for None, or &#x60;\&quot;11\&quot;&#x60; for Group Policies ACL | [optional] 
**HostMode** | Pointer to **string** | Choose the Host Mode for the access policy. | [optional] 
**AccessPolicyType** | Pointer to **string** | Access Type of the policy. Automatically &#39;Hybrid authentication&#39; when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**IncreaseAccessSpeed** | Pointer to **bool** | Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is &#39;Hybrid Authentication. | [optional] 
**GuestVlanId** | Pointer to **int32** | ID for the guest VLAN allow unauthorized devices access to limited network resources | [optional] 
**Dot1x** | Pointer to [**GetNetworkSwitchAccessPolicies200ResponseInnerDot1x**](GetNetworkSwitchAccessPolicies200ResponseInnerDot1x.md) |  | [optional] 
**VoiceVlanClients** | Pointer to **bool** | CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**UrlRedirectWalledGardenEnabled** | Pointer to **bool** | Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication | [optional] 
**UrlRedirectWalledGardenRanges** | Pointer to **[]string** | IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 

## Methods

### NewGetNetworkSwitchAccessPolicies200ResponseInner

`func NewGetNetworkSwitchAccessPolicies200ResponseInner() *GetNetworkSwitchAccessPolicies200ResponseInner`

NewGetNetworkSwitchAccessPolicies200ResponseInner instantiates a new GetNetworkSwitchAccessPolicies200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchAccessPolicies200ResponseInnerWithDefaults

`func NewGetNetworkSwitchAccessPolicies200ResponseInnerWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInner`

NewGetNetworkSwitchAccessPolicies200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRadiusServers

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusServers() []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusServersOk() (*[]GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusServers(v []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadius

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadius() GetNetworkSwitchAccessPolicies200ResponseInnerRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadius(v GetNetworkSwitchAccessPolicies200ResponseInnerRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.

### HasRadiusTestingEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusTestingEnabled() bool`

HasRadiusTestingEnabled returns a boolean if a field has been set.

### GetRadiusCoaSupportEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusCoaSupportEnabled() bool`

GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaSupportEnabledOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusCoaSupportEnabledOk() (*bool, bool)`

GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaSupportEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusCoaSupportEnabled(v bool)`

SetRadiusCoaSupportEnabled sets RadiusCoaSupportEnabled field to given value.

### HasRadiusCoaSupportEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusCoaSupportEnabled() bool`

HasRadiusCoaSupportEnabled returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingServers() []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingServersOk() (*[]GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusAccountingServers(v []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusGroupAttribute

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusGroupAttribute() string`

GetRadiusGroupAttribute returns the RadiusGroupAttribute field if non-nil, zero value otherwise.

### GetRadiusGroupAttributeOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusGroupAttributeOk() (*string, bool)`

GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroupAttribute

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusGroupAttribute(v string)`

SetRadiusGroupAttribute sets RadiusGroupAttribute field to given value.

### HasRadiusGroupAttribute

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusGroupAttribute() bool`

HasRadiusGroupAttribute returns a boolean if a field has been set.

### GetHostMode

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.

### HasHostMode

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasHostMode() bool`

HasHostMode returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetIncreaseAccessSpeed

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetIncreaseAccessSpeed() bool`

GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field if non-nil, zero value otherwise.

### GetIncreaseAccessSpeedOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetIncreaseAccessSpeedOk() (*bool, bool)`

GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseAccessSpeed

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetIncreaseAccessSpeed(v bool)`

SetIncreaseAccessSpeed sets IncreaseAccessSpeed field to given value.

### HasIncreaseAccessSpeed

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasIncreaseAccessSpeed() bool`

HasIncreaseAccessSpeed returns a boolean if a field has been set.

### GetGuestVlanId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetGuestVlanId() int32`

GetGuestVlanId returns the GuestVlanId field if non-nil, zero value otherwise.

### GetGuestVlanIdOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetGuestVlanIdOk() (*int32, bool)`

GetGuestVlanIdOk returns a tuple with the GuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetGuestVlanId(v int32)`

SetGuestVlanId sets GuestVlanId field to given value.

### HasGuestVlanId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasGuestVlanId() bool`

HasGuestVlanId returns a boolean if a field has been set.

### GetDot1x

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetDot1x() GetNetworkSwitchAccessPolicies200ResponseInnerDot1x`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetDot1xOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerDot1x, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetDot1x(v GetNetworkSwitchAccessPolicies200ResponseInnerDot1x)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetVoiceVlanClients

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetVoiceVlanClients() bool`

GetVoiceVlanClients returns the VoiceVlanClients field if non-nil, zero value otherwise.

### GetVoiceVlanClientsOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetVoiceVlanClientsOk() (*bool, bool)`

GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanClients

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetVoiceVlanClients(v bool)`

SetVoiceVlanClients sets VoiceVlanClients field to given value.

### HasVoiceVlanClients

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasVoiceVlanClients() bool`

HasVoiceVlanClients returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenEnabled() bool`

GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenEnabledOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool)`

GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetUrlRedirectWalledGardenEnabled(v bool)`

SetUrlRedirectWalledGardenEnabled sets UrlRedirectWalledGardenEnabled field to given value.

### HasUrlRedirectWalledGardenEnabled

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasUrlRedirectWalledGardenEnabled() bool`

HasUrlRedirectWalledGardenEnabled returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenRanges

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenRanges() []string`

GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenRangesOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenRangesOk() (*[]string, bool)`

GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenRanges

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetUrlRedirectWalledGardenRanges(v []string)`

SetUrlRedirectWalledGardenRanges sets UrlRedirectWalledGardenRanges field to given value.

### HasUrlRedirectWalledGardenRanges

`func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasUrlRedirectWalledGardenRanges() bool`

HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


