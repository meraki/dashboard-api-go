# NetworksNetworkIdApplianceVpnBgpNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IPv4 address of the neighbor | [optional] 
**Ipv6** | Pointer to [**NetworksNetworkIdApplianceVpnBgpIpv6**](NetworksNetworkIdApplianceVpnBgpIpv6.md) |  | [optional] 
**RemoteAsNumber** | **int32** | Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295. | 
**ReceiveLimit** | Pointer to **int32** | The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 4294967295. When absent, it defaults to 0. | [optional] 
**AllowTransit** | Pointer to **bool** | When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false. | [optional] 
**EbgpHoldTimer** | **int32** | The EBGP hold timer in seconds for each neighbor. The EBGP hold timer must be an integer between 12 and 240. | 
**EbgpMultihop** | **int32** | Configure this if the neighbor is not adjacent. The EBGP multi-hop must be an integer between 1 and 255. | 

## Methods

### NewNetworksNetworkIdApplianceVpnBgpNeighbors

`func NewNetworksNetworkIdApplianceVpnBgpNeighbors(remoteAsNumber int32, ebgpHoldTimer int32, ebgpMultihop int32, ) *NetworksNetworkIdApplianceVpnBgpNeighbors`

NewNetworksNetworkIdApplianceVpnBgpNeighbors instantiates a new NetworksNetworkIdApplianceVpnBgpNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceVpnBgpNeighborsWithDefaults

`func NewNetworksNetworkIdApplianceVpnBgpNeighborsWithDefaults() *NetworksNetworkIdApplianceVpnBgpNeighbors`

NewNetworksNetworkIdApplianceVpnBgpNeighborsWithDefaults instantiates a new NetworksNetworkIdApplianceVpnBgpNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIpv6() NetworksNetworkIdApplianceVpnBgpIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetIpv6Ok() (*NetworksNetworkIdApplianceVpnBgpIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetIpv6(v NetworksNetworkIdApplianceVpnBgpIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.


### GetReceiveLimit

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetAllowTransit

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetAllowTransit() bool`

GetAllowTransit returns the AllowTransit field if non-nil, zero value otherwise.

### GetAllowTransitOk

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetAllowTransitOk() (*bool, bool)`

GetAllowTransitOk returns a tuple with the AllowTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransit

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetAllowTransit(v bool)`

SetAllowTransit sets AllowTransit field to given value.

### HasAllowTransit

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) HasAllowTransit() bool`

HasAllowTransit returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.


### GetEbgpMultihop

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *NetworksNetworkIdApplianceVpnBgpNeighbors) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


