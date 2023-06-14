# GetOrganizationClientsOverview200ResponseUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to [**GetOrganizationClientsOverview200ResponseUsageOverall**](GetOrganizationClientsOverview200ResponseUsageOverall.md) |  | [optional] 
**Average** | Pointer to **float32** | Average data usage (in kb) of each client in organization | [optional] 

## Methods

### NewGetOrganizationClientsOverview200ResponseUsage

`func NewGetOrganizationClientsOverview200ResponseUsage() *GetOrganizationClientsOverview200ResponseUsage`

NewGetOrganizationClientsOverview200ResponseUsage instantiates a new GetOrganizationClientsOverview200ResponseUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationClientsOverview200ResponseUsageWithDefaults

`func NewGetOrganizationClientsOverview200ResponseUsageWithDefaults() *GetOrganizationClientsOverview200ResponseUsage`

NewGetOrganizationClientsOverview200ResponseUsageWithDefaults instantiates a new GetOrganizationClientsOverview200ResponseUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *GetOrganizationClientsOverview200ResponseUsage) GetOverall() GetOrganizationClientsOverview200ResponseUsageOverall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *GetOrganizationClientsOverview200ResponseUsage) GetOverallOk() (*GetOrganizationClientsOverview200ResponseUsageOverall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *GetOrganizationClientsOverview200ResponseUsage) SetOverall(v GetOrganizationClientsOverview200ResponseUsageOverall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *GetOrganizationClientsOverview200ResponseUsage) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetAverage

`func (o *GetOrganizationClientsOverview200ResponseUsage) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *GetOrganizationClientsOverview200ResponseUsage) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *GetOrganizationClientsOverview200ResponseUsage) SetAverage(v float32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *GetOrganizationClientsOverview200ResponseUsage) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


