# InlineResponse20092ApBandSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39; or &#39;5ghz&#39;. Defaults to dual. | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band. Can be either true or false. Defaults to true. | [optional] 

## Methods

### NewInlineResponse20092ApBandSettings

`func NewInlineResponse20092ApBandSettings() *InlineResponse20092ApBandSettings`

NewInlineResponse20092ApBandSettings instantiates a new InlineResponse20092ApBandSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20092ApBandSettingsWithDefaults

`func NewInlineResponse20092ApBandSettingsWithDefaults() *InlineResponse20092ApBandSettings`

NewInlineResponse20092ApBandSettingsWithDefaults instantiates a new InlineResponse20092ApBandSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandOperationMode

`func (o *InlineResponse20092ApBandSettings) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *InlineResponse20092ApBandSettings) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *InlineResponse20092ApBandSettings) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *InlineResponse20092ApBandSettings) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *InlineResponse20092ApBandSettings) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *InlineResponse20092ApBandSettings) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *InlineResponse20092ApBandSettings) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *InlineResponse20092ApBandSettings) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


