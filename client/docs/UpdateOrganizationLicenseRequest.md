# UpdateOrganizationLicenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSerial** | Pointer to **string** | The serial number of the device to assign this license to. Set this to  null to unassign the license. If a different license is already active on the device, this parameter will control queueing/dequeuing this license. | [optional] 

## Methods

### NewUpdateOrganizationLicenseRequest

`func NewUpdateOrganizationLicenseRequest() *UpdateOrganizationLicenseRequest`

NewUpdateOrganizationLicenseRequest instantiates a new UpdateOrganizationLicenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationLicenseRequestWithDefaults

`func NewUpdateOrganizationLicenseRequestWithDefaults() *UpdateOrganizationLicenseRequest`

NewUpdateOrganizationLicenseRequestWithDefaults instantiates a new UpdateOrganizationLicenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSerial

`func (o *UpdateOrganizationLicenseRequest) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *UpdateOrganizationLicenseRequest) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *UpdateOrganizationLicenseRequest) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *UpdateOrganizationLicenseRequest) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


