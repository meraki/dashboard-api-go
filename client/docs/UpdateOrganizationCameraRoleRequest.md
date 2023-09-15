# UpdateOrganizationCameraRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the new role. Must be unique. | [optional] 
**AppliedOnDevices** | Pointer to [**[]CreateOrganizationCameraRoleRequestAppliedOnDevicesInner**](CreateOrganizationCameraRoleRequestAppliedOnDevicesInner.md) | Device tag on which this specified permission is applied. | [optional] 
**AppliedOnNetworks** | Pointer to [**[]UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner**](UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner.md) | Network tag on which this specified permission is applied. | [optional] 
**AppliedOrgWide** | Pointer to [**[]CreateOrganizationCameraRoleRequestAppliedOrgWideInner**](CreateOrganizationCameraRoleRequestAppliedOrgWideInner.md) | Permissions to be applied org wide. | [optional] 

## Methods

### NewUpdateOrganizationCameraRoleRequest

`func NewUpdateOrganizationCameraRoleRequest() *UpdateOrganizationCameraRoleRequest`

NewUpdateOrganizationCameraRoleRequest instantiates a new UpdateOrganizationCameraRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationCameraRoleRequestWithDefaults

`func NewUpdateOrganizationCameraRoleRequestWithDefaults() *UpdateOrganizationCameraRoleRequest`

NewUpdateOrganizationCameraRoleRequestWithDefaults instantiates a new UpdateOrganizationCameraRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrganizationCameraRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationCameraRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationCameraRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationCameraRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAppliedOnDevices

`func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnDevices() []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner`

GetAppliedOnDevices returns the AppliedOnDevices field if non-nil, zero value otherwise.

### GetAppliedOnDevicesOk

`func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnDevicesOk() (*[]CreateOrganizationCameraRoleRequestAppliedOnDevicesInner, bool)`

GetAppliedOnDevicesOk returns a tuple with the AppliedOnDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOnDevices

`func (o *UpdateOrganizationCameraRoleRequest) SetAppliedOnDevices(v []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner)`

SetAppliedOnDevices sets AppliedOnDevices field to given value.

### HasAppliedOnDevices

`func (o *UpdateOrganizationCameraRoleRequest) HasAppliedOnDevices() bool`

HasAppliedOnDevices returns a boolean if a field has been set.

### GetAppliedOnNetworks

`func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnNetworks() []UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner`

GetAppliedOnNetworks returns the AppliedOnNetworks field if non-nil, zero value otherwise.

### GetAppliedOnNetworksOk

`func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOnNetworksOk() (*[]UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner, bool)`

GetAppliedOnNetworksOk returns a tuple with the AppliedOnNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOnNetworks

`func (o *UpdateOrganizationCameraRoleRequest) SetAppliedOnNetworks(v []UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner)`

SetAppliedOnNetworks sets AppliedOnNetworks field to given value.

### HasAppliedOnNetworks

`func (o *UpdateOrganizationCameraRoleRequest) HasAppliedOnNetworks() bool`

HasAppliedOnNetworks returns a boolean if a field has been set.

### GetAppliedOrgWide

`func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOrgWide() []CreateOrganizationCameraRoleRequestAppliedOrgWideInner`

GetAppliedOrgWide returns the AppliedOrgWide field if non-nil, zero value otherwise.

### GetAppliedOrgWideOk

`func (o *UpdateOrganizationCameraRoleRequest) GetAppliedOrgWideOk() (*[]CreateOrganizationCameraRoleRequestAppliedOrgWideInner, bool)`

GetAppliedOrgWideOk returns a tuple with the AppliedOrgWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOrgWide

`func (o *UpdateOrganizationCameraRoleRequest) SetAppliedOrgWide(v []CreateOrganizationCameraRoleRequestAppliedOrgWideInner)`

SetAppliedOrgWide sets AppliedOrgWide field to given value.

### HasAppliedOrgWide

`func (o *UpdateOrganizationCameraRoleRequest) HasAppliedOrgWide() bool`

HasAppliedOrgWide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


