# GetDeviceSwitchPorts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The identifier of the switch port. | [optional] 
**Name** | Pointer to **string** | The name of the switch port. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the switch port. | [optional] 
**Enabled** | Pointer to **bool** | The status of the switch port. | [optional] 
**PoeEnabled** | Pointer to **bool** | The PoE status of the switch port. | [optional] 
**Type** | Pointer to **string** | The type of the switch port (&#39;trunk&#39; or &#39;access&#39;). | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the switch port. A null value will clear the value set for trunk ports. | [optional] 
**VoiceVlan** | Pointer to **int32** | The voice VLAN of the switch port. Only applicable to access ports. | [optional] 
**AllowedVlans** | Pointer to **string** | The VLANs allowed on the switch port. Only applicable to trunk ports. | [optional] 
**IsolationEnabled** | Pointer to **bool** | The isolation status of the switch port. | [optional] 
**RstpEnabled** | Pointer to **bool** | The rapid spanning tree protocol status. | [optional] 
**StpGuard** | Pointer to **string** | The state of the STP guard (&#39;disabled&#39;, &#39;root guard&#39;, &#39;bpdu guard&#39; or &#39;loop guard&#39;). | [optional] 
**LinkNegotiation** | Pointer to **string** | The link speed for the switch port. | [optional] 
**LinkNegotiationCapabilities** | Pointer to **[]string** | Available link speeds for the switch port. | [optional] 
**PortScheduleId** | Pointer to **string** | The ID of the port schedule. A value of null will clear the port schedule. | [optional] 
**Udld** | Pointer to **string** | The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only. | [optional] 
**AccessPolicyType** | Pointer to **string** | The type of the access policy of the switch port. Only applicable to access ports. Can be one of &#39;Open&#39;, &#39;Custom access policy&#39;, &#39;MAC allow list&#39; or &#39;Sticky MAC allow list&#39;. | [optional] 
**AccessPolicyNumber** | Pointer to **int32** | The number of a custom access policy to configure on the switch port. Only applicable when &#39;accessPolicyType&#39; is &#39;Custom access policy&#39;. | [optional] 
**MacAllowList** | Pointer to **[]string** | Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when &#39;accessPolicyType&#39; is &#39;MAC allow list&#39;. | [optional] 
**StickyMacAllowList** | Pointer to **[]string** | The initial list of MAC addresses for sticky Mac allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StickyMacAllowListLimit** | Pointer to **int32** | The maximum number of MAC addresses for sticky MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StormControlEnabled** | Pointer to **bool** | The storm control status of the switch port. | [optional] 
**AdaptivePolicyGroupId** | Pointer to **string** | The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile. | [optional] 
**PeerSgtCapable** | Pointer to **bool** | If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile. | [optional] 
**FlexibleStackingEnabled** | Pointer to **bool** | For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled. | [optional] 
**DaiTrusted** | Pointer to **bool** | If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic. | [optional] 
**Profile** | Pointer to [**GetDeviceSwitchPorts200ResponseInnerProfile**](GetDeviceSwitchPorts200ResponseInnerProfile.md) |  | [optional] 

## Methods

### NewGetDeviceSwitchPorts200ResponseInner

`func NewGetDeviceSwitchPorts200ResponseInner() *GetDeviceSwitchPorts200ResponseInner`

NewGetDeviceSwitchPorts200ResponseInner instantiates a new GetDeviceSwitchPorts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchPorts200ResponseInnerWithDefaults

`func NewGetDeviceSwitchPorts200ResponseInnerWithDefaults() *GetDeviceSwitchPorts200ResponseInner`

NewGetDeviceSwitchPorts200ResponseInnerWithDefaults instantiates a new GetDeviceSwitchPorts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetDeviceSwitchPorts200ResponseInner) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GetDeviceSwitchPorts200ResponseInner) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetName

`func (o *GetDeviceSwitchPorts200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeviceSwitchPorts200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeviceSwitchPorts200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *GetDeviceSwitchPorts200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetDeviceSwitchPorts200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetDeviceSwitchPorts200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoeEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPoeEnabled() bool`

GetPoeEnabled returns the PoeEnabled field if non-nil, zero value otherwise.

### GetPoeEnabledOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPoeEnabledOk() (*bool, bool)`

GetPoeEnabledOk returns a tuple with the PoeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) SetPoeEnabled(v bool)`

SetPoeEnabled sets PoeEnabled field to given value.

### HasPoeEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) HasPoeEnabled() bool`

HasPoeEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetDeviceSwitchPorts200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDeviceSwitchPorts200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDeviceSwitchPorts200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *GetDeviceSwitchPorts200ResponseInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetDeviceSwitchPorts200ResponseInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetDeviceSwitchPorts200ResponseInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVoiceVlan

`func (o *GetDeviceSwitchPorts200ResponseInner) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *GetDeviceSwitchPorts200ResponseInner) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *GetDeviceSwitchPorts200ResponseInner) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *GetDeviceSwitchPorts200ResponseInner) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *GetDeviceSwitchPorts200ResponseInner) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetIsolationEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) GetIsolationEnabled() bool`

