# OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | 
**SrcPort** | Pointer to **string** | Comma-separated list of source port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**SrcCidr** | **string** | Comma-separated list of source IP address(es) (in IP or CIDR notation), or &#39;any&#39; (FQDN not supported) | 
**DestPort** | Pointer to **string** | Comma-separated list of destination port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**DestCidr** | **string** | Comma-separated list of destination IP address(es) (in IP or CIDR notation) or &#39;any&#39; (FQDN not supported) | 
**SyslogEnabled** | Pointer to **bool** | Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional) | [optional] 

## Methods

### NewOrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules

`func NewOrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules(policy string, protocol string, srcCidr string, destCidr string, ) *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules`

NewOrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules instantiates a new OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRulesWithDefaults

`func NewOrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRulesWithDefaults() *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules`

NewOrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRulesWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSrcPort

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcCidr

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.


### GetDestPort

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.


### GetSyslogEnabled

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


