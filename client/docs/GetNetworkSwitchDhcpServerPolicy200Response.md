# GetNetworkSwitchDhcpServerPolicy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**GetNetworkSwitchDhcpServerPolicy200ResponseAlerts**](GetNetworkSwitchDhcpServerPolicy200ResponseAlerts.md) |  | [optional] 
**DefaultPolicy** | Pointer to **string** | &#39;allow&#39; or &#39;block&#39; new DHCP servers. Default value is &#39;allow&#39;. | [optional] 
**BlockedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries. | [optional] 
**AllowedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries. | [optional] 
**ArpInspection** | Pointer to [**GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection**](GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection.md) |  | [optional] 

## Methods

### NewGetNetworkSwitchDhcpServerPolicy200Response

`func NewGetNetworkSwitchDhcpServerPolicy200Response() *GetNetworkSwitchDhcpServerPolicy200Response`

NewGetNetworkSwitchDhcpServerPolicy200Response instantiates a new GetNetworkSwitchDhcpServerPolicy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchDhcpServerPolicy200ResponseWithDefaults

`func NewGetNetworkSwitchDhcpServerPolicy200ResponseWithDefaults() *GetNetworkSwitchDhcpServerPolicy200Response`

NewGetNetworkSwitchDhcpServerPolicy200ResponseWithDefaults instantiates a new GetNetworkSwitchDhcpServerPolicy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAlerts() GetNetworkSwitchDhcpServerPolicy200ResponseAlerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAlertsOk() (*GetNetworkSwitchDhcpServerPolicy200ResponseAlerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetAlerts(v GetNetworkSwitchDhcpServerPolicy200ResponseAlerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetBlockedServers

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetBlockedServers() []string`

GetBlockedServers returns the BlockedServers field if non-nil, zero value otherwise.

### GetBlockedServersOk

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetBlockedServersOk() (*[]string, bool)`

GetBlockedServersOk returns a tuple with the BlockedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServers

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetBlockedServers(v []string)`

SetBlockedServers sets BlockedServers field to given value.

### HasBlockedServers

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasBlockedServers() bool`

HasBlockedServers returns a boolean if a field has been set.

### GetAllowedServers

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetArpInspection

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetArpInspection() GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection`

GetArpInspection returns the ArpInspection field if non-nil, zero value otherwise.

### GetArpInspectionOk

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetArpInspectionOk() (*GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection, bool)`

GetArpInspectionOk returns a tuple with the ArpInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpInspection

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetArpInspection(v GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection)`

SetArpInspection sets ArpInspection field to given value.

### HasArpInspection

`func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasArpInspection() bool`

HasArpInspection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


