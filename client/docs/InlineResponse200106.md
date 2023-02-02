# InlineResponse200106

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestOrganizationId** | Pointer to **string** | The ID of the organization to move the licenses to | [optional] 
**LicenseIds** | Pointer to **[]string** | A list of IDs of licenses to move to the new organization | [optional] 

## Methods

### NewInlineResponse200106

`func NewInlineResponse200106() *InlineResponse200106`

NewInlineResponse200106 instantiates a new InlineResponse200106 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200106WithDefaults

`func NewInlineResponse200106WithDefaults() *InlineResponse200106`

NewInlineResponse200106WithDefaults instantiates a new InlineResponse200106 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestOrganizationId

`func (o *InlineResponse200106) GetDestOrganizationId() string`

GetDestOrganizationId returns the DestOrganizationId field if non-nil, zero value otherwise.

### GetDestOrganizationIdOk

`func (o *InlineResponse200106) GetDestOrganizationIdOk() (*string, bool)`

GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOrganizationId

`func (o *InlineResponse200106) SetDestOrganizationId(v string)`

SetDestOrganizationId sets DestOrganizationId field to given value.

### HasDestOrganizationId

`func (o *InlineResponse200106) HasDestOrganizationId() bool`

HasDestOrganizationId returns a boolean if a field has been set.

### GetLicenseIds

`func (o *InlineResponse200106) GetLicenseIds() []string`

GetLicenseIds returns the LicenseIds field if non-nil, zero value otherwise.

### GetLicenseIdsOk

`func (o *InlineResponse200106) GetLicenseIdsOk() (*[]string, bool)`

GetLicenseIdsOk returns a tuple with the LicenseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseIds

`func (o *InlineResponse200106) SetLicenseIds(v []string)`

SetLicenseIds sets LicenseIds field to given value.

### HasLicenseIds

`func (o *InlineResponse200106) HasLicenseIds() bool`

HasLicenseIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


