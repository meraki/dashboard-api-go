# UpdateOrganizationAdaptivePolicyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the group | [optional] 
**Sgt** | Pointer to **int32** | SGT value of the group | [optional] 
**Description** | Pointer to **string** | Description of the group | [optional] 
**PolicyObjects** | Pointer to [**[]CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner**](CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner.md) | The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group&#39;s SGT value if no other tagging scheme is being used (each requires one unique attribute) | [optional] 

## Methods

### NewUpdateOrganizationAdaptivePolicyGroupRequest

`func NewUpdateOrganizationAdaptivePolicyGroupRequest() *UpdateOrganizationAdaptivePolicyGroupRequest`

NewUpdateOrganizationAdaptivePolicyGroupRequest instantiates a new UpdateOrganizationAdaptivePolicyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationAdaptivePolicyGroupRequestWithDefaults

`func NewUpdateOrganizationAdaptivePolicyGroupRequestWithDefaults() *UpdateOrganizationAdaptivePolicyGroupRequest`

NewUpdateOrganizationAdaptivePolicyGroupRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSgt

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetSgt(v int32)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjects() []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjectsOk() (*[]CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetPolicyObjects(v []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


