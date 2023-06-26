# InlineObject215

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the combined network | 
**NetworkIds** | **[]string** | A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network | 
**EnrollmentString** | Pointer to **string** | A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by &#39;-network_type&#39;. If left empty, all exisitng enrollment strings will be deleted. | [optional] 

## Methods

### NewInlineObject215

`func NewInlineObject215(name string, networkIds []string, ) *InlineObject215`

NewInlineObject215 instantiates a new InlineObject215 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject215WithDefaults

`func NewInlineObject215WithDefaults() *InlineObject215`

NewInlineObject215WithDefaults instantiates a new InlineObject215 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject215) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject215) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject215) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkIds

`func (o *InlineObject215) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *InlineObject215) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *InlineObject215) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.


### GetEnrollmentString

`func (o *InlineObject215) GetEnrollmentString() string`

GetEnrollmentString returns the EnrollmentString field if non-nil, zero value otherwise.

### GetEnrollmentStringOk

`func (o *InlineObject215) GetEnrollmentStringOk() (*string, bool)`

GetEnrollmentStringOk returns a tuple with the EnrollmentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentString

`func (o *InlineObject215) SetEnrollmentString(v string)`

SetEnrollmentString sets EnrollmentString field to given value.

### HasEnrollmentString

`func (o *InlineObject215) HasEnrollmentString() bool`

HasEnrollmentString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


