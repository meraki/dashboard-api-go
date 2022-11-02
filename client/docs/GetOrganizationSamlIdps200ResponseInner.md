# GetOrganizationSamlIdps200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpId** | Pointer to **string** | ID associated with the SAML Identity Provider (IdP) | [optional] 
**ConsumerUrl** | Pointer to **string** | URL that is consuming SAML Identity Provider (IdP) | [optional] 
**X509certSha1Fingerprint** | Pointer to **string** | Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation. | [optional] 
**SloLogoutUrl** | Pointer to **string** | Dashboard will redirect users to this URL when they sign out. | [optional] 

## Methods

### NewGetOrganizationSamlIdps200ResponseInner

`func NewGetOrganizationSamlIdps200ResponseInner() *GetOrganizationSamlIdps200ResponseInner`

NewGetOrganizationSamlIdps200ResponseInner instantiates a new GetOrganizationSamlIdps200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSamlIdps200ResponseInnerWithDefaults

`func NewGetOrganizationSamlIdps200ResponseInnerWithDefaults() *GetOrganizationSamlIdps200ResponseInner`

NewGetOrganizationSamlIdps200ResponseInnerWithDefaults instantiates a new GetOrganizationSamlIdps200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpId

`func (o *GetOrganizationSamlIdps200ResponseInner) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *GetOrganizationSamlIdps200ResponseInner) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *GetOrganizationSamlIdps200ResponseInner) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *GetOrganizationSamlIdps200ResponseInner) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### GetConsumerUrl

`func (o *GetOrganizationSamlIdps200ResponseInner) GetConsumerUrl() string`

GetConsumerUrl returns the ConsumerUrl field if non-nil, zero value otherwise.

### GetConsumerUrlOk

`func (o *GetOrganizationSamlIdps200ResponseInner) GetConsumerUrlOk() (*string, bool)`

GetConsumerUrlOk returns a tuple with the ConsumerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUrl

`func (o *GetOrganizationSamlIdps200ResponseInner) SetConsumerUrl(v string)`

SetConsumerUrl sets ConsumerUrl field to given value.

### HasConsumerUrl

`func (o *GetOrganizationSamlIdps200ResponseInner) HasConsumerUrl() bool`

HasConsumerUrl returns a boolean if a field has been set.

### GetX509certSha1Fingerprint

`func (o *GetOrganizationSamlIdps200ResponseInner) GetX509certSha1Fingerprint() string`

GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field if non-nil, zero value otherwise.

### GetX509certSha1FingerprintOk

`func (o *GetOrganizationSamlIdps200ResponseInner) GetX509certSha1FingerprintOk() (*string, bool)`

GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509certSha1Fingerprint

`func (o *GetOrganizationSamlIdps200ResponseInner) SetX509certSha1Fingerprint(v string)`

SetX509certSha1Fingerprint sets X509certSha1Fingerprint field to given value.

### HasX509certSha1Fingerprint

`func (o *GetOrganizationSamlIdps200ResponseInner) HasX509certSha1Fingerprint() bool`

HasX509certSha1Fingerprint returns a boolean if a field has been set.

### GetSloLogoutUrl

`func (o *GetOrganizationSamlIdps200ResponseInner) GetSloLogoutUrl() string`

GetSloLogoutUrl returns the SloLogoutUrl field if non-nil, zero value otherwise.

### GetSloLogoutUrlOk

`func (o *GetOrganizationSamlIdps200ResponseInner) GetSloLogoutUrlOk() (*string, bool)`

GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloLogoutUrl

`func (o *GetOrganizationSamlIdps200ResponseInner) SetSloLogoutUrl(v string)`

SetSloLogoutUrl sets SloLogoutUrl field to given value.

### HasSloLogoutUrl

`func (o *GetOrganizationSamlIdps200ResponseInner) HasSloLogoutUrl() bool`

HasSloLogoutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


