# InlineResponse20019Value1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of &#39;applicationCategory&#39; or &#39;application&#39; type traffic filter | [optional] 
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | Pointer to [**InlineResponse20019Value1Source**](InlineResponse20019Value1Source.md) |  | [optional] 
**Destination** | Pointer to [**InlineResponse20019Value1Destination**](InlineResponse20019Value1Destination.md) |  | [optional] 

## Methods

### NewInlineResponse20019Value1

`func NewInlineResponse20019Value1() *InlineResponse20019Value1`

NewInlineResponse20019Value1 instantiates a new InlineResponse20019Value1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019Value1WithDefaults

`func NewInlineResponse20019Value1WithDefaults() *InlineResponse20019Value1`

NewInlineResponse20019Value1WithDefaults instantiates a new InlineResponse20019Value1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20019Value1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20019Value1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20019Value1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20019Value1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20019Value1) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20019Value1) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20019Value1) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20019Value1) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20019Value1) GetSource() InlineResponse20019Value1Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20019Value1) GetSourceOk() (*InlineResponse20019Value1Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20019Value1) SetSource(v InlineResponse20019Value1Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *InlineResponse20019Value1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *InlineResponse20019Value1) GetDestination() InlineResponse20019Value1Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20019Value1) GetDestinationOk() (*InlineResponse20019Value1Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20019Value1) SetDestination(v InlineResponse20019Value1Destination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineResponse20019Value1) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


