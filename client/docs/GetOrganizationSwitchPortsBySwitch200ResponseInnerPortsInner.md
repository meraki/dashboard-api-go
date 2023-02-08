# GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner

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
**RstpEnabled** | Pointer to **bool** | The rapid spanning tree protocol status. | [optional] 
**StpGuard** | Pointer to **string** | The state of the STP guard (&#39;disabled&#39;, &#39;root guard&#39;, &#39;bpdu guard&#39; or &#39;loop guard&#39;). | [optional] 
**LinkNegotiation** | Pointer to **string** | The link speed for the switch port. | [optional] 
**AccessPolicyType** | Pointer to **string** | The type of the access policy of the switch port. Only applicable to access ports. Can be one of &#39;Open&#39;, &#39;Custom access policy&#39;, &#39;MAC allow list&#39; or &#39;Sticky MAC allow list&#39;. | [optional] 
**StickyMacAllowList** | Pointer to **[]string** | The initial list of MAC addresses for sticky Mac allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StickyMacAllowListLimit** | Pointer to **int32** | The maximum number of MAC addresses for sticky MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 

## Methods

### NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner

`func NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner() *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner`

NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInnerWithDefaults

`func NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInnerWithDefaults() *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner`

NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInnerWithDefaults instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoeEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPoeEnabled() bool`

GetPoeEnabled returns the PoeEnabled field if non-nil, zero value otherwise.

### GetPoeEnabledOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPoeEnabledOk() (*bool, bool)`

GetPoeEnabledOk returns a tuple with the PoeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetPoeEnabled(v bool)`

SetPoeEnabled sets PoeEnabled field to given value.

### HasPoeEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasPoeEnabled() bool`

HasPoeEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVoiceVlan

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetRstpEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpGuard

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStpGuard() string`

GetStpGuard returns the StpGuard field if non-nil, zero value otherwise.

### GetStpGuardOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStpGuardOk() (*string, bool)`

GetStpGuardOk returns a tuple with the StpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpGuard

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetStpGuard(v string)`

SetStpGuard sets StpGuard field to given value.

### HasStpGuard

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasStpGuard() bool`

HasStpGuard returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetStickyMacAllowList

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowList() []string`

GetStickyMacAllowList returns the StickyMacAllowList field if non-nil, zero value otherwise.

### GetStickyMacAllowListOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowListOk() (*[]string, bool)`

GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowList

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetStickyMacAllowList(v []string)`

SetStickyMacAllowList sets StickyMacAllowList field to given value.

### HasStickyMacAllowList

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasStickyMacAllowList() bool`

HasStickyMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowListLimit

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowListLimit() int32`

GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field if non-nil, zero value otherwise.

### GetStickyMacAllowListLimitOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowListLimitOk() (*int32, bool)`

GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowListLimit

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetStickyMacAllowListLimit(v int32)`

SetStickyMacAllowListLimit sets StickyMacAllowListLimit field to given value.

### HasStickyMacAllowListLimit

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasStickyMacAllowListLimit() bool`

HasStickyMacAllowListLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


