# UpdateNetworkApplianceVpnBgpRequestNeighborsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IPv4 address of the neighbor | [optional] 
**Ipv6** | Pointer to [**UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6**](UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6.md) |  | [optional] 
**RemoteAsNumber** | **int32** | Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295. | 
**ReceiveLimit** | Pointer to **int32** | The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 4294967295. When absent, it defaults to 0. | [optional] 
**AllowTransit** | Pointer to **bool** | When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false. | [optional] 
**EbgpHoldTimer** | **int32** | The EBGP hold timer in seconds for each neighbor. The EBGP hold timer must be an integer between 12 and 240. | 
**EbgpMultihop** | **int32** | Configure this if the neighbor is not adjacent. The EBGP multi-hop must be an integer between 1 and 255. | 

## Methods

### NewUpdateNetworkApplianceVpnBgpRequestNeighborsInner

`func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInner(remoteAsNumber int32, ebgpHoldTimer int32, ebgpMultihop int32, ) *UpdateNetworkApplianceVpnBgpRequestNeighborsInner`

NewUpdateNetworkApplianceVpnBgpRequestNeighborsInner instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerWithDefaults

`func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerWithDefaults() *UpdateNetworkApplianceVpnBgpRequestNeighborsInner`

NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerWithDefaults instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIpv6() UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetIpv6Ok() (*UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetIpv6(v UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.


### GetReceiveLimit

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetAllowTransit

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetAllowTransit() bool`

GetAllowTransit returns the AllowTransit field if non-nil, zero value otherwise.

### GetAllowTransitOk

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetAllowTransitOk() (*bool, bool)`

GetAllowTransitOk returns a tuple with the AllowTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransit

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetAllowTransit(v bool)`

SetAllowTransit sets AllowTransit field to given value.

### HasAllowTransit

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) HasAllowTransit() bool`

HasAllowTransit returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.


### GetEbgpMultihop

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInner) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


