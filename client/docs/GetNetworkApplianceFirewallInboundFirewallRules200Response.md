# GetNetworkApplianceFirewallInboundFirewallRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner**](GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner.md) | An ordered array of the firewall rules (not including the default rule) | [optional] 
**SyslogDefaultRule** | Pointer to **bool** | Log the special default rule (boolean value - enable only if you&#39;ve configured a syslog server) (optional) | [optional] 

## Methods

### NewGetNetworkApplianceFirewallInboundFirewallRules200Response

`func NewGetNetworkApplianceFirewallInboundFirewallRules200Response() *GetNetworkApplianceFirewallInboundFirewallRules200Response`

NewGetNetworkApplianceFirewallInboundFirewallRules200Response instantiates a new GetNetworkApplianceFirewallInboundFirewallRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseWithDefaults

`func NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseWithDefaults() *GetNetworkApplianceFirewallInboundFirewallRules200Response`

NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseWithDefaults instantiates a new GetNetworkApplianceFirewallInboundFirewallRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetRules() []GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetRulesOk() (*[]GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) SetRules(v []GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSyslogDefaultRule

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetSyslogDefaultRule() bool`

GetSyslogDefaultRule returns the SyslogDefaultRule field if non-nil, zero value otherwise.

### GetSyslogDefaultRuleOk

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetSyslogDefaultRuleOk() (*bool, bool)`

GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogDefaultRule

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) SetSyslogDefaultRule(v bool)`

SetSyslogDefaultRule sets SyslogDefaultRule field to given value.

### HasSyslogDefaultRule

`func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) HasSyslogDefaultRule() bool`

HasSyslogDefaultRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


