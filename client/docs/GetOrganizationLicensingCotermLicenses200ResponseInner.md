# GetOrganizationLicensingCotermLicenses200ResponseInner

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
**Editions** | Pointer to [**[]GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner**](GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner.md) | The editions of the license for each relevant product type | [optional] 
**Counts** | Pointer to [**[]GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner**](GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner.md) | The counts of the license by model type | [optional] 

## Methods

### NewGetOrganizationLicensingCotermLicenses200ResponseInner

`func NewGetOrganizationLicensingCotermLicenses200ResponseInner() *GetOrganizationLicensingCotermLicenses200ResponseInner`

NewGetOrganizationLicensingCotermLicenses200ResponseInner instantiates a new GetOrganizationLicensingCotermLicenses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationLicensingCotermLicenses200ResponseInnerWithDefaults

`func NewGetOrganizationLicensingCotermLicenses200ResponseInnerWithDefaults() *GetOrganizationLicensingCotermLicenses200ResponseInner`

NewGetOrganizationLicensingCotermLicenses200ResponseInnerWithDefaults instantiates a new GetOrganizationLicensingCotermLicenses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDuration

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMode

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetClaimedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetClaimedAt() time.Time`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetClaimedAtOk() (*time.Time, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetClaimedAt(v time.Time)`

SetClaimedAt sets ClaimedAt field to given value.

### HasClaimedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasClaimedAt() bool`

HasClaimedAt returns a boolean if a field has been set.

### GetInvalidated

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidated() bool`

GetInvalidated returns the Invalidated field if non-nil, zero value otherwise.

### GetInvalidatedOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidatedOk() (*bool, bool)`

GetInvalidatedOk returns a tuple with the Invalidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidated

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetInvalidated(v bool)`

SetInvalidated sets Invalidated field to given value.

### HasInvalidated

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasInvalidated() bool`

HasInvalidated returns a boolean if a field has been set.

### GetInvalidatedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidatedAt() time.Time`

GetInvalidatedAt returns the InvalidatedAt field if non-nil, zero value otherwise.

### GetInvalidatedAtOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetInvalidatedAtOk() (*time.Time, bool)`

GetInvalidatedAtOk returns a tuple with the InvalidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidatedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetInvalidatedAt(v time.Time)`

SetInvalidatedAt sets InvalidatedAt field to given value.

### HasInvalidatedAt

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasInvalidatedAt() bool`

HasInvalidatedAt returns a boolean if a field has been set.

### GetExpired

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetEditions

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetEditions() []GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner`

GetEditions returns the Editions field if non-nil, zero value otherwise.

### GetEditionsOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetEditionsOk() (*[]GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner, bool)`

GetEditionsOk returns a tuple with the Editions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditions

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetEditions(v []GetOrganizationLicensingCotermLicenses200ResponseInnerEditionsInner)`

SetEditions sets Editions field to given value.

### HasEditions

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasEditions() bool`

HasEditions returns a boolean if a field has been set.

### GetCounts

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetCounts() []GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) GetCountsOk() (*[]GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) SetCounts(v []GetOrganizationLicensingCotermLicenses200ResponseInnerCountsInner)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *GetOrganizationLicensingCotermLicenses200ResponseInner) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


