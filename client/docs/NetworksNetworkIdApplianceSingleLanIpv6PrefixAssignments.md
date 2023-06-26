# NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autonomous** | Pointer to **bool** | Auto assign a /64 prefix from the origin to the VLAN | [optional] 
**StaticPrefix** | Pointer to **string** | Manual configuration of a /64 prefix on the VLAN | [optional] 
**StaticApplianceIp6** | Pointer to **string** | Manual configuration of the IPv6 Appliance IP | [optional] 
**Origin** | Pointer to [**NetworksNetworkIdApplianceSingleLanIpv6Origin**](NetworksNetworkIdApplianceSingleLanIpv6Origin.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments

`func NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments() *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments`

NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments instantiates a new NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignmentsWithDefaults

`func NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignmentsWithDefaults() *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments`

NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignmentsWithDefaults instantiates a new NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonomous

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetAutonomous() bool`

GetAutonomous returns the Autonomous field if non-nil, zero value otherwise.

### GetAutonomousOk

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetAutonomousOk() (*bool, bool)`

GetAutonomousOk returns a tuple with the Autonomous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomous

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetAutonomous(v bool)`

SetAutonomous sets Autonomous field to given value.

### HasAutonomous

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasAutonomous() bool`

HasAutonomous returns a boolean if a field has been set.

### GetStaticPrefix

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticPrefix() string`

GetStaticPrefix returns the StaticPrefix field if non-nil, zero value otherwise.

### GetStaticPrefixOk

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticPrefixOk() (*string, bool)`

GetStaticPrefixOk returns a tuple with the StaticPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticPrefix

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetStaticPrefix(v string)`

SetStaticPrefix sets StaticPrefix field to given value.

### HasStaticPrefix

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasStaticPrefix() bool`

HasStaticPrefix returns a boolean if a field has been set.

### GetStaticApplianceIp6

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticApplianceIp6() string`

GetStaticApplianceIp6 returns the StaticApplianceIp6 field if non-nil, zero value otherwise.

### GetStaticApplianceIp6Ok

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticApplianceIp6Ok() (*string, bool)`

GetStaticApplianceIp6Ok returns a tuple with the StaticApplianceIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticApplianceIp6

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetStaticApplianceIp6(v string)`

SetStaticApplianceIp6 sets StaticApplianceIp6 field to given value.

### HasStaticApplianceIp6

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasStaticApplianceIp6() bool`

HasStaticApplianceIp6 returns a boolean if a field has been set.

### GetOrigin

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetOrigin() NetworksNetworkIdApplianceSingleLanIpv6Origin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetOriginOk() (*NetworksNetworkIdApplianceSingleLanIpv6Origin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetOrigin(v NetworksNetworkIdApplianceSingleLanIpv6Origin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


