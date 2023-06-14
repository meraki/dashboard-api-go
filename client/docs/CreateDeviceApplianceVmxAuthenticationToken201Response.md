# CreateDeviceApplianceVmxAuthenticationToken201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The newly generated authentication token for the vMX instance | [optional] 
**ExpiresAt** | Pointer to **string** | The expiration time for the token, in ISO 8601 format | [optional] 

## Methods

### NewCreateDeviceApplianceVmxAuthenticationToken201Response

`func NewCreateDeviceApplianceVmxAuthenticationToken201Response() *CreateDeviceApplianceVmxAuthenticationToken201Response`

NewCreateDeviceApplianceVmxAuthenticationToken201Response instantiates a new CreateDeviceApplianceVmxAuthenticationToken201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceApplianceVmxAuthenticationToken201ResponseWithDefaults

`func NewCreateDeviceApplianceVmxAuthenticationToken201ResponseWithDefaults() *CreateDeviceApplianceVmxAuthenticationToken201Response`

NewCreateDeviceApplianceVmxAuthenticationToken201ResponseWithDefaults instantiates a new CreateDeviceApplianceVmxAuthenticationToken201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


