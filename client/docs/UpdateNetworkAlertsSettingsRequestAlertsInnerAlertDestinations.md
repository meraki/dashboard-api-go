# UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | A list of emails that will receive information about the alert | [optional] 
**AllAdmins** | Pointer to **bool** | If true, then all network admins will receive emails for this alert | [optional] 
**Snmp** | Pointer to **bool** | If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network | [optional] 
**HttpServerIds** | Pointer to **[]string** | A list of HTTP server IDs to send a Webhook to for this alert | [optional] 

## Methods

### NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations

`func NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations() *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations`

NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations instantiates a new UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinationsWithDefaults

`func NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinationsWithDefaults() *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations`

NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinationsWithDefaults instantiates a new UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetAllAdmins

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetAllAdmins() bool`

GetAllAdmins returns the AllAdmins field if non-nil, zero value otherwise.

### GetAllAdminsOk

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetAllAdminsOk() (*bool, bool)`

GetAllAdminsOk returns a tuple with the AllAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAdmins

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetAllAdmins(v bool)`

SetAllAdmins sets AllAdmins field to given value.

### HasAllAdmins

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasAllAdmins() bool`

HasAllAdmins returns a boolean if a field has been set.

### GetSnmp

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetSnmp() bool`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetSnmpOk() (*bool, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetSnmp(v bool)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetHttpServerIds

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetHttpServerIds() []string`

GetHttpServerIds returns the HttpServerIds field if non-nil, zero value otherwise.

### GetHttpServerIdsOk

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetHttpServerIdsOk() (*[]string, bool)`

GetHttpServerIdsOk returns a tuple with the HttpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerIds

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetHttpServerIds(v []string)`

SetHttpServerIds sets HttpServerIds field to given value.

### HasHttpServerIds

`func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasHttpServerIds() bool`

HasHttpServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


