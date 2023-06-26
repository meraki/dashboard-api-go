# GetOrganizationInsightApplications200ResponseInnerThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Threshold type (static or smart) | [optional] 
**ByNetwork** | Pointer to [**[]GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner**](GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner.md) | Threshold for each network | [optional] 

## Methods

### NewGetOrganizationInsightApplications200ResponseInnerThresholds

`func NewGetOrganizationInsightApplications200ResponseInnerThresholds() *GetOrganizationInsightApplications200ResponseInnerThresholds`

NewGetOrganizationInsightApplications200ResponseInnerThresholds instantiates a new GetOrganizationInsightApplications200ResponseInnerThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationInsightApplications200ResponseInnerThresholdsWithDefaults

`func NewGetOrganizationInsightApplications200ResponseInnerThresholdsWithDefaults() *GetOrganizationInsightApplications200ResponseInnerThresholds`

NewGetOrganizationInsightApplications200ResponseInnerThresholdsWithDefaults instantiates a new GetOrganizationInsightApplications200ResponseInnerThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) HasType() bool`

HasType returns a boolean if a field has been set.

### GetByNetwork

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetByNetwork() []GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner`

GetByNetwork returns the ByNetwork field if non-nil, zero value otherwise.

### GetByNetworkOk

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetByNetworkOk() (*[]GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner, bool)`

GetByNetworkOk returns a tuple with the ByNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByNetwork

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) SetByNetwork(v []GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner)`

SetByNetwork sets ByNetwork field to given value.

### HasByNetwork

`func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) HasByNetwork() bool`

HasByNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


