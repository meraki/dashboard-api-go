# GetOrganizationSwitchPortsBySwitch200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the switch | [optional] 
**Serial** | Pointer to **string** | Serial of the switch | [optional] 
**Mac** | Pointer to **string** | MAC address of the switch | [optional] 
**Model** | Pointer to **string** | Model of the switch | [optional] 
**Network** | Pointer to [**GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork**](GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork.md) |  | [optional] 
**Ports** | Pointer to [**[]GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner**](GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner.md) | Ports belonging to the switch | [optional] 

## Methods

### NewGetOrganizationSwitchPortsBySwitch200ResponseInner

`func NewGetOrganizationSwitchPortsBySwitch200ResponseInner() *GetOrganizationSwitchPortsBySwitch200ResponseInner`

NewGetOrganizationSwitchPortsBySwitch200ResponseInner instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSwitchPortsBySwitch200ResponseInnerWithDefaults

`func NewGetOrganizationSwitchPortsBySwitch200ResponseInnerWithDefaults() *GetOrganizationSwitchPortsBySwitch200ResponseInner`

NewGetOrganizationSwitchPortsBySwitch200ResponseInnerWithDefaults instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetNetwork() GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetNetworkOk() (*GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetNetwork(v GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPorts

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetPorts() []GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetPortsOk() (*[]GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetPorts(v []GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


