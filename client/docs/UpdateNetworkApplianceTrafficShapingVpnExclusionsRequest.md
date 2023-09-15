# UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**[]UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner**](UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner.md) | Custom VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 
**MajorApplications** | Pointer to [**[]UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner**](UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner.md) | Major Application based VPN exclusion rules. Pass an empty array to clear existing rules. | [optional] 

## Methods

### NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest

`func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest`

NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestWithDefaults

`func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest`

NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetCustom() []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetCustomOk() (*[]UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) SetCustom(v []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetMajorApplications

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetMajorApplications() []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetMajorApplicationsOk() (*[]UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) SetMajorApplications(v []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner)`

SetMajorApplications sets MajorApplications field to given value.

### HasMajorApplications

`func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) HasMajorApplications() bool`

HasMajorApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


