# InlineObject43

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Set mode to &#39;disabled&#39;/&#39;detection&#39;/&#39;prevention&#39; (optional - omitting will leave current config unchanged) | [optional] 
**IdsRulesets** | Pointer to **string** | Set the detection ruleset &#39;connectivity&#39;/&#39;balanced&#39;/&#39;security&#39; (optional - omitting will leave current config unchanged). Default value is &#39;balanced&#39; if none currently saved | [optional] 
**ProtectedNetworks** | Pointer to [**NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks**](NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks.md) |  | [optional] 

## Methods

### NewInlineObject43

`func NewInlineObject43() *InlineObject43`

NewInlineObject43 instantiates a new InlineObject43 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject43WithDefaults

`func NewInlineObject43WithDefaults() *InlineObject43`

NewInlineObject43WithDefaults instantiates a new InlineObject43 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject43) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject43) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject43) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject43) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIdsRulesets

`func (o *InlineObject43) GetIdsRulesets() string`

GetIdsRulesets returns the IdsRulesets field if non-nil, zero value otherwise.

### GetIdsRulesetsOk

`func (o *InlineObject43) GetIdsRulesetsOk() (*string, bool)`

GetIdsRulesetsOk returns a tuple with the IdsRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsRulesets

`func (o *InlineObject43) SetIdsRulesets(v string)`

SetIdsRulesets sets IdsRulesets field to given value.

### HasIdsRulesets

`func (o *InlineObject43) HasIdsRulesets() bool`

HasIdsRulesets returns a boolean if a field has been set.

### GetProtectedNetworks

`func (o *InlineObject43) GetProtectedNetworks() NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks`

GetProtectedNetworks returns the ProtectedNetworks field if non-nil, zero value otherwise.

### GetProtectedNetworksOk

`func (o *InlineObject43) GetProtectedNetworksOk() (*NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks, bool)`

GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedNetworks

`func (o *InlineObject43) SetProtectedNetworks(v NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks)`

SetProtectedNetworks sets ProtectedNetworks field to given value.

### HasProtectedNetworks

`func (o *InlineObject43) HasProtectedNetworks() bool`

HasProtectedNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


