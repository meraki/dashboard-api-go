# UpdateDeviceSwitchRoutingInterfaceDhcpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **string** | The DHCP mode options for the switch interface (&#39;dhcpDisabled&#39;, &#39;dhcpRelay&#39; or &#39;dhcpServer&#39;) | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The DHCP relay server IPs to which DHCP packets would get relayed for the switch interface | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The DHCP lease time config for the dhcp server running on switch interface (&#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;) | [optional] 
**DnsNameserversOption** | Pointer to **string** | The DHCP name server option for the dhcp server running on the switch interface (&#39;googlePublicDns&#39;, &#39;openDns&#39; or &#39;custom&#39;) | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | The DHCP name server IPs when DHCP name server option is &#39;custom&#39; | [optional] 
**BootOptionsEnabled** | Pointer to **bool** | Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch interface | [optional] 
**BootNextServer** | Pointer to **string** | The PXE boot server IP for the DHCP server running on the switch interface | [optional] 
**BootFileName** | Pointer to **string** | The PXE boot server filename for the DHCP server running on the switch interface | [optional] 
**DhcpOptions** | Pointer to [**[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner**](UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner.md) | Array of DHCP options consisting of code, type and value for the DHCP server running on the switch interface | [optional] 
**ReservedIpRanges** | Pointer to [**[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner**](UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch interface | [optional] 
**FixedIpAssignments** | Pointer to [**[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner**](UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner.md) | Array of DHCP fixed IP assignments for the DHCP server running on the switch interface | [optional] 

## Methods

### NewUpdateDeviceSwitchRoutingInterfaceDhcpRequest

`func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequest() *UpdateDeviceSwitchRoutingInterfaceDhcpRequest`

NewUpdateDeviceSwitchRoutingInterfaceDhcpRequest instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestWithDefaults

`func NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestWithDefaults() *UpdateDeviceSwitchRoutingInterfaceDhcpRequest`

NewUpdateDeviceSwitchRoutingInterfaceDhcpRequestWithDefaults instantiates a new UpdateDeviceSwitchRoutingInterfaceDhcpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpMode() string`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpModeOk() (*string, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetDhcpMode(v string)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameserversOption

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDnsNameserversOption() string`

GetDnsNameserversOption returns the DnsNameserversOption field if non-nil, zero value otherwise.

### GetDnsNameserversOptionOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDnsNameserversOptionOk() (*string, bool)`

GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameserversOption

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetDnsNameserversOption(v string)`

SetDnsNameserversOption sets DnsNameserversOption field to given value.

### HasDnsNameserversOption

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasDnsNameserversOption() bool`

HasDnsNameserversOption returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.

### GetBootOptionsEnabled

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetBootOptionsEnabled() bool`

GetBootOptionsEnabled returns the BootOptionsEnabled field if non-nil, zero value otherwise.

### GetBootOptionsEnabledOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetBootOptionsEnabledOk() (*bool, bool)`

GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionsEnabled

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetBootOptionsEnabled(v bool)`

SetBootOptionsEnabled sets BootOptionsEnabled field to given value.

### HasBootOptionsEnabled

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasBootOptionsEnabled() bool`

HasBootOptionsEnabled returns a boolean if a field has been set.

### GetBootNextServer

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetBootNextServer() string`

GetBootNextServer returns the BootNextServer field if non-nil, zero value otherwise.

### GetBootNextServerOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetBootNextServerOk() (*string, bool)`

GetBootNextServerOk returns a tuple with the BootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNextServer

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetBootNextServer(v string)`

SetBootNextServer sets BootNextServer field to given value.

### HasBootNextServer

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasBootNextServer() bool`

HasBootNextServer returns a boolean if a field has been set.

### GetBootFileName

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpOptions() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetDhcpOptionsOk() (*[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetDhcpOptions(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetReservedIpRanges() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetReservedIpRangesOk() (*[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetReservedIpRanges(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetFixedIpAssignments() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) GetFixedIpAssignmentsOk() (*[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) SetFixedIpAssignments(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *UpdateDeviceSwitchRoutingInterfaceDhcpRequest) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


