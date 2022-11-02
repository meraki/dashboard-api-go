# CreateDeviceSwitchRoutingStaticRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name or description for layer 3 static route | [optional] 
**Subnet** | **string** | The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24) | 
**NextHopIp** | **string** | IP address of the next hop device to which the device sends its traffic for the subnet | 
**AdvertiseViaOspfEnabled** | Pointer to **bool** | Option to advertise static route via OSPF | [optional] 
**PreferOverOspfRoutesEnabled** | Pointer to **bool** | Option to prefer static route over OSPF routes | [optional] 

## Methods

### NewCreateDeviceSwitchRoutingStaticRouteRequest

`func NewCreateDeviceSwitchRoutingStaticRouteRequest(subnet string, nextHopIp string, ) *CreateDeviceSwitchRoutingStaticRouteRequest`

NewCreateDeviceSwitchRoutingStaticRouteRequest instantiates a new CreateDeviceSwitchRoutingStaticRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceSwitchRoutingStaticRouteRequestWithDefaults

`func NewCreateDeviceSwitchRoutingStaticRouteRequestWithDefaults() *CreateDeviceSwitchRoutingStaticRouteRequest`

NewCreateDeviceSwitchRoutingStaticRouteRequestWithDefaults instantiates a new CreateDeviceSwitchRoutingStaticRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetNextHopIp

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.


### GetAdvertiseViaOspfEnabled

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetAdvertiseViaOspfEnabled() bool`

GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field if non-nil, zero value otherwise.

### GetAdvertiseViaOspfEnabledOk

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetAdvertiseViaOspfEnabledOk() (*bool, bool)`

GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseViaOspfEnabled

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetAdvertiseViaOspfEnabled(v bool)`

SetAdvertiseViaOspfEnabled sets AdvertiseViaOspfEnabled field to given value.

### HasAdvertiseViaOspfEnabled

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) HasAdvertiseViaOspfEnabled() bool`

HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.

### GetPreferOverOspfRoutesEnabled

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetPreferOverOspfRoutesEnabled() bool`

GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field if non-nil, zero value otherwise.

### GetPreferOverOspfRoutesEnabledOk

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetPreferOverOspfRoutesEnabledOk() (*bool, bool)`

GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOverOspfRoutesEnabled

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetPreferOverOspfRoutesEnabled(v bool)`

SetPreferOverOspfRoutesEnabled sets PreferOverOspfRoutesEnabled field to given value.

### HasPreferOverOspfRoutesEnabled

`func (o *CreateDeviceSwitchRoutingStaticRouteRequest) HasPreferOverOspfRoutesEnabled() bool`

HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


