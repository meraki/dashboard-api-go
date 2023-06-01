# InlineObject222

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V2cEnabled** | Pointer to **bool** | Boolean indicating whether SNMP version 2c is enabled for the organization. | [optional] 
**V3Enabled** | Pointer to **bool** | Boolean indicating whether SNMP version 3 is enabled for the organization. | [optional] 
**V3AuthMode** | Pointer to **string** | The SNMP version 3 authentication mode. Can be either &#39;MD5&#39; or &#39;SHA&#39;. | [optional] 
**V3AuthPass** | Pointer to **string** | The SNMP version 3 authentication password. Must be at least 8 characters if specified. | [optional] 
**V3PrivMode** | Pointer to **string** | The SNMP version 3 privacy mode. Can be either &#39;DES&#39; or &#39;AES128&#39;. | [optional] 
**V3PrivPass** | Pointer to **string** | The SNMP version 3 privacy password. Must be at least 8 characters if specified. | [optional] 
**PeerIps** | Pointer to **[]string** | The list of IPv4 addresses that are allowed to access the SNMP server. | [optional] 

## Methods

### NewInlineObject222

`func NewInlineObject222() *InlineObject222`

NewInlineObject222 instantiates a new InlineObject222 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject222WithDefaults

`func NewInlineObject222WithDefaults() *InlineObject222`

NewInlineObject222WithDefaults instantiates a new InlineObject222 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV2cEnabled

`func (o *InlineObject222) GetV2cEnabled() bool`

GetV2cEnabled returns the V2cEnabled field if non-nil, zero value otherwise.

### GetV2cEnabledOk

`func (o *InlineObject222) GetV2cEnabledOk() (*bool, bool)`

GetV2cEnabledOk returns a tuple with the V2cEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2cEnabled

`func (o *InlineObject222) SetV2cEnabled(v bool)`

SetV2cEnabled sets V2cEnabled field to given value.

### HasV2cEnabled

`func (o *InlineObject222) HasV2cEnabled() bool`

HasV2cEnabled returns a boolean if a field has been set.

### GetV3Enabled

`func (o *InlineObject222) GetV3Enabled() bool`

GetV3Enabled returns the V3Enabled field if non-nil, zero value otherwise.

### GetV3EnabledOk

`func (o *InlineObject222) GetV3EnabledOk() (*bool, bool)`

GetV3EnabledOk returns a tuple with the V3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Enabled

`func (o *InlineObject222) SetV3Enabled(v bool)`

SetV3Enabled sets V3Enabled field to given value.

### HasV3Enabled

`func (o *InlineObject222) HasV3Enabled() bool`

HasV3Enabled returns a boolean if a field has been set.

### GetV3AuthMode

`func (o *InlineObject222) GetV3AuthMode() string`

GetV3AuthMode returns the V3AuthMode field if non-nil, zero value otherwise.

### GetV3AuthModeOk

`func (o *InlineObject222) GetV3AuthModeOk() (*string, bool)`

GetV3AuthModeOk returns a tuple with the V3AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3AuthMode

`func (o *InlineObject222) SetV3AuthMode(v string)`

SetV3AuthMode sets V3AuthMode field to given value.

### HasV3AuthMode

`func (o *InlineObject222) HasV3AuthMode() bool`

HasV3AuthMode returns a boolean if a field has been set.

### GetV3AuthPass

`func (o *InlineObject222) GetV3AuthPass() string`

GetV3AuthPass returns the V3AuthPass field if non-nil, zero value otherwise.

### GetV3AuthPassOk

`func (o *InlineObject222) GetV3AuthPassOk() (*string, bool)`

GetV3AuthPassOk returns a tuple with the V3AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3AuthPass

`func (o *InlineObject222) SetV3AuthPass(v string)`

SetV3AuthPass sets V3AuthPass field to given value.

### HasV3AuthPass

`func (o *InlineObject222) HasV3AuthPass() bool`

HasV3AuthPass returns a boolean if a field has been set.

### GetV3PrivMode

`func (o *InlineObject222) GetV3PrivMode() string`

GetV3PrivMode returns the V3PrivMode field if non-nil, zero value otherwise.

### GetV3PrivModeOk

`func (o *InlineObject222) GetV3PrivModeOk() (*string, bool)`

GetV3PrivModeOk returns a tuple with the V3PrivMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3PrivMode

`func (o *InlineObject222) SetV3PrivMode(v string)`

SetV3PrivMode sets V3PrivMode field to given value.

### HasV3PrivMode

`func (o *InlineObject222) HasV3PrivMode() bool`

HasV3PrivMode returns a boolean if a field has been set.

### GetV3PrivPass

`func (o *InlineObject222) GetV3PrivPass() string`

GetV3PrivPass returns the V3PrivPass field if non-nil, zero value otherwise.

### GetV3PrivPassOk

`func (o *InlineObject222) GetV3PrivPassOk() (*string, bool)`

GetV3PrivPassOk returns a tuple with the V3PrivPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3PrivPass

`func (o *InlineObject222) SetV3PrivPass(v string)`

SetV3PrivPass sets V3PrivPass field to given value.

### HasV3PrivPass

`func (o *InlineObject222) HasV3PrivPass() bool`

HasV3PrivPass returns a boolean if a field has been set.

### GetPeerIps

`func (o *InlineObject222) GetPeerIps() []string`

GetPeerIps returns the PeerIps field if non-nil, zero value otherwise.

### GetPeerIpsOk

`func (o *InlineObject222) GetPeerIpsOk() (*[]string, bool)`

GetPeerIpsOk returns a tuple with the PeerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIps

`func (o *InlineObject222) SetPeerIps(v []string)`

SetPeerIps sets PeerIps field to given value.

### HasPeerIps

`func (o *InlineObject222) HasPeerIps() bool`

HasPeerIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


