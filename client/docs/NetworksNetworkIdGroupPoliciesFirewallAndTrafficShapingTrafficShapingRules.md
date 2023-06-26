# NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | [**[]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions**](NetworksNetworkIdApplianceTrafficShapingRulesDefinitions.md) |     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.  | 
**PerClientBandwidthLimits** | Pointer to [**NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits**](NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits.md) |  | [optional] 
**DscpTagValue** | Pointer to **int32** |     The DSCP tag applied by your rule. null means &#39;Do not change DSCP tag&#39;.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.  | [optional] 
**PcpTagValue** | Pointer to **int32** |     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means &#39;Do not set PCP tag&#39;.  | [optional] 
**Priority** | Pointer to **string** |     A string, indicating the priority level for packets bound to your rule.     Can be &#39;low&#39;, &#39;normal&#39; or &#39;high&#39;.  | [optional] 

## Methods

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules(definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, ) *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesWithDefaults

`func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules`

NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDefinitions() []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDefinitionsOk() (*[]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetDefinitions(v []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions)`

SetDefinitions sets Definitions field to given value.


### GetPerClientBandwidthLimits

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPerClientBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits`

GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitsOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPerClientBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits, bool)`

GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimits

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetPerClientBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits)`

SetPerClientBandwidthLimits sets PerClientBandwidthLimits field to given value.

### HasPerClientBandwidthLimits

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasPerClientBandwidthLimits() bool`

HasPerClientBandwidthLimits returns a boolean if a field has been set.

### GetDscpTagValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDscpTagValue() int32`

GetDscpTagValue returns the DscpTagValue field if non-nil, zero value otherwise.

### GetDscpTagValueOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDscpTagValueOk() (*int32, bool)`

GetDscpTagValueOk returns a tuple with the DscpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpTagValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetDscpTagValue(v int32)`

SetDscpTagValue sets DscpTagValue field to given value.

### HasDscpTagValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasDscpTagValue() bool`

HasDscpTagValue returns a boolean if a field has been set.

### GetPcpTagValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPcpTagValue() int32`

GetPcpTagValue returns the PcpTagValue field if non-nil, zero value otherwise.

### GetPcpTagValueOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPcpTagValueOk() (*int32, bool)`

GetPcpTagValueOk returns a tuple with the PcpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcpTagValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetPcpTagValue(v int32)`

SetPcpTagValue sets PcpTagValue field to given value.

### HasPcpTagValue

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasPcpTagValue() bool`

HasPcpTagValue returns a boolean if a field has been set.

### GetPriority

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


