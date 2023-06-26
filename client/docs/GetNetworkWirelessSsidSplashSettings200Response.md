# GetNetworkWirelessSsidSplashSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | SSID number | [optional] 
**SplashPage** | Pointer to **string** | The type of splash page for this SSID | [optional] 
**UseSplashUrl** | Pointer to **bool** | Boolean indicating whether the users will be redirected to the custom splash url | [optional] 
**SplashUrl** | Pointer to **string** | The custom splash URL of the click-through splash page. | [optional] 
**SplashTimeout** | Pointer to **int32** | Splash timeout in minutes. | [optional] 
**RedirectUrl** | Pointer to **string** | The custom redirect URL where the users will go after the splash page. | [optional] 
**UseRedirectUrl** | Pointer to **bool** | The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. | [optional] 
**WelcomeMessage** | Pointer to **string** | The welcome message for the users on the splash page. | [optional] 
**SplashLogo** | Pointer to [**GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo**](GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo.md) |  | [optional] 
**SplashImage** | Pointer to [**GetNetworkWirelessSsidSplashSettings200ResponseSplashImage**](GetNetworkWirelessSsidSplashSettings200ResponseSplashImage.md) |  | [optional] 
**SplashPrepaidFront** | Pointer to [**GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront**](GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront.md) |  | [optional] 
**GuestSponsorship** | Pointer to [**GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship**](GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship.md) |  | [optional] 
**BlockAllTrafficBeforeSignOn** | Pointer to **bool** | How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged. | [optional] 
**ControllerDisconnectionBehavior** | Pointer to **string** | How login attempts should be handled when the controller is unreachable. | [optional] 
**AllowSimultaneousLogins** | Pointer to **bool** | Whether or not to allow simultaneous logins from different devices. | [optional] 
**Billing** | Pointer to [**GetNetworkWirelessSsidSplashSettings200ResponseBilling**](GetNetworkWirelessSsidSplashSettings200ResponseBilling.md) |  | [optional] 
**SentryEnrollment** | Pointer to [**GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment**](GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment.md) |  | [optional] 
**SelfRegistration** | Pointer to [**GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration**](GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration.md) |  | [optional] 

## Methods

### NewGetNetworkWirelessSsidSplashSettings200Response

`func NewGetNetworkWirelessSsidSplashSettings200Response() *GetNetworkWirelessSsidSplashSettings200Response`

NewGetNetworkWirelessSsidSplashSettings200Response instantiates a new GetNetworkWirelessSsidSplashSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsidSplashSettings200ResponseWithDefaults

`func NewGetNetworkWirelessSsidSplashSettings200ResponseWithDefaults() *GetNetworkWirelessSsidSplashSettings200Response`

NewGetNetworkWirelessSsidSplashSettings200ResponseWithDefaults instantiates a new GetNetworkWirelessSsidSplashSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetSplashPage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetUseSplashUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseSplashUrl() bool`

GetUseSplashUrl returns the UseSplashUrl field if non-nil, zero value otherwise.

### GetUseSplashUrlOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseSplashUrlOk() (*bool, bool)`

GetUseSplashUrlOk returns a tuple with the UseSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplashUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetUseSplashUrl(v bool)`

SetUseSplashUrl sets UseSplashUrl field to given value.

### HasUseSplashUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasUseSplashUrl() bool`

HasUseSplashUrl returns a boolean if a field has been set.

### GetSplashUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashUrl() string`

GetSplashUrl returns the SplashUrl field if non-nil, zero value otherwise.

### GetSplashUrlOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashUrlOk() (*string, bool)`

GetSplashUrlOk returns a tuple with the SplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashUrl(v string)`

SetSplashUrl sets SplashUrl field to given value.

### HasSplashUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashUrl() bool`

HasSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashTimeout() int32`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashTimeoutOk() (*int32, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashTimeout(v int32)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetUseRedirectUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseRedirectUrl() bool`

GetUseRedirectUrl returns the UseRedirectUrl field if non-nil, zero value otherwise.

### GetUseRedirectUrlOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetUseRedirectUrlOk() (*bool, bool)`

GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRedirectUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetUseRedirectUrl(v bool)`

SetUseRedirectUrl sets UseRedirectUrl field to given value.

### HasUseRedirectUrl

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasUseRedirectUrl() bool`

HasUseRedirectUrl returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### GetSplashLogo

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashLogo() GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo`

GetSplashLogo returns the SplashLogo field if non-nil, zero value otherwise.

### GetSplashLogoOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashLogoOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo, bool)`

GetSplashLogoOk returns a tuple with the SplashLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashLogo

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashLogo(v GetNetworkWirelessSsidSplashSettings200ResponseSplashLogo)`

SetSplashLogo sets SplashLogo field to given value.

### HasSplashLogo

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashLogo() bool`

HasSplashLogo returns a boolean if a field has been set.

