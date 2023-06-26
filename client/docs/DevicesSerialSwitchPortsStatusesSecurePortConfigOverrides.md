# DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the  (&#39;trunk&#39; or &#39;access&#39;). | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the . A null value will clear the value set for trunk ports. | [optional] 
**VoiceVlan** | Pointer to **int32** | The voice VLAN of the . Only applicable to access ports. | [optional] 
**AllowedVlans** | Pointer to **string** | The VLANs allowed on the . Only applicable to trunk ports. | [optional] 

## Methods

### NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides

`func NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides() *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides`

NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides instantiates a new DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverridesWithDefaults

`func NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverridesWithDefaults() *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides`

NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverridesWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVoiceVlan

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


