# GetNetworkApplianceTrafficShapingUplinkSelection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveActiveAutoVpnEnabled** | Pointer to **bool** | Whether active-active AutoVPN is enabled | [optional] 
**DefaultUplink** | Pointer to **string** | The default uplink. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | [optional] 
**LoadBalancingEnabled** | Pointer to **bool** | Whether load balancing is enabled | [optional] 
**FailoverAndFailback** | Pointer to [**GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback**](GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback.md) |  | [optional] 
**WanTrafficUplinkPreferences** | Pointer to [**[]GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner**](GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner.md) | Uplink preference rules for WAN traffic | [optional] 
**VpnTrafficUplinkPreferences** | Pointer to [**[]GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner**](GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner.md) | Uplink preference rules for VPN traffic | [optional] 

## Methods

### NewGetNetworkApplianceTrafficShapingUplinkSelection200Response

`func NewGetNetworkApplianceTrafficShapingUplinkSelection200Response() *GetNetworkApplianceTrafficShapingUplinkSelection200Response`

NewGetNetworkApplianceTrafficShapingUplinkSelection200Response instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWithDefaults

`func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200Response`

NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveActiveAutoVpnEnabled

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetActiveActiveAutoVpnEnabled() bool`

GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field if non-nil, zero value otherwise.

### GetActiveActiveAutoVpnEnabledOk

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetActiveActiveAutoVpnEnabledOk() (*bool, bool)`

GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActiveAutoVpnEnabled

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetActiveActiveAutoVpnEnabled(v bool)`

SetActiveActiveAutoVpnEnabled sets ActiveActiveAutoVpnEnabled field to given value.

### HasActiveActiveAutoVpnEnabled

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasActiveActiveAutoVpnEnabled() bool`

HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.

### GetDefaultUplink

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetDefaultUplink() string`

GetDefaultUplink returns the DefaultUplink field if non-nil, zero value otherwise.

### GetDefaultUplinkOk

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetDefaultUplinkOk() (*string, bool)`

GetDefaultUplinkOk returns a tuple with the DefaultUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUplink

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetDefaultUplink(v string)`

SetDefaultUplink sets DefaultUplink field to given value.

### HasDefaultUplink

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasDefaultUplink() bool`

HasDefaultUplink returns a boolean if a field has been set.

### GetLoadBalancingEnabled

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetLoadBalancingEnabled() bool`

GetLoadBalancingEnabled returns the LoadBalancingEnabled field if non-nil, zero value otherwise.

### GetLoadBalancingEnabledOk

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetLoadBalancingEnabledOk() (*bool, bool)`

GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingEnabled

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetLoadBalancingEnabled(v bool)`

SetLoadBalancingEnabled sets LoadBalancingEnabled field to given value.

### HasLoadBalancingEnabled

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasLoadBalancingEnabled() bool`

HasLoadBalancingEnabled returns a boolean if a field has been set.

### GetFailoverAndFailback

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetFailoverAndFailback() GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback`

GetFailoverAndFailback returns the FailoverAndFailback field if non-nil, zero value otherwise.

### GetFailoverAndFailbackOk

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetFailoverAndFailbackOk() (*GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback, bool)`

GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAndFailback

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetFailoverAndFailback(v GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback)`

SetFailoverAndFailback sets FailoverAndFailback field to given value.

### HasFailoverAndFailback

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasFailoverAndFailback() bool`

HasFailoverAndFailback returns a boolean if a field has been set.

### GetWanTrafficUplinkPreferences

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetWanTrafficUplinkPreferences() []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner`

GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetWanTrafficUplinkPreferencesOk

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetWanTrafficUplinkPreferencesOk() (*[]GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner, bool)`

GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanTrafficUplinkPreferences

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetWanTrafficUplinkPreferences(v []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner)`

SetWanTrafficUplinkPreferences sets WanTrafficUplinkPreferences field to given value.

### HasWanTrafficUplinkPreferences

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasWanTrafficUplinkPreferences() bool`

HasWanTrafficUplinkPreferences returns a boolean if a field has been set.

### GetVpnTrafficUplinkPreferences

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetVpnTrafficUplinkPreferences() []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner`

GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetVpnTrafficUplinkPreferencesOk

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetVpnTrafficUplinkPreferencesOk() (*[]GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner, bool)`

GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTrafficUplinkPreferences

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetVpnTrafficUplinkPreferences(v []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner)`

SetVpnTrafficUplinkPreferences sets VpnTrafficUplinkPreferences field to given value.

### HasVpnTrafficUplinkPreferences

`func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasVpnTrafficUplinkPreferences() bool`

HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


