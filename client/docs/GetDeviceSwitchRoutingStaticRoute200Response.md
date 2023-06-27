# GetDeviceSwitchRoutingStaticRoute200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticRouteId** | Pointer to **string** | The identifier of a layer 3 static route | [optional] 
**Name** | Pointer to **string** | The name or description of the layer 3 static route | [optional] 
**Subnet** | **string** | The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24) | 
**NextHopIp** | **string** |  The IP address of the router to which traffic for this destination network should be sent | 
**AdvertiseViaOspfEnabled** | Pointer to **bool** | Option to advertise static routes via OSPF | [optional] 
**PreferOverOspfRoutesEnabled** | Pointer to **bool** | Option to prefer static routes over OSPF routes | [optional] 

## Methods

### NewGetDeviceSwitchRoutingStaticRoute200Response

`func NewGetDeviceSwitchRoutingStaticRoute200Response(subnet string, nextHopIp string, ) *GetDeviceSwitchRoutingStaticRoute200Response`

NewGetDeviceSwitchRoutingStaticRoute200Response instantiates a new GetDeviceSwitchRoutingStaticRoute200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchRoutingStaticRoute200ResponseWithDefaults

`func NewGetDeviceSwitchRoutingStaticRoute200ResponseWithDefaults() *GetDeviceSwitchRoutingStaticRoute200Response`

NewGetDeviceSwitchRoutingStaticRoute200ResponseWithDefaults instantiates a new GetDeviceSwitchRoutingStaticRoute200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticRouteId

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetStaticRouteId() string`

GetStaticRouteId returns the StaticRouteId field if non-nil, zero value otherwise.

### GetStaticRouteIdOk

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetStaticRouteIdOk() (*string, bool)`

GetStaticRouteIdOk returns a tuple with the StaticRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRouteId

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetStaticRouteId(v string)`

SetStaticRouteId sets StaticRouteId field to given value.

### HasStaticRouteId

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasStaticRouteId() bool`

HasStaticRouteId returns a boolean if a field has been set.

### GetName

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetNextHopIp

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.


### GetAdvertiseViaOspfEnabled

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetAdvertiseViaOspfEnabled() bool`

GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field if non-nil, zero value otherwise.

### GetAdvertiseViaOspfEnabledOk

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetAdvertiseViaOspfEnabledOk() (*bool, bool)`

GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseViaOspfEnabled

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetAdvertiseViaOspfEnabled(v bool)`

SetAdvertiseViaOspfEnabled sets AdvertiseViaOspfEnabled field to given value.

### HasAdvertiseViaOspfEnabled

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasAdvertiseViaOspfEnabled() bool`

HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.

### GetPreferOverOspfRoutesEnabled

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetPreferOverOspfRoutesEnabled() bool`

GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field if non-nil, zero value otherwise.

### GetPreferOverOspfRoutesEnabledOk

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) GetPreferOverOspfRoutesEnabledOk() (*bool, bool)`

GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOverOspfRoutesEnabled

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) SetPreferOverOspfRoutesEnabled(v bool)`

SetPreferOverOspfRoutesEnabled sets PreferOverOspfRoutesEnabled field to given value.

### HasPreferOverOspfRoutesEnabled

`func (o *GetDeviceSwitchRoutingStaticRoute200Response) HasPreferOverOspfRoutesEnabled() bool`

HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


