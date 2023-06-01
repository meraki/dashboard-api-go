# InlineResponse20083

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeshingEnabled** | Pointer to **bool** | Toggle for enabling or disabling meshing in a network | [optional] 
**Ipv6BridgeEnabled** | Pointer to **bool** | Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode) | [optional] 
**LocationAnalyticsEnabled** | Pointer to **bool** | Toggle for enabling or disabling location analytics for your network | [optional] 
**UpgradeStrategy** | Pointer to **string** | The upgrade strategy to apply to the network. Must be one of &#39;minimizeUpgradeTime&#39; or &#39;minimizeClientDowntime&#39;. Requires firmware version MR 26.8 or higher&#39; | [optional] 
**LedLightsOn** | Pointer to **bool** | Toggle for enabling or disabling LED lights on all APs in the network (making them run dark) | [optional] 
**NamedVlans** | Pointer to [**InlineResponse20083NamedVlans**](InlineResponse20083NamedVlans.md) |  | [optional] 

## Methods

### NewInlineResponse20083

`func NewInlineResponse20083() *InlineResponse20083`

NewInlineResponse20083 instantiates a new InlineResponse20083 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20083WithDefaults

`func NewInlineResponse20083WithDefaults() *InlineResponse20083`

NewInlineResponse20083WithDefaults instantiates a new InlineResponse20083 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeshingEnabled

`func (o *InlineResponse20083) GetMeshingEnabled() bool`

GetMeshingEnabled returns the MeshingEnabled field if non-nil, zero value otherwise.

### GetMeshingEnabledOk

`func (o *InlineResponse20083) GetMeshingEnabledOk() (*bool, bool)`

GetMeshingEnabledOk returns a tuple with the MeshingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshingEnabled

`func (o *InlineResponse20083) SetMeshingEnabled(v bool)`

SetMeshingEnabled sets MeshingEnabled field to given value.

### HasMeshingEnabled

`func (o *InlineResponse20083) HasMeshingEnabled() bool`

HasMeshingEnabled returns a boolean if a field has been set.

### GetIpv6BridgeEnabled

`func (o *InlineResponse20083) GetIpv6BridgeEnabled() bool`

GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field if non-nil, zero value otherwise.

### GetIpv6BridgeEnabledOk

`func (o *InlineResponse20083) GetIpv6BridgeEnabledOk() (*bool, bool)`

GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6BridgeEnabled

`func (o *InlineResponse20083) SetIpv6BridgeEnabled(v bool)`

SetIpv6BridgeEnabled sets Ipv6BridgeEnabled field to given value.

### HasIpv6BridgeEnabled

`func (o *InlineResponse20083) HasIpv6BridgeEnabled() bool`

HasIpv6BridgeEnabled returns a boolean if a field has been set.

### GetLocationAnalyticsEnabled

`func (o *InlineResponse20083) GetLocationAnalyticsEnabled() bool`

GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field if non-nil, zero value otherwise.

### GetLocationAnalyticsEnabledOk

`func (o *InlineResponse20083) GetLocationAnalyticsEnabledOk() (*bool, bool)`

GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnalyticsEnabled

`func (o *InlineResponse20083) SetLocationAnalyticsEnabled(v bool)`

SetLocationAnalyticsEnabled sets LocationAnalyticsEnabled field to given value.

### HasLocationAnalyticsEnabled

`func (o *InlineResponse20083) HasLocationAnalyticsEnabled() bool`

HasLocationAnalyticsEnabled returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *InlineResponse20083) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *InlineResponse20083) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *InlineResponse20083) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *InlineResponse20083) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.

### GetLedLightsOn

`func (o *InlineResponse20083) GetLedLightsOn() bool`

GetLedLightsOn returns the LedLightsOn field if non-nil, zero value otherwise.

### GetLedLightsOnOk

`func (o *InlineResponse20083) GetLedLightsOnOk() (*bool, bool)`

GetLedLightsOnOk returns a tuple with the LedLightsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedLightsOn

`func (o *InlineResponse20083) SetLedLightsOn(v bool)`

SetLedLightsOn sets LedLightsOn field to given value.

### HasLedLightsOn

`func (o *InlineResponse20083) HasLedLightsOn() bool`

HasLedLightsOn returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineResponse20083) GetNamedVlans() InlineResponse20083NamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineResponse20083) GetNamedVlansOk() (*InlineResponse20083NamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineResponse20083) SetNamedVlans(v InlineResponse20083NamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineResponse20083) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


