package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type CellularGatewayService service

type GetOrganizationCellularGatewayUplinkStatusesQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //A list of network IDs. The returned devices will be filtered to only include these networks.
	Serials       []string `url:"serials[],omitempty"`     //A list of serial numbers. The returned devices will be filtered to only include these serials.
	Iccids        []string `url:"iccids[],omitempty"`      //A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
}

type ResponseCellularGatewayGetDeviceCellularGatewayLan struct {
	DeviceLanIP        string                                                                  `json:"deviceLanIp,omitempty"`        //
	DeviceName         string                                                                  `json:"deviceName,omitempty"`         //
	DeviceSubnet       string                                                                  `json:"deviceSubnet,omitempty"`       //
	FixedIPAssignments *[]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIPAssignments `json:"fixedIpAssignments,omitempty"` //
	ReservedIPRanges   *[]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIPRanges   `json:"reservedIpRanges,omitempty"`   //
}
type ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   //
	Mac  string `json:"mac,omitempty"`  //
	Name string `json:"name,omitempty"` //
}
type ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` //
	End     string `json:"end,omitempty"`     //
	Start   string `json:"start,omitempty"`   //
}
type ResponseCellularGatewayUpdateDeviceCellularGatewayLan interface{}
type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules struct {
	Rules *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules `json:"rules,omitempty"` //
}
type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules struct {
	Access     string `json:"access,omitempty"`     //
	LanIP      string `json:"lanIp,omitempty"`      //
	LocalPort  string `json:"localPort,omitempty"`  //
	Name       string `json:"name,omitempty"`       //
	Protocol   string `json:"protocol,omitempty"`   //
	PublicPort string `json:"publicPort,omitempty"` //
	Uplink     string `json:"uplink,omitempty"`     //
}
type ResponseCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules interface{}
type ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations struct {
	Destinations *[]ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` //
}
type ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     //
	Description string `json:"description,omitempty"` //
	IP          string `json:"ip,omitempty"`          //
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations interface{}
type ResponseCellularGatewayGetNetworkCellularGatewayDhcp struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime,omitempty"`        // DHCP Lease time for all MG in the network.
	DNSCustomNameservers []string `json:"dnsCustomNameservers,omitempty"` // List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.
	DNSNameservers       string   `json:"dnsNameservers,omitempty"`       // DNS name servers mode for all MG in the network.
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime,omitempty"`        // DHCP Lease time for all MG in the network.
	DNSCustomNameservers []string `json:"dnsCustomNameservers,omitempty"` // List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.
	DNSNameservers       string   `json:"dnsNameservers,omitempty"`       // DNS name servers mode for all MG in the network.
}
type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool struct {
	Cidr           string                                                               `json:"cidr,omitempty"`           //
	DeploymentMode string                                                               `json:"deploymentMode,omitempty"` //
	Mask           *int                                                                 `json:"mask,omitempty"`           //
	Subnets        *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets `json:"subnets,omitempty"`        //
}
type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets struct {
	ApplianceIP string `json:"applianceIp,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Serial      string `json:"serial,omitempty"`      //
	Subnet      string `json:"subnet,omitempty"`      //
}
type ResponseCellularGatewayUpdateNetworkCellularGatewaySubnetPool interface{}
type ResponseCellularGatewayGetNetworkCellularGatewayUplink struct {
	BandwidthLimits *ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits `json:"bandwidthLimits,omitempty"` //
}
type ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` //
	LimitUp   *int `json:"limitUp,omitempty"`   //
}
type ResponseCellularGatewayUpdateNetworkCellularGatewayUplink interface{}
type ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses []ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses // Array of ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses
type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses struct {
	LastReportedAt string                                                                            `json:"lastReportedAt,omitempty"` // Last reported time for the device
	Model          string                                                                            `json:"model,omitempty"`          // Device model
	NetworkID      string                                                                            `json:"networkId,omitempty"`      // Network Id
	Serial         string                                                                            `json:"serial,omitempty"`         // Serial number of the device
	Uplinks        *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks `json:"uplinks,omitempty"`        // Uplinks info
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks struct {
	Apn            string                                                                                    `json:"apn,omitempty"`            // Access Point Name
	ConnectionType string                                                                                    `json:"connectionType,omitempty"` // Connection Type
	DNS1           string                                                                                    `json:"dns1,omitempty"`           // Primary DNS IP
	DNS2           string                                                                                    `json:"dns2,omitempty"`           // Secondary DNS IP
	Gateway        string                                                                                    `json:"gateway,omitempty"`        // Gateway IP
	Iccid          string                                                                                    `json:"iccid,omitempty"`          // Integrated Circuit Card Identification Number
	Interface      string                                                                                    `json:"interface,omitempty"`      // Uplink interface
	IP             string                                                                                    `json:"ip,omitempty"`             // Uplink IP
	Model          string                                                                                    `json:"model,omitempty"`          // Uplink model
	Provider       string                                                                                    `json:"provider,omitempty"`       // Network Provider
	PublicIP       string                                                                                    `json:"publicIp,omitempty"`       // Public IP
	SignalStat     *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat `json:"signalStat,omitempty"`     // Tower Signal Status
	SignalType     string                                                                                    `json:"signalType,omitempty"`     // Signal Type
	Status         string                                                                                    `json:"status,omitempty"`         // Uplink status
}
type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat struct {
	Rsrp string `json:"rsrp,omitempty"` // Reference Signal Received Power
	Rsrq string `json:"rsrq,omitempty"` // Reference Signal Received Quality
}
type RequestCellularGatewayUpdateDeviceCellularGatewayLan struct {
	FixedIPAssignments *[]RequestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments `json:"fixedIpAssignments,omitempty"` // list of all fixed IP assignments for a single MG
	ReservedIPRanges   *[]RequestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges   `json:"reservedIpRanges,omitempty"`   // list of all reserved IP ranges for a single MG
}
type RequestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments struct {
	IP   string `json:"ip,omitempty"`   // The IP address you want to assign to a specific server or device
	Mac  string `json:"mac,omitempty"`  // The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	Name string `json:"name,omitempty"` // A descriptive name of the assignment
}
type RequestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges struct {
	Comment string `json:"comment,omitempty"` // Comment explaining the reserved IP range
	End     string `json:"end,omitempty"`     // Ending IP included in the reserved range of IPs
	Start   string `json:"start,omitempty"`   // Starting IP included in the reserved range of IPs
}
type RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules struct {
	Rules *[]RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules `json:"rules,omitempty"` // An array of port forwarding params
}
type RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules struct {
	Access     string   `json:"access,omitempty"`     // *any* or *restricted*. Specify the right to make inbound connections on the specified ports or port ranges. If *restricted*, a list of allowed IPs is mandatory.
	AllowedIPs []string `json:"allowedIps,omitempty"` // An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
	LanIP      string   `json:"lanIp,omitempty"`      // The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LocalPort  string   `json:"localPort,omitempty"`  // A port or port ranges that will receive the forwarded traffic from the WAN
	Name       string   `json:"name,omitempty"`       // A descriptive name for the rule
	Protocol   string   `json:"protocol,omitempty"`   // TCP or UDP
	PublicPort string   `json:"publicPort,omitempty"` // A port or port ranges that will be forwarded to the host on the LAN
}
type RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations struct {
	Destinations *[]RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations `json:"destinations,omitempty"` // The list of connectivity monitoring destinations
}
type RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	Default     *bool  `json:"default,omitempty"`     // Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Description string `json:"description,omitempty"` // Description of the testing destination. Optional, defaults to an empty string
	IP          string `json:"ip,omitempty"`          // The IP address to test connectivity with
}
type RequestCellularGatewayUpdateNetworkCellularGatewayDhcp struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime,omitempty"`        // DHCP Lease time for all MG of the network. Possible values are '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'.
	DNSCustomNameservers []string `json:"dnsCustomNameservers,omitempty"` // list of fixed IPs representing the the DNS Name servers when the mode is 'custom'
	DNSNameservers       string   `json:"dnsNameservers,omitempty"`       // DNS name servers mode for all MG of the network. Possible values are: 'upstream_dns', 'google_dns', 'opendns', 'custom'.
}
type RequestCellularGatewayUpdateNetworkCellularGatewaySubnetPool struct {
	Cidr string `json:"cidr,omitempty"` // CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	Mask *int   `json:"mask,omitempty"` // Mask used for the subnet of all MGs in  this network.
}
type RequestCellularGatewayUpdateNetworkCellularGatewayUplink struct {
	BandwidthLimits *RequestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits `json:"bandwidthLimits,omitempty"` // The bandwidth settings for the 'cellular' uplink
}
type RequestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits struct {
	LimitDown *int `json:"limitDown,omitempty"` // The maximum download limit (integer, in Kbps). null indicates no limit
	LimitUp   *int `json:"limitUp,omitempty"`   // The maximum upload limit (integer, in Kbps). null indicates no limit
}

