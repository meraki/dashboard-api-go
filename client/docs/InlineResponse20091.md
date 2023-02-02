# InlineResponse20091

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
**Profile** | Pointer to [**DevicesSerialSwitchPortsProfile**](DevicesSerialSwitchPortsProfile.md) |  | [optional] 

## Methods

### NewInlineResponse20091

`func NewInlineResponse20091() *InlineResponse20091`

NewInlineResponse20091 instantiates a new InlineResponse20091 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20091WithDefaults

`func NewInlineResponse20091WithDefaults() *InlineResponse20091`

NewInlineResponse20091WithDefaults instantiates a new InlineResponse20091 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse20091) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20091) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20091) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20091) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20091) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20091) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20091) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20091) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20091) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20091) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20091) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20091) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20091) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20091) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20091) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20091) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoeEnabled

`func (o *InlineResponse20091) GetPoeEnabled() bool`

GetPoeEnabled returns the PoeEnabled field if non-nil, zero value otherwise.

### GetPoeEnabledOk

`func (o *InlineResponse20091) GetPoeEnabledOk() (*bool, bool)`

GetPoeEnabledOk returns a tuple with the PoeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeEnabled

`func (o *InlineResponse20091) SetPoeEnabled(v bool)`

SetPoeEnabled sets PoeEnabled field to given value.

### HasPoeEnabled

`func (o *InlineResponse20091) HasPoeEnabled() bool`

HasPoeEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20091) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20091) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20091) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20091) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20091) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20091) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20091) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20091) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVoiceVlan

`func (o *InlineResponse20091) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *InlineResponse20091) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *InlineResponse20091) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *InlineResponse20091) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *InlineResponse20091) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *InlineResponse20091) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *InlineResponse20091) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *InlineResponse20091) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetIsolationEnabled

`func (o *InlineResponse20091) GetIsolationEnabled() bool`

GetIsolationEnabled returns the IsolationEnabled field if non-nil, zero value otherwise.

### GetIsolationEnabledOk

`func (o *InlineResponse20091) GetIsolationEnabledOk() (*bool, bool)`

GetIsolationEnabledOk returns a tuple with the IsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationEnabled

`func (o *InlineResponse20091) SetIsolationEnabled(v bool)`

SetIsolationEnabled sets IsolationEnabled field to given value.

### HasIsolationEnabled

`func (o *InlineResponse20091) HasIsolationEnabled() bool`

HasIsolationEnabled returns a boolean if a field has been set.

### GetRstpEnabled

`func (o *InlineResponse20091) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineResponse20091) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineResponse20091) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineResponse20091) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpGuard

`func (o *InlineResponse20091) GetStpGuard() string`

GetStpGuard returns the StpGuard field if non-nil, zero value otherwise.

### GetStpGuardOk

`func (o *InlineResponse20091) GetStpGuardOk() (*string, bool)`

GetStpGuardOk returns a tuple with the StpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpGuard

`func (o *InlineResponse20091) SetStpGuard(v string)`

SetStpGuard sets StpGuard field to given value.

### HasStpGuard

`func (o *InlineResponse20091) HasStpGuard() bool`

HasStpGuard returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineResponse20091) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineResponse20091) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineResponse20091) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineResponse20091) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetLinkNegotiationCapabilities

`func (o *InlineResponse20091) GetLinkNegotiationCapabilities() []string`

GetLinkNegotiationCapabilities returns the LinkNegotiationCapabilities field if non-nil, zero value otherwise.

### GetLinkNegotiationCapabilitiesOk

`func (o *InlineResponse20091) GetLinkNegotiationCapabilitiesOk() (*[]string, bool)`

GetLinkNegotiationCapabilitiesOk returns a tuple with the LinkNegotiationCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiationCapabilities

`func (o *InlineResponse20091) SetLinkNegotiationCapabilities(v []string)`

SetLinkNegotiationCapabilities sets LinkNegotiationCapabilities field to given value.

### HasLinkNegotiationCapabilities

`func (o *InlineResponse20091) HasLinkNegotiationCapabilities() bool`

HasLinkNegotiationCapabilities returns a boolean if a field has been set.

### GetPortScheduleId

`func (o *InlineResponse20091) GetPortScheduleId() string`

GetPortScheduleId returns the PortScheduleId field if non-nil, zero value otherwise.

### GetPortScheduleIdOk

`func (o *InlineResponse20091) GetPortScheduleIdOk() (*string, bool)`

GetPortScheduleIdOk returns a tuple with the PortScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScheduleId

`func (o *InlineResponse20091) SetPortScheduleId(v string)`

SetPortScheduleId sets PortScheduleId field to given value.

### HasPortScheduleId

`func (o *InlineResponse20091) HasPortScheduleId() bool`

HasPortScheduleId returns a boolean if a field has been set.

### GetUdld

`func (o *InlineResponse20091) GetUdld() string`

GetUdld returns the Udld field if non-nil, zero value otherwise.

### GetUdldOk

`func (o *InlineResponse20091) GetUdldOk() (*string, bool)`

GetUdldOk returns a tuple with the Udld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdld

`func (o *InlineResponse20091) SetUdld(v string)`

