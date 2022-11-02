# UpdateNetworkWirelessSsidBonjourForwardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on this SSID. | [optional] 
**Rules** | Pointer to [**[]CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner**](CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner.md) | List of bonjour forwarding rules. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidBonjourForwardingRequest

`func NewUpdateNetworkWirelessSsidBonjourForwardingRequest() *UpdateNetworkWirelessSsidBonjourForwardingRequest`

NewUpdateNetworkWirelessSsidBonjourForwardingRequest instantiates a new UpdateNetworkWirelessSsidBonjourForwardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidBonjourForwardingRequestWithDefaults

`func NewUpdateNetworkWirelessSsidBonjourForwardingRequestWithDefaults() *UpdateNetworkWirelessSsidBonjourForwardingRequest`

NewUpdateNetworkWirelessSsidBonjourForwardingRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidBonjourForwardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetRules() []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetRulesOk() (*[]CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) SetRules(v []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


