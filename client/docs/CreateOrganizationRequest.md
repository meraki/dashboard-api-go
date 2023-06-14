# CreateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the organization | 
**Management** | Pointer to [**GetOrganizations200ResponseInnerManagement**](GetOrganizations200ResponseInnerManagement.md) |  | [optional] 

## Methods

### NewCreateOrganizationRequest

`func NewCreateOrganizationRequest(name string, ) *CreateOrganizationRequest`

NewCreateOrganizationRequest instantiates a new CreateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationRequestWithDefaults

`func NewCreateOrganizationRequestWithDefaults() *CreateOrganizationRequest`

NewCreateOrganizationRequestWithDefaults instantiates a new CreateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetManagement

`func (o *CreateOrganizationRequest) GetManagement() GetOrganizations200ResponseInnerManagement`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *CreateOrganizationRequest) GetManagementOk() (*GetOrganizations200ResponseInnerManagement, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *CreateOrganizationRequest) SetManagement(v GetOrganizations200ResponseInnerManagement)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *CreateOrganizationRequest) HasManagement() bool`

HasManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


