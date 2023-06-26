# UpdateNetworkApplianceStaticRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the static route | [optional] 
**Subnet** | Pointer to **string** | The subnet of the static route | [optional] 
**GatewayIp** | Pointer to **string** | The gateway IP (next hop) of the static route | [optional] 
**GatewayVlanId** | Pointer to **string** | The gateway IP (next hop) VLAN ID of the static route | [optional] 
**Enabled** | Pointer to **bool** | The enabled state of the static route | [optional] 
**FixedIpAssignments** | Pointer to **map[string]interface{}** | The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain \&quot;ip\&quot; and \&quot;name\&quot; string fields. See the sample request/response for more details. | [optional] 
**ReservedIpRanges** | Pointer to [**[]UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner**](UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner.md) | The DHCP reserved IP ranges on the static route | [optional] 

## Methods

### NewUpdateNetworkApplianceStaticRouteRequest

`func NewUpdateNetworkApplianceStaticRouteRequest() *UpdateNetworkApplianceStaticRouteRequest`

NewUpdateNetworkApplianceStaticRouteRequest instantiates a new UpdateNetworkApplianceStaticRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceStaticRouteRequestWithDefaults

`func NewUpdateNetworkApplianceStaticRouteRequestWithDefaults() *UpdateNetworkApplianceStaticRouteRequest`

NewUpdateNetworkApplianceStaticRouteRequestWithDefaults instantiates a new UpdateNetworkApplianceStaticRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkApplianceStaticRouteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkApplianceStaticRouteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateNetworkApplianceStaticRouteRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateNetworkApplianceStaticRouteRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGatewayIp

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *UpdateNetworkApplianceStaticRouteRequest) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *UpdateNetworkApplianceStaticRouteRequest) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayVlanId

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayVlanId() string`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetGatewayVlanIdOk() (*string, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *UpdateNetworkApplianceStaticRouteRequest) SetGatewayVlanId(v string)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *UpdateNetworkApplianceStaticRouteRequest) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkApplianceStaticRouteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkApplianceStaticRouteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetFixedIpAssignments() map[string]interface{}`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetFixedIpAssignmentsOk() (*map[string]interface{}, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *UpdateNetworkApplianceStaticRouteRequest) SetFixedIpAssignments(v map[string]interface{})`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *UpdateNetworkApplianceStaticRouteRequest) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetReservedIpRanges() []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *UpdateNetworkApplianceStaticRouteRequest) GetReservedIpRangesOk() (*[]UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *UpdateNetworkApplianceStaticRouteRequest) SetReservedIpRanges(v []UpdateNetworkApplianceStaticRouteRequestReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *UpdateNetworkApplianceStaticRouteRequest) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


