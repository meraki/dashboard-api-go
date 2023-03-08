# InlineObject112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**NetworksNetworkIdSwitchDhcpServerPolicyAlerts**](NetworksNetworkIdSwitchDhcpServerPolicyAlerts.md) |  | [optional] 
**DefaultPolicy** | Pointer to **string** | &#39;allow&#39; or &#39;block&#39; new DHCP servers. Default value is &#39;allow&#39;. | [optional] 
**AllowedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries. | [optional] 
**BlockedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries. | [optional] 
**ArpInspection** | Pointer to [**NetworksNetworkIdSwitchDhcpServerPolicyArpInspection**](NetworksNetworkIdSwitchDhcpServerPolicyArpInspection.md) |  | [optional] 

## Methods

### NewInlineObject112

`func NewInlineObject112() *InlineObject112`

NewInlineObject112 instantiates a new InlineObject112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject112WithDefaults

`func NewInlineObject112WithDefaults() *InlineObject112`

NewInlineObject112WithDefaults instantiates a new InlineObject112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *InlineObject112) GetAlerts() NetworksNetworkIdSwitchDhcpServerPolicyAlerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject112) GetAlertsOk() (*NetworksNetworkIdSwitchDhcpServerPolicyAlerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject112) SetAlerts(v NetworksNetworkIdSwitchDhcpServerPolicyAlerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject112) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *InlineObject112) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *InlineObject112) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *InlineObject112) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *InlineObject112) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetAllowedServers

`func (o *InlineObject112) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *InlineObject112) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *InlineObject112) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *InlineObject112) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetBlockedServers

`func (o *InlineObject112) GetBlockedServers() []string`

GetBlockedServers returns the BlockedServers field if non-nil, zero value otherwise.

### GetBlockedServersOk

`func (o *InlineObject112) GetBlockedServersOk() (*[]string, bool)`

GetBlockedServersOk returns a tuple with the BlockedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServers

`func (o *InlineObject112) SetBlockedServers(v []string)`

SetBlockedServers sets BlockedServers field to given value.

### HasBlockedServers

`func (o *InlineObject112) HasBlockedServers() bool`

HasBlockedServers returns a boolean if a field has been set.

### GetArpInspection

`func (o *InlineObject112) GetArpInspection() NetworksNetworkIdSwitchDhcpServerPolicyArpInspection`

GetArpInspection returns the ArpInspection field if non-nil, zero value otherwise.

### GetArpInspectionOk

`func (o *InlineObject112) GetArpInspectionOk() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspection, bool)`

GetArpInspectionOk returns a tuple with the ArpInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpInspection

`func (o *InlineObject112) SetArpInspection(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspection)`

SetArpInspection sets ArpInspection field to given value.

### HasArpInspection

`func (o *InlineObject112) HasArpInspection() bool`

HasArpInspection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


