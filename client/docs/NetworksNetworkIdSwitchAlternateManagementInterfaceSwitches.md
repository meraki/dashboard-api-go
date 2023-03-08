# NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | Switch serial number | 
**AlternateManagementIp** | **string** | Switch alternative management IP. To remove a prior IP setting, provide an empty string | 
**SubnetMask** | Pointer to **string** | Switch subnet mask must be in IP format. Only and must be specified for Polaris switches | [optional] 
**Gateway** | Pointer to **string** | Switch gateway must be in IP format. Only and must be specified for Polaris switches | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches

`func NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches(serial string, alternateManagementIp string, ) *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches`

NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches instantiates a new NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitchesWithDefaults

`func NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitchesWithDefaults() *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches`

NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitchesWithDefaults instantiates a new NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetAlternateManagementIp

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetAlternateManagementIp() string`

GetAlternateManagementIp returns the AlternateManagementIp field if non-nil, zero value otherwise.

### GetAlternateManagementIpOk

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetAlternateManagementIpOk() (*string, bool)`

GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateManagementIp

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetAlternateManagementIp(v string)`

SetAlternateManagementIp sets AlternateManagementIp field to given value.


### GetSubnetMask

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetGateway

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


