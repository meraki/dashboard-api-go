# UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveActiveAutoVpnEnabled** | Pointer to **bool** | Toggle for enabling or disabling active-active AutoVPN | [optional] 
**DefaultUplink** | Pointer to **string** | The default uplink. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | [optional] 
**LoadBalancingEnabled** | Pointer to **bool** | Toggle for enabling or disabling load balancing | [optional] 
**WanTrafficUplinkPreferences** | Pointer to [**[]UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner**](UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner.md) | Array of uplink preference rules for WAN traffic | [optional] 
**VpnTrafficUplinkPreferences** | Pointer to [**[]UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner**](UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner.md) | Array of uplink preference rules for VPN traffic | [optional] 

## Methods

### NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest

`func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest`

NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWithDefaults

`func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest`

NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveActiveAutoVpnEnabled

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetActiveActiveAutoVpnEnabled() bool`

GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field if non-nil, zero value otherwise.

### GetActiveActiveAutoVpnEnabledOk

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetActiveActiveAutoVpnEnabledOk() (*bool, bool)`

GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActiveAutoVpnEnabled

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetActiveActiveAutoVpnEnabled(v bool)`

SetActiveActiveAutoVpnEnabled sets ActiveActiveAutoVpnEnabled field to given value.

### HasActiveActiveAutoVpnEnabled

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasActiveActiveAutoVpnEnabled() bool`

HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.

### GetDefaultUplink

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetDefaultUplink() string`

GetDefaultUplink returns the DefaultUplink field if non-nil, zero value otherwise.

### GetDefaultUplinkOk

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetDefaultUplinkOk() (*string, bool)`

GetDefaultUplinkOk returns a tuple with the DefaultUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUplink

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetDefaultUplink(v string)`

SetDefaultUplink sets DefaultUplink field to given value.

### HasDefaultUplink

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasDefaultUplink() bool`

HasDefaultUplink returns a boolean if a field has been set.

### GetLoadBalancingEnabled

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetLoadBalancingEnabled() bool`

GetLoadBalancingEnabled returns the LoadBalancingEnabled field if non-nil, zero value otherwise.

### GetLoadBalancingEnabledOk

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetLoadBalancingEnabledOk() (*bool, bool)`

GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingEnabled

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetLoadBalancingEnabled(v bool)`

SetLoadBalancingEnabled sets LoadBalancingEnabled field to given value.

### HasLoadBalancingEnabled

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasLoadBalancingEnabled() bool`

HasLoadBalancingEnabled returns a boolean if a field has been set.

### GetWanTrafficUplinkPreferences

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetWanTrafficUplinkPreferences() []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner`

GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetWanTrafficUplinkPreferencesOk

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetWanTrafficUplinkPreferencesOk() (*[]UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner, bool)`

GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanTrafficUplinkPreferences

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetWanTrafficUplinkPreferences(v []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestWanTrafficUplinkPreferencesInner)`

SetWanTrafficUplinkPreferences sets WanTrafficUplinkPreferences field to given value.

### HasWanTrafficUplinkPreferences

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasWanTrafficUplinkPreferences() bool`

HasWanTrafficUplinkPreferences returns a boolean if a field has been set.

### GetVpnTrafficUplinkPreferences

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetVpnTrafficUplinkPreferences() []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner`

GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetVpnTrafficUplinkPreferencesOk

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) GetVpnTrafficUplinkPreferencesOk() (*[]UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner, bool)`

GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTrafficUplinkPreferences

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) SetVpnTrafficUplinkPreferences(v []UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInner)`

SetVpnTrafficUplinkPreferences sets VpnTrafficUplinkPreferences field to given value.

### HasVpnTrafficUplinkPreferences

`func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest) HasVpnTrafficUplinkPreferences() bool`

HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


