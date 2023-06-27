# CreateNetworkAppliancePrefixesDelegatedStaticRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** | A static IPv6 prefix | 
**Origin** | [**CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin**](CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin.md) |  | 
**Description** | Pointer to **string** | A name or description for the prefix | [optional] 

## Methods

### NewCreateNetworkAppliancePrefixesDelegatedStaticRequest

`func NewCreateNetworkAppliancePrefixesDelegatedStaticRequest(prefix string, origin CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, ) *CreateNetworkAppliancePrefixesDelegatedStaticRequest`

NewCreateNetworkAppliancePrefixesDelegatedStaticRequest instantiates a new CreateNetworkAppliancePrefixesDelegatedStaticRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults

`func NewCreateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults() *CreateNetworkAppliancePrefixesDelegatedStaticRequest`

NewCreateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults instantiates a new CreateNetworkAppliancePrefixesDelegatedStaticRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetOrigin

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetOrigin() CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetOriginOk() (*CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) SetOrigin(v CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin)`

SetOrigin sets Origin field to given value.


### GetDescription

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkAppliancePrefixesDelegatedStaticRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


