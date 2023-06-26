# GetNetworkSwitchAccessControlLists200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | Pointer to **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | [optional] 
**IpVersion** | Pointer to **string** | IP address version | [optional] 
**Protocol** | Pointer to **string** | The type of protocol | [optional] 
**SrcCidr** | Pointer to **string** | Source IP address (in IP or CIDR notation) | [optional] 
**SrcPort** | Pointer to **string** | Source port | [optional] 
**DstCidr** | Pointer to **string** | Destination IP address (in IP or CIDR notation) | [optional] 
**DstPort** | Pointer to **string** | Destination port | [optional] 
**Vlan** | Pointer to **string** | ncoming traffic VLAN | [optional] 

## Methods

### NewGetNetworkSwitchAccessControlLists200ResponseRulesInner

`func NewGetNetworkSwitchAccessControlLists200ResponseRulesInner() *GetNetworkSwitchAccessControlLists200ResponseRulesInner`

NewGetNetworkSwitchAccessControlLists200ResponseRulesInner instantiates a new GetNetworkSwitchAccessControlLists200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchAccessControlLists200ResponseRulesInnerWithDefaults

`func NewGetNetworkSwitchAccessControlLists200ResponseRulesInnerWithDefaults() *GetNetworkSwitchAccessControlLists200ResponseRulesInner`

NewGetNetworkSwitchAccessControlLists200ResponseRulesInnerWithDefaults instantiates a new GetNetworkSwitchAccessControlLists200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetIpVersion

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcCidr

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.

### HasSrcCidr

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasSrcCidr() bool`

HasSrcCidr returns a boolean if a field has been set.

### GetSrcPort

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstCidr

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetDstCidr() string`

GetDstCidr returns the DstCidr field if non-nil, zero value otherwise.

### GetDstCidrOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetDstCidrOk() (*string, bool)`

GetDstCidrOk returns a tuple with the DstCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstCidr

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetDstCidr(v string)`

SetDstCidr sets DstCidr field to given value.

### HasDstCidr

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasDstCidr() bool`

HasDstCidr returns a boolean if a field has been set.

### GetDstPort

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetVlan

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetNetworkSwitchAccessControlLists200ResponseRulesInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


