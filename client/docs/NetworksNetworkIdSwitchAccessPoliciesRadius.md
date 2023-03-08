# NetworksNetworkIdSwitchAccessPoliciesRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriticalAuth** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth**](NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth.md) |  | [optional] 
**FailedAuthVlanId** | Pointer to **int32** | VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**ReAuthenticationInterval** | Pointer to **int32** | Re-authentication period in seconds. Will be null if hostMode is Multi-Auth | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadius

`func NewNetworksNetworkIdSwitchAccessPoliciesRadius() *NetworksNetworkIdSwitchAccessPoliciesRadius`

NewNetworksNetworkIdSwitchAccessPoliciesRadius instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadius`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCriticalAuth() NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth`

GetCriticalAuth returns the CriticalAuth field if non-nil, zero value otherwise.

### GetCriticalAuthOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCriticalAuthOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth, bool)`

GetCriticalAuthOk returns a tuple with the CriticalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetCriticalAuth(v NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth)`

SetCriticalAuth sets CriticalAuth field to given value.

### HasCriticalAuth

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasCriticalAuth() bool`

HasCriticalAuth returns a boolean if a field has been set.

### GetFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthVlanId() int32`

GetFailedAuthVlanId returns the FailedAuthVlanId field if non-nil, zero value otherwise.

### GetFailedAuthVlanIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthVlanIdOk() (*int32, bool)`

GetFailedAuthVlanIdOk returns a tuple with the FailedAuthVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetFailedAuthVlanId(v int32)`

SetFailedAuthVlanId sets FailedAuthVlanId field to given value.

### HasFailedAuthVlanId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasFailedAuthVlanId() bool`

HasFailedAuthVlanId returns a boolean if a field has been set.

### GetReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetReAuthenticationInterval() int32`

GetReAuthenticationInterval returns the ReAuthenticationInterval field if non-nil, zero value otherwise.

### GetReAuthenticationIntervalOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetReAuthenticationIntervalOk() (*int32, bool)`

GetReAuthenticationIntervalOk returns a tuple with the ReAuthenticationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetReAuthenticationInterval(v int32)`

SetReAuthenticationInterval sets ReAuthenticationInterval field to given value.

### HasReAuthenticationInterval

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasReAuthenticationInterval() bool`

HasReAuthenticationInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