SetUdld sets Udld field to given value.

### HasUdld

`func (o *InlineResponse20091) HasUdld() bool`

HasUdld returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *InlineResponse20091) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineResponse20091) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineResponse20091) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineResponse20091) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetAccessPolicyNumber

`func (o *InlineResponse20091) GetAccessPolicyNumber() int32`

GetAccessPolicyNumber returns the AccessPolicyNumber field if non-nil, zero value otherwise.

### GetAccessPolicyNumberOk

`func (o *InlineResponse20091) GetAccessPolicyNumberOk() (*int32, bool)`

GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyNumber

`func (o *InlineResponse20091) SetAccessPolicyNumber(v int32)`

SetAccessPolicyNumber sets AccessPolicyNumber field to given value.

### HasAccessPolicyNumber

`func (o *InlineResponse20091) HasAccessPolicyNumber() bool`

HasAccessPolicyNumber returns a boolean if a field has been set.

### GetMacAllowList

`func (o *InlineResponse20091) GetMacAllowList() []string`

GetMacAllowList returns the MacAllowList field if non-nil, zero value otherwise.

### GetMacAllowListOk

`func (o *InlineResponse20091) GetMacAllowListOk() (*[]string, bool)`

GetMacAllowListOk returns a tuple with the MacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAllowList

`func (o *InlineResponse20091) SetMacAllowList(v []string)`

SetMacAllowList sets MacAllowList field to given value.

### HasMacAllowList

`func (o *InlineResponse20091) HasMacAllowList() bool`

HasMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowList

`func (o *InlineResponse20091) GetStickyMacAllowList() []string`

GetStickyMacAllowList returns the StickyMacAllowList field if non-nil, zero value otherwise.

### GetStickyMacAllowListOk

`func (o *InlineResponse20091) GetStickyMacAllowListOk() (*[]string, bool)`

GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowList

`func (o *InlineResponse20091) SetStickyMacAllowList(v []string)`

SetStickyMacAllowList sets StickyMacAllowList field to given value.

### HasStickyMacAllowList

`func (o *InlineResponse20091) HasStickyMacAllowList() bool`

HasStickyMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowListLimit

`func (o *InlineResponse20091) GetStickyMacAllowListLimit() int32`

GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field if non-nil, zero value otherwise.

### GetStickyMacAllowListLimitOk

`func (o *InlineResponse20091) GetStickyMacAllowListLimitOk() (*int32, bool)`

GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowListLimit

`func (o *InlineResponse20091) SetStickyMacAllowListLimit(v int32)`

SetStickyMacAllowListLimit sets StickyMacAllowListLimit field to given value.

### HasStickyMacAllowListLimit

`func (o *InlineResponse20091) HasStickyMacAllowListLimit() bool`

HasStickyMacAllowListLimit returns a boolean if a field has been set.

### GetStormControlEnabled

`func (o *InlineResponse20091) GetStormControlEnabled() bool`

GetStormControlEnabled returns the StormControlEnabled field if non-nil, zero value otherwise.

### GetStormControlEnabledOk

`func (o *InlineResponse20091) GetStormControlEnabledOk() (*bool, bool)`

GetStormControlEnabledOk returns a tuple with the StormControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControlEnabled

`func (o *InlineResponse20091) SetStormControlEnabled(v bool)`

SetStormControlEnabled sets StormControlEnabled field to given value.

### HasStormControlEnabled

`func (o *InlineResponse20091) HasStormControlEnabled() bool`

HasStormControlEnabled returns a boolean if a field has been set.

### GetFlexibleStackingEnabled

`func (o *InlineResponse20091) GetFlexibleStackingEnabled() bool`

GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field if non-nil, zero value otherwise.

### GetFlexibleStackingEnabledOk

`func (o *InlineResponse20091) GetFlexibleStackingEnabledOk() (*bool, bool)`

GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleStackingEnabled

`func (o *InlineResponse20091) SetFlexibleStackingEnabled(v bool)`

SetFlexibleStackingEnabled sets FlexibleStackingEnabled field to given value.

### HasFlexibleStackingEnabled

`func (o *InlineResponse20091) HasFlexibleStackingEnabled() bool`

HasFlexibleStackingEnabled returns a boolean if a field has been set.

### GetDaiTrusted

`func (o *InlineResponse20091) GetDaiTrusted() bool`

GetDaiTrusted returns the DaiTrusted field if non-nil, zero value otherwise.

### GetDaiTrustedOk

`func (o *InlineResponse20091) GetDaiTrustedOk() (*bool, bool)`

GetDaiTrustedOk returns a tuple with the DaiTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaiTrusted

`func (o *InlineResponse20091) SetDaiTrusted(v bool)`

SetDaiTrusted sets DaiTrusted field to given value.

### HasDaiTrusted

`func (o *InlineResponse20091) HasDaiTrusted() bool`

HasDaiTrusted returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse20091) GetProfile() DevicesSerialSwitchPortsProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse20091) GetProfileOk() (*DevicesSerialSwitchPortsProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse20091) SetProfile(v DevicesSerialSwitchPortsProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineResponse20091) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


