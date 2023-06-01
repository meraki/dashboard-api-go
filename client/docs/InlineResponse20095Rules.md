# InlineResponse20095Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule | [optional] 
**Policy** | Pointer to **string** | &#39;Deny&#39; traffic specified by this rule | [optional] 
**Protocol** | Pointer to **string** | The type of protocol | [optional] 
**SrcPort** | Pointer to **string** | List of source ports | [optional] 
**SrcCidr** | Pointer to **string** | List of source IP addresses | [optional] 
**DestPort** | Pointer to **string** | List of destination ports | [optional] 
**DestCidr** | Pointer to **string** | List of destination IP addresses | [optional] 
**SyslogEnabled** | Pointer to **bool** | Flag indicating whether the rule should be logged to syslog | [optional] 

## Methods

### NewInlineResponse20095Rules

`func NewInlineResponse20095Rules() *InlineResponse20095Rules`

NewInlineResponse20095Rules instantiates a new InlineResponse20095Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20095RulesWithDefaults

`func NewInlineResponse20095RulesWithDefaults() *InlineResponse20095Rules`

NewInlineResponse20095RulesWithDefaults instantiates a new InlineResponse20095Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineResponse20095Rules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse20095Rules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse20095Rules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse20095Rules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse20095Rules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse20095Rules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse20095Rules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse20095Rules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20095Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20095Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20095Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20095Rules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *InlineResponse20095Rules) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InlineResponse20095Rules) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InlineResponse20095Rules) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InlineResponse20095Rules) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcCidr

`func (o *InlineResponse20095Rules) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *InlineResponse20095Rules) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *InlineResponse20095Rules) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.

### HasSrcCidr

`func (o *InlineResponse20095Rules) HasSrcCidr() bool`

HasSrcCidr returns a boolean if a field has been set.

### GetDestPort

`func (o *InlineResponse20095Rules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *InlineResponse20095Rules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *InlineResponse20095Rules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *InlineResponse20095Rules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *InlineResponse20095Rules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *InlineResponse20095Rules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *InlineResponse20095Rules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.

### HasDestCidr

`func (o *InlineResponse20095Rules) HasDestCidr() bool`

HasDestCidr returns a boolean if a field has been set.

### GetSyslogEnabled

`func (o *InlineResponse20095Rules) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *InlineResponse20095Rules) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *InlineResponse20095Rules) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *InlineResponse20095Rules) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


