# InlineObject95

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the sensor alert profile. | 
**Schedule** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesSchedule1**](NetworksNetworkIdSensorAlertsProfilesSchedule1.md) |  | [optional] 
**Conditions** | [**[]NetworksNetworkIdSensorAlertsProfilesConditions**](NetworksNetworkIdSensorAlertsProfilesConditions.md) | List of conditions that will cause the profile to send an alert. | 
**Recipients** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesRecipients**](NetworksNetworkIdSensorAlertsProfilesRecipients.md) |  | [optional] 
**Serials** | Pointer to **[]string** | List of device serials assigned to this sensor alert profile. | [optional] 

## Methods

### NewInlineObject95

`func NewInlineObject95(name string, conditions []NetworksNetworkIdSensorAlertsProfilesConditions, ) *InlineObject95`

NewInlineObject95 instantiates a new InlineObject95 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject95WithDefaults

`func NewInlineObject95WithDefaults() *InlineObject95`

NewInlineObject95WithDefaults instantiates a new InlineObject95 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject95) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject95) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject95) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *InlineObject95) GetSchedule() NetworksNetworkIdSensorAlertsProfilesSchedule1`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineObject95) GetScheduleOk() (*NetworksNetworkIdSensorAlertsProfilesSchedule1, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineObject95) SetSchedule(v NetworksNetworkIdSensorAlertsProfilesSchedule1)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineObject95) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConditions

`func (o *InlineObject95) GetConditions() []NetworksNetworkIdSensorAlertsProfilesConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InlineObject95) GetConditionsOk() (*[]NetworksNetworkIdSensorAlertsProfilesConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InlineObject95) SetConditions(v []NetworksNetworkIdSensorAlertsProfilesConditions)`

SetConditions sets Conditions field to given value.


### GetRecipients

`func (o *InlineObject95) GetRecipients() NetworksNetworkIdSensorAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject95) GetRecipientsOk() (*NetworksNetworkIdSensorAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject95) SetRecipients(v NetworksNetworkIdSensorAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineObject95) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject95) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject95) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject95) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject95) HasSerials() bool`

HasSerials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


