# UpdateNetworkApplianceVpnSiteToSiteVpnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The site-to-site VPN mode. Can be one of &#39;none&#39;, &#39;spoke&#39; or &#39;hub&#39; | 
**Hubs** | Pointer to [**[]UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner**](UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner.md) | The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required. | [optional] 
**Subnets** | Pointer to [**[]UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner**](UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner.md) | The list of subnets and their VPN presence. | [optional] 

## Methods

### NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest

`func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest(mode string, ) *UpdateNetworkApplianceVpnSiteToSiteVpnRequest`

NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestWithDefaults

`func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestWithDefaults() *UpdateNetworkApplianceVpnSiteToSiteVpnRequest`

NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestWithDefaults instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetHubs

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetHubs() []UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetHubsOk() (*[]UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) SetHubs(v []UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetSubnets() []UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetSubnetsOk() (*[]UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) SetSubnets(v []UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


