# UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | The IP protocol used for the address | [optional] 
**AssignmentMode** | Pointer to **string** | The type of address assignment. Either static or dynamic. | [optional] 
**Address** | Pointer to **string** | The IP address configured for the alternate management interface | [optional] 
**Gateway** | Pointer to **string** | The gateway address configured for the alternate managment interface | [optional] 
**Prefix** | Pointer to **string** | The IPv6 prefix of the interface. Required if IPv6 object is included. | [optional] 
**Nameservers** | Pointer to [**UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers**](UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers.md) |  | [optional] 

## Methods

### NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner

`func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner`

NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInnerWithDefaults

`func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInnerWithDefaults() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner`

NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInnerWithDefaults instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAssignmentMode

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAssignmentMode() string`

GetAssignmentMode returns the AssignmentMode field if non-nil, zero value otherwise.

### GetAssignmentModeOk

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAssignmentModeOk() (*string, bool)`

GetAssignmentModeOk returns a tuple with the AssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentMode

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetAssignmentMode(v string)`

SetAssignmentMode sets AssignmentMode field to given value.

### HasAssignmentMode

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasAssignmentMode() bool`

HasAssignmentMode returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPrefix

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNameservers

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetNameservers() UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetNameserversOk() (*UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetNameservers(v UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


