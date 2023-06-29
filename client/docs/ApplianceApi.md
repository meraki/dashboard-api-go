# \ApplianceApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceApplianceVmxAuthenticationToken**](ApplianceApi.md#CreateDeviceApplianceVmxAuthenticationToken) | **Post** /devices/{serial}/appliance/vmx/authenticationToken | Generate a new vMX authentication token
[**CreateNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#CreateNetworkAppliancePrefixesDelegatedStatic) | **Post** /networks/{networkId}/appliance/prefixes/delegated/statics | Add a static delegated prefix from a network
[**CreateNetworkApplianceRfProfile**](ApplianceApi.md#CreateNetworkApplianceRfProfile) | **Post** /networks/{networkId}/appliance/rfProfiles | Creates new RF profile for this network
[**CreateNetworkApplianceStaticRoute**](ApplianceApi.md#CreateNetworkApplianceStaticRoute) | **Post** /networks/{networkId}/appliance/staticRoutes | Add a static route for an MX or teleworker network
[**CreateNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#CreateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Post** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | Add a custom performance class for an MX network
[**CreateNetworkApplianceVlan**](ApplianceApi.md#CreateNetworkApplianceVlan) | **Post** /networks/{networkId}/appliance/vlans | Add a VLAN
[**DeleteNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#DeleteNetworkAppliancePrefixesDelegatedStatic) | **Delete** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Delete a static delegated prefix from a network
[**DeleteNetworkApplianceRfProfile**](ApplianceApi.md#DeleteNetworkApplianceRfProfile) | **Delete** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Delete a RF Profile
[**DeleteNetworkApplianceStaticRoute**](ApplianceApi.md#DeleteNetworkApplianceStaticRoute) | **Delete** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Delete a static route from an MX or teleworker network
[**DeleteNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#DeleteNetworkApplianceTrafficShapingCustomPerformanceClass) | **Delete** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Delete a custom performance class from an MX network
[**DeleteNetworkApplianceVlan**](ApplianceApi.md#DeleteNetworkApplianceVlan) | **Delete** /networks/{networkId}/appliance/vlans/{vlanId} | Delete a VLAN from a network
[**GetDeviceApplianceDhcpSubnets**](ApplianceApi.md#GetDeviceApplianceDhcpSubnets) | **Get** /devices/{serial}/appliance/dhcp/subnets | Return the DHCP subnet information for an appliance
[**GetDeviceAppliancePerformance**](ApplianceApi.md#GetDeviceAppliancePerformance) | **Get** /devices/{serial}/appliance/performance | Return the performance score for a single MX
[**GetDeviceAppliancePrefixesDelegated**](ApplianceApi.md#GetDeviceAppliancePrefixesDelegated) | **Get** /devices/{serial}/appliance/prefixes/delegated | Return current delegated IPv6 prefixes on an appliance.
[**GetDeviceAppliancePrefixesDelegatedVlanAssignments**](ApplianceApi.md#GetDeviceAppliancePrefixesDelegatedVlanAssignments) | **Get** /devices/{serial}/appliance/prefixes/delegated/vlanAssignments | Return prefixes assigned to all IPv6 enabled VLANs on an appliance.
[**GetDeviceApplianceRadioSettings**](ApplianceApi.md#GetDeviceApplianceRadioSettings) | **Get** /devices/{serial}/appliance/radio/settings | Return the radio settings of an appliance
[**GetDeviceApplianceUplinksSettings**](ApplianceApi.md#GetDeviceApplianceUplinksSettings) | **Get** /devices/{serial}/appliance/uplinks/settings | Return the uplink settings for an MX appliance
[**GetNetworkApplianceClientSecurityEvents**](ApplianceApi.md#GetNetworkApplianceClientSecurityEvents) | **Get** /networks/{networkId}/appliance/clients/{clientId}/security/events | List the security events for a client
[**GetNetworkApplianceConnectivityMonitoringDestinations**](ApplianceApi.md#GetNetworkApplianceConnectivityMonitoringDestinations) | **Get** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Return the connectivity testing destinations for an MX network
[**GetNetworkApplianceContentFiltering**](ApplianceApi.md#GetNetworkApplianceContentFiltering) | **Get** /networks/{networkId}/appliance/contentFiltering | Return the content filtering settings for an MX network
[**GetNetworkApplianceContentFilteringCategories**](ApplianceApi.md#GetNetworkApplianceContentFilteringCategories) | **Get** /networks/{networkId}/appliance/contentFiltering/categories | List all available content filtering categories for an MX network
[**GetNetworkApplianceFirewallCellularFirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Return the cellular firewall rules for an MX network
[**GetNetworkApplianceFirewallFirewalledService**](ApplianceApi.md#GetNetworkApplianceFirewallFirewalledService) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Return the accessibility settings of the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)
[**GetNetworkApplianceFirewallFirewalledServices**](ApplianceApi.md#GetNetworkApplianceFirewallFirewalledServices) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices | List the appliance services and their accessibility rules
[**GetNetworkApplianceFirewallInboundCellularFirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallInboundCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Return the inbound cellular firewall rules for an MX network
[**GetNetworkApplianceFirewallInboundFirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallInboundFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Return the inbound firewall rules for an MX network
[**GetNetworkApplianceFirewallL3FirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallL3FirewallRules) | **Get** /networks/{networkId}/appliance/firewall/l3FirewallRules | Return the L3 firewall rules for an MX network
[**GetNetworkApplianceFirewallL7FirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallL7FirewallRules) | **Get** /networks/{networkId}/appliance/firewall/l7FirewallRules | List the MX L7 firewall rules for an MX network
[**GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories**](ApplianceApi.md#GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories) | **Get** /networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories | Return the L7 firewall application categories and their associated applications for an MX network
[**GetNetworkApplianceFirewallOneToManyNatRules**](ApplianceApi.md#GetNetworkApplianceFirewallOneToManyNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Return the 1:Many NAT mapping rules for an MX network
[**GetNetworkApplianceFirewallOneToOneNatRules**](ApplianceApi.md#GetNetworkApplianceFirewallOneToOneNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Return the 1:1 NAT mapping rules for an MX network
[**GetNetworkApplianceFirewallPortForwardingRules**](ApplianceApi.md#GetNetworkApplianceFirewallPortForwardingRules) | **Get** /networks/{networkId}/appliance/firewall/portForwardingRules | Return the port forwarding rules for an MX network
[**GetNetworkApplianceFirewallSettings**](ApplianceApi.md#GetNetworkApplianceFirewallSettings) | **Get** /networks/{networkId}/appliance/firewall/settings | Return the firewall settings for this network
[**GetNetworkAppliancePort**](ApplianceApi.md#GetNetworkAppliancePort) | **Get** /networks/{networkId}/appliance/ports/{portId} | Return per-port VLAN settings for a single MX port.
[**GetNetworkAppliancePorts**](ApplianceApi.md#GetNetworkAppliancePorts) | **Get** /networks/{networkId}/appliance/ports | List per-port VLAN settings for all ports of a MX.
[**GetNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#GetNetworkAppliancePrefixesDelegatedStatic) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Return a static delegated prefix from a network
[**GetNetworkAppliancePrefixesDelegatedStatics**](ApplianceApi.md#GetNetworkAppliancePrefixesDelegatedStatics) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics | List static delegated prefixes for a network
[**GetNetworkApplianceRfProfile**](ApplianceApi.md#GetNetworkApplianceRfProfile) | **Get** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkApplianceRfProfiles**](ApplianceApi.md#GetNetworkApplianceRfProfiles) | **Get** /networks/{networkId}/appliance/rfProfiles | List the RF profiles for this network
[**GetNetworkApplianceSecurityEvents**](ApplianceApi.md#GetNetworkApplianceSecurityEvents) | **Get** /networks/{networkId}/appliance/security/events | List the security events for a network
[**GetNetworkApplianceSecurityIntrusion**](ApplianceApi.md#GetNetworkApplianceSecurityIntrusion) | **Get** /networks/{networkId}/appliance/security/intrusion | Returns all supported intrusion settings for an MX network
[**GetNetworkApplianceSecurityMalware**](ApplianceApi.md#GetNetworkApplianceSecurityMalware) | **Get** /networks/{networkId}/appliance/security/malware | Returns all supported malware settings for an MX network
[**GetNetworkApplianceSettings**](ApplianceApi.md#GetNetworkApplianceSettings) | **Get** /networks/{networkId}/appliance/settings | Return the appliance settings for a network
[**GetNetworkApplianceSingleLan**](ApplianceApi.md#GetNetworkApplianceSingleLan) | **Get** /networks/{networkId}/appliance/singleLan | Return single LAN configuration
[**GetNetworkApplianceSsid**](ApplianceApi.md#GetNetworkApplianceSsid) | **Get** /networks/{networkId}/appliance/ssids/{number} | Return a single MX SSID
[**GetNetworkApplianceSsids**](ApplianceApi.md#GetNetworkApplianceSsids) | **Get** /networks/{networkId}/appliance/ssids | List the MX SSIDs in a network
[**GetNetworkApplianceStaticRoute**](ApplianceApi.md#GetNetworkApplianceStaticRoute) | **Get** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Return a static route for an MX or teleworker network
[**GetNetworkApplianceStaticRoutes**](ApplianceApi.md#GetNetworkApplianceStaticRoutes) | **Get** /networks/{networkId}/appliance/staticRoutes | List the static routes for an MX or teleworker network
[**GetNetworkApplianceTrafficShaping**](ApplianceApi.md#GetNetworkApplianceTrafficShaping) | **Get** /networks/{networkId}/appliance/trafficShaping | Display the traffic shaping settings for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#GetNetworkApplianceTrafficShapingCustomPerformanceClass) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Return a custom performance class for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses**](ApplianceApi.md#GetNetworkApplianceTrafficShapingCustomPerformanceClasses) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | List all custom performance classes for an MX network
[**GetNetworkApplianceTrafficShapingRules**](ApplianceApi.md#GetNetworkApplianceTrafficShapingRules) | **Get** /networks/{networkId}/appliance/trafficShaping/rules | Display the traffic shaping settings rules for an MX network
[**GetNetworkApplianceTrafficShapingUplinkBandwidth**](ApplianceApi.md#GetNetworkApplianceTrafficShapingUplinkBandwidth) | **Get** /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth | Returns the uplink bandwidth limits for your MX network
[**GetNetworkApplianceTrafficShapingUplinkSelection**](ApplianceApi.md#GetNetworkApplianceTrafficShapingUplinkSelection) | **Get** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Show uplink selection settings for an MX network
[**GetNetworkApplianceUplinksUsageHistory**](ApplianceApi.md#GetNetworkApplianceUplinksUsageHistory) | **Get** /networks/{networkId}/appliance/uplinks/usageHistory | Get the sent and received bytes for each uplink of a network.
[**GetNetworkApplianceVlan**](ApplianceApi.md#GetNetworkApplianceVlan) | **Get** /networks/{networkId}/appliance/vlans/{vlanId} | Return a VLAN
[**GetNetworkApplianceVlans**](ApplianceApi.md#GetNetworkApplianceVlans) | **Get** /networks/{networkId}/appliance/vlans | List the VLANs for an MX network
[**GetNetworkApplianceVlansSettings**](ApplianceApi.md#GetNetworkApplianceVlansSettings) | **Get** /networks/{networkId}/appliance/vlans/settings | Returns the enabled status of VLANs for the network
[**GetNetworkApplianceVpnBgp**](ApplianceApi.md#GetNetworkApplianceVpnBgp) | **Get** /networks/{networkId}/appliance/vpn/bgp | Return a Hub BGP Configuration
[**GetNetworkApplianceVpnSiteToSiteVpn**](ApplianceApi.md#GetNetworkApplianceVpnSiteToSiteVpn) | **Get** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Return the site-to-site VPN settings of a network
[**GetNetworkApplianceWarmSpare**](ApplianceApi.md#GetNetworkApplianceWarmSpare) | **Get** /networks/{networkId}/appliance/warmSpare | Return MX warm spare settings
[**GetOrganizationApplianceSecurityEvents**](ApplianceApi.md#GetOrganizationApplianceSecurityEvents) | **Get** /organizations/{organizationId}/appliance/security/events | List the security events for an organization
[**GetOrganizationApplianceSecurityIntrusion**](ApplianceApi.md#GetOrganizationApplianceSecurityIntrusion) | **Get** /organizations/{organizationId}/appliance/security/intrusion | Returns all supported intrusion settings for an organization
[**GetOrganizationApplianceUplinkStatuses**](ApplianceApi.md#GetOrganizationApplianceUplinkStatuses) | **Get** /organizations/{organizationId}/appliance/uplink/statuses | List the uplink status of every Meraki MX and Z series appliances in the organization
[**GetOrganizationApplianceVpnStats**](ApplianceApi.md#GetOrganizationApplianceVpnStats) | **Get** /organizations/{organizationId}/appliance/vpn/stats | Show VPN history stat for networks in an organization
[**GetOrganizationApplianceVpnStatuses**](ApplianceApi.md#GetOrganizationApplianceVpnStatuses) | **Get** /organizations/{organizationId}/appliance/vpn/statuses | Show VPN status for networks in an organization
[**GetOrganizationApplianceVpnThirdPartyVPNPeers**](ApplianceApi.md#GetOrganizationApplianceVpnThirdPartyVPNPeers) | **Get** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Return the third party VPN peers for an organization
[**GetOrganizationApplianceVpnVpnFirewallRules**](ApplianceApi.md#GetOrganizationApplianceVpnVpnFirewallRules) | **Get** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Return the firewall rules for an organization&#39;s site-to-site VPN
[**SwapNetworkApplianceWarmSpare**](ApplianceApi.md#SwapNetworkApplianceWarmSpare) | **Post** /networks/{networkId}/appliance/warmSpare/swap | Swap MX primary and warm spare appliances
[**UpdateDeviceApplianceRadioSettings**](ApplianceApi.md#UpdateDeviceApplianceRadioSettings) | **Put** /devices/{serial}/appliance/radio/settings | Update the radio settings of an appliance
[**UpdateDeviceApplianceUplinksSettings**](ApplianceApi.md#UpdateDeviceApplianceUplinksSettings) | **Put** /devices/{serial}/appliance/uplinks/settings | Update the uplink settings for an MX appliance
[**UpdateNetworkApplianceConnectivityMonitoringDestinations**](ApplianceApi.md#UpdateNetworkApplianceConnectivityMonitoringDestinations) | **Put** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Update the connectivity testing destinations for an MX network
[**UpdateNetworkApplianceContentFiltering**](ApplianceApi.md#UpdateNetworkApplianceContentFiltering) | **Put** /networks/{networkId}/appliance/contentFiltering | Update the content filtering settings for an MX network
[**UpdateNetworkApplianceFirewallCellularFirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Update the cellular firewall rules of an MX network
[**UpdateNetworkApplianceFirewallFirewalledService**](ApplianceApi.md#UpdateNetworkApplianceFirewallFirewalledService) | **Put** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Updates the accessibility settings for the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)
[**UpdateNetworkApplianceFirewallInboundCellularFirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallInboundCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Update the inbound cellular firewall rules of an MX network
[**UpdateNetworkApplianceFirewallInboundFirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallInboundFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Update the inbound firewall rules of an MX network
[**UpdateNetworkApplianceFirewallL3FirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallL3FirewallRules) | **Put** /networks/{networkId}/appliance/firewall/l3FirewallRules | Update the L3 firewall rules of an MX network
[**UpdateNetworkApplianceFirewallL7FirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallL7FirewallRules) | **Put** /networks/{networkId}/appliance/firewall/l7FirewallRules | Update the MX L7 firewall rules for an MX network
[**UpdateNetworkApplianceFirewallOneToManyNatRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallOneToManyNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Set the 1:Many NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallOneToOneNatRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallOneToOneNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Set the 1:1 NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallPortForwardingRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallPortForwardingRules) | **Put** /networks/{networkId}/appliance/firewall/portForwardingRules | Update the port forwarding rules for an MX network
[**UpdateNetworkApplianceFirewallSettings**](ApplianceApi.md#UpdateNetworkApplianceFirewallSettings) | **Put** /networks/{networkId}/appliance/firewall/settings | Update the firewall settings for this network
[**UpdateNetworkAppliancePort**](ApplianceApi.md#UpdateNetworkAppliancePort) | **Put** /networks/{networkId}/appliance/ports/{portId} | Update the per-port VLAN settings for a single MX port.
[**UpdateNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#UpdateNetworkAppliancePrefixesDelegatedStatic) | **Put** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Update a static delegated prefix from a network
[**UpdateNetworkApplianceRfProfile**](ApplianceApi.md#UpdateNetworkApplianceRfProfile) | **Put** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Updates specified RF profile for this network
[**UpdateNetworkApplianceSecurityIntrusion**](ApplianceApi.md#UpdateNetworkApplianceSecurityIntrusion) | **Put** /networks/{networkId}/appliance/security/intrusion | Set the supported intrusion settings for an MX network
[**UpdateNetworkApplianceSecurityMalware**](ApplianceApi.md#UpdateNetworkApplianceSecurityMalware) | **Put** /networks/{networkId}/appliance/security/malware | Set the supported malware settings for an MX network
[**UpdateNetworkApplianceSettings**](ApplianceApi.md#UpdateNetworkApplianceSettings) | **Put** /networks/{networkId}/appliance/settings | Update the appliance settings for a network
[**UpdateNetworkApplianceSingleLan**](ApplianceApi.md#UpdateNetworkApplianceSingleLan) | **Put** /networks/{networkId}/appliance/singleLan | Update single LAN configuration
[**UpdateNetworkApplianceSsid**](ApplianceApi.md#UpdateNetworkApplianceSsid) | **Put** /networks/{networkId}/appliance/ssids/{number} | Update the attributes of an MX SSID
[**UpdateNetworkApplianceStaticRoute**](ApplianceApi.md#UpdateNetworkApplianceStaticRoute) | **Put** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Update a static route for an MX or teleworker network
[**UpdateNetworkApplianceTrafficShaping**](ApplianceApi.md#UpdateNetworkApplianceTrafficShaping) | **Put** /networks/{networkId}/appliance/trafficShaping | Update the traffic shaping settings for an MX network
[**UpdateNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Put** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Update a custom performance class for an MX network
[**UpdateNetworkApplianceTrafficShapingRules**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingRules) | **Put** /networks/{networkId}/appliance/trafficShaping/rules | Update the traffic shaping settings rules for an MX network
[**UpdateNetworkApplianceTrafficShapingUplinkBandwidth**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingUplinkBandwidth) | **Put** /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth | Updates the uplink bandwidth settings for your MX network.
[**UpdateNetworkApplianceTrafficShapingUplinkSelection**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingUplinkSelection) | **Put** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Update uplink selection settings for an MX network
[**UpdateNetworkApplianceVlan**](ApplianceApi.md#UpdateNetworkApplianceVlan) | **Put** /networks/{networkId}/appliance/vlans/{vlanId} | Update a VLAN
[**UpdateNetworkApplianceVlansSettings**](ApplianceApi.md#UpdateNetworkApplianceVlansSettings) | **Put** /networks/{networkId}/appliance/vlans/settings | Enable/Disable VLANs for the given network
[**UpdateNetworkApplianceVpnBgp**](ApplianceApi.md#UpdateNetworkApplianceVpnBgp) | **Put** /networks/{networkId}/appliance/vpn/bgp | Update a Hub BGP Configuration
[**UpdateNetworkApplianceVpnSiteToSiteVpn**](ApplianceApi.md#UpdateNetworkApplianceVpnSiteToSiteVpn) | **Put** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Update the site-to-site VPN settings of a network
[**UpdateNetworkApplianceWarmSpare**](ApplianceApi.md#UpdateNetworkApplianceWarmSpare) | **Put** /networks/{networkId}/appliance/warmSpare | Update MX warm spare settings
[**UpdateOrganizationApplianceSecurityIntrusion**](ApplianceApi.md#UpdateOrganizationApplianceSecurityIntrusion) | **Put** /organizations/{organizationId}/appliance/security/intrusion | Sets supported intrusion settings for an organization
[**UpdateOrganizationApplianceVpnThirdPartyVPNPeers**](ApplianceApi.md#UpdateOrganizationApplianceVpnThirdPartyVPNPeers) | **Put** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Update the third party VPN peers for an organization
[**UpdateOrganizationApplianceVpnVpnFirewallRules**](ApplianceApi.md#UpdateOrganizationApplianceVpnVpnFirewallRules) | **Put** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Update the firewall rules of an organization&#39;s site-to-site VPN



## CreateDeviceApplianceVmxAuthenticationToken

> CreateDeviceApplianceVmxAuthenticationToken201Response CreateDeviceApplianceVmxAuthenticationToken(ctx, serial).Execute()

Generate a new vMX authentication token



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateDeviceApplianceVmxAuthenticationToken(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateDeviceApplianceVmxAuthenticationToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceApplianceVmxAuthenticationToken`: CreateDeviceApplianceVmxAuthenticationToken201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateDeviceApplianceVmxAuthenticationToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceApplianceVmxAuthenticationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDeviceApplianceVmxAuthenticationToken201Response**](CreateDeviceApplianceVmxAuthenticationToken201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} CreateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId).CreateNetworkAppliancePrefixesDelegatedStaticRequest(createNetworkAppliancePrefixesDelegatedStaticRequest).Execute()

Add a static delegated prefix from a network



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
    createNetworkAppliancePrefixesDelegatedStaticRequest := *openapiclient.NewCreateNetworkAppliancePrefixesDelegatedStaticRequest("Prefix_example", *openapiclient.NewCreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin()) // CreateNetworkAppliancePrefixesDelegatedStaticRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId).CreateNetworkAppliancePrefixesDelegatedStaticRequest(createNetworkAppliancePrefixesDelegatedStaticRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkAppliancePrefixesDelegatedStaticRequest** | [**CreateNetworkAppliancePrefixesDelegatedStaticRequest**](CreateNetworkAppliancePrefixesDelegatedStaticRequest.md) |  | 

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


## CreateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner CreateNetworkApplianceRfProfile(ctx, networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()

Creates new RF profile for this network



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
    createNetworkApplianceRfProfileRequest := *openapiclient.NewCreateNetworkApplianceRfProfileRequest("Name_example") // CreateNetworkApplianceRfProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceRfProfile(context.Background(), networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceRfProfileRequest** | [**CreateNetworkApplianceRfProfileRequest**](CreateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceStaticRoute

> map[string]interface{} CreateNetworkApplianceStaticRoute(ctx, networkId).CreateNetworkApplianceStaticRouteRequest(createNetworkApplianceStaticRouteRequest).Execute()

Add a static route for an MX or teleworker network



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
    createNetworkApplianceStaticRouteRequest := *openapiclient.NewCreateNetworkApplianceStaticRouteRequest("Name_example", "Subnet_example", "GatewayIp_example") // CreateNetworkApplianceStaticRouteRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceStaticRoute(context.Background(), networkId).CreateNetworkApplianceStaticRouteRequest(createNetworkApplianceStaticRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceStaticRoute`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceStaticRouteRequest** | [**CreateNetworkApplianceStaticRouteRequest**](CreateNetworkApplianceStaticRouteRequest.md) |  | 

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


## CreateNetworkApplianceTrafficShapingCustomPerformanceClass

> map[string]interface{} CreateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(createNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()

Add a custom performance class for an MX network



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
    createNetworkApplianceTrafficShapingCustomPerformanceClassRequest := *openapiclient.NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest("Name_example") // CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(createNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceTrafficShapingCustomPerformanceClassRequest** | [**CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

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


## CreateNetworkApplianceVlan

> CreateNetworkApplianceVlan201Response CreateNetworkApplianceVlan(ctx, networkId).CreateNetworkApplianceVlanRequest(createNetworkApplianceVlanRequest).Execute()

Add a VLAN



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
    createNetworkApplianceVlanRequest := *openapiclient.NewCreateNetworkApplianceVlanRequest("Id_example", "Name_example") // CreateNetworkApplianceVlanRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceVlan(context.Background(), networkId).CreateNetworkApplianceVlanRequest(createNetworkApplianceVlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceVlan`: CreateNetworkApplianceVlan201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceVlanRequest** | [**CreateNetworkApplianceVlanRequest**](CreateNetworkApplianceVlanRequest.md) |  | 

### Return type

[**CreateNetworkApplianceVlan201Response**](CreateNetworkApplianceVlan201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAppliancePrefixesDelegatedStatic

> DeleteNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Delete a static delegated prefix from a network



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
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplianceApi.DeleteNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


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


## DeleteNetworkApplianceRfProfile

> DeleteNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplianceApi.DeleteNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceRfProfileRequest struct via the builder pattern


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


## DeleteNetworkApplianceStaticRoute

> DeleteNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Delete a static route from an MX or teleworker network



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
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplianceApi.DeleteNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceStaticRouteRequest struct via the builder pattern


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


## DeleteNetworkApplianceTrafficShapingCustomPerformanceClass

> DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Delete a custom performance class from an MX network



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
    customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplianceApi.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


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


## DeleteNetworkApplianceVlan

> DeleteNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Delete a VLAN from a network



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
    vlanId := "vlanId_example" // string | Vlan ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplianceApi.DeleteNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceVlanRequest struct via the builder pattern


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


## GetDeviceApplianceDhcpSubnets

> []map[string]interface{} GetDeviceApplianceDhcpSubnets(ctx, serial).Execute()

Return the DHCP subnet information for an appliance



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceApplianceDhcpSubnets(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceApplianceDhcpSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceApplianceDhcpSubnets`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceApplianceDhcpSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceDhcpSubnetsRequest struct via the builder pattern


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


## GetDeviceAppliancePerformance

> map[string]interface{} GetDeviceAppliancePerformance(ctx, serial).Execute()

Return the performance score for a single MX



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceAppliancePerformance(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceAppliancePerformance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAppliancePerformance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceAppliancePerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePerformanceRequest struct via the builder pattern


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


## GetDeviceAppliancePrefixesDelegated

> []map[string]interface{} GetDeviceAppliancePrefixesDelegated(ctx, serial).Execute()

Return current delegated IPv6 prefixes on an appliance.



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceAppliancePrefixesDelegated(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceAppliancePrefixesDelegated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAppliancePrefixesDelegated`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceAppliancePrefixesDelegated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedRequest struct via the builder pattern


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


## GetDeviceAppliancePrefixesDelegatedVlanAssignments

> []map[string]interface{} GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx, serial).Execute()

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceAppliancePrefixesDelegatedVlanAssignments(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceAppliancePrefixesDelegatedVlanAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAppliancePrefixesDelegatedVlanAssignments`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceAppliancePrefixesDelegatedVlanAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct via the builder pattern


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


## GetDeviceApplianceRadioSettings

> GetDeviceApplianceRadioSettings200Response GetDeviceApplianceRadioSettings(ctx, serial).Execute()

Return the radio settings of an appliance



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceApplianceRadioSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceApplianceRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceApplianceRadioSettings`: GetDeviceApplianceRadioSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceApplianceRadioSettings200Response**](GetDeviceApplianceRadioSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceApplianceUplinksSettings

> GetDeviceApplianceUplinksSettings200Response GetDeviceApplianceUplinksSettings(ctx, serial).Execute()

Return the uplink settings for an MX appliance



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceApplianceUplinksSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceApplianceUplinksSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceApplianceUplinksSettings`: GetDeviceApplianceUplinksSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceApplianceUplinksSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceUplinksSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceApplianceUplinksSettings200Response**](GetDeviceApplianceUplinksSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceClientSecurityEvents

> []map[string]interface{} GetNetworkApplianceClientSecurityEvents(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a client



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceClientSecurityEvents(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceClientSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceClientSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceClientSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceClientSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

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


## GetNetworkApplianceConnectivityMonitoringDestinations

> map[string]interface{} GetNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).Execute()

Return the connectivity testing destinations for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceConnectivityMonitoringDestinations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


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


## GetNetworkApplianceContentFiltering

> map[string]interface{} GetNetworkApplianceContentFiltering(ctx, networkId).Execute()

Return the content filtering settings for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceContentFiltering(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceContentFiltering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceContentFiltering`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceContentFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceContentFilteringRequest struct via the builder pattern


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


## GetNetworkApplianceContentFilteringCategories

> map[string]interface{} GetNetworkApplianceContentFilteringCategories(ctx, networkId).Execute()

List all available content filtering categories for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceContentFilteringCategories(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceContentFilteringCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceContentFilteringCategories`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceContentFilteringCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceContentFilteringCategoriesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallCellularFirewallRules

> map[string]interface{} GetNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).Execute()

Return the cellular firewall rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallCellularFirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallFirewalledService

> map[string]interface{} GetNetworkApplianceFirewallFirewalledService(ctx, networkId, service).Execute()

Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')



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
    service := "service_example" // string | Service

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallFirewalledService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallFirewalledService`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallFirewalledServices

> []map[string]interface{} GetNetworkApplianceFirewallFirewalledServices(ctx, networkId).Execute()

List the appliance services and their accessibility rules



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallFirewalledServices(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallFirewalledServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallFirewalledServices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallFirewalledServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServicesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallInboundCellularFirewallRules

> []map[string]interface{} GetNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).Execute()

Return the inbound cellular firewall rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallInboundCellularFirewallRules`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallInboundFirewallRules

> map[string]interface{} GetNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).Execute()

Return the inbound firewall rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallInboundFirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallL3FirewallRules

> map[string]interface{} GetNetworkApplianceFirewallL3FirewallRules(ctx, networkId).Execute()

Return the L3 firewall rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallL3FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallL7FirewallRules

> map[string]interface{} GetNetworkApplianceFirewallL7FirewallRules(ctx, networkId).Execute()

List the MX L7 firewall rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallL7FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL7FirewallRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories

> GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(ctx, networkId).Execute()

Return the L7 firewall application categories and their associated applications for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response**](GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallOneToManyNatRules

> map[string]interface{} GetNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).Execute()

Return the 1:Many NAT mapping rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallOneToManyNatRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallOneToOneNatRules

> map[string]interface{} GetNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).Execute()

Return the 1:1 NAT mapping rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallOneToOneNatRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallPortForwardingRules

> map[string]interface{} GetNetworkApplianceFirewallPortForwardingRules(ctx, networkId).Execute()

Return the port forwarding rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallPortForwardingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallSettings

> map[string]interface{} GetNetworkApplianceFirewallSettings(ctx, networkId).Execute()

Return the firewall settings for this network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallSettingsRequest struct via the builder pattern


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


## GetNetworkAppliancePort

> GetNetworkAppliancePorts200ResponseInner GetNetworkAppliancePort(ctx, networkId, portId).Execute()

Return per-port VLAN settings for a single MX port.



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
    portId := "portId_example" // string | Port ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePort(context.Background(), networkId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePort`: GetNetworkAppliancePorts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePorts

> []GetNetworkAppliancePorts200ResponseInner GetNetworkAppliancePorts(ctx, networkId).Execute()

List per-port VLAN settings for all ports of a MX.



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePorts(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePorts`: []GetNetworkAppliancePorts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatic

> GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner GetNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Return a static delegated prefix from a network



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
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePrefixesDelegatedStatic`: GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner**](GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatics

> []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner GetNetworkAppliancePrefixesDelegatedStatics(ctx, networkId).Execute()

List static delegated prefixes for a network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatics(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePrefixesDelegatedStatics`: []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner**](GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner GetNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfiles

> GetNetworkApplianceRfProfiles200Response GetNetworkApplianceRfProfiles(ctx, networkId).Execute()

List the RF profiles for this network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceRfProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceRfProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceRfProfiles`: GetNetworkApplianceRfProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceRfProfiles200Response**](GetNetworkApplianceRfProfiles200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityEvents

> []map[string]interface{} GetNetworkApplianceSecurityEvents(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a network



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSecurityEvents(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

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


## GetNetworkApplianceSecurityIntrusion

> map[string]interface{} GetNetworkApplianceSecurityIntrusion(ctx, networkId).Execute()

Returns all supported intrusion settings for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSecurityIntrusion(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityIntrusion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


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


## GetNetworkApplianceSecurityMalware

> map[string]interface{} GetNetworkApplianceSecurityMalware(ctx, networkId).Execute()

Returns all supported malware settings for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSecurityMalware(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSecurityMalware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityMalware`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityMalwareRequest struct via the builder pattern


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


## GetNetworkApplianceSettings

> GetNetworkApplianceSettings200Response GetNetworkApplianceSettings(ctx, networkId).Execute()

Return the appliance settings for a network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSettings`: GetNetworkApplianceSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceSettings200Response**](GetNetworkApplianceSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSingleLan

> GetNetworkApplianceSingleLan200Response GetNetworkApplianceSingleLan(ctx, networkId).Execute()

Return single LAN configuration



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSingleLan(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSingleLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSingleLan`: GetNetworkApplianceSingleLan200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSingleLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSingleLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceSingleLan200Response**](GetNetworkApplianceSingleLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSsid

> GetNetworkApplianceSsids200ResponseInner GetNetworkApplianceSsid(ctx, networkId, number).Execute()

Return a single MX SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSsid(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSsid`: GetNetworkApplianceSsids200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceSsids200ResponseInner**](GetNetworkApplianceSsids200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSsids

> []GetNetworkApplianceSsids200ResponseInner GetNetworkApplianceSsids(ctx, networkId).Execute()

List the MX SSIDs in a network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSsids(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSsids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSsids`: []GetNetworkApplianceSsids200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceSsids200ResponseInner**](GetNetworkApplianceSsids200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceStaticRoute

> map[string]interface{} GetNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Return a static route for an MX or teleworker network



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
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceStaticRoute`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRouteRequest struct via the builder pattern


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


## GetNetworkApplianceStaticRoutes

> []map[string]interface{} GetNetworkApplianceStaticRoutes(ctx, networkId).Execute()

List the static routes for an MX or teleworker network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceStaticRoutes(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceStaticRoutes`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRoutesRequest struct via the builder pattern


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


## GetNetworkApplianceTrafficShaping

> map[string]interface{} GetNetworkApplianceTrafficShaping(ctx, networkId).Execute()

Display the traffic shaping settings for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShaping(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShaping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShaping`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShaping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRequest struct via the builder pattern


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


## GetNetworkApplianceTrafficShapingCustomPerformanceClass

> map[string]interface{} GetNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Return a custom performance class for an MX network



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
    customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingCustomPerformanceClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


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


## GetNetworkApplianceTrafficShapingCustomPerformanceClasses

> []map[string]interface{} GetNetworkApplianceTrafficShapingCustomPerformanceClasses(ctx, networkId).Execute()

List all custom performance classes for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest struct via the builder pattern


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


## GetNetworkApplianceTrafficShapingRules

> map[string]interface{} GetNetworkApplianceTrafficShapingRules(ctx, networkId).Execute()

Display the traffic shaping settings rules for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


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


## GetNetworkApplianceTrafficShapingUplinkBandwidth

> GetNetworkApplianceTrafficShapingUplinkBandwidth200Response GetNetworkApplianceTrafficShapingUplinkBandwidth(ctx, networkId).Execute()

Returns the uplink bandwidth limits for your MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkBandwidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingUplinkBandwidth`: GetNetworkApplianceTrafficShapingUplinkBandwidth200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkBandwidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingUplinkBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceTrafficShapingUplinkBandwidth200Response**](GetNetworkApplianceTrafficShapingUplinkBandwidth200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingUplinkSelection

> GetNetworkApplianceTrafficShapingUplinkSelection200Response GetNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).Execute()

Show uplink selection settings for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingUplinkSelection`: GetNetworkApplianceTrafficShapingUplinkSelection200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceTrafficShapingUplinkSelection200Response**](GetNetworkApplianceTrafficShapingUplinkSelection200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceUplinksUsageHistory

> []map[string]interface{} GetNetworkApplianceUplinksUsageHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Get the sent and received bytes for each uplink of a network.



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceUplinksUsageHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceUplinksUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceUplinksUsageHistory`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceUplinksUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceUplinksUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60. | 

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


## GetNetworkApplianceVlan

> GetNetworkApplianceVlans200ResponseInner GetNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Return a VLAN



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
    vlanId := "vlanId_example" // string | Vlan ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlan`: GetNetworkApplianceVlans200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlans

> []GetNetworkApplianceVlans200ResponseInner GetNetworkApplianceVlans(ctx, networkId).Execute()

List the VLANs for an MX network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVlans(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlans`: []GetNetworkApplianceVlans200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlansSettings

> map[string]interface{} GetNetworkApplianceVlansSettings(ctx, networkId).Execute()

Returns the enabled status of VLANs for the network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVlansSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVlansSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlansSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansSettingsRequest struct via the builder pattern


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


## GetNetworkApplianceVpnBgp

> map[string]interface{} GetNetworkApplianceVpnBgp(ctx, networkId).Execute()

Return a Hub BGP Configuration



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVpnBgp(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVpnBgp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVpnBgp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnBgpRequest struct via the builder pattern


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


## GetNetworkApplianceVpnSiteToSiteVpn

> GetNetworkApplianceVpnSiteToSiteVpn200Response GetNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).Execute()

Return the site-to-site VPN settings of a network



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVpnSiteToSiteVpn`: GetNetworkApplianceVpnSiteToSiteVpn200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceVpnSiteToSiteVpn200Response**](GetNetworkApplianceVpnSiteToSiteVpn200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceWarmSpare

> map[string]interface{} GetNetworkApplianceWarmSpare(ctx, networkId).Execute()

Return MX warm spare settings



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
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceWarmSpareRequest struct via the builder pattern


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


## GetOrganizationApplianceSecurityEvents

> []map[string]interface{} GetOrganizationApplianceSecurityEvents(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for an organization



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceSecurityEvents(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

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


## GetOrganizationApplianceSecurityIntrusion

> map[string]interface{} GetOrganizationApplianceSecurityIntrusion(ctx, organizationId).Execute()

Returns all supported intrusion settings for an organization



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

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


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


## GetOrganizationApplianceUplinkStatuses

> []map[string]interface{} GetOrganizationApplianceUplinkStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()

List the uplink status of every Meraki MX and Z series appliances in the organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned devices will be filtered to only include these networks. (optional)
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
    iccids := []string{"Inner_example"} // []string | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceUplinkStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceUplinkStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinkStatuses`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceUplinkStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinkStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned devices will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **iccids** | **[]string** | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. | 

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


## GetOrganizationApplianceVpnStats

> []map[string]interface{} GetOrganizationApplianceVpnStats(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).T0(t0).T1(t1).Timespan(timespan).Execute()

Show VPN history stat for networks in an organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnStats(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnStats`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 
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


## GetOrganizationApplianceVpnStatuses

> []map[string]interface{} GetOrganizationApplianceVpnStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Show VPN status for networks in an organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnStatuses`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 

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


## GetOrganizationApplianceVpnThirdPartyVPNPeers

> GetOrganizationApplianceVpnThirdPartyVPNPeers200Response GetOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).Execute()

Return the third party VPN peers for an organization



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

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnThirdPartyVPNPeers`: GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationApplianceVpnThirdPartyVPNPeers200Response**](GetOrganizationApplianceVpnThirdPartyVPNPeers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnVpnFirewallRules

> map[string]interface{} GetOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).Execute()

Return the firewall rules for an organization's site-to-site VPN



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

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnVpnFirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


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


## SwapNetworkApplianceWarmSpare

> map[string]interface{} SwapNetworkApplianceWarmSpare(ctx, networkId).Execute()

Swap MX primary and warm spare appliances



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
    resp, r, err := apiClient.ApplianceApi.SwapNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.SwapNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapNetworkApplianceWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.SwapNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapNetworkApplianceWarmSpareRequest struct via the builder pattern


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


## UpdateDeviceApplianceRadioSettings

> GetDeviceApplianceRadioSettings200Response UpdateDeviceApplianceRadioSettings(ctx, serial).UpdateDeviceApplianceRadioSettingsRequest(updateDeviceApplianceRadioSettingsRequest).Execute()

Update the radio settings of an appliance



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
    serial := "serial_example" // string | Serial
    updateDeviceApplianceRadioSettingsRequest := *openapiclient.NewUpdateDeviceApplianceRadioSettingsRequest() // UpdateDeviceApplianceRadioSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateDeviceApplianceRadioSettings(context.Background(), serial).UpdateDeviceApplianceRadioSettingsRequest(updateDeviceApplianceRadioSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateDeviceApplianceRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceApplianceRadioSettings`: GetDeviceApplianceRadioSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceRadioSettingsRequest** | [**UpdateDeviceApplianceRadioSettingsRequest**](UpdateDeviceApplianceRadioSettingsRequest.md) |  | 

### Return type

[**GetDeviceApplianceRadioSettings200Response**](GetDeviceApplianceRadioSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceApplianceUplinksSettings

> GetDeviceApplianceUplinksSettings200Response UpdateDeviceApplianceUplinksSettings(ctx, serial).UpdateDeviceApplianceUplinksSettingsRequest(updateDeviceApplianceUplinksSettingsRequest).Execute()

Update the uplink settings for an MX appliance



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
    serial := "serial_example" // string | Serial
    updateDeviceApplianceUplinksSettingsRequest := *openapiclient.NewUpdateDeviceApplianceUplinksSettingsRequest(*openapiclient.NewUpdateDeviceApplianceUplinksSettingsRequestInterfaces()) // UpdateDeviceApplianceUplinksSettingsRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateDeviceApplianceUplinksSettings(context.Background(), serial).UpdateDeviceApplianceUplinksSettingsRequest(updateDeviceApplianceUplinksSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateDeviceApplianceUplinksSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceApplianceUplinksSettings`: GetDeviceApplianceUplinksSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateDeviceApplianceUplinksSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceUplinksSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceUplinksSettingsRequest** | [**UpdateDeviceApplianceUplinksSettingsRequest**](UpdateDeviceApplianceUplinksSettingsRequest.md) |  | 

### Return type

[**GetDeviceApplianceUplinksSettings200Response**](GetDeviceApplianceUplinksSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceConnectivityMonitoringDestinations

> map[string]interface{} UpdateNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest(updateNetworkApplianceConnectivityMonitoringDestinationsRequest).Execute()

Update the connectivity testing destinations for an MX network



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
    updateNetworkApplianceConnectivityMonitoringDestinationsRequest := *openapiclient.NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest() // UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest(updateNetworkApplianceConnectivityMonitoringDestinationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceConnectivityMonitoringDestinations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceConnectivityMonitoringDestinationsRequest** | [**UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest**](UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest.md) |  | 

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


## UpdateNetworkApplianceContentFiltering

> map[string]interface{} UpdateNetworkApplianceContentFiltering(ctx, networkId).UpdateNetworkApplianceContentFilteringRequest(updateNetworkApplianceContentFilteringRequest).Execute()

Update the content filtering settings for an MX network



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
    updateNetworkApplianceContentFilteringRequest := *openapiclient.NewUpdateNetworkApplianceContentFilteringRequest() // UpdateNetworkApplianceContentFilteringRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceContentFiltering(context.Background(), networkId).UpdateNetworkApplianceContentFilteringRequest(updateNetworkApplianceContentFilteringRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceContentFiltering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceContentFiltering`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceContentFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceContentFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceContentFilteringRequest** | [**UpdateNetworkApplianceContentFilteringRequest**](UpdateNetworkApplianceContentFilteringRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallCellularFirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()

Update the cellular firewall rules of an MX network



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
    updateNetworkApplianceFirewallCellularFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequest() // UpdateNetworkApplianceFirewallCellularFirewallRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallCellularFirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallCellularFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallCellularFirewallRulesRequest**](UpdateNetworkApplianceFirewallCellularFirewallRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallFirewalledService

> map[string]interface{} UpdateNetworkApplianceFirewallFirewalledService(ctx, networkId, service).UpdateNetworkApplianceFirewallFirewalledServiceRequest(updateNetworkApplianceFirewallFirewalledServiceRequest).Execute()

Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')



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
    service := "service_example" // string | Service
    updateNetworkApplianceFirewallFirewalledServiceRequest := *openapiclient.NewUpdateNetworkApplianceFirewallFirewalledServiceRequest("Access_example") // UpdateNetworkApplianceFirewallFirewalledServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).UpdateNetworkApplianceFirewallFirewalledServiceRequest(updateNetworkApplianceFirewallFirewalledServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallFirewalledService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallFirewalledService`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceFirewallFirewalledServiceRequest** | [**UpdateNetworkApplianceFirewallFirewalledServiceRequest**](UpdateNetworkApplianceFirewallFirewalledServiceRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallInboundCellularFirewallRules

> []map[string]interface{} UpdateNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()

Update the inbound cellular firewall rules of an MX network



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
    updateNetworkApplianceFirewallCellularFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequest() // UpdateNetworkApplianceFirewallCellularFirewallRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallCellularFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallCellularFirewallRulesRequest**](UpdateNetworkApplianceFirewallCellularFirewallRulesRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallInboundFirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()

Update the inbound firewall rules of an MX network



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
    updateNetworkApplianceFirewallInboundFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequest() // UpdateNetworkApplianceFirewallInboundFirewallRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallInboundFirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallInboundFirewallRulesRequest**](UpdateNetworkApplianceFirewallInboundFirewallRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallL3FirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallL3FirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()

Update the L3 firewall rules of an MX network



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
    updateNetworkApplianceFirewallInboundFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequest() // UpdateNetworkApplianceFirewallInboundFirewallRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallL3FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallInboundFirewallRulesRequest**](UpdateNetworkApplianceFirewallInboundFirewallRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallL7FirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallL7FirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallL7FirewallRulesRequest(updateNetworkApplianceFirewallL7FirewallRulesRequest).Execute()

Update the MX L7 firewall rules for an MX network



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
    updateNetworkApplianceFirewallL7FirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallL7FirewallRulesRequest() // UpdateNetworkApplianceFirewallL7FirewallRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallL7FirewallRulesRequest(updateNetworkApplianceFirewallL7FirewallRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallL7FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallL7FirewallRulesRequest** | [**UpdateNetworkApplianceFirewallL7FirewallRulesRequest**](UpdateNetworkApplianceFirewallL7FirewallRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallOneToManyNatRules

> map[string]interface{} UpdateNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToManyNatRulesRequest(updateNetworkApplianceFirewallOneToManyNatRulesRequest).Execute()

Set the 1:Many NAT mapping rules for an MX network



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
    updateNetworkApplianceFirewallOneToManyNatRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner("PublicIp_example", "Uplink_example", []openapiclient.UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner()})}) // UpdateNetworkApplianceFirewallOneToManyNatRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToManyNatRulesRequest(updateNetworkApplianceFirewallOneToManyNatRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallOneToManyNatRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToManyNatRulesRequest** | [**UpdateNetworkApplianceFirewallOneToManyNatRulesRequest**](UpdateNetworkApplianceFirewallOneToManyNatRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallOneToOneNatRules

> map[string]interface{} UpdateNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToOneNatRulesRequest(updateNetworkApplianceFirewallOneToOneNatRulesRequest).Execute()

Set the 1:1 NAT mapping rules for an MX network



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
    updateNetworkApplianceFirewallOneToOneNatRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner("LanIp_example")}) // UpdateNetworkApplianceFirewallOneToOneNatRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToOneNatRulesRequest(updateNetworkApplianceFirewallOneToOneNatRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallOneToOneNatRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToOneNatRulesRequest** | [**UpdateNetworkApplianceFirewallOneToOneNatRulesRequest**](UpdateNetworkApplianceFirewallOneToOneNatRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallPortForwardingRules

> map[string]interface{} UpdateNetworkApplianceFirewallPortForwardingRules(ctx, networkId).UpdateNetworkApplianceFirewallPortForwardingRulesRequest(updateNetworkApplianceFirewallPortForwardingRulesRequest).Execute()

Update the port forwarding rules for an MX network



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
    updateNetworkApplianceFirewallPortForwardingRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner("LanIp_example", "PublicPort_example", "LocalPort_example", []string{"AllowedIps_example"}, "Protocol_example")}) // UpdateNetworkApplianceFirewallPortForwardingRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).UpdateNetworkApplianceFirewallPortForwardingRulesRequest(updateNetworkApplianceFirewallPortForwardingRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallPortForwardingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallPortForwardingRulesRequest** | [**UpdateNetworkApplianceFirewallPortForwardingRulesRequest**](UpdateNetworkApplianceFirewallPortForwardingRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallSettings

> map[string]interface{} UpdateNetworkApplianceFirewallSettings(ctx, networkId).UpdateNetworkApplianceFirewallSettingsRequest(updateNetworkApplianceFirewallSettingsRequest).Execute()

Update the firewall settings for this network



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
    updateNetworkApplianceFirewallSettingsRequest := *openapiclient.NewUpdateNetworkApplianceFirewallSettingsRequest() // UpdateNetworkApplianceFirewallSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallSettings(context.Background(), networkId).UpdateNetworkApplianceFirewallSettingsRequest(updateNetworkApplianceFirewallSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallSettingsRequest** | [**UpdateNetworkApplianceFirewallSettingsRequest**](UpdateNetworkApplianceFirewallSettingsRequest.md) |  | 

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


## UpdateNetworkAppliancePort

> GetNetworkAppliancePorts200ResponseInner UpdateNetworkAppliancePort(ctx, networkId, portId).UpdateNetworkAppliancePortRequest(updateNetworkAppliancePortRequest).Execute()

Update the per-port VLAN settings for a single MX port.



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
    portId := "portId_example" // string | Port ID
    updateNetworkAppliancePortRequest := *openapiclient.NewUpdateNetworkAppliancePortRequest() // UpdateNetworkAppliancePortRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkAppliancePort(context.Background(), networkId, portId).UpdateNetworkAppliancePortRequest(updateNetworkAppliancePortRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkAppliancePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAppliancePort`: GetNetworkAppliancePorts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePortRequest** | [**UpdateNetworkAppliancePortRequest**](UpdateNetworkAppliancePortRequest.md) |  | 

### Return type

[**GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} UpdateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStaticRequest(updateNetworkAppliancePrefixesDelegatedStaticRequest).Execute()

Update a static delegated prefix from a network



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
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID
    updateNetworkAppliancePrefixesDelegatedStaticRequest := *openapiclient.NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest() // UpdateNetworkAppliancePrefixesDelegatedStaticRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStaticRequest(updateNetworkAppliancePrefixesDelegatedStaticRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePrefixesDelegatedStaticRequest** | [**UpdateNetworkAppliancePrefixesDelegatedStaticRequest**](UpdateNetworkAppliancePrefixesDelegatedStaticRequest.md) |  | 

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


## UpdateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner UpdateNetworkApplianceRfProfile(ctx, networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()

Updates specified RF profile for this network



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID
    updateNetworkApplianceRfProfileRequest := *openapiclient.NewUpdateNetworkApplianceRfProfileRequest() // UpdateNetworkApplianceRfProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceRfProfileRequest** | [**UpdateNetworkApplianceRfProfileRequest**](UpdateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityIntrusion

> map[string]interface{} UpdateNetworkApplianceSecurityIntrusion(ctx, networkId).UpdateNetworkApplianceSecurityIntrusionRequest(updateNetworkApplianceSecurityIntrusionRequest).Execute()

Set the supported intrusion settings for an MX network



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
    updateNetworkApplianceSecurityIntrusionRequest := *openapiclient.NewUpdateNetworkApplianceSecurityIntrusionRequest() // UpdateNetworkApplianceSecurityIntrusionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSecurityIntrusion(context.Background(), networkId).UpdateNetworkApplianceSecurityIntrusionRequest(updateNetworkApplianceSecurityIntrusionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSecurityIntrusion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityIntrusionRequest** | [**UpdateNetworkApplianceSecurityIntrusionRequest**](UpdateNetworkApplianceSecurityIntrusionRequest.md) |  | 

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


## UpdateNetworkApplianceSecurityMalware

> map[string]interface{} UpdateNetworkApplianceSecurityMalware(ctx, networkId).UpdateNetworkApplianceSecurityMalwareRequest(updateNetworkApplianceSecurityMalwareRequest).Execute()

Set the supported malware settings for an MX network



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
    updateNetworkApplianceSecurityMalwareRequest := *openapiclient.NewUpdateNetworkApplianceSecurityMalwareRequest("Mode_example") // UpdateNetworkApplianceSecurityMalwareRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSecurityMalware(context.Background(), networkId).UpdateNetworkApplianceSecurityMalwareRequest(updateNetworkApplianceSecurityMalwareRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSecurityMalware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSecurityMalware`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityMalwareRequest** | [**UpdateNetworkApplianceSecurityMalwareRequest**](UpdateNetworkApplianceSecurityMalwareRequest.md) |  | 

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


## UpdateNetworkApplianceSettings

> GetNetworkApplianceSettings200Response UpdateNetworkApplianceSettings(ctx, networkId).UpdateNetworkApplianceSettingsRequest(updateNetworkApplianceSettingsRequest).Execute()

Update the appliance settings for a network



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
    updateNetworkApplianceSettingsRequest := *openapiclient.NewUpdateNetworkApplianceSettingsRequest() // UpdateNetworkApplianceSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSettings(context.Background(), networkId).UpdateNetworkApplianceSettingsRequest(updateNetworkApplianceSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSettings`: GetNetworkApplianceSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSettingsRequest** | [**UpdateNetworkApplianceSettingsRequest**](UpdateNetworkApplianceSettingsRequest.md) |  | 

### Return type

[**GetNetworkApplianceSettings200Response**](GetNetworkApplianceSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSingleLan

> GetNetworkApplianceSingleLan200Response UpdateNetworkApplianceSingleLan(ctx, networkId).UpdateNetworkApplianceSingleLanRequest(updateNetworkApplianceSingleLanRequest).Execute()

Update single LAN configuration



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
    updateNetworkApplianceSingleLanRequest := *openapiclient.NewUpdateNetworkApplianceSingleLanRequest() // UpdateNetworkApplianceSingleLanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSingleLan(context.Background(), networkId).UpdateNetworkApplianceSingleLanRequest(updateNetworkApplianceSingleLanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSingleLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSingleLan`: GetNetworkApplianceSingleLan200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSingleLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSingleLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSingleLanRequest** | [**UpdateNetworkApplianceSingleLanRequest**](UpdateNetworkApplianceSingleLanRequest.md) |  | 

### Return type

[**GetNetworkApplianceSingleLan200Response**](GetNetworkApplianceSingleLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSsid

> GetNetworkApplianceSsids200ResponseInner UpdateNetworkApplianceSsid(ctx, networkId, number).UpdateNetworkApplianceSsidRequest(updateNetworkApplianceSsidRequest).Execute()

Update the attributes of an MX SSID



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
    number := "number_example" // string | Number
    updateNetworkApplianceSsidRequest := *openapiclient.NewUpdateNetworkApplianceSsidRequest() // UpdateNetworkApplianceSsidRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSsid(context.Background(), networkId, number).UpdateNetworkApplianceSsidRequest(updateNetworkApplianceSsidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSsid`: GetNetworkApplianceSsids200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceSsidRequest** | [**UpdateNetworkApplianceSsidRequest**](UpdateNetworkApplianceSsidRequest.md) |  | 

### Return type

[**GetNetworkApplianceSsids200ResponseInner**](GetNetworkApplianceSsids200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceStaticRoute

> map[string]interface{} UpdateNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).UpdateNetworkApplianceStaticRouteRequest(updateNetworkApplianceStaticRouteRequest).Execute()

Update a static route for an MX or teleworker network



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
    staticRouteId := "staticRouteId_example" // string | Static route ID
    updateNetworkApplianceStaticRouteRequest := *openapiclient.NewUpdateNetworkApplianceStaticRouteRequest() // UpdateNetworkApplianceStaticRouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).UpdateNetworkApplianceStaticRouteRequest(updateNetworkApplianceStaticRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceStaticRoute`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceStaticRouteRequest** | [**UpdateNetworkApplianceStaticRouteRequest**](UpdateNetworkApplianceStaticRouteRequest.md) |  | 

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


## UpdateNetworkApplianceTrafficShaping

> map[string]interface{} UpdateNetworkApplianceTrafficShaping(ctx, networkId).UpdateNetworkApplianceTrafficShapingRequest(updateNetworkApplianceTrafficShapingRequest).Execute()

Update the traffic shaping settings for an MX network



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
    updateNetworkApplianceTrafficShapingRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingRequest() // UpdateNetworkApplianceTrafficShapingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShaping(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRequest(updateNetworkApplianceTrafficShapingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShaping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShaping`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShaping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingRequest** | [**UpdateNetworkApplianceTrafficShapingRequest**](UpdateNetworkApplianceTrafficShapingRequest.md) |  | 

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


## UpdateNetworkApplianceTrafficShapingCustomPerformanceClass

> map[string]interface{} UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()

Update a custom performance class for an MX network



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
    customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID
    updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest() // UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest** | [**UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

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


## UpdateNetworkApplianceTrafficShapingRules

> map[string]interface{} UpdateNetworkApplianceTrafficShapingRules(ctx, networkId).UpdateNetworkApplianceTrafficShapingRulesRequest(updateNetworkApplianceTrafficShapingRulesRequest).Execute()

Update the traffic shaping settings rules for an MX network



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
    updateNetworkApplianceTrafficShapingRulesRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingRulesRequest() // UpdateNetworkApplianceTrafficShapingRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingRules(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRulesRequest(updateNetworkApplianceTrafficShapingRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingRulesRequest** | [**UpdateNetworkApplianceTrafficShapingRulesRequest**](UpdateNetworkApplianceTrafficShapingRulesRequest.md) |  | 

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


## UpdateNetworkApplianceTrafficShapingUplinkBandwidth

> map[string]interface{} UpdateNetworkApplianceTrafficShapingUplinkBandwidth(ctx, networkId).UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest(updateNetworkApplianceTrafficShapingUplinkBandwidthRequest).Execute()

Updates the uplink bandwidth settings for your MX network.



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
    updateNetworkApplianceTrafficShapingUplinkBandwidthRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest() // UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest(updateNetworkApplianceTrafficShapingUplinkBandwidthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkBandwidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingUplinkBandwidth`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkBandwidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingUplinkBandwidthRequest** | [**UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest**](UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest.md) |  | 

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


## UpdateNetworkApplianceTrafficShapingUplinkSelection

> GetNetworkApplianceTrafficShapingUplinkSelection200Response UpdateNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest(updateNetworkApplianceTrafficShapingUplinkSelectionRequest).Execute()

Update uplink selection settings for an MX network



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
    updateNetworkApplianceTrafficShapingUplinkSelectionRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest() // UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest(updateNetworkApplianceTrafficShapingUplinkSelectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingUplinkSelection`: GetNetworkApplianceTrafficShapingUplinkSelection200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingUplinkSelectionRequest** | [**UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest**](UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest.md) |  | 

### Return type

[**GetNetworkApplianceTrafficShapingUplinkSelection200Response**](GetNetworkApplianceTrafficShapingUplinkSelection200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlan

> GetNetworkApplianceVlans200ResponseInner UpdateNetworkApplianceVlan(ctx, networkId, vlanId).UpdateNetworkApplianceVlanRequest(updateNetworkApplianceVlanRequest).Execute()

Update a VLAN



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
    vlanId := "vlanId_example" // string | Vlan ID
    updateNetworkApplianceVlanRequest := *openapiclient.NewUpdateNetworkApplianceVlanRequest() // UpdateNetworkApplianceVlanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVlan(context.Background(), networkId, vlanId).UpdateNetworkApplianceVlanRequest(updateNetworkApplianceVlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVlan`: GetNetworkApplianceVlans200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceVlanRequest** | [**UpdateNetworkApplianceVlanRequest**](UpdateNetworkApplianceVlanRequest.md) |  | 

### Return type

[**GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlansSettings

> map[string]interface{} UpdateNetworkApplianceVlansSettings(ctx, networkId).UpdateNetworkApplianceVlansSettingsRequest(updateNetworkApplianceVlansSettingsRequest).Execute()

Enable/Disable VLANs for the given network



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
    updateNetworkApplianceVlansSettingsRequest := *openapiclient.NewUpdateNetworkApplianceVlansSettingsRequest() // UpdateNetworkApplianceVlansSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVlansSettings(context.Background(), networkId).UpdateNetworkApplianceVlansSettingsRequest(updateNetworkApplianceVlansSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVlansSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVlansSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVlansSettingsRequest** | [**UpdateNetworkApplianceVlansSettingsRequest**](UpdateNetworkApplianceVlansSettingsRequest.md) |  | 

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


## UpdateNetworkApplianceVpnBgp

> map[string]interface{} UpdateNetworkApplianceVpnBgp(ctx, networkId).UpdateNetworkApplianceVpnBgpRequest(updateNetworkApplianceVpnBgpRequest).Execute()

Update a Hub BGP Configuration



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
    updateNetworkApplianceVpnBgpRequest := *openapiclient.NewUpdateNetworkApplianceVpnBgpRequest(false) // UpdateNetworkApplianceVpnBgpRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVpnBgp(context.Background(), networkId).UpdateNetworkApplianceVpnBgpRequest(updateNetworkApplianceVpnBgpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVpnBgp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVpnBgp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnBgpRequest** | [**UpdateNetworkApplianceVpnBgpRequest**](UpdateNetworkApplianceVpnBgpRequest.md) |  | 

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


## UpdateNetworkApplianceVpnSiteToSiteVpn

> GetNetworkApplianceVpnSiteToSiteVpn200Response UpdateNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).UpdateNetworkApplianceVpnSiteToSiteVpnRequest(updateNetworkApplianceVpnSiteToSiteVpnRequest).Execute()

Update the site-to-site VPN settings of a network



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
    updateNetworkApplianceVpnSiteToSiteVpnRequest := *openapiclient.NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest("Mode_example") // UpdateNetworkApplianceVpnSiteToSiteVpnRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).UpdateNetworkApplianceVpnSiteToSiteVpnRequest(updateNetworkApplianceVpnSiteToSiteVpnRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVpnSiteToSiteVpn`: GetNetworkApplianceVpnSiteToSiteVpn200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnSiteToSiteVpnRequest** | [**UpdateNetworkApplianceVpnSiteToSiteVpnRequest**](UpdateNetworkApplianceVpnSiteToSiteVpnRequest.md) |  | 

### Return type

[**GetNetworkApplianceVpnSiteToSiteVpn200Response**](GetNetworkApplianceVpnSiteToSiteVpn200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceWarmSpare

> map[string]interface{} UpdateNetworkApplianceWarmSpare(ctx, networkId).UpdateNetworkApplianceWarmSpareRequest(updateNetworkApplianceWarmSpareRequest).Execute()

Update MX warm spare settings



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
    updateNetworkApplianceWarmSpareRequest := *openapiclient.NewUpdateNetworkApplianceWarmSpareRequest(false) // UpdateNetworkApplianceWarmSpareRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceWarmSpare(context.Background(), networkId).UpdateNetworkApplianceWarmSpareRequest(updateNetworkApplianceWarmSpareRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceWarmSpareRequest** | [**UpdateNetworkApplianceWarmSpareRequest**](UpdateNetworkApplianceWarmSpareRequest.md) |  | 

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


## UpdateOrganizationApplianceSecurityIntrusion

> map[string]interface{} UpdateOrganizationApplianceSecurityIntrusion(ctx, organizationId).UpdateOrganizationApplianceSecurityIntrusionRequest(updateOrganizationApplianceSecurityIntrusionRequest).Execute()

Sets supported intrusion settings for an organization



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
    updateOrganizationApplianceSecurityIntrusionRequest := *openapiclient.NewUpdateOrganizationApplianceSecurityIntrusionRequest([]openapiclient.UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner{*openapiclient.NewUpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner("RuleId_example")}) // UpdateOrganizationApplianceSecurityIntrusionRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).UpdateOrganizationApplianceSecurityIntrusionRequest(updateOrganizationApplianceSecurityIntrusionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceSecurityIntrusionRequest** | [**UpdateOrganizationApplianceSecurityIntrusionRequest**](UpdateOrganizationApplianceSecurityIntrusionRequest.md) |  | 

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


## UpdateOrganizationApplianceVpnThirdPartyVPNPeers

> GetOrganizationApplianceVpnThirdPartyVPNPeers200Response UpdateOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(updateOrganizationApplianceVpnThirdPartyVPNPeersRequest).Execute()

Update the third party VPN peers for an organization



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
    updateOrganizationApplianceVpnThirdPartyVPNPeersRequest := *openapiclient.NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest([]openapiclient.UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner{*openapiclient.NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner("Name_example", []string{"PrivateSubnets_example"}, "Secret_example")}) // UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(updateOrganizationApplianceVpnThirdPartyVPNPeersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnThirdPartyVPNPeersRequest** | [**UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest**](UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest.md) |  | 

### Return type

[**GetOrganizationApplianceVpnThirdPartyVPNPeers200Response**](GetOrganizationApplianceVpnThirdPartyVPNPeers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnVpnFirewallRules

> UpdateOrganizationApplianceVpnVpnFirewallRules200Response UpdateOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).UpdateOrganizationApplianceVpnVpnFirewallRulesRequest(updateOrganizationApplianceVpnVpnFirewallRulesRequest).Execute()

Update the firewall rules of an organization's site-to-site VPN



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
    updateOrganizationApplianceVpnVpnFirewallRulesRequest := *openapiclient.NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequest() // UpdateOrganizationApplianceVpnVpnFirewallRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).UpdateOrganizationApplianceVpnVpnFirewallRulesRequest(updateOrganizationApplianceVpnVpnFirewallRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnVpnFirewallRules`: UpdateOrganizationApplianceVpnVpnFirewallRules200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnVpnFirewallRulesRequest** | [**UpdateOrganizationApplianceVpnVpnFirewallRulesRequest**](UpdateOrganizationApplianceVpnVpnFirewallRulesRequest.md) |  | 

### Return type

[**UpdateOrganizationApplianceVpnVpnFirewallRules200Response**](UpdateOrganizationApplianceVpnVpnFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

