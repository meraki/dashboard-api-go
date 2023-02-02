# NetworksNetworkIdSwitchMtuOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Switches** | Pointer to **[]string** | List of switch serials. Applicable only for switch network. | [optional] 
**SwitchProfiles** | Pointer to **[]string** | List of switch profile IDs. Applicable only for template network. | [optional] 
**MtuSize** | **int32** | MTU size for the switches or switch profiles. | 

## Methods

### NewNetworksNetworkIdSwitchMtuOverrides

`func NewNetworksNetworkIdSwitchMtuOverrides(mtuSize int32, ) *NetworksNetworkIdSwitchMtuOverrides`

NewNetworksNetworkIdSwitchMtuOverrides instantiates a new NetworksNetworkIdSwitchMtuOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchMtuOverridesWithDefaults

`func NewNetworksNetworkIdSwitchMtuOverridesWithDefaults() *NetworksNetworkIdSwitchMtuOverrides`

NewNetworksNetworkIdSwitchMtuOverridesWithDefaults instantiates a new NetworksNetworkIdSwitchMtuOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitches

`func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *NetworksNetworkIdSwitchMtuOverrides) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *NetworksNetworkIdSwitchMtuOverrides) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitchProfiles() []string`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitchProfilesOk() (*[]string, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *NetworksNetworkIdSwitchMtuOverrides) SetSwitchProfiles(v []string)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *NetworksNetworkIdSwitchMtuOverrides) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### GetMtuSize

`func (o *NetworksNetworkIdSwitchMtuOverrides) GetMtuSize() int32`

GetMtuSize returns the MtuSize field if non-nil, zero value otherwise.

### GetMtuSizeOk

`func (o *NetworksNetworkIdSwitchMtuOverrides) GetMtuSizeOk() (*int32, bool)`

GetMtuSizeOk returns a tuple with the MtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuSize

`func (o *NetworksNetworkIdSwitchMtuOverrides) SetMtuSize(v int32)`

SetMtuSize sets MtuSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


