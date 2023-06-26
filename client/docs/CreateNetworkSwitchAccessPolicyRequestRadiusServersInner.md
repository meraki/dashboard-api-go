# CreateNetworkSwitchAccessPolicyRequestRadiusServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Public IP address of the RADIUS server | 
**Port** | **int32** | UDP port that the RADIUS server listens on for access requests | 
**Secret** | **string** | RADIUS client shared secret | 

## Methods

### NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInner

`func NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInner(host string, port int32, secret string, ) *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner`

NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInner instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInnerWithDefaults

`func NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInnerWithDefaults() *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner`

NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInnerWithDefaults instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecret

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


