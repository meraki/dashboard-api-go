# NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataVlanId** | Pointer to **int32** | VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**VoiceVlanId** | Pointer to **int32** | VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**SuspendPortBounce** | Pointer to **bool** | Enable to suspend port bounce when RADIUS servers are unreachable | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth() *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataVlanId() int32`

GetDataVlanId returns the DataVlanId field if non-nil, zero value otherwise.

### GetDataVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataVlanIdOk() (*int32, bool)`

GetDataVlanIdOk returns a tuple with the DataVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetDataVlanId(v int32)`

SetDataVlanId sets DataVlanId field to given value.

### HasDataVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasDataVlanId() bool`

HasDataVlanId returns a boolean if a field has been set.

### GetVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceVlanId() int32`

GetVoiceVlanId returns the VoiceVlanId field if non-nil, zero value otherwise.

### GetVoiceVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceVlanIdOk() (*int32, bool)`

GetVoiceVlanIdOk returns a tuple with the VoiceVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetVoiceVlanId(v int32)`

SetVoiceVlanId sets VoiceVlanId field to given value.

### HasVoiceVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasVoiceVlanId() bool`

HasVoiceVlanId returns a boolean if a field has been set.

### GetSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetSuspendPortBounce() bool`

GetSuspendPortBounce returns the SuspendPortBounce field if non-nil, zero value otherwise.

### GetSuspendPortBounceOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetSuspendPortBounceOk() (*bool, bool)`

GetSuspendPortBounceOk returns a tuple with the SuspendPortBounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetSuspendPortBounce(v bool)`

SetSuspendPortBounce sets SuspendPortBounce field to given value.

### HasSuspendPortBounce

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasSuspendPortBounce() bool`

HasSuspendPortBounce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


