# UpdateOrganizationAdaptivePolicySettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledNetworks** | Pointer to **[]string** | List of network IDs with adaptive policy enabled | [optional] 

## Methods

### NewUpdateOrganizationAdaptivePolicySettingsRequest

`func NewUpdateOrganizationAdaptivePolicySettingsRequest() *UpdateOrganizationAdaptivePolicySettingsRequest`

NewUpdateOrganizationAdaptivePolicySettingsRequest instantiates a new UpdateOrganizationAdaptivePolicySettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationAdaptivePolicySettingsRequestWithDefaults

`func NewUpdateOrganizationAdaptivePolicySettingsRequestWithDefaults() *UpdateOrganizationAdaptivePolicySettingsRequest`

NewUpdateOrganizationAdaptivePolicySettingsRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicySettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledNetworks

`func (o *UpdateOrganizationAdaptivePolicySettingsRequest) GetEnabledNetworks() []string`

GetEnabledNetworks returns the EnabledNetworks field if non-nil, zero value otherwise.

### GetEnabledNetworksOk

`func (o *UpdateOrganizationAdaptivePolicySettingsRequest) GetEnabledNetworksOk() (*[]string, bool)`

GetEnabledNetworksOk returns a tuple with the EnabledNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledNetworks

`func (o *UpdateOrganizationAdaptivePolicySettingsRequest) SetEnabledNetworks(v []string)`

SetEnabledNetworks sets EnabledNetworks field to given value.

### HasEnabledNetworks

`func (o *UpdateOrganizationAdaptivePolicySettingsRequest) HasEnabledNetworks() bool`

HasEnabledNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


