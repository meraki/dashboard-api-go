# GetNetworkVlanProfiles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iname** | Pointer to **string** | IName of the VLAN profile | [optional] 
**Name** | Pointer to **string** | Name of the profile, string length must be from 1 to 255 characters | [optional] 
**IsDefault** | Pointer to **bool** | Boolean indicating the default VLAN Profile for any device that does not have a profile explicitly assigned | [optional] 
**VlanNames** | Pointer to [**[]GetNetworkVlanProfiles200ResponseInnerVlanNamesInner**](GetNetworkVlanProfiles200ResponseInnerVlanNamesInner.md) | An array of named VLANs | [optional] 
**VlanGroups** | Pointer to [**[]GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner**](GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner.md) | An array of named VLANs | [optional] 

## Methods

### NewGetNetworkVlanProfiles200ResponseInner

`func NewGetNetworkVlanProfiles200ResponseInner() *GetNetworkVlanProfiles200ResponseInner`

NewGetNetworkVlanProfiles200ResponseInner instantiates a new GetNetworkVlanProfiles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkVlanProfiles200ResponseInnerWithDefaults

`func NewGetNetworkVlanProfiles200ResponseInnerWithDefaults() *GetNetworkVlanProfiles200ResponseInner`

NewGetNetworkVlanProfiles200ResponseInnerWithDefaults instantiates a new GetNetworkVlanProfiles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIname

`func (o *GetNetworkVlanProfiles200ResponseInner) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *GetNetworkVlanProfiles200ResponseInner) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *GetNetworkVlanProfiles200ResponseInner) SetIname(v string)`

SetIname sets Iname field to given value.

### HasIname

`func (o *GetNetworkVlanProfiles200ResponseInner) HasIname() bool`

HasIname returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkVlanProfiles200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkVlanProfiles200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkVlanProfiles200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkVlanProfiles200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetNetworkVlanProfiles200ResponseInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetNetworkVlanProfiles200ResponseInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetNetworkVlanProfiles200ResponseInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetNetworkVlanProfiles200ResponseInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetVlanNames

`func (o *GetNetworkVlanProfiles200ResponseInner) GetVlanNames() []GetNetworkVlanProfiles200ResponseInnerVlanNamesInner`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *GetNetworkVlanProfiles200ResponseInner) GetVlanNamesOk() (*[]GetNetworkVlanProfiles200ResponseInnerVlanNamesInner, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *GetNetworkVlanProfiles200ResponseInner) SetVlanNames(v []GetNetworkVlanProfiles200ResponseInnerVlanNamesInner)`

SetVlanNames sets VlanNames field to given value.

### HasVlanNames

`func (o *GetNetworkVlanProfiles200ResponseInner) HasVlanNames() bool`

HasVlanNames returns a boolean if a field has been set.

### GetVlanGroups

`func (o *GetNetworkVlanProfiles200ResponseInner) GetVlanGroups() []GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *GetNetworkVlanProfiles200ResponseInner) GetVlanGroupsOk() (*[]GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *GetNetworkVlanProfiles200ResponseInner) SetVlanGroups(v []GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner)`

SetVlanGroups sets VlanGroups field to given value.

### HasVlanGroups

`func (o *GetNetworkVlanProfiles200ResponseInner) HasVlanGroups() bool`

HasVlanGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


