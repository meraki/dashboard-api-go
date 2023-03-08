# OrganizationsOrganizationIdAdminsNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The network ID | 
**Access** | **string** | The privilege of the dashboard administrator on the network. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;guest-ambassador&#39; or &#39;monitor-only&#39; | 

## Methods

### NewOrganizationsOrganizationIdAdminsNetworks

`func NewOrganizationsOrganizationIdAdminsNetworks(id string, access string, ) *OrganizationsOrganizationIdAdminsNetworks`

NewOrganizationsOrganizationIdAdminsNetworks instantiates a new OrganizationsOrganizationIdAdminsNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAdminsNetworksWithDefaults

`func NewOrganizationsOrganizationIdAdminsNetworksWithDefaults() *OrganizationsOrganizationIdAdminsNetworks`

NewOrganizationsOrganizationIdAdminsNetworksWithDefaults instantiates a new OrganizationsOrganizationIdAdminsNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationsOrganizationIdAdminsNetworks) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationsOrganizationIdAdminsNetworks) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationsOrganizationIdAdminsNetworks) SetId(v string)`

SetId sets Id field to given value.


### GetAccess

`func (o *OrganizationsOrganizationIdAdminsNetworks) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *OrganizationsOrganizationIdAdminsNetworks) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *OrganizationsOrganizationIdAdminsNetworks) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


