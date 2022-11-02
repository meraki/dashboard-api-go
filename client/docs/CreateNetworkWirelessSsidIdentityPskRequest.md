# CreateNetworkWirelessSsidIdentityPskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Identity PSK | 
**Passphrase** | Pointer to **string** | The passphrase for client authentication. If left blank, one will be auto-generated. | [optional] 
**GroupPolicyId** | **string** | The group policy to be applied to clients | 

## Methods

### NewCreateNetworkWirelessSsidIdentityPskRequest

`func NewCreateNetworkWirelessSsidIdentityPskRequest(name string, groupPolicyId string, ) *CreateNetworkWirelessSsidIdentityPskRequest`

NewCreateNetworkWirelessSsidIdentityPskRequest instantiates a new CreateNetworkWirelessSsidIdentityPskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessSsidIdentityPskRequestWithDefaults

`func NewCreateNetworkWirelessSsidIdentityPskRequestWithDefaults() *CreateNetworkWirelessSsidIdentityPskRequest`

NewCreateNetworkWirelessSsidIdentityPskRequestWithDefaults instantiates a new CreateNetworkWirelessSsidIdentityPskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassphrase

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *CreateNetworkWirelessSsidIdentityPskRequest) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


