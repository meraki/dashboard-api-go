# DevicesSerialSwitchRoutingInterfacesIpv61

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentMode** | Pointer to **string** | The IPv6 assignment mode for the interface. Can be either &#39;eui-64&#39; or &#39;static&#39;. | [optional] 
**Prefix** | Pointer to **string** | The IPv6 prefix of the interface. Required if IPv6 object is included. | [optional] 
**Address** | Pointer to **string** | The IPv6 address of the interface. Required if assignmentMode is &#39;static&#39;. Must not be included if           assignmentMode is &#39;eui-64&#39;. | [optional] 
**Gateway** | Pointer to **string** | The IPv6 default gateway of the interface. Required if prefix is defined and this is the first           interface with IPv6 configured for the switch. | [optional] 

## Methods

### NewDevicesSerialSwitchRoutingInterfacesIpv61

`func NewDevicesSerialSwitchRoutingInterfacesIpv61() *DevicesSerialSwitchRoutingInterfacesIpv61`

NewDevicesSerialSwitchRoutingInterfacesIpv61 instantiates a new DevicesSerialSwitchRoutingInterfacesIpv61 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchRoutingInterfacesIpv61WithDefaults

`func NewDevicesSerialSwitchRoutingInterfacesIpv61WithDefaults() *DevicesSerialSwitchRoutingInterfacesIpv61`

NewDevicesSerialSwitchRoutingInterfacesIpv61WithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesIpv61 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentMode

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetAssignmentMode() string`

GetAssignmentMode returns the AssignmentMode field if non-nil, zero value otherwise.

### GetAssignmentModeOk

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetAssignmentModeOk() (*string, bool)`

GetAssignmentModeOk returns a tuple with the AssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentMode

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) SetAssignmentMode(v string)`

SetAssignmentMode sets AssignmentMode field to given value.

### HasAssignmentMode

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) HasAssignmentMode() bool`

HasAssignmentMode returns a boolean if a field has been set.

### GetPrefix

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAddress

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DevicesSerialSwitchRoutingInterfacesIpv61) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


