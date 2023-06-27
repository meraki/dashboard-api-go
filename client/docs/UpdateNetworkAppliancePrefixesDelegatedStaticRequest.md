# UpdateNetworkAppliancePrefixesDelegatedStaticRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | A static IPv6 prefix | [optional] 
**Origin** | Pointer to [**CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin**](CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin.md) |  | [optional] 
**Description** | Pointer to **string** | A name or description for the prefix | [optional] 

## Methods

### NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest

`func NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest() *UpdateNetworkAppliancePrefixesDelegatedStaticRequest`

NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest instantiates a new UpdateNetworkAppliancePrefixesDelegatedStaticRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults

`func NewUpdateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults() *UpdateNetworkAppliancePrefixesDelegatedStaticRequest`

NewUpdateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults instantiates a new UpdateNetworkAppliancePrefixesDelegatedStaticRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetOrigin

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetOrigin() CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetOriginOk() (*CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) SetOrigin(v CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


