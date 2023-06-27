# GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL to the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Mac** | Pointer to **string** | The mac address of the device | [optional] 
**Lldp** | Pointer to [**GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerLldp**](GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerLldp.md) |  | [optional] 
**Clients** | Pointer to [**[]GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner**](GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner.md) | Clients related to the device | [optional] 

## Methods

### NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInner

`func NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInner() *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner`

NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInner instantiates a new GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerWithDefaults

`func NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerWithDefaults() *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner`

NewGetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductType

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLldp

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetLldp() GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerLldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetLldpOk() (*GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerLldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) SetLldp(v GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerLldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetClients

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetClients() []GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) GetClientsOk() (*[]GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) SetClients(v []GetNetworkHealthAlerts200ResponseInnerScopeDevicesInnerClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *GetNetworkHealthAlerts200ResponseInnerScopeDevicesInner) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


