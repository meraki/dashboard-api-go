# UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSubnet** | **string** | The CIDR notation subnet used within the VPN | 
**UseVpn** | Pointer to **bool** | Indicates the presence of the subnet in the VPN | [optional] 

## Methods

### NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner

`func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner(localSubnet string, ) *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner`

NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInnerWithDefaults

`func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInnerWithDefaults() *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner`

NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInnerWithDefaults instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSubnet

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.


### GetUseVpn

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetUseVpn() bool`

GetUseVpn returns the UseVpn field if non-nil, zero value otherwise.

### GetUseVpnOk

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) GetUseVpnOk() (*bool, bool)`

GetUseVpnOk returns a tuple with the UseVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVpn

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) SetUseVpn(v bool)`

SetUseVpn sets UseVpn field to given value.

### HasUseVpn

`func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) HasUseVpn() bool`

HasUseVpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


