# UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | 
**SrcPort** | Pointer to **string** | Comma-separated list of source port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**SrcCidr** | **string** | Comma-separated list of source IP address(es) (in IP or CIDR notation), or &#39;any&#39; (note: FQDN not supported for source addresses) | 
**DestPort** | Pointer to **string** | Comma-separated list of destination port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**DestCidr** | **string** | Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or &#39;any&#39; | 
**SyslogEnabled** | Pointer to **bool** | Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional) | [optional] 

## Methods

### NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner

`func NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner(policy string, protocol string, srcCidr string, destCidr string, ) *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner`

NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner instantiates a new UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInnerWithDefaults

`func NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner`

NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSrcPort

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcCidr

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.


### GetDestPort

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.


### GetSyslogEnabled

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


