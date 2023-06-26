# UpdateNetworkAlertsSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**UpdateNetworkAlertsSettingsRequestDefaultDestinations**](UpdateNetworkAlertsSettingsRequestDefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]UpdateNetworkAlertsSettingsRequestAlertsInner**](UpdateNetworkAlertsSettingsRequestAlertsInner.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 

## Methods

### NewUpdateNetworkAlertsSettingsRequest

`func NewUpdateNetworkAlertsSettingsRequest() *UpdateNetworkAlertsSettingsRequest`

NewUpdateNetworkAlertsSettingsRequest instantiates a new UpdateNetworkAlertsSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkAlertsSettingsRequestWithDefaults

`func NewUpdateNetworkAlertsSettingsRequestWithDefaults() *UpdateNetworkAlertsSettingsRequest`

NewUpdateNetworkAlertsSettingsRequestWithDefaults instantiates a new UpdateNetworkAlertsSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *UpdateNetworkAlertsSettingsRequest) GetDefaultDestinations() UpdateNetworkAlertsSettingsRequestDefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *UpdateNetworkAlertsSettingsRequest) GetDefaultDestinationsOk() (*UpdateNetworkAlertsSettingsRequestDefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *UpdateNetworkAlertsSettingsRequest) SetDefaultDestinations(v UpdateNetworkAlertsSettingsRequestDefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *UpdateNetworkAlertsSettingsRequest) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *UpdateNetworkAlertsSettingsRequest) GetAlerts() []UpdateNetworkAlertsSettingsRequestAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *UpdateNetworkAlertsSettingsRequest) GetAlertsOk() (*[]UpdateNetworkAlertsSettingsRequestAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *UpdateNetworkAlertsSettingsRequest) SetAlerts(v []UpdateNetworkAlertsSettingsRequestAlertsInner)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *UpdateNetworkAlertsSettingsRequest) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


