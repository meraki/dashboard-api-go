# GetNetworkCellularGatewayDhcp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpLeaseTime** | Pointer to **string** | DHCP Lease time for all MG in the network. | [optional] 
**DnsNameservers** | Pointer to **string** | DNS name servers mode for all MG in the network. | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | List of fixed IPs representing the the DNS Name servers when the mode is &#39;custom&#39;. | [optional] 

## Methods

### NewGetNetworkCellularGatewayDhcp200Response

`func NewGetNetworkCellularGatewayDhcp200Response() *GetNetworkCellularGatewayDhcp200Response`

NewGetNetworkCellularGatewayDhcp200Response instantiates a new GetNetworkCellularGatewayDhcp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkCellularGatewayDhcp200ResponseWithDefaults

`func NewGetNetworkCellularGatewayDhcp200ResponseWithDefaults() *GetNetworkCellularGatewayDhcp200Response`

NewGetNetworkCellularGatewayDhcp200ResponseWithDefaults instantiates a new GetNetworkCellularGatewayDhcp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpLeaseTime

`func (o *GetNetworkCellularGatewayDhcp200Response) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *GetNetworkCellularGatewayDhcp200Response) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *GetNetworkCellularGatewayDhcp200Response) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *GetNetworkCellularGatewayDhcp200Response) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameservers

`func (o *GetNetworkCellularGatewayDhcp200Response) GetDnsNameservers() string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *GetNetworkCellularGatewayDhcp200Response) GetDnsNameserversOk() (*string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *GetNetworkCellularGatewayDhcp200Response) SetDnsNameservers(v string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *GetNetworkCellularGatewayDhcp200Response) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *GetNetworkCellularGatewayDhcp200Response) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *GetNetworkCellularGatewayDhcp200Response) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *GetNetworkCellularGatewayDhcp200Response) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *GetNetworkCellularGatewayDhcp200Response) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


