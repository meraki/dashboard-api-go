# GetNetworkSwitchMtu200ResponseOverridesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Switches** | Pointer to **[]string** | List of switch serials. Applicable only for switch network. | [optional] 
**SwitchProfiles** | Pointer to **[]string** | List of switch profile IDs. Applicable only for template network. | [optional] 
**MtuSize** | **int32** | MTU size for the switches or switch profiles. | 

## Methods

### NewGetNetworkSwitchMtu200ResponseOverridesInner

`func NewGetNetworkSwitchMtu200ResponseOverridesInner(mtuSize int32, ) *GetNetworkSwitchMtu200ResponseOverridesInner`

NewGetNetworkSwitchMtu200ResponseOverridesInner instantiates a new GetNetworkSwitchMtu200ResponseOverridesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchMtu200ResponseOverridesInnerWithDefaults

`func NewGetNetworkSwitchMtu200ResponseOverridesInnerWithDefaults() *GetNetworkSwitchMtu200ResponseOverridesInner`

NewGetNetworkSwitchMtu200ResponseOverridesInnerWithDefaults instantiates a new GetNetworkSwitchMtu200ResponseOverridesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitches

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitchProfiles() []string`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitchProfilesOk() (*[]string, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) SetSwitchProfiles(v []string)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### GetMtuSize

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetMtuSize() int32`

GetMtuSize returns the MtuSize field if non-nil, zero value otherwise.

### GetMtuSizeOk

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetMtuSizeOk() (*int32, bool)`

GetMtuSizeOk returns a tuple with the MtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuSize

`func (o *GetNetworkSwitchMtu200ResponseOverridesInner) SetMtuSize(v int32)`

SetMtuSize sets MtuSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


