# UpdateNetworkApplianceSingleLanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | The subnet of the single LAN configuration | [optional] 
**ApplianceIp** | Pointer to **string** | The appliance IP address of the single LAN | [optional] 
**Ipv6** | Pointer to [**UpdateNetworkApplianceSingleLanRequestIpv6**](UpdateNetworkApplianceSingleLanRequestIpv6.md) |  | [optional] 
**MandatoryDhcp** | Pointer to [**UpdateNetworkApplianceSingleLanRequestMandatoryDhcp**](UpdateNetworkApplianceSingleLanRequestMandatoryDhcp.md) |  | [optional] 

## Methods

### NewUpdateNetworkApplianceSingleLanRequest

`func NewUpdateNetworkApplianceSingleLanRequest() *UpdateNetworkApplianceSingleLanRequest`

NewUpdateNetworkApplianceSingleLanRequest instantiates a new UpdateNetworkApplianceSingleLanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceSingleLanRequestWithDefaults

`func NewUpdateNetworkApplianceSingleLanRequestWithDefaults() *UpdateNetworkApplianceSingleLanRequest`

NewUpdateNetworkApplianceSingleLanRequestWithDefaults instantiates a new UpdateNetworkApplianceSingleLanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *UpdateNetworkApplianceSingleLanRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateNetworkApplianceSingleLanRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateNetworkApplianceSingleLanRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateNetworkApplianceSingleLanRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *UpdateNetworkApplianceSingleLanRequest) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *UpdateNetworkApplianceSingleLanRequest) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *UpdateNetworkApplianceSingleLanRequest) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *UpdateNetworkApplianceSingleLanRequest) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetIpv6

`func (o *UpdateNetworkApplianceSingleLanRequest) GetIpv6() UpdateNetworkApplianceSingleLanRequestIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *UpdateNetworkApplianceSingleLanRequest) GetIpv6Ok() (*UpdateNetworkApplianceSingleLanRequestIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *UpdateNetworkApplianceSingleLanRequest) SetIpv6(v UpdateNetworkApplianceSingleLanRequestIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *UpdateNetworkApplianceSingleLanRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *UpdateNetworkApplianceSingleLanRequest) GetMandatoryDhcp() UpdateNetworkApplianceSingleLanRequestMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *UpdateNetworkApplianceSingleLanRequest) GetMandatoryDhcpOk() (*UpdateNetworkApplianceSingleLanRequestMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *UpdateNetworkApplianceSingleLanRequest) SetMandatoryDhcp(v UpdateNetworkApplianceSingleLanRequestMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *UpdateNetworkApplianceSingleLanRequest) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


