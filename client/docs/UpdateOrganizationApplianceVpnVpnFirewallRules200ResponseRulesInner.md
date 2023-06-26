# UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule | [optional] 
**Policy** | Pointer to **string** | &#39;Deny&#39; traffic specified by this rule | [optional] 
**Protocol** | Pointer to **string** | The type of protocol | [optional] 
**SrcPort** | Pointer to **string** | List of source ports | [optional] 
**SrcCidr** | Pointer to **string** | List of source IP addresses | [optional] 
**DestPort** | Pointer to **string** | List of destination ports | [optional] 
**DestCidr** | Pointer to **string** | List of destination IP addresses | [optional] 
**SyslogEnabled** | Pointer to **bool** | Flag indicating whether the rule should be logged to syslog | [optional] 

## Methods

### NewUpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner

`func NewUpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner() *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner`

NewUpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner instantiates a new UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInnerWithDefaults

`func NewUpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInnerWithDefaults() *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner`

NewUpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInnerWithDefaults instantiates a new UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProtocol

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.

### HasSrcCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasSrcCidr() bool`

HasSrcCidr returns a boolean if a field has been set.

### GetDestPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.

### HasDestCidr

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasDestCidr() bool`

HasDestCidr returns a boolean if a field has been set.

### GetSyslogEnabled

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *UpdateOrganizationApplianceVpnVpnFirewallRules200ResponseRulesInner) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


