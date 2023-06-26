# GetOrganizationLicenses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | License ID | [optional] 
**LicenseType** | Pointer to **string** | License type | [optional] 
**LicenseKey** | Pointer to **string** | License key | [optional] 
**OrderNumber** | Pointer to **string** | Order number | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device the license is assigned to | [optional] 
**NetworkId** | Pointer to **string** | ID of the network the license is assigned to | [optional] 
**State** | Pointer to **string** | The state of the license. All queued licenses have a status of &#x60;recentlyQueued&#x60;. | [optional] 
**SeatCount** | Pointer to **int32** | The number of seats of the license. Only applicable to SM licenses. | [optional] 
**TotalDurationInDays** | Pointer to **int32** | The duration of the license plus all permanently queued licenses associated with it | [optional] 
**DurationInDays** | Pointer to **int32** | The duration of the individual license | [optional] 
**PermanentlyQueuedLicenses** | Pointer to [**[]GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner**](GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner.md) | DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial&#x3D; to retrieved queued licenses for a given device. | [optional] 
**ClaimDate** | Pointer to **string** | The date the license was claimed into the organization | [optional] 
**ActivationDate** | Pointer to **string** | The date the license started burning | [optional] 
**ExpirationDate** | Pointer to **string** | The date the license will expire | [optional] 
**HeadLicenseId** | Pointer to **string** | The id of the head license this license is queued behind. If there is no head license, it returns nil. | [optional] 

## Methods

### NewGetOrganizationLicenses200ResponseInner

`func NewGetOrganizationLicenses200ResponseInner() *GetOrganizationLicenses200ResponseInner`

NewGetOrganizationLicenses200ResponseInner instantiates a new GetOrganizationLicenses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationLicenses200ResponseInnerWithDefaults

`func NewGetOrganizationLicenses200ResponseInnerWithDefaults() *GetOrganizationLicenses200ResponseInner`

NewGetOrganizationLicenses200ResponseInnerWithDefaults instantiates a new GetOrganizationLicenses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationLicenses200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationLicenses200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationLicenses200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationLicenses200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseType

`func (o *GetOrganizationLicenses200ResponseInner) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *GetOrganizationLicenses200ResponseInner) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *GetOrganizationLicenses200ResponseInner) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *GetOrganizationLicenses200ResponseInner) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseKey

`func (o *GetOrganizationLicenses200ResponseInner) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *GetOrganizationLicenses200ResponseInner) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *GetOrganizationLicenses200ResponseInner) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *GetOrganizationLicenses200ResponseInner) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetOrderNumber

