# GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autonomous** | Pointer to **bool** | Auto assign a /64 prefix from the origin to the VLAN | [optional] 
**StaticPrefix** | Pointer to **string** | Manual configuration of a /64 prefix on the VLAN | [optional] 
**StaticApplianceIp6** | Pointer to **string** | Manual configuration of the IPv6 Appliance IP | [optional] 
**Origin** | Pointer to [**CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin**](CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin.md) |  | [optional] 

## Methods

### NewGetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner

`func NewGetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner() *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner`

NewGetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner instantiates a new GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInnerWithDefaults

`func NewGetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInnerWithDefaults() *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner`

NewGetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInnerWithDefaults instantiates a new GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonomous

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetAutonomous() bool`

GetAutonomous returns the Autonomous field if non-nil, zero value otherwise.

### GetAutonomousOk

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetAutonomousOk() (*bool, bool)`

GetAutonomousOk returns a tuple with the Autonomous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomous

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) SetAutonomous(v bool)`

SetAutonomous sets Autonomous field to given value.

### HasAutonomous

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) HasAutonomous() bool`

HasAutonomous returns a boolean if a field has been set.

### GetStaticPrefix

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetStaticPrefix() string`

GetStaticPrefix returns the StaticPrefix field if non-nil, zero value otherwise.

### GetStaticPrefixOk

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetStaticPrefixOk() (*string, bool)`

GetStaticPrefixOk returns a tuple with the StaticPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticPrefix

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) SetStaticPrefix(v string)`

SetStaticPrefix sets StaticPrefix field to given value.

### HasStaticPrefix

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) HasStaticPrefix() bool`

HasStaticPrefix returns a boolean if a field has been set.

### GetStaticApplianceIp6

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetStaticApplianceIp6() string`

GetStaticApplianceIp6 returns the StaticApplianceIp6 field if non-nil, zero value otherwise.

### GetStaticApplianceIp6Ok

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetStaticApplianceIp6Ok() (*string, bool)`

GetStaticApplianceIp6Ok returns a tuple with the StaticApplianceIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticApplianceIp6

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) SetStaticApplianceIp6(v string)`

SetStaticApplianceIp6 sets StaticApplianceIp6 field to given value.

### HasStaticApplianceIp6

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) HasStaticApplianceIp6() bool`

HasStaticApplianceIp6 returns a boolean if a field has been set.

### GetOrigin

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetOrigin() CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) GetOriginOk() (*CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) SetOrigin(v CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


