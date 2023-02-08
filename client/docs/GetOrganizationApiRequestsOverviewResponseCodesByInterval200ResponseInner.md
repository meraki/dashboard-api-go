# GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the access period | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the access period | [optional] 
**Counts** | Pointer to [**[]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner**](GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner.md) | list of response codes and a count of how many requests had that code in the given time period | [optional] 

## Methods

### NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner

`func NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner`

NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner instantiates a new GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerWithDefaults

`func NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerWithDefaults() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner`

NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerWithDefaults instantiates a new GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetCounts

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetCounts() []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetCountsOk() (*[]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) SetCounts(v []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


