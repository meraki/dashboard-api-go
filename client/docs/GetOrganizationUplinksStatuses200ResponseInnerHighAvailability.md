# GetOrganizationUplinksStatuses200ResponseInnerHighAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether High Availability is enabled for the device. For devices that do not support HA, this will be &#39;false&#39; | [optional] 
**Role** | Pointer to **string** | The HA role of the device on the network. For devices that do not support HA, this will be &#39;primary&#39; | [optional] 

## Methods

### NewGetOrganizationUplinksStatuses200ResponseInnerHighAvailability

`func NewGetOrganizationUplinksStatuses200ResponseInnerHighAvailability() *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability`

NewGetOrganizationUplinksStatuses200ResponseInnerHighAvailability instantiates a new GetOrganizationUplinksStatuses200ResponseInnerHighAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationUplinksStatuses200ResponseInnerHighAvailabilityWithDefaults

`func NewGetOrganizationUplinksStatuses200ResponseInnerHighAvailabilityWithDefaults() *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability`

NewGetOrganizationUplinksStatuses200ResponseInnerHighAvailabilityWithDefaults instantiates a new GetOrganizationUplinksStatuses200ResponseInnerHighAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRole

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetOrganizationUplinksStatuses200ResponseInnerHighAvailability) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


