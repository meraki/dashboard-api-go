# OrganizationsManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]OrganizationsManagementDetails**](OrganizationsManagementDetails.md) | Details related to organization management, possibly empty | [optional] 

## Methods

### NewOrganizationsManagement

`func NewOrganizationsManagement() *OrganizationsManagement`

NewOrganizationsManagement instantiates a new OrganizationsManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsManagementWithDefaults

`func NewOrganizationsManagementWithDefaults() *OrganizationsManagement`

NewOrganizationsManagementWithDefaults instantiates a new OrganizationsManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *OrganizationsManagement) GetDetails() []OrganizationsManagementDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OrganizationsManagement) GetDetailsOk() (*[]OrganizationsManagementDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OrganizationsManagement) SetDetails(v []OrganizationsManagementDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *OrganizationsManagement) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


