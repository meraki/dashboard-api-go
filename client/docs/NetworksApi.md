# \NetworksApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindNetwork**](NetworksApi.md#BindNetwork) | **Post** /networks/{networkId}/bind | Bind a network to a template.
[**ClaimNetworkDevices**](NetworksApi.md#ClaimNetworkDevices) | **Post** /networks/{networkId}/devices/claim | Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requsts against that device to succeed)
[**CombineOrganizationNetworks**](NetworksApi.md#CombineOrganizationNetworks) | **Post** /organizations/{organizationId}/networks/combine | Combine multiple networks into a single network
[**CreateNetworkFirmwareUpgradesRollback**](NetworksApi.md#CreateNetworkFirmwareUpgradesRollback) | **Post** /networks/{networkId}/firmwareUpgrades/rollbacks | Rollback a Firmware Upgrade For A Network
[**CreateNetworkFirmwareUpgradesStagedEvent**](NetworksApi.md#CreateNetworkFirmwareUpgradesStagedEvent) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events | Create a Staged Upgrade Event for a network
[**CreateNetworkFirmwareUpgradesStagedGroup**](NetworksApi.md#CreateNetworkFirmwareUpgradesStagedGroup) | **Post** /networks/{networkId}/firmwareUpgrades/staged/groups | Create a Staged Upgrade Group for a network
[**CreateNetworkFloorPlan**](NetworksApi.md#CreateNetworkFloorPlan) | **Post** /networks/{networkId}/floorPlans | Upload a floor plan
[**CreateNetworkGroupPolicy**](NetworksApi.md#CreateNetworkGroupPolicy) | **Post** /networks/{networkId}/groupPolicies | Create a group policy
[**CreateNetworkMerakiAuthUser**](NetworksApi.md#CreateNetworkMerakiAuthUser) | **Post** /networks/{networkId}/merakiAuthUsers | Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)
[**CreateNetworkMqttBroker**](NetworksApi.md#CreateNetworkMqttBroker) | **Post** /networks/{networkId}/mqttBrokers | Add an MQTT broker
[**CreateNetworkPiiRequest**](NetworksApi.md#CreateNetworkPiiRequest) | **Post** /networks/{networkId}/pii/requests | Submit a new delete or restrict processing PII request
[**CreateNetworkWebhooksHttpServer**](NetworksApi.md#CreateNetworkWebhooksHttpServer) | **Post** /networks/{networkId}/webhooks/httpServers | Add an HTTP server to a network
[**CreateNetworkWebhooksPayloadTemplate**](NetworksApi.md#CreateNetworkWebhooksPayloadTemplate) | **Post** /networks/{networkId}/webhooks/payloadTemplates | Create a webhook payload template for a network
[**CreateNetworkWebhooksWebhookTest**](NetworksApi.md#CreateNetworkWebhooksWebhookTest) | **Post** /networks/{networkId}/webhooks/webhookTests | Send a test webhook for a network
[**CreateOrganizationNetwork**](NetworksApi.md#CreateOrganizationNetwork) | **Post** /organizations/{organizationId}/networks | Create a network
[**DeferNetworkFirmwareUpgradesStagedEvents**](NetworksApi.md#DeferNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/defer | Postpone by 1 week all pending staged upgrade stages for a network
[**DeleteNetwork**](NetworksApi.md#DeleteNetwork) | **Delete** /networks/{networkId} | Delete a network
[**DeleteNetworkFirmwareUpgradesStagedGroup**](NetworksApi.md#DeleteNetworkFirmwareUpgradesStagedGroup) | **Delete** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Delete a Staged Upgrade Group
[**DeleteNetworkFloorPlan**](NetworksApi.md#DeleteNetworkFloorPlan) | **Delete** /networks/{networkId}/floorPlans/{floorPlanId} | Destroy a floor plan
[**DeleteNetworkGroupPolicy**](NetworksApi.md#DeleteNetworkGroupPolicy) | **Delete** /networks/{networkId}/groupPolicies/{groupPolicyId} | Delete a group policy
[**DeleteNetworkMerakiAuthUser**](NetworksApi.md#DeleteNetworkMerakiAuthUser) | **Delete** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Deauthorize a user
[**DeleteNetworkMqttBroker**](NetworksApi.md#DeleteNetworkMqttBroker) | **Delete** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Delete an MQTT broker
[**DeleteNetworkPiiRequest**](NetworksApi.md#DeleteNetworkPiiRequest) | **Delete** /networks/{networkId}/pii/requests/{requestId} | Delete a restrict processing PII request
[**DeleteNetworkWebhooksHttpServer**](NetworksApi.md#DeleteNetworkWebhooksHttpServer) | **Delete** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Delete an HTTP server from a network
[**DeleteNetworkWebhooksPayloadTemplate**](NetworksApi.md#DeleteNetworkWebhooksPayloadTemplate) | **Delete** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Destroy a webhook payload template for a network
[**GetNetwork**](NetworksApi.md#GetNetwork) | **Get** /networks/{networkId} | Return a network
[**GetNetworkAlertsHistory**](NetworksApi.md#GetNetworkAlertsHistory) | **Get** /networks/{networkId}/alerts/history | Return the alert history for this network
[**GetNetworkAlertsSettings**](NetworksApi.md#GetNetworkAlertsSettings) | **Get** /networks/{networkId}/alerts/settings | Return the alert configuration for this network
[**GetNetworkBluetoothClient**](NetworksApi.md#GetNetworkBluetoothClient) | **Get** /networks/{networkId}/bluetoothClients/{bluetoothClientId} | Return a Bluetooth client
[**GetNetworkBluetoothClients**](NetworksApi.md#GetNetworkBluetoothClients) | **Get** /networks/{networkId}/bluetoothClients | List the Bluetooth clients seen by APs in this network
[**GetNetworkClient**](NetworksApi.md#GetNetworkClient) | **Get** /networks/{networkId}/clients/{clientId} | Return the client associated with the given identifier
[**GetNetworkClientPolicy**](NetworksApi.md#GetNetworkClientPolicy) | **Get** /networks/{networkId}/clients/{clientId}/policy | Return the policy assigned to a client on the network
[**GetNetworkClientSplashAuthorizationStatus**](NetworksApi.md#GetNetworkClientSplashAuthorizationStatus) | **Get** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Return the splash authorization for a client, for each SSID they&#39;ve associated with through splash
[**GetNetworkClientTrafficHistory**](NetworksApi.md#GetNetworkClientTrafficHistory) | **Get** /networks/{networkId}/clients/{clientId}/trafficHistory | Return the client&#39;s network traffic data over time
[**GetNetworkClientUsageHistory**](NetworksApi.md#GetNetworkClientUsageHistory) | **Get** /networks/{networkId}/clients/{clientId}/usageHistory | Return the client&#39;s daily usage history
[**GetNetworkClients**](NetworksApi.md#GetNetworkClients) | **Get** /networks/{networkId}/clients | List the clients that have used this network in the timespan
[**GetNetworkClientsApplicationUsage**](NetworksApi.md#GetNetworkClientsApplicationUsage) | **Get** /networks/{networkId}/clients/applicationUsage | Return the application usage data for clients
[**GetNetworkClientsBandwidthUsageHistory**](NetworksApi.md#GetNetworkClientsBandwidthUsageHistory) | **Get** /networks/{networkId}/clients/bandwidthUsageHistory | Returns a timeseries of total traffic consumption rates for all clients on a network within a given timespan, in megabits per second.
[**GetNetworkClientsOverview**](NetworksApi.md#GetNetworkClientsOverview) | **Get** /networks/{networkId}/clients/overview | Return overview statistics for network clients
[**GetNetworkClientsUsageHistories**](NetworksApi.md#GetNetworkClientsUsageHistories) | **Get** /networks/{networkId}/clients/usageHistories | Return the usage histories for clients
[**GetNetworkDevices**](NetworksApi.md#GetNetworkDevices) | **Get** /networks/{networkId}/devices | List the devices in a network
[**GetNetworkEvents**](NetworksApi.md#GetNetworkEvents) | **Get** /networks/{networkId}/events | List the events for the network
[**GetNetworkEventsEventTypes**](NetworksApi.md#GetNetworkEventsEventTypes) | **Get** /networks/{networkId}/events/eventTypes | List the event type to human-readable description
[**GetNetworkFirmwareUpgrades**](NetworksApi.md#GetNetworkFirmwareUpgrades) | **Get** /networks/{networkId}/firmwareUpgrades | Get firmware upgrade information for a network
[**GetNetworkFirmwareUpgradesStagedEvents**](NetworksApi.md#GetNetworkFirmwareUpgradesStagedEvents) | **Get** /networks/{networkId}/firmwareUpgrades/staged/events | Get the Staged Upgrade Event from a network
[**GetNetworkFirmwareUpgradesStagedGroup**](NetworksApi.md#GetNetworkFirmwareUpgradesStagedGroup) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Get a Staged Upgrade Group from a network
[**GetNetworkFirmwareUpgradesStagedGroups**](NetworksApi.md#GetNetworkFirmwareUpgradesStagedGroups) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups | List of Staged Upgrade Groups in a network
[**GetNetworkFirmwareUpgradesStagedStages**](NetworksApi.md#GetNetworkFirmwareUpgradesStagedStages) | **Get** /networks/{networkId}/firmwareUpgrades/staged/stages | Order of Staged Upgrade Groups in a network
[**GetNetworkFloorPlan**](NetworksApi.md#GetNetworkFloorPlan) | **Get** /networks/{networkId}/floorPlans/{floorPlanId} | Find a floor plan by ID
[**GetNetworkFloorPlans**](NetworksApi.md#GetNetworkFloorPlans) | **Get** /networks/{networkId}/floorPlans | List the floor plans that belong to your network
[**GetNetworkGroupPolicies**](NetworksApi.md#GetNetworkGroupPolicies) | **Get** /networks/{networkId}/groupPolicies | List the group policies in a network
[**GetNetworkGroupPolicy**](NetworksApi.md#GetNetworkGroupPolicy) | **Get** /networks/{networkId}/groupPolicies/{groupPolicyId} | Display a group policy
[**GetNetworkHealthAlerts**](NetworksApi.md#GetNetworkHealthAlerts) | **Get** /networks/{networkId}/health/alerts | Return all global alerts on this network
[**GetNetworkMerakiAuthUser**](NetworksApi.md#GetNetworkMerakiAuthUser) | **Get** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Return the Meraki Auth splash guest, RADIUS, or client VPN user
[**GetNetworkMerakiAuthUsers**](NetworksApi.md#GetNetworkMerakiAuthUsers) | **Get** /networks/{networkId}/merakiAuthUsers | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)
[**GetNetworkMqttBroker**](NetworksApi.md#GetNetworkMqttBroker) | **Get** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Return an MQTT broker
[**GetNetworkMqttBrokers**](NetworksApi.md#GetNetworkMqttBrokers) | **Get** /networks/{networkId}/mqttBrokers | List the MQTT brokers for this network
[**GetNetworkNetflow**](NetworksApi.md#GetNetworkNetflow) | **Get** /networks/{networkId}/netflow | Return the NetFlow traffic reporting settings for a network
[**GetNetworkNetworkHealthChannelUtilization**](NetworksApi.md#GetNetworkNetworkHealthChannelUtilization) | **Get** /networks/{networkId}/networkHealth/channelUtilization | Get the channel utilization over each radio for all APs in a network.
[**GetNetworkPiiPiiKeys**](NetworksApi.md#GetNetworkPiiPiiKeys) | **Get** /networks/{networkId}/pii/piiKeys | List the keys required to access Personally Identifiable Information (PII) for a given identifier
[**GetNetworkPiiRequest**](NetworksApi.md#GetNetworkPiiRequest) | **Get** /networks/{networkId}/pii/requests/{requestId} | Return a PII request
[**GetNetworkPiiRequests**](NetworksApi.md#GetNetworkPiiRequests) | **Get** /networks/{networkId}/pii/requests | List the PII requests for this network or organization
[**GetNetworkPiiSmDevicesForKey**](NetworksApi.md#GetNetworkPiiSmDevicesForKey) | **Get** /networks/{networkId}/pii/smDevicesForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier
[**GetNetworkPiiSmOwnersForKey**](NetworksApi.md#GetNetworkPiiSmOwnersForKey) | **Get** /networks/{networkId}/pii/smOwnersForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier
[**GetNetworkPoliciesByClient**](NetworksApi.md#GetNetworkPoliciesByClient) | **Get** /networks/{networkId}/policies/byClient | Get policies for all clients with policies
[**GetNetworkSettings**](NetworksApi.md#GetNetworkSettings) | **Get** /networks/{networkId}/settings | Return the settings for a network
[**GetNetworkSnmp**](NetworksApi.md#GetNetworkSnmp) | **Get** /networks/{networkId}/snmp | Return the SNMP settings for a network
[**GetNetworkSplashLoginAttempts**](NetworksApi.md#GetNetworkSplashLoginAttempts) | **Get** /networks/{networkId}/splashLoginAttempts | List the splash login attempts for a network
[**GetNetworkSyslogServers**](NetworksApi.md#GetNetworkSyslogServers) | **Get** /networks/{networkId}/syslogServers | List the syslog servers for a network
[**GetNetworkTopologyLinkLayer**](NetworksApi.md#GetNetworkTopologyLinkLayer) | **Get** /networks/{networkId}/topology/linkLayer | List the LLDP and CDP information for all discovered devices and connections in a network.
[**GetNetworkTraffic**](NetworksApi.md#GetNetworkTraffic) | **Get** /networks/{networkId}/traffic | Return the traffic analysis data for this network
[**GetNetworkTrafficAnalysis**](NetworksApi.md#GetNetworkTrafficAnalysis) | **Get** /networks/{networkId}/trafficAnalysis | Return the traffic analysis settings for a network
[**GetNetworkTrafficShapingApplicationCategories**](NetworksApi.md#GetNetworkTrafficShapingApplicationCategories) | **Get** /networks/{networkId}/trafficShaping/applicationCategories | Returns the application categories for traffic shaping rules.
[**GetNetworkTrafficShapingDscpTaggingOptions**](NetworksApi.md#GetNetworkTrafficShapingDscpTaggingOptions) | **Get** /networks/{networkId}/trafficShaping/dscpTaggingOptions | Returns the available DSCP tagging options for your traffic shaping rules.
[**GetNetworkWebhooksHttpServer**](NetworksApi.md#GetNetworkWebhooksHttpServer) | **Get** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Return an HTTP server for a network
[**GetNetworkWebhooksHttpServers**](NetworksApi.md#GetNetworkWebhooksHttpServers) | **Get** /networks/{networkId}/webhooks/httpServers | List the HTTP servers for a network
[**GetNetworkWebhooksPayloadTemplate**](NetworksApi.md#GetNetworkWebhooksPayloadTemplate) | **Get** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Get the webhook payload template for a network
[**GetNetworkWebhooksPayloadTemplates**](NetworksApi.md#GetNetworkWebhooksPayloadTemplates) | **Get** /networks/{networkId}/webhooks/payloadTemplates | List the webhook payload templates for a network
[**GetNetworkWebhooksWebhookTest**](NetworksApi.md#GetNetworkWebhooksWebhookTest) | **Get** /networks/{networkId}/webhooks/webhookTests/{webhookTestId} | Return the status of a webhook test for a network
[**GetOrganizationInventoryOnboardingCloudMonitoringNetworks**](NetworksApi.md#GetOrganizationInventoryOnboardingCloudMonitoringNetworks) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks | Returns list of networks eligible for adding cloud monitored device
[**GetOrganizationNetworks**](NetworksApi.md#GetOrganizationNetworks) | **Get** /organizations/{organizationId}/networks | List the networks that the user has privileges on in an organization
[**ProvisionNetworkClients**](NetworksApi.md#ProvisionNetworkClients) | **Post** /networks/{networkId}/clients/provision | Provisions a client with a name and policy
[**RemoveNetworkDevices**](NetworksApi.md#RemoveNetworkDevices) | **Post** /networks/{networkId}/devices/remove | Remove a single device
[**RollbacksNetworkFirmwareUpgradesStagedEvents**](NetworksApi.md#RollbacksNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/rollbacks | Rollback a Staged Upgrade Event for a network
[**SplitNetwork**](NetworksApi.md#SplitNetwork) | **Post** /networks/{networkId}/split | Split a combined network into individual networks for each type of device
[**UnbindNetwork**](NetworksApi.md#UnbindNetwork) | **Post** /networks/{networkId}/unbind | Unbind a network from a template.
[**UpdateNetwork**](NetworksApi.md#UpdateNetwork) | **Put** /networks/{networkId} | Update a network
[**UpdateNetworkAlertsSettings**](NetworksApi.md#UpdateNetworkAlertsSettings) | **Put** /networks/{networkId}/alerts/settings | Update the alert configuration for this network
[**UpdateNetworkClientPolicy**](NetworksApi.md#UpdateNetworkClientPolicy) | **Put** /networks/{networkId}/clients/{clientId}/policy | Update the policy assigned to a client on the network
[**UpdateNetworkClientSplashAuthorizationStatus**](NetworksApi.md#UpdateNetworkClientSplashAuthorizationStatus) | **Put** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Update a client&#39;s splash authorization
[**UpdateNetworkFirmwareUpgrades**](NetworksApi.md#UpdateNetworkFirmwareUpgrades) | **Put** /networks/{networkId}/firmwareUpgrades | Update firmware upgrade information for a network
[**UpdateNetworkFirmwareUpgradesStagedEvents**](NetworksApi.md#UpdateNetworkFirmwareUpgradesStagedEvents) | **Put** /networks/{networkId}/firmwareUpgrades/staged/events | Update the Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgradesStagedGroup**](NetworksApi.md#UpdateNetworkFirmwareUpgradesStagedGroup) | **Put** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Update a Staged Upgrade Group for a network
[**UpdateNetworkFirmwareUpgradesStagedStages**](NetworksApi.md#UpdateNetworkFirmwareUpgradesStagedStages) | **Put** /networks/{networkId}/firmwareUpgrades/staged/stages | Assign Staged Upgrade Group order in the sequence.
[**UpdateNetworkFloorPlan**](NetworksApi.md#UpdateNetworkFloorPlan) | **Put** /networks/{networkId}/floorPlans/{floorPlanId} | Update a floor plan&#39;s geolocation and other meta data
[**UpdateNetworkGroupPolicy**](NetworksApi.md#UpdateNetworkGroupPolicy) | **Put** /networks/{networkId}/groupPolicies/{groupPolicyId} | Update a group policy
[**UpdateNetworkMerakiAuthUser**](NetworksApi.md#UpdateNetworkMerakiAuthUser) | **Put** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)
[**UpdateNetworkMqttBroker**](NetworksApi.md#UpdateNetworkMqttBroker) | **Put** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Update an MQTT broker
[**UpdateNetworkNetflow**](NetworksApi.md#UpdateNetworkNetflow) | **Put** /networks/{networkId}/netflow | Update the NetFlow traffic reporting settings for a network
[**UpdateNetworkSettings**](NetworksApi.md#UpdateNetworkSettings) | **Put** /networks/{networkId}/settings | Update the settings for a network
[**UpdateNetworkSnmp**](NetworksApi.md#UpdateNetworkSnmp) | **Put** /networks/{networkId}/snmp | Update the SNMP settings for a network
[**UpdateNetworkSyslogServers**](NetworksApi.md#UpdateNetworkSyslogServers) | **Put** /networks/{networkId}/syslogServers | Update the syslog servers for a network
[**UpdateNetworkTrafficAnalysis**](NetworksApi.md#UpdateNetworkTrafficAnalysis) | **Put** /networks/{networkId}/trafficAnalysis | Update the traffic analysis settings for a network
[**UpdateNetworkWebhooksHttpServer**](NetworksApi.md#UpdateNetworkWebhooksHttpServer) | **Put** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Update an HTTP server
[**UpdateNetworkWebhooksPayloadTemplate**](NetworksApi.md#UpdateNetworkWebhooksPayloadTemplate) | **Put** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Update a webhook payload template for a network
[**VmxNetworkDevicesClaim**](NetworksApi.md#VmxNetworkDevicesClaim) | **Post** /networks/{networkId}/devices/claim/vmx | Claim a vMX into a network



## BindNetwork

> map[string]interface{} BindNetwork(ctx, networkId).BindNetworkRequest(bindNetworkRequest).Execute()

Bind a network to a template.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    bindNetworkRequest := *openapiclient.NewBindNetworkRequest("ConfigTemplateId_example") // BindNetworkRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.BindNetwork(context.Background(), networkId).BindNetworkRequest(bindNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.BindNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindNetwork`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.BindNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindNetworkRequest** | [**BindNetworkRequest**](BindNetworkRequest.md) |  | 

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


## ClaimNetworkDevices

> ClaimNetworkDevices(ctx, networkId).ClaimNetworkDevicesRequest(claimNetworkDevicesRequest).Execute()

Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requsts against that device to succeed)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    claimNetworkDevicesRequest := *openapiclient.NewClaimNetworkDevicesRequest([]string{"Serials_example"}) // ClaimNetworkDevicesRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.ClaimNetworkDevices(context.Background(), networkId).ClaimNetworkDevicesRequest(claimNetworkDevicesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ClaimNetworkDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimNetworkDevicesRequest** | [**ClaimNetworkDevicesRequest**](ClaimNetworkDevicesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CombineOrganizationNetworks

> CombineOrganizationNetworks200Response CombineOrganizationNetworks(ctx, organizationId).CombineOrganizationNetworksRequest(combineOrganizationNetworksRequest).Execute()

Combine multiple networks into a single network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    combineOrganizationNetworksRequest := *openapiclient.NewCombineOrganizationNetworksRequest("Name_example", []string{"NetworkIds_example"}) // CombineOrganizationNetworksRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CombineOrganizationNetworks(context.Background(), organizationId).CombineOrganizationNetworksRequest(combineOrganizationNetworksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CombineOrganizationNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CombineOrganizationNetworks`: CombineOrganizationNetworks200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CombineOrganizationNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCombineOrganizationNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **combineOrganizationNetworksRequest** | [**CombineOrganizationNetworksRequest**](CombineOrganizationNetworksRequest.md) |  | 

### Return type

[**CombineOrganizationNetworks200Response**](CombineOrganizationNetworks200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesRollback

> CreateNetworkFirmwareUpgradesRollback200Response CreateNetworkFirmwareUpgradesRollback(ctx, networkId).CreateNetworkFirmwareUpgradesRollbackRequest(createNetworkFirmwareUpgradesRollbackRequest).Execute()

Rollback a Firmware Upgrade For A Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkFirmwareUpgradesRollbackRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesRollbackRequest([]openapiclient.CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner{*openapiclient.NewCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner("Category_example", "Comment_example")}) // CreateNetworkFirmwareUpgradesRollbackRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkFirmwareUpgradesRollback(context.Background(), networkId).CreateNetworkFirmwareUpgradesRollbackRequest(createNetworkFirmwareUpgradesRollbackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkFirmwareUpgradesRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesRollback`: CreateNetworkFirmwareUpgradesRollback200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkFirmwareUpgradesRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesRollbackRequest** | [**CreateNetworkFirmwareUpgradesRollbackRequest**](CreateNetworkFirmwareUpgradesRollbackRequest.md) |  | 

### Return type

[**CreateNetworkFirmwareUpgradesRollback200Response**](CreateNetworkFirmwareUpgradesRollback200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesStagedEvent

> GetNetworkFirmwareUpgradesStagedEvents200Response CreateNetworkFirmwareUpgradesStagedEvent(ctx, networkId).CreateNetworkFirmwareUpgradesStagedEventRequest(createNetworkFirmwareUpgradesStagedEventRequest).Execute()

Create a Staged Upgrade Event for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkFirmwareUpgradesStagedEventRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedEventRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // CreateNetworkFirmwareUpgradesStagedEventRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkFirmwareUpgradesStagedEvent(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedEventRequest(createNetworkFirmwareUpgradesStagedEventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkFirmwareUpgradesStagedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesStagedEvent`: GetNetworkFirmwareUpgradesStagedEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkFirmwareUpgradesStagedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedEventRequest** | [**CreateNetworkFirmwareUpgradesStagedEventRequest**](CreateNetworkFirmwareUpgradesStagedEventRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesStagedGroup

> map[string]interface{} CreateNetworkFirmwareUpgradesStagedGroup(ctx, networkId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()

Create a Staged Upgrade Group for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkFirmwareUpgradesStagedGroupRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedGroupRequest("Name_example", false) // CreateNetworkFirmwareUpgradesStagedGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedGroupRequest** | [**CreateNetworkFirmwareUpgradesStagedGroupRequest**](CreateNetworkFirmwareUpgradesStagedGroupRequest.md) |  | 

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


## CreateNetworkFloorPlan

> map[string]interface{} CreateNetworkFloorPlan(ctx, networkId).CreateNetworkFloorPlanRequest(createNetworkFloorPlanRequest).Execute()

Upload a floor plan



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkFloorPlanRequest := *openapiclient.NewCreateNetworkFloorPlanRequest("Name_example", string(123)) // CreateNetworkFloorPlanRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkFloorPlan(context.Background(), networkId).CreateNetworkFloorPlanRequest(createNetworkFloorPlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFloorPlan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFloorPlanRequest** | [**CreateNetworkFloorPlanRequest**](CreateNetworkFloorPlanRequest.md) |  | 

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


## CreateNetworkGroupPolicy

> map[string]interface{} CreateNetworkGroupPolicy(ctx, networkId).CreateNetworkGroupPolicyRequest(createNetworkGroupPolicyRequest).Execute()

Create a group policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkGroupPolicyRequest := *openapiclient.NewCreateNetworkGroupPolicyRequest("Name_example") // CreateNetworkGroupPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkGroupPolicy(context.Background(), networkId).CreateNetworkGroupPolicyRequest(createNetworkGroupPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkGroupPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkGroupPolicyRequest** | [**CreateNetworkGroupPolicyRequest**](CreateNetworkGroupPolicyRequest.md) |  | 

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


## CreateNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner CreateNetworkMerakiAuthUser(ctx, networkId).CreateNetworkMerakiAuthUserRequest(createNetworkMerakiAuthUserRequest).Execute()

Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkMerakiAuthUserRequest := *openapiclient.NewCreateNetworkMerakiAuthUserRequest("Email_example", []openapiclient.CreateNetworkMerakiAuthUserRequestAuthorizationsInner{*openapiclient.NewCreateNetworkMerakiAuthUserRequestAuthorizationsInner()}) // CreateNetworkMerakiAuthUserRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkMerakiAuthUser(context.Background(), networkId).CreateNetworkMerakiAuthUserRequest(createNetworkMerakiAuthUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMerakiAuthUserRequest** | [**CreateNetworkMerakiAuthUserRequest**](CreateNetworkMerakiAuthUserRequest.md) |  | 

### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkMqttBroker

> map[string]interface{} CreateNetworkMqttBroker(ctx, networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()

Add an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkMqttBrokerRequest := *openapiclient.NewCreateNetworkMqttBrokerRequest("Name_example", "Host_example", int32(123)) // CreateNetworkMqttBrokerRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkMqttBroker(context.Background(), networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkMqttBroker`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMqttBrokerRequest** | [**CreateNetworkMqttBrokerRequest**](CreateNetworkMqttBrokerRequest.md) |  | 

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


## CreateNetworkPiiRequest

> map[string]interface{} CreateNetworkPiiRequest(ctx, networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()

Submit a new delete or restrict processing PII request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkPiiRequestRequest := *openapiclient.NewCreateNetworkPiiRequestRequest() // CreateNetworkPiiRequestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkPiiRequest(context.Background(), networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPiiRequest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPiiRequestRequest** | [**CreateNetworkPiiRequestRequest**](CreateNetworkPiiRequestRequest.md) |  | 

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


## CreateNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner CreateNetworkWebhooksHttpServer(ctx, networkId).CreateNetworkWebhooksHttpServerRequest(createNetworkWebhooksHttpServerRequest).Execute()

Add an HTTP server to a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkWebhooksHttpServerRequest := *openapiclient.NewCreateNetworkWebhooksHttpServerRequest("Name_example", "Url_example") // CreateNetworkWebhooksHttpServerRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkWebhooksHttpServer(context.Background(), networkId).CreateNetworkWebhooksHttpServerRequest(createNetworkWebhooksHttpServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksHttpServerRequest** | [**CreateNetworkWebhooksHttpServerRequest**](CreateNetworkWebhooksHttpServerRequest.md) |  | 

### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner CreateNetworkWebhooksPayloadTemplate(ctx, networkId).CreateNetworkWebhooksPayloadTemplateRequest(createNetworkWebhooksPayloadTemplateRequest).Execute()

Create a webhook payload template for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkWebhooksPayloadTemplateRequest := *openapiclient.NewCreateNetworkWebhooksPayloadTemplateRequest("Name_example") // CreateNetworkWebhooksPayloadTemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkWebhooksPayloadTemplate(context.Background(), networkId).CreateNetworkWebhooksPayloadTemplateRequest(createNetworkWebhooksPayloadTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksPayloadTemplateRequest** | [**CreateNetworkWebhooksPayloadTemplateRequest**](CreateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWebhooksWebhookTest

> CreateNetworkWebhooksWebhookTest201Response CreateNetworkWebhooksWebhookTest(ctx, networkId).CreateNetworkWebhooksWebhookTestRequest(createNetworkWebhooksWebhookTestRequest).Execute()

Send a test webhook for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkWebhooksWebhookTestRequest := *openapiclient.NewCreateNetworkWebhooksWebhookTestRequest("Url_example") // CreateNetworkWebhooksWebhookTestRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateNetworkWebhooksWebhookTest(context.Background(), networkId).CreateNetworkWebhooksWebhookTestRequest(createNetworkWebhooksWebhookTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksWebhookTestRequest** | [**CreateNetworkWebhooksWebhookTestRequest**](CreateNetworkWebhooksWebhookTestRequest.md) |  | 

### Return type

[**CreateNetworkWebhooksWebhookTest201Response**](CreateNetworkWebhooksWebhookTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationNetwork

> GetNetwork200Response CreateOrganizationNetwork(ctx, organizationId).CreateOrganizationNetworkRequest(createOrganizationNetworkRequest).Execute()

Create a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationNetworkRequest := *openapiclient.NewCreateOrganizationNetworkRequest("Name_example", []string{"ProductTypes_example"}) // CreateOrganizationNetworkRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.CreateOrganizationNetwork(context.Background(), organizationId).CreateOrganizationNetworkRequest(createOrganizationNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateOrganizationNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationNetwork`: GetNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateOrganizationNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationNetworkRequest** | [**CreateOrganizationNetworkRequest**](CreateOrganizationNetworkRequest.md) |  | 

### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeferNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response DeferNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Postpone by 1 week all pending staged upgrade stages for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.DeferNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeferNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeferNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeferNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeferNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> DeleteNetwork(ctx, networkId).Execute()

Delete a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetwork(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


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


## DeleteNetworkFirmwareUpgradesStagedGroup

> DeleteNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Delete a Staged Upgrade Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    groupId := "groupId_example" // string | Group ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


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


## DeleteNetworkFloorPlan

> DeleteNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Destroy a floor plan



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    floorPlanId := "floorPlanId_example" // string | Floor plan ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFloorPlanRequest struct via the builder pattern


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


## DeleteNetworkGroupPolicy

> DeleteNetworkGroupPolicy(ctx, networkId, groupPolicyId).Execute()

Delete a group policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    groupPolicyId := "groupPolicyId_example" // string | Group policy ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupPolicyRequest struct via the builder pattern


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


## DeleteNetworkMerakiAuthUser

> DeleteNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Execute()

Deauthorize a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMerakiAuthUserRequest struct via the builder pattern


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


## DeleteNetworkMqttBroker

> DeleteNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Delete an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMqttBrokerRequest struct via the builder pattern


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


## DeleteNetworkPiiRequest

> DeleteNetworkPiiRequest(ctx, networkId, requestId).Execute()

Delete a restrict processing PII request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    requestId := "requestId_example" // string | Request ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPiiRequestRequest struct via the builder pattern


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


## DeleteNetworkWebhooksHttpServer

> DeleteNetworkWebhooksHttpServer(ctx, networkId, httpServerId).Execute()

Delete an HTTP server from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    httpServerId := "httpServerId_example" // string | Http server ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWebhooksHttpServerRequest struct via the builder pattern


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


## DeleteNetworkWebhooksPayloadTemplate

> DeleteNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Destroy a webhook payload template for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.DeleteNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


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


## GetNetwork

> GetNetwork200Response GetNetwork(ctx, networkId).Execute()

Return a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetwork(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetwork`: GetNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsHistory

> []GetNetworkAlertsHistory200ResponseInner GetNetworkAlertsHistory(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the alert history for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkAlertsHistory(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkAlertsHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAlertsHistory`: []GetNetworkAlertsHistory200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkAlertsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkAlertsHistory200ResponseInner**](GetNetworkAlertsHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsSettings

> map[string]interface{} GetNetworkAlertsSettings(ctx, networkId).Execute()

Return the alert configuration for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkAlertsSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkAlertsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAlertsSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsSettingsRequest struct via the builder pattern


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


## GetNetworkBluetoothClient

> map[string]interface{} GetNetworkBluetoothClient(ctx, networkId, bluetoothClientId).IncludeConnectivityHistory(includeConnectivityHistory).ConnectivityHistoryTimespan(connectivityHistoryTimespan).Execute()

Return a Bluetooth client



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    bluetoothClientId := "bluetoothClientId_example" // string | Bluetooth client ID
    includeConnectivityHistory := true // bool | Include the connectivity history for this client (optional)
    connectivityHistoryTimespan := int32(56) // int32 | The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkBluetoothClient(context.Background(), networkId, bluetoothClientId).IncludeConnectivityHistory(includeConnectivityHistory).ConnectivityHistoryTimespan(connectivityHistoryTimespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkBluetoothClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkBluetoothClient`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkBluetoothClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**bluetoothClientId** | **string** | Bluetooth client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBluetoothClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeConnectivityHistory** | **bool** | Include the connectivity history for this client | 
 **connectivityHistoryTimespan** | **int32** | The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used. | 

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


## GetNetworkBluetoothClients

> []map[string]interface{} GetNetworkBluetoothClients(ctx, networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).IncludeConnectivityHistory(includeConnectivityHistory).Execute()

List the Bluetooth clients seen by APs in this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 7 days from today. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 7 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 5 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    includeConnectivityHistory := true // bool | Include the connectivity history for this client (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkBluetoothClients(context.Background(), networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).IncludeConnectivityHistory(includeConnectivityHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkBluetoothClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkBluetoothClients`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkBluetoothClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBluetoothClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 7 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 7 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 5 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **includeConnectivityHistory** | **bool** | Include the connectivity history for this client | 

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


## GetNetworkClient

> GetNetworkClient200Response GetNetworkClient(ctx, networkId, clientId).Execute()

Return the client associated with the given identifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClient(context.Background(), networkId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClient`: GetNetworkClient200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkClient200Response**](GetNetworkClient200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientPolicy

> map[string]interface{} GetNetworkClientPolicy(ctx, networkId, clientId).Execute()

Return the policy assigned to a client on the network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientPolicy(context.Background(), networkId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientPolicyRequest struct via the builder pattern


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


## GetNetworkClientSplashAuthorizationStatus

> map[string]interface{} GetNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).Execute()

Return the splash authorization for a client, for each SSID they've associated with through splash



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientSplashAuthorizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientSplashAuthorizationStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


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


## GetNetworkClientTrafficHistory

> []map[string]interface{} GetNetworkClientTrafficHistory(ctx, networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the client's network traffic data over time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientTrafficHistory(context.Background(), networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientTrafficHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientTrafficHistory`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientTrafficHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientTrafficHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

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


## GetNetworkClientUsageHistory

> []map[string]interface{} GetNetworkClientUsageHistory(ctx, networkId, clientId).Execute()

Return the client's daily usage history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientUsageHistory(context.Background(), networkId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientUsageHistory`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientUsageHistoryRequest struct via the builder pattern


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


## GetNetworkClients

> GetNetworkClients200Response GetNetworkClients(ctx, networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Statuses(statuses).Ip(ip).Ip6(ip6).Ip6Local(ip6Local).Mac(mac).Os(os).PskGroup(pskGroup).Description(description).Vlan(vlan).RecentDeviceConnections(recentDeviceConnections).Execute()

List the clients that have used this network in the timespan



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    statuses := []string{"Statuses_example"} // []string | Filters clients based on status. Can be one of 'Online' or 'Offline'. (optional)
    ip := "ip_example" // string | Filters clients based on a partial or full match for the ip address field. (optional)
    ip6 := "ip6_example" // string | Filters clients based on a partial or full match for the ip6 address field. (optional)
    ip6Local := "ip6Local_example" // string | Filters clients based on a partial or full match for the ip6Local address field. (optional)
    mac := "mac_example" // string | Filters clients based on a partial or full match for the mac address field. (optional)
    os := "os_example" // string | Filters clients based on a partial or full match for the os (operating system) field. (optional)
    pskGroup := "pskGroup_example" // string | Filters clients based on partial or full match for the iPSK name field. (optional)
    description := "description_example" // string | Filters clients based on a partial or full match for the description field. (optional)
    vlan := "vlan_example" // string | Filters clients based on the full match for the VLAN field. (optional)
    recentDeviceConnections := []string{"RecentDeviceConnections_example"} // []string | Filters clients based on recent connection type. Can be one of 'Wired' or 'Wireless'. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClients(context.Background(), networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Statuses(statuses).Ip(ip).Ip6(ip6).Ip6Local(ip6Local).Mac(mac).Os(os).PskGroup(pskGroup).Description(description).Vlan(vlan).RecentDeviceConnections(recentDeviceConnections).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClients`: GetNetworkClients200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **statuses** | **[]string** | Filters clients based on status. Can be one of &#39;Online&#39; or &#39;Offline&#39;. | 
 **ip** | **string** | Filters clients based on a partial or full match for the ip address field. | 
 **ip6** | **string** | Filters clients based on a partial or full match for the ip6 address field. | 
 **ip6Local** | **string** | Filters clients based on a partial or full match for the ip6Local address field. | 
 **mac** | **string** | Filters clients based on a partial or full match for the mac address field. | 
 **os** | **string** | Filters clients based on a partial or full match for the os (operating system) field. | 
 **pskGroup** | **string** | Filters clients based on partial or full match for the iPSK name field. | 
 **description** | **string** | Filters clients based on a partial or full match for the description field. | 
 **vlan** | **string** | Filters clients based on the full match for the VLAN field. | 
 **recentDeviceConnections** | **[]string** | Filters clients based on recent connection type. Can be one of &#39;Wired&#39; or &#39;Wireless&#39;. | 

### Return type

[**GetNetworkClients200Response**](GetNetworkClients200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientsApplicationUsage

> []map[string]interface{} GetNetworkClientsApplicationUsage(ctx, networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the application usage data for clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clients := "clients_example" // string | A list of client keys, MACs or IPs separated by comma.
    ssidNumber := int32(56) // int32 | An SSID number to include. If not specified, eveusage histories application usagents for all SSIDs will be returned. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientsApplicationUsage(context.Background(), networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientsApplicationUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientsApplicationUsage`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientsApplicationUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsApplicationUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clients** | **string** | A list of client keys, MACs or IPs separated by comma. | 
 **ssidNumber** | **int32** | An SSID number to include. If not specified, eveusage histories application usagents for all SSIDs will be returned. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

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


## GetNetworkClientsBandwidthUsageHistory

> []map[string]interface{} GetNetworkClientsBandwidthUsageHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns a timeseries of total traffic consumption rates for all clients on a network within a given timespan, in megabits per second.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientsBandwidthUsageHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientsBandwidthUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientsBandwidthUsageHistory`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientsBandwidthUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsBandwidthUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

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


## GetNetworkClientsOverview

> map[string]interface{} GetNetworkClientsOverview(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Return overview statistics for network clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientsOverview(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientsOverview`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800. | 

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


## GetNetworkClientsUsageHistories

> []map[string]interface{} GetNetworkClientsUsageHistories(ctx, networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the usage histories for clients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clients := "clients_example" // string | A list of client keys, MACs or IPs separated by comma.
    ssidNumber := int32(56) // int32 | An SSID number to include. If not specified, events for all SSIDs will be returned. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkClientsUsageHistories(context.Background(), networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkClientsUsageHistories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientsUsageHistories`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkClientsUsageHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsUsageHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clients** | **string** | A list of client keys, MACs or IPs separated by comma. | 
 **ssidNumber** | **int32** | An SSID number to include. If not specified, events for all SSIDs will be returned. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

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


## GetNetworkDevices

> []map[string]interface{} GetNetworkDevices(ctx, networkId).Execute()

List the devices in a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkDevices(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicesRequest struct via the builder pattern


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


## GetNetworkEvents

> GetNetworkEvents200Response GetNetworkEvents(ctx, networkId).ProductType(productType).IncludedEventTypes(includedEventTypes).ExcludedEventTypes(excludedEventTypes).DeviceMac(deviceMac).DeviceSerial(deviceSerial).DeviceName(deviceName).ClientIp(clientIp).ClientMac(clientMac).ClientName(clientName).SmDeviceMac(smDeviceMac).SmDeviceName(smDeviceName).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the events for the network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    productType := "productType_example" // string | The product type to fetch events for. This parameter is required for networks with multiple device types. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and cloudGateway (optional)
    includedEventTypes := []string{"Inner_example"} // []string | A list of event types. The returned events will be filtered to only include events with these types. (optional)
    excludedEventTypes := []string{"Inner_example"} // []string | A list of event types. The returned events will be filtered to exclude events with these types. (optional)
    deviceMac := "deviceMac_example" // string | The MAC address of the Meraki device which the list of events will be filtered with (optional)
    deviceSerial := "deviceSerial_example" // string | The serial of the Meraki device which the list of events will be filtered with (optional)
    deviceName := "deviceName_example" // string | The name of the Meraki device which the list of events will be filtered with (optional)
    clientIp := "clientIp_example" // string | The IP of the client which the list of events will be filtered with. Only supported for track-by-IP networks. (optional)
    clientMac := "clientMac_example" // string | The MAC address of the client which the list of events will be filtered with. Only supported for track-by-MAC networks. (optional)
    clientName := "clientName_example" // string | The name, or partial name, of the client which the list of events will be filtered with (optional)
    smDeviceMac := "smDeviceMac_example" // string | The MAC address of the Systems Manager device which the list of events will be filtered with (optional)
    smDeviceName := "smDeviceName_example" // string | The name of the Systems Manager device which the list of events will be filtered with (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkEvents(context.Background(), networkId).ProductType(productType).IncludedEventTypes(includedEventTypes).ExcludedEventTypes(excludedEventTypes).DeviceMac(deviceMac).DeviceSerial(deviceSerial).DeviceName(deviceName).ClientIp(clientIp).ClientMac(clientMac).ClientName(clientName).SmDeviceMac(smDeviceMac).SmDeviceName(smDeviceName).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkEvents`: GetNetworkEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productType** | **string** | The product type to fetch events for. This parameter is required for networks with multiple device types. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and cloudGateway | 
 **includedEventTypes** | **[]string** | A list of event types. The returned events will be filtered to only include events with these types. | 
 **excludedEventTypes** | **[]string** | A list of event types. The returned events will be filtered to exclude events with these types. | 
 **deviceMac** | **string** | The MAC address of the Meraki device which the list of events will be filtered with | 
 **deviceSerial** | **string** | The serial of the Meraki device which the list of events will be filtered with | 
 **deviceName** | **string** | The name of the Meraki device which the list of events will be filtered with | 
 **clientIp** | **string** | The IP of the client which the list of events will be filtered with. Only supported for track-by-IP networks. | 
 **clientMac** | **string** | The MAC address of the client which the list of events will be filtered with. Only supported for track-by-MAC networks. | 
 **clientName** | **string** | The name, or partial name, of the client which the list of events will be filtered with | 
 **smDeviceMac** | **string** | The MAC address of the Systems Manager device which the list of events will be filtered with | 
 **smDeviceName** | **string** | The name of the Systems Manager device which the list of events will be filtered with | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetNetworkEvents200Response**](GetNetworkEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEventsEventTypes

> []GetNetworkEventsEventTypes200ResponseInner GetNetworkEventsEventTypes(ctx, networkId).Execute()

List the event type to human-readable description



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkEventsEventTypes(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkEventsEventTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkEventsEventTypes`: []GetNetworkEventsEventTypes200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkEventsEventTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEventsEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkEventsEventTypes200ResponseInner**](GetNetworkEventsEventTypes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgrades

> GetNetworkFirmwareUpgrades200Response GetNetworkFirmwareUpgrades(ctx, networkId).Execute()

Get firmware upgrade information for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkFirmwareUpgrades(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgrades`: GetNetworkFirmwareUpgrades200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgrades200Response**](GetNetworkFirmwareUpgrades200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response GetNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Get the Staged Upgrade Event from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroup

> GetNetworkFirmwareUpgradesStagedGroups200ResponseInner GetNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Get a Staged Upgrade Group from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    groupId := "groupId_example" // string | Group ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroup`: GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroups

> []GetNetworkFirmwareUpgradesStagedGroups200ResponseInner GetNetworkFirmwareUpgradesStagedGroups(ctx, networkId).Execute()

List of Staged Upgrade Groups in a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirmwareUpgradesStagedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroups`: []GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirmwareUpgradesStagedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedStages

> []GetNetworkFirmwareUpgradesStagedStages200ResponseInner GetNetworkFirmwareUpgradesStagedStages(ctx, networkId).Execute()

Order of Staged Upgrade Groups in a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirmwareUpgradesStagedStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedStages`: []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkFirmwareUpgradesStagedStages200ResponseInner**](GetNetworkFirmwareUpgradesStagedStages200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFloorPlan

> map[string]interface{} GetNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Find a floor plan by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    floorPlanId := "floorPlanId_example" // string | Floor plan ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFloorPlan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlanRequest struct via the builder pattern


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


## GetNetworkFloorPlans

> []map[string]interface{} GetNetworkFloorPlans(ctx, networkId).Execute()

List the floor plans that belong to your network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkFloorPlans(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFloorPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFloorPlans`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFloorPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlansRequest struct via the builder pattern


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


## GetNetworkGroupPolicies

> []map[string]interface{} GetNetworkGroupPolicies(ctx, networkId).Execute()

List the group policies in a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkGroupPolicies(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroupPolicies`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPoliciesRequest struct via the builder pattern


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


## GetNetworkGroupPolicy

> map[string]interface{} GetNetworkGroupPolicy(ctx, networkId, groupPolicyId).Execute()

Display a group policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    groupPolicyId := "groupPolicyId_example" // string | Group policy ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroupPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPolicyRequest struct via the builder pattern


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


## GetNetworkHealthAlerts

> []GetNetworkHealthAlerts200ResponseInner GetNetworkHealthAlerts(ctx, networkId).Execute()

Return all global alerts on this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkHealthAlerts(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkHealthAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkHealthAlerts`: []GetNetworkHealthAlerts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkHealthAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkHealthAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkHealthAlerts200ResponseInner**](GetNetworkHealthAlerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner GetNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Execute()

Return the Meraki Auth splash guest, RADIUS, or client VPN user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMerakiAuthUsers

> []GetNetworkMerakiAuthUsers200ResponseInner GetNetworkMerakiAuthUsers(ctx, networkId).Execute()

List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkMerakiAuthUsers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkMerakiAuthUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMerakiAuthUsers`: []GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkMerakiAuthUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMqttBroker

> map[string]interface{} GetNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Return an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMqttBroker`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokerRequest struct via the builder pattern


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


## GetNetworkMqttBrokers

> []map[string]interface{} GetNetworkMqttBrokers(ctx, networkId).Execute()

List the MQTT brokers for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkMqttBrokers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkMqttBrokers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMqttBrokers`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkMqttBrokers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokersRequest struct via the builder pattern


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


## GetNetworkNetflow

> map[string]interface{} GetNetworkNetflow(ctx, networkId).Execute()

Return the NetFlow traffic reporting settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkNetflow(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkNetflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkNetflow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkNetflowRequest struct via the builder pattern


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


## GetNetworkNetworkHealthChannelUtilization

> []map[string]interface{} GetNetworkNetworkHealthChannelUtilization(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Get the channel utilization over each radio for all APs in a network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 600. The default is 600. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkNetworkHealthChannelUtilization(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkNetworkHealthChannelUtilization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkNetworkHealthChannelUtilization`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkNetworkHealthChannelUtilization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkNetworkHealthChannelUtilizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 600. The default is 600. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

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


## GetNetworkPiiPiiKeys

> map[string]interface{} GetNetworkPiiPiiKeys(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

List the keys required to access Personally Identifiable Information (PII) for a given identifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    username := "username_example" // string | The username of a Systems Manager user (optional)
    email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
    mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
    serial := "serial_example" // string | The serial of a Systems Manager device (optional)
    imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
    bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkPiiPiiKeys(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPiiPiiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiPiiKeys`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPiiPiiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiPiiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

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


## GetNetworkPiiRequest

> map[string]interface{} GetNetworkPiiRequest(ctx, networkId, requestId).Execute()

Return a PII request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    requestId := "requestId_example" // string | Request ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiRequest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestRequest struct via the builder pattern


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


## GetNetworkPiiRequests

> []map[string]interface{} GetNetworkPiiRequests(ctx, networkId).Execute()

List the PII requests for this network or organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkPiiRequests(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPiiRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiRequests`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPiiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestsRequest struct via the builder pattern


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


## GetNetworkPiiSmDevicesForKey

> map[string]interface{} GetNetworkPiiSmDevicesForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    username := "username_example" // string | The username of a Systems Manager user (optional)
    email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
    mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
    serial := "serial_example" // string | The serial of a Systems Manager device (optional)
    imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
    bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkPiiSmDevicesForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPiiSmDevicesForKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiSmDevicesForKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPiiSmDevicesForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmDevicesForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

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


## GetNetworkPiiSmOwnersForKey

> map[string]interface{} GetNetworkPiiSmOwnersForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    username := "username_example" // string | The username of a Systems Manager user (optional)
    email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
    mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
    serial := "serial_example" // string | The serial of a Systems Manager device (optional)
    imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
    bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkPiiSmOwnersForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPiiSmOwnersForKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiSmOwnersForKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPiiSmOwnersForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmOwnersForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

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


## GetNetworkPoliciesByClient

> []GetNetworkPoliciesByClient200ResponseInner GetNetworkPoliciesByClient(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).Execute()

Get policies for all clients with policies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkPoliciesByClient(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPoliciesByClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPoliciesByClient`: []GetNetworkPoliciesByClient200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPoliciesByClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoliciesByClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]GetNetworkPoliciesByClient200ResponseInner**](GetNetworkPoliciesByClient200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSettings

> GetNetworkSettings200Response GetNetworkSettings(ctx, networkId).Execute()

Return the settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSettings`: GetNetworkSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSettings200Response**](GetNetworkSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSnmp

> map[string]interface{} GetNetworkSnmp(ctx, networkId).Execute()

Return the SNMP settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkSnmp(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSnmpRequest struct via the builder pattern


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


## GetNetworkSplashLoginAttempts

> []map[string]interface{} GetNetworkSplashLoginAttempts(ctx, networkId).SsidNumber(ssidNumber).LoginIdentifier(loginIdentifier).Timespan(timespan).Execute()

List the splash login attempts for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    ssidNumber := int32(56) // int32 | Only return the login attempts for the specified SSID (optional)
    loginIdentifier := "loginIdentifier_example" // string | The username, email, or phone number used during login (optional)
    timespan := int32(56) // int32 | The timespan, in seconds, for the login attempts. The period will be from [timespan] seconds ago until now. The maximum timespan is 3 months (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkSplashLoginAttempts(context.Background(), networkId).SsidNumber(ssidNumber).LoginIdentifier(loginIdentifier).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkSplashLoginAttempts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSplashLoginAttempts`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkSplashLoginAttempts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSplashLoginAttemptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssidNumber** | **int32** | Only return the login attempts for the specified SSID | 
 **loginIdentifier** | **string** | The username, email, or phone number used during login | 
 **timespan** | **int32** | The timespan, in seconds, for the login attempts. The period will be from [timespan] seconds ago until now. The maximum timespan is 3 months | 

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


## GetNetworkSyslogServers

> GetNetworkSyslogServers200Response GetNetworkSyslogServers(ctx, networkId).Execute()

List the syslog servers for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkSyslogServers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkSyslogServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSyslogServers`: GetNetworkSyslogServers200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSyslogServers200Response**](GetNetworkSyslogServers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTopologyLinkLayer

> map[string]interface{} GetNetworkTopologyLinkLayer(ctx, networkId).Execute()

List the LLDP and CDP information for all discovered devices and connections in a network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkTopologyLinkLayer(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkTopologyLinkLayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTopologyLinkLayer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkTopologyLinkLayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTopologyLinkLayerRequest struct via the builder pattern


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


## GetNetworkTraffic

> []map[string]interface{} GetNetworkTraffic(ctx, networkId).T0(t0).Timespan(timespan).DeviceType(deviceType).Execute()

Return the traffic analysis data for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 30 days. (optional)
    deviceType := "deviceType_example" // string | Filter the data by device type: 'combined', 'wireless', 'switch' or 'appliance'. Defaults to 'combined'. When using 'combined', for each rule the data will come from the device type with the most usage. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkTraffic(context.Background(), networkId).T0(t0).Timespan(timespan).DeviceType(deviceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkTraffic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTraffic`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 30 days. | 
 **deviceType** | **string** | Filter the data by device type: &#39;combined&#39;, &#39;wireless&#39;, &#39;switch&#39; or &#39;appliance&#39;. Defaults to &#39;combined&#39;. When using &#39;combined&#39;, for each rule the data will come from the device type with the most usage. | 

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


## GetNetworkTrafficAnalysis

> map[string]interface{} GetNetworkTrafficAnalysis(ctx, networkId).Execute()

Return the traffic analysis settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkTrafficAnalysis(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkTrafficAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTrafficAnalysis`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficAnalysisRequest struct via the builder pattern


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


## GetNetworkTrafficShapingApplicationCategories

> map[string]interface{} GetNetworkTrafficShapingApplicationCategories(ctx, networkId).Execute()

Returns the application categories for traffic shaping rules.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkTrafficShapingApplicationCategories(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkTrafficShapingApplicationCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTrafficShapingApplicationCategories`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkTrafficShapingApplicationCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficShapingApplicationCategoriesRequest struct via the builder pattern


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


## GetNetworkTrafficShapingDscpTaggingOptions

> []map[string]interface{} GetNetworkTrafficShapingDscpTaggingOptions(ctx, networkId).Execute()

Returns the available DSCP tagging options for your traffic shaping rules.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkTrafficShapingDscpTaggingOptions(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkTrafficShapingDscpTaggingOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTrafficShapingDscpTaggingOptions`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkTrafficShapingDscpTaggingOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficShapingDscpTaggingOptionsRequest struct via the builder pattern


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


## GetNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner GetNetworkWebhooksHttpServer(ctx, networkId, httpServerId).Execute()

Return an HTTP server for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    httpServerId := "httpServerId_example" // string | Http server ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksHttpServers

> []GetNetworkWebhooksHttpServers200ResponseInner GetNetworkWebhooksHttpServers(ctx, networkId).Execute()

List the HTTP servers for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkWebhooksHttpServers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkWebhooksHttpServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksHttpServers`: []GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkWebhooksHttpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Get the webhook payload template for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplates

> []GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplates(ctx, networkId).Execute()

List the webhook payload templates for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkWebhooksPayloadTemplates(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkWebhooksPayloadTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksPayloadTemplates`: []GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkWebhooksPayloadTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksWebhookTest

> CreateNetworkWebhooksWebhookTest201Response GetNetworkWebhooksWebhookTest(ctx, networkId, webhookTestId).Execute()

Return the status of a webhook test for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    webhookTestId := "webhookTestId_example" // string | Webhook test ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetNetworkWebhooksWebhookTest(context.Background(), networkId, webhookTestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**webhookTestId** | **string** | Webhook test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateNetworkWebhooksWebhookTest201Response**](CreateNetworkWebhooksWebhookTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringNetworks

> []GetNetwork200Response GetOrganizationInventoryOnboardingCloudMonitoringNetworks(ctx, organizationId).DeviceType(deviceType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns list of networks eligible for adding cloud monitored device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    deviceType := "deviceType_example" // string | Device Type switch or wireless controller
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).DeviceType(deviceType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: []GetNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceType** | **string** | Device Type switch or wireless controller | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationNetworks

> []GetNetwork200Response GetOrganizationNetworks(ctx, organizationId).ConfigTemplateId(configTemplateId).IsBoundToConfigTemplate(isBoundToConfigTemplate).Tags(tags).TagsFilterType(tagsFilterType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the networks that the user has privileges on in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    configTemplateId := "configTemplateId_example" // string | An optional parameter that is the ID of a config template. Will return all networks bound to that template. (optional)
    isBoundToConfigTemplate := true // bool | An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.GetOrganizationNetworks(context.Background(), organizationId).ConfigTemplateId(configTemplateId).IsBoundToConfigTemplate(isBoundToConfigTemplate).Tags(tags).TagsFilterType(tagsFilterType).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetOrganizationNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationNetworks`: []GetNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetOrganizationNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configTemplateId** | **string** | An optional parameter that is the ID of a config template. Will return all networks bound to that template. | 
 **isBoundToConfigTemplate** | **bool** | An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false. | 
 **tags** | **[]string** | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionNetworkClients

> map[string]interface{} ProvisionNetworkClients(ctx, networkId).ProvisionNetworkClientsRequest(provisionNetworkClientsRequest).Execute()

Provisions a client with a name and policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    provisionNetworkClientsRequest := *openapiclient.NewProvisionNetworkClientsRequest([]openapiclient.ProvisionNetworkClientsRequestClientsInner{*openapiclient.NewProvisionNetworkClientsRequestClientsInner("Mac_example")}, "DevicePolicy_example") // ProvisionNetworkClientsRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.ProvisionNetworkClients(context.Background(), networkId).ProvisionNetworkClientsRequest(provisionNetworkClientsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ProvisionNetworkClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionNetworkClients`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ProvisionNetworkClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionNetworkClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisionNetworkClientsRequest** | [**ProvisionNetworkClientsRequest**](ProvisionNetworkClientsRequest.md) |  | 

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


## RemoveNetworkDevices

> RemoveNetworkDevices(ctx, networkId).RemoveNetworkDevicesRequest(removeNetworkDevicesRequest).Execute()

Remove a single device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    removeNetworkDevicesRequest := *openapiclient.NewRemoveNetworkDevicesRequest("Serial_example") // RemoveNetworkDevicesRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworksApi.RemoveNetworkDevices(context.Background(), networkId).RemoveNetworkDevicesRequest(removeNetworkDevicesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.RemoveNetworkDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeNetworkDevicesRequest** | [**RemoveNetworkDevicesRequest**](RemoveNetworkDevicesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbacksNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response RollbacksNetworkFirmwareUpgradesStagedEvents(ctx, networkId).RollbacksNetworkFirmwareUpgradesStagedEventsRequest(rollbacksNetworkFirmwareUpgradesStagedEventsRequest).Execute()

Rollback a Staged Upgrade Event for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    rollbacksNetworkFirmwareUpgradesStagedEventsRequest := *openapiclient.NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // RollbacksNetworkFirmwareUpgradesStagedEventsRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.RollbacksNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).RollbacksNetworkFirmwareUpgradesStagedEventsRequest(rollbacksNetworkFirmwareUpgradesStagedEventsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.RollbacksNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbacksNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.RollbacksNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbacksNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rollbacksNetworkFirmwareUpgradesStagedEventsRequest** | [**RollbacksNetworkFirmwareUpgradesStagedEventsRequest**](RollbacksNetworkFirmwareUpgradesStagedEventsRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SplitNetwork

> SplitNetwork200Response SplitNetwork(ctx, networkId).Execute()

Split a combined network into individual networks for each type of device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.SplitNetwork(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.SplitNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SplitNetwork`: SplitNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.SplitNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SplitNetwork200Response**](SplitNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindNetwork

> GetNetwork200Response UnbindNetwork(ctx, networkId).UnbindNetworkRequest(unbindNetworkRequest).Execute()

Unbind a network from a template.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    unbindNetworkRequest := *openapiclient.NewUnbindNetworkRequest() // UnbindNetworkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UnbindNetwork(context.Background(), networkId).UnbindNetworkRequest(unbindNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UnbindNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnbindNetwork`: GetNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UnbindNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unbindNetworkRequest** | [**UnbindNetworkRequest**](UnbindNetworkRequest.md) |  | 

### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> GetNetwork200Response UpdateNetwork(ctx, networkId).UpdateNetworkRequest(updateNetworkRequest).Execute()

Update a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkRequest := *openapiclient.NewUpdateNetworkRequest() // UpdateNetworkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetwork(context.Background(), networkId).UpdateNetworkRequest(updateNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetwork`: GetNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkRequest** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md) |  | 

### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAlertsSettings

> map[string]interface{} UpdateNetworkAlertsSettings(ctx, networkId).UpdateNetworkAlertsSettingsRequest(updateNetworkAlertsSettingsRequest).Execute()

Update the alert configuration for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkAlertsSettingsRequest := *openapiclient.NewUpdateNetworkAlertsSettingsRequest() // UpdateNetworkAlertsSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkAlertsSettings(context.Background(), networkId).UpdateNetworkAlertsSettingsRequest(updateNetworkAlertsSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkAlertsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAlertsSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAlertsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkAlertsSettingsRequest** | [**UpdateNetworkAlertsSettingsRequest**](UpdateNetworkAlertsSettingsRequest.md) |  | 

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


## UpdateNetworkClientPolicy

> map[string]interface{} UpdateNetworkClientPolicy(ctx, networkId, clientId).UpdateNetworkClientPolicyRequest(updateNetworkClientPolicyRequest).Execute()

Update the policy assigned to a client on the network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID
    updateNetworkClientPolicyRequest := *openapiclient.NewUpdateNetworkClientPolicyRequest("DevicePolicy_example") // UpdateNetworkClientPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkClientPolicy(context.Background(), networkId, clientId).UpdateNetworkClientPolicyRequest(updateNetworkClientPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkClientPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkClientPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientPolicyRequest** | [**UpdateNetworkClientPolicyRequest**](UpdateNetworkClientPolicyRequest.md) |  | 

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


## UpdateNetworkClientSplashAuthorizationStatus

> map[string]interface{} UpdateNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).UpdateNetworkClientSplashAuthorizationStatusRequest(updateNetworkClientSplashAuthorizationStatusRequest).Execute()

Update a client's splash authorization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID
    updateNetworkClientSplashAuthorizationStatusRequest := *openapiclient.NewUpdateNetworkClientSplashAuthorizationStatusRequest(*openapiclient.NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids()) // UpdateNetworkClientSplashAuthorizationStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).UpdateNetworkClientSplashAuthorizationStatusRequest(updateNetworkClientSplashAuthorizationStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkClientSplashAuthorizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkClientSplashAuthorizationStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientSplashAuthorizationStatusRequest** | [**UpdateNetworkClientSplashAuthorizationStatusRequest**](UpdateNetworkClientSplashAuthorizationStatusRequest.md) |  | 

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


## UpdateNetworkFirmwareUpgrades

> GetNetworkFirmwareUpgrades200Response UpdateNetworkFirmwareUpgrades(ctx, networkId).UpdateNetworkFirmwareUpgradesRequest(updateNetworkFirmwareUpgradesRequest).Execute()

Update firmware upgrade information for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkFirmwareUpgradesRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesRequest() // UpdateNetworkFirmwareUpgradesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkFirmwareUpgrades(context.Background(), networkId).UpdateNetworkFirmwareUpgradesRequest(updateNetworkFirmwareUpgradesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgrades`: GetNetworkFirmwareUpgrades200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesRequest** | [**UpdateNetworkFirmwareUpgradesRequest**](UpdateNetworkFirmwareUpgradesRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgrades200Response**](GetNetworkFirmwareUpgrades200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response UpdateNetworkFirmwareUpgradesStagedEvents(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedEventsRequest(updateNetworkFirmwareUpgradesStagedEventsRequest).Execute()

Update the Staged Upgrade Event for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkFirmwareUpgradesStagedEventsRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // UpdateNetworkFirmwareUpgradesStagedEventsRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedEventsRequest(updateNetworkFirmwareUpgradesStagedEventsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedEventsRequest** | [**UpdateNetworkFirmwareUpgradesStagedEventsRequest**](UpdateNetworkFirmwareUpgradesStagedEventsRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedGroup

> map[string]interface{} UpdateNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()

Update a Staged Upgrade Group for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    groupId := "groupId_example" // string | Group ID
    createNetworkFirmwareUpgradesStagedGroupRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedGroupRequest("Name_example", false) // CreateNetworkFirmwareUpgradesStagedGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkFirmwareUpgradesStagedGroupRequest** | [**CreateNetworkFirmwareUpgradesStagedGroupRequest**](CreateNetworkFirmwareUpgradesStagedGroupRequest.md) |  | 

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


## UpdateNetworkFirmwareUpgradesStagedStages

> []GetNetworkFirmwareUpgradesStagedStages200ResponseInner UpdateNetworkFirmwareUpgradesStagedStages(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedStagesRequest(updateNetworkFirmwareUpgradesStagedStagesRequest).Execute()

Assign Staged Upgrade Group order in the sequence.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkFirmwareUpgradesStagedStagesRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesStagedStagesRequest() // UpdateNetworkFirmwareUpgradesStagedStagesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedStagesRequest(updateNetworkFirmwareUpgradesStagedStagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkFirmwareUpgradesStagedStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedStages`: []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedStagesRequest** | [**UpdateNetworkFirmwareUpgradesStagedStagesRequest**](UpdateNetworkFirmwareUpgradesStagedStagesRequest.md) |  | 

### Return type

[**[]GetNetworkFirmwareUpgradesStagedStages200ResponseInner**](GetNetworkFirmwareUpgradesStagedStages200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFloorPlan

> map[string]interface{} UpdateNetworkFloorPlan(ctx, networkId, floorPlanId).UpdateNetworkFloorPlanRequest(updateNetworkFloorPlanRequest).Execute()

Update a floor plan's geolocation and other meta data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    floorPlanId := "floorPlanId_example" // string | Floor plan ID
    updateNetworkFloorPlanRequest := *openapiclient.NewUpdateNetworkFloorPlanRequest() // UpdateNetworkFloorPlanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkFloorPlan(context.Background(), networkId, floorPlanId).UpdateNetworkFloorPlanRequest(updateNetworkFloorPlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFloorPlan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkFloorPlanRequest** | [**UpdateNetworkFloorPlanRequest**](UpdateNetworkFloorPlanRequest.md) |  | 

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


## UpdateNetworkGroupPolicy

> map[string]interface{} UpdateNetworkGroupPolicy(ctx, networkId, groupPolicyId).UpdateNetworkGroupPolicyRequest(updateNetworkGroupPolicyRequest).Execute()

Update a group policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    groupPolicyId := "groupPolicyId_example" // string | Group policy ID
    updateNetworkGroupPolicyRequest := *openapiclient.NewUpdateNetworkGroupPolicyRequest() // UpdateNetworkGroupPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).UpdateNetworkGroupPolicyRequest(updateNetworkGroupPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkGroupPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkGroupPolicyRequest** | [**UpdateNetworkGroupPolicyRequest**](UpdateNetworkGroupPolicyRequest.md) |  | 

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


## UpdateNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner UpdateNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUserRequest(updateNetworkMerakiAuthUserRequest).Execute()

Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID
    updateNetworkMerakiAuthUserRequest := *openapiclient.NewUpdateNetworkMerakiAuthUserRequest() // UpdateNetworkMerakiAuthUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUserRequest(updateNetworkMerakiAuthUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMerakiAuthUserRequest** | [**UpdateNetworkMerakiAuthUserRequest**](UpdateNetworkMerakiAuthUserRequest.md) |  | 

### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkMqttBroker

> map[string]interface{} UpdateNetworkMqttBroker(ctx, networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()

Update an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID
    updateNetworkMqttBrokerRequest := *openapiclient.NewUpdateNetworkMqttBrokerRequest() // UpdateNetworkMqttBrokerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkMqttBroker`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMqttBrokerRequest** | [**UpdateNetworkMqttBrokerRequest**](UpdateNetworkMqttBrokerRequest.md) |  | 

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


## UpdateNetworkNetflow

> map[string]interface{} UpdateNetworkNetflow(ctx, networkId).UpdateNetworkNetflowRequest(updateNetworkNetflowRequest).Execute()

Update the NetFlow traffic reporting settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkNetflowRequest := *openapiclient.NewUpdateNetworkNetflowRequest() // UpdateNetworkNetflowRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkNetflow(context.Background(), networkId).UpdateNetworkNetflowRequest(updateNetworkNetflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkNetflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkNetflow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkNetflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkNetflowRequest** | [**UpdateNetworkNetflowRequest**](UpdateNetworkNetflowRequest.md) |  | 

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


## UpdateNetworkSettings

> GetNetworkSettings200Response UpdateNetworkSettings(ctx, networkId).UpdateNetworkSettingsRequest(updateNetworkSettingsRequest).Execute()

Update the settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkSettingsRequest := *openapiclient.NewUpdateNetworkSettingsRequest() // UpdateNetworkSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkSettings(context.Background(), networkId).UpdateNetworkSettingsRequest(updateNetworkSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSettings`: GetNetworkSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSettingsRequest** | [**UpdateNetworkSettingsRequest**](UpdateNetworkSettingsRequest.md) |  | 

### Return type

[**GetNetworkSettings200Response**](GetNetworkSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSnmp

> map[string]interface{} UpdateNetworkSnmp(ctx, networkId).UpdateNetworkSnmpRequest(updateNetworkSnmpRequest).Execute()

Update the SNMP settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkSnmpRequest := *openapiclient.NewUpdateNetworkSnmpRequest() // UpdateNetworkSnmpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkSnmp(context.Background(), networkId).UpdateNetworkSnmpRequest(updateNetworkSnmpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSnmpRequest** | [**UpdateNetworkSnmpRequest**](UpdateNetworkSnmpRequest.md) |  | 

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


## UpdateNetworkSyslogServers

> GetNetworkSyslogServers200Response UpdateNetworkSyslogServers(ctx, networkId).UpdateNetworkSyslogServersRequest(updateNetworkSyslogServersRequest).Execute()

Update the syslog servers for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkSyslogServersRequest := *openapiclient.NewUpdateNetworkSyslogServersRequest([]openapiclient.UpdateNetworkSyslogServersRequestServersInner{*openapiclient.NewUpdateNetworkSyslogServersRequestServersInner("Host_example", int32(123), []string{"Roles_example"})}) // UpdateNetworkSyslogServersRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkSyslogServers(context.Background(), networkId).UpdateNetworkSyslogServersRequest(updateNetworkSyslogServersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkSyslogServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSyslogServers`: GetNetworkSyslogServers200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSyslogServersRequest** | [**UpdateNetworkSyslogServersRequest**](UpdateNetworkSyslogServersRequest.md) |  | 

### Return type

[**GetNetworkSyslogServers200Response**](GetNetworkSyslogServers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkTrafficAnalysis

> map[string]interface{} UpdateNetworkTrafficAnalysis(ctx, networkId).UpdateNetworkTrafficAnalysisRequest(updateNetworkTrafficAnalysisRequest).Execute()

Update the traffic analysis settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkTrafficAnalysisRequest := *openapiclient.NewUpdateNetworkTrafficAnalysisRequest() // UpdateNetworkTrafficAnalysisRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkTrafficAnalysis(context.Background(), networkId).UpdateNetworkTrafficAnalysisRequest(updateNetworkTrafficAnalysisRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkTrafficAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkTrafficAnalysis`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkTrafficAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkTrafficAnalysisRequest** | [**UpdateNetworkTrafficAnalysisRequest**](UpdateNetworkTrafficAnalysisRequest.md) |  | 

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


## UpdateNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner UpdateNetworkWebhooksHttpServer(ctx, networkId, httpServerId).UpdateNetworkWebhooksHttpServerRequest(updateNetworkWebhooksHttpServerRequest).Execute()

Update an HTTP server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    httpServerId := "httpServerId_example" // string | Http server ID
    updateNetworkWebhooksHttpServerRequest := *openapiclient.NewUpdateNetworkWebhooksHttpServerRequest() // UpdateNetworkWebhooksHttpServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).UpdateNetworkWebhooksHttpServerRequest(updateNetworkWebhooksHttpServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksHttpServerRequest** | [**UpdateNetworkWebhooksHttpServerRequest**](UpdateNetworkWebhooksHttpServerRequest.md) |  | 

### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner UpdateNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplateRequest(updateNetworkWebhooksPayloadTemplateRequest).Execute()

Update a webhook payload template for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID
    updateNetworkWebhooksPayloadTemplateRequest := *openapiclient.NewUpdateNetworkWebhooksPayloadTemplateRequest() // UpdateNetworkWebhooksPayloadTemplateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.UpdateNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplateRequest(updateNetworkWebhooksPayloadTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksPayloadTemplateRequest** | [**UpdateNetworkWebhooksPayloadTemplateRequest**](UpdateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmxNetworkDevicesClaim

> map[string]interface{} VmxNetworkDevicesClaim(ctx, networkId).VmxNetworkDevicesClaimRequest(vmxNetworkDevicesClaimRequest).Execute()

Claim a vMX into a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    vmxNetworkDevicesClaimRequest := *openapiclient.NewVmxNetworkDevicesClaimRequest("Size_example") // VmxNetworkDevicesClaimRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksApi.VmxNetworkDevicesClaim(context.Background(), networkId).VmxNetworkDevicesClaimRequest(vmxNetworkDevicesClaimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.VmxNetworkDevicesClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmxNetworkDevicesClaim`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.VmxNetworkDevicesClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmxNetworkDevicesClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmxNetworkDevicesClaimRequest** | [**VmxNetworkDevicesClaimRequest**](VmxNetworkDevicesClaimRequest.md) |  | 

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

