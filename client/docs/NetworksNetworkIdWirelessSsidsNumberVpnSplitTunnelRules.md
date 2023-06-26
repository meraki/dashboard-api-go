# NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol for this split tunnel rule. | [optional] 
**DestCidr** | **string** | Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or &#39;any&#39;. | 
**DestPort** | Pointer to **string** | Destination port for this split tunnel rule, (integer in the range 1-65535), or &#39;any&#39;. | [optional] 
**Policy** | **string** | Traffic policy specified for this split tunnel rule, &#39;allow&#39; or &#39;deny&#39;. | 
**Comment** | Pointer to **string** | Description for this split tunnel rule (optional). | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules

`func NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules(destCidr string, policy string, ) *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules`

NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRulesWithDefaults

`func NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRulesWithDefaults() *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules`

NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRulesWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestCidr

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.


### GetDestPort

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetPolicy

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetComment

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


