# UpdateOrganizationAdaptivePolicyAclRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the adaptive policy ACL | [optional] 
**Description** | Pointer to **string** | Description of the adaptive policy ACL | [optional] 
**Rules** | Pointer to [**[]CreateOrganizationAdaptivePolicyAclRequestRulesInner**](CreateOrganizationAdaptivePolicyAclRequestRulesInner.md) | An ordered array of the adaptive policy ACL rules. An empty array will clear the rules. | [optional] 
**IpVersion** | Pointer to **string** | IP version of adpative policy ACL. One of: &#39;any&#39;, &#39;ipv4&#39; or &#39;ipv6&#39; | [optional] 

## Methods

### NewUpdateOrganizationAdaptivePolicyAclRequest

`func NewUpdateOrganizationAdaptivePolicyAclRequest() *UpdateOrganizationAdaptivePolicyAclRequest`

NewUpdateOrganizationAdaptivePolicyAclRequest instantiates a new UpdateOrganizationAdaptivePolicyAclRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationAdaptivePolicyAclRequestWithDefaults

`func NewUpdateOrganizationAdaptivePolicyAclRequestWithDefaults() *UpdateOrganizationAdaptivePolicyAclRequest`

NewUpdateOrganizationAdaptivePolicyAclRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicyAclRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetRules() []CreateOrganizationAdaptivePolicyAclRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetRulesOk() (*[]CreateOrganizationAdaptivePolicyAclRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetRules(v []CreateOrganizationAdaptivePolicyAclRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetIpVersion

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *UpdateOrganizationAdaptivePolicyAclRequest) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


