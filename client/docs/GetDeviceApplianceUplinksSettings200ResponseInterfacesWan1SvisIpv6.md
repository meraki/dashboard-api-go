# GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentMode** | Pointer to **string** | The assignment mode for this SVI. Applies only when PPPoE is disabled. | [optional] 
**Address** | Pointer to **string** | Static address that will override the one(s) received by SLAAC. | [optional] 
**Gateway** | Pointer to **string** | Static gateway that will override the one received by autoconf. | [optional] 
**Nameservers** | Pointer to [**GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers**](GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers.md) |  | [optional] 

## Methods

### NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6

`func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6`

NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6WithDefaults

`func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6WithDefaults() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6`

NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6WithDefaults instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentMode

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAssignmentMode() string`

GetAssignmentMode returns the AssignmentMode field if non-nil, zero value otherwise.

### GetAssignmentModeOk

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAssignmentModeOk() (*string, bool)`

GetAssignmentModeOk returns a tuple with the AssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentMode

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetAssignmentMode(v string)`

SetAssignmentMode sets AssignmentMode field to given value.

### HasAssignmentMode

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasAssignmentMode() bool`

HasAssignmentMode returns a boolean if a field has been set.

### GetAddress

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNameservers

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetNameservers() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) GetNameserversOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) SetNameservers(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv6) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


