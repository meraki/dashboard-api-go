# GetNetworkSwitchStackRoutingInterfaceDhcp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **string** | The DHCP mode options for the switch stack interface (&#39;dhcpDisabled&#39;, &#39;dhcpRelay&#39; or &#39;dhcpServer&#39;) | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The DHCP lease time config for the dhcp server running on the switch stack interface (&#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;) | [optional] 
**DnsNameserversOption** | Pointer to **string** | The DHCP name server option for the dhcp server running on the switch stack interface (&#39;googlePublicDns&#39;, &#39;openDns&#39; or &#39;custom&#39;) | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | The DHCP name server IPs when DHCP name server option is &#39;custom&#39; | [optional] 
**BootOptionsEnabled** | Pointer to **bool** | Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface | [optional] 
**BootNextServer** | Pointer to **string** | The PXE boot server IP for the DHCP server running on the switch stack interface | [optional] 
**BootFileName** | Pointer to **string** | The PXE boot server file name for the DHCP server running on the switch stack interface | [optional] 
**DhcpOptions** | Pointer to [**[]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner**](GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner.md) | Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface | [optional] 
**ReservedIpRanges** | Pointer to [**[]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner**](GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 
**FixedIpAssignments** | Pointer to [**[]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner**](GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 

## Methods

### NewGetNetworkSwitchStackRoutingInterfaceDhcp200Response

`func NewGetNetworkSwitchStackRoutingInterfaceDhcp200Response() *GetNetworkSwitchStackRoutingInterfaceDhcp200Response`

NewGetNetworkSwitchStackRoutingInterfaceDhcp200Response instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseWithDefaults

`func NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseWithDefaults() *GetNetworkSwitchStackRoutingInterfaceDhcp200Response`

NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseWithDefaults instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpMode() string`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpModeOk() (*string, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDhcpMode(v string)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameserversOption

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsNameserversOption() string`

GetDnsNameserversOption returns the DnsNameserversOption field if non-nil, zero value otherwise.

### GetDnsNameserversOptionOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsNameserversOptionOk() (*string, bool)`

GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameserversOption

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDnsNameserversOption(v string)`

SetDnsNameserversOption sets DnsNameserversOption field to given value.

### HasDnsNameserversOption

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDnsNameserversOption() bool`

HasDnsNameserversOption returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.

### GetBootOptionsEnabled

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootOptionsEnabled() bool`

GetBootOptionsEnabled returns the BootOptionsEnabled field if non-nil, zero value otherwise.

### GetBootOptionsEnabledOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootOptionsEnabledOk() (*bool, bool)`

GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionsEnabled

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetBootOptionsEnabled(v bool)`

SetBootOptionsEnabled sets BootOptionsEnabled field to given value.

### HasBootOptionsEnabled

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasBootOptionsEnabled() bool`

HasBootOptionsEnabled returns a boolean if a field has been set.

### GetBootNextServer

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootNextServer() string`

GetBootNextServer returns the BootNextServer field if non-nil, zero value otherwise.

### GetBootNextServerOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootNextServerOk() (*string, bool)`

GetBootNextServerOk returns a tuple with the BootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNextServer

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetBootNextServer(v string)`

SetBootNextServer sets BootNextServer field to given value.

### HasBootNextServer

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasBootNextServer() bool`

HasBootNextServer returns a boolean if a field has been set.

### GetBootFileName

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpOptions() []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetDhcpOptionsOk() (*[]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetDhcpOptions(v []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseDhcpOptionsInner)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetReservedIpRanges() []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetReservedIpRangesOk() (*[]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetReservedIpRanges(v []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetFixedIpAssignments() []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) GetFixedIpAssignmentsOk() (*[]GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) SetFixedIpAssignments(v []GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200Response) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


