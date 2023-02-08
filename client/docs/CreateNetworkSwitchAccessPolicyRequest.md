# CreateNetworkSwitchAccessPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the access policy | 
**RadiusServers** | [**[]CreateNetworkSwitchAccessPolicyRequestRadiusServersInner**](CreateNetworkSwitchAccessPolicyRequestRadiusServersInner.md) | List of RADIUS servers to require connecting devices to authenticate against before granting network access | 
**Radius** | Pointer to [**GetNetworkSwitchAccessPolicies200ResponseInnerRadius**](GetNetworkSwitchAccessPolicies200ResponseInnerRadius.md) |  | [optional] 
**RadiusTestingEnabled** | **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | 
**RadiusCoaSupportEnabled** | **bool** | Change of authentication for RADIUS re-authentication and disconnection | 
**RadiusAccountingEnabled** | **bool** | Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients | 
**RadiusAccountingServers** | Pointer to [**[]CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner**](CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner.md) | List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access | [optional] 
**RadiusGroupAttribute** | Pointer to **string** | Acceptable values are &#x60;\&quot;\&quot;&#x60; for None, or &#x60;\&quot;11\&quot;&#x60; for Group Policies ACL | [optional] 
**HostMode** | **string** | Choose the Host Mode for the access policy. | 
**AccessPolicyType** | Pointer to **string** | Access Type of the policy. Automatically &#39;Hybrid authentication&#39; when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**IncreaseAccessSpeed** | Pointer to **bool** | Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is &#39;Hybrid Authentication. | [optional] 
**GuestVlanId** | Pointer to **int32** | ID for the guest VLAN allow unauthorized devices access to limited network resources | [optional] 
**Dot1x** | Pointer to [**GetNetworkSwitchAccessPolicies200ResponseInnerDot1x**](GetNetworkSwitchAccessPolicies200ResponseInnerDot1x.md) |  | [optional] 
**VoiceVlanClients** | Pointer to **bool** | CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**UrlRedirectWalledGardenEnabled** | **bool** | Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | 
**UrlRedirectWalledGardenRanges** | Pointer to **[]string** | IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 

## Methods

### NewCreateNetworkSwitchAccessPolicyRequest

`func NewCreateNetworkSwitchAccessPolicyRequest(name string, radiusServers []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner, radiusTestingEnabled bool, radiusCoaSupportEnabled bool, radiusAccountingEnabled bool, hostMode string, urlRedirectWalledGardenEnabled bool, ) *CreateNetworkSwitchAccessPolicyRequest`

NewCreateNetworkSwitchAccessPolicyRequest instantiates a new CreateNetworkSwitchAccessPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchAccessPolicyRequestWithDefaults

`func NewCreateNetworkSwitchAccessPolicyRequestWithDefaults() *CreateNetworkSwitchAccessPolicyRequest`

NewCreateNetworkSwitchAccessPolicyRequestWithDefaults instantiates a new CreateNetworkSwitchAccessPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRadiusServers

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusServers() []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusServersOk() (*[]CreateNetworkSwitchAccessPolicyRequestRadiusServersInner, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusServers(v []CreateNetworkSwitchAccessPolicyRequestRadiusServersInner)`

SetRadiusServers sets RadiusServers field to given value.


### GetRadius

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadius() GetNetworkSwitchAccessPolicies200ResponseInnerRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadius(v GetNetworkSwitchAccessPolicies200ResponseInnerRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.


### GetRadiusCoaSupportEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusCoaSupportEnabled() bool`

GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaSupportEnabledOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusCoaSupportEnabledOk() (*bool, bool)`

GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaSupportEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusCoaSupportEnabled(v bool)`

SetRadiusCoaSupportEnabled sets RadiusCoaSupportEnabled field to given value.


### GetRadiusAccountingEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.


### GetRadiusAccountingServers

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingServers() []CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusAccountingServersOk() (*[]CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusAccountingServers(v []CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusGroupAttribute

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusGroupAttribute() string`

GetRadiusGroupAttribute returns the RadiusGroupAttribute field if non-nil, zero value otherwise.

### GetRadiusGroupAttributeOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetRadiusGroupAttributeOk() (*string, bool)`

GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroupAttribute

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetRadiusGroupAttribute(v string)`

SetRadiusGroupAttribute sets RadiusGroupAttribute field to given value.

### HasRadiusGroupAttribute

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasRadiusGroupAttribute() bool`

HasRadiusGroupAttribute returns a boolean if a field has been set.

### GetHostMode

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.


### GetAccessPolicyType

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetIncreaseAccessSpeed

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetIncreaseAccessSpeed() bool`

GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field if non-nil, zero value otherwise.

### GetIncreaseAccessSpeedOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetIncreaseAccessSpeedOk() (*bool, bool)`

GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseAccessSpeed

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetIncreaseAccessSpeed(v bool)`

SetIncreaseAccessSpeed sets IncreaseAccessSpeed field to given value.

### HasIncreaseAccessSpeed

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasIncreaseAccessSpeed() bool`

HasIncreaseAccessSpeed returns a boolean if a field has been set.

### GetGuestVlanId

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetGuestVlanId() int32`

GetGuestVlanId returns the GuestVlanId field if non-nil, zero value otherwise.

### GetGuestVlanIdOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetGuestVlanIdOk() (*int32, bool)`

GetGuestVlanIdOk returns a tuple with the GuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanId

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetGuestVlanId(v int32)`

SetGuestVlanId sets GuestVlanId field to given value.

### HasGuestVlanId

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasGuestVlanId() bool`

HasGuestVlanId returns a boolean if a field has been set.

### GetDot1x

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetDot1x() GetNetworkSwitchAccessPolicies200ResponseInnerDot1x`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetDot1xOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerDot1x, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetDot1x(v GetNetworkSwitchAccessPolicies200ResponseInnerDot1x)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetVoiceVlanClients

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetVoiceVlanClients() bool`

GetVoiceVlanClients returns the VoiceVlanClients field if non-nil, zero value otherwise.

### GetVoiceVlanClientsOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetVoiceVlanClientsOk() (*bool, bool)`

GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanClients

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetVoiceVlanClients(v bool)`

SetVoiceVlanClients sets VoiceVlanClients field to given value.

### HasVoiceVlanClients

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasVoiceVlanClients() bool`

HasVoiceVlanClients returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenEnabled() bool`

GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenEnabledOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool)`

GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenEnabled

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetUrlRedirectWalledGardenEnabled(v bool)`

SetUrlRedirectWalledGardenEnabled sets UrlRedirectWalledGardenEnabled field to given value.


### GetUrlRedirectWalledGardenRanges

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenRanges() []string`

GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenRangesOk

`func (o *CreateNetworkSwitchAccessPolicyRequest) GetUrlRedirectWalledGardenRangesOk() (*[]string, bool)`

GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenRanges

`func (o *CreateNetworkSwitchAccessPolicyRequest) SetUrlRedirectWalledGardenRanges(v []string)`

SetUrlRedirectWalledGardenRanges sets UrlRedirectWalledGardenRanges field to given value.

### HasUrlRedirectWalledGardenRanges

`func (o *CreateNetworkSwitchAccessPolicyRequest) HasUrlRedirectWalledGardenRanges() bool`

HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


