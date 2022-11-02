# MoveOrganizationLicensesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestOrganizationId** | **string** | The ID of the organization to move the licenses to | 
**LicenseIds** | **[]string** | A list of IDs of licenses to move to the new organization | 

## Methods

### NewMoveOrganizationLicensesRequest

`func NewMoveOrganizationLicensesRequest(destOrganizationId string, licenseIds []string, ) *MoveOrganizationLicensesRequest`

NewMoveOrganizationLicensesRequest instantiates a new MoveOrganizationLicensesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveOrganizationLicensesRequestWithDefaults

`func NewMoveOrganizationLicensesRequestWithDefaults() *MoveOrganizationLicensesRequest`

NewMoveOrganizationLicensesRequestWithDefaults instantiates a new MoveOrganizationLicensesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestOrganizationId

`func (o *MoveOrganizationLicensesRequest) GetDestOrganizationId() string`

GetDestOrganizationId returns the DestOrganizationId field if non-nil, zero value otherwise.

### GetDestOrganizationIdOk

`func (o *MoveOrganizationLicensesRequest) GetDestOrganizationIdOk() (*string, bool)`

GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOrganizationId

`func (o *MoveOrganizationLicensesRequest) SetDestOrganizationId(v string)`

SetDestOrganizationId sets DestOrganizationId field to given value.


### GetLicenseIds

`func (o *MoveOrganizationLicensesRequest) GetLicenseIds() []string`

GetLicenseIds returns the LicenseIds field if non-nil, zero value otherwise.

### GetLicenseIdsOk

`func (o *MoveOrganizationLicensesRequest) GetLicenseIdsOk() (*[]string, bool)`

GetLicenseIdsOk returns a tuple with the LicenseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseIds

`func (o *MoveOrganizationLicensesRequest) SetLicenseIds(v []string)`

SetLicenseIds sets LicenseIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


