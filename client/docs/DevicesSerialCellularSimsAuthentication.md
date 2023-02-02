# DevicesSerialCellularSimsAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | APN auth type. | [optional] [default to "none"]
**Username** | Pointer to **string** | APN username, if type is set. | [optional] 
**Password** | Pointer to **string** | APN password, if type is set (if APN password is not supplied, the password is left unchanged). | [optional] 

## Methods

### NewDevicesSerialCellularSimsAuthentication

`func NewDevicesSerialCellularSimsAuthentication() *DevicesSerialCellularSimsAuthentication`

NewDevicesSerialCellularSimsAuthentication instantiates a new DevicesSerialCellularSimsAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialCellularSimsAuthenticationWithDefaults

`func NewDevicesSerialCellularSimsAuthenticationWithDefaults() *DevicesSerialCellularSimsAuthentication`

NewDevicesSerialCellularSimsAuthenticationWithDefaults instantiates a new DevicesSerialCellularSimsAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DevicesSerialCellularSimsAuthentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevicesSerialCellularSimsAuthentication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevicesSerialCellularSimsAuthentication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevicesSerialCellularSimsAuthentication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *DevicesSerialCellularSimsAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DevicesSerialCellularSimsAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DevicesSerialCellularSimsAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DevicesSerialCellularSimsAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *DevicesSerialCellularSimsAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DevicesSerialCellularSimsAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DevicesSerialCellularSimsAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DevicesSerialCellularSimsAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


