# ProvisionNetworkClientsRequestClientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | The MAC address of the client. Required. | 
**Name** | Pointer to **string** | The display name for the client. Optional. Limited to 255 bytes. | [optional] 

## Methods

### NewProvisionNetworkClientsRequestClientsInner

`func NewProvisionNetworkClientsRequestClientsInner(mac string, ) *ProvisionNetworkClientsRequestClientsInner`

NewProvisionNetworkClientsRequestClientsInner instantiates a new ProvisionNetworkClientsRequestClientsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionNetworkClientsRequestClientsInnerWithDefaults

`func NewProvisionNetworkClientsRequestClientsInnerWithDefaults() *ProvisionNetworkClientsRequestClientsInner`

NewProvisionNetworkClientsRequestClientsInnerWithDefaults instantiates a new ProvisionNetworkClientsRequestClientsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *ProvisionNetworkClientsRequestClientsInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ProvisionNetworkClientsRequestClientsInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ProvisionNetworkClientsRequestClientsInner) SetMac(v string)`

SetMac sets Mac field to given value.


### GetName

`func (o *ProvisionNetworkClientsRequestClientsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionNetworkClientsRequestClientsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionNetworkClientsRequestClientsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionNetworkClientsRequestClientsInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


