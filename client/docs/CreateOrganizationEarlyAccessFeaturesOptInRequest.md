# CreateOrganizationEarlyAccessFeaturesOptInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | **string** | Short name of the early access feature | 
**LimitScopeToNetworks** | Pointer to **[]string** | A list of network IDs to apply the opt-in to | [optional] 

## Methods

### NewCreateOrganizationEarlyAccessFeaturesOptInRequest

`func NewCreateOrganizationEarlyAccessFeaturesOptInRequest(shortName string, ) *CreateOrganizationEarlyAccessFeaturesOptInRequest`

NewCreateOrganizationEarlyAccessFeaturesOptInRequest instantiates a new CreateOrganizationEarlyAccessFeaturesOptInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationEarlyAccessFeaturesOptInRequestWithDefaults

`func NewCreateOrganizationEarlyAccessFeaturesOptInRequestWithDefaults() *CreateOrganizationEarlyAccessFeaturesOptInRequest`

NewCreateOrganizationEarlyAccessFeaturesOptInRequestWithDefaults instantiates a new CreateOrganizationEarlyAccessFeaturesOptInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetLimitScopeToNetworks

`func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetLimitScopeToNetworks() []string`

GetLimitScopeToNetworks returns the LimitScopeToNetworks field if non-nil, zero value otherwise.

### GetLimitScopeToNetworksOk

`func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) GetLimitScopeToNetworksOk() (*[]string, bool)`

GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitScopeToNetworks

`func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) SetLimitScopeToNetworks(v []string)`

SetLimitScopeToNetworks sets LimitScopeToNetworks field to given value.

### HasLimitScopeToNetworks

`func (o *CreateOrganizationEarlyAccessFeaturesOptInRequest) HasLimitScopeToNetworks() bool`

HasLimitScopeToNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


