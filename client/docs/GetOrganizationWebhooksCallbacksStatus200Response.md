# GetOrganizationWebhooksCallbacksStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackId** | Pointer to **string** | The ID of the callback | [optional] 
**Status** | Pointer to **string** | The status of the callback | [optional] 
**Errors** | Pointer to **[]string** | The errors returned by the callback | [optional] 
**CreatedBy** | Pointer to [**GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy**](GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy.md) |  | [optional] 
**Webhook** | Pointer to [**GetOrganizationWebhooksCallbacksStatus200ResponseWebhook**](GetOrganizationWebhooksCallbacksStatus200ResponseWebhook.md) |  | [optional] 

## Methods

### NewGetOrganizationWebhooksCallbacksStatus200Response

`func NewGetOrganizationWebhooksCallbacksStatus200Response() *GetOrganizationWebhooksCallbacksStatus200Response`

NewGetOrganizationWebhooksCallbacksStatus200Response instantiates a new GetOrganizationWebhooksCallbacksStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWebhooksCallbacksStatus200ResponseWithDefaults

`func NewGetOrganizationWebhooksCallbacksStatus200ResponseWithDefaults() *GetOrganizationWebhooksCallbacksStatus200Response`

NewGetOrganizationWebhooksCallbacksStatus200ResponseWithDefaults instantiates a new GetOrganizationWebhooksCallbacksStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackId

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.

### HasCallbackId

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasCallbackId() bool`

HasCallbackId returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCreatedBy() GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCreatedByOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetCreatedBy(v GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetWebhook

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetWebhook() GetOrganizationWebhooksCallbacksStatus200ResponseWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetWebhookOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetWebhook(v GetOrganizationWebhooksCallbacksStatus200ResponseWebhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


