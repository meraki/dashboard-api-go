# CreateOrganizationCameraRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new role. Must be unique. This parameter is required. | 
**AppliedOnDevices** | Pointer to [**[]CreateOrganizationCameraRoleRequestAppliedOnDevicesInner**](CreateOrganizationCameraRoleRequestAppliedOnDevicesInner.md) | Device tag on which this specified permission is applied. | [optional] 
**AppliedOnNetworks** | Pointer to [**[]CreateOrganizationCameraRoleRequestAppliedOnNetworksInner**](CreateOrganizationCameraRoleRequestAppliedOnNetworksInner.md) | Network tag on which this specified permission is applied. | [optional] 
**AppliedOrgWide** | Pointer to [**[]CreateOrganizationCameraRoleRequestAppliedOrgWideInner**](CreateOrganizationCameraRoleRequestAppliedOrgWideInner.md) | Permissions to be applied org wide. | [optional] 

## Methods

### NewCreateOrganizationCameraRoleRequest

`func NewCreateOrganizationCameraRoleRequest(name string, ) *CreateOrganizationCameraRoleRequest`

NewCreateOrganizationCameraRoleRequest instantiates a new CreateOrganizationCameraRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationCameraRoleRequestWithDefaults

`func NewCreateOrganizationCameraRoleRequestWithDefaults() *CreateOrganizationCameraRoleRequest`

NewCreateOrganizationCameraRoleRequestWithDefaults instantiates a new CreateOrganizationCameraRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationCameraRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationCameraRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationCameraRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAppliedOnDevices

`func (o *CreateOrganizationCameraRoleRequest) GetAppliedOnDevices() []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner`

GetAppliedOnDevices returns the AppliedOnDevices field if non-nil, zero value otherwise.

### GetAppliedOnDevicesOk

`func (o *CreateOrganizationCameraRoleRequest) GetAppliedOnDevicesOk() (*[]CreateOrganizationCameraRoleRequestAppliedOnDevicesInner, bool)`

GetAppliedOnDevicesOk returns a tuple with the AppliedOnDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOnDevices

`func (o *CreateOrganizationCameraRoleRequest) SetAppliedOnDevices(v []CreateOrganizationCameraRoleRequestAppliedOnDevicesInner)`

SetAppliedOnDevices sets AppliedOnDevices field to given value.

### HasAppliedOnDevices

`func (o *CreateOrganizationCameraRoleRequest) HasAppliedOnDevices() bool`

HasAppliedOnDevices returns a boolean if a field has been set.

### GetAppliedOnNetworks

`func (o *CreateOrganizationCameraRoleRequest) GetAppliedOnNetworks() []CreateOrganizationCameraRoleRequestAppliedOnNetworksInner`

GetAppliedOnNetworks returns the AppliedOnNetworks field if non-nil, zero value otherwise.

### GetAppliedOnNetworksOk

`func (o *CreateOrganizationCameraRoleRequest) GetAppliedOnNetworksOk() (*[]CreateOrganizationCameraRoleRequestAppliedOnNetworksInner, bool)`

GetAppliedOnNetworksOk returns a tuple with the AppliedOnNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOnNetworks

`func (o *CreateOrganizationCameraRoleRequest) SetAppliedOnNetworks(v []CreateOrganizationCameraRoleRequestAppliedOnNetworksInner)`

SetAppliedOnNetworks sets AppliedOnNetworks field to given value.

### HasAppliedOnNetworks

`func (o *CreateOrganizationCameraRoleRequest) HasAppliedOnNetworks() bool`

HasAppliedOnNetworks returns a boolean if a field has been set.

### GetAppliedOrgWide

`func (o *CreateOrganizationCameraRoleRequest) GetAppliedOrgWide() []CreateOrganizationCameraRoleRequestAppliedOrgWideInner`

GetAppliedOrgWide returns the AppliedOrgWide field if non-nil, zero value otherwise.

### GetAppliedOrgWideOk

`func (o *CreateOrganizationCameraRoleRequest) GetAppliedOrgWideOk() (*[]CreateOrganizationCameraRoleRequestAppliedOrgWideInner, bool)`

GetAppliedOrgWideOk returns a tuple with the AppliedOrgWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOrgWide

`func (o *CreateOrganizationCameraRoleRequest) SetAppliedOrgWide(v []CreateOrganizationCameraRoleRequestAppliedOrgWideInner)`

SetAppliedOrgWide sets AppliedOrgWide field to given value.

### HasAppliedOrgWide

`func (o *CreateOrganizationCameraRoleRequest) HasAppliedOrgWide() bool`

HasAppliedOrgWide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


