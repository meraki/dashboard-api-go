# BindAdministeredLicensingSubscriptionSubscription200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Subscription ID | [optional] 
**Networks** | Pointer to [**[]BindAdministeredLicensingSubscriptionSubscription200ResponseNetworksInner**](BindAdministeredLicensingSubscriptionSubscription200ResponseNetworksInner.md) | Unbound networks | [optional] 
**Errors** | Pointer to **[]string** | Array of errors if failed | [optional] 
**InsufficientEntitlements** | Pointer to [**[]BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner**](BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner.md) | A list of entitlements required to successfully bind the networks to the subscription | [optional] 

## Methods

### NewBindAdministeredLicensingSubscriptionSubscription200Response

`func NewBindAdministeredLicensingSubscriptionSubscription200Response() *BindAdministeredLicensingSubscriptionSubscription200Response`

NewBindAdministeredLicensingSubscriptionSubscription200Response instantiates a new BindAdministeredLicensingSubscriptionSubscription200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindAdministeredLicensingSubscriptionSubscription200ResponseWithDefaults

`func NewBindAdministeredLicensingSubscriptionSubscription200ResponseWithDefaults() *BindAdministeredLicensingSubscriptionSubscription200Response`

NewBindAdministeredLicensingSubscriptionSubscription200ResponseWithDefaults instantiates a new BindAdministeredLicensingSubscriptionSubscription200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetNetworks

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetNetworks() []BindAdministeredLicensingSubscriptionSubscription200ResponseNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetNetworksOk() (*[]BindAdministeredLicensingSubscriptionSubscription200ResponseNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) SetNetworks(v []BindAdministeredLicensingSubscriptionSubscription200ResponseNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetErrors

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetInsufficientEntitlements

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetInsufficientEntitlements() []BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner`

GetInsufficientEntitlements returns the InsufficientEntitlements field if non-nil, zero value otherwise.

### GetInsufficientEntitlementsOk

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) GetInsufficientEntitlementsOk() (*[]BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner, bool)`

GetInsufficientEntitlementsOk returns a tuple with the InsufficientEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientEntitlements

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) SetInsufficientEntitlements(v []BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner)`

SetInsufficientEntitlements sets InsufficientEntitlements field to given value.

### HasInsufficientEntitlements

`func (o *BindAdministeredLicensingSubscriptionSubscription200Response) HasInsufficientEntitlements() bool`

HasInsufficientEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


