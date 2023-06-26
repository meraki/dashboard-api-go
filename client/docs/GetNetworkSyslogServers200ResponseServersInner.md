# GetNetworkSyslogServers200ResponseServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The IP address of the syslog server | [optional] 
**Port** | Pointer to **int32** | The port of the syslog server | [optional] 
**Roles** | Pointer to **[]string** | A list of roles for the syslog server. Options (case-insensitive): &#39;Wireless event log&#39;, &#39;Appliance event log&#39;, &#39;Switch event log&#39;, &#39;Air Marshal events&#39;, &#39;Flows&#39;, &#39;URLs&#39;, &#39;IDS alerts&#39;, &#39;Security events&#39; | [optional] 

## Methods

### NewGetNetworkSyslogServers200ResponseServersInner

`func NewGetNetworkSyslogServers200ResponseServersInner() *GetNetworkSyslogServers200ResponseServersInner`

NewGetNetworkSyslogServers200ResponseServersInner instantiates a new GetNetworkSyslogServers200ResponseServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSyslogServers200ResponseServersInnerWithDefaults

`func NewGetNetworkSyslogServers200ResponseServersInnerWithDefaults() *GetNetworkSyslogServers200ResponseServersInner`

NewGetNetworkSyslogServers200ResponseServersInnerWithDefaults instantiates a new GetNetworkSyslogServers200ResponseServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *GetNetworkSyslogServers200ResponseServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetNetworkSyslogServers200ResponseServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetNetworkSyslogServers200ResponseServersInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetNetworkSyslogServers200ResponseServersInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *GetNetworkSyslogServers200ResponseServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetNetworkSyslogServers200ResponseServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetNetworkSyslogServers200ResponseServersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetNetworkSyslogServers200ResponseServersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRoles

`func (o *GetNetworkSyslogServers200ResponseServersInner) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetNetworkSyslogServers200ResponseServersInner) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetNetworkSyslogServers200ResponseServersInner) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GetNetworkSyslogServers200ResponseServersInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


