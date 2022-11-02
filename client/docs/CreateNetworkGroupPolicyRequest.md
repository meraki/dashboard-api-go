# CreateNetworkGroupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for your group policy. Required. | 
**Scheduling** | Pointer to [**CreateNetworkGroupPolicyRequestScheduling**](CreateNetworkGroupPolicyRequestScheduling.md) |  | [optional] 
**Bandwidth** | Pointer to [**CreateNetworkGroupPolicyRequestBandwidth**](CreateNetworkGroupPolicyRequestBandwidth.md) |  | [optional] 
**FirewallAndTrafficShaping** | Pointer to [**CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping**](CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping.md) |  | [optional] 
**ContentFiltering** | Pointer to [**CreateNetworkGroupPolicyRequestContentFiltering**](CreateNetworkGroupPolicyRequestContentFiltering.md) |  | [optional] 
**SplashAuthSettings** | Pointer to **string** | Whether clients bound to your policy will bypass splash authorization or behave according to the network&#39;s rules. Can be one of &#39;network default&#39; or &#39;bypass&#39;. Only available if your network has a wireless configuration. | [optional] 
**VlanTagging** | Pointer to [**CreateNetworkGroupPolicyRequestVlanTagging**](CreateNetworkGroupPolicyRequestVlanTagging.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**CreateNetworkGroupPolicyRequestBonjourForwarding**](CreateNetworkGroupPolicyRequestBonjourForwarding.md) |  | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequest

`func NewCreateNetworkGroupPolicyRequest(name string, ) *CreateNetworkGroupPolicyRequest`

NewCreateNetworkGroupPolicyRequest instantiates a new CreateNetworkGroupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestWithDefaults

`func NewCreateNetworkGroupPolicyRequestWithDefaults() *CreateNetworkGroupPolicyRequest`

NewCreateNetworkGroupPolicyRequestWithDefaults instantiates a new CreateNetworkGroupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkGroupPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkGroupPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkGroupPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScheduling

`func (o *CreateNetworkGroupPolicyRequest) GetScheduling() CreateNetworkGroupPolicyRequestScheduling`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *CreateNetworkGroupPolicyRequest) GetSchedulingOk() (*CreateNetworkGroupPolicyRequestScheduling, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *CreateNetworkGroupPolicyRequest) SetScheduling(v CreateNetworkGroupPolicyRequestScheduling)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *CreateNetworkGroupPolicyRequest) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.

### GetBandwidth

`func (o *CreateNetworkGroupPolicyRequest) GetBandwidth() CreateNetworkGroupPolicyRequestBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CreateNetworkGroupPolicyRequest) GetBandwidthOk() (*CreateNetworkGroupPolicyRequestBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CreateNetworkGroupPolicyRequest) SetBandwidth(v CreateNetworkGroupPolicyRequestBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *CreateNetworkGroupPolicyRequest) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetFirewallAndTrafficShaping

`func (o *CreateNetworkGroupPolicyRequest) GetFirewallAndTrafficShaping() CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping`

GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field if non-nil, zero value otherwise.

### GetFirewallAndTrafficShapingOk

`func (o *CreateNetworkGroupPolicyRequest) GetFirewallAndTrafficShapingOk() (*CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping, bool)`

GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallAndTrafficShaping

`func (o *CreateNetworkGroupPolicyRequest) SetFirewallAndTrafficShaping(v CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping)`

SetFirewallAndTrafficShaping sets FirewallAndTrafficShaping field to given value.

### HasFirewallAndTrafficShaping

`func (o *CreateNetworkGroupPolicyRequest) HasFirewallAndTrafficShaping() bool`

HasFirewallAndTrafficShaping returns a boolean if a field has been set.

### GetContentFiltering

`func (o *CreateNetworkGroupPolicyRequest) GetContentFiltering() CreateNetworkGroupPolicyRequestContentFiltering`

GetContentFiltering returns the ContentFiltering field if non-nil, zero value otherwise.

### GetContentFilteringOk

`func (o *CreateNetworkGroupPolicyRequest) GetContentFilteringOk() (*CreateNetworkGroupPolicyRequestContentFiltering, bool)`

GetContentFilteringOk returns a tuple with the ContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFiltering

`func (o *CreateNetworkGroupPolicyRequest) SetContentFiltering(v CreateNetworkGroupPolicyRequestContentFiltering)`

SetContentFiltering sets ContentFiltering field to given value.

### HasContentFiltering

`func (o *CreateNetworkGroupPolicyRequest) HasContentFiltering() bool`

HasContentFiltering returns a boolean if a field has been set.

### GetSplashAuthSettings

`func (o *CreateNetworkGroupPolicyRequest) GetSplashAuthSettings() string`

GetSplashAuthSettings returns the SplashAuthSettings field if non-nil, zero value otherwise.

### GetSplashAuthSettingsOk

`func (o *CreateNetworkGroupPolicyRequest) GetSplashAuthSettingsOk() (*string, bool)`

GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashAuthSettings

`func (o *CreateNetworkGroupPolicyRequest) SetSplashAuthSettings(v string)`

SetSplashAuthSettings sets SplashAuthSettings field to given value.

### HasSplashAuthSettings

`func (o *CreateNetworkGroupPolicyRequest) HasSplashAuthSettings() bool`

HasSplashAuthSettings returns a boolean if a field has been set.

### GetVlanTagging

`func (o *CreateNetworkGroupPolicyRequest) GetVlanTagging() CreateNetworkGroupPolicyRequestVlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *CreateNetworkGroupPolicyRequest) GetVlanTaggingOk() (*CreateNetworkGroupPolicyRequestVlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *CreateNetworkGroupPolicyRequest) SetVlanTagging(v CreateNetworkGroupPolicyRequestVlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *CreateNetworkGroupPolicyRequest) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *CreateNetworkGroupPolicyRequest) GetBonjourForwarding() CreateNetworkGroupPolicyRequestBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *CreateNetworkGroupPolicyRequest) GetBonjourForwardingOk() (*CreateNetworkGroupPolicyRequestBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *CreateNetworkGroupPolicyRequest) SetBonjourForwarding(v CreateNetworkGroupPolicyRequestBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *CreateNetworkGroupPolicyRequest) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


