# MoveOrganizationLicensingCotermLicenses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainderLicenses** | Pointer to [**[]GetOrganizationLicensingCotermLicenses200ResponseInner**](GetOrganizationLicensingCotermLicenses200ResponseInner.md) | Remainder licenses created in the source organization as a result of moving a subset of the counts of a license | [optional] 
**MovedLicenses** | Pointer to [**[]GetOrganizationLicensingCotermLicenses200ResponseInner**](GetOrganizationLicensingCotermLicenses200ResponseInner.md) | Newly moved licenses created in the destination organization of the license move operation | [optional] 

## Methods

### NewMoveOrganizationLicensingCotermLicenses200Response

`func NewMoveOrganizationLicensingCotermLicenses200Response() *MoveOrganizationLicensingCotermLicenses200Response`

NewMoveOrganizationLicensingCotermLicenses200Response instantiates a new MoveOrganizationLicensingCotermLicenses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveOrganizationLicensingCotermLicenses200ResponseWithDefaults

`func NewMoveOrganizationLicensingCotermLicenses200ResponseWithDefaults() *MoveOrganizationLicensingCotermLicenses200Response`

NewMoveOrganizationLicensingCotermLicenses200ResponseWithDefaults instantiates a new MoveOrganizationLicensingCotermLicenses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainderLicenses

`func (o *MoveOrganizationLicensingCotermLicenses200Response) GetRemainderLicenses() []GetOrganizationLicensingCotermLicenses200ResponseInner`

GetRemainderLicenses returns the RemainderLicenses field if non-nil, zero value otherwise.

### GetRemainderLicensesOk

`func (o *MoveOrganizationLicensingCotermLicenses200Response) GetRemainderLicensesOk() (*[]GetOrganizationLicensingCotermLicenses200ResponseInner, bool)`

GetRemainderLicensesOk returns a tuple with the RemainderLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainderLicenses

`func (o *MoveOrganizationLicensingCotermLicenses200Response) SetRemainderLicenses(v []GetOrganizationLicensingCotermLicenses200ResponseInner)`

SetRemainderLicenses sets RemainderLicenses field to given value.

### HasRemainderLicenses

`func (o *MoveOrganizationLicensingCotermLicenses200Response) HasRemainderLicenses() bool`

HasRemainderLicenses returns a boolean if a field has been set.

### GetMovedLicenses

`func (o *MoveOrganizationLicensingCotermLicenses200Response) GetMovedLicenses() []GetOrganizationLicensingCotermLicenses200ResponseInner`

GetMovedLicenses returns the MovedLicenses field if non-nil, zero value otherwise.

### GetMovedLicensesOk

`func (o *MoveOrganizationLicensingCotermLicenses200Response) GetMovedLicensesOk() (*[]GetOrganizationLicensingCotermLicenses200ResponseInner, bool)`

GetMovedLicensesOk returns a tuple with the MovedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedLicenses

`func (o *MoveOrganizationLicensingCotermLicenses200Response) SetMovedLicenses(v []GetOrganizationLicensingCotermLicenses200ResponseInner)`

SetMovedLicenses sets MovedLicenses field to given value.

### HasMovedLicenses

`func (o *MoveOrganizationLicensingCotermLicenses200Response) HasMovedLicenses() bool`

HasMovedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


