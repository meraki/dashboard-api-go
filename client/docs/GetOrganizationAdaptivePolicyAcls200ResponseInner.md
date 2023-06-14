# GetOrganizationAdaptivePolicyAcls200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclId** | Pointer to **string** | ID of the adaptive policy ACL | [optional] 
**Name** | Pointer to **string** | Name of the adaptive policy ACL | [optional] 
**Description** | Pointer to **string** | Description of the adaptive policy ACL | [optional] 
**IpVersion** | Pointer to **string** | IP version of adpative policy ACL | [optional] 
**Rules** | Pointer to [**[]GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner**](GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner.md) | An ordered array of the adaptive policy ACL rules | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the adaptive policy ACL was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the adaptive policy ACL was last updated | [optional] 

## Methods

### NewGetOrganizationAdaptivePolicyAcls200ResponseInner

`func NewGetOrganizationAdaptivePolicyAcls200ResponseInner() *GetOrganizationAdaptivePolicyAcls200ResponseInner`

NewGetOrganizationAdaptivePolicyAcls200ResponseInner instantiates a new GetOrganizationAdaptivePolicyAcls200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAdaptivePolicyAcls200ResponseInnerWithDefaults

`func NewGetOrganizationAdaptivePolicyAcls200ResponseInnerWithDefaults() *GetOrganizationAdaptivePolicyAcls200ResponseInner`

NewGetOrganizationAdaptivePolicyAcls200ResponseInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyAcls200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclId

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpVersion

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetRules

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetRules() []GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetRulesOk() (*[]GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetRules(v []GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


