# InlineResponse20065

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Switch serial. | [optional] 
**Name** | Pointer to **string** | Switch name. | [optional] 
**Url** | Pointer to **string** | Url link to switch. | [optional] 
**SupportsInspection** | Pointer to **bool** | Whether this switch supports Dynamic ARP Inspection. | [optional] 
**HasTrustedPort** | Pointer to **bool** | Whether this switch has a trusted DAI port. Always false if supportsInspection is false. | [optional] 

## Methods

### NewInlineResponse20065

`func NewInlineResponse20065() *InlineResponse20065`

NewInlineResponse20065 instantiates a new InlineResponse20065 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20065WithDefaults

`func NewInlineResponse20065WithDefaults() *InlineResponse20065`

NewInlineResponse20065WithDefaults instantiates a new InlineResponse20065 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20065) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20065) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20065) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20065) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20065) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20065) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20065) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20065) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20065) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20065) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20065) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20065) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSupportsInspection

`func (o *InlineResponse20065) GetSupportsInspection() bool`

GetSupportsInspection returns the SupportsInspection field if non-nil, zero value otherwise.

### GetSupportsInspectionOk

`func (o *InlineResponse20065) GetSupportsInspectionOk() (*bool, bool)`

GetSupportsInspectionOk returns a tuple with the SupportsInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsInspection

`func (o *InlineResponse20065) SetSupportsInspection(v bool)`

SetSupportsInspection sets SupportsInspection field to given value.

### HasSupportsInspection

`func (o *InlineResponse20065) HasSupportsInspection() bool`

HasSupportsInspection returns a boolean if a field has been set.

### GetHasTrustedPort

`func (o *InlineResponse20065) GetHasTrustedPort() bool`

GetHasTrustedPort returns the HasTrustedPort field if non-nil, zero value otherwise.

### GetHasTrustedPortOk

`func (o *InlineResponse20065) GetHasTrustedPortOk() (*bool, bool)`

GetHasTrustedPortOk returns a tuple with the HasTrustedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTrustedPort

`func (o *InlineResponse20065) SetHasTrustedPort(v bool)`

SetHasTrustedPort sets HasTrustedPort field to given value.

### HasHasTrustedPort

`func (o *InlineResponse20065) HasHasTrustedPort() bool`

HasHasTrustedPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


