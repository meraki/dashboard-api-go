# CreateOrganizationSamlIdpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X509certSha1Fingerprint** | **string** | Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation. | 
**SloLogoutUrl** | Pointer to **string** | Dashboard will redirect users to this URL when they sign out. | [optional] 

## Methods

### NewCreateOrganizationSamlIdpRequest

`func NewCreateOrganizationSamlIdpRequest(x509certSha1Fingerprint string, ) *CreateOrganizationSamlIdpRequest`

NewCreateOrganizationSamlIdpRequest instantiates a new CreateOrganizationSamlIdpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationSamlIdpRequestWithDefaults

`func NewCreateOrganizationSamlIdpRequestWithDefaults() *CreateOrganizationSamlIdpRequest`

NewCreateOrganizationSamlIdpRequestWithDefaults instantiates a new CreateOrganizationSamlIdpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX509certSha1Fingerprint

`func (o *CreateOrganizationSamlIdpRequest) GetX509certSha1Fingerprint() string`

GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field if non-nil, zero value otherwise.

### GetX509certSha1FingerprintOk

`func (o *CreateOrganizationSamlIdpRequest) GetX509certSha1FingerprintOk() (*string, bool)`

GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509certSha1Fingerprint

`func (o *CreateOrganizationSamlIdpRequest) SetX509certSha1Fingerprint(v string)`

SetX509certSha1Fingerprint sets X509certSha1Fingerprint field to given value.


### GetSloLogoutUrl

`func (o *CreateOrganizationSamlIdpRequest) GetSloLogoutUrl() string`

GetSloLogoutUrl returns the SloLogoutUrl field if non-nil, zero value otherwise.

### GetSloLogoutUrlOk

`func (o *CreateOrganizationSamlIdpRequest) GetSloLogoutUrlOk() (*string, bool)`

GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloLogoutUrl

`func (o *CreateOrganizationSamlIdpRequest) SetSloLogoutUrl(v string)`

SetSloLogoutUrl sets SloLogoutUrl field to given value.

### HasSloLogoutUrl

`func (o *CreateOrganizationSamlIdpRequest) HasSloLogoutUrl() bool`

HasSloLogoutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


