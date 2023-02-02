# InlineResponse200110ApiAuthenticationIpRestrictionsForKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges. | [optional] 
**Ranges** | Pointer to **[]string** | List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets. | [optional] 

## Methods

### NewInlineResponse200110ApiAuthenticationIpRestrictionsForKeys

`func NewInlineResponse200110ApiAuthenticationIpRestrictionsForKeys() *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys`

NewInlineResponse200110ApiAuthenticationIpRestrictionsForKeys instantiates a new InlineResponse200110ApiAuthenticationIpRestrictionsForKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110ApiAuthenticationIpRestrictionsForKeysWithDefaults

`func NewInlineResponse200110ApiAuthenticationIpRestrictionsForKeysWithDefaults() *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys`

NewInlineResponse200110ApiAuthenticationIpRestrictionsForKeysWithDefaults instantiates a new InlineResponse200110ApiAuthenticationIpRestrictionsForKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) GetRanges() []string`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) GetRangesOk() (*[]string, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) SetRanges(v []string)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *InlineResponse200110ApiAuthenticationIpRestrictionsForKeys) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


