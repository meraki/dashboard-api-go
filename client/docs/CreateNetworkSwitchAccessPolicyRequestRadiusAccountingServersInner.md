# CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Public IP address of the RADIUS accounting server | 
**Port** | **int32** | UDP port that the RADIUS Accounting server listens on for access requests | 
**Secret** | **string** | RADIUS client shared secret | 

## Methods

### NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner

`func NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner(host string, port int32, secret string, ) *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner`

NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults

`func NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults() *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner`

NewCreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecret

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


