# GetNetworkApplianceVpnSiteToSiteVpn200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner**](GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner**](GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner.md) | The list of subnets and their VPN presence. | [optional] 

## Methods

### NewGetNetworkApplianceVpnSiteToSiteVpn200Response

`func NewGetNetworkApplianceVpnSiteToSiteVpn200Response() *GetNetworkApplianceVpnSiteToSiteVpn200Response`

NewGetNetworkApplianceVpnSiteToSiteVpn200Response instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseWithDefaults

`func NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseWithDefaults() *GetNetworkApplianceVpnSiteToSiteVpn200Response`

NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseWithDefaults instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetHubs() []GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetHubsOk() (*[]GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) SetHubs(v []GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetSubnets() []GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) GetSubnetsOk() (*[]GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) SetSubnets(v []GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GetNetworkApplianceVpnSiteToSiteVpn200Response) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


