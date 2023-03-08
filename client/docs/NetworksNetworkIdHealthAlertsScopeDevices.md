# NetworksNetworkIdHealthAlertsScopeDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL to the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Mac** | Pointer to **string** | The mac address of the device | [optional] 
**Lldp** | Pointer to [**NetworksNetworkIdHealthAlertsScopeLldp**](NetworksNetworkIdHealthAlertsScopeLldp.md) |  | [optional] 
**Clients** | Pointer to [**[]NetworksNetworkIdHealthAlertsScopeClients**](NetworksNetworkIdHealthAlertsScopeClients.md) | Clients related to the device | [optional] 

## Methods

### NewNetworksNetworkIdHealthAlertsScopeDevices

`func NewNetworksNetworkIdHealthAlertsScopeDevices() *NetworksNetworkIdHealthAlertsScopeDevices`

NewNetworksNetworkIdHealthAlertsScopeDevices instantiates a new NetworksNetworkIdHealthAlertsScopeDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdHealthAlertsScopeDevicesWithDefaults

`func NewNetworksNetworkIdHealthAlertsScopeDevicesWithDefaults() *NetworksNetworkIdHealthAlertsScopeDevices`

NewNetworksNetworkIdHealthAlertsScopeDevicesWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScopeDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductType

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLldp

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetLldp() NetworksNetworkIdHealthAlertsScopeLldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetLldpOk() (*NetworksNetworkIdHealthAlertsScopeLldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetLldp(v NetworksNetworkIdHealthAlertsScopeLldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetClients

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetClients() []NetworksNetworkIdHealthAlertsScopeClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) GetClientsOk() (*[]NetworksNetworkIdHealthAlertsScopeClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) SetClients(v []NetworksNetworkIdHealthAlertsScopeClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *NetworksNetworkIdHealthAlertsScopeDevices) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


