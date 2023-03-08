# InlineObject121

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | **int32** | The VLAN of the incoming packet. A null value will match any VLAN. | 
**Protocol** | Pointer to **string** | The protocol of the incoming packet. Can be one of \&quot;ANY\&quot;, \&quot;TCP\&quot; or \&quot;UDP\&quot;. Default value is \&quot;ANY\&quot; | [optional] 
**SrcPort** | Pointer to **int32** | The source port of the incoming packet. Applicable only if protocol is TCP or UDP. | [optional] 
**SrcPortRange** | Pointer to **string** | The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80 | [optional] 
**DstPort** | Pointer to **int32** | The destination port of the incoming packet. Applicable only if protocol is TCP or UDP. | [optional] 
**DstPortRange** | Pointer to **string** | The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80 | [optional] 
**Dscp** | Pointer to **int32** | DSCP tag. Set this to -1 to trust incoming DSCP. Default value is 0 | [optional] 

## Methods

### NewInlineObject121

`func NewInlineObject121(vlan int32, ) *InlineObject121`

NewInlineObject121 instantiates a new InlineObject121 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject121WithDefaults

`func NewInlineObject121WithDefaults() *InlineObject121`

NewInlineObject121WithDefaults instantiates a new InlineObject121 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineObject121) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject121) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject121) SetVlan(v int32)`

SetVlan sets Vlan field to given value.


### GetProtocol

`func (o *InlineObject121) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineObject121) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineObject121) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineObject121) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *InlineObject121) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InlineObject121) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InlineObject121) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InlineObject121) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcPortRange

`func (o *InlineObject121) GetSrcPortRange() string`

GetSrcPortRange returns the SrcPortRange field if non-nil, zero value otherwise.

### GetSrcPortRangeOk

`func (o *InlineObject121) GetSrcPortRangeOk() (*string, bool)`

GetSrcPortRangeOk returns a tuple with the SrcPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPortRange

`func (o *InlineObject121) SetSrcPortRange(v string)`

SetSrcPortRange sets SrcPortRange field to given value.

### HasSrcPortRange

`func (o *InlineObject121) HasSrcPortRange() bool`

HasSrcPortRange returns a boolean if a field has been set.

### GetDstPort

`func (o *InlineObject121) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *InlineObject121) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *InlineObject121) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *InlineObject121) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetDstPortRange

`func (o *InlineObject121) GetDstPortRange() string`

GetDstPortRange returns the DstPortRange field if non-nil, zero value otherwise.

### GetDstPortRangeOk

`func (o *InlineObject121) GetDstPortRangeOk() (*string, bool)`

GetDstPortRangeOk returns a tuple with the DstPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPortRange

`func (o *InlineObject121) SetDstPortRange(v string)`

SetDstPortRange sets DstPortRange field to given value.

### HasDstPortRange

`func (o *InlineObject121) HasDstPortRange() bool`

HasDstPortRange returns a boolean if a field has been set.

### GetDscp

`func (o *InlineObject121) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *InlineObject121) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *InlineObject121) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *InlineObject121) HasDscp() bool`

HasDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


