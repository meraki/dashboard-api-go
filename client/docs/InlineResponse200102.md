# InlineResponse200102

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Model** | Pointer to **string** | Model type of the device | [optional] 
**NetworkId** | Pointer to **string** | Network Id of the device | [optional] 
**OrderNumber** | Pointer to **string** | Order number of the device | [optional] 
**ClaimedAt** | Pointer to **time.Time** | Claimed time of the device | [optional] 
**LicenseExpirationDate** | Pointer to **time.Time** | License expiration date of the device | [optional] 
**Tags** | Pointer to **[]string** | Device tags | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 

## Methods

### NewInlineResponse200102

`func NewInlineResponse200102() *InlineResponse200102`

NewInlineResponse200102 instantiates a new InlineResponse200102 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102WithDefaults

`func NewInlineResponse200102WithDefaults() *InlineResponse200102`

NewInlineResponse200102WithDefaults instantiates a new InlineResponse200102 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse200102) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200102) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200102) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200102) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200102) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200102) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200102) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200102) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200102) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200102) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200102) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200102) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200102) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200102) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200102) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200102) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200102) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200102) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200102) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200102) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *InlineResponse200102) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *InlineResponse200102) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *InlineResponse200102) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *InlineResponse200102) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetClaimedAt

`func (o *InlineResponse200102) GetClaimedAt() time.Time`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *InlineResponse200102) GetClaimedAtOk() (*time.Time, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *InlineResponse200102) SetClaimedAt(v time.Time)`

SetClaimedAt sets ClaimedAt field to given value.

### HasClaimedAt

`func (o *InlineResponse200102) HasClaimedAt() bool`

HasClaimedAt returns a boolean if a field has been set.

### GetLicenseExpirationDate

`func (o *InlineResponse200102) GetLicenseExpirationDate() time.Time`

GetLicenseExpirationDate returns the LicenseExpirationDate field if non-nil, zero value otherwise.

### GetLicenseExpirationDateOk

`func (o *InlineResponse200102) GetLicenseExpirationDateOk() (*time.Time, bool)`

GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationDate

`func (o *InlineResponse200102) SetLicenseExpirationDate(v time.Time)`

SetLicenseExpirationDate sets LicenseExpirationDate field to given value.

### HasLicenseExpirationDate

`func (o *InlineResponse200102) HasLicenseExpirationDate() bool`

HasLicenseExpirationDate returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200102) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200102) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200102) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200102) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200102) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200102) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200102) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200102) HasProductType() bool`

HasProductType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