//GetDeviceCellularGatewayLan Show the LAN Settings of a MG
/* Show the LAN Settings of a MG

@param serial serial path parameter.


*/
func (s *CellularGatewayService) GetDeviceCellularGatewayLan(serial string) (*ResponseCellularGatewayGetDeviceCellularGatewayLan, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/lan"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetDeviceCellularGatewayLan{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCellularGatewayLan")
	}

	result := response.Result().(*ResponseCellularGatewayGetDeviceCellularGatewayLan)
	return result, response, err

}

//GetDeviceCellularGatewayPortForwardingRules Returns the port forwarding rules for a single MG.
/* Returns the port forwarding rules for a single MG.

@param serial serial path parameter.


*/
func (s *CellularGatewayService) GetDeviceCellularGatewayPortForwardingRules(serial string) (*ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/portForwardingRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCellularGatewayPortForwardingRules")
	}

	result := response.Result().(*ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules)
	return result, response, err

}

//GetNetworkCellularGatewayConnectivityMonitoringDestinations Return the connectivity testing destinations for an MG network
/* Return the connectivity testing destinations for an MG network

@param networkID networkId path parameter. Network ID


*/
func (s *CellularGatewayService) GetNetworkCellularGatewayConnectivityMonitoringDestinations(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/connectivityMonitoringDestinations"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewayConnectivityMonitoringDestinations")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations)
	return result, response, err

}

