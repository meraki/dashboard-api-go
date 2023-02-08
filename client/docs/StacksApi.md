# \StacksApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworkSwitchStack**](StacksApi.md#AddNetworkSwitchStack) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/add | Add a switch to a stack
[**CreateNetworkSwitchStack**](StacksApi.md#CreateNetworkSwitchStack) | **Post** /networks/{networkId}/switch/stacks | Create a stack
[**CreateNetworkSwitchStackRoutingInterface**](StacksApi.md#CreateNetworkSwitchStackRoutingInterface) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | Create a layer 3 interface for a switch stack
[**CreateNetworkSwitchStackRoutingStaticRoute**](StacksApi.md#CreateNetworkSwitchStackRoutingStaticRoute) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | Create a layer 3 static route for a switch stack
[**DeleteNetworkSwitchStack**](StacksApi.md#DeleteNetworkSwitchStack) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId} | Delete a stack
[**DeleteNetworkSwitchStackRoutingInterface**](StacksApi.md#DeleteNetworkSwitchStackRoutingInterface) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Delete a layer 3 interface from a switch stack
[**DeleteNetworkSwitchStackRoutingStaticRoute**](StacksApi.md#DeleteNetworkSwitchStackRoutingStaticRoute) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch stack
[**GetNetworkSwitchStack**](StacksApi.md#GetNetworkSwitchStack) | **Get** /networks/{networkId}/switch/stacks/{switchStackId} | Show a switch stack
[**GetNetworkSwitchStackRoutingInterface**](StacksApi.md#GetNetworkSwitchStackRoutingInterface) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Return a layer 3 interface from a switch stack
[**GetNetworkSwitchStackRoutingInterfaceDhcp**](StacksApi.md#GetNetworkSwitchStackRoutingInterfaceDhcp) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch stack
[**GetNetworkSwitchStackRoutingInterfaces**](StacksApi.md#GetNetworkSwitchStackRoutingInterfaces) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | List layer 3 interfaces for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoute**](StacksApi.md#GetNetworkSwitchStackRoutingStaticRoute) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoutes**](StacksApi.md#GetNetworkSwitchStackRoutingStaticRoutes) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | List layer 3 static routes for a switch stack
[**GetNetworkSwitchStacks**](StacksApi.md#GetNetworkSwitchStacks) | **Get** /networks/{networkId}/switch/stacks | List the switch stacks in a network
[**RemoveNetworkSwitchStack**](StacksApi.md#RemoveNetworkSwitchStack) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/remove | Remove a switch from a stack
[**UpdateNetworkSwitchStackRoutingInterface**](StacksApi.md#UpdateNetworkSwitchStackRoutingInterface) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch stack
[**UpdateNetworkSwitchStackRoutingInterfaceDhcp**](StacksApi.md#UpdateNetworkSwitchStackRoutingInterfaceDhcp) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch stack
[**UpdateNetworkSwitchStackRoutingStaticRoute**](StacksApi.md#UpdateNetworkSwitchStackRoutingStaticRoute) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch stack



## AddNetworkSwitchStack

> map[string]interface{} AddNetworkSwitchStack(ctx, networkId, switchStackId).AddNetworkSwitchStack(addNetworkSwitchStack).Execute()

Add a switch to a stack



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
    switchStackId := "switchStackId_example" // string | 
    addNetworkSwitchStack := *openapiclient.NewAddNetworkSwitchStackRequest("Serial_example") // AddNetworkSwitchStackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.AddNetworkSwitchStack(context.Background(), networkId, switchStackId).AddNetworkSwitchStack(addNetworkSwitchStack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.AddNetworkSwitchStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNetworkSwitchStack`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.AddNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addNetworkSwitchStack** | [**AddNetworkSwitchStackRequest**](AddNetworkSwitchStackRequest.md) |  | 

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


## CreateNetworkSwitchStack

> map[string]interface{} CreateNetworkSwitchStack(ctx, networkId).CreateNetworkSwitchStack(createNetworkSwitchStack).Execute()

Create a stack



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
    createNetworkSwitchStack := *openapiclient.NewCreateNetworkSwitchStackRequest("Name_example", []string{"Serials_example"}) // CreateNetworkSwitchStackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.CreateNetworkSwitchStack(context.Background(), networkId).CreateNetworkSwitchStack(createNetworkSwitchStack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.CreateNetworkSwitchStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchStack`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.CreateNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchStack** | [**CreateNetworkSwitchStackRequest**](CreateNetworkSwitchStackRequest.md) |  | 

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


## CreateNetworkSwitchStackRoutingInterface

> map[string]interface{} CreateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId).CreateNetworkSwitchStackRoutingInterface(createNetworkSwitchStackRoutingInterface).Execute()

Create a layer 3 interface for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    createNetworkSwitchStackRoutingInterface := *openapiclient.NewCreateNetworkSwitchStackRoutingInterfaceRequest("Name_example", int32(123)) // CreateNetworkSwitchStackRoutingInterfaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.CreateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId).CreateNetworkSwitchStackRoutingInterface(createNetworkSwitchStackRoutingInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.CreateNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchStackRoutingInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.CreateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchStackRoutingInterface** | [**CreateNetworkSwitchStackRoutingInterfaceRequest**](CreateNetworkSwitchStackRoutingInterfaceRequest.md) |  | 

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


## CreateNetworkSwitchStackRoutingStaticRoute

> map[string]interface{} CreateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId).CreateNetworkSwitchStackRoutingStaticRoute(createNetworkSwitchStackRoutingStaticRoute).Execute()

Create a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    createNetworkSwitchStackRoutingStaticRoute := *openapiclient.NewCreateDeviceSwitchRoutingStaticRouteRequest("Subnet_example", "NextHopIp_example") // CreateDeviceSwitchRoutingStaticRouteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).CreateNetworkSwitchStackRoutingStaticRoute(createNetworkSwitchStackRoutingStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.CreateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchStackRoutingStaticRoute`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.CreateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchStackRoutingStaticRoute** | [**CreateDeviceSwitchRoutingStaticRouteRequest**](CreateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

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


## DeleteNetworkSwitchStack

> DeleteNetworkSwitchStack(ctx, networkId, switchStackId).Execute()

Delete a stack



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
    switchStackId := "switchStackId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.DeleteNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.DeleteNetworkSwitchStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRequest struct via the builder pattern


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


## DeleteNetworkSwitchStackRoutingInterface

> DeleteNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Delete a layer 3 interface from a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    interfaceId := "interfaceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.DeleteNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.DeleteNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**interfaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


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


## DeleteNetworkSwitchStackRoutingStaticRoute

> DeleteNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Delete a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    staticRouteId := "staticRouteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.DeleteNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**staticRouteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


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


## GetNetworkSwitchStack

> map[string]interface{} GetNetworkSwitchStack(ctx, networkId, switchStackId).Execute()

Show a switch stack



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
    switchStackId := "switchStackId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.GetNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetNetworkSwitchStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStack`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRequest struct via the builder pattern


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


## GetNetworkSwitchStackRoutingInterface

> map[string]interface{} GetNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface from a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    interfaceId := "interfaceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.GetNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**interfaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


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


## GetNetworkSwitchStackRoutingInterfaceDhcp

> map[string]interface{} GetNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface DHCP configuration for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    interfaceId := "interfaceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingInterfaceDhcp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**interfaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


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


## GetNetworkSwitchStackRoutingInterfaces

> []map[string]interface{} GetNetworkSwitchStackRoutingInterfaces(ctx, networkId, switchStackId).Execute()

List layer 3 interfaces for a switch stack



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
    switchStackId := "switchStackId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.GetNetworkSwitchStackRoutingInterfaces(context.Background(), networkId, switchStackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetNetworkSwitchStackRoutingInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingInterfaces`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetNetworkSwitchStackRoutingInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfacesRequest struct via the builder pattern


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


## GetNetworkSwitchStackRoutingStaticRoute

> map[string]interface{} GetNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Return a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    staticRouteId := "staticRouteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingStaticRoute`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**staticRouteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


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


## GetNetworkSwitchStackRoutingStaticRoutes

> []map[string]interface{} GetNetworkSwitchStackRoutingStaticRoutes(ctx, networkId, switchStackId).Execute()

List layer 3 static routes for a switch stack



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
    switchStackId := "switchStackId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetNetworkSwitchStackRoutingStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingStaticRoutes`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetNetworkSwitchStackRoutingStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRoutesRequest struct via the builder pattern


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


## GetNetworkSwitchStacks

> []map[string]interface{} GetNetworkSwitchStacks(ctx, networkId).Execute()

List the switch stacks in a network



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
    resp, r, err := apiClient.StacksApi.GetNetworkSwitchStacks(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.GetNetworkSwitchStacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStacks`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.GetNetworkSwitchStacks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStacksRequest struct via the builder pattern


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


## RemoveNetworkSwitchStack

> map[string]interface{} RemoveNetworkSwitchStack(ctx, networkId, switchStackId).RemoveNetworkSwitchStack(removeNetworkSwitchStack).Execute()

Remove a switch from a stack



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
    switchStackId := "switchStackId_example" // string | 
    removeNetworkSwitchStack := *openapiclient.NewRemoveNetworkSwitchStackRequest("Serial_example") // RemoveNetworkSwitchStackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.RemoveNetworkSwitchStack(context.Background(), networkId, switchStackId).RemoveNetworkSwitchStack(removeNetworkSwitchStack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.RemoveNetworkSwitchStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveNetworkSwitchStack`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.RemoveNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeNetworkSwitchStack** | [**RemoveNetworkSwitchStackRequest**](RemoveNetworkSwitchStackRequest.md) |  | 

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


## UpdateNetworkSwitchStackRoutingInterface

> map[string]interface{} UpdateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterface(updateNetworkSwitchStackRoutingInterface).Execute()

Update a layer 3 interface for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    interfaceId := "interfaceId_example" // string | 
    updateNetworkSwitchStackRoutingInterface := *openapiclient.NewUpdateNetworkSwitchStackRoutingInterfaceRequest() // UpdateNetworkSwitchStackRoutingInterfaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.UpdateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterface(updateNetworkSwitchStackRoutingInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.UpdateNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStackRoutingInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.UpdateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**interfaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterface** | [**UpdateNetworkSwitchStackRoutingInterfaceRequest**](UpdateNetworkSwitchStackRoutingInterfaceRequest.md) |  | 

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


## UpdateNetworkSwitchStackRoutingInterfaceDhcp

> map[string]interface{} UpdateNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcp(updateNetworkSwitchStackRoutingInterfaceDhcp).Execute()

Update a layer 3 interface DHCP configuration for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    interfaceId := "interfaceId_example" // string | 
    updateNetworkSwitchStackRoutingInterfaceDhcp := *openapiclient.NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest() // UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcp(updateNetworkSwitchStackRoutingInterfaceDhcp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStackRoutingInterfaceDhcp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**interfaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterfaceDhcp** | [**UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest**](UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest.md) |  | 

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


## UpdateNetworkSwitchStackRoutingStaticRoute

> map[string]interface{} UpdateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).UpdateNetworkSwitchStackRoutingStaticRoute(updateNetworkSwitchStackRoutingStaticRoute).Execute()

Update a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | 
    staticRouteId := "staticRouteId_example" // string | 
    updateNetworkSwitchStackRoutingStaticRoute := *openapiclient.NewUpdateDeviceSwitchRoutingStaticRouteRequest() // UpdateDeviceSwitchRoutingStaticRouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StacksApi.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).UpdateNetworkSwitchStackRoutingStaticRoute(updateNetworkSwitchStackRoutingStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StacksApi.UpdateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStackRoutingStaticRoute`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StacksApi.UpdateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**switchStackId** | **string** |  | 
**staticRouteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingStaticRoute** | [**UpdateDeviceSwitchRoutingStaticRouteRequest**](UpdateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

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

