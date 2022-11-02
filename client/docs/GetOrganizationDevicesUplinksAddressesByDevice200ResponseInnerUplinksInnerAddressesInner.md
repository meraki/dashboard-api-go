# GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Type of address for the device uplink. Available options are: ipv4, ipv6. | [optional] 
**AssignmentMode** | Pointer to **string** | Indicates how the device uplink address is assigned. Available options are: static, dynamic. | [optional] 
**Address** | Pointer to **string** | Device uplink address. | [optional] 
**Gateway** | Pointer to **string** | Device uplink gateway address. | [optional] 
**Public** | Pointer to [**GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic**](GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic.md) |  | [optional] 

## Methods

### NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner

`func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner`

NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerWithDefaults

`func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerWithDefaults() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner`

NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAssignmentMode

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAssignmentMode() string`

GetAssignmentMode returns the AssignmentMode field if non-nil, zero value otherwise.

### GetAssignmentModeOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAssignmentModeOk() (*string, bool)`

GetAssignmentModeOk returns a tuple with the AssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentMode

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetAssignmentMode(v string)`

SetAssignmentMode sets AssignmentMode field to given value.

### HasAssignmentMode

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasAssignmentMode() bool`

HasAssignmentMode returns a boolean if a field has been set.

### GetAddress

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPublic

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetPublic() GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetPublicOk() (*GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetPublic(v GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


