# MoveOrganizationLicensesSeatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestOrganizationId** | **string** | The ID of the organization to move the SM seats to | 
**LicenseId** | **string** | The ID of the SM license to move the seats from | 
**SeatCount** | **int32** | The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license | 

## Methods

### NewMoveOrganizationLicensesSeatsRequest

`func NewMoveOrganizationLicensesSeatsRequest(destOrganizationId string, licenseId string, seatCount int32, ) *MoveOrganizationLicensesSeatsRequest`

NewMoveOrganizationLicensesSeatsRequest instantiates a new MoveOrganizationLicensesSeatsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveOrganizationLicensesSeatsRequestWithDefaults

`func NewMoveOrganizationLicensesSeatsRequestWithDefaults() *MoveOrganizationLicensesSeatsRequest`

NewMoveOrganizationLicensesSeatsRequestWithDefaults instantiates a new MoveOrganizationLicensesSeatsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestOrganizationId

`func (o *MoveOrganizationLicensesSeatsRequest) GetDestOrganizationId() string`

GetDestOrganizationId returns the DestOrganizationId field if non-nil, zero value otherwise.

### GetDestOrganizationIdOk

`func (o *MoveOrganizationLicensesSeatsRequest) GetDestOrganizationIdOk() (*string, bool)`

GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOrganizationId

`func (o *MoveOrganizationLicensesSeatsRequest) SetDestOrganizationId(v string)`

SetDestOrganizationId sets DestOrganizationId field to given value.


### GetLicenseId

`func (o *MoveOrganizationLicensesSeatsRequest) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *MoveOrganizationLicensesSeatsRequest) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *MoveOrganizationLicensesSeatsRequest) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetSeatCount

`func (o *MoveOrganizationLicensesSeatsRequest) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *MoveOrganizationLicensesSeatsRequest) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *MoveOrganizationLicensesSeatsRequest) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


