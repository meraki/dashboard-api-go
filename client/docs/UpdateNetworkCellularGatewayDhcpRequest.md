# UpdateNetworkCellularGatewayDhcpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpLeaseTime** | Pointer to **string** | DHCP Lease time for all MG of the network. It can be &#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;. | [optional] 
**DnsNameservers** | Pointer to **string** | DNS name servers mode for all MG of the network. It can take 4 different values: &#39;upstream_dns&#39;, &#39;google_dns&#39;, &#39;opendns&#39;, &#39;custom&#39;. | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | list of fixed IP representing the the DNS Name servers when the mode is &#39;custom&#39; | [optional] 

## Methods

### NewUpdateNetworkCellularGatewayDhcpRequest

`func NewUpdateNetworkCellularGatewayDhcpRequest() *UpdateNetworkCellularGatewayDhcpRequest`

NewUpdateNetworkCellularGatewayDhcpRequest instantiates a new UpdateNetworkCellularGatewayDhcpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkCellularGatewayDhcpRequestWithDefaults

`func NewUpdateNetworkCellularGatewayDhcpRequestWithDefaults() *UpdateNetworkCellularGatewayDhcpRequest`

NewUpdateNetworkCellularGatewayDhcpRequestWithDefaults instantiates a new UpdateNetworkCellularGatewayDhcpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpLeaseTime

`func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *UpdateNetworkCellularGatewayDhcpRequest) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *UpdateNetworkCellularGatewayDhcpRequest) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameservers

`func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsNameservers() string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsNameserversOk() (*string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *UpdateNetworkCellularGatewayDhcpRequest) SetDnsNameservers(v string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *UpdateNetworkCellularGatewayDhcpRequest) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *UpdateNetworkCellularGatewayDhcpRequest) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *UpdateNetworkCellularGatewayDhcpRequest) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *UpdateNetworkCellularGatewayDhcpRequest) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


