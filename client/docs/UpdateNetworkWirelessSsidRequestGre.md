# UpdateNetworkWirelessSsidRequestGre

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concentrator** | Pointer to [**UpdateNetworkWirelessSsidRequestGreConcentrator**](UpdateNetworkWirelessSsidRequestGreConcentrator.md) |  | [optional] 
**Key** | Pointer to **int32** | Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel. | [optional] 

## Methods

### NewUpdateNetworkWirelessSsidRequestGre

`func NewUpdateNetworkWirelessSsidRequestGre() *UpdateNetworkWirelessSsidRequestGre`

NewUpdateNetworkWirelessSsidRequestGre instantiates a new UpdateNetworkWirelessSsidRequestGre object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSsidRequestGreWithDefaults

`func NewUpdateNetworkWirelessSsidRequestGreWithDefaults() *UpdateNetworkWirelessSsidRequestGre`

NewUpdateNetworkWirelessSsidRequestGreWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestGre object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcentrator

`func (o *UpdateNetworkWirelessSsidRequestGre) GetConcentrator() UpdateNetworkWirelessSsidRequestGreConcentrator`

GetConcentrator returns the Concentrator field if non-nil, zero value otherwise.

### GetConcentratorOk

`func (o *UpdateNetworkWirelessSsidRequestGre) GetConcentratorOk() (*UpdateNetworkWirelessSsidRequestGreConcentrator, bool)`

GetConcentratorOk returns a tuple with the Concentrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcentrator

`func (o *UpdateNetworkWirelessSsidRequestGre) SetConcentrator(v UpdateNetworkWirelessSsidRequestGreConcentrator)`

SetConcentrator sets Concentrator field to given value.

### HasConcentrator

`func (o *UpdateNetworkWirelessSsidRequestGre) HasConcentrator() bool`

HasConcentrator returns a boolean if a field has been set.

### GetKey

`func (o *UpdateNetworkWirelessSsidRequestGre) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateNetworkWirelessSsidRequestGre) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateNetworkWirelessSsidRequestGre) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateNetworkWirelessSsidRequestGre) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


