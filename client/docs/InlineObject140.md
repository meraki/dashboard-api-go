# InlineObject140

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |     The traffic analysis mode for the network. Can be one of &#39;disabled&#39; (do not collect traffic types),     &#39;basic&#39; (collect generic traffic categories), or &#39;detailed&#39; (collect destination hostnames).  | [optional] 
**CustomPieChartItems** | Pointer to [**[]NetworksNetworkIdTrafficAnalysisCustomPieChartItems**](NetworksNetworkIdTrafficAnalysisCustomPieChartItems.md) | The list of items that make up the custom pie chart for traffic reporting. | [optional] 

## Methods

### NewInlineObject140

`func NewInlineObject140() *InlineObject140`

NewInlineObject140 instantiates a new InlineObject140 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject140WithDefaults

`func NewInlineObject140WithDefaults() *InlineObject140`

NewInlineObject140WithDefaults instantiates a new InlineObject140 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject140) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject140) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject140) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject140) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCustomPieChartItems

`func (o *InlineObject140) GetCustomPieChartItems() []NetworksNetworkIdTrafficAnalysisCustomPieChartItems`

GetCustomPieChartItems returns the CustomPieChartItems field if non-nil, zero value otherwise.

### GetCustomPieChartItemsOk

`func (o *InlineObject140) GetCustomPieChartItemsOk() (*[]NetworksNetworkIdTrafficAnalysisCustomPieChartItems, bool)`

GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPieChartItems

`func (o *InlineObject140) SetCustomPieChartItems(v []NetworksNetworkIdTrafficAnalysisCustomPieChartItems)`

SetCustomPieChartItems sets CustomPieChartItems field to given value.

### HasCustomPieChartItems

`func (o *InlineObject140) HasCustomPieChartItems() bool`

HasCustomPieChartItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


