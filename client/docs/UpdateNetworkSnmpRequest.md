# UpdateNetworkSnmpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | The SNMP community string. Only relevant if &#39;access&#39; is set to &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]UpdateNetworkSnmpRequestUsersInner**](UpdateNetworkSnmpRequestUsersInner.md) | The list of SNMP users. Only relevant if &#39;access&#39; is set to &#39;users&#39;. | [optional] 

## Methods

### NewUpdateNetworkSnmpRequest

`func NewUpdateNetworkSnmpRequest() *UpdateNetworkSnmpRequest`

NewUpdateNetworkSnmpRequest instantiates a new UpdateNetworkSnmpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSnmpRequestWithDefaults

`func NewUpdateNetworkSnmpRequestWithDefaults() *UpdateNetworkSnmpRequest`

NewUpdateNetworkSnmpRequestWithDefaults instantiates a new UpdateNetworkSnmpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *UpdateNetworkSnmpRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateNetworkSnmpRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateNetworkSnmpRequest) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UpdateNetworkSnmpRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *UpdateNetworkSnmpRequest) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *UpdateNetworkSnmpRequest) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *UpdateNetworkSnmpRequest) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *UpdateNetworkSnmpRequest) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *UpdateNetworkSnmpRequest) GetUsers() []UpdateNetworkSnmpRequestUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateNetworkSnmpRequest) GetUsersOk() (*[]UpdateNetworkSnmpRequestUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateNetworkSnmpRequest) SetUsers(v []UpdateNetworkSnmpRequestUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateNetworkSnmpRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


