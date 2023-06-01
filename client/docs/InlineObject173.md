# InlineObject173

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the group | 
**Sgt** | **int32** | SGT value of the group | 
**Description** | Pointer to **string** | Description of the group (default: \&quot;\&quot;) | [optional] 
**PolicyObjects** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects**](OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects.md) | The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group&#39;s SGT value if no other tagging scheme is being used (each requires one unique attribute) (default: []) | [optional] 

## Methods

### NewInlineObject173

`func NewInlineObject173(name string, sgt int32, ) *InlineObject173`

NewInlineObject173 instantiates a new InlineObject173 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject173WithDefaults

`func NewInlineObject173WithDefaults() *InlineObject173`

NewInlineObject173WithDefaults instantiates a new InlineObject173 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject173) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject173) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject173) SetName(v string)`

SetName sets Name field to given value.


### GetSgt

`func (o *InlineObject173) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineObject173) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineObject173) SetSgt(v int32)`

SetSgt sets Sgt field to given value.


### GetDescription

`func (o *InlineObject173) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject173) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject173) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject173) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineObject173) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineObject173) GetPolicyObjectsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineObject173) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineObject173) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


