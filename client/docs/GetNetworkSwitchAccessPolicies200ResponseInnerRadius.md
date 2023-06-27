# GetNetworkSwitchAccessPolicies200ResponseInnerRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriticalAuth** | Pointer to [**GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth**](GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth.md) |  | [optional] 
**FailedAuthVlanId** | Pointer to **int32** | VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth | [optional] 
**ReAuthenticationInterval** | Pointer to **int32** | Re-authentication period in seconds. Will be null if hostMode is Multi-Auth | [optional] 

## Methods

### NewGetNetworkSwitchAccessPolicies200ResponseInnerRadius

`func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadius() *GetNetworkSwitchAccessPolicies200ResponseInnerRadius`

NewGetNetworkSwitchAccessPolicies200ResponseInnerRadius instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusWithDefaults

`func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerRadius`

NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriticalAuth

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) GetCriticalAuth() GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth`

GetCriticalAuth returns the CriticalAuth field if non-nil, zero value otherwise.

### GetCriticalAuthOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) GetCriticalAuthOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth, bool)`

GetCriticalAuthOk returns a tuple with the CriticalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAuth

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) SetCriticalAuth(v GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth)`

SetCriticalAuth sets CriticalAuth field to given value.

### HasCriticalAuth

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) HasCriticalAuth() bool`

HasCriticalAuth returns a boolean if a field has been set.

### GetFailedAuthVlanId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) GetFailedAuthVlanId() int32`

GetFailedAuthVlanId returns the FailedAuthVlanId field if non-nil, zero value otherwise.

### GetFailedAuthVlanIdOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) GetFailedAuthVlanIdOk() (*int32, bool)`

GetFailedAuthVlanIdOk returns a tuple with the FailedAuthVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthVlanId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) SetFailedAuthVlanId(v int32)`

SetFailedAuthVlanId sets FailedAuthVlanId field to given value.

### HasFailedAuthVlanId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) HasFailedAuthVlanId() bool`

HasFailedAuthVlanId returns a boolean if a field has been set.

### GetReAuthenticationInterval

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) GetReAuthenticationInterval() int32`

GetReAuthenticationInterval returns the ReAuthenticationInterval field if non-nil, zero value otherwise.

### GetReAuthenticationIntervalOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) GetReAuthenticationIntervalOk() (*int32, bool)`

GetReAuthenticationIntervalOk returns a tuple with the ReAuthenticationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAuthenticationInterval

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) SetReAuthenticationInterval(v int32)`

SetReAuthenticationInterval sets ReAuthenticationInterval field to given value.

### HasReAuthenticationInterval

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadius) HasReAuthenticationInterval() bool`

HasReAuthenticationInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


