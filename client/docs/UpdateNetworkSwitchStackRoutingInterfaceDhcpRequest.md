# UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **string** | The DHCP mode options for the switch stack interface (&#39;dhcpDisabled&#39;, &#39;dhcpRelay&#39; or &#39;dhcpServer&#39;) | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The DHCP lease time config for the dhcp server running on switch stack interface (&#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;) | [optional] 
**DnsNameserversOption** | Pointer to **string** | The DHCP name server option for the dhcp server running on the switch stack interface (&#39;googlePublicDns&#39;, &#39;openDns&#39; or &#39;custom&#39;) | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | The DHCP name server IPs when DHCP name server option is &#39;custom&#39; | [optional] 
**BootOptionsEnabled** | Pointer to **bool** | Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface | [optional] 
**BootNextServer** | Pointer to **string** | The PXE boot server IP for the DHCP server running on the switch stack interface | [optional] 
**BootFileName** | Pointer to **string** | The PXE boot server file name for the DHCP server running on the switch stack interface | [optional] 
**DhcpOptions** | Pointer to [**[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner**](UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner.md) | Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface | [optional] 
**ReservedIpRanges** | Pointer to [**[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner**](UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 
**FixedIpAssignments** | Pointer to [**[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner**](UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner.md) | Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface | [optional] 

## Methods

### NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest

`func NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest() *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest`

NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest instantiates a new UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequestWithDefaults

`func NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequestWithDefaults() *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest`

NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequestWithDefaults instantiates a new UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpMode() string`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpModeOk() (*string, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpMode(v string)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameserversOption

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsNameserversOption() string`

GetDnsNameserversOption returns the DnsNameserversOption field if non-nil, zero value otherwise.

### GetDnsNameserversOptionOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsNameserversOptionOk() (*string, bool)`

GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameserversOption

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDnsNameserversOption(v string)`

SetDnsNameserversOption sets DnsNameserversOption field to given value.

### HasDnsNameserversOption

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDnsNameserversOption() bool`

HasDnsNameserversOption returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.

### GetBootOptionsEnabled

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootOptionsEnabled() bool`

GetBootOptionsEnabled returns the BootOptionsEnabled field if non-nil, zero value otherwise.

### GetBootOptionsEnabledOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootOptionsEnabledOk() (*bool, bool)`

GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionsEnabled

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetBootOptionsEnabled(v bool)`

SetBootOptionsEnabled sets BootOptionsEnabled field to given value.

### HasBootOptionsEnabled

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasBootOptionsEnabled() bool`

HasBootOptionsEnabled returns a boolean if a field has been set.

### GetBootNextServer

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootNextServer() string`

GetBootNextServer returns the BootNextServer field if non-nil, zero value otherwise.

### GetBootNextServerOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootNextServerOk() (*string, bool)`

GetBootNextServerOk returns a tuple with the BootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNextServer

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetBootNextServer(v string)`

SetBootNextServer sets BootNextServer field to given value.

### HasBootNextServer

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasBootNextServer() bool`

HasBootNextServer returns a boolean if a field has been set.

### GetBootFileName

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpOptions() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpOptionsOk() (*[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpOptions(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetReservedIpRanges() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetReservedIpRangesOk() (*[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetReservedIpRanges(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetFixedIpAssignments() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetFixedIpAssignmentsOk() (*[]UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetFixedIpAssignments(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


