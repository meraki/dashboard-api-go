# GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges. | [optional] 
**Ranges** | Pointer to **[]string** | List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets. | [optional] 

## Methods

### NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys

`func NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys() *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys`

NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys instantiates a new GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeysWithDefaults

`func NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeysWithDefaults() *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys`

NewGetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeysWithDefaults instantiates a new GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetRanges() []string`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) GetRangesOk() (*[]string, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) SetRanges(v []string)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *GetOrganizationLoginSecurity200ResponseApiAuthenticationIpRestrictionsForKeys) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


