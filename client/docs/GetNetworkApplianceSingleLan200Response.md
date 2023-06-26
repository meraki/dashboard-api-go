# GetNetworkApplianceSingleLan200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | The subnet of the single LAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the single LAN | [optional] 
**MandatoryDhcp** | Pointer to [**GetNetworkApplianceSingleLan200ResponseMandatoryDhcp**](GetNetworkApplianceSingleLan200ResponseMandatoryDhcp.md) |  | [optional] 
**Ipv6** | Pointer to [**GetNetworkApplianceSingleLan200ResponseIpv6**](GetNetworkApplianceSingleLan200ResponseIpv6.md) |  | [optional] 

## Methods

### NewGetNetworkApplianceSingleLan200Response

`func NewGetNetworkApplianceSingleLan200Response() *GetNetworkApplianceSingleLan200Response`

NewGetNetworkApplianceSingleLan200Response instantiates a new GetNetworkApplianceSingleLan200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceSingleLan200ResponseWithDefaults

`func NewGetNetworkApplianceSingleLan200ResponseWithDefaults() *GetNetworkApplianceSingleLan200Response`

NewGetNetworkApplianceSingleLan200ResponseWithDefaults instantiates a new GetNetworkApplianceSingleLan200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *GetNetworkApplianceSingleLan200Response) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetNetworkApplianceSingleLan200Response) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetNetworkApplianceSingleLan200Response) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *GetNetworkApplianceSingleLan200Response) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *GetNetworkApplianceSingleLan200Response) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *GetNetworkApplianceSingleLan200Response) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *GetNetworkApplianceSingleLan200Response) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *GetNetworkApplianceSingleLan200Response) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *GetNetworkApplianceSingleLan200Response) GetMandatoryDhcp() GetNetworkApplianceSingleLan200ResponseMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *GetNetworkApplianceSingleLan200Response) GetMandatoryDhcpOk() (*GetNetworkApplianceSingleLan200ResponseMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *GetNetworkApplianceSingleLan200Response) SetMandatoryDhcp(v GetNetworkApplianceSingleLan200ResponseMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *GetNetworkApplianceSingleLan200Response) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.

### GetIpv6

`func (o *GetNetworkApplianceSingleLan200Response) GetIpv6() GetNetworkApplianceSingleLan200ResponseIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GetNetworkApplianceSingleLan200Response) GetIpv6Ok() (*GetNetworkApplianceSingleLan200ResponseIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GetNetworkApplianceSingleLan200Response) SetIpv6(v GetNetworkApplianceSingleLan200ResponseIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *GetNetworkApplianceSingleLan200Response) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


