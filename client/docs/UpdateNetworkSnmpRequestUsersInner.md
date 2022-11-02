# UpdateNetworkSnmpRequestUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username for the SNMP user. Required. | 
**Passphrase** | **string** | The passphrase for the SNMP user. Required. | 

## Methods

### NewUpdateNetworkSnmpRequestUsersInner

`func NewUpdateNetworkSnmpRequestUsersInner(username string, passphrase string, ) *UpdateNetworkSnmpRequestUsersInner`

NewUpdateNetworkSnmpRequestUsersInner instantiates a new UpdateNetworkSnmpRequestUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSnmpRequestUsersInnerWithDefaults

`func NewUpdateNetworkSnmpRequestUsersInnerWithDefaults() *UpdateNetworkSnmpRequestUsersInner`

NewUpdateNetworkSnmpRequestUsersInnerWithDefaults instantiates a new UpdateNetworkSnmpRequestUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateNetworkSnmpRequestUsersInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateNetworkSnmpRequestUsersInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateNetworkSnmpRequestUsersInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassphrase

`func (o *UpdateNetworkSnmpRequestUsersInner) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *UpdateNetworkSnmpRequestUsersInner) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *UpdateNetworkSnmpRequestUsersInner) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


