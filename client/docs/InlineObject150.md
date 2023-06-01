# InlineObject150

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanningEnabled** | Pointer to **bool** | Whether APs will scan for Bluetooth enabled clients. | [optional] 
**AdvertisingEnabled** | Pointer to **bool** | Whether APs will advertise beacons. | [optional] 
**Uuid** | Pointer to **string** | The UUID to be used in the beacon identifier. | [optional] 
**MajorMinorAssignmentMode** | Pointer to **string** | The way major and minor number should be assigned to nodes in the network. (&#39;Unique&#39;, &#39;Non-unique&#39;) | [optional] 
**Major** | Pointer to **int32** | The major number to be used in the beacon identifier. Only valid in &#39;Non-unique&#39; mode. | [optional] 
**Minor** | Pointer to **int32** | The minor number to be used in the beacon identifier. Only valid in &#39;Non-unique&#39; mode. | [optional] 

## Methods

### NewInlineObject150

`func NewInlineObject150() *InlineObject150`

NewInlineObject150 instantiates a new InlineObject150 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject150WithDefaults

`func NewInlineObject150WithDefaults() *InlineObject150`

NewInlineObject150WithDefaults instantiates a new InlineObject150 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanningEnabled

`func (o *InlineObject150) GetScanningEnabled() bool`

GetScanningEnabled returns the ScanningEnabled field if non-nil, zero value otherwise.

### GetScanningEnabledOk

`func (o *InlineObject150) GetScanningEnabledOk() (*bool, bool)`

GetScanningEnabledOk returns a tuple with the ScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanningEnabled

`func (o *InlineObject150) SetScanningEnabled(v bool)`

SetScanningEnabled sets ScanningEnabled field to given value.

### HasScanningEnabled

`func (o *InlineObject150) HasScanningEnabled() bool`

HasScanningEnabled returns a boolean if a field has been set.

### GetAdvertisingEnabled

`func (o *InlineObject150) GetAdvertisingEnabled() bool`

GetAdvertisingEnabled returns the AdvertisingEnabled field if non-nil, zero value otherwise.

### GetAdvertisingEnabledOk

`func (o *InlineObject150) GetAdvertisingEnabledOk() (*bool, bool)`

GetAdvertisingEnabledOk returns a tuple with the AdvertisingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisingEnabled

`func (o *InlineObject150) SetAdvertisingEnabled(v bool)`

SetAdvertisingEnabled sets AdvertisingEnabled field to given value.

### HasAdvertisingEnabled

`func (o *InlineObject150) HasAdvertisingEnabled() bool`

HasAdvertisingEnabled returns a boolean if a field has been set.

### GetUuid

`func (o *InlineObject150) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InlineObject150) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InlineObject150) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InlineObject150) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMajorMinorAssignmentMode

`func (o *InlineObject150) GetMajorMinorAssignmentMode() string`

GetMajorMinorAssignmentMode returns the MajorMinorAssignmentMode field if non-nil, zero value otherwise.

### GetMajorMinorAssignmentModeOk

`func (o *InlineObject150) GetMajorMinorAssignmentModeOk() (*string, bool)`

GetMajorMinorAssignmentModeOk returns a tuple with the MajorMinorAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorMinorAssignmentMode

`func (o *InlineObject150) SetMajorMinorAssignmentMode(v string)`

SetMajorMinorAssignmentMode sets MajorMinorAssignmentMode field to given value.

### HasMajorMinorAssignmentMode

`func (o *InlineObject150) HasMajorMinorAssignmentMode() bool`

HasMajorMinorAssignmentMode returns a boolean if a field has been set.

### GetMajor

`func (o *InlineObject150) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *InlineObject150) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *InlineObject150) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *InlineObject150) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *InlineObject150) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *InlineObject150) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *InlineObject150) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *InlineObject150) HasMinor() bool`

HasMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


