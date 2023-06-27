# UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | [**[]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner**](UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner.md) |     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.  | 
**PerClientBandwidthLimits** | Pointer to [**UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits**](UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits.md) |  | [optional] 
**DscpTagValue** | Pointer to **int32** |     The DSCP tag applied by your rule. null means &#39;Do not change DSCP tag&#39;.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.  | [optional] 
**PcpTagValue** | Pointer to **int32** |     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means &#39;Do not set PCP tag&#39;.  | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner

`func NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner(definitions []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner, ) *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner`

NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner instantiates a new UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInnerWithDefaults

`func NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInnerWithDefaults() *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner`

NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetDefinitions() []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetDefinitionsOk() (*[]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) SetDefinitions(v []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner)`

SetDefinitions sets Definitions field to given value.


### GetPerClientBandwidthLimits

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetPerClientBandwidthLimits() UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits`

GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitsOk

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetPerClientBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits, bool)`

GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimits

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) SetPerClientBandwidthLimits(v UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits)`

SetPerClientBandwidthLimits sets PerClientBandwidthLimits field to given value.

### HasPerClientBandwidthLimits

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) HasPerClientBandwidthLimits() bool`

HasPerClientBandwidthLimits returns a boolean if a field has been set.

### GetDscpTagValue

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetDscpTagValue() int32`

GetDscpTagValue returns the DscpTagValue field if non-nil, zero value otherwise.

### GetDscpTagValueOk

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetDscpTagValueOk() (*int32, bool)`

GetDscpTagValueOk returns a tuple with the DscpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpTagValue

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) SetDscpTagValue(v int32)`

SetDscpTagValue sets DscpTagValue field to given value.

### HasDscpTagValue

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) HasDscpTagValue() bool`

HasDscpTagValue returns a boolean if a field has been set.

### GetPcpTagValue

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetPcpTagValue() int32`

GetPcpTagValue returns the PcpTagValue field if non-nil, zero value otherwise.

### GetPcpTagValueOk

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) GetPcpTagValueOk() (*int32, bool)`

GetPcpTagValueOk returns a tuple with the PcpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcpTagValue

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) SetPcpTagValue(v int32)`

SetPcpTagValue sets PcpTagValue field to given value.

### HasPcpTagValue

`func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequestRulesInner) HasPcpTagValue() bool`

HasPcpTagValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


