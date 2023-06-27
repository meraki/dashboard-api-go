# UpdateNetworkSwitchRoutingOspfRequestAreasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaId** | **string** | OSPF area ID | 
**AreaName** | **string** | Name of the OSPF area | 
**AreaType** | **string** | Area types in OSPF. Must be one of: [\&quot;normal\&quot;, \&quot;stub\&quot;, \&quot;nssa\&quot;] | 

## Methods

### NewUpdateNetworkSwitchRoutingOspfRequestAreasInner

`func NewUpdateNetworkSwitchRoutingOspfRequestAreasInner(areaId string, areaName string, areaType string, ) *UpdateNetworkSwitchRoutingOspfRequestAreasInner`

NewUpdateNetworkSwitchRoutingOspfRequestAreasInner instantiates a new UpdateNetworkSwitchRoutingOspfRequestAreasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchRoutingOspfRequestAreasInnerWithDefaults

`func NewUpdateNetworkSwitchRoutingOspfRequestAreasInnerWithDefaults() *UpdateNetworkSwitchRoutingOspfRequestAreasInner`

NewUpdateNetworkSwitchRoutingOspfRequestAreasInnerWithDefaults instantiates a new UpdateNetworkSwitchRoutingOspfRequestAreasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaId

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.


### GetAreaName

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaName() string`

GetAreaName returns the AreaName field if non-nil, zero value otherwise.

### GetAreaNameOk

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaNameOk() (*string, bool)`

GetAreaNameOk returns a tuple with the AreaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaName

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) SetAreaName(v string)`

SetAreaName sets AreaName field to given value.


### GetAreaType

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaType() string`

GetAreaType returns the AreaType field if non-nil, zero value otherwise.

### GetAreaTypeOk

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) GetAreaTypeOk() (*string, bool)`

GetAreaTypeOk returns a tuple with the AreaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaType

`func (o *UpdateNetworkSwitchRoutingOspfRequestAreasInner) SetAreaType(v string)`

SetAreaType sets AreaType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


