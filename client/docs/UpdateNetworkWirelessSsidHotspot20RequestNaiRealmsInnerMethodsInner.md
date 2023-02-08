# UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of method | [optional] 
**AuthenticationTypes** | Pointer to **map[string]interface{}** | The authentication types for the method. These should be formatted as an object with the EAP method category in camelcase as the key and the list of types as the value (nonEapInnerAuthentication: Reserved, PAP, CHAP, MSCHAP, MSCHAPV2; eapInnerAuthentication: EAP-TLS, EAP-SIM, EAP-AKA, EAP-TTLS with MSCHAPv2; credentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, none, Reserved, Vendor Specific; tunneledEapMethodCredentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, Reserved, Anonymous, Vendor Specific) | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner

`func NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner`

NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner instantiates a new UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInnerWithDefaults

`func NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInnerWithDefaults() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner`

NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthenticationTypes

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetAuthenticationTypes() map[string]interface{}`

GetAuthenticationTypes returns the AuthenticationTypes field if non-nil, zero value otherwise.

### GetAuthenticationTypesOk

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) GetAuthenticationTypesOk() (*map[string]interface{}, bool)`

GetAuthenticationTypesOk returns a tuple with the AuthenticationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationTypes

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) SetAuthenticationTypes(v map[string]interface{})`

SetAuthenticationTypes sets AuthenticationTypes field to given value.

### HasAuthenticationTypes

`func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) HasAuthenticationTypes() bool`

HasAuthenticationTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


