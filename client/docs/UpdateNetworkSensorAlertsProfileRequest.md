# UpdateNetworkSensorAlertsProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the sensor alert profile. | [optional] 
**Schedule** | Pointer to [**CreateNetworkSensorAlertsProfileRequestSchedule**](CreateNetworkSensorAlertsProfileRequestSchedule.md) |  | [optional] 
**Conditions** | Pointer to [**[]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner**](GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner.md) | List of conditions that will cause the profile to send an alert. | [optional] 
**Recipients** | Pointer to [**GetNetworkSensorAlertsProfiles200ResponseInnerRecipients**](GetNetworkSensorAlertsProfiles200ResponseInnerRecipients.md) |  | [optional] 
**Serials** | Pointer to **[]string** | List of device serials assigned to this sensor alert profile. | [optional] 

## Methods

### NewUpdateNetworkSensorAlertsProfileRequest

`func NewUpdateNetworkSensorAlertsProfileRequest() *UpdateNetworkSensorAlertsProfileRequest`

NewUpdateNetworkSensorAlertsProfileRequest instantiates a new UpdateNetworkSensorAlertsProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSensorAlertsProfileRequestWithDefaults

`func NewUpdateNetworkSensorAlertsProfileRequestWithDefaults() *UpdateNetworkSensorAlertsProfileRequest`

NewUpdateNetworkSensorAlertsProfileRequestWithDefaults instantiates a new UpdateNetworkSensorAlertsProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkSensorAlertsProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkSensorAlertsProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetSchedule() CreateNetworkSensorAlertsProfileRequestSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetScheduleOk() (*CreateNetworkSensorAlertsProfileRequestSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateNetworkSensorAlertsProfileRequest) SetSchedule(v CreateNetworkSensorAlertsProfileRequestSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateNetworkSensorAlertsProfileRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConditions

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetConditions() []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetConditionsOk() (*[]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateNetworkSensorAlertsProfileRequest) SetConditions(v []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UpdateNetworkSensorAlertsProfileRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetRecipients

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetRecipients() GetNetworkSensorAlertsProfiles200ResponseInnerRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetRecipientsOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *UpdateNetworkSensorAlertsProfileRequest) SetRecipients(v GetNetworkSensorAlertsProfiles200ResponseInnerRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *UpdateNetworkSensorAlertsProfileRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSerials

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *UpdateNetworkSensorAlertsProfileRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *UpdateNetworkSensorAlertsProfileRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *UpdateNetworkSensorAlertsProfileRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


