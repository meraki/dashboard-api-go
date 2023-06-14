# UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | IP address of your Active Directory server. | 
**Port** | Pointer to **int32** | (Optional) UDP port the Active Directory server listens on. By default, uses port 3268. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner

`func NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner(host string, ) *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner`

NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInnerWithDefaults

`func NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInnerWithDefaults() *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner`

NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


