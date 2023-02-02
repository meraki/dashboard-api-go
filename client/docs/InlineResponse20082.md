# InlineResponse20082

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | Pointer to **string** | Database ID for the admin user who made the API request. | [optional] 
**Method** | Pointer to **string** | HTTP method used in the API request. | [optional] 
**Host** | Pointer to **string** | The host which the API request was directed at. | [optional] 
**Path** | Pointer to **string** | The API request path. | [optional] 
**QueryString** | Pointer to **string** | The query string sent with the API request. | [optional] 
**UserAgent** | Pointer to **string** | The API request user agent. | [optional] 
**Ts** | Pointer to **time.Time** | Timestamp, in iso8601 format, indicating when the API request was made. | [optional] 
**ResponseCode** | Pointer to **int32** | API request response code. | [optional] 
**SourceIp** | Pointer to **string** | Public IP address from which the API request was made. | [optional] 
**Version** | Pointer to **int32** | API version of the endpoint. | [optional] 
**OperationId** | Pointer to **string** | Operation ID for the endpoint. | [optional] 

## Methods

### NewInlineResponse20082

`func NewInlineResponse20082() *InlineResponse20082`

NewInlineResponse20082 instantiates a new InlineResponse20082 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20082WithDefaults

`func NewInlineResponse20082WithDefaults() *InlineResponse20082`

NewInlineResponse20082WithDefaults instantiates a new InlineResponse20082 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *InlineResponse20082) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *InlineResponse20082) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *InlineResponse20082) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *InlineResponse20082) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetMethod

`func (o *InlineResponse20082) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineResponse20082) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineResponse20082) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineResponse20082) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetHost

`func (o *InlineResponse20082) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineResponse20082) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineResponse20082) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineResponse20082) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPath

`func (o *InlineResponse20082) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InlineResponse20082) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InlineResponse20082) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *InlineResponse20082) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetQueryString

`func (o *InlineResponse20082) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *InlineResponse20082) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *InlineResponse20082) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *InlineResponse20082) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetUserAgent

`func (o *InlineResponse20082) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *InlineResponse20082) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *InlineResponse20082) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *InlineResponse20082) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse20082) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse20082) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse20082) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse20082) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetResponseCode

`func (o *InlineResponse20082) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *InlineResponse20082) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *InlineResponse20082) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *InlineResponse20082) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetSourceIp

`func (o *InlineResponse20082) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *InlineResponse20082) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *InlineResponse20082) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *InlineResponse20082) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetVersion

`func (o *InlineResponse20082) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineResponse20082) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineResponse20082) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineResponse20082) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOperationId

`func (o *InlineResponse20082) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *InlineResponse20082) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *InlineResponse20082) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *InlineResponse20082) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


