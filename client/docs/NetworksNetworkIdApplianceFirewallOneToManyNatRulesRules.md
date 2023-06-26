# NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIp** | **string** | The IP address that will be used to access the internal resource from the WAN | 
**Uplink** | **string** | The physical WAN interface on which the traffic will arrive (&#39;internet1&#39; or, if available, &#39;internet2&#39;) | 
**PortRules** | [**[]NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules**](NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules.md) | An array of associated forwarding rules | 

## Methods

### NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules

`func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules(publicIp string, uplink string, portRules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules, ) *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules`

NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRulesWithDefaults

`func NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRulesWithDefaults() *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules`

NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIp

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### GetUplink

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) SetUplink(v string)`

SetUplink sets Uplink field to given value.


### GetPortRules

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPortRules() []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules`

GetPortRules returns the PortRules field if non-nil, zero value otherwise.

### GetPortRulesOk

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) GetPortRulesOk() (*[]NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules, bool)`

GetPortRulesOk returns a tuple with the PortRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRules

`func (o *NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) SetPortRules(v []NetworksNetworkIdApplianceFirewallOneToManyNatRulesPortRules)`

SetPortRules sets PortRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


