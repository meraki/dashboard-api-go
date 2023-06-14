# GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the  (&#39;trunk&#39; or &#39;access&#39;). | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the . A null value will clear the value set for trunk ports. | [optional] 
**VoiceVlan** | Pointer to **int32** | The voice VLAN of the . Only applicable to access ports. | [optional] 
**AllowedVlans** | Pointer to **string** | The VLANs allowed on the . Only applicable to trunk ports. | [optional] 

## Methods

### NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides

`func NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides() *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides`

NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverridesWithDefaults

`func NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverridesWithDefaults() *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides`

NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverridesWithDefaults instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVoiceVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


