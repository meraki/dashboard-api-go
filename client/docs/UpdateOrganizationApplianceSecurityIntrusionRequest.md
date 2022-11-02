# UpdateOrganizationApplianceSecurityIntrusionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRules** | [**[]UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner**](UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner.md) | Sets a list of specific SNORT signatures to allow | 

## Methods

### NewUpdateOrganizationApplianceSecurityIntrusionRequest

`func NewUpdateOrganizationApplianceSecurityIntrusionRequest(allowedRules []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner, ) *UpdateOrganizationApplianceSecurityIntrusionRequest`

NewUpdateOrganizationApplianceSecurityIntrusionRequest instantiates a new UpdateOrganizationApplianceSecurityIntrusionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationApplianceSecurityIntrusionRequestWithDefaults

`func NewUpdateOrganizationApplianceSecurityIntrusionRequestWithDefaults() *UpdateOrganizationApplianceSecurityIntrusionRequest`

NewUpdateOrganizationApplianceSecurityIntrusionRequestWithDefaults instantiates a new UpdateOrganizationApplianceSecurityIntrusionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedRules

`func (o *UpdateOrganizationApplianceSecurityIntrusionRequest) GetAllowedRules() []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner`

GetAllowedRules returns the AllowedRules field if non-nil, zero value otherwise.

### GetAllowedRulesOk

`func (o *UpdateOrganizationApplianceSecurityIntrusionRequest) GetAllowedRulesOk() (*[]UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner, bool)`

GetAllowedRulesOk returns a tuple with the AllowedRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRules

`func (o *UpdateOrganizationApplianceSecurityIntrusionRequest) SetAllowedRules(v []UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner)`

SetAllowedRules sets AllowedRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


