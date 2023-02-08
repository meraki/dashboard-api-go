# CreateOrganizationPolicyObjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only) | 
**Category** | **string** | Category of a policy object (one of: adaptivePolicy, network) | 
**Type** | **string** | Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask) | 
**Cidr** | Pointer to **string** | CIDR Value of a policy object (e.g. 10.11.12.1/24\&quot;) | [optional] 
**Fqdn** | Pointer to **string** | Fully qualified domain name of policy object (e.g. \&quot;example.com\&quot;) | [optional] 
**Mask** | Pointer to **string** | Mask of a policy object (e.g. \&quot;255.255.0.0\&quot;) | [optional] 
**Ip** | Pointer to **string** | IP Address of a policy object (e.g. \&quot;1.2.3.4\&quot;) | [optional] 
**GroupIds** | Pointer to **[]int32** | The IDs of policy object groups the policy object belongs to | [optional] 

## Methods

### NewCreateOrganizationPolicyObjectRequest

`func NewCreateOrganizationPolicyObjectRequest(name string, category string, type_ string, ) *CreateOrganizationPolicyObjectRequest`

NewCreateOrganizationPolicyObjectRequest instantiates a new CreateOrganizationPolicyObjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationPolicyObjectRequestWithDefaults

`func NewCreateOrganizationPolicyObjectRequestWithDefaults() *CreateOrganizationPolicyObjectRequest`

NewCreateOrganizationPolicyObjectRequestWithDefaults instantiates a new CreateOrganizationPolicyObjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationPolicyObjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationPolicyObjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationPolicyObjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *CreateOrganizationPolicyObjectRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateOrganizationPolicyObjectRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateOrganizationPolicyObjectRequest) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetType

`func (o *CreateOrganizationPolicyObjectRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrganizationPolicyObjectRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrganizationPolicyObjectRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCidr

`func (o *CreateOrganizationPolicyObjectRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateOrganizationPolicyObjectRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateOrganizationPolicyObjectRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CreateOrganizationPolicyObjectRequest) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetFqdn

`func (o *CreateOrganizationPolicyObjectRequest) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateOrganizationPolicyObjectRequest) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateOrganizationPolicyObjectRequest) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *CreateOrganizationPolicyObjectRequest) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetMask

`func (o *CreateOrganizationPolicyObjectRequest) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *CreateOrganizationPolicyObjectRequest) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *CreateOrganizationPolicyObjectRequest) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *CreateOrganizationPolicyObjectRequest) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetIp

`func (o *CreateOrganizationPolicyObjectRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateOrganizationPolicyObjectRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateOrganizationPolicyObjectRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateOrganizationPolicyObjectRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetGroupIds

`func (o *CreateOrganizationPolicyObjectRequest) GetGroupIds() []int32`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CreateOrganizationPolicyObjectRequest) GetGroupIdsOk() (*[]int32, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CreateOrganizationPolicyObjectRequest) SetGroupIds(v []int32)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *CreateOrganizationPolicyObjectRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


