# InlineResponse20019VpnTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20019TrafficFilters1**](InlineResponse20019TrafficFilters1.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39;, &#39;wan2&#39;, &#39;bestForVoIP&#39;, &#39;loadBalancing&#39; or &#39;defaultUplink&#39; | 
**FailOverCriterion** | Pointer to **string** | Fail over criterion for uplink preference rule. Must be one of: &#39;poorPerformance&#39; or &#39;uplinkDown&#39; | [optional] 
**PerformanceClass** | Pointer to [**InlineResponse20019PerformanceClass**](InlineResponse20019PerformanceClass.md) |  | [optional] 

## Methods

### NewInlineResponse20019VpnTrafficUplinkPreferences

`func NewInlineResponse20019VpnTrafficUplinkPreferences(trafficFilters []InlineResponse20019TrafficFilters1, preferredUplink string, ) *InlineResponse20019VpnTrafficUplinkPreferences`

NewInlineResponse20019VpnTrafficUplinkPreferences instantiates a new InlineResponse20019VpnTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019VpnTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20019VpnTrafficUplinkPreferencesWithDefaults() *InlineResponse20019VpnTrafficUplinkPreferences`

NewInlineResponse20019VpnTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20019VpnTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20019TrafficFilters1`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20019TrafficFilters1, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20019TrafficFilters1)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetFailOverCriterion

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetFailOverCriterion() string`

GetFailOverCriterion returns the FailOverCriterion field if non-nil, zero value otherwise.

### GetFailOverCriterionOk

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool)`

GetFailOverCriterionOk returns a tuple with the FailOverCriterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOverCriterion

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) SetFailOverCriterion(v string)`

SetFailOverCriterion sets FailOverCriterion field to given value.

### HasFailOverCriterion

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) HasFailOverCriterion() bool`

HasFailOverCriterion returns a boolean if a field has been set.

### GetPerformanceClass

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetPerformanceClass() InlineResponse20019PerformanceClass`

GetPerformanceClass returns the PerformanceClass field if non-nil, zero value otherwise.

### GetPerformanceClassOk

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) GetPerformanceClassOk() (*InlineResponse20019PerformanceClass, bool)`

GetPerformanceClassOk returns a tuple with the PerformanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceClass

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) SetPerformanceClass(v InlineResponse20019PerformanceClass)`

SetPerformanceClass sets PerformanceClass field to given value.

### HasPerformanceClass

`func (o *InlineResponse20019VpnTrafficUplinkPreferences) HasPerformanceClass() bool`

HasPerformanceClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


