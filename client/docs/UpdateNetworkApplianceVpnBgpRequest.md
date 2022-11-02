# UpdateNetworkApplianceVpnBgpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Boolean value to enable or disable the BGP configuration. When BGP is enabled, the asNumber (ASN) will be autopopulated with the preconfigured ASN at other Hubs or a default value if there is no ASN configured. | 
**AsNumber** | Pointer to **int32** | An Autonomous System Number (ASN) is required if you are to run BGP and peer with another BGP Speaker outside of the Auto VPN domain. This ASN will be applied to the entire Auto VPN domain. The entire 4-byte ASN range is supported. So, the ASN must be an integer between 1 and 4294967295. When absent, this field is not updated. If no value exists then it defaults to 64512. | [optional] 
**IbgpHoldTimer** | Pointer to **int32** | The IBGP holdtimer in seconds. The IBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240. | [optional] 
**Neighbors** | Pointer to [**[]UpdateNetworkApplianceVpnBgpRequestNeighborsInner**](UpdateNetworkApplianceVpnBgpRequestNeighborsInner.md) | List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated. | [optional] 

## Methods

### NewUpdateNetworkApplianceVpnBgpRequest

`func NewUpdateNetworkApplianceVpnBgpRequest(enabled bool, ) *UpdateNetworkApplianceVpnBgpRequest`

NewUpdateNetworkApplianceVpnBgpRequest instantiates a new UpdateNetworkApplianceVpnBgpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceVpnBgpRequestWithDefaults

`func NewUpdateNetworkApplianceVpnBgpRequestWithDefaults() *UpdateNetworkApplianceVpnBgpRequest`

NewUpdateNetworkApplianceVpnBgpRequestWithDefaults instantiates a new UpdateNetworkApplianceVpnBgpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkApplianceVpnBgpRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsNumber

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *UpdateNetworkApplianceVpnBgpRequest) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *UpdateNetworkApplianceVpnBgpRequest) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetIbgpHoldTimer

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetIbgpHoldTimer() int32`

GetIbgpHoldTimer returns the IbgpHoldTimer field if non-nil, zero value otherwise.

### GetIbgpHoldTimerOk

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetIbgpHoldTimerOk() (*int32, bool)`

GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpHoldTimer

`func (o *UpdateNetworkApplianceVpnBgpRequest) SetIbgpHoldTimer(v int32)`

SetIbgpHoldTimer sets IbgpHoldTimer field to given value.

### HasIbgpHoldTimer

`func (o *UpdateNetworkApplianceVpnBgpRequest) HasIbgpHoldTimer() bool`

HasIbgpHoldTimer returns a boolean if a field has been set.

### GetNeighbors

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetNeighbors() []UpdateNetworkApplianceVpnBgpRequestNeighborsInner`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *UpdateNetworkApplianceVpnBgpRequest) GetNeighborsOk() (*[]UpdateNetworkApplianceVpnBgpRequestNeighborsInner, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *UpdateNetworkApplianceVpnBgpRequest) SetNeighbors(v []UpdateNetworkApplianceVpnBgpRequestNeighborsInner)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *UpdateNetworkApplianceVpnBgpRequest) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


