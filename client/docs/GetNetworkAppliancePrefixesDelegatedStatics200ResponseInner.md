# GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticDelegatedPrefixId** | Pointer to **string** | Static delegated prefix id. | [optional] 
**Prefix** | Pointer to **string** | IPv6 prefix/prefix length. | [optional] 
**Origin** | Pointer to [**GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin**](GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin.md) |  | [optional] 
**Description** | Pointer to **string** | Identifying description for the prefix. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Prefix creation time. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Prefix Updated time. | [optional] 

## Methods

### NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner

`func NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner`

NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner instantiates a new GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerWithDefaults

`func NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerWithDefaults() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner`

NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerWithDefaults instantiates a new GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticDelegatedPrefixId

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetStaticDelegatedPrefixId() string`

GetStaticDelegatedPrefixId returns the StaticDelegatedPrefixId field if non-nil, zero value otherwise.

### GetStaticDelegatedPrefixIdOk

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetStaticDelegatedPrefixIdOk() (*string, bool)`

GetStaticDelegatedPrefixIdOk returns a tuple with the StaticDelegatedPrefixId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDelegatedPrefixId

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetStaticDelegatedPrefixId(v string)`

SetStaticDelegatedPrefixId sets StaticDelegatedPrefixId field to given value.

### HasStaticDelegatedPrefixId

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasStaticDelegatedPrefixId() bool`

HasStaticDelegatedPrefixId returns a boolean if a field has been set.

### GetPrefix

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetOrigin

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetOrigin() GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetOriginOk() (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetOrigin(v GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


