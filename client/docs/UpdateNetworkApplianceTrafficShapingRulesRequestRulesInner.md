# UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | [**[]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner**](UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner.md) |     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.  | 
**PerClientBandwidthLimits** | Pointer to [**UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits**](UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits.md) |  | [optional] 
**DscpTagValue** | Pointer to **int32** |     The DSCP tag applied by your rule. null means &#39;Do not change DSCP tag&#39;.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.  | [optional] 
**Priority** | Pointer to **string** |     A string, indicating the priority level for packets bound to your rule.     Can be &#39;low&#39;, &#39;normal&#39; or &#39;high&#39;.  | [optional] 

## Methods

### NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInner

`func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInner(definitions []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner, ) *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner`

NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInner instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerWithDefaults

`func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner`

NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetDefinitions() []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetDefinitionsOk() (*[]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) SetDefinitions(v []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner)`

SetDefinitions sets Definitions field to given value.


### GetPerClientBandwidthLimits

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetPerClientBandwidthLimits() UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits`

GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitsOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetPerClientBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits, bool)`

GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimits

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) SetPerClientBandwidthLimits(v UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits)`

SetPerClientBandwidthLimits sets PerClientBandwidthLimits field to given value.

### HasPerClientBandwidthLimits

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) HasPerClientBandwidthLimits() bool`

HasPerClientBandwidthLimits returns a boolean if a field has been set.

### GetDscpTagValue

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetDscpTagValue() int32`

GetDscpTagValue returns the DscpTagValue field if non-nil, zero value otherwise.

### GetDscpTagValueOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetDscpTagValueOk() (*int32, bool)`

GetDscpTagValueOk returns a tuple with the DscpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpTagValue

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) SetDscpTagValue(v int32)`

SetDscpTagValue sets DscpTagValue field to given value.

### HasDscpTagValue

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) HasDscpTagValue() bool`

HasDscpTagValue returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


