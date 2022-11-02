# \QosRulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchQosRule**](QosRulesApi.md#CreateNetworkSwitchQosRule) | **Post** /networks/{networkId}/switch/qosRules | Add a quality of service rule
[**DeleteNetworkSwitchQosRule**](QosRulesApi.md#DeleteNetworkSwitchQosRule) | **Delete** /networks/{networkId}/switch/qosRules/{qosRuleId} | Delete a quality of service rule
[**GetNetworkSwitchQosRule**](QosRulesApi.md#GetNetworkSwitchQosRule) | **Get** /networks/{networkId}/switch/qosRules/{qosRuleId} | Return a quality of service rule
[**GetNetworkSwitchQosRules**](QosRulesApi.md#GetNetworkSwitchQosRules) | **Get** /networks/{networkId}/switch/qosRules | List quality of service rules
[**GetNetworkSwitchQosRulesOrder**](QosRulesApi.md#GetNetworkSwitchQosRulesOrder) | **Get** /networks/{networkId}/switch/qosRules/order | Return the quality of service rule IDs by order in which they will be processed by the switch
[**UpdateNetworkSwitchQosRule**](QosRulesApi.md#UpdateNetworkSwitchQosRule) | **Put** /networks/{networkId}/switch/qosRules/{qosRuleId} | Update a quality of service rule
[**UpdateNetworkSwitchQosRulesOrder**](QosRulesApi.md#UpdateNetworkSwitchQosRulesOrder) | **Put** /networks/{networkId}/switch/qosRules/order | Update the order in which the rules should be processed by the switch



## CreateNetworkSwitchQosRule

> map[string]interface{} CreateNetworkSwitchQosRule(ctx, networkId).CreateNetworkSwitchQosRule(createNetworkSwitchQosRule).Execute()

Add a quality of service rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 
    createNetworkSwitchQosRule := *openapiclient.NewCreateNetworkSwitchQosRuleRequest(int32(123)) // CreateNetworkSwitchQosRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QosRulesApi.CreateNetworkSwitchQosRule(context.Background(), networkId).CreateNetworkSwitchQosRule(createNetworkSwitchQosRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QosRulesApi.CreateNetworkSwitchQosRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchQosRule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QosRulesApi.CreateNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchQosRule** | [**CreateNetworkSwitchQosRuleRequest**](CreateNetworkSwitchQosRuleRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchQosRule

> DeleteNetworkSwitchQosRule(ctx, networkId, qosRuleId).Execute()

Delete a quality of service rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 
    qosRuleId := "qosRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QosRulesApi.DeleteNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QosRulesApi.DeleteNetworkSwitchQosRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**qosRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRule

> map[string]interface{} GetNetworkSwitchQosRule(ctx, networkId, qosRuleId).Execute()

Return a quality of service rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 
    qosRuleId := "qosRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QosRulesApi.GetNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QosRulesApi.GetNetworkSwitchQosRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchQosRule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QosRulesApi.GetNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**qosRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRules

> []map[string]interface{} GetNetworkSwitchQosRules(ctx, networkId).Execute()

List quality of service rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QosRulesApi.GetNetworkSwitchQosRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QosRulesApi.GetNetworkSwitchQosRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchQosRules`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QosRulesApi.GetNetworkSwitchQosRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRulesOrder

> map[string]interface{} GetNetworkSwitchQosRulesOrder(ctx, networkId).Execute()

Return the quality of service rule IDs by order in which they will be processed by the switch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QosRulesApi.GetNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QosRulesApi.GetNetworkSwitchQosRulesOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchQosRulesOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QosRulesApi.GetNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchQosRule

> map[string]interface{} UpdateNetworkSwitchQosRule(ctx, networkId, qosRuleId).UpdateNetworkSwitchQosRule(updateNetworkSwitchQosRule).Execute()

Update a quality of service rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 
    qosRuleId := "qosRuleId_example" // string | 
    updateNetworkSwitchQosRule := *openapiclient.NewUpdateNetworkSwitchQosRuleRequest() // UpdateNetworkSwitchQosRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QosRulesApi.UpdateNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).UpdateNetworkSwitchQosRule(updateNetworkSwitchQosRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QosRulesApi.UpdateNetworkSwitchQosRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchQosRule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QosRulesApi.UpdateNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**qosRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchQosRule** | [**UpdateNetworkSwitchQosRuleRequest**](UpdateNetworkSwitchQosRuleRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchQosRulesOrder

> map[string]interface{} UpdateNetworkSwitchQosRulesOrder(ctx, networkId).UpdateNetworkSwitchQosRulesOrder(updateNetworkSwitchQosRulesOrder).Execute()

Update the order in which the rules should be processed by the switch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 
    updateNetworkSwitchQosRulesOrder := *openapiclient.NewUpdateNetworkSwitchQosRulesOrderRequest([]string{"RuleIds_example"}) // UpdateNetworkSwitchQosRulesOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QosRulesApi.UpdateNetworkSwitchQosRulesOrder(context.Background(), networkId).UpdateNetworkSwitchQosRulesOrder(updateNetworkSwitchQosRulesOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QosRulesApi.UpdateNetworkSwitchQosRulesOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchQosRulesOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QosRulesApi.UpdateNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchQosRulesOrder** | [**UpdateNetworkSwitchQosRulesOrderRequest**](UpdateNetworkSwitchQosRulesOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

