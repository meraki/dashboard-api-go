# UpdateNetworkApplianceSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientTrackingMethod** | Pointer to **string** | Client tracking method of a network | [optional] 
**DeploymentMode** | Pointer to **string** | Deployment mode of a network | [optional] 
**DynamicDns** | Pointer to [**UpdateNetworkApplianceSettingsRequestDynamicDns**](UpdateNetworkApplianceSettingsRequestDynamicDns.md) |  | [optional] 

## Methods

### NewUpdateNetworkApplianceSettingsRequest

`func NewUpdateNetworkApplianceSettingsRequest() *UpdateNetworkApplianceSettingsRequest`

NewUpdateNetworkApplianceSettingsRequest instantiates a new UpdateNetworkApplianceSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceSettingsRequestWithDefaults

`func NewUpdateNetworkApplianceSettingsRequestWithDefaults() *UpdateNetworkApplianceSettingsRequest`

NewUpdateNetworkApplianceSettingsRequestWithDefaults instantiates a new UpdateNetworkApplianceSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTrackingMethod

`func (o *UpdateNetworkApplianceSettingsRequest) GetClientTrackingMethod() string`

GetClientTrackingMethod returns the ClientTrackingMethod field if non-nil, zero value otherwise.

### GetClientTrackingMethodOk

`func (o *UpdateNetworkApplianceSettingsRequest) GetClientTrackingMethodOk() (*string, bool)`

GetClientTrackingMethodOk returns a tuple with the ClientTrackingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTrackingMethod

`func (o *UpdateNetworkApplianceSettingsRequest) SetClientTrackingMethod(v string)`

SetClientTrackingMethod sets ClientTrackingMethod field to given value.

### HasClientTrackingMethod

`func (o *UpdateNetworkApplianceSettingsRequest) HasClientTrackingMethod() bool`

HasClientTrackingMethod returns a boolean if a field has been set.

### GetDeploymentMode

`func (o *UpdateNetworkApplianceSettingsRequest) GetDeploymentMode() string`

GetDeploymentMode returns the DeploymentMode field if non-nil, zero value otherwise.

### GetDeploymentModeOk

`func (o *UpdateNetworkApplianceSettingsRequest) GetDeploymentModeOk() (*string, bool)`

GetDeploymentModeOk returns a tuple with the DeploymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMode

`func (o *UpdateNetworkApplianceSettingsRequest) SetDeploymentMode(v string)`

SetDeploymentMode sets DeploymentMode field to given value.

### HasDeploymentMode

`func (o *UpdateNetworkApplianceSettingsRequest) HasDeploymentMode() bool`

HasDeploymentMode returns a boolean if a field has been set.

### GetDynamicDns

`func (o *UpdateNetworkApplianceSettingsRequest) GetDynamicDns() UpdateNetworkApplianceSettingsRequestDynamicDns`

GetDynamicDns returns the DynamicDns field if non-nil, zero value otherwise.

### GetDynamicDnsOk

`func (o *UpdateNetworkApplianceSettingsRequest) GetDynamicDnsOk() (*UpdateNetworkApplianceSettingsRequestDynamicDns, bool)`

GetDynamicDnsOk returns a tuple with the DynamicDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicDns

`func (o *UpdateNetworkApplianceSettingsRequest) SetDynamicDns(v UpdateNetworkApplianceSettingsRequestDynamicDns)`

SetDynamicDns sets DynamicDns field to given value.

### HasDynamicDns

`func (o *UpdateNetworkApplianceSettingsRequest) HasDynamicDns() bool`

HasDynamicDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


