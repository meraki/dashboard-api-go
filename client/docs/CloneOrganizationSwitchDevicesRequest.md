# CloneOrganizationSwitchDevicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceSerial** | **string** | Serial number of the source switch (must be on a network not bound to a template) | 
**TargetSerials** | **[]string** | Array of serial numbers of one or more target switches (must be on a network not bound to a template) | 

## Methods

### NewCloneOrganizationSwitchDevicesRequest

`func NewCloneOrganizationSwitchDevicesRequest(sourceSerial string, targetSerials []string, ) *CloneOrganizationSwitchDevicesRequest`

NewCloneOrganizationSwitchDevicesRequest instantiates a new CloneOrganizationSwitchDevicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneOrganizationSwitchDevicesRequestWithDefaults

`func NewCloneOrganizationSwitchDevicesRequestWithDefaults() *CloneOrganizationSwitchDevicesRequest`

NewCloneOrganizationSwitchDevicesRequestWithDefaults instantiates a new CloneOrganizationSwitchDevicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSerial

`func (o *CloneOrganizationSwitchDevicesRequest) GetSourceSerial() string`

GetSourceSerial returns the SourceSerial field if non-nil, zero value otherwise.

### GetSourceSerialOk

`func (o *CloneOrganizationSwitchDevicesRequest) GetSourceSerialOk() (*string, bool)`

GetSourceSerialOk returns a tuple with the SourceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSerial

`func (o *CloneOrganizationSwitchDevicesRequest) SetSourceSerial(v string)`

SetSourceSerial sets SourceSerial field to given value.


### GetTargetSerials

`func (o *CloneOrganizationSwitchDevicesRequest) GetTargetSerials() []string`

GetTargetSerials returns the TargetSerials field if non-nil, zero value otherwise.

### GetTargetSerialsOk

`func (o *CloneOrganizationSwitchDevicesRequest) GetTargetSerialsOk() (*[]string, bool)`

GetTargetSerialsOk returns a tuple with the TargetSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSerials

`func (o *CloneOrganizationSwitchDevicesRequest) SetTargetSerials(v []string)`

SetTargetSerials sets TargetSerials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


