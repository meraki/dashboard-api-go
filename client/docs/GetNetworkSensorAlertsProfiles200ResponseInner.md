# GetNetworkSensorAlertsProfiles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the sensor alert profile. | [optional] 
**Name** | Pointer to **string** | Name of the sensor alert profile. | [optional] 
**Schedule** | Pointer to [**GetNetworkSensorAlertsProfiles200ResponseInnerSchedule**](GetNetworkSensorAlertsProfiles200ResponseInnerSchedule.md) |  | [optional] 
**Conditions** | [**[]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner**](GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner.md) | List of conditions that will cause the profile to send an alert. | 
**Recipients** | Pointer to [**GetNetworkSensorAlertsProfiles200ResponseInnerRecipients**](GetNetworkSensorAlertsProfiles200ResponseInnerRecipients.md) |  | [optional] 
**Serials** | Pointer to **[]string** | List of device serials assigned to this sensor alert profile. | [optional] 

## Methods

### NewGetNetworkSensorAlertsProfiles200ResponseInner

`func NewGetNetworkSensorAlertsProfiles200ResponseInner(conditions []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, ) *GetNetworkSensorAlertsProfiles200ResponseInner`

NewGetNetworkSensorAlertsProfiles200ResponseInner instantiates a new GetNetworkSensorAlertsProfiles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSensorAlertsProfiles200ResponseInnerWithDefaults

`func NewGetNetworkSensorAlertsProfiles200ResponseInnerWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInner`

NewGetNetworkSensorAlertsProfiles200ResponseInnerWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetSchedule() GetNetworkSensorAlertsProfiles200ResponseInnerSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetScheduleOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetSchedule(v GetNetworkSensorAlertsProfiles200ResponseInnerSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConditions

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetConditions() []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetConditionsOk() (*[]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetConditions(v []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner)`

SetConditions sets Conditions field to given value.


### GetRecipients

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetRecipients() GetNetworkSensorAlertsProfiles200ResponseInnerRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetRecipientsOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetRecipients(v GetNetworkSensorAlertsProfiles200ResponseInnerRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSerials

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *GetNetworkSensorAlertsProfiles200ResponseInner) HasSerials() bool`

HasSerials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


