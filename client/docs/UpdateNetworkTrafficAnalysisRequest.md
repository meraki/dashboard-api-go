# UpdateNetworkTrafficAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |     The traffic analysis mode for the network. Can be one of &#39;disabled&#39; (do not collect traffic types),     &#39;basic&#39; (collect generic traffic categories), or &#39;detailed&#39; (collect destination hostnames).  | [optional] 
**CustomPieChartItems** | Pointer to [**[]UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner**](UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner.md) | The list of items that make up the custom pie chart for traffic reporting. | [optional] 

## Methods

### NewUpdateNetworkTrafficAnalysisRequest

`func NewUpdateNetworkTrafficAnalysisRequest() *UpdateNetworkTrafficAnalysisRequest`

NewUpdateNetworkTrafficAnalysisRequest instantiates a new UpdateNetworkTrafficAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkTrafficAnalysisRequestWithDefaults

`func NewUpdateNetworkTrafficAnalysisRequestWithDefaults() *UpdateNetworkTrafficAnalysisRequest`

NewUpdateNetworkTrafficAnalysisRequestWithDefaults instantiates a new UpdateNetworkTrafficAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *UpdateNetworkTrafficAnalysisRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateNetworkTrafficAnalysisRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateNetworkTrafficAnalysisRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateNetworkTrafficAnalysisRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCustomPieChartItems

`func (o *UpdateNetworkTrafficAnalysisRequest) GetCustomPieChartItems() []UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner`

GetCustomPieChartItems returns the CustomPieChartItems field if non-nil, zero value otherwise.

### GetCustomPieChartItemsOk

`func (o *UpdateNetworkTrafficAnalysisRequest) GetCustomPieChartItemsOk() (*[]UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner, bool)`

GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPieChartItems

`func (o *UpdateNetworkTrafficAnalysisRequest) SetCustomPieChartItems(v []UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner)`

SetCustomPieChartItems sets CustomPieChartItems field to given value.

### HasCustomPieChartItems

`func (o *UpdateNetworkTrafficAnalysisRequest) HasCustomPieChartItems() bool`

HasCustomPieChartItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


