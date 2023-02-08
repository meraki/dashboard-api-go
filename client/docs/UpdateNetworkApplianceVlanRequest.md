# UpdateNetworkApplianceVlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the VLAN | [optional] 
**Subnet** | Pointer to **string** | The subnet of the VLAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the VLAN | [optional] 
**GroupPolicyId** | Pointer to **string** | The id of the desired group policy to apply to the VLAN | [optional] 
**VpnNatSubnet** | Pointer to **string** | The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN | [optional] 
**DhcpHandling** | Pointer to **string** | The appliance&#39;s handling of DHCP requests on this VLAN. One of: &#39;Run a DHCP server&#39;, &#39;Relay DHCP to another server&#39; or &#39;Do not respond to DHCP requests&#39; | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The IPs of the DHCP servers that DHCP requests should be relayed to | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: &#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39; | [optional] 
**DhcpBootOptionsEnabled** | Pointer to **bool** | Use DHCP boot options specified in other properties | [optional] 
**DhcpBootNextServer** | Pointer to **string** | DHCP boot option to direct boot clients to the server to load the boot file from | [optional] 
**DhcpBootFilename** | Pointer to **string** | DHCP boot option for boot filename | [optional] 
**FixedIpAssignments** | Pointer to **map[string]interface{}** | The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain \&quot;ip\&quot; and \&quot;name\&quot; string fields. See the sample request/response for more details. | [optional] 
**ReservedIpRanges** | Pointer to [**[]UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner**](UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner.md) | The DHCP reserved IP ranges on the VLAN | [optional] 
**DnsNameservers** | Pointer to **string** | The DNS nameservers used for DHCP responses, either \&quot;upstream_dns\&quot;, \&quot;google_dns\&quot;, \&quot;opendns\&quot;, or a newline seperated string of IP addresses or domain names | [optional] 
**DhcpOptions** | Pointer to [**[]GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner**](GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner.md) | The list of DHCP options that will be included in DHCP responses. Each object in the list should have \&quot;code\&quot;, \&quot;type\&quot;, and \&quot;value\&quot; properties. | [optional] 
**TemplateVlanType** | Pointer to **string** | Type of subnetting of the VLAN. Applicable only for template network. | [optional] 
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all bound to the template networks. Applicable only for template network. | [optional] 
**Ipv6** | Pointer to [**UpdateNetworkApplianceSingleLanRequestIpv6**](UpdateNetworkApplianceSingleLanRequestIpv6.md) |  | [optional] 
**MandatoryDhcp** | Pointer to [**GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp**](GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp.md) |  | [optional] 

## Methods

### NewUpdateNetworkApplianceVlanRequest

`func NewUpdateNetworkApplianceVlanRequest() *UpdateNetworkApplianceVlanRequest`

NewUpdateNetworkApplianceVlanRequest instantiates a new UpdateNetworkApplianceVlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceVlanRequestWithDefaults

`func NewUpdateNetworkApplianceVlanRequestWithDefaults() *UpdateNetworkApplianceVlanRequest`

NewUpdateNetworkApplianceVlanRequestWithDefaults instantiates a new UpdateNetworkApplianceVlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkApplianceVlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkApplianceVlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkApplianceVlanRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkApplianceVlanRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *UpdateNetworkApplianceVlanRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateNetworkApplianceVlanRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateNetworkApplianceVlanRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateNetworkApplianceVlanRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *UpdateNetworkApplianceVlanRequest) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *UpdateNetworkApplianceVlanRequest) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *UpdateNetworkApplianceVlanRequest) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *UpdateNetworkApplianceVlanRequest) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *UpdateNetworkApplianceVlanRequest) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *UpdateNetworkApplianceVlanRequest) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *UpdateNetworkApplianceVlanRequest) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *UpdateNetworkApplianceVlanRequest) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetVpnNatSubnet

`func (o *UpdateNetworkApplianceVlanRequest) GetVpnNatSubnet() string`

GetVpnNatSubnet returns the VpnNatSubnet field if non-nil, zero value otherwise.

### GetVpnNatSubnetOk

`func (o *UpdateNetworkApplianceVlanRequest) GetVpnNatSubnetOk() (*string, bool)`

GetVpnNatSubnetOk returns a tuple with the VpnNatSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnNatSubnet

`func (o *UpdateNetworkApplianceVlanRequest) SetVpnNatSubnet(v string)`

SetVpnNatSubnet sets VpnNatSubnet field to given value.

### HasVpnNatSubnet

`func (o *UpdateNetworkApplianceVlanRequest) HasVpnNatSubnet() bool`

HasVpnNatSubnet returns a boolean if a field has been set.

### GetDhcpHandling

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpHandling() string`

GetDhcpHandling returns the DhcpHandling field if non-nil, zero value otherwise.

### GetDhcpHandlingOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpHandlingOk() (*string, bool)`

GetDhcpHandlingOk returns a tuple with the DhcpHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHandling

`func (o *UpdateNetworkApplianceVlanRequest) SetDhcpHandling(v string)`

SetDhcpHandling sets DhcpHandling field to given value.

### HasDhcpHandling

`func (o *UpdateNetworkApplianceVlanRequest) HasDhcpHandling() bool`

