# CreateNetworkVlanProfileRequestVlanGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the VLAN, string length must be from 1 to 32 characters | 
**VlanIds** | **string** | Comma-separated VLAN IDs or ID ranges | 

## Methods

### NewCreateNetworkVlanProfileRequestVlanGroupsInner

`func NewCreateNetworkVlanProfileRequestVlanGroupsInner(name string, vlanIds string, ) *CreateNetworkVlanProfileRequestVlanGroupsInner`

NewCreateNetworkVlanProfileRequestVlanGroupsInner instantiates a new CreateNetworkVlanProfileRequestVlanGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkVlanProfileRequestVlanGroupsInnerWithDefaults

`func NewCreateNetworkVlanProfileRequestVlanGroupsInnerWithDefaults() *CreateNetworkVlanProfileRequestVlanGroupsInner`

NewCreateNetworkVlanProfileRequestVlanGroupsInnerWithDefaults instantiates a new CreateNetworkVlanProfileRequestVlanGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) SetName(v string)`

SetName sets Name field to given value.


### GetVlanIds

`func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetVlanIds() string`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) GetVlanIdsOk() (*string, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *CreateNetworkVlanProfileRequestVlanGroupsInner) SetVlanIds(v string)`

SetVlanIds sets VlanIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


