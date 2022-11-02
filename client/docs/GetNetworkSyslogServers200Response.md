# GetNetworkSyslogServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]GetNetworkSyslogServers200ResponseServersInner**](GetNetworkSyslogServers200ResponseServersInner.md) | List of the syslog servers for this network | [optional] 

## Methods

### NewGetNetworkSyslogServers200Response

`func NewGetNetworkSyslogServers200Response() *GetNetworkSyslogServers200Response`

NewGetNetworkSyslogServers200Response instantiates a new GetNetworkSyslogServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSyslogServers200ResponseWithDefaults

`func NewGetNetworkSyslogServers200ResponseWithDefaults() *GetNetworkSyslogServers200Response`

NewGetNetworkSyslogServers200ResponseWithDefaults instantiates a new GetNetworkSyslogServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *GetNetworkSyslogServers200Response) GetServers() []GetNetworkSyslogServers200ResponseServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetNetworkSyslogServers200Response) GetServersOk() (*[]GetNetworkSyslogServers200ResponseServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetNetworkSyslogServers200Response) SetServers(v []GetNetworkSyslogServers200ResponseServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetNetworkSyslogServers200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


