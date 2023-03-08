# InlineResponse20097

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The device MAC address. | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesNetwork**](OrganizationsOrganizationIdDevicesAvailabilitiesNetwork.md) |  | [optional] 
**ProductType** | Pointer to **string** | Device product type. | [optional] 
**Serial** | Pointer to **string** | The device serial number. | [optional] 
**Status** | Pointer to **string** | Status of the device. Possible values are: online, alerting, offline, dormant. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device. | [optional] 

## Methods

### NewInlineResponse20097

`func NewInlineResponse20097() *InlineResponse20097`

NewInlineResponse20097 instantiates a new InlineResponse20097 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20097WithDefaults

`func NewInlineResponse20097WithDefaults() *InlineResponse20097`

NewInlineResponse20097WithDefaults instantiates a new InlineResponse20097 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse20097) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse20097) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse20097) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse20097) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20097) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20097) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20097) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20097) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20097) GetNetwork() OrganizationsOrganizationIdDevicesAvailabilitiesNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20097) GetNetworkOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20097) SetNetwork(v OrganizationsOrganizationIdDevicesAvailabilitiesNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20097) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse20097) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse20097) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse20097) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse20097) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse20097) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20097) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20097) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20097) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20097) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20097) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20097) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20097) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20097) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20097) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20097) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20097) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


