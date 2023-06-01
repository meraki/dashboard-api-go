# InlineResponse20019

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveActiveAutoVpnEnabled** | Pointer to **bool** | Whether active-active AutoVPN is enabled | [optional] 
**DefaultUplink** | Pointer to **string** | The default uplink. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | [optional] 
**LoadBalancingEnabled** | Pointer to **bool** | Whether load balancing is enabled | [optional] 
**FailoverAndFailback** | Pointer to [**InlineResponse20019FailoverAndFailback**](InlineResponse20019FailoverAndFailback.md) |  | [optional] 
**WanTrafficUplinkPreferences** | Pointer to [**[]InlineResponse20019WanTrafficUplinkPreferences**](InlineResponse20019WanTrafficUplinkPreferences.md) | Uplink preference rules for WAN traffic | [optional] 
**VpnTrafficUplinkPreferences** | Pointer to [**[]InlineResponse20019VpnTrafficUplinkPreferences**](InlineResponse20019VpnTrafficUplinkPreferences.md) | Uplink preference rules for VPN traffic | [optional] 

## Methods

### NewInlineResponse20019

`func NewInlineResponse20019() *InlineResponse20019`

NewInlineResponse20019 instantiates a new InlineResponse20019 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019WithDefaults

`func NewInlineResponse20019WithDefaults() *InlineResponse20019`

NewInlineResponse20019WithDefaults instantiates a new InlineResponse20019 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveActiveAutoVpnEnabled

`func (o *InlineResponse20019) GetActiveActiveAutoVpnEnabled() bool`

GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field if non-nil, zero value otherwise.

### GetActiveActiveAutoVpnEnabledOk

`func (o *InlineResponse20019) GetActiveActiveAutoVpnEnabledOk() (*bool, bool)`

GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActiveAutoVpnEnabled

`func (o *InlineResponse20019) SetActiveActiveAutoVpnEnabled(v bool)`

SetActiveActiveAutoVpnEnabled sets ActiveActiveAutoVpnEnabled field to given value.

### HasActiveActiveAutoVpnEnabled

`func (o *InlineResponse20019) HasActiveActiveAutoVpnEnabled() bool`

HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.

### GetDefaultUplink

`func (o *InlineResponse20019) GetDefaultUplink() string`

GetDefaultUplink returns the DefaultUplink field if non-nil, zero value otherwise.

### GetDefaultUplinkOk

`func (o *InlineResponse20019) GetDefaultUplinkOk() (*string, bool)`

GetDefaultUplinkOk returns a tuple with the DefaultUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUplink

`func (o *InlineResponse20019) SetDefaultUplink(v string)`

SetDefaultUplink sets DefaultUplink field to given value.

### HasDefaultUplink

`func (o *InlineResponse20019) HasDefaultUplink() bool`

HasDefaultUplink returns a boolean if a field has been set.

### GetLoadBalancingEnabled

`func (o *InlineResponse20019) GetLoadBalancingEnabled() bool`

GetLoadBalancingEnabled returns the LoadBalancingEnabled field if non-nil, zero value otherwise.

### GetLoadBalancingEnabledOk

`func (o *InlineResponse20019) GetLoadBalancingEnabledOk() (*bool, bool)`

GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingEnabled

`func (o *InlineResponse20019) SetLoadBalancingEnabled(v bool)`

SetLoadBalancingEnabled sets LoadBalancingEnabled field to given value.

### HasLoadBalancingEnabled

`func (o *InlineResponse20019) HasLoadBalancingEnabled() bool`

HasLoadBalancingEnabled returns a boolean if a field has been set.

### GetFailoverAndFailback

`func (o *InlineResponse20019) GetFailoverAndFailback() InlineResponse20019FailoverAndFailback`

GetFailoverAndFailback returns the FailoverAndFailback field if non-nil, zero value otherwise.

### GetFailoverAndFailbackOk

`func (o *InlineResponse20019) GetFailoverAndFailbackOk() (*InlineResponse20019FailoverAndFailback, bool)`

GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAndFailback

`func (o *InlineResponse20019) SetFailoverAndFailback(v InlineResponse20019FailoverAndFailback)`

SetFailoverAndFailback sets FailoverAndFailback field to given value.

### HasFailoverAndFailback

`func (o *InlineResponse20019) HasFailoverAndFailback() bool`

HasFailoverAndFailback returns a boolean if a field has been set.

### GetWanTrafficUplinkPreferences

`func (o *InlineResponse20019) GetWanTrafficUplinkPreferences() []InlineResponse20019WanTrafficUplinkPreferences`

GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetWanTrafficUplinkPreferencesOk

`func (o *InlineResponse20019) GetWanTrafficUplinkPreferencesOk() (*[]InlineResponse20019WanTrafficUplinkPreferences, bool)`

GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanTrafficUplinkPreferences

`func (o *InlineResponse20019) SetWanTrafficUplinkPreferences(v []InlineResponse20019WanTrafficUplinkPreferences)`

SetWanTrafficUplinkPreferences sets WanTrafficUplinkPreferences field to given value.

### HasWanTrafficUplinkPreferences

`func (o *InlineResponse20019) HasWanTrafficUplinkPreferences() bool`

HasWanTrafficUplinkPreferences returns a boolean if a field has been set.

### GetVpnTrafficUplinkPreferences

`func (o *InlineResponse20019) GetVpnTrafficUplinkPreferences() []InlineResponse20019VpnTrafficUplinkPreferences`

GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetVpnTrafficUplinkPreferencesOk

`func (o *InlineResponse20019) GetVpnTrafficUplinkPreferencesOk() (*[]InlineResponse20019VpnTrafficUplinkPreferences, bool)`

GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTrafficUplinkPreferences

`func (o *InlineResponse20019) SetVpnTrafficUplinkPreferences(v []InlineResponse20019VpnTrafficUplinkPreferences)`

SetVpnTrafficUplinkPreferences sets VpnTrafficUplinkPreferences field to given value.

### HasVpnTrafficUplinkPreferences

`func (o *InlineResponse20019) HasVpnTrafficUplinkPreferences() bool`

HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