### GetSplashImage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashImage() GetNetworkWirelessSsidSplashSettings200ResponseSplashImage`

GetSplashImage returns the SplashImage field if non-nil, zero value otherwise.

### GetSplashImageOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashImageOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSplashImage, bool)`

GetSplashImageOk returns a tuple with the SplashImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashImage(v GetNetworkWirelessSsidSplashSettings200ResponseSplashImage)`

SetSplashImage sets SplashImage field to given value.

### HasSplashImage

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashImage() bool`

HasSplashImage returns a boolean if a field has been set.

### GetSplashPrepaidFront

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPrepaidFront() GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront`

GetSplashPrepaidFront returns the SplashPrepaidFront field if non-nil, zero value otherwise.

### GetSplashPrepaidFrontOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSplashPrepaidFrontOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront, bool)`

GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPrepaidFront

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSplashPrepaidFront(v GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront)`

SetSplashPrepaidFront sets SplashPrepaidFront field to given value.

### HasSplashPrepaidFront

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSplashPrepaidFront() bool`

HasSplashPrepaidFront returns a boolean if a field has been set.

### GetGuestSponsorship

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetGuestSponsorship() GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship`

GetGuestSponsorship returns the GuestSponsorship field if non-nil, zero value otherwise.

### GetGuestSponsorshipOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetGuestSponsorshipOk() (*GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship, bool)`

GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSponsorship

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetGuestSponsorship(v GetNetworkWirelessSsidSplashSettings200ResponseGuestSponsorship)`

SetGuestSponsorship sets GuestSponsorship field to given value.

### HasGuestSponsorship

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasGuestSponsorship() bool`

HasGuestSponsorship returns a boolean if a field has been set.

### GetBlockAllTrafficBeforeSignOn

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBlockAllTrafficBeforeSignOn() bool`

GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field if non-nil, zero value otherwise.

### GetBlockAllTrafficBeforeSignOnOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool)`

GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllTrafficBeforeSignOn

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetBlockAllTrafficBeforeSignOn(v bool)`

SetBlockAllTrafficBeforeSignOn sets BlockAllTrafficBeforeSignOn field to given value.

### HasBlockAllTrafficBeforeSignOn

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasBlockAllTrafficBeforeSignOn() bool`

HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.

### GetControllerDisconnectionBehavior

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetControllerDisconnectionBehavior() string`

GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field if non-nil, zero value otherwise.

### GetControllerDisconnectionBehaviorOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetControllerDisconnectionBehaviorOk() (*string, bool)`

GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDisconnectionBehavior

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetControllerDisconnectionBehavior(v string)`

SetControllerDisconnectionBehavior sets ControllerDisconnectionBehavior field to given value.

### HasControllerDisconnectionBehavior

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasControllerDisconnectionBehavior() bool`

HasControllerDisconnectionBehavior returns a boolean if a field has been set.

### GetAllowSimultaneousLogins

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetAllowSimultaneousLogins() bool`

GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field if non-nil, zero value otherwise.

### GetAllowSimultaneousLoginsOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetAllowSimultaneousLoginsOk() (*bool, bool)`

GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneousLogins

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetAllowSimultaneousLogins(v bool)`

SetAllowSimultaneousLogins sets AllowSimultaneousLogins field to given value.

### HasAllowSimultaneousLogins

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasAllowSimultaneousLogins() bool`

HasAllowSimultaneousLogins returns a boolean if a field has been set.

### GetBilling

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBilling() GetNetworkWirelessSsidSplashSettings200ResponseBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetBillingOk() (*GetNetworkWirelessSsidSplashSettings200ResponseBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetBilling(v GetNetworkWirelessSsidSplashSettings200ResponseBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetSentryEnrollment

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSentryEnrollment() GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment`

GetSentryEnrollment returns the SentryEnrollment field if non-nil, zero value otherwise.

### GetSentryEnrollmentOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSentryEnrollmentOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment, bool)`

GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryEnrollment

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSentryEnrollment(v GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment)`

SetSentryEnrollment sets SentryEnrollment field to given value.

### HasSentryEnrollment

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSentryEnrollment() bool`

HasSentryEnrollment returns a boolean if a field has been set.

### GetSelfRegistration

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSelfRegistration() GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration`

GetSelfRegistration returns the SelfRegistration field if non-nil, zero value otherwise.

### GetSelfRegistrationOk

`func (o *GetNetworkWirelessSsidSplashSettings200Response) GetSelfRegistrationOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration, bool)`

GetSelfRegistrationOk returns a tuple with the SelfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegistration

`func (o *GetNetworkWirelessSsidSplashSettings200Response) SetSelfRegistration(v GetNetworkWirelessSsidSplashSettings200ResponseSelfRegistration)`

SetSelfRegistration sets SelfRegistration field to given value.

### HasSelfRegistration

`func (o *GetNetworkWirelessSsidSplashSettings200Response) HasSelfRegistration() bool`

HasSelfRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


