# CreateOrganizationAdaptivePolicyAclRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the adaptive policy ACL | 
**Description** | Pointer to **string** | Description of the adaptive policy ACL | [optional] [default to ""]
**Rules** | [**[]CreateOrganizationAdaptivePolicyAclRequestRulesInner**](CreateOrganizationAdaptivePolicyAclRequestRulesInner.md) | An ordered array of the adaptive policy ACL rules. | 
**IpVersion** | **string** | IP version of adpative policy ACL. One of: &#39;any&#39;, &#39;ipv4&#39; or &#39;ipv6&#39; | 

## Methods

### NewCreateOrganizationAdaptivePolicyAclRequest

`func NewCreateOrganizationAdaptivePolicyAclRequest(name string, rules []CreateOrganizationAdaptivePolicyAclRequestRulesInner, ipVersion string, ) *CreateOrganizationAdaptivePolicyAclRequest`

NewCreateOrganizationAdaptivePolicyAclRequest instantiates a new CreateOrganizationAdaptivePolicyAclRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationAdaptivePolicyAclRequestWithDefaults

`func NewCreateOrganizationAdaptivePolicyAclRequestWithDefaults() *CreateOrganizationAdaptivePolicyAclRequest`

NewCreateOrganizationAdaptivePolicyAclRequestWithDefaults instantiates a new CreateOrganizationAdaptivePolicyAclRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationAdaptivePolicyAclRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationAdaptivePolicyAclRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationAdaptivePolicyAclRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetRules() []CreateOrganizationAdaptivePolicyAclRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetRulesOk() (*[]CreateOrganizationAdaptivePolicyAclRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateOrganizationAdaptivePolicyAclRequest) SetRules(v []CreateOrganizationAdaptivePolicyAclRequestRulesInner)`

SetRules sets Rules field to given value.


### GetIpVersion

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *CreateOrganizationAdaptivePolicyAclRequest) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *CreateOrganizationAdaptivePolicyAclRequest) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