HasDhcpHandling returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *UpdateNetworkApplianceVlanRequest) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *UpdateNetworkApplianceVlanRequest) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *UpdateNetworkApplianceVlanRequest) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *UpdateNetworkApplianceVlanRequest) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDhcpBootOptionsEnabled

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpBootOptionsEnabled() bool`

GetDhcpBootOptionsEnabled returns the DhcpBootOptionsEnabled field if non-nil, zero value otherwise.

### GetDhcpBootOptionsEnabledOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpBootOptionsEnabledOk() (*bool, bool)`

GetDhcpBootOptionsEnabledOk returns a tuple with the DhcpBootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootOptionsEnabled

`func (o *UpdateNetworkApplianceVlanRequest) SetDhcpBootOptionsEnabled(v bool)`

SetDhcpBootOptionsEnabled sets DhcpBootOptionsEnabled field to given value.

### HasDhcpBootOptionsEnabled

`func (o *UpdateNetworkApplianceVlanRequest) HasDhcpBootOptionsEnabled() bool`

HasDhcpBootOptionsEnabled returns a boolean if a field has been set.

### GetDhcpBootNextServer

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpBootNextServer() string`

GetDhcpBootNextServer returns the DhcpBootNextServer field if non-nil, zero value otherwise.

### GetDhcpBootNextServerOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpBootNextServerOk() (*string, bool)`

GetDhcpBootNextServerOk returns a tuple with the DhcpBootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootNextServer

`func (o *UpdateNetworkApplianceVlanRequest) SetDhcpBootNextServer(v string)`

SetDhcpBootNextServer sets DhcpBootNextServer field to given value.

### HasDhcpBootNextServer

`func (o *UpdateNetworkApplianceVlanRequest) HasDhcpBootNextServer() bool`

HasDhcpBootNextServer returns a boolean if a field has been set.

### GetDhcpBootFilename

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpBootFilename() string`

GetDhcpBootFilename returns the DhcpBootFilename field if non-nil, zero value otherwise.

### GetDhcpBootFilenameOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpBootFilenameOk() (*string, bool)`

GetDhcpBootFilenameOk returns a tuple with the DhcpBootFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootFilename

`func (o *UpdateNetworkApplianceVlanRequest) SetDhcpBootFilename(v string)`

SetDhcpBootFilename sets DhcpBootFilename field to given value.

### HasDhcpBootFilename

`func (o *UpdateNetworkApplianceVlanRequest) HasDhcpBootFilename() bool`

HasDhcpBootFilename returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *UpdateNetworkApplianceVlanRequest) GetFixedIpAssignments() map[string]interface{}`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *UpdateNetworkApplianceVlanRequest) GetFixedIpAssignmentsOk() (*map[string]interface{}, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *UpdateNetworkApplianceVlanRequest) SetFixedIpAssignments(v map[string]interface{})`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *UpdateNetworkApplianceVlanRequest) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *UpdateNetworkApplianceVlanRequest) GetReservedIpRanges() []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *UpdateNetworkApplianceVlanRequest) GetReservedIpRangesOk() (*[]UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *UpdateNetworkApplianceVlanRequest) SetReservedIpRanges(v []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *UpdateNetworkApplianceVlanRequest) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetDnsNameservers

`func (o *UpdateNetworkApplianceVlanRequest) GetDnsNameservers() string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDnsNameserversOk() (*string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *UpdateNetworkApplianceVlanRequest) SetDnsNameservers(v string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *UpdateNetworkApplianceVlanRequest) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpOptions() []GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *UpdateNetworkApplianceVlanRequest) GetDhcpOptionsOk() (*[]GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *UpdateNetworkApplianceVlanRequest) SetDhcpOptions(v []GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *UpdateNetworkApplianceVlanRequest) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetTemplateVlanType

`func (o *UpdateNetworkApplianceVlanRequest) GetTemplateVlanType() string`

GetTemplateVlanType returns the TemplateVlanType field if non-nil, zero value otherwise.

### GetTemplateVlanTypeOk

`func (o *UpdateNetworkApplianceVlanRequest) GetTemplateVlanTypeOk() (*string, bool)`

GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVlanType

`func (o *UpdateNetworkApplianceVlanRequest) SetTemplateVlanType(v string)`

SetTemplateVlanType sets TemplateVlanType field to given value.

### HasTemplateVlanType

`func (o *UpdateNetworkApplianceVlanRequest) HasTemplateVlanType() bool`

HasTemplateVlanType returns a boolean if a field has been set.

### GetCidr

`func (o *UpdateNetworkApplianceVlanRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateNetworkApplianceVlanRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateNetworkApplianceVlanRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateNetworkApplianceVlanRequest) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *UpdateNetworkApplianceVlanRequest) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *UpdateNetworkApplianceVlanRequest) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *UpdateNetworkApplianceVlanRequest) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *UpdateNetworkApplianceVlanRequest) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetIpv6

`func (o *UpdateNetworkApplianceVlanRequest) GetIpv6() UpdateNetworkApplianceSingleLanRequestIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *UpdateNetworkApplianceVlanRequest) GetIpv6Ok() (*UpdateNetworkApplianceSingleLanRequestIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *UpdateNetworkApplianceVlanRequest) SetIpv6(v UpdateNetworkApplianceSingleLanRequestIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *UpdateNetworkApplianceVlanRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *UpdateNetworkApplianceVlanRequest) GetMandatoryDhcp() GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *UpdateNetworkApplianceVlanRequest) GetMandatoryDhcpOk() (*GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *UpdateNetworkApplianceVlanRequest) SetMandatoryDhcp(v GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *UpdateNetworkApplianceVlanRequest) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


