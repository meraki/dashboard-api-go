# NetworksNetworkIdSyslogServersServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The IP address of the syslog server | 
**Port** | **int32** | The port of the syslog server | 
**Roles** | **[]string** | A list of roles for the syslog server. Options (case-insensitive): &#39;Wireless event log&#39;, &#39;Appliance event log&#39;, &#39;Switch event log&#39;, &#39;Air Marshal events&#39;, &#39;Flows&#39;, &#39;URLs&#39;, &#39;IDS alerts&#39;, &#39;Security events&#39; | 

## Methods

### NewNetworksNetworkIdSyslogServersServers

`func NewNetworksNetworkIdSyslogServersServers(host string, port int32, roles []string, ) *NetworksNetworkIdSyslogServersServers`

NewNetworksNetworkIdSyslogServersServers instantiates a new NetworksNetworkIdSyslogServersServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSyslogServersServersWithDefaults

`func NewNetworksNetworkIdSyslogServersServersWithDefaults() *NetworksNetworkIdSyslogServersServers`

NewNetworksNetworkIdSyslogServersServersWithDefaults instantiates a new NetworksNetworkIdSyslogServersServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NetworksNetworkIdSyslogServersServers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdSyslogServersServers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdSyslogServersServers) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *NetworksNetworkIdSyslogServersServers) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdSyslogServersServers) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdSyslogServersServers) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRoles

`func (o *NetworksNetworkIdSyslogServersServers) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NetworksNetworkIdSyslogServersServers) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NetworksNetworkIdSyslogServersServers) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


