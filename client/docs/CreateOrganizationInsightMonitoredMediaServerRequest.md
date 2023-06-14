# CreateOrganizationInsightMonitoredMediaServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the VoIP provider | 
**Address** | **string** | The IP address (IPv4 only) or hostname of the media server to monitor | 
**BestEffortMonitoringEnabled** | Pointer to **bool** | Indicates that if the media server doesn&#39;t respond to ICMP pings, the nearest hop will be used in its stead. | [optional] 

## Methods

### NewCreateOrganizationInsightMonitoredMediaServerRequest

`func NewCreateOrganizationInsightMonitoredMediaServerRequest(name string, address string, ) *CreateOrganizationInsightMonitoredMediaServerRequest`

NewCreateOrganizationInsightMonitoredMediaServerRequest instantiates a new CreateOrganizationInsightMonitoredMediaServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInsightMonitoredMediaServerRequestWithDefaults

`func NewCreateOrganizationInsightMonitoredMediaServerRequestWithDefaults() *CreateOrganizationInsightMonitoredMediaServerRequest`

NewCreateOrganizationInsightMonitoredMediaServerRequestWithDefaults instantiates a new CreateOrganizationInsightMonitoredMediaServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBestEffortMonitoringEnabled

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetBestEffortMonitoringEnabled() bool`

GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field if non-nil, zero value otherwise.

### GetBestEffortMonitoringEnabledOk

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) GetBestEffortMonitoringEnabledOk() (*bool, bool)`

GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestEffortMonitoringEnabled

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) SetBestEffortMonitoringEnabled(v bool)`

SetBestEffortMonitoringEnabled sets BestEffortMonitoringEnabled field to given value.

### HasBestEffortMonitoringEnabled

`func (o *CreateOrganizationInsightMonitoredMediaServerRequest) HasBestEffortMonitoringEnabled() bool`

HasBestEffortMonitoringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


