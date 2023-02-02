# NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseDefault** | Pointer to **bool** | true/false whether to use special IPv4 addresses: https://tools.ietf.org/html/rfc5735 (required). Default value is true if none currently saved | [optional] 
**IncludedCidr** | Pointer to **[]string** | list of IP addresses or subnets being protected (required if &#39;useDefault&#39; is false) | [optional] 
**ExcludedCidr** | Pointer to **[]string** | list of IP addresses or subnets being excluded from protection (required if &#39;useDefault&#39; is false) | [optional] 

## Methods

### NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks

`func NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks() *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks`

NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks instantiates a new NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworksWithDefaults

`func NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworksWithDefaults() *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks`

NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworksWithDefaults instantiates a new NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseDefault

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetUseDefault() bool`

GetUseDefault returns the UseDefault field if non-nil, zero value otherwise.

### GetUseDefaultOk

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetUseDefaultOk() (*bool, bool)`

GetUseDefaultOk returns a tuple with the UseDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefault

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) SetUseDefault(v bool)`

SetUseDefault sets UseDefault field to given value.

### HasUseDefault

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) HasUseDefault() bool`

HasUseDefault returns a boolean if a field has been set.

### GetIncludedCidr

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetIncludedCidr() []string`

GetIncludedCidr returns the IncludedCidr field if non-nil, zero value otherwise.

### GetIncludedCidrOk

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetIncludedCidrOk() (*[]string, bool)`

GetIncludedCidrOk returns a tuple with the IncludedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedCidr

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) SetIncludedCidr(v []string)`

SetIncludedCidr sets IncludedCidr field to given value.

### HasIncludedCidr

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) HasIncludedCidr() bool`

HasIncludedCidr returns a boolean if a field has been set.

### GetExcludedCidr

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetExcludedCidr() []string`

GetExcludedCidr returns the ExcludedCidr field if non-nil, zero value otherwise.

### GetExcludedCidrOk

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetExcludedCidrOk() (*[]string, bool)`

GetExcludedCidrOk returns a tuple with the ExcludedCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCidr

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) SetExcludedCidr(v []string)`

SetExcludedCidr sets ExcludedCidr field to given value.

### HasExcludedCidr

`func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) HasExcludedCidr() bool`

HasExcludedCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


