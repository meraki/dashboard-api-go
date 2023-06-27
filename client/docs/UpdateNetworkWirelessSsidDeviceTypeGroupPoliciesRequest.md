# UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID device type group policies are enabled. | [optional] 
**DeviceTypePolicies** | Pointer to [**[]UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner**](UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner.md) | List of device type policies. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest

`func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest`

NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest instantiates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestWithDefaults

`func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestWithDefaults() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest`

NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDeviceTypePolicies

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetDeviceTypePolicies() []UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner`

GetDeviceTypePolicies returns the DeviceTypePolicies field if non-nil, zero value otherwise.

### GetDeviceTypePoliciesOk

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetDeviceTypePoliciesOk() (*[]UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner, bool)`

GetDeviceTypePoliciesOk returns a tuple with the DeviceTypePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypePolicies

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) SetDeviceTypePolicies(v []UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner)`

SetDeviceTypePolicies sets DeviceTypePolicies field to given value.

### HasDeviceTypePolicies

`func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) HasDeviceTypePolicies() bool`

HasDeviceTypePolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


