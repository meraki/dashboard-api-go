# UpdateNetworkSwitchDhcpServerPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**UpdateNetworkSwitchDhcpServerPolicyRequestAlerts**](UpdateNetworkSwitchDhcpServerPolicyRequestAlerts.md) |  | [optional] 
**DefaultPolicy** | Pointer to **string** | &#39;allow&#39; or &#39;block&#39; new DHCP servers. Default value is &#39;allow&#39;. | [optional] 
**AllowedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries. | [optional] 
**BlockedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries. | [optional] 
**ArpInspection** | Pointer to [**UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection**](UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection.md) |  | [optional] 

## Methods

### NewUpdateNetworkSwitchDhcpServerPolicyRequest

`func NewUpdateNetworkSwitchDhcpServerPolicyRequest() *UpdateNetworkSwitchDhcpServerPolicyRequest`

NewUpdateNetworkSwitchDhcpServerPolicyRequest instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchDhcpServerPolicyRequestWithDefaults

`func NewUpdateNetworkSwitchDhcpServerPolicyRequestWithDefaults() *UpdateNetworkSwitchDhcpServerPolicyRequest`

NewUpdateNetworkSwitchDhcpServerPolicyRequestWithDefaults instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAlerts() UpdateNetworkSwitchDhcpServerPolicyRequestAlerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAlertsOk() (*UpdateNetworkSwitchDhcpServerPolicyRequestAlerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetAlerts(v UpdateNetworkSwitchDhcpServerPolicyRequestAlerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetAllowedServers

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetBlockedServers

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetBlockedServers() []string`

GetBlockedServers returns the BlockedServers field if non-nil, zero value otherwise.

### GetBlockedServersOk

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetBlockedServersOk() (*[]string, bool)`

GetBlockedServersOk returns a tuple with the BlockedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServers

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetBlockedServers(v []string)`

SetBlockedServers sets BlockedServers field to given value.

### HasBlockedServers

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasBlockedServers() bool`

HasBlockedServers returns a boolean if a field has been set.

### GetArpInspection

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetArpInspection() UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection`

GetArpInspection returns the ArpInspection field if non-nil, zero value otherwise.

### GetArpInspectionOk

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetArpInspectionOk() (*UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection, bool)`

GetArpInspectionOk returns a tuple with the ArpInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpInspection

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetArpInspection(v UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection)`

SetArpInspection sets ArpInspection field to given value.

### HasArpInspection

`func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasArpInspection() bool`

HasArpInspection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


