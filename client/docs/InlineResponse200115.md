# InlineResponse200115

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Monitored media server id | [optional] 
**Name** | Pointer to **string** | The name of the VoIP provider | [optional] 
**Address** | Pointer to **string** | The IP address (IPv4 only) or hostname of the media server to monitor | [optional] 
**BestEffortMonitoringEnabled** | Pointer to **bool** | Indicates that if the media server doesn&#39;t respond to ICMP pings, the nearest hop will be used in its stead | [optional] 

## Methods

### NewInlineResponse200115

`func NewInlineResponse200115() *InlineResponse200115`

NewInlineResponse200115 instantiates a new InlineResponse200115 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200115WithDefaults

`func NewInlineResponse200115WithDefaults() *InlineResponse200115`

NewInlineResponse200115WithDefaults instantiates a new InlineResponse200115 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200115) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200115) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200115) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200115) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200115) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200115) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200115) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200115) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200115) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200115) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200115) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200115) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBestEffortMonitoringEnabled

`func (o *InlineResponse200115) GetBestEffortMonitoringEnabled() bool`

GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field if non-nil, zero value otherwise.

### GetBestEffortMonitoringEnabledOk

`func (o *InlineResponse200115) GetBestEffortMonitoringEnabledOk() (*bool, bool)`

GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestEffortMonitoringEnabled

`func (o *InlineResponse200115) SetBestEffortMonitoringEnabled(v bool)`

SetBestEffortMonitoringEnabled sets BestEffortMonitoringEnabled field to given value.

### HasBestEffortMonitoringEnabled

`func (o *InlineResponse200115) HasBestEffortMonitoringEnabled() bool`

HasBestEffortMonitoringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