GetIsolationEnabled returns the IsolationEnabled field if non-nil, zero value otherwise.

### GetIsolationEnabledOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetIsolationEnabledOk() (*bool, bool)`

GetIsolationEnabledOk returns a tuple with the IsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) SetIsolationEnabled(v bool)`

SetIsolationEnabled sets IsolationEnabled field to given value.

### HasIsolationEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) HasIsolationEnabled() bool`

HasIsolationEnabled returns a boolean if a field has been set.

### GetRstpEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpGuard

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStpGuard() string`

GetStpGuard returns the StpGuard field if non-nil, zero value otherwise.

### GetStpGuardOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStpGuardOk() (*string, bool)`

GetStpGuardOk returns a tuple with the StpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpGuard

`func (o *GetDeviceSwitchPorts200ResponseInner) SetStpGuard(v string)`

SetStpGuard sets StpGuard field to given value.

### HasStpGuard

`func (o *GetDeviceSwitchPorts200ResponseInner) HasStpGuard() bool`

HasStpGuard returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *GetDeviceSwitchPorts200ResponseInner) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *GetDeviceSwitchPorts200ResponseInner) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *GetDeviceSwitchPorts200ResponseInner) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetLinkNegotiationCapabilities

`func (o *GetDeviceSwitchPorts200ResponseInner) GetLinkNegotiationCapabilities() []string`

GetLinkNegotiationCapabilities returns the LinkNegotiationCapabilities field if non-nil, zero value otherwise.

### GetLinkNegotiationCapabilitiesOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetLinkNegotiationCapabilitiesOk() (*[]string, bool)`

GetLinkNegotiationCapabilitiesOk returns a tuple with the LinkNegotiationCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiationCapabilities

`func (o *GetDeviceSwitchPorts200ResponseInner) SetLinkNegotiationCapabilities(v []string)`

SetLinkNegotiationCapabilities sets LinkNegotiationCapabilities field to given value.

### HasLinkNegotiationCapabilities

`func (o *GetDeviceSwitchPorts200ResponseInner) HasLinkNegotiationCapabilities() bool`

HasLinkNegotiationCapabilities returns a boolean if a field has been set.

### GetPortScheduleId

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPortScheduleId() string`

GetPortScheduleId returns the PortScheduleId field if non-nil, zero value otherwise.

### GetPortScheduleIdOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPortScheduleIdOk() (*string, bool)`

GetPortScheduleIdOk returns a tuple with the PortScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScheduleId

`func (o *GetDeviceSwitchPorts200ResponseInner) SetPortScheduleId(v string)`

SetPortScheduleId sets PortScheduleId field to given value.

### HasPortScheduleId

`func (o *GetDeviceSwitchPorts200ResponseInner) HasPortScheduleId() bool`

HasPortScheduleId returns a boolean if a field has been set.

### GetUdld

`func (o *GetDeviceSwitchPorts200ResponseInner) GetUdld() string`

GetUdld returns the Udld field if non-nil, zero value otherwise.

### GetUdldOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetUdldOk() (*string, bool)`

GetUdldOk returns a tuple with the Udld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdld

`func (o *GetDeviceSwitchPorts200ResponseInner) SetUdld(v string)`

SetUdld sets Udld field to given value.

### HasUdld

`func (o *GetDeviceSwitchPorts200ResponseInner) HasUdld() bool`

HasUdld returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *GetDeviceSwitchPorts200ResponseInner) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *GetDeviceSwitchPorts200ResponseInner) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetAccessPolicyNumber

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAccessPolicyNumber() int32`

GetAccessPolicyNumber returns the AccessPolicyNumber field if non-nil, zero value otherwise.

### GetAccessPolicyNumberOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAccessPolicyNumberOk() (*int32, bool)`

GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyNumber

`func (o *GetDeviceSwitchPorts200ResponseInner) SetAccessPolicyNumber(v int32)`

SetAccessPolicyNumber sets AccessPolicyNumber field to given value.

### HasAccessPolicyNumber

`func (o *GetDeviceSwitchPorts200ResponseInner) HasAccessPolicyNumber() bool`

HasAccessPolicyNumber returns a boolean if a field has been set.

### GetMacAllowList

`func (o *GetDeviceSwitchPorts200ResponseInner) GetMacAllowList() []string`

GetMacAllowList returns the MacAllowList field if non-nil, zero value otherwise.

### GetMacAllowListOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetMacAllowListOk() (*[]string, bool)`

GetMacAllowListOk returns a tuple with the MacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAllowList

`func (o *GetDeviceSwitchPorts200ResponseInner) SetMacAllowList(v []string)`

SetMacAllowList sets MacAllowList field to given value.

### HasMacAllowList

`func (o *GetDeviceSwitchPorts200ResponseInner) HasMacAllowList() bool`

HasMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowList

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStickyMacAllowList() []string`

GetStickyMacAllowList returns the StickyMacAllowList field if non-nil, zero value otherwise.

### GetStickyMacAllowListOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStickyMacAllowListOk() (*[]string, bool)`

GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowList

`func (o *GetDeviceSwitchPorts200ResponseInner) SetStickyMacAllowList(v []string)`

SetStickyMacAllowList sets StickyMacAllowList field to given value.

### HasStickyMacAllowList

`func (o *GetDeviceSwitchPorts200ResponseInner) HasStickyMacAllowList() bool`

HasStickyMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowListLimit

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStickyMacAllowListLimit() int32`

GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field if non-nil, zero value otherwise.

### GetStickyMacAllowListLimitOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStickyMacAllowListLimitOk() (*int32, bool)`

GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowListLimit

`func (o *GetDeviceSwitchPorts200ResponseInner) SetStickyMacAllowListLimit(v int32)`

SetStickyMacAllowListLimit sets StickyMacAllowListLimit field to given value.

### HasStickyMacAllowListLimit

`func (o *GetDeviceSwitchPorts200ResponseInner) HasStickyMacAllowListLimit() bool`

HasStickyMacAllowListLimit returns a boolean if a field has been set.

### GetStormControlEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStormControlEnabled() bool`

GetStormControlEnabled returns the StormControlEnabled field if non-nil, zero value otherwise.

### GetStormControlEnabledOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetStormControlEnabledOk() (*bool, bool)`

GetStormControlEnabledOk returns a tuple with the StormControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControlEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) SetStormControlEnabled(v bool)`

SetStormControlEnabled sets StormControlEnabled field to given value.

### HasStormControlEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) HasStormControlEnabled() bool`

HasStormControlEnabled returns a boolean if a field has been set.

### GetAdaptivePolicyGroupId

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAdaptivePolicyGroupId() string`

GetAdaptivePolicyGroupId returns the AdaptivePolicyGroupId field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupIdOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetAdaptivePolicyGroupIdOk() (*string, bool)`

GetAdaptivePolicyGroupIdOk returns a tuple with the AdaptivePolicyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroupId

`func (o *GetDeviceSwitchPorts200ResponseInner) SetAdaptivePolicyGroupId(v string)`

SetAdaptivePolicyGroupId sets AdaptivePolicyGroupId field to given value.

### HasAdaptivePolicyGroupId

`func (o *GetDeviceSwitchPorts200ResponseInner) HasAdaptivePolicyGroupId() bool`

HasAdaptivePolicyGroupId returns a boolean if a field has been set.

### GetPeerSgtCapable

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPeerSgtCapable() bool`

GetPeerSgtCapable returns the PeerSgtCapable field if non-nil, zero value otherwise.

### GetPeerSgtCapableOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetPeerSgtCapableOk() (*bool, bool)`

GetPeerSgtCapableOk returns a tuple with the PeerSgtCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSgtCapable

`func (o *GetDeviceSwitchPorts200ResponseInner) SetPeerSgtCapable(v bool)`

SetPeerSgtCapable sets PeerSgtCapable field to given value.

### HasPeerSgtCapable

`func (o *GetDeviceSwitchPorts200ResponseInner) HasPeerSgtCapable() bool`

HasPeerSgtCapable returns a boolean if a field has been set.

### GetFlexibleStackingEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) GetFlexibleStackingEnabled() bool`

GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field if non-nil, zero value otherwise.

### GetFlexibleStackingEnabledOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetFlexibleStackingEnabledOk() (*bool, bool)`

GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleStackingEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) SetFlexibleStackingEnabled(v bool)`

SetFlexibleStackingEnabled sets FlexibleStackingEnabled field to given value.

### HasFlexibleStackingEnabled

`func (o *GetDeviceSwitchPorts200ResponseInner) HasFlexibleStackingEnabled() bool`

HasFlexibleStackingEnabled returns a boolean if a field has been set.

### GetDaiTrusted

`func (o *GetDeviceSwitchPorts200ResponseInner) GetDaiTrusted() bool`

GetDaiTrusted returns the DaiTrusted field if non-nil, zero value otherwise.

### GetDaiTrustedOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetDaiTrustedOk() (*bool, bool)`

GetDaiTrustedOk returns a tuple with the DaiTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaiTrusted

`func (o *GetDeviceSwitchPorts200ResponseInner) SetDaiTrusted(v bool)`

SetDaiTrusted sets DaiTrusted field to given value.

### HasDaiTrusted

`func (o *GetDeviceSwitchPorts200ResponseInner) HasDaiTrusted() bool`

HasDaiTrusted returns a boolean if a field has been set.

### GetProfile

`func (o *GetDeviceSwitchPorts200ResponseInner) GetProfile() GetDeviceSwitchPorts200ResponseInnerProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetDeviceSwitchPorts200ResponseInner) GetProfileOk() (*GetDeviceSwitchPorts200ResponseInnerProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetDeviceSwitchPorts200ResponseInner) SetProfile(v GetDeviceSwitchPorts200ResponseInnerProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetDeviceSwitchPorts200ResponseInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


