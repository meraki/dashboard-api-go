# MoveNetworkSmDevices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | The Meraki Ids of the set of devices. | [optional] 
**NewNetwork** | Pointer to **string** | The network to which the devices was moved. | [optional] 

## Methods

### NewMoveNetworkSmDevices200Response

`func NewMoveNetworkSmDevices200Response() *MoveNetworkSmDevices200Response`

NewMoveNetworkSmDevices200Response instantiates a new MoveNetworkSmDevices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveNetworkSmDevices200ResponseWithDefaults

`func NewMoveNetworkSmDevices200ResponseWithDefaults() *MoveNetworkSmDevices200Response`

NewMoveNetworkSmDevices200ResponseWithDefaults instantiates a new MoveNetworkSmDevices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *MoveNetworkSmDevices200Response) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MoveNetworkSmDevices200Response) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MoveNetworkSmDevices200Response) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MoveNetworkSmDevices200Response) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetNewNetwork

`func (o *MoveNetworkSmDevices200Response) GetNewNetwork() string`

GetNewNetwork returns the NewNetwork field if non-nil, zero value otherwise.

### GetNewNetworkOk

`func (o *MoveNetworkSmDevices200Response) GetNewNetworkOk() (*string, bool)`

GetNewNetworkOk returns a tuple with the NewNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNetwork

`func (o *MoveNetworkSmDevices200Response) SetNewNetwork(v string)`

SetNewNetwork sets NewNetwork field to given value.

### HasNewNetwork

`func (o *MoveNetworkSmDevices200Response) HasNewNetwork() bool`

HasNewNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


