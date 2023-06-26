# CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogEvent** | **string** | The type of log event this is recording, e.g. download or opening a banner | 
**Timestamp** | **int32** | A JavaScript UTC datetime stamp for when the even occurred | 
**TargetOS** | Pointer to **string** | The name of the onboarding distro being downloaded | [optional] 
**Request** | Pointer to **string** | Used to describe if this event was the result of a redirect. E.g. a query param if an info banner is being used | [optional] 

## Methods

### NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest

`func NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(logEvent string, timestamp int32, ) *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest`

NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequestWithDefaults

`func NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequestWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest`

NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequestWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogEvent

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetLogEvent() string`

GetLogEvent returns the LogEvent field if non-nil, zero value otherwise.

### GetLogEventOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetLogEventOk() (*string, bool)`

GetLogEventOk returns a tuple with the LogEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEvent

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetLogEvent(v string)`

SetLogEvent sets LogEvent field to given value.


### GetTimestamp

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTargetOS

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTargetOS() string`

GetTargetOS returns the TargetOS field if non-nil, zero value otherwise.

### GetTargetOSOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTargetOSOk() (*string, bool)`

GetTargetOSOk returns a tuple with the TargetOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOS

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetTargetOS(v string)`

SetTargetOS sets TargetOS field to given value.

### HasTargetOS

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) HasTargetOS() bool`

HasTargetOS returns a boolean if a field has been set.

### GetRequest

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


