# InlineResponse20022WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20022TrafficFilters**](InlineResponse20022TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | 

## Methods

### NewInlineResponse20022WanTrafficUplinkPreferences

`func NewInlineResponse20022WanTrafficUplinkPreferences(trafficFilters []InlineResponse20022TrafficFilters, preferredUplink string, ) *InlineResponse20022WanTrafficUplinkPreferences`

NewInlineResponse20022WanTrafficUplinkPreferences instantiates a new InlineResponse20022WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20022WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20022WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20022WanTrafficUplinkPreferences`

NewInlineResponse20022WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20022WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20022WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20022TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20022WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20022TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20022WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20022TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20022WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20022WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20022WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


