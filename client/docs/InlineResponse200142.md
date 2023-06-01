# InlineResponse200142

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | Pointer to **string** | Type of alert that the webhook is delivering | [optional] 
**LoggedAt** | Pointer to **time.Time** | When the webhook log was created, in ISO8601 format | [optional] 
**NetworkId** | Pointer to **string** | Network ID for the webhook log | [optional] 
**OrganizationId** | Pointer to **string** | ID for the webhook log&#39;s organization | [optional] 
**ResponseCode** | Pointer to **int32** | Response code from the webhook | [optional] 
**ResponseDuration** | Pointer to **int32** | Duration of the response, in milliseconds | [optional] 
**SentAt** | Pointer to **time.Time** | When the webhook was sent, in ISO8601 format | [optional] 
**Url** | Pointer to **string** | URL where the webhook was sent | [optional] 

## Methods

### NewInlineResponse200142

`func NewInlineResponse200142() *InlineResponse200142`

NewInlineResponse200142 instantiates a new InlineResponse200142 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200142WithDefaults

`func NewInlineResponse200142WithDefaults() *InlineResponse200142`

NewInlineResponse200142WithDefaults instantiates a new InlineResponse200142 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertType

`func (o *InlineResponse200142) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *InlineResponse200142) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *InlineResponse200142) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *InlineResponse200142) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetLoggedAt

`func (o *InlineResponse200142) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *InlineResponse200142) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *InlineResponse200142) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.

### HasLoggedAt

`func (o *InlineResponse200142) HasLoggedAt() bool`

HasLoggedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200142) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200142) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200142) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200142) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse200142) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse200142) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse200142) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse200142) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetResponseCode

`func (o *InlineResponse200142) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *InlineResponse200142) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *InlineResponse200142) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *InlineResponse200142) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetResponseDuration

`func (o *InlineResponse200142) GetResponseDuration() int32`

GetResponseDuration returns the ResponseDuration field if non-nil, zero value otherwise.

### GetResponseDurationOk

`func (o *InlineResponse200142) GetResponseDurationOk() (*int32, bool)`

GetResponseDurationOk returns a tuple with the ResponseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDuration

`func (o *InlineResponse200142) SetResponseDuration(v int32)`

SetResponseDuration sets ResponseDuration field to given value.

### HasResponseDuration

`func (o *InlineResponse200142) HasResponseDuration() bool`

HasResponseDuration returns a boolean if a field has been set.

### GetSentAt

`func (o *InlineResponse200142) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InlineResponse200142) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InlineResponse200142) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InlineResponse200142) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200142) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200142) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200142) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200142) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


