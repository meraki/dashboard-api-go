# GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autonomous** | Pointer to **bool** | Auto assign a /64 prefix from the origin to the single LAN | [optional] 
**StaticPrefix** | Pointer to **string** | Manual configuration of a /64 prefix on the single LAN | [optional] 
**StaticApplianceIp6** | Pointer to **string** | Manual configuration of the IPv6 Appliance IP | [optional] 
**Origin** | Pointer to [**CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin**](CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin.md) |  | [optional] 

## Methods

### NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner

`func NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner() *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner`

NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner instantiates a new GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInnerWithDefaults

`func NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInnerWithDefaults() *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner`

NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInnerWithDefaults instantiates a new GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonomous

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetAutonomous() bool`

GetAutonomous returns the Autonomous field if non-nil, zero value otherwise.

### GetAutonomousOk

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetAutonomousOk() (*bool, bool)`

GetAutonomousOk returns a tuple with the Autonomous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomous

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetAutonomous(v bool)`

SetAutonomous sets Autonomous field to given value.

### HasAutonomous

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasAutonomous() bool`

HasAutonomous returns a boolean if a field has been set.

### GetStaticPrefix

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticPrefix() string`

GetStaticPrefix returns the StaticPrefix field if non-nil, zero value otherwise.

### GetStaticPrefixOk

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticPrefixOk() (*string, bool)`

GetStaticPrefixOk returns a tuple with the StaticPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticPrefix

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetStaticPrefix(v string)`

SetStaticPrefix sets StaticPrefix field to given value.

### HasStaticPrefix

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasStaticPrefix() bool`

HasStaticPrefix returns a boolean if a field has been set.

### GetStaticApplianceIp6

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticApplianceIp6() string`

GetStaticApplianceIp6 returns the StaticApplianceIp6 field if non-nil, zero value otherwise.

### GetStaticApplianceIp6Ok

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticApplianceIp6Ok() (*string, bool)`

GetStaticApplianceIp6Ok returns a tuple with the StaticApplianceIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticApplianceIp6

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetStaticApplianceIp6(v string)`

SetStaticApplianceIp6 sets StaticApplianceIp6 field to given value.

### HasStaticApplianceIp6

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasStaticApplianceIp6() bool`

HasStaticApplianceIp6 returns a boolean if a field has been set.

### GetOrigin

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetOrigin() CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetOriginOk() (*CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetOrigin(v CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


