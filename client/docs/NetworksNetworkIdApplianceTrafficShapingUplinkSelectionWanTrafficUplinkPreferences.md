# NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters.md) | Array of traffic filters for this uplink preference rule | 
**PreferredUplink** | **string** | Preferred uplink for this uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences(trafficFilters []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters, preferredUplink string, ) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesWithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetTrafficFilters() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) SetTrafficFilters(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


