# InlineResponse20017Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse20017ValueSource**](InlineResponse20017ValueSource.md) |  | 
**Destination** | [**InlineResponse20017ValueDestination**](InlineResponse20017ValueDestination.md) |  | 

## Methods

### NewInlineResponse20017Value

`func NewInlineResponse20017Value(source InlineResponse20017ValueSource, destination InlineResponse20017ValueDestination, ) *InlineResponse20017Value`

NewInlineResponse20017Value instantiates a new InlineResponse20017Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20017ValueWithDefaults

`func NewInlineResponse20017ValueWithDefaults() *InlineResponse20017Value`

NewInlineResponse20017ValueWithDefaults instantiates a new InlineResponse20017Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20017Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20017Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20017Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20017Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20017Value) GetSource() InlineResponse20017ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20017Value) GetSourceOk() (*InlineResponse20017ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20017Value) SetSource(v InlineResponse20017ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse20017Value) GetDestination() InlineResponse20017ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20017Value) GetDestinationOk() (*InlineResponse20017ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20017Value) SetDestination(v InlineResponse20017ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


