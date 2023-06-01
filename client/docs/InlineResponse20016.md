# InlineResponse20016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | The subnet of the single LAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the single LAN | [optional] 
**MandatoryDhcp** | Pointer to [**InlineResponse20016MandatoryDhcp**](InlineResponse20016MandatoryDhcp.md) |  | [optional] 
**Ipv6** | Pointer to [**InlineResponse20016Ipv6**](InlineResponse20016Ipv6.md) |  | [optional] 

## Methods

### NewInlineResponse20016

`func NewInlineResponse20016() *InlineResponse20016`

NewInlineResponse20016 instantiates a new InlineResponse20016 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20016WithDefaults

`func NewInlineResponse20016WithDefaults() *InlineResponse20016`

NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *InlineResponse20016) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20016) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20016) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20016) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineResponse20016) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineResponse20016) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineResponse20016) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineResponse20016) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *InlineResponse20016) GetMandatoryDhcp() InlineResponse20016MandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *InlineResponse20016) GetMandatoryDhcpOk() (*InlineResponse20016MandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *InlineResponse20016) SetMandatoryDhcp(v InlineResponse20016MandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *InlineResponse20016) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20016) GetIpv6() InlineResponse20016Ipv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20016) GetIpv6Ok() (*InlineResponse20016Ipv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20016) SetIpv6(v InlineResponse20016Ipv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20016) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


