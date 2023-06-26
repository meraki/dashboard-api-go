# UpdateNetworkSwitchAccessControlListsRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional). | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule. | 
**IpVersion** | Pointer to **string** | IP address version (must be &#39;any&#39;, &#39;ipv4&#39; or &#39;ipv6&#39;). Applicable only if network supports IPv6. Default value is &#39;ipv4&#39;. | [optional] 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, or &#39;any&#39;). | 
**SrcCidr** | **string** | Source IP address (in IP or CIDR notation) or &#39;any&#39;. | 
**SrcPort** | Pointer to **string** | Source port. Must be in the range  of 1-65535 or &#39;any&#39;. Default is &#39;any&#39;. | [optional] 
**DstCidr** | **string** | Destination IP address (in IP or CIDR notation) or &#39;any&#39;. | 
**DstPort** | Pointer to **string** | Destination port. Must be in the range of 1-65535 or &#39;any&#39;. Default is &#39;any&#39;. | [optional] 
**Vlan** | Pointer to **string** | Incoming traffic VLAN. Must be in the range of 1-4095 or &#39;any&#39;. Default is &#39;any&#39;. | [optional] 

## Methods

### NewUpdateNetworkSwitchAccessControlListsRequestRulesInner

`func NewUpdateNetworkSwitchAccessControlListsRequestRulesInner(policy string, protocol string, srcCidr string, dstCidr string, ) *UpdateNetworkSwitchAccessControlListsRequestRulesInner`

NewUpdateNetworkSwitchAccessControlListsRequestRulesInner instantiates a new UpdateNetworkSwitchAccessControlListsRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchAccessControlListsRequestRulesInnerWithDefaults

`func NewUpdateNetworkSwitchAccessControlListsRequestRulesInnerWithDefaults() *UpdateNetworkSwitchAccessControlListsRequestRulesInner`

NewUpdateNetworkSwitchAccessControlListsRequestRulesInnerWithDefaults instantiates a new UpdateNetworkSwitchAccessControlListsRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetIpVersion

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetProtocol

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSrcCidr

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.


### GetSrcPort

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstCidr

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetDstCidr() string`

GetDstCidr returns the DstCidr field if non-nil, zero value otherwise.

### GetDstCidrOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetDstCidrOk() (*string, bool)`

GetDstCidrOk returns a tuple with the DstCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstCidr

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetDstCidr(v string)`

SetDstCidr sets DstCidr field to given value.


### GetDstPort

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetVlan

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *UpdateNetworkSwitchAccessControlListsRequestRulesInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


