# UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | MD5 authentication key index. Key index must be between 1 to 255 | [optional] 
**Passphrase** | Pointer to **string** | MD5 authentication passphrase | [optional] 

## Methods

### NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey

`func NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey() *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey`

NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey instantiates a new UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKeyWithDefaults

`func NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKeyWithDefaults() *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey`

NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKeyWithDefaults instantiates a new UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassphrase

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


