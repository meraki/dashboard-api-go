# UpdateNetworkWirelessAlternateManagementInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable alternate management interface | [optional] 
**VlanId** | Pointer to **int32** | Alternate management interface VLAN, must be between 1 and 4094 | [optional] 
**Protocols** | Pointer to **[]string** | Can be one or more of the following values: &#39;radius&#39;, &#39;snmp&#39;, &#39;syslog&#39; or &#39;ldap&#39; | [optional] 
**AccessPoints** | Pointer to [**[]UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner**](UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner.md) | Array of access point serial number and IP assignment. Note: accessPoints IP assignment is not applicable for template networks, in other words, do not put &#39;accessPoints&#39; in the body when updating template networks. Also, an empty &#39;accessPoints&#39; array will remove all previous static IP assignments | [optional] 

## Methods

### NewUpdateNetworkWirelessAlternateManagementInterfaceRequest

`func NewUpdateNetworkWirelessAlternateManagementInterfaceRequest() *UpdateNetworkWirelessAlternateManagementInterfaceRequest`

NewUpdateNetworkWirelessAlternateManagementInterfaceRequest instantiates a new UpdateNetworkWirelessAlternateManagementInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessAlternateManagementInterfaceRequestWithDefaults

`func NewUpdateNetworkWirelessAlternateManagementInterfaceRequestWithDefaults() *UpdateNetworkWirelessAlternateManagementInterfaceRequest`

NewUpdateNetworkWirelessAlternateManagementInterfaceRequestWithDefaults instantiates a new UpdateNetworkWirelessAlternateManagementInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetProtocols

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetAccessPoints

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetAccessPoints() []UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner`

GetAccessPoints returns the AccessPoints field if non-nil, zero value otherwise.

### GetAccessPointsOk

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) GetAccessPointsOk() (*[]UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner, bool)`

GetAccessPointsOk returns a tuple with the AccessPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoints

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) SetAccessPoints(v []UpdateNetworkWirelessAlternateManagementInterfaceRequestAccessPointsInner)`

SetAccessPoints sets AccessPoints field to given value.

### HasAccessPoints

`func (o *UpdateNetworkWirelessAlternateManagementInterfaceRequest) HasAccessPoints() bool`

HasAccessPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


