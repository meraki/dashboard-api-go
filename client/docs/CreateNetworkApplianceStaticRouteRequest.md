# CreateNetworkApplianceStaticRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new static route | 
**Subnet** | **string** | The subnet of the static route | 
**GatewayIp** | **string** | The gateway IP (next hop) of the static route | 
**GatewayVlanId** | Pointer to **string** | The gateway IP (next hop) VLAN ID of the static route | [optional] 

## Methods

### NewCreateNetworkApplianceStaticRouteRequest

`func NewCreateNetworkApplianceStaticRouteRequest(name string, subnet string, gatewayIp string, ) *CreateNetworkApplianceStaticRouteRequest`

NewCreateNetworkApplianceStaticRouteRequest instantiates a new CreateNetworkApplianceStaticRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkApplianceStaticRouteRequestWithDefaults

`func NewCreateNetworkApplianceStaticRouteRequestWithDefaults() *CreateNetworkApplianceStaticRouteRequest`

NewCreateNetworkApplianceStaticRouteRequestWithDefaults instantiates a new CreateNetworkApplianceStaticRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkApplianceStaticRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkApplianceStaticRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkApplianceStaticRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *CreateNetworkApplianceStaticRouteRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateNetworkApplianceStaticRouteRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateNetworkApplianceStaticRouteRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetGatewayIp

`func (o *CreateNetworkApplianceStaticRouteRequest) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *CreateNetworkApplianceStaticRouteRequest) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *CreateNetworkApplianceStaticRouteRequest) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.


### GetGatewayVlanId

`func (o *CreateNetworkApplianceStaticRouteRequest) GetGatewayVlanId() string`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *CreateNetworkApplianceStaticRouteRequest) GetGatewayVlanIdOk() (*string, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *CreateNetworkApplianceStaticRouteRequest) SetGatewayVlanId(v string)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *CreateNetworkApplianceStaticRouteRequest) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