`func (o *GetOrganizationLicenses200ResponseInner) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *GetOrganizationLicenses200ResponseInner) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *GetOrganizationLicenses200ResponseInner) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *GetOrganizationLicenses200ResponseInner) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *GetOrganizationLicenses200ResponseInner) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *GetOrganizationLicenses200ResponseInner) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *GetOrganizationLicenses200ResponseInner) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *GetOrganizationLicenses200ResponseInner) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetOrganizationLicenses200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationLicenses200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationLicenses200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationLicenses200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetState

`func (o *GetOrganizationLicenses200ResponseInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetOrganizationLicenses200ResponseInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetOrganizationLicenses200ResponseInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetOrganizationLicenses200ResponseInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSeatCount

`func (o *GetOrganizationLicenses200ResponseInner) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *GetOrganizationLicenses200ResponseInner) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *GetOrganizationLicenses200ResponseInner) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.

### HasSeatCount

`func (o *GetOrganizationLicenses200ResponseInner) HasSeatCount() bool`

HasSeatCount returns a boolean if a field has been set.

### GetTotalDurationInDays

`func (o *GetOrganizationLicenses200ResponseInner) GetTotalDurationInDays() int32`

GetTotalDurationInDays returns the TotalDurationInDays field if non-nil, zero value otherwise.

### GetTotalDurationInDaysOk

`func (o *GetOrganizationLicenses200ResponseInner) GetTotalDurationInDaysOk() (*int32, bool)`

GetTotalDurationInDaysOk returns a tuple with the TotalDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationInDays

`func (o *GetOrganizationLicenses200ResponseInner) SetTotalDurationInDays(v int32)`

SetTotalDurationInDays sets TotalDurationInDays field to given value.

### HasTotalDurationInDays

`func (o *GetOrganizationLicenses200ResponseInner) HasTotalDurationInDays() bool`

HasTotalDurationInDays returns a boolean if a field has been set.

### GetDurationInDays

`func (o *GetOrganizationLicenses200ResponseInner) GetDurationInDays() int32`

GetDurationInDays returns the DurationInDays field if non-nil, zero value otherwise.

### GetDurationInDaysOk

`func (o *GetOrganizationLicenses200ResponseInner) GetDurationInDaysOk() (*int32, bool)`

GetDurationInDaysOk returns a tuple with the DurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInDays

`func (o *GetOrganizationLicenses200ResponseInner) SetDurationInDays(v int32)`

SetDurationInDays sets DurationInDays field to given value.

### HasDurationInDays

`func (o *GetOrganizationLicenses200ResponseInner) HasDurationInDays() bool`

HasDurationInDays returns a boolean if a field has been set.

### GetPermanentlyQueuedLicenses

`func (o *GetOrganizationLicenses200ResponseInner) GetPermanentlyQueuedLicenses() []GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner`

GetPermanentlyQueuedLicenses returns the PermanentlyQueuedLicenses field if non-nil, zero value otherwise.

### GetPermanentlyQueuedLicensesOk

`func (o *GetOrganizationLicenses200ResponseInner) GetPermanentlyQueuedLicensesOk() (*[]GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner, bool)`

GetPermanentlyQueuedLicensesOk returns a tuple with the PermanentlyQueuedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentlyQueuedLicenses

`func (o *GetOrganizationLicenses200ResponseInner) SetPermanentlyQueuedLicenses(v []GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner)`

SetPermanentlyQueuedLicenses sets PermanentlyQueuedLicenses field to given value.

### HasPermanentlyQueuedLicenses

`func (o *GetOrganizationLicenses200ResponseInner) HasPermanentlyQueuedLicenses() bool`

HasPermanentlyQueuedLicenses returns a boolean if a field has been set.

### GetClaimDate

`func (o *GetOrganizationLicenses200ResponseInner) GetClaimDate() string`

GetClaimDate returns the ClaimDate field if non-nil, zero value otherwise.

### GetClaimDateOk

`func (o *GetOrganizationLicenses200ResponseInner) GetClaimDateOk() (*string, bool)`

GetClaimDateOk returns a tuple with the ClaimDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimDate

`func (o *GetOrganizationLicenses200ResponseInner) SetClaimDate(v string)`

SetClaimDate sets ClaimDate field to given value.

### HasClaimDate

`func (o *GetOrganizationLicenses200ResponseInner) HasClaimDate() bool`

HasClaimDate returns a boolean if a field has been set.

### GetActivationDate

`func (o *GetOrganizationLicenses200ResponseInner) GetActivationDate() string`

GetActivationDate returns the ActivationDate field if non-nil, zero value otherwise.

### GetActivationDateOk

`func (o *GetOrganizationLicenses200ResponseInner) GetActivationDateOk() (*string, bool)`

GetActivationDateOk returns a tuple with the ActivationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationDate

`func (o *GetOrganizationLicenses200ResponseInner) SetActivationDate(v string)`

SetActivationDate sets ActivationDate field to given value.

### HasActivationDate

`func (o *GetOrganizationLicenses200ResponseInner) HasActivationDate() bool`

HasActivationDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *GetOrganizationLicenses200ResponseInner) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GetOrganizationLicenses200ResponseInner) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GetOrganizationLicenses200ResponseInner) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *GetOrganizationLicenses200ResponseInner) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHeadLicenseId

`func (o *GetOrganizationLicenses200ResponseInner) GetHeadLicenseId() string`

GetHeadLicenseId returns the HeadLicenseId field if non-nil, zero value otherwise.

### GetHeadLicenseIdOk

`func (o *GetOrganizationLicenses200ResponseInner) GetHeadLicenseIdOk() (*string, bool)`

GetHeadLicenseIdOk returns a tuple with the HeadLicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadLicenseId

`func (o *GetOrganizationLicenses200ResponseInner) SetHeadLicenseId(v string)`

SetHeadLicenseId sets HeadLicenseId field to given value.

### HasHeadLicenseId

`func (o *GetOrganizationLicenses200ResponseInner) HasHeadLicenseId() bool`

HasHeadLicenseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


