# GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp, in iso8601 format, at which the event happened | [optional] 
**Device** | Pointer to [**GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice**](GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice.md) |  | [optional] 
**Details** | Pointer to [**GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails**](GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails.md) |  | [optional] 
**Network** | Pointer to [**GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork**](GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork.md) |  | [optional] 

## Methods

### NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner

`func NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner`

NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner instantiates a new GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerWithDefaults

`func NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerWithDefaults() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner`

NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDevice

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDevice() GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDeviceOk() (*GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetDevice(v GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDetails

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDetails() GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDetailsOk() (*GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetDetails(v GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetNetwork() GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetNetwork(v GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


