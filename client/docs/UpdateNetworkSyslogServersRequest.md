# UpdateNetworkSyslogServersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]UpdateNetworkSyslogServersRequestServersInner**](UpdateNetworkSyslogServersRequestServersInner.md) | A list of the syslog servers for this network | 

## Methods

### NewUpdateNetworkSyslogServersRequest

`func NewUpdateNetworkSyslogServersRequest(servers []UpdateNetworkSyslogServersRequestServersInner, ) *UpdateNetworkSyslogServersRequest`

NewUpdateNetworkSyslogServersRequest instantiates a new UpdateNetworkSyslogServersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSyslogServersRequestWithDefaults

`func NewUpdateNetworkSyslogServersRequestWithDefaults() *UpdateNetworkSyslogServersRequest`

NewUpdateNetworkSyslogServersRequestWithDefaults instantiates a new UpdateNetworkSyslogServersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *UpdateNetworkSyslogServersRequest) GetServers() []UpdateNetworkSyslogServersRequestServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UpdateNetworkSyslogServersRequest) GetServersOk() (*[]UpdateNetworkSyslogServersRequestServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UpdateNetworkSyslogServersRequest) SetServers(v []UpdateNetworkSyslogServersRequestServersInner)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


