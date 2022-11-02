# UpdateNetworkWirelessSsidSplashSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplashUrl** | Pointer to **string** | [optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see &#39;useSplashUrl&#39; | [optional] 
**UseSplashUrl** | Pointer to **bool** | [optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID&#39;s access control settings, it may not be possible to use the custom splash URL. | [optional] 
**SplashTimeout** | Pointer to **int32** | Splash timeout in minutes. This will determine how often users will see the splash page. | [optional] 
**RedirectUrl** | Pointer to **string** | The custom redirect URL where the users will go after the splash page. | [optional] 
**UseRedirectUrl** | Pointer to **bool** | The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true. | [optional] 
**WelcomeMessage** | Pointer to **string** | The welcome message for the users on the splash page. | [optional] 
**SplashLogo** | Pointer to [**UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo**](UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo.md) |  | [optional] 
**SplashImage** | Pointer to [**UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage**](UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage.md) |  | [optional] 
**SplashPrepaidFront** | Pointer to [**UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront**](UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront.md) |  | [optional] 
**BlockAllTrafficBeforeSignOn** | Pointer to **bool** | How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged. | [optional] 
**ControllerDisconnectionBehavior** | Pointer to **string** | How login attempts should be handled when the controller is unreachable. Can be either &#39;open&#39;, &#39;restricted&#39;, or &#39;default&#39;. | [optional] 
**AllowSimultaneousLogins** | Pointer to **bool** | Whether or not to allow simultaneous logins from different devices. | [optional] 
**GuestSponsorship** | Pointer to [**UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship**](UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship.md) |  | [optional] 
**Billing** | Pointer to [**UpdateNetworkWirelessSsidSplashSettingsRequestBilling**](UpdateNetworkWirelessSsidSplashSettingsRequestBilling.md) |  | [optional] 
**SentryEnrollment** | Pointer to [**UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment**](UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidSplashSettingsRequest

`func NewUpdateNetworkWirelessSsidSplashSettingsRequest() *UpdateNetworkWirelessSsidSplashSettingsRequest`

NewUpdateNetworkWirelessSsidSplashSettingsRequest instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidSplashSettingsRequestWithDefaults

`func NewUpdateNetworkWirelessSsidSplashSettingsRequestWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequest`

NewUpdateNetworkWirelessSsidSplashSettingsRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSplashUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashUrl() string`

GetSplashUrl returns the SplashUrl field if non-nil, zero value otherwise.

### GetSplashUrlOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashUrlOk() (*string, bool)`

GetSplashUrlOk returns a tuple with the SplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashUrl(v string)`

SetSplashUrl sets SplashUrl field to given value.

### HasSplashUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashUrl() bool`

HasSplashUrl returns a boolean if a field has been set.

### GetUseSplashUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseSplashUrl() bool`

GetUseSplashUrl returns the UseSplashUrl field if non-nil, zero value otherwise.

### GetUseSplashUrlOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseSplashUrlOk() (*bool, bool)`

GetUseSplashUrlOk returns a tuple with the UseSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplashUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetUseSplashUrl(v bool)`

SetUseSplashUrl sets UseSplashUrl field to given value.

### HasUseSplashUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasUseSplashUrl() bool`

HasUseSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashTimeout() int32`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashTimeoutOk() (*int32, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashTimeout(v int32)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetUseRedirectUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseRedirectUrl() bool`

GetUseRedirectUrl returns the UseRedirectUrl field if non-nil, zero value otherwise.

### GetUseRedirectUrlOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseRedirectUrlOk() (*bool, bool)`

GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRedirectUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetUseRedirectUrl(v bool)`

SetUseRedirectUrl sets UseRedirectUrl field to given value.

### HasUseRedirectUrl

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasUseRedirectUrl() bool`

HasUseRedirectUrl returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### GetSplashLogo

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashLogo() UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo`

GetSplashLogo returns the SplashLogo field if non-nil, zero value otherwise.

### GetSplashLogoOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashLogoOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo, bool)`

GetSplashLogoOk returns a tuple with the SplashLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashLogo

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashLogo(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo)`

SetSplashLogo sets SplashLogo field to given value.

### HasSplashLogo

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashLogo() bool`

HasSplashLogo returns a boolean if a field has been set.

### GetSplashImage

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashImage() UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage`

GetSplashImage returns the SplashImage field if non-nil, zero value otherwise.

### GetSplashImageOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashImageOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage, bool)`

GetSplashImageOk returns a tuple with the SplashImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImage

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashImage(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage)`

SetSplashImage sets SplashImage field to given value.

### HasSplashImage

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashImage() bool`

HasSplashImage returns a boolean if a field has been set.

### GetSplashPrepaidFront

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashPrepaidFront() UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront`

GetSplashPrepaidFront returns the SplashPrepaidFront field if non-nil, zero value otherwise.

### GetSplashPrepaidFrontOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashPrepaidFrontOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront, bool)`

GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPrepaidFront

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashPrepaidFront(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront)`

SetSplashPrepaidFront sets SplashPrepaidFront field to given value.

### HasSplashPrepaidFront

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashPrepaidFront() bool`

HasSplashPrepaidFront returns a boolean if a field has been set.

### GetBlockAllTrafficBeforeSignOn

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBlockAllTrafficBeforeSignOn() bool`

GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field if non-nil, zero value otherwise.

### GetBlockAllTrafficBeforeSignOnOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool)`

GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllTrafficBeforeSignOn

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetBlockAllTrafficBeforeSignOn(v bool)`

SetBlockAllTrafficBeforeSignOn sets BlockAllTrafficBeforeSignOn field to given value.

### HasBlockAllTrafficBeforeSignOn

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasBlockAllTrafficBeforeSignOn() bool`

HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.

### GetControllerDisconnectionBehavior

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetControllerDisconnectionBehavior() string`

GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field if non-nil, zero value otherwise.

### GetControllerDisconnectionBehaviorOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetControllerDisconnectionBehaviorOk() (*string, bool)`

GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDisconnectionBehavior

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetControllerDisconnectionBehavior(v string)`

SetControllerDisconnectionBehavior sets ControllerDisconnectionBehavior field to given value.

### HasControllerDisconnectionBehavior

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasControllerDisconnectionBehavior() bool`

HasControllerDisconnectionBehavior returns a boolean if a field has been set.

### GetAllowSimultaneousLogins

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetAllowSimultaneousLogins() bool`

GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field if non-nil, zero value otherwise.

### GetAllowSimultaneousLoginsOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetAllowSimultaneousLoginsOk() (*bool, bool)`

GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneousLogins

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetAllowSimultaneousLogins(v bool)`

SetAllowSimultaneousLogins sets AllowSimultaneousLogins field to given value.

### HasAllowSimultaneousLogins

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasAllowSimultaneousLogins() bool`

HasAllowSimultaneousLogins returns a boolean if a field has been set.

### GetGuestSponsorship

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetGuestSponsorship() UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship`

GetGuestSponsorship returns the GuestSponsorship field if non-nil, zero value otherwise.

### GetGuestSponsorshipOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetGuestSponsorshipOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship, bool)`

GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSponsorship

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetGuestSponsorship(v UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship)`

SetGuestSponsorship sets GuestSponsorship field to given value.

### HasGuestSponsorship

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasGuestSponsorship() bool`

HasGuestSponsorship returns a boolean if a field has been set.

### GetBilling

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBilling() UpdateNetworkWirelessSsidSplashSettingsRequestBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBillingOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetBilling(v UpdateNetworkWirelessSsidSplashSettingsRequestBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetSentryEnrollment

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSentryEnrollment() UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment`

GetSentryEnrollment returns the SentryEnrollment field if non-nil, zero value otherwise.

### GetSentryEnrollmentOk

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSentryEnrollmentOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment, bool)`

GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryEnrollment

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSentryEnrollment(v UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment)`

SetSentryEnrollment sets SentryEnrollment field to given value.

### HasSentryEnrollment

`func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSentryEnrollment() bool`

HasSentryEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


