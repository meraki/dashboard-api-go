# CreateNetworkSensorAlertsProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the sensor alert profile. | 
**Schedule** | Pointer to [**CreateNetworkSensorAlertsProfileRequestSchedule**](CreateNetworkSensorAlertsProfileRequestSchedule.md) |  | [optional] 
**Conditions** | [**[]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner**](GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner.md) | List of conditions that will cause the profile to send an alert. | 
**Recipients** | Pointer to [**GetNetworkSensorAlertsProfiles200ResponseInnerRecipients**](GetNetworkSensorAlertsProfiles200ResponseInnerRecipients.md) |  | [optional] 
**Serials** | Pointer to **[]string** | List of device serials assigned to this sensor alert profile. | [optional] 

## Methods

### NewCreateNetworkSensorAlertsProfileRequest

`func NewCreateNetworkSensorAlertsProfileRequest(name string, conditions []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, ) *CreateNetworkSensorAlertsProfileRequest`

NewCreateNetworkSensorAlertsProfileRequest instantiates a new CreateNetworkSensorAlertsProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSensorAlertsProfileRequestWithDefaults

`func NewCreateNetworkSensorAlertsProfileRequestWithDefaults() *CreateNetworkSensorAlertsProfileRequest`

NewCreateNetworkSensorAlertsProfileRequestWithDefaults instantiates a new CreateNetworkSensorAlertsProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkSensorAlertsProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkSensorAlertsProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkSensorAlertsProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *CreateNetworkSensorAlertsProfileRequest) GetSchedule() CreateNetworkSensorAlertsProfileRequestSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateNetworkSensorAlertsProfileRequest) GetScheduleOk() (*CreateNetworkSensorAlertsProfileRequestSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateNetworkSensorAlertsProfileRequest) SetSchedule(v CreateNetworkSensorAlertsProfileRequestSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateNetworkSensorAlertsProfileRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConditions

`func (o *CreateNetworkSensorAlertsProfileRequest) GetConditions() []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateNetworkSensorAlertsProfileRequest) GetConditionsOk() (*[]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateNetworkSensorAlertsProfileRequest) SetConditions(v []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner)`

SetConditions sets Conditions field to given value.


### GetRecipients

`func (o *CreateNetworkSensorAlertsProfileRequest) GetRecipients() GetNetworkSensorAlertsProfiles200ResponseInnerRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateNetworkSensorAlertsProfileRequest) GetRecipientsOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateNetworkSensorAlertsProfileRequest) SetRecipients(v GetNetworkSensorAlertsProfiles200ResponseInnerRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CreateNetworkSensorAlertsProfileRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSerials

`func (o *CreateNetworkSensorAlertsProfileRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *CreateNetworkSensorAlertsProfileRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *CreateNetworkSensorAlertsProfileRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *CreateNetworkSensorAlertsProfileRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


