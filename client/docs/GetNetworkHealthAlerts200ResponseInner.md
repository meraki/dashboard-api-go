# GetNetworkHealthAlerts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alert identifier. Value can be empty | [optional] 
**Category** | Pointer to **string** | Category of the alert | [optional] 
**Type** | Pointer to **string** | Alert type | [optional] 
**Severity** | Pointer to **string** | Severity of the alert | [optional] 
**Scope** | Pointer to [**GetNetworkHealthAlerts200ResponseInnerScope**](GetNetworkHealthAlerts200ResponseInnerScope.md) |  | [optional] 

## Methods

### NewGetNetworkHealthAlerts200ResponseInner

`func NewGetNetworkHealthAlerts200ResponseInner() *GetNetworkHealthAlerts200ResponseInner`

NewGetNetworkHealthAlerts200ResponseInner instantiates a new GetNetworkHealthAlerts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkHealthAlerts200ResponseInnerWithDefaults

`func NewGetNetworkHealthAlerts200ResponseInnerWithDefaults() *GetNetworkHealthAlerts200ResponseInner`

NewGetNetworkHealthAlerts200ResponseInnerWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkHealthAlerts200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkHealthAlerts200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkHealthAlerts200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkHealthAlerts200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *GetNetworkHealthAlerts200ResponseInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNetworkHealthAlerts200ResponseInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNetworkHealthAlerts200ResponseInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetNetworkHealthAlerts200ResponseInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkHealthAlerts200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkHealthAlerts200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkHealthAlerts200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkHealthAlerts200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *GetNetworkHealthAlerts200ResponseInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetNetworkHealthAlerts200ResponseInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetNetworkHealthAlerts200ResponseInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetNetworkHealthAlerts200ResponseInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetScope

`func (o *GetNetworkHealthAlerts200ResponseInner) GetScope() GetNetworkHealthAlerts200ResponseInnerScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetNetworkHealthAlerts200ResponseInner) GetScopeOk() (*GetNetworkHealthAlerts200ResponseInnerScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetNetworkHealthAlerts200ResponseInner) SetScope(v GetNetworkHealthAlerts200ResponseInnerScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetNetworkHealthAlerts200ResponseInner) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


