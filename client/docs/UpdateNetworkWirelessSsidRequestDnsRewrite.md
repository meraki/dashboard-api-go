# UpdateNetworkWirelessSsidRequestDnsRewrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | User specified DNS servers (up to two servers) | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestDnsRewrite

`func NewUpdateNetworkWirelessSsidRequestDnsRewrite() *UpdateNetworkWirelessSsidRequestDnsRewrite`

NewUpdateNetworkWirelessSsidRequestDnsRewrite instantiates a new UpdateNetworkWirelessSsidRequestDnsRewrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestDnsRewriteWithDefaults

`func NewUpdateNetworkWirelessSsidRequestDnsRewriteWithDefaults() *UpdateNetworkWirelessSsidRequestDnsRewrite`

NewUpdateNetworkWirelessSsidRequestDnsRewriteWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestDnsRewrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


