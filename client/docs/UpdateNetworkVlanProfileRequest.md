# UpdateNetworkVlanProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the profile, string length must be from 1 to 255 characters | 
**VlanNames** | [**[]CreateNetworkVlanProfileRequestVlanNamesInner**](CreateNetworkVlanProfileRequestVlanNamesInner.md) | An array of named VLANs | 
**VlanGroups** | [**[]CreateNetworkVlanProfileRequestVlanGroupsInner**](CreateNetworkVlanProfileRequestVlanGroupsInner.md) | An array of VLAN groups | 

## Methods

### NewUpdateNetworkVlanProfileRequest

`func NewUpdateNetworkVlanProfileRequest(name string, vlanNames []CreateNetworkVlanProfileRequestVlanNamesInner, vlanGroups []CreateNetworkVlanProfileRequestVlanGroupsInner, ) *UpdateNetworkVlanProfileRequest`

NewUpdateNetworkVlanProfileRequest instantiates a new UpdateNetworkVlanProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkVlanProfileRequestWithDefaults

`func NewUpdateNetworkVlanProfileRequestWithDefaults() *UpdateNetworkVlanProfileRequest`

NewUpdateNetworkVlanProfileRequestWithDefaults instantiates a new UpdateNetworkVlanProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkVlanProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkVlanProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkVlanProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVlanNames

`func (o *UpdateNetworkVlanProfileRequest) GetVlanNames() []CreateNetworkVlanProfileRequestVlanNamesInner`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *UpdateNetworkVlanProfileRequest) GetVlanNamesOk() (*[]CreateNetworkVlanProfileRequestVlanNamesInner, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *UpdateNetworkVlanProfileRequest) SetVlanNames(v []CreateNetworkVlanProfileRequestVlanNamesInner)`

SetVlanNames sets VlanNames field to given value.


### GetVlanGroups

`func (o *UpdateNetworkVlanProfileRequest) GetVlanGroups() []CreateNetworkVlanProfileRequestVlanGroupsInner`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *UpdateNetworkVlanProfileRequest) GetVlanGroupsOk() (*[]CreateNetworkVlanProfileRequestVlanGroupsInner, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *UpdateNetworkVlanProfileRequest) SetVlanGroups(v []CreateNetworkVlanProfileRequestVlanGroupsInner)`

SetVlanGroups sets VlanGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


