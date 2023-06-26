# NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | Serial number of access point to be configured with alternate management IP | 
**AlternateManagementIp** | **string** | Wireless alternate management interface device IP. Provide an empty string to remove alternate management IP assignment | 
**SubnetMask** | Pointer to **string** | Subnet mask must be in IP format | [optional] 
**Gateway** | Pointer to **string** | Gateway must be in IP format | [optional] 
**Dns1** | Pointer to **string** | Primary DNS must be in IP format | [optional] 
**Dns2** | Pointer to **string** | Optional secondary DNS must be in IP format | [optional] 

## Methods

### NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints

`func NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints(serial string, alternateManagementIp string, ) *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints`

NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints instantiates a new NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPointsWithDefaults

`func NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPointsWithDefaults() *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints`

NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPointsWithDefaults instantiates a new NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetAlternateManagementIp

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetAlternateManagementIp() string`

GetAlternateManagementIp returns the AlternateManagementIp field if non-nil, zero value otherwise.

### GetAlternateManagementIpOk

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetAlternateManagementIpOk() (*string, bool)`

GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateManagementIp

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetAlternateManagementIp(v string)`

SetAlternateManagementIp sets AlternateManagementIp field to given value.


### GetSubnetMask

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetGateway

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetDns1

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


