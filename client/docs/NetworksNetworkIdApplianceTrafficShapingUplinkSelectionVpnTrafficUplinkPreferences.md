# NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1.md) | Array of traffic filters for this uplink preference rule | 
**PreferredUplink** | **string** | Preferred uplink for this uplink preference rule. Must be one of: &#39;wan1&#39;, &#39;wan2&#39;, &#39;bestForVoIP&#39;, &#39;loadBalancing&#39; or &#39;defaultUplink&#39; | 
**FailOverCriterion** | Pointer to **string** | Fail over criterion for this uplink preference rule. Must be one of: &#39;poorPerformance&#39; or &#39;uplinkDown&#39; | [optional] 
**PerformanceClass** | Pointer to [**NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences(trafficFilters []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1, preferredUplink string, ) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesWithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetTrafficFilters() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetTrafficFilters(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetFailOverCriterion

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetFailOverCriterion() string`

GetFailOverCriterion returns the FailOverCriterion field if non-nil, zero value otherwise.

### GetFailOverCriterionOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool)`

GetFailOverCriterionOk returns a tuple with the FailOverCriterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOverCriterion

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetFailOverCriterion(v string)`

SetFailOverCriterion sets FailOverCriterion field to given value.

### HasFailOverCriterion

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) HasFailOverCriterion() bool`

HasFailOverCriterion returns a boolean if a field has been set.

### GetPerformanceClass

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPerformanceClass() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass`

GetPerformanceClass returns the PerformanceClass field if non-nil, zero value otherwise.

### GetPerformanceClassOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPerformanceClassOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass, bool)`

GetPerformanceClassOk returns a tuple with the PerformanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceClass

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetPerformanceClass(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass)`

SetPerformanceClass sets PerformanceClass field to given value.

### HasPerformanceClass

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) HasPerformanceClass() bool`

HasPerformanceClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


