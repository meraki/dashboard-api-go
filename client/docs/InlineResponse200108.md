# InlineResponse200108

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the license | [optional] 
**OrganizationId** | Pointer to **string** | The ID of the organization that the license is claimed in | [optional] 
**Duration** | Pointer to **int32** | The duration (term length) of the license, measured in days | [optional] 
**Mode** | Pointer to **string** | The operation mode of the license when it was claimed | [optional] 
**StartedAt** | Pointer to **time.Time** | When the license&#39;s term began (approximately the date when the license was created) | [optional] 
**ClaimedAt** | Pointer to **time.Time** | When the license was claimed into the organization | [optional] 
**Invalidated** | Pointer to **bool** | Flag to indicated that the license is invalidated | [optional] 
**InvalidatedAt** | Pointer to **time.Time** | When the license was invalidated. Will be null for active licenses | [optional] 
**Expired** | Pointer to **bool** | Flag to indicate if the license is expired | [optional] 
**Editions** | Pointer to [**[]OrganizationsOrganizationIdLicensingCotermLicensesEditions**](OrganizationsOrganizationIdLicensingCotermLicensesEditions.md) | The editions of the license for each relevant product type | [optional] 
**Counts** | Pointer to [**[]OrganizationsOrganizationIdLicensingCotermLicensesCounts**](OrganizationsOrganizationIdLicensingCotermLicensesCounts.md) | The counts of the license by model type | [optional] 

## Methods

### NewInlineResponse200108

`func NewInlineResponse200108() *InlineResponse200108`

NewInlineResponse200108 instantiates a new InlineResponse200108 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200108WithDefaults

`func NewInlineResponse200108WithDefaults() *InlineResponse200108`

NewInlineResponse200108WithDefaults instantiates a new InlineResponse200108 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *InlineResponse200108) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InlineResponse200108) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InlineResponse200108) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InlineResponse200108) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse200108) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse200108) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse200108) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse200108) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDuration

`func (o *InlineResponse200108) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse200108) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse200108) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse200108) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse200108) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200108) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200108) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200108) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetStartedAt

`func (o *InlineResponse200108) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InlineResponse200108) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InlineResponse200108) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *InlineResponse200108) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetClaimedAt

`func (o *InlineResponse200108) GetClaimedAt() time.Time`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *InlineResponse200108) GetClaimedAtOk() (*time.Time, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *InlineResponse200108) SetClaimedAt(v time.Time)`

SetClaimedAt sets ClaimedAt field to given value.

### HasClaimedAt

`func (o *InlineResponse200108) HasClaimedAt() bool`

HasClaimedAt returns a boolean if a field has been set.

### GetInvalidated

`func (o *InlineResponse200108) GetInvalidated() bool`

GetInvalidated returns the Invalidated field if non-nil, zero value otherwise.

### GetInvalidatedOk

`func (o *InlineResponse200108) GetInvalidatedOk() (*bool, bool)`

GetInvalidatedOk returns a tuple with the Invalidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidated

`func (o *InlineResponse200108) SetInvalidated(v bool)`

SetInvalidated sets Invalidated field to given value.

### HasInvalidated

`func (o *InlineResponse200108) HasInvalidated() bool`

HasInvalidated returns a boolean if a field has been set.

### GetInvalidatedAt

`func (o *InlineResponse200108) GetInvalidatedAt() time.Time`

GetInvalidatedAt returns the InvalidatedAt field if non-nil, zero value otherwise.

### GetInvalidatedAtOk

`func (o *InlineResponse200108) GetInvalidatedAtOk() (*time.Time, bool)`

GetInvalidatedAtOk returns a tuple with the InvalidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidatedAt

`func (o *InlineResponse200108) SetInvalidatedAt(v time.Time)`

SetInvalidatedAt sets InvalidatedAt field to given value.

### HasInvalidatedAt

`func (o *InlineResponse200108) HasInvalidatedAt() bool`

HasInvalidatedAt returns a boolean if a field has been set.

### GetExpired

`func (o *InlineResponse200108) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *InlineResponse200108) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *InlineResponse200108) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *InlineResponse200108) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetEditions

`func (o *InlineResponse200108) GetEditions() []OrganizationsOrganizationIdLicensingCotermLicensesEditions`

GetEditions returns the Editions field if non-nil, zero value otherwise.

### GetEditionsOk

`func (o *InlineResponse200108) GetEditionsOk() (*[]OrganizationsOrganizationIdLicensingCotermLicensesEditions, bool)`

GetEditionsOk returns a tuple with the Editions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditions

`func (o *InlineResponse200108) SetEditions(v []OrganizationsOrganizationIdLicensingCotermLicensesEditions)`

SetEditions sets Editions field to given value.

### HasEditions

`func (o *InlineResponse200108) HasEditions() bool`

HasEditions returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200108) GetCounts() []OrganizationsOrganizationIdLicensingCotermLicensesCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200108) GetCountsOk() (*[]OrganizationsOrganizationIdLicensingCotermLicensesCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200108) SetCounts(v []OrganizationsOrganizationIdLicensingCotermLicensesCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200108) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


