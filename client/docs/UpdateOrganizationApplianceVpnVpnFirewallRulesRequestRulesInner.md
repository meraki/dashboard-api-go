# UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner

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

### NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner

`func NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner(policy string, protocol string, srcCidr string, destCidr string, ) *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner`

NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner instantiates a new UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInnerWithDefaults

`func NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInnerWithDefaults() *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner`

NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInnerWithDefaults instantiates a new UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSrcPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.


### GetDestPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.


### GetSyslogEnabled

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRulesRequestRulesInner) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