//GetNetworkCellularGatewayDhcp List common DHCP settings of MGs
/* List common DHCP settings of MGs

@param networkID networkId path parameter. Network ID


*/
func (s *CellularGatewayService) GetNetworkCellularGatewayDhcp(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewayDhcp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/dhcp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewayDhcp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewayDhcp")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewayDhcp)
	return result, response, err

}

//GetNetworkCellularGatewaySubnetPool Return the subnet pool and mask configured for MGs in the network.
/* Return the subnet pool and mask configured for MGs in the network.

@param networkID networkId path parameter. Network ID


*/
func (s *CellularGatewayService) GetNetworkCellularGatewaySubnetPool(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/subnetPool"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewaySubnetPool")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool)
	return result, response, err

}

//GetNetworkCellularGatewayUplink Returns the uplink settings for your MG network.
/* Returns the uplink settings for your MG network.

@param networkID networkId path parameter. Network ID


*/
func (s *CellularGatewayService) GetNetworkCellularGatewayUplink(networkID string) (*ResponseCellularGatewayGetNetworkCellularGatewayUplink, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/uplink"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCellularGatewayGetNetworkCellularGatewayUplink{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCellularGatewayUplink")
	}

	result := response.Result().(*ResponseCellularGatewayGetNetworkCellularGatewayUplink)
	return result, response, err

}

//GetOrganizationCellularGatewayUplinkStatuses List the uplink status of every Meraki MG cellular gateway in the organization
/* List the uplink status of every Meraki MG cellular gateway in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCellularGatewayUplinkStatusesQueryParams Filtering parameter


*/
func (s *CellularGatewayService) GetOrganizationCellularGatewayUplinkStatuses(organizationID string, getOrganizationCellularGatewayUplinkStatusesQueryParams *GetOrganizationCellularGatewayUplinkStatusesQueryParams) (*ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/cellularGateway/uplink/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationCellularGatewayUplinkStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCellularGatewayUplinkStatuses")
	}

	result := response.Result().(*ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses)
	return result, response, err

}

//UpdateDeviceCellularGatewayLan Update the LAN Settings for a single MG.
/* Update the LAN Settings for a single MG.

@param serial serial path parameter.
*/
func (s *CellularGatewayService) UpdateDeviceCellularGatewayLan(serial string, requestCellularGatewayUpdateDeviceCellularGatewayLan *RequestCellularGatewayUpdateDeviceCellularGatewayLan) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/lan"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateDeviceCellularGatewayLan).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCellularGatewayLan")
	}

	return response, err

}

//UpdateDeviceCellularGatewayPortForwardingRules Updates the port forwarding rules for a single MG.
/* Updates the port forwarding rules for a single MG.

@param serial serial path parameter.
*/
func (s *CellularGatewayService) UpdateDeviceCellularGatewayPortForwardingRules(serial string, requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules *RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellularGateway/portForwardingRules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCellularGatewayPortForwardingRules")
	}

	return response, err

}

//UpdateNetworkCellularGatewayConnectivityMonitoringDestinations Update the connectivity testing destinations for an MG network
/* Update the connectivity testing destinations for an MG network

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(networkID string, requestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations *RequestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/connectivityMonitoringDestinations"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewayConnectivityMonitoringDestinations).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkCellularGatewayConnectivityMonitoringDestinations")
	}

	return response, err

}

//UpdateNetworkCellularGatewayDhcp Update common DHCP settings of MGs
/* Update common DHCP settings of MGs

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewayDhcp(networkID string, requestCellularGatewayUpdateNetworkCellularGatewayDhcp *RequestCellularGatewayUpdateNetworkCellularGatewayDhcp) (*ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/dhcp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewayDhcp).
		SetResult(&ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkCellularGatewayDhcp")
	}

	result := response.Result().(*ResponseCellularGatewayUpdateNetworkCellularGatewayDhcp)
	return result, response, err

}

//UpdateNetworkCellularGatewaySubnetPool Update the subnet pool and mask configuration for MGs in the network.
/* Update the subnet pool and mask configuration for MGs in the network.

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewaySubnetPool(networkID string, requestCellularGatewayUpdateNetworkCellularGatewaySubnetPool *RequestCellularGatewayUpdateNetworkCellularGatewaySubnetPool) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/subnetPool"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewaySubnetPool).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkCellularGatewaySubnetPool")
	}

	return response, err

}

//UpdateNetworkCellularGatewayUplink Updates the uplink settings for your MG network.
/* Updates the uplink settings for your MG network.

@param networkID networkId path parameter. Network ID
*/
func (s *CellularGatewayService) UpdateNetworkCellularGatewayUplink(networkID string, requestCellularGatewayUpdateNetworkCellularGatewayUplink *RequestCellularGatewayUpdateNetworkCellularGatewayUplink) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/cellularGateway/uplink"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCellularGatewayUpdateNetworkCellularGatewayUplink).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkCellularGatewayUplink")
	}

	return response, err

}
