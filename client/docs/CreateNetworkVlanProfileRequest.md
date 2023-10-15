# CreateNetworkVlanProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the profile, string length must be from 1 to 255 characters | 
**VlanNames** | [**[]CreateNetworkVlanProfileRequestVlanNamesInner**](CreateNetworkVlanProfileRequestVlanNamesInner.md) | An array of named VLANs | 
**VlanGroups** | [**[]CreateNetworkVlanProfileRequestVlanGroupsInner**](CreateNetworkVlanProfileRequestVlanGroupsInner.md) | An array of VLAN groups | 
**Iname** | **string** | IName of the profile | 

## Methods

### NewCreateNetworkVlanProfileRequest

`func NewCreateNetworkVlanProfileRequest(name string, vlanNames []CreateNetworkVlanProfileRequestVlanNamesInner, vlanGroups []CreateNetworkVlanProfileRequestVlanGroupsInner, iname string, ) *CreateNetworkVlanProfileRequest`

NewCreateNetworkVlanProfileRequest instantiates a new CreateNetworkVlanProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkVlanProfileRequestWithDefaults

`func NewCreateNetworkVlanProfileRequestWithDefaults() *CreateNetworkVlanProfileRequest`

NewCreateNetworkVlanProfileRequestWithDefaults instantiates a new CreateNetworkVlanProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkVlanProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkVlanProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkVlanProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVlanNames

`func (o *CreateNetworkVlanProfileRequest) GetVlanNames() []CreateNetworkVlanProfileRequestVlanNamesInner`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *CreateNetworkVlanProfileRequest) GetVlanNamesOk() (*[]CreateNetworkVlanProfileRequestVlanNamesInner, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *CreateNetworkVlanProfileRequest) SetVlanNames(v []CreateNetworkVlanProfileRequestVlanNamesInner)`

SetVlanNames sets VlanNames field to given value.


### GetVlanGroups

`func (o *CreateNetworkVlanProfileRequest) GetVlanGroups() []CreateNetworkVlanProfileRequestVlanGroupsInner`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *CreateNetworkVlanProfileRequest) GetVlanGroupsOk() (*[]CreateNetworkVlanProfileRequestVlanGroupsInner, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *CreateNetworkVlanProfileRequest) SetVlanGroups(v []CreateNetworkVlanProfileRequestVlanGroupsInner)`

SetVlanGroups sets VlanGroups field to given value.


### GetIname

`func (o *CreateNetworkVlanProfileRequest) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *CreateNetworkVlanProfileRequest) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *CreateNetworkVlanProfileRequest) SetIname(v string)`

SetIname sets Iname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


