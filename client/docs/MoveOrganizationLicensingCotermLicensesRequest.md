# MoveOrganizationLicensingCotermLicensesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**MoveOrganizationLicensingCotermLicensesRequestDestination**](MoveOrganizationLicensingCotermLicensesRequestDestination.md) |  | 
**Licenses** | [**[]MoveOrganizationLicensingCotermLicensesRequestLicensesInner**](MoveOrganizationLicensingCotermLicensesRequestLicensesInner.md) | The list of licenses to move | 

## Methods

### NewMoveOrganizationLicensingCotermLicensesRequest

`func NewMoveOrganizationLicensingCotermLicensesRequest(destination MoveOrganizationLicensingCotermLicensesRequestDestination, licenses []MoveOrganizationLicensingCotermLicensesRequestLicensesInner, ) *MoveOrganizationLicensingCotermLicensesRequest`

NewMoveOrganizationLicensingCotermLicensesRequest instantiates a new MoveOrganizationLicensingCotermLicensesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveOrganizationLicensingCotermLicensesRequestWithDefaults

`func NewMoveOrganizationLicensingCotermLicensesRequestWithDefaults() *MoveOrganizationLicensingCotermLicensesRequest`

NewMoveOrganizationLicensingCotermLicensesRequestWithDefaults instantiates a new MoveOrganizationLicensingCotermLicensesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *MoveOrganizationLicensingCotermLicensesRequest) GetDestination() MoveOrganizationLicensingCotermLicensesRequestDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MoveOrganizationLicensingCotermLicensesRequest) GetDestinationOk() (*MoveOrganizationLicensingCotermLicensesRequestDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MoveOrganizationLicensingCotermLicensesRequest) SetDestination(v MoveOrganizationLicensingCotermLicensesRequestDestination)`

SetDestination sets Destination field to given value.


### GetLicenses

`func (o *MoveOrganizationLicensingCotermLicensesRequest) GetLicenses() []MoveOrganizationLicensingCotermLicensesRequestLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *MoveOrganizationLicensingCotermLicensesRequest) GetLicensesOk() (*[]MoveOrganizationLicensingCotermLicensesRequestLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *MoveOrganizationLicensingCotermLicensesRequest) SetLicenses(v []MoveOrganizationLicensingCotermLicensesRequestLicensesInner)`

SetLicenses sets Licenses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


