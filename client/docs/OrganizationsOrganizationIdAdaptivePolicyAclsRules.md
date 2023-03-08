# OrganizationsOrganizationIdAdaptivePolicyAclsRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule. | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39; or &#39;any&#39;). | 
**SrcPort** | Pointer to **string** | Source port. Must be in the format of single port: &#39;1&#39;, port list: &#39;1,2&#39; or port range: &#39;1-10&#39;, and in the range of 1-65535, or &#39;any&#39;. Default is &#39;any&#39;. | [optional] 
**DstPort** | Pointer to **string** | Destination port. Must be in the format of single port: &#39;1&#39;, port list: &#39;1,2&#39; or port range: &#39;1-10&#39;, and in the range of 1-65535, or &#39;any&#39;. Default is &#39;any&#39;. | [optional] 

## Methods

### NewOrganizationsOrganizationIdAdaptivePolicyAclsRules

`func NewOrganizationsOrganizationIdAdaptivePolicyAclsRules(policy string, protocol string, ) *OrganizationsOrganizationIdAdaptivePolicyAclsRules`

NewOrganizationsOrganizationIdAdaptivePolicyAclsRules instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAdaptivePolicyAclsRulesWithDefaults

`func NewOrganizationsOrganizationIdAdaptivePolicyAclsRulesWithDefaults() *OrganizationsOrganizationIdAdaptivePolicyAclsRules`

NewOrganizationsOrganizationIdAdaptivePolicyAclsRulesWithDefaults instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSrcPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


