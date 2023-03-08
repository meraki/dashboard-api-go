# NetworksNetworkIdAlertsSettingsAlerts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of alert | 
**Enabled** | Pointer to **bool** | A boolean depicting if the alert is turned on or off | [optional] 
**AlertDestinations** | Pointer to [**NetworksNetworkIdAlertsSettingsAlertDestinations**](NetworksNetworkIdAlertsSettingsAlertDestinations.md) |  | [optional] 
**Filters** | Pointer to **map[string]interface{}** | A hash of specific configuration data for the alert. Only filters specific to the alert will be updated. | [optional] 

## Methods

### NewNetworksNetworkIdAlertsSettingsAlerts

`func NewNetworksNetworkIdAlertsSettingsAlerts(type_ string, ) *NetworksNetworkIdAlertsSettingsAlerts`

NewNetworksNetworkIdAlertsSettingsAlerts instantiates a new NetworksNetworkIdAlertsSettingsAlerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdAlertsSettingsAlertsWithDefaults

`func NewNetworksNetworkIdAlertsSettingsAlertsWithDefaults() *NetworksNetworkIdAlertsSettingsAlerts`

NewNetworksNetworkIdAlertsSettingsAlertsWithDefaults instantiates a new NetworksNetworkIdAlertsSettingsAlerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworksNetworkIdAlertsSettingsAlerts) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdAlertsSettingsAlerts) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdAlertsSettingsAlerts) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAlertDestinations

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetAlertDestinations() NetworksNetworkIdAlertsSettingsAlertDestinations`

GetAlertDestinations returns the AlertDestinations field if non-nil, zero value otherwise.

### GetAlertDestinationsOk

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetAlertDestinationsOk() (*NetworksNetworkIdAlertsSettingsAlertDestinations, bool)`

GetAlertDestinationsOk returns a tuple with the AlertDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDestinations

`func (o *NetworksNetworkIdAlertsSettingsAlerts) SetAlertDestinations(v NetworksNetworkIdAlertsSettingsAlertDestinations)`

SetAlertDestinations sets AlertDestinations field to given value.

### HasAlertDestinations

`func (o *NetworksNetworkIdAlertsSettingsAlerts) HasAlertDestinations() bool`

HasAlertDestinations returns a boolean if a field has been set.

### GetFilters

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *NetworksNetworkIdAlertsSettingsAlerts) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *NetworksNetworkIdAlertsSettingsAlerts) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *NetworksNetworkIdAlertsSettingsAlerts) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


