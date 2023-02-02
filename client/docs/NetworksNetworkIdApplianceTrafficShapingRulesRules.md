# NetworksNetworkIdApplianceTrafficShapingRulesRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | [**[]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions**](NetworksNetworkIdApplianceTrafficShapingRulesDefinitions.md) |     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.  | 
**PerClientBandwidthLimits** | Pointer to [**NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits**](NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits.md) |  | [optional] 
**DscpTagValue** | Pointer to **int32** |     The DSCP tag applied by your rule. null means &#39;Do not change DSCP tag&#39;.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.  | [optional] 
**Priority** | Pointer to **string** |     A string, indicating the priority level for packets bound to your rule.     Can be &#39;low&#39;, &#39;normal&#39; or &#39;high&#39;.  | [optional] 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingRulesRules

`func NewNetworksNetworkIdApplianceTrafficShapingRulesRules(definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, ) *NetworksNetworkIdApplianceTrafficShapingRulesRules`

NewNetworksNetworkIdApplianceTrafficShapingRulesRules instantiates a new NetworksNetworkIdApplianceTrafficShapingRulesRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingRulesRulesWithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingRulesRulesWithDefaults() *NetworksNetworkIdApplianceTrafficShapingRulesRules`

NewNetworksNetworkIdApplianceTrafficShapingRulesRulesWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingRulesRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDefinitions() []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDefinitionsOk() (*[]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetDefinitions(v []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions)`

SetDefinitions sets Definitions field to given value.


### GetPerClientBandwidthLimits

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPerClientBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits`

GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitsOk

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPerClientBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits, bool)`

GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimits

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetPerClientBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits)`

SetPerClientBandwidthLimits sets PerClientBandwidthLimits field to given value.

### HasPerClientBandwidthLimits

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) HasPerClientBandwidthLimits() bool`

HasPerClientBandwidthLimits returns a boolean if a field has been set.

### GetDscpTagValue

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDscpTagValue() int32`

GetDscpTagValue returns the DscpTagValue field if non-nil, zero value otherwise.

### GetDscpTagValueOk

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetDscpTagValueOk() (*int32, bool)`

GetDscpTagValueOk returns a tuple with the DscpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpTagValue

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetDscpTagValue(v int32)`

SetDscpTagValue sets DscpTagValue field to given value.

### HasDscpTagValue

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) HasDscpTagValue() bool`

HasDscpTagValue returns a boolean if a field has been set.

### GetPriority

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworksNetworkIdApplianceTrafficShapingRulesRules) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


