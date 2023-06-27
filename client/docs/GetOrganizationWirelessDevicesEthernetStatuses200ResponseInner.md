# GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number of the AP | [optional] 
**Name** | Pointer to **string** | The name of the AP | [optional] 
**Network** | Pointer to [**GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork**](GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork.md) |  | [optional] 
**Power** | Pointer to [**GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower**](GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower.md) |  | [optional] 
**Ports** | Pointer to [**[]GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner**](GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner.md) | List of port details | [optional] 
**Aggregation** | Pointer to [**GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation**](GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation.md) |  | [optional] 

## Methods

### NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInner

`func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInner() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner`

NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInner instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerWithDefaults

`func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner`

NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPower

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetPower() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetPowerOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) SetPower(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower)`

SetPower sets Power field to given value.

### HasPower

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPorts

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetPorts() []GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetPortsOk() (*[]GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) SetPorts(v []GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetAggregation

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetAggregation() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) GetAggregationOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) SetAggregation(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


