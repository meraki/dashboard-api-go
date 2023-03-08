# InlineResponse20061Rules

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

### NewInlineResponse20061Rules

`func NewInlineResponse20061Rules() *InlineResponse20061Rules`

NewInlineResponse20061Rules instantiates a new InlineResponse20061Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20061RulesWithDefaults

`func NewInlineResponse20061RulesWithDefaults() *InlineResponse20061Rules`

NewInlineResponse20061RulesWithDefaults instantiates a new InlineResponse20061Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineResponse20061Rules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse20061Rules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse20061Rules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse20061Rules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse20061Rules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse20061Rules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse20061Rules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse20061Rules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse20061Rules) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse20061Rules) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse20061Rules) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse20061Rules) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20061Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20061Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20061Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20061Rules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcCidr

`func (o *InlineResponse20061Rules) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *InlineResponse20061Rules) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *InlineResponse20061Rules) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.

### HasSrcCidr

`func (o *InlineResponse20061Rules) HasSrcCidr() bool`

HasSrcCidr returns a boolean if a field has been set.

### GetSrcPort

`func (o *InlineResponse20061Rules) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InlineResponse20061Rules) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InlineResponse20061Rules) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InlineResponse20061Rules) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstCidr

`func (o *InlineResponse20061Rules) GetDstCidr() string`

GetDstCidr returns the DstCidr field if non-nil, zero value otherwise.

### GetDstCidrOk

`func (o *InlineResponse20061Rules) GetDstCidrOk() (*string, bool)`

GetDstCidrOk returns a tuple with the DstCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstCidr

`func (o *InlineResponse20061Rules) SetDstCidr(v string)`

SetDstCidr sets DstCidr field to given value.

### HasDstCidr

`func (o *InlineResponse20061Rules) HasDstCidr() bool`

HasDstCidr returns a boolean if a field has been set.

### GetDstPort

`func (o *InlineResponse20061Rules) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *InlineResponse20061Rules) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *InlineResponse20061Rules) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *InlineResponse20061Rules) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20061Rules) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20061Rules) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20061Rules) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20061Rules) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


