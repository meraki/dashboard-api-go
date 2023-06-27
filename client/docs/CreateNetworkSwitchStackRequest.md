# CreateNetworkSwitchStackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new stack | 
**Serials** | **[]string** | An array of switch serials to be added into the new stack | 

## Methods

### NewCreateNetworkSwitchStackRequest

`func NewCreateNetworkSwitchStackRequest(name string, serials []string, ) *CreateNetworkSwitchStackRequest`

NewCreateNetworkSwitchStackRequest instantiates a new CreateNetworkSwitchStackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkSwitchStackRequestWithDefaults

`func NewCreateNetworkSwitchStackRequestWithDefaults() *CreateNetworkSwitchStackRequest`

NewCreateNetworkSwitchStackRequestWithDefaults instantiates a new CreateNetworkSwitchStackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkSwitchStackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkSwitchStackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkSwitchStackRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSerials

`func (o *CreateNetworkSwitchStackRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *CreateNetworkSwitchStackRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *CreateNetworkSwitchStackRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


