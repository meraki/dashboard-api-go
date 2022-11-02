# UpdateNetworkWirelessSsidEapOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | General EAP timeout in seconds. | [optional] 
**Identity** | Pointer to [**UpdateNetworkWirelessSsidEapOverrideRequestIdentity**](UpdateNetworkWirelessSsidEapOverrideRequestIdentity.md) |  | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of general EAP retries. | [optional] 
**EapolKey** | Pointer to [**UpdateNetworkWirelessSsidEapOverrideRequestEapolKey**](UpdateNetworkWirelessSsidEapOverrideRequestEapolKey.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidEapOverrideRequest

`func NewUpdateNetworkWirelessSsidEapOverrideRequest() *UpdateNetworkWirelessSsidEapOverrideRequest`

NewUpdateNetworkWirelessSsidEapOverrideRequest instantiates a new UpdateNetworkWirelessSsidEapOverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidEapOverrideRequestWithDefaults

`func NewUpdateNetworkWirelessSsidEapOverrideRequestWithDefaults() *UpdateNetworkWirelessSsidEapOverrideRequest`

NewUpdateNetworkWirelessSsidEapOverrideRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidEapOverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetIdentity

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetIdentity() UpdateNetworkWirelessSsidEapOverrideRequestIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetIdentityOk() (*UpdateNetworkWirelessSsidEapOverrideRequestIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetIdentity(v UpdateNetworkWirelessSsidEapOverrideRequestIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaxRetries

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetEapolKey

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetEapolKey() UpdateNetworkWirelessSsidEapOverrideRequestEapolKey`

GetEapolKey returns the EapolKey field if non-nil, zero value otherwise.

### GetEapolKeyOk

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetEapolKeyOk() (*UpdateNetworkWirelessSsidEapOverrideRequestEapolKey, bool)`

GetEapolKeyOk returns a tuple with the EapolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapolKey

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetEapolKey(v UpdateNetworkWirelessSsidEapOverrideRequestEapolKey)`

SetEapolKey sets EapolKey field to given value.

### HasEapolKey

`func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasEapolKey() bool`

HasEapolKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


