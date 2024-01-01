# GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Subscription&#39;s ID | [optional] 
**Name** | Pointer to **string** | Subscription name | [optional] 
**Description** | Pointer to **string** | Subscription description | [optional] 
**Status** | Pointer to **string** | Subscription status | [optional] 
**StartDate** | Pointer to **time.Time** | Subscription start date | [optional] 
**EndDate** | Pointer to **time.Time** | Subscription expiration date | [optional] 
**WebOrderId** | Pointer to **string** | Web order id | [optional] 
**ProductTypes** | Pointer to **[]string** | Products the subscription has entitlements for | [optional] 
**Entitlements** | Pointer to [**[]GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner**](GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner.md) | Entitlement info | [optional] 
**Counts** | Pointer to [**GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts**](GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts.md) |  | [optional] 

## Methods

### NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner

`func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner`

NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInner instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerWithDefaults

`func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerWithDefaults() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner`

NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerWithDefaults instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetName

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetWebOrderId

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetWebOrderId() string`

GetWebOrderId returns the WebOrderId field if non-nil, zero value otherwise.

### GetWebOrderIdOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetWebOrderIdOk() (*string, bool)`

GetWebOrderIdOk returns a tuple with the WebOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebOrderId

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetWebOrderId(v string)`

SetWebOrderId sets WebOrderId field to given value.

### HasWebOrderId

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasWebOrderId() bool`

HasWebOrderId returns a boolean if a field has been set.

### GetProductTypes

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetEntitlements

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEntitlements() []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetEntitlementsOk() (*[]GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetEntitlements(v []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInner)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetCounts

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetCounts() GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) GetCountsOk() (*GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) SetCounts(v GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


