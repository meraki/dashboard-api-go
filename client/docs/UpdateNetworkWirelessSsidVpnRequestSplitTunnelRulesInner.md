# UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol for this split tunnel rule. | [optional] 
**DestCidr** | **string** | Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or &#39;any&#39;. | 
**DestPort** | Pointer to **string** | Destination port for this split tunnel rule, (integer in the range 1-65535), or &#39;any&#39;. | [optional] 
**Policy** | **string** | Traffic policy specified for this split tunnel rule, &#39;allow&#39; or &#39;deny&#39;. | 
**Comment** | Pointer to **string** | Description for this split tunnel rule (optional). | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner

`func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner(destCidr string, policy string, ) *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner`

NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInnerWithDefaults

`func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInnerWithDefaults() *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner`

NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestCidr

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.


### GetDestPort

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetComment

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


