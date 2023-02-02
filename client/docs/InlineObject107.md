# InlineObject107

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | The SNMP community string. Only relevant if &#39;access&#39; is set to &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]NetworksNetworkIdSnmpUsers**](NetworksNetworkIdSnmpUsers.md) | The list of SNMP users. Only relevant if &#39;access&#39; is set to &#39;users&#39;. | [optional] 

## Methods

### NewInlineObject107

`func NewInlineObject107() *InlineObject107`

NewInlineObject107 instantiates a new InlineObject107 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject107WithDefaults

`func NewInlineObject107WithDefaults() *InlineObject107`

NewInlineObject107WithDefaults instantiates a new InlineObject107 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineObject107) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineObject107) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineObject107) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineObject107) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *InlineObject107) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *InlineObject107) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *InlineObject107) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *InlineObject107) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *InlineObject107) GetUsers() []NetworksNetworkIdSnmpUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineObject107) GetUsersOk() (*[]NetworksNetworkIdSnmpUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineObject107) SetUsers(v []NetworksNetworkIdSnmpUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineObject107) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


