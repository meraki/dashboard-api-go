# InlineResponse20029

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alert identifier. Value can be empty | [optional] 
**Category** | Pointer to **string** | Category of the alert | [optional] 
**Type** | Pointer to **string** | Alert type | [optional] 
**Severity** | Pointer to **string** | Severity of the alert | [optional] 
**Scope** | Pointer to [**NetworksNetworkIdHealthAlertsScope**](NetworksNetworkIdHealthAlertsScope.md) |  | [optional] 

## Methods

### NewInlineResponse20029

`func NewInlineResponse20029() *InlineResponse20029`

NewInlineResponse20029 instantiates a new InlineResponse20029 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029WithDefaults

`func NewInlineResponse20029WithDefaults() *InlineResponse20029`

NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20029) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20029) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20029) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20029) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse20029) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20029) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20029) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20029) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20029) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20029) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20029) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20029) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *InlineResponse20029) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InlineResponse20029) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InlineResponse20029) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InlineResponse20029) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetScope

`func (o *InlineResponse20029) GetScope() NetworksNetworkIdHealthAlertsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineResponse20029) GetScopeOk() (*NetworksNetworkIdHealthAlertsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineResponse20029) SetScope(v NetworksNetworkIdHealthAlertsScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineResponse20029) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


