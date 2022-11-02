# GetNetworkSmTrustedAccessConfigs200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | device ID | [optional] 
**SsidName** | Pointer to **string** | SSID name | [optional] 
**Name** | Pointer to **string** | device name | [optional] 
**Scope** | Pointer to **string** | scope | [optional] 
**Tags** | Pointer to **[]string** | device tags | [optional] 
**AccessStartAt** | Pointer to **time.Time** | time that access starts | [optional] 
**AccessEndAt** | Pointer to **time.Time** | time that access ends | [optional] 

## Methods

### NewGetNetworkSmTrustedAccessConfigs200ResponseInner

`func NewGetNetworkSmTrustedAccessConfigs200ResponseInner() *GetNetworkSmTrustedAccessConfigs200ResponseInner`

NewGetNetworkSmTrustedAccessConfigs200ResponseInner instantiates a new GetNetworkSmTrustedAccessConfigs200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSmTrustedAccessConfigs200ResponseInnerWithDefaults

`func NewGetNetworkSmTrustedAccessConfigs200ResponseInnerWithDefaults() *GetNetworkSmTrustedAccessConfigs200ResponseInner`

NewGetNetworkSmTrustedAccessConfigs200ResponseInnerWithDefaults instantiates a new GetNetworkSmTrustedAccessConfigs200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSsidName

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTags

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAccessStartAt

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessStartAt() time.Time`

GetAccessStartAt returns the AccessStartAt field if non-nil, zero value otherwise.

### GetAccessStartAtOk

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessStartAtOk() (*time.Time, bool)`

GetAccessStartAtOk returns a tuple with the AccessStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStartAt

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetAccessStartAt(v time.Time)`

SetAccessStartAt sets AccessStartAt field to given value.

### HasAccessStartAt

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasAccessStartAt() bool`

HasAccessStartAt returns a boolean if a field has been set.

### GetAccessEndAt

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessEndAt() time.Time`

GetAccessEndAt returns the AccessEndAt field if non-nil, zero value otherwise.

### GetAccessEndAtOk

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessEndAtOk() (*time.Time, bool)`

GetAccessEndAtOk returns a tuple with the AccessEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessEndAt

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetAccessEndAt(v time.Time)`

SetAccessEndAt sets AccessEndAt field to given value.

### HasAccessEndAt

`func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasAccessEndAt() bool`

HasAccessEndAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


