# DevicesSerialCellularGatewayPortForwardingRulesRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**LanIp** | **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | 
**PublicPort** | **string** | A port or port ranges that will be forwarded to the host on the LAN | 
**LocalPort** | **string** | A port or port ranges that will receive the forwarded traffic from the WAN | 
**AllowedIps** | Pointer to **[]string** | An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges. | [optional] 
**Protocol** | **string** | TCP or UDP | 
**Access** | **string** | &#x60;any&#x60; or &#x60;restricted&#x60;. Specify the right to make inbound connections on the specified ports or port ranges. If &#x60;restricted&#x60;, a list of allowed IPs is mandatory. | 

## Methods

### NewDevicesSerialCellularGatewayPortForwardingRulesRules

`func NewDevicesSerialCellularGatewayPortForwardingRulesRules(lanIp string, publicPort string, localPort string, protocol string, access string, ) *DevicesSerialCellularGatewayPortForwardingRulesRules`

NewDevicesSerialCellularGatewayPortForwardingRulesRules instantiates a new DevicesSerialCellularGatewayPortForwardingRulesRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialCellularGatewayPortForwardingRulesRulesWithDefaults

`func NewDevicesSerialCellularGatewayPortForwardingRulesRulesWithDefaults() *DevicesSerialCellularGatewayPortForwardingRulesRules`

NewDevicesSerialCellularGatewayPortForwardingRulesRulesWithDefaults instantiates a new DevicesSerialCellularGatewayPortForwardingRulesRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanIp

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.


### GetPublicPort

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetPublicPort() string`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetPublicPortOk() (*string, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetPublicPort(v string)`

SetPublicPort sets PublicPort field to given value.


### GetLocalPort

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.


### GetAllowedIps

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### GetProtocol

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetAccess

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DevicesSerialCellularGatewayPortForwardingRulesRules) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


