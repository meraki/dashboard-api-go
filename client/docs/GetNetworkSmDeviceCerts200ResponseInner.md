# GetNetworkSmDeviceCerts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the certificate. | [optional] 
**NotValidAfter** | Pointer to **string** | The date after which the certificate is no longer valid. | [optional] 
**NotValidBefore** | Pointer to **string** | The date before which the certificate is not valid. | [optional] 
**CertPem** | Pointer to **string** | The PEM of the certificate. | [optional] 
**DeviceId** | Pointer to **string** | The Meraki managed device Id. | [optional] 
**Issuer** | Pointer to **string** | The certificate issuer. | [optional] 
**Subject** | Pointer to **string** | The subject of the certificate. | [optional] 
**Id** | Pointer to **string** | The Meraki Id of the certificate record. | [optional] 

## Methods

### NewGetNetworkSmDeviceCerts200ResponseInner

`func NewGetNetworkSmDeviceCerts200ResponseInner() *GetNetworkSmDeviceCerts200ResponseInner`

NewGetNetworkSmDeviceCerts200ResponseInner instantiates a new GetNetworkSmDeviceCerts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDeviceCerts200ResponseInnerWithDefaults

`func NewGetNetworkSmDeviceCerts200ResponseInnerWithDefaults() *GetNetworkSmDeviceCerts200ResponseInner`

NewGetNetworkSmDeviceCerts200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceCerts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotValidAfter

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetNotValidAfter() string`

GetNotValidAfter returns the NotValidAfter field if non-nil, zero value otherwise.

### GetNotValidAfterOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetNotValidAfterOk() (*string, bool)`

GetNotValidAfterOk returns a tuple with the NotValidAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotValidAfter

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetNotValidAfter(v string)`

SetNotValidAfter sets NotValidAfter field to given value.

### HasNotValidAfter

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasNotValidAfter() bool`

HasNotValidAfter returns a boolean if a field has been set.

### GetNotValidBefore

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetNotValidBefore() string`

GetNotValidBefore returns the NotValidBefore field if non-nil, zero value otherwise.

### GetNotValidBeforeOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetNotValidBeforeOk() (*string, bool)`

GetNotValidBeforeOk returns a tuple with the NotValidBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotValidBefore

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetNotValidBefore(v string)`

SetNotValidBefore sets NotValidBefore field to given value.

### HasNotValidBefore

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasNotValidBefore() bool`

HasNotValidBefore returns a boolean if a field has been set.

### GetCertPem

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetCertPem() string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetCertPemOk() (*string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetCertPem(v string)`

SetCertPem sets CertPem field to given value.

### HasCertPem

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasCertPem() bool`

HasCertPem returns a boolean if a field has been set.

### GetDeviceId

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetIssuer

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetId

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSmDeviceCerts200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSmDeviceCerts200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSmDeviceCerts200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


