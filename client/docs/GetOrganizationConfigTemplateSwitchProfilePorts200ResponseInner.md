# GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The identifier of the switch profile port. | [optional] 
**Name** | Pointer to **string** | The name of the switch profile port. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the switch profile port. | [optional] 
**Enabled** | Pointer to **bool** | The status of the switch profile port. | [optional] 
**PoeEnabled** | Pointer to **bool** | The PoE status of the switch profile port. | [optional] 
**Type** | Pointer to **string** | The type of the switch profile port (&#39;trunk&#39; or &#39;access&#39;). | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the switch profile port. A null value will clear the value set for trunk ports. | [optional] 
**VoiceVlan** | Pointer to **int32** | The voice VLAN of the switch profile port. Only applicable to access ports. | [optional] 
**AllowedVlans** | Pointer to **string** | The VLANs allowed on the switch profile port. Only applicable to trunk ports. | [optional] 
**IsolationEnabled** | Pointer to **bool** | The isolation status of the switch profile port. | [optional] 
**RstpEnabled** | Pointer to **bool** | The rapid spanning tree protocol status. | [optional] 
**StpGuard** | Pointer to **string** | The state of the STP guard (&#39;disabled&#39;, &#39;root guard&#39;, &#39;bpdu guard&#39; or &#39;loop guard&#39;). | [optional] 
**LinkNegotiation** | Pointer to **string** | The link speed for the switch profile port. | [optional] 
**LinkNegotiationCapabilities** | Pointer to **[]string** | Available link speeds for the switch profile port. | [optional] 
**PortScheduleId** | Pointer to **string** | The ID of the port schedule. A value of null will clear the port schedule. | [optional] 
**Udld** | Pointer to **string** | The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only. | [optional] 
**AccessPolicyType** | Pointer to **string** | The type of the access policy of the switch profile port. Only applicable to access ports. Can be one of &#39;Open&#39;, &#39;Custom access policy&#39;, &#39;MAC allow list&#39; or &#39;Sticky MAC allow list&#39;. | [optional] 
**AccessPolicyNumber** | Pointer to **int32** | The number of a custom access policy to configure on the switch profile port. Only applicable when &#39;accessPolicyType&#39; is &#39;Custom access policy&#39;. | [optional] 
**MacAllowList** | Pointer to **[]string** | Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when &#39;accessPolicyType&#39; is &#39;MAC allow list&#39;. | [optional] 
**StickyMacAllowList** | Pointer to **[]string** | The initial list of MAC addresses for sticky Mac allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StickyMacAllowListLimit** | Pointer to **int32** | The maximum number of MAC addresses for sticky MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StormControlEnabled** | Pointer to **bool** | The storm control status of the switch profile port. | [optional] 
**FlexibleStackingEnabled** | Pointer to **bool** | For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled. | [optional] 
**DaiTrusted** | Pointer to **bool** | If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic. | [optional] 
**Profile** | Pointer to [**GetDeviceSwitchPorts200ResponseInnerProfile**](GetDeviceSwitchPorts200ResponseInnerProfile.md) |  | [optional] 

## Methods

### NewGetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner

`func NewGetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner() *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner`

NewGetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner instantiates a new GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationConfigTemplateSwitchProfilePorts200ResponseInnerWithDefaults

`func NewGetOrganizationConfigTemplateSwitchProfilePorts200ResponseInnerWithDefaults() *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner`

NewGetOrganizationConfigTemplateSwitchProfilePorts200ResponseInnerWithDefaults instantiates a new GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoeEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetPoeEnabled() bool`

GetPoeEnabled returns the PoeEnabled field if non-nil, zero value otherwise.

### GetPoeEnabledOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetPoeEnabledOk() (*bool, bool)`

GetPoeEnabledOk returns a tuple with the PoeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetPoeEnabled(v bool)`

SetPoeEnabled sets PoeEnabled field to given value.

### HasPoeEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasPoeEnabled() bool`

HasPoeEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVoiceVlan

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetIsolationEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetIsolationEnabled() bool`

GetIsolationEnabled returns the IsolationEnabled field if non-nil, zero value otherwise.

### GetIsolationEnabledOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetIsolationEnabledOk() (*bool, bool)`

GetIsolationEnabledOk returns a tuple with the IsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetIsolationEnabled(v bool)`

SetIsolationEnabled sets IsolationEnabled field to given value.

### HasIsolationEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasIsolationEnabled() bool`

HasIsolationEnabled returns a boolean if a field has been set.

### GetRstpEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpGuard

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStpGuard() string`

GetStpGuard returns the StpGuard field if non-nil, zero value otherwise.

### GetStpGuardOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStpGuardOk() (*string, bool)`

GetStpGuardOk returns a tuple with the StpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpGuard

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetStpGuard(v string)`

SetStpGuard sets StpGuard field to given value.

### HasStpGuard

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasStpGuard() bool`

HasStpGuard returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetLinkNegotiationCapabilities

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetLinkNegotiationCapabilities() []string`

GetLinkNegotiationCapabilities returns the LinkNegotiationCapabilities field if non-nil, zero value otherwise.

### GetLinkNegotiationCapabilitiesOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetLinkNegotiationCapabilitiesOk() (*[]string, bool)`

GetLinkNegotiationCapabilitiesOk returns a tuple with the LinkNegotiationCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiationCapabilities

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetLinkNegotiationCapabilities(v []string)`

SetLinkNegotiationCapabilities sets LinkNegotiationCapabilities field to given value.

### HasLinkNegotiationCapabilities

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasLinkNegotiationCapabilities() bool`

HasLinkNegotiationCapabilities returns a boolean if a field has been set.

### GetPortScheduleId

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetPortScheduleId() string`

GetPortScheduleId returns the PortScheduleId field if non-nil, zero value otherwise.

### GetPortScheduleIdOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetPortScheduleIdOk() (*string, bool)`

GetPortScheduleIdOk returns a tuple with the PortScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScheduleId

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetPortScheduleId(v string)`

SetPortScheduleId sets PortScheduleId field to given value.

### HasPortScheduleId

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasPortScheduleId() bool`

HasPortScheduleId returns a boolean if a field has been set.

### GetUdld

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetUdld() string`

GetUdld returns the Udld field if non-nil, zero value otherwise.

### GetUdldOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetUdldOk() (*string, bool)`

GetUdldOk returns a tuple with the Udld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdld

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetUdld(v string)`

SetUdld sets Udld field to given value.

### HasUdld

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasUdld() bool`

HasUdld returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetAccessPolicyNumber

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetAccessPolicyNumber() int32`

GetAccessPolicyNumber returns the AccessPolicyNumber field if non-nil, zero value otherwise.

### GetAccessPolicyNumberOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetAccessPolicyNumberOk() (*int32, bool)`

GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyNumber

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetAccessPolicyNumber(v int32)`

SetAccessPolicyNumber sets AccessPolicyNumber field to given value.

### HasAccessPolicyNumber

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasAccessPolicyNumber() bool`

HasAccessPolicyNumber returns a boolean if a field has been set.

### GetMacAllowList

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetMacAllowList() []string`

GetMacAllowList returns the MacAllowList field if non-nil, zero value otherwise.

### GetMacAllowListOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetMacAllowListOk() (*[]string, bool)`

GetMacAllowListOk returns a tuple with the MacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAllowList

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetMacAllowList(v []string)`

SetMacAllowList sets MacAllowList field to given value.

### HasMacAllowList

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasMacAllowList() bool`

HasMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowList

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStickyMacAllowList() []string`

GetStickyMacAllowList returns the StickyMacAllowList field if non-nil, zero value otherwise.

### GetStickyMacAllowListOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStickyMacAllowListOk() (*[]string, bool)`

GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowList

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetStickyMacAllowList(v []string)`

SetStickyMacAllowList sets StickyMacAllowList field to given value.

### HasStickyMacAllowList

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasStickyMacAllowList() bool`

HasStickyMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowListLimit

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStickyMacAllowListLimit() int32`

GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field if non-nil, zero value otherwise.

### GetStickyMacAllowListLimitOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStickyMacAllowListLimitOk() (*int32, bool)`

GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowListLimit

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetStickyMacAllowListLimit(v int32)`

SetStickyMacAllowListLimit sets StickyMacAllowListLimit field to given value.

### HasStickyMacAllowListLimit

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasStickyMacAllowListLimit() bool`

HasStickyMacAllowListLimit returns a boolean if a field has been set.

### GetStormControlEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStormControlEnabled() bool`

GetStormControlEnabled returns the StormControlEnabled field if non-nil, zero value otherwise.

### GetStormControlEnabledOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetStormControlEnabledOk() (*bool, bool)`

GetStormControlEnabledOk returns a tuple with the StormControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControlEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetStormControlEnabled(v bool)`

SetStormControlEnabled sets StormControlEnabled field to given value.

### HasStormControlEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasStormControlEnabled() bool`

HasStormControlEnabled returns a boolean if a field has been set.

### GetFlexibleStackingEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetFlexibleStackingEnabled() bool`

GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field if non-nil, zero value otherwise.

### GetFlexibleStackingEnabledOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetFlexibleStackingEnabledOk() (*bool, bool)`

GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleStackingEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetFlexibleStackingEnabled(v bool)`

SetFlexibleStackingEnabled sets FlexibleStackingEnabled field to given value.

### HasFlexibleStackingEnabled

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasFlexibleStackingEnabled() bool`

HasFlexibleStackingEnabled returns a boolean if a field has been set.

### GetDaiTrusted

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetDaiTrusted() bool`

GetDaiTrusted returns the DaiTrusted field if non-nil, zero value otherwise.

### GetDaiTrustedOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetDaiTrustedOk() (*bool, bool)`

GetDaiTrustedOk returns a tuple with the DaiTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaiTrusted

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetDaiTrusted(v bool)`

SetDaiTrusted sets DaiTrusted field to given value.

### HasDaiTrusted

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasDaiTrusted() bool`

HasDaiTrusted returns a boolean if a field has been set.

### GetProfile

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetProfile() GetDeviceSwitchPorts200ResponseInnerProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) GetProfileOk() (*GetDeviceSwitchPorts200ResponseInnerProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) SetProfile(v GetDeviceSwitchPorts200ResponseInnerProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


