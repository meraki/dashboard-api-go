# UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of definition. Can be one of &#39;application&#39;, &#39;applicationCategory&#39;, &#39;host&#39;, &#39;port&#39;, &#39;ipRange&#39; or &#39;localNet&#39;. | 
**Value** | **string** |     If \&quot;type\&quot; is &#39;host&#39;, &#39;port&#39;, &#39;ipRange&#39; or &#39;localNet&#39;, then \&quot;value\&quot; must be a string, matching either     a hostname (e.g. \&quot;somesite.com\&quot;), a port (e.g. 8080), or an IP range (\&quot;192.1.0.0\&quot;,     \&quot;192.1.0.0/16\&quot;, or \&quot;10.1.0.0/16:80\&quot;). &#39;localNet&#39; also supports CIDR notation, excluding     custom ports.      If \&quot;type\&quot; is &#39;application&#39; or &#39;applicationCategory&#39;, then \&quot;value\&quot; must be an object     with the structure { \&quot;id\&quot;: \&quot;meraki:layer7/...\&quot; }, where \&quot;id\&quot; is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint).  | 

## Methods

### NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner

`func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner(type_ string, value string, ) *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner`

NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInnerWithDefaults

`func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner`

NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


