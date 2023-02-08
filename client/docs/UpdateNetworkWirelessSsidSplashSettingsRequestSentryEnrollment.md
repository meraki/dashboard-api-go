# UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemsManagerNetwork** | Pointer to [**UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork**](UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork.md) |  | [optional] 
**Strength** | Pointer to **string** | The strength of the enforcement of selected system types. Must be one of: &#39;focused&#39;, &#39;click-through&#39;, and &#39;strict&#39;. | [optional] 
**EnforcedSystems** | Pointer to **[]string** | The system types that the Sentry enforces. Must be included in: &#39;iOS, &#39;Android&#39;, &#39;macOS&#39;, and &#39;Windows&#39;. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment

`func NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment() *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment`

NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentWithDefaults

`func NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment`

NewUpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemsManagerNetwork

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetSystemsManagerNetwork() UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork`

GetSystemsManagerNetwork returns the SystemsManagerNetwork field if non-nil, zero value otherwise.

### GetSystemsManagerNetworkOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetSystemsManagerNetworkOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork, bool)`

GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsManagerNetwork

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) SetSystemsManagerNetwork(v UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollmentSystemsManagerNetwork)`

SetSystemsManagerNetwork sets SystemsManagerNetwork field to given value.

### HasSystemsManagerNetwork

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) HasSystemsManagerNetwork() bool`

HasSystemsManagerNetwork returns a boolean if a field has been set.

### GetStrength

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetEnforcedSystems

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetEnforcedSystems() []string`

GetEnforcedSystems returns the EnforcedSystems field if non-nil, zero value otherwise.

### GetEnforcedSystemsOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) GetEnforcedSystemsOk() (*[]string, bool)`

GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedSystems

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) SetEnforcedSystems(v []string)`

SetEnforcedSystems sets EnforcedSystems field to given value.

### HasEnforcedSystems

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) HasEnforcedSystems() bool`

HasEnforcedSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


