# UpdateNetworkApplianceSsidRequestRadiusServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The IP address of your RADIUS server. | [optional] 
**Port** | Pointer to **int32** | The UDP port your RADIUS servers listens on for Access-requests. | [optional] 
**Secret** | Pointer to **string** | The RADIUS client shared secret. | [optional] 

## Methods

### NewUpdateNetworkApplianceSsidRequestRadiusServersInner

`func NewUpdateNetworkApplianceSsidRequestRadiusServersInner() *UpdateNetworkApplianceSsidRequestRadiusServersInner`

NewUpdateNetworkApplianceSsidRequestRadiusServersInner instantiates a new UpdateNetworkApplianceSsidRequestRadiusServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceSsidRequestRadiusServersInnerWithDefaults

`func NewUpdateNetworkApplianceSsidRequestRadiusServersInnerWithDefaults() *UpdateNetworkApplianceSsidRequestRadiusServersInner`

NewUpdateNetworkApplianceSsidRequestRadiusServersInnerWithDefaults instantiates a new UpdateNetworkApplianceSsidRequestRadiusServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateNetworkApplianceSsidRequestRadiusServersInner) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


