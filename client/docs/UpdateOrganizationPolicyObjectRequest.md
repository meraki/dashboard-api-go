# UpdateOrganizationPolicyObjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only) | [optional] 
**Cidr** | Pointer to **string** | CIDR Value of a policy object (e.g. 10.11.12.1/24\&quot;) | [optional] 
**Fqdn** | Pointer to **string** | Fully qualified domain name of policy object (e.g. \&quot;example.com\&quot;) | [optional] 
**Mask** | Pointer to **string** | Mask of a policy object (e.g. \&quot;255.255.0.0\&quot;) | [optional] 
**Ip** | Pointer to **string** | IP Address of a policy object (e.g. \&quot;1.2.3.4\&quot;) | [optional] 
**GroupIds** | Pointer to **[]int32** | The IDs of policy object groups the policy object belongs to | [optional] 

## Methods

### NewUpdateOrganizationPolicyObjectRequest

`func NewUpdateOrganizationPolicyObjectRequest() *UpdateOrganizationPolicyObjectRequest`

NewUpdateOrganizationPolicyObjectRequest instantiates a new UpdateOrganizationPolicyObjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationPolicyObjectRequestWithDefaults

`func NewUpdateOrganizationPolicyObjectRequestWithDefaults() *UpdateOrganizationPolicyObjectRequest`

NewUpdateOrganizationPolicyObjectRequestWithDefaults instantiates a new UpdateOrganizationPolicyObjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrganizationPolicyObjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationPolicyObjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationPolicyObjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationPolicyObjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCidr

`func (o *UpdateOrganizationPolicyObjectRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateOrganizationPolicyObjectRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateOrganizationPolicyObjectRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateOrganizationPolicyObjectRequest) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetFqdn

`func (o *UpdateOrganizationPolicyObjectRequest) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *UpdateOrganizationPolicyObjectRequest) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *UpdateOrganizationPolicyObjectRequest) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *UpdateOrganizationPolicyObjectRequest) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetMask

`func (o *UpdateOrganizationPolicyObjectRequest) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *UpdateOrganizationPolicyObjectRequest) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *UpdateOrganizationPolicyObjectRequest) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *UpdateOrganizationPolicyObjectRequest) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetIp

`func (o *UpdateOrganizationPolicyObjectRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UpdateOrganizationPolicyObjectRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UpdateOrganizationPolicyObjectRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *UpdateOrganizationPolicyObjectRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetGroupIds

`func (o *UpdateOrganizationPolicyObjectRequest) GetGroupIds() []int32`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UpdateOrganizationPolicyObjectRequest) GetGroupIdsOk() (*[]int32, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UpdateOrganizationPolicyObjectRequest) SetGroupIds(v []int32)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UpdateOrganizationPolicyObjectRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


