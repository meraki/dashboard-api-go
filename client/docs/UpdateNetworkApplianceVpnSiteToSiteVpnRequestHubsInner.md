# UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HubId** | **string** | The network ID of the hub. | 
**UseDefaultRoute** | Pointer to **bool** | Only valid in &#39;spoke&#39; mode. Indicates whether default route traffic should be sent to this hub. | [optional] 

## Methods

### NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner

`func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner(hubId string, ) *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner`

NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInnerWithDefaults

`func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInnerWithDefaults() *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner`

NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInnerWithDefaults instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHubId

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetHubId() string`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetHubIdOk() (*string, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubId

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) SetHubId(v string)`

SetHubId sets HubId field to given value.


### GetUseDefaultRoute

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetUseDefaultRoute() bool`

GetUseDefaultRoute returns the UseDefaultRoute field if non-nil, zero value otherwise.

### GetUseDefaultRouteOk

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetUseDefaultRouteOk() (*bool, bool)`

GetUseDefaultRouteOk returns a tuple with the UseDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultRoute

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) SetUseDefaultRoute(v bool)`

SetUseDefaultRoute sets UseDefaultRoute field to given value.

### HasUseDefaultRoute

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) HasUseDefaultRoute() bool`

HasUseDefaultRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


