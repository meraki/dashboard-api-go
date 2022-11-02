# GetNetworkSmDeviceConnectivity200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstSeenAt** | Pointer to **string** | When the device was first seen as connected to the internet in each connection. | [optional] 
**LastSeenAt** | Pointer to **string** | When the device was last seen as connected to the internet in each connection. | [optional] 

## Methods

### NewGetNetworkSmDeviceConnectivity200ResponseInner

`func NewGetNetworkSmDeviceConnectivity200ResponseInner() *GetNetworkSmDeviceConnectivity200ResponseInner`

NewGetNetworkSmDeviceConnectivity200ResponseInner instantiates a new GetNetworkSmDeviceConnectivity200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmDeviceConnectivity200ResponseInnerWithDefaults

`func NewGetNetworkSmDeviceConnectivity200ResponseInnerWithDefaults() *GetNetworkSmDeviceConnectivity200ResponseInner`

NewGetNetworkSmDeviceConnectivity200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceConnectivity200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstSeenAt

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) GetFirstSeenAt() string`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) GetFirstSeenAtOk() (*string, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) SetFirstSeenAt(v string)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *GetNetworkSmDeviceConnectivity200ResponseInner) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


