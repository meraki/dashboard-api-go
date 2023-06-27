# GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether Secure Port is turned on for this port. | [optional] 
**Active** | Pointer to **bool** | Whether Secure Port is currently active for this port. | [optional] 
**AuthenticationStatus** | Pointer to **string** | The current Secure Port status. | [optional] 
**ConfigOverrides** | Pointer to [**GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides**](GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides.md) |  | [optional] 

## Methods

### NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePort

`func NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePort() *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort`

NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePort instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortWithDefaults

`func NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortWithDefaults() *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort`

NewGetDeviceSwitchPortsStatuses200ResponseInnerSecurePortWithDefaults instantiates a new GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetActive

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAuthenticationStatus

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetAuthenticationStatus() string`

GetAuthenticationStatus returns the AuthenticationStatus field if non-nil, zero value otherwise.

### GetAuthenticationStatusOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetAuthenticationStatusOk() (*string, bool)`

GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatus

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) SetAuthenticationStatus(v string)`

SetAuthenticationStatus sets AuthenticationStatus field to given value.

### HasAuthenticationStatus

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) HasAuthenticationStatus() bool`

HasAuthenticationStatus returns a boolean if a field has been set.

### GetConfigOverrides

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetConfigOverrides() GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides`

GetConfigOverrides returns the ConfigOverrides field if non-nil, zero value otherwise.

### GetConfigOverridesOk

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) GetConfigOverridesOk() (*GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides, bool)`

GetConfigOverridesOk returns a tuple with the ConfigOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOverrides

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) SetConfigOverrides(v GetDeviceSwitchPortsStatuses200ResponseInnerSecurePortConfigOverrides)`

SetConfigOverrides sets ConfigOverrides field to given value.

### HasConfigOverrides

`func (o *GetDeviceSwitchPortsStatuses200ResponseInnerSecurePort) HasConfigOverrides() bool`

HasConfigOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


