# UpdateNetworkSwitchAlternateManagementInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set | [optional] 
**VlanId** | Pointer to **int32** | Alternate management VLAN, must be between 1 and 4094 | [optional] 
**Protocols** | Pointer to **[]string** | Can be one or more of the following values: &#39;radius&#39;, &#39;snmp&#39; or &#39;syslog&#39; | [optional] 
**Switches** | Pointer to [**[]UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner**](UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner.md) | Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put &#39;switches&#39; in the body when updating template networks. Also, an empty &#39;switches&#39; array will remove all previous assignments | [optional] 

## Methods

### NewUpdateNetworkSwitchAlternateManagementInterfaceRequest

`func NewUpdateNetworkSwitchAlternateManagementInterfaceRequest() *UpdateNetworkSwitchAlternateManagementInterfaceRequest`

NewUpdateNetworkSwitchAlternateManagementInterfaceRequest instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchAlternateManagementInterfaceRequestWithDefaults

`func NewUpdateNetworkSwitchAlternateManagementInterfaceRequestWithDefaults() *UpdateNetworkSwitchAlternateManagementInterfaceRequest`

NewUpdateNetworkSwitchAlternateManagementInterfaceRequestWithDefaults instantiates a new UpdateNetworkSwitchAlternateManagementInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetProtocols

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetSwitches

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetSwitches() []UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) GetSwitchesOk() (*[]UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) SetSwitches(v []UpdateNetworkSwitchAlternateManagementInterfaceRequestSwitchesInner)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *UpdateNetworkSwitchAlternateManagementInterfaceRequest) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


