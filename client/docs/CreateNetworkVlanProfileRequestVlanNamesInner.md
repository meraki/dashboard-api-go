# CreateNetworkVlanProfileRequestVlanNamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the VLAN, string length must be from 1 to 32 characters | 
**VlanId** | **string** | VLAN ID | 
**AdaptivePolicyGroup** | Pointer to [**CreateNetworkVlanProfileRequestVlanNamesInnerAdaptivePolicyGroup**](CreateNetworkVlanProfileRequestVlanNamesInnerAdaptivePolicyGroup.md) |  | [optional] 

## Methods

### NewCreateNetworkVlanProfileRequestVlanNamesInner

`func NewCreateNetworkVlanProfileRequestVlanNamesInner(name string, vlanId string, ) *CreateNetworkVlanProfileRequestVlanNamesInner`

NewCreateNetworkVlanProfileRequestVlanNamesInner instantiates a new CreateNetworkVlanProfileRequestVlanNamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkVlanProfileRequestVlanNamesInnerWithDefaults

`func NewCreateNetworkVlanProfileRequestVlanNamesInnerWithDefaults() *CreateNetworkVlanProfileRequestVlanNamesInner`

NewCreateNetworkVlanProfileRequestVlanNamesInnerWithDefaults instantiates a new CreateNetworkVlanProfileRequestVlanNamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) SetName(v string)`

SetName sets Name field to given value.


### GetVlanId

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.


### GetAdaptivePolicyGroup

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) GetAdaptivePolicyGroup() CreateNetworkVlanProfileRequestVlanNamesInnerAdaptivePolicyGroup`

GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupOk

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) GetAdaptivePolicyGroupOk() (*CreateNetworkVlanProfileRequestVlanNamesInnerAdaptivePolicyGroup, bool)`

GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroup

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) SetAdaptivePolicyGroup(v CreateNetworkVlanProfileRequestVlanNamesInnerAdaptivePolicyGroup)`

SetAdaptivePolicyGroup sets AdaptivePolicyGroup field to given value.

### HasAdaptivePolicyGroup

`func (o *CreateNetworkVlanProfileRequestVlanNamesInner) HasAdaptivePolicyGroup() bool`

HasAdaptivePolicyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


