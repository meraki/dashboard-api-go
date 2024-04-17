package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DevicesService service

type GetDeviceClientsQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
}
type GetDeviceLossAndLatencyHistoryQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 60 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
	Resolution int     `url:"resolution,omitempty"` //The time resolution in seconds for returned data. The valid resolutions are: 60, 600, 3600, 86400. The default is 60.
	Uplink     string  `url:"uplink,omitempty"`     //The WAN uplink used to obtain the requested stats. Valid uplinks are wan1, wan2, wan3, cellular. The default is wan1.
	IP         string  `url:"ip,omitempty"`         //The destination IP used to obtain the requested stats. This is required.
}
type GetOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Filter results by network.
	Serials       []string `url:"serials[],omitempty"`     //Filter results by device.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days.
	Interval      int      `url:"interval,omitempty"`      //The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
}
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Filter results by network.
	Serials       []string `url:"serials[],omitempty"`     //Filter results by device.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days.
	Interval      int      `url:"interval,omitempty"`      //The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
}
type GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Filter results by network.
	Serials       []string `url:"serials[],omitempty"`     //Filter results by device.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Interval      int      `url:"interval,omitempty"`      //The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
}
type GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Filter results by network.
	Serials       []string `url:"serials[],omitempty"`     //Filter results by device.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Interval      int      `url:"interval,omitempty"`      //The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
}
type GetOrganizationWirelessDevicesPacketLossByClientQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Filter results by network.
	SSIDs         []string `url:"ssids[],omitempty"`       //Filter results by SSID number.
	Bands         []string `url:"bands[],omitempty"`       //Filter results by band. Valid bands are: 2.4, 5, and 6.
	Macs          []string `url:"macs[],omitempty"`        //Filter results by client mac address(es).
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.
}
type GetOrganizationWirelessDevicesPacketLossByDeviceQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Filter results by network.
	Serials       []string `url:"serials[],omitempty"`     //Filter results by device.
	SSIDs         []string `url:"ssids[],omitempty"`       //Filter results by SSID number.
	Bands         []string `url:"bands[],omitempty"`       //Filter results by band. Valid bands are: 2.4, 5, and 6.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.
}
type GetOrganizationWirelessDevicesPacketLossByNetworkQueryParams struct {
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Filter results by network.
	Serials       []string `url:"serials[],omitempty"`     //Filter results by device.
	SSIDs         []string `url:"ssids[],omitempty"`       //Filter results by SSID number.
	Bands         []string `url:"bands[],omitempty"`       //Filter results by band. Valid bands are: 2.4, 5, and 6.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.
}

type ResponseDevicesGetDevice struct {
	Address     string                             `json:"address,omitempty"`     // Physical address of the device
	Details     *[]ResponseDevicesGetDeviceDetails `json:"details,omitempty"`     // Additional device information
	Firmware    string                             `json:"firmware,omitempty"`    // Firmware version of the device
	Imei        string                             `json:"imei,omitempty"`        // IMEI of the device, if applicable
	LanIP       string                             `json:"lanIp,omitempty"`       // LAN IP address of the device
	Lat         *float64                           `json:"lat,omitempty"`         // Latitude of the device
	Lng         *float64                           `json:"lng,omitempty"`         // Longitude of the device
	Mac         string                             `json:"mac,omitempty"`         // MAC address of the device
	Model       string                             `json:"model,omitempty"`       // Model of the device
	Name        string                             `json:"name,omitempty"`        // Name of the device
	NetworkID   string                             `json:"networkId,omitempty"`   // ID of the network the device belongs to
	Notes       string                             `json:"notes,omitempty"`       // Notes for the device, limited to 255 characters
	ProductType string                             `json:"productType,omitempty"` // Product type of the device
	Serial      string                             `json:"serial,omitempty"`      // Serial number of the device
	Tags        []string                           `json:"tags,omitempty"`        // List of tags assigned to the device
}
type ResponseDevicesGetDeviceDetails struct {
	Name  string `json:"name,omitempty"`  // Additional property name
	Value string `json:"value,omitempty"` // Additional property value
}
type ResponseDevicesUpdateDevice interface{}
type ResponseDevicesBlinkDeviceLeds struct {
	Duration *int `json:"duration,omitempty"` // The duration in seconds. Will be between 5 and 120. Default is 20 seconds
	Duty     *int `json:"duty,omitempty"`     // The duty cycle as the percent active. Will be between 10 and 90. Default is 50
	Period   *int `json:"period,omitempty"`   // The period in milliseconds. Will be between 100 and 1000. Default is 160 milliseconds
}
type ResponseDevicesGetDeviceCellularSims struct {
	Sims *[]ResponseDevicesGetDeviceCellularSimsSims `json:"sims,omitempty"` //
}
type ResponseDevicesGetDeviceCellularSimsSims struct {
	Apns      *[]ResponseDevicesGetDeviceCellularSimsSimsApns `json:"apns,omitempty"`      //
	IsPrimary *bool                                           `json:"isPrimary,omitempty"` //
	Slot      string                                          `json:"slot,omitempty"`      //
}
type ResponseDevicesGetDeviceCellularSimsSimsApns struct {
	AllowedIPTypes []string                                                    `json:"allowedIpTypes,omitempty"` //
	Authentication *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication `json:"authentication,omitempty"` //
	Name           string                                                      `json:"name,omitempty"`           //
}
type ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication struct {
	Type     string `json:"type,omitempty"`     //
	Username string `json:"username,omitempty"` //
}
type ResponseDevicesUpdateDeviceCellularSims interface{}
type ResponseDevicesGetDeviceClients []ResponseItemDevicesGetDeviceClients // Array of ResponseDevicesGetDeviceClients
type ResponseItemDevicesGetDeviceClients struct {
	AdaptivePolicyGroup string                                    `json:"adaptivePolicyGroup,omitempty"` // A description of the adaptive policy group
	Description         string                                    `json:"description,omitempty"`         // Short description of the client
	DhcpHostname        string                                    `json:"dhcpHostname,omitempty"`        // The client's DHCP hostname
	ID                  string                                    `json:"id,omitempty"`                  // The ID of the client
	IP                  string                                    `json:"ip,omitempty"`                  // The IP address of the client
	Mac                 string                                    `json:"mac,omitempty"`                 // The MAC address of the client
	MdnsName            string                                    `json:"mdnsName,omitempty"`            // The client's MDNS name
	NamedVLAN           string                                    `json:"namedVlan,omitempty"`           // The owner-assigned name of the VLAN the client is connected to
	Switchport          string                                    `json:"switchport,omitempty"`          // The name of the switchport with clients on it, if the device is a switch
	Usage               *ResponseItemDevicesGetDeviceClientsUsage `json:"usage,omitempty"`               // Client usage data for sent and received
	User                string                                    `json:"user,omitempty"`                // The client user's name
	VLAN                string                                    `json:"vlan,omitempty"`                // The client-assigned name of the VLAN the client is connected to
}
type ResponseItemDevicesGetDeviceClientsUsage struct {
	Recv *float64 `json:"recv,omitempty"` // Usage received by the client
	Sent *float64 `json:"sent,omitempty"` // Usage sent by the client
}
type ResponseDevicesCreateDeviceLiveToolsArpTable struct {
	ArpTableID string                                                `json:"arpTableId,omitempty"` // Id of the ARP table request. Used to check the status of the request.
	Callback   *ResponseDevicesCreateDeviceLiveToolsArpTableCallback `json:"callback,omitempty"`   // Information for callback used to send back results
	Request    *ResponseDevicesCreateDeviceLiveToolsArpTableRequest  `json:"request,omitempty"`    // ARP table request parameters
	Status     string                                                `json:"status,omitempty"`     // Status of the ARP table request.
	URL        string                                                `json:"url,omitempty"`        // GET this url to check the status of your ARP table request.
}
type ResponseDevicesCreateDeviceLiveToolsArpTableCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesCreateDeviceLiveToolsArpTableRequest struct {
	Serial string `json:"serial,omitempty"` // Device serial number
}
type ResponseDevicesGetDeviceLiveToolsArpTable struct {
	ArpTableID string                                              `json:"arpTableId,omitempty"` // Id of the ARP table request. Used to check the status of the request.
	Entries    *[]ResponseDevicesGetDeviceLiveToolsArpTableEntries `json:"entries,omitempty"`    // The ARP table entries
	Error      string                                              `json:"error,omitempty"`      // An error message for a failed execution
	Request    *ResponseDevicesGetDeviceLiveToolsArpTableRequest   `json:"request,omitempty"`    // ARP table request parameters
	Status     string                                              `json:"status,omitempty"`     // Status of the ARP table request.
	URL        string                                              `json:"url,omitempty"`        // GET this url to check the status of your ARP table request.
}
type ResponseDevicesGetDeviceLiveToolsArpTableEntries struct {
	IP            string `json:"ip,omitempty"`            // The IP address of the ARP table entry
	LastUpdatedAt string `json:"lastUpdatedAt,omitempty"` // Time of the last update of the ARP table entry
	Mac           string `json:"mac,omitempty"`           // The MAC address of the ARP table entry
	VLANID        *int   `json:"vlanId,omitempty"`        // The VLAN ID of the ARP table entry
}
type ResponseDevicesGetDeviceLiveToolsArpTableRequest struct {
	Serial string `json:"serial,omitempty"` // Device serial number
}
type ResponseDevicesCreateDeviceLiveToolsCableTest struct {
	CableTestID string                                                 `json:"cableTestId,omitempty"` // Id of the cable test request. Used to check the status of the request.
	Callback    *ResponseDevicesCreateDeviceLiveToolsCableTestCallback `json:"callback,omitempty"`    // Information for callback used to send back results
	Request     *ResponseDevicesCreateDeviceLiveToolsCableTestRequest  `json:"request,omitempty"`     // Cable test request parameters
	Status      string                                                 `json:"status,omitempty"`      // Status of the cable test request.
	URL         string                                                 `json:"url,omitempty"`         // GET this url to check the status of your cable test request.
}
type ResponseDevicesCreateDeviceLiveToolsCableTestCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesCreateDeviceLiveToolsCableTestRequest struct {
	Ports  []string `json:"ports,omitempty"`  // A list of ports for which to perform the cable test.
	Serial string   `json:"serial,omitempty"` // Device serial number
}
type ResponseDevicesGetDeviceLiveToolsCableTest struct {
	CableTestID string                                               `json:"cableTestId,omitempty"` // Id of the cable test request. Used to check the status of the request.
	Error       string                                               `json:"error,omitempty"`       // An error message for a failed execution
	Request     *ResponseDevicesGetDeviceLiveToolsCableTestRequest   `json:"request,omitempty"`     // Cable test request parameters
	Results     *[]ResponseDevicesGetDeviceLiveToolsCableTestResults `json:"results,omitempty"`     // Results of the cable test request, one for each requested port.
	Status      string                                               `json:"status,omitempty"`      // Status of the cable test request.
	URL         string                                               `json:"url,omitempty"`         // GET this url to check the status of your cable test request.
}
type ResponseDevicesGetDeviceLiveToolsCableTestRequest struct {
	Ports  []string `json:"ports,omitempty"`  // A list of ports for which to perform the cable test.
	Serial string   `json:"serial,omitempty"` // Device serial number
}
type ResponseDevicesGetDeviceLiveToolsCableTestResults struct {
	Error     string                                                    `json:"error,omitempty"`     // If an error occurred during the cable test, the error message will be populated here.
	Pairs     *[]ResponseDevicesGetDeviceLiveToolsCableTestResultsPairs `json:"pairs,omitempty"`     // Results for each twisted pair within the cable.
	Port      string                                                    `json:"port,omitempty"`      // The port for which the test was performed.
	SpeedMbps *int                                                      `json:"speedMbps,omitempty"` // Speed in Mbps.  A speed of 0 indicates the port is down or the port speed is automatic.
	Status    string                                                    `json:"status,omitempty"`    // The current status of the port. If the cable test is still being performed on the port, "in-progress" is used. If an error occurred during the cable test, "error" is used and the error property will be populated.
}
type ResponseDevicesGetDeviceLiveToolsCableTestResultsPairs struct {
	Index        *int   `json:"index,omitempty"`        // The index of the twisted pair tested.
	LengthMeters *int   `json:"lengthMeters,omitempty"` // The detected length of the twisted pair.
	Status       string `json:"status,omitempty"`       // The test result of the twisted pair tested.
}
type ResponseDevicesCreateDeviceLiveToolsPing struct {
	Callback *ResponseDevicesCreateDeviceLiveToolsPingCallback `json:"callback,omitempty"` // Information for callback used to send back results
	PingID   string                                            `json:"pingId,omitempty"`   // Id to check the status of your ping request.
	Request  *ResponseDevicesCreateDeviceLiveToolsPingRequest  `json:"request,omitempty"`  // Ping request parameters
	Status   string                                            `json:"status,omitempty"`   // Status of the ping request.
	URL      string                                            `json:"url,omitempty"`      // GET this url to check the status of your ping request.
}
type ResponseDevicesCreateDeviceLiveToolsPingCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesCreateDeviceLiveToolsPingRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPing struct {
	PingID  string                                        `json:"pingId,omitempty"`  // Id to check the status of your ping request.
	Request *ResponseDevicesGetDeviceLiveToolsPingRequest `json:"request,omitempty"` // Ping request parameters
	Results *ResponseDevicesGetDeviceLiveToolsPingResults `json:"results,omitempty"` // Results of the ping request.
	Status  string                                        `json:"status,omitempty"`  // Status of the ping request.
	URL     string                                        `json:"url,omitempty"`     // GET this url to check the status of your ping request.
}
type ResponseDevicesGetDeviceLiveToolsPingRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPingResults struct {
	Latencies *ResponseDevicesGetDeviceLiveToolsPingResultsLatencies `json:"latencies,omitempty"` // Packet latency stats
	Loss      *ResponseDevicesGetDeviceLiveToolsPingResultsLoss      `json:"loss,omitempty"`      // Lost packets
	Received  *int                                                   `json:"received,omitempty"`  // Number of packets received
	Replies   *[]ResponseDevicesGetDeviceLiveToolsPingResultsReplies `json:"replies,omitempty"`   // Received packets
	Sent      *int                                                   `json:"sent,omitempty"`      // Number of packets sent
}
type ResponseDevicesGetDeviceLiveToolsPingResultsLatencies struct {
	Average *float64 `json:"average,omitempty"` // Average latency
	Maximum *float64 `json:"maximum,omitempty"` // Maximum latency
	Minimum *float64 `json:"minimum,omitempty"` // Minimum latency
}
type ResponseDevicesGetDeviceLiveToolsPingResultsLoss struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of packets lost
}
type ResponseDevicesGetDeviceLiveToolsPingResultsReplies struct {
	Latency    *float64 `json:"latency,omitempty"`    // Latency of the packet in milliseconds
	SequenceID *int     `json:"sequenceId,omitempty"` // Sequence ID of the packet
	Size       *int     `json:"size,omitempty"`       // Size of the packet in bytes
}
type ResponseDevicesCreateDeviceLiveToolsPingDevice struct {
	Callback *ResponseDevicesCreateDeviceLiveToolsPingDeviceCallback `json:"callback,omitempty"` // Information for callback used to send back results
	PingID   string                                                  `json:"pingId,omitempty"`   // Id to check the status of your ping request.
	Request  *ResponseDevicesCreateDeviceLiveToolsPingDeviceRequest  `json:"request,omitempty"`  // Ping request parameters
	Status   string                                                  `json:"status,omitempty"`   // Status of the ping request.
	URL      string                                                  `json:"url,omitempty"`      // GET this url to check the status of your ping request.
}
type ResponseDevicesCreateDeviceLiveToolsPingDeviceCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesCreateDeviceLiveToolsPingDeviceRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPingDevice struct {
	Callback *ResponseDevicesGetDeviceLiveToolsPingDeviceCallback `json:"callback,omitempty"` // Information for callback used to send back results
	PingID   string                                               `json:"pingId,omitempty"`   // Id to check the status of your ping request.
	Request  *ResponseDevicesGetDeviceLiveToolsPingDeviceRequest  `json:"request,omitempty"`  // Ping request parameters
	Results  *ResponseDevicesGetDeviceLiveToolsPingDeviceResults  `json:"results,omitempty"`  // Results of the ping request.
	Status   string                                               `json:"status,omitempty"`   // Status of the ping request.
	URL      string                                               `json:"url,omitempty"`      // GET this url to check the status of your ping request.
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceRequest struct {
	Count  *int   `json:"count,omitempty"`  // Number of pings to send
	Serial string `json:"serial,omitempty"` // Device serial number
	Target string `json:"target,omitempty"` // IP address or FQDN to ping
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResults struct {
	Latencies *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies `json:"latencies,omitempty"` // Packet latency stats
	Loss      *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss      `json:"loss,omitempty"`      // Lost packets
	Received  *int                                                         `json:"received,omitempty"`  // Number of packets received
	Replies   *[]ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies `json:"replies,omitempty"`   // Received packets
	Sent      *int                                                         `json:"sent,omitempty"`      // Number of packets sent
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies struct {
	Average *float64 `json:"average,omitempty"` // Average latency
	Maximum *float64 `json:"maximum,omitempty"` // Maximum latency
	Minimum *float64 `json:"minimum,omitempty"` // Minimum latency
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of packets lost
}
type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies struct {
	Latency    *float64 `json:"latency,omitempty"`    // Latency of the packet in milliseconds
	SequenceID *int     `json:"sequenceId,omitempty"` // Sequence ID of the packet
	Size       *int     `json:"size,omitempty"`       // Size of the packet in bytes
}
type ResponseDevicesCreateDeviceLiveToolsThroughputTest struct {
	Callback         *ResponseDevicesCreateDeviceLiveToolsThroughputTestCallback `json:"callback,omitempty"`         // Information for callback used to send back results
	Error            string                                                      `json:"error,omitempty"`            // Description of the error.
	Request          *ResponseDevicesCreateDeviceLiveToolsThroughputTestRequest  `json:"request,omitempty"`          // The parameters of the throughput test request
	Result           *ResponseDevicesCreateDeviceLiveToolsThroughputTestResult   `json:"result,omitempty"`           // Result of the throughput test request
	Status           string                                                      `json:"status,omitempty"`           // Status of the throughput test request
	ThroughputTestID string                                                      `json:"throughputTestId,omitempty"` // ID of throughput test job
	URL              string                                                      `json:"url,omitempty"`              // GET this url to check the status of your throughput test request
}
type ResponseDevicesCreateDeviceLiveToolsThroughputTestCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesCreateDeviceLiveToolsThroughputTestRequest struct {
	Serial string `json:"serial,omitempty"` // Device serial number
}
type ResponseDevicesCreateDeviceLiveToolsThroughputTestResult struct {
	Speeds *ResponseDevicesCreateDeviceLiveToolsThroughputTestResultSpeeds `json:"speeds,omitempty"` // Shows the speeds (Mbps)
}
type ResponseDevicesCreateDeviceLiveToolsThroughputTestResultSpeeds struct {
	Downstream *int `json:"downstream,omitempty"` // Shows the download speed from shard (Mbps)
}
type ResponseDevicesGetDeviceLiveToolsThroughputTest struct {
	Error            string                                                  `json:"error,omitempty"`            // Description of the error.
	Request          *ResponseDevicesGetDeviceLiveToolsThroughputTestRequest `json:"request,omitempty"`          // The parameters of the throughput test request
	Result           *ResponseDevicesGetDeviceLiveToolsThroughputTestResult  `json:"result,omitempty"`           // Result of the throughput test request
	Status           string                                                  `json:"status,omitempty"`           // Status of the throughput test request
	ThroughputTestID string                                                  `json:"throughputTestId,omitempty"` // ID of throughput test job
	URL              string                                                  `json:"url,omitempty"`              // GET this url to check the status of your throughput test request
}
type ResponseDevicesGetDeviceLiveToolsThroughputTestRequest struct {
	Serial string `json:"serial,omitempty"` // Device serial number
}
type ResponseDevicesGetDeviceLiveToolsThroughputTestResult struct {
	Speeds *ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeeds `json:"speeds,omitempty"` // Shows the speeds (Mbps)
}
type ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeeds struct {
	Downstream *int `json:"downstream,omitempty"` // Shows the download speed from shard (Mbps)
}
type ResponseDevicesCreateDeviceLiveToolsWakeOnLan struct {
	Callback    *ResponseDevicesCreateDeviceLiveToolsWakeOnLanCallback `json:"callback,omitempty"`    // Information for callback used to send back results
	Error       string                                                 `json:"error,omitempty"`       // An error message for a failed execution
	Request     *ResponseDevicesCreateDeviceLiveToolsWakeOnLanRequest  `json:"request,omitempty"`     // The parameters of the Wake-on-LAN request
	Status      string                                                 `json:"status,omitempty"`      // Status of the Wake-on-LAN request
	URL         string                                                 `json:"url,omitempty"`         // GET this url to check the status of your ping request
	WakeOnLanID string                                                 `json:"wakeOnLanId,omitempty"` // ID of the Wake-on-LAN job
}
type ResponseDevicesCreateDeviceLiveToolsWakeOnLanCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesCreateDeviceLiveToolsWakeOnLanRequest struct {
	Mac    string `json:"mac,omitempty"`    // The target's MAC address
	Serial string `json:"serial,omitempty"` // Device serial number
	VLANID *int   `json:"vlanId,omitempty"` // The target's VLAN (1 to 4094)
}
type ResponseDevicesGetDeviceLiveToolsWakeOnLan struct {
	Error       string                                             `json:"error,omitempty"`       // An error message for a failed execution
	Request     *ResponseDevicesGetDeviceLiveToolsWakeOnLanRequest `json:"request,omitempty"`     // The parameters of the Wake-on-LAN request
	Status      string                                             `json:"status,omitempty"`      // Status of the Wake-on-LAN request
	URL         string                                             `json:"url,omitempty"`         // GET this url to check the status of your ping request
	WakeOnLanID string                                             `json:"wakeOnLanId,omitempty"` // ID of the Wake-on-LAN job
}
type ResponseDevicesGetDeviceLiveToolsWakeOnLanRequest struct {
	Mac    string `json:"mac,omitempty"`    // The target's MAC address
	Serial string `json:"serial,omitempty"` // Device serial number
	VLANID *int   `json:"vlanId,omitempty"` // The target's VLAN (1 to 4094)
}
type ResponseDevicesGetDeviceLldpCdp struct {
	Ports     *ResponseDevicesGetDeviceLldpCdpPorts `json:"ports,omitempty"`     // Mapping of ports to lldp and/or cdp information
	SourceMac string                                `json:"sourceMac,omitempty"` // Source MAC address
}
type ResponseDevicesGetDeviceLldpCdpPorts struct {
	Status12 *ResponseDevicesGetDeviceLldpCdpPorts12 `json:"12,omitempty"` //
	Status8  *ResponseDevicesGetDeviceLldpCdpPorts8  `json:"8,omitempty"`  //
}
type ResponseDevicesGetDeviceLldpCdpPorts12 struct {
	Cdp  *ResponseDevicesGetDeviceLldpCdpPorts12Cdp  `json:"cdp,omitempty"`  //
	Lldp *ResponseDevicesGetDeviceLldpCdpPorts12Lldp `json:"lldp,omitempty"` //
}
type ResponseDevicesGetDeviceLldpCdpPorts12Cdp struct {
	Address    string `json:"address,omitempty"`    //
	DeviceID   string `json:"deviceId,omitempty"`   //
	PortID     string `json:"portId,omitempty"`     //
	SourcePort string `json:"sourcePort,omitempty"` //
}
type ResponseDevicesGetDeviceLldpCdpPorts12Lldp struct {
	ManagementAddress string `json:"managementAddress,omitempty"` //
	PortID            string `json:"portId,omitempty"`            //
	SourcePort        string `json:"sourcePort,omitempty"`        //
	SystemName        string `json:"systemName,omitempty"`        //
}
type ResponseDevicesGetDeviceLldpCdpPorts8 struct {
	Cdp *ResponseDevicesGetDeviceLldpCdpPorts8Cdp `json:"cdp,omitempty"` //
}
type ResponseDevicesGetDeviceLldpCdpPorts8Cdp struct {
	Address    string `json:"address,omitempty"`    //
	DeviceID   string `json:"deviceId,omitempty"`   //
	PortID     string `json:"portId,omitempty"`     //
	SourcePort string `json:"sourcePort,omitempty"` //
}
type ResponseDevicesGetDeviceLossAndLatencyHistory []ResponseItemDevicesGetDeviceLossAndLatencyHistory // Array of ResponseDevicesGetDeviceLossAndLatencyHistory
type ResponseItemDevicesGetDeviceLossAndLatencyHistory struct {
	EndTime     string   `json:"endTime,omitempty"`     // End time of the sample
	Goodput     *int     `json:"goodput,omitempty"`     // Number of useful information bits delivered
	Jitter      *float64 `json:"jitter,omitempty"`      // Jitter, in milliseconds
	LatencyMs   *float64 `json:"latencyMs,omitempty"`   // Latency in milliseconds
	LossPercent *float64 `json:"lossPercent,omitempty"` // Percentage of packets lost
	StartTime   string   `json:"startTime,omitempty"`   // Start time of the sample
}
type ResponseDevicesGetDeviceManagementInterface struct {
	DdnsHostnames *ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames `json:"ddnsHostnames,omitempty"` // Dynamic DNS hostnames.
	Wan1          *ResponseDevicesGetDeviceManagementInterfaceWan1          `json:"wan1,omitempty"`          // WAN 1 settings
	Wan2          *ResponseDevicesGetDeviceManagementInterfaceWan2          `json:"wan2,omitempty"`          // WAN 2 settings (only for MX devices)
}
type ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames struct {
	ActiveDdnsHostname string `json:"activeDdnsHostname,omitempty"` // Active dynamic DNS hostname.
	DdnsHostnameWan1   string `json:"ddnsHostnameWan1,omitempty"`   // WAN 1 dynamic DNS hostname.
	DdnsHostnameWan2   string `json:"ddnsHostnameWan2,omitempty"`   // WAN 2 dynamic DNS hostname.
}
type ResponseDevicesGetDeviceManagementInterfaceWan1 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type ResponseDevicesGetDeviceManagementInterfaceWan2 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type ResponseDevicesUpdateDeviceManagementInterface struct {
	DdnsHostnames *ResponseDevicesUpdateDeviceManagementInterfaceDdnsHostnames `json:"ddnsHostnames,omitempty"` // Dynamic DNS hostnames.
	Wan1          *ResponseDevicesUpdateDeviceManagementInterfaceWan1          `json:"wan1,omitempty"`          // WAN 1 settings
	Wan2          *ResponseDevicesUpdateDeviceManagementInterfaceWan2          `json:"wan2,omitempty"`          // WAN 2 settings (only for MX devices)
}
type ResponseDevicesUpdateDeviceManagementInterfaceDdnsHostnames struct {
	ActiveDdnsHostname string `json:"activeDdnsHostname,omitempty"` // Active dynamic DNS hostname.
	DdnsHostnameWan1   string `json:"ddnsHostnameWan1,omitempty"`   // WAN 1 dynamic DNS hostname.
	DdnsHostnameWan2   string `json:"ddnsHostnameWan2,omitempty"`   // WAN 2 dynamic DNS hostname.
}
type ResponseDevicesUpdateDeviceManagementInterfaceWan1 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type ResponseDevicesUpdateDeviceManagementInterfaceWan2 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type ResponseDevicesRebootDevice struct {
	Success *bool `json:"success,omitempty"` // Shows the success of the reboot
}
type ResponseDevicesGetNetworkDevices []ResponseItemDevicesGetNetworkDevices // Array of ResponseDevicesGetNetworkDevices
type ResponseItemDevicesGetNetworkDevices struct {
	Address     string                                         `json:"address,omitempty"`     // Physical address of the device
	Details     *[]ResponseItemDevicesGetNetworkDevicesDetails `json:"details,omitempty"`     // Additional device information
	Firmware    string                                         `json:"firmware,omitempty"`    // Firmware version of the device
	Imei        string                                         `json:"imei,omitempty"`        // IMEI of the device, if applicable
	LanIP       string                                         `json:"lanIp,omitempty"`       // LAN IP address of the device
	Lat         *float64                                       `json:"lat,omitempty"`         // Latitude of the device
	Lng         *float64                                       `json:"lng,omitempty"`         // Longitude of the device
	Mac         string                                         `json:"mac,omitempty"`         // MAC address of the device
	Model       string                                         `json:"model,omitempty"`       // Model of the device
	Name        string                                         `json:"name,omitempty"`        // Name of the device
	NetworkID   string                                         `json:"networkId,omitempty"`   // ID of the network the device belongs to
	Notes       string                                         `json:"notes,omitempty"`       // Notes for the device, limited to 255 characters
	ProductType string                                         `json:"productType,omitempty"` // Product type of the device
	Serial      string                                         `json:"serial,omitempty"`      // Serial number of the device
	Tags        []string                                       `json:"tags,omitempty"`        // List of tags assigned to the device
}
type ResponseItemDevicesGetNetworkDevicesDetails struct {
	Name  string `json:"name,omitempty"`  // Additional property name
	Value string `json:"value,omitempty"` // Additional property value
}
type ResponseDevicesClaimNetworkDevices struct {
	Serials []string `json:"serials,omitempty"` // The serials of the devices
}
type ResponseDevicesVmxNetworkDevicesClaim struct {
	Address     string                                          `json:"address,omitempty"`     // Physical address of the device
	Details     *[]ResponseDevicesVmxNetworkDevicesClaimDetails `json:"details,omitempty"`     // Additional device information
	Firmware    string                                          `json:"firmware,omitempty"`    // Firmware version of the device
	Imei        string                                          `json:"imei,omitempty"`        // IMEI of the device, if applicable
	LanIP       string                                          `json:"lanIp,omitempty"`       // LAN IP address of the device
	Lat         *float64                                        `json:"lat,omitempty"`         // Latitude of the device
	Lng         *float64                                        `json:"lng,omitempty"`         // Longitude of the device
	Mac         string                                          `json:"mac,omitempty"`         // MAC address of the device
	Model       string                                          `json:"model,omitempty"`       // Model of the device
	Name        string                                          `json:"name,omitempty"`        // Name of the device
	NetworkID   string                                          `json:"networkId,omitempty"`   // ID of the network the device belongs to
	Notes       string                                          `json:"notes,omitempty"`       // Notes for the device, limited to 255 characters
	ProductType string                                          `json:"productType,omitempty"` // Product type of the device
	Serial      string                                          `json:"serial,omitempty"`      // Serial number of the device
	Tags        []string                                        `json:"tags,omitempty"`        // List of tags assigned to the device
}
type ResponseDevicesVmxNetworkDevicesClaimDetails struct {
	Name  string `json:"name,omitempty"`  // Additional property name
	Value string `json:"value,omitempty"` // Additional property value
}
type ResponseDevicesCheckinNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of devices.
}
type ResponseDevicesUpdateNetworkSmDevicesFields []ResponseItemDevicesUpdateNetworkSmDevicesFields // Array of ResponseDevicesUpdateNetworkSmDevicesFields
type ResponseItemDevicesUpdateNetworkSmDevicesFields struct {
	ID      string `json:"id,omitempty"`      // The Meraki Id of the device record.
	Name    string `json:"name,omitempty"`    // The name of the device.
	Notes   string `json:"notes,omitempty"`   // Notes associated with the device.
	Serial  string `json:"serial,omitempty"`  // The device serial.
	WifiMac string `json:"wifiMac,omitempty"` // The MAC of the device.
}
type ResponseDevicesLockNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of devices.
}
type ResponseDevicesModifyNetworkSmDevicesTags []ResponseItemDevicesModifyNetworkSmDevicesTags // Array of ResponseDevicesModifyNetworkSmDevicesTags
type ResponseItemDevicesModifyNetworkSmDevicesTags struct {
	ID      string   `json:"id,omitempty"`      // The Meraki Id of the device record.
	Serial  string   `json:"serial,omitempty"`  // The device serial.
	Tags    []string `json:"tags,omitempty"`    // An array of tags associated with the device.
	WifiMac string   `json:"wifiMac,omitempty"` // The MAC of the device.
}
type ResponseDevicesMoveNetworkSmDevices struct {
	IDs        []string `json:"ids,omitempty"`        // The Meraki Ids of the set of devices.
	NewNetwork string   `json:"newNetwork,omitempty"` // The network to which the devices was moved.
}
type ResponseDevicesRebootNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of endpoints.
}
type ResponseDevicesShutdownNetworkSmDevices struct {
	IDs []string `json:"ids,omitempty"` // The Meraki Ids of the set of endpoints.
}
type ResponseDevicesWipeNetworkSmDevices struct {
	ID string `json:"id,omitempty"` // The Meraki Id of the devices.
}
type ResponseDevicesGetNetworkSmDeviceCellularUsageHistory []ResponseItemDevicesGetNetworkSmDeviceCellularUsageHistory // Array of ResponseDevicesGetNetworkSmDeviceCellularUsageHistory
type ResponseItemDevicesGetNetworkSmDeviceCellularUsageHistory struct {
	Received *float64 `json:"received,omitempty"` // The amount of cellular data received by the device.
	Sent     *float64 `json:"sent,omitempty"`     // The amount of cellular sent received by the device.
	Ts       string   `json:"ts,omitempty"`       // When the cellular usage data was collected.
}
type ResponseDevicesGetNetworkSmDeviceCerts []ResponseItemDevicesGetNetworkSmDeviceCerts // Array of ResponseDevicesGetNetworkSmDeviceCerts
type ResponseItemDevicesGetNetworkSmDeviceCerts struct {
	CertPem        string `json:"certPem,omitempty"`        // The PEM of the certificate.
	DeviceID       string `json:"deviceId,omitempty"`       // The Meraki managed device Id.
	ID             string `json:"id,omitempty"`             // The Meraki Id of the certificate record.
	Issuer         string `json:"issuer,omitempty"`         // The certificate issuer.
	Name           string `json:"name,omitempty"`           // The name of the certificate.
	NotValidAfter  string `json:"notValidAfter,omitempty"`  // The date after which the certificate is no longer valid.
	NotValidBefore string `json:"notValidBefore,omitempty"` // The date before which the certificate is not valid.
	Subject        string `json:"subject,omitempty"`        // The subject of the certificate.
}
type ResponseDevicesGetNetworkSmDeviceDeviceProfiles []ResponseItemDevicesGetNetworkSmDeviceDeviceProfiles // Array of ResponseDevicesGetNetworkSmDeviceDeviceProfiles
type ResponseItemDevicesGetNetworkSmDeviceDeviceProfiles struct {
	DeviceID          string `json:"deviceId,omitempty"`          // The Meraki managed device Id.
	ID                string `json:"id,omitempty"`                // The numerical Meraki Id of the profile.
	IsEncrypted       *bool  `json:"isEncrypted,omitempty"`       // A boolean indicating if the profile is encrypted.
	IsManaged         *bool  `json:"isManaged,omitempty"`         // Whether or not the profile is managed by Meraki.
	Name              string `json:"name,omitempty"`              // The name of the profile.
	ProfileData       string `json:"profileData,omitempty"`       // A string containing a JSON object with the profile data.
	ProfileIDentifier string `json:"profileIdentifier,omitempty"` // The identifier of the profile.
	Version           string `json:"version,omitempty"`           // The verison of the profile.
}
type ResponseDevicesGetNetworkSmDeviceNetworkAdapters []ResponseItemDevicesGetNetworkSmDeviceNetworkAdapters // Array of ResponseDevicesGetNetworkSmDeviceNetworkAdapters
type ResponseItemDevicesGetNetworkSmDeviceNetworkAdapters struct {
	DhcpServer string `json:"dhcpServer,omitempty"` // The IP address of the DCHP Server.
	DNSServer  string `json:"dnsServer,omitempty"`  // The IP address of the DNS Server.
	Gateway    string `json:"gateway,omitempty"`    // The IP address of the Gateway.
	ID         string `json:"id,omitempty"`         // The Meraki Id of the network adapter record.
	IP         string `json:"ip,omitempty"`         // The IP address of the network adapter.
	Mac        string `json:"mac,omitempty"`        // The MAC associated with the network adapter.
	Name       string `json:"name,omitempty"`       // The name of the newtwork adapter.
	Subnet     string `json:"subnet,omitempty"`     // The subnet for the network adapter.
}
type ResponseDevicesGetNetworkSmDeviceRestrictions []ResponseItemDevicesGetNetworkSmDeviceRestrictions // Array of ResponseDevicesGetNetworkSmDeviceRestrictions
type ResponseItemDevicesGetNetworkSmDeviceRestrictions struct {
	Profile      string                                                         `json:"profile,omitempty"`      //
	Restrictions *ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictions `json:"restrictions,omitempty"` //
}
type ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictions struct {
	MyRestriction *ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictionsMyRestriction `json:"myRestriction,omitempty"` //
}
type ResponseItemDevicesGetNetworkSmDeviceRestrictionsRestrictionsMyRestriction struct {
	Value *bool `json:"value,omitempty"` //
}
type ResponseDevicesGetNetworkSmDeviceSecurityCenters []ResponseItemDevicesGetNetworkSmDeviceSecurityCenters // Array of ResponseDevicesGetNetworkSmDeviceSecurityCenters
type ResponseItemDevicesGetNetworkSmDeviceSecurityCenters struct {
	AntiVirusName        string `json:"antiVirusName,omitempty"`        // The name of the Antivirus.
	FireWallName         string `json:"fireWallName,omitempty"`         // The name of the Firewall.
	HasAntiVirus         *bool  `json:"hasAntiVirus,omitempty"`         // Boolean indicating if the device has Antivirus.
	HasFireWallInstalled *bool  `json:"hasFireWallInstalled,omitempty"` // Boolean indicating if the device has a Firewall installed.
	ID                   string `json:"id,omitempty"`                   // The Meraki identifier for the security center record.
	IsAutoLoginDisabled  *bool  `json:"isAutoLoginDisabled,omitempty"`  // Boolean indicating if the device has auto login disabled.
	IsDiskEncrypted      *bool  `json:"isDiskEncrypted,omitempty"`      // Boolean indicating if the device has disk encryption.
	IsFireWallEnabled    *bool  `json:"isFireWallEnabled,omitempty"`    // Boolean indicating if the device has a Firewall enabled.
	IsRooted             *bool  `json:"isRooted,omitempty"`             // Boolean indicating if the device is rooted.
	RunningProcs         string `json:"runningProcs,omitempty"`         // A comma seperated list of procs running on the device.
}
type ResponseDevicesGetNetworkSmDeviceSoftwares []ResponseItemDevicesGetNetworkSmDeviceSoftwares // Array of ResponseDevicesGetNetworkSmDeviceSoftwares
type ResponseItemDevicesGetNetworkSmDeviceSoftwares struct {
	AppID             string `json:"appId,omitempty"`             // The Meraki managed application Id for this record on a particular device.
	BundleSize        *int   `json:"bundleSize,omitempty"`        // The size of the software bundle.
	CreatedAt         string `json:"createdAt,omitempty"`         // When the Meraki record for the software was created.
	DeviceID          string `json:"deviceId,omitempty"`          // The Meraki managed device Id.
	DynamicSize       *int   `json:"dynamicSize,omitempty"`       // The size of the data stored in the application.
	ID                string `json:"id,omitempty"`                // The Meraki software Id.
	IDentifier        string `json:"identifier,omitempty"`        // Software bundle identifier.
	InstalledAt       string `json:"installedAt,omitempty"`       // When the Software was installed on the device.
	IosRedemptionCode *bool  `json:"iosRedemptionCode,omitempty"` // A boolean indicating whether or not an iOS redemption code was used.
	IsManaged         *bool  `json:"isManaged,omitempty"`         // A boolean indicating whether or not the software is managed by Meraki.
	ItunesID          string `json:"itunesId,omitempty"`          // The itunes numerical identifier.
	LicenseKey        string `json:"licenseKey,omitempty"`        // The license key associated with this software installation.
	Name              string `json:"name,omitempty"`              // The name of the software.
	Path              string `json:"path,omitempty"`              // The path on the device where the software record is located.
	RedemptionCode    *int   `json:"redemptionCode,omitempty"`    // The redemption code used for this software.
	ShortVersion      string `json:"shortVersion,omitempty"`      // Short version notation for the software.
	Status            string `json:"status,omitempty"`            // The management status of the software.
	ToInstall         *bool  `json:"toInstall,omitempty"`         // A boolean indicating this software record should be installed on the associated device.
	ToUninstall       *bool  `json:"toUninstall,omitempty"`       // A boolean indicating this software record should be uninstalled on the associated device.
	UninstalledAt     string `json:"uninstalledAt,omitempty"`     // When the record was uninstalled from the device.
	UpdatedAt         string `json:"updatedAt,omitempty"`         // When the record was last updated by Meraki.
	Vendor            string `json:"vendor,omitempty"`            // The vendor of the software.
	Version           string `json:"version,omitempty"`           // Full version notation for the software.
}
type ResponseDevicesUnenrollNetworkSmDevice struct {
	Success *bool `json:"success,omitempty"` // Boolean indicating whether the operation was completed successfully.
}
type ResponseDevicesGetNetworkSmDeviceWLANLists []ResponseItemDevicesGetNetworkSmDeviceWLANLists // Array of ResponseDevicesGetNetworkSmDeviceWlanLists
type ResponseItemDevicesGetNetworkSmDeviceWLANLists struct {
	CreatedAt string `json:"createdAt,omitempty"` // When the Meraki record for the wlanList was created.
	ID        string `json:"id,omitempty"`        // The Meraki managed Id of the wlanList record.
	Xml       string `json:"xml,omitempty"`       // An XML string containing the WLAN List for the device.
}

type ResponseOrganizationsCloneOrganizationSwitchDevices interface{}
type ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDevice []ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDevice // Array of ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDevice
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDevice struct {
	ByBand  *[]ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	Mac     string                                                                                     `json:"mac,omitempty"`     // The MAC address of the device.
	Network *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	Serial  string                                                                                     `json:"serial,omitempty"`  // The serial number for the device.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand struct {
	Band    string                                                                                          `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetwork []ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetwork // Array of ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetwork
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetwork struct {
	ByBand  *[]ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	Network *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand struct {
	Band    string                                                                                           `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval []ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval // Array of ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval struct {
	ByBand  *[]ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	EndTs   string                                                                                                      `json:"endTs,omitempty"`   // The end time of the channel utilization interval.
	Mac     string                                                                                                      `json:"mac,omitempty"`     // The MAC address of the device.
	Network *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	Serial  string                                                                                                      `json:"serial,omitempty"`  // The serial number for the device.
	StartTs string                                                                                                      `json:"startTs,omitempty"` // The start time of the channel utilization interval.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBand struct {
	Band    string                                                                                                           `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval []ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval // Array of ResponseOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval struct {
	ByBand  *[]ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	EndTs   string                                                                                                       `json:"endTs,omitempty"`   // The end time of the channel utilization interval.
	Network *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	StartTs string                                                                                                       `json:"startTs,omitempty"` // The start time of the channel utilization interval.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBand struct {
	Band    string                                                                                                            `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseOrganizationsGetOrganizationWirelessDevicesPacketLossByClient []ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClient // Array of ResponseOrganizationsGetOrganizationWirelessDevicesPacketLossByClient
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClient struct {
	Client     *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientClient     `json:"client,omitempty"`     // Client.
	Downstream *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientClient struct {
	ID  string `json:"id,omitempty"`  // Client ID.
	Mac string `json:"mac,omitempty"` // MAC address.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByClientUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type ResponseOrganizationsGetOrganizationWirelessDevicesPacketLossByDevice []ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDevice // Array of ResponseOrganizationsGetOrganizationWirelessDevicesPacketLossByDevice
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDevice struct {
	Device     *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceDevice     `json:"device,omitempty"`     // Device.
	Downstream *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceDevice struct {
	Mac    string `json:"mac,omitempty"`    // MAC address
	Name   string `json:"name,omitempty"`   // Name
	Serial string `json:"serial,omitempty"` // Serial Number
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByDeviceUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type ResponseOrganizationsGetOrganizationWirelessDevicesPacketLossByNetwork []ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetwork // Array of ResponseOrganizationsGetOrganizationWirelessDevicesPacketLossByNetwork
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetwork struct {
	Downstream *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetworkDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetworkNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetworkUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetworkDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetworkNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemOrganizationsGetOrganizationWirelessDevicesPacketLossByNetworkUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type RequestDevicesUpdateDevice struct {
	Address         string   `json:"address,omitempty"`         // The address of a device
	FloorPlanID     string   `json:"floorPlanId,omitempty"`     // The floor plan to associate to this device. null disassociates the device from the floorplan.
	Lat             *float64 `json:"lat,omitempty"`             // The latitude of a device
	Lng             *float64 `json:"lng,omitempty"`             // The longitude of a device
	MoveMapMarker   *bool    `json:"moveMapMarker,omitempty"`   // Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
	Name            string   `json:"name,omitempty"`            // The name of a device
	Notes           string   `json:"notes,omitempty"`           // The notes for the device. String. Limited to 255 characters.
	SwitchProfileID string   `json:"switchProfileId,omitempty"` // The ID of a switch template to bind to the device (for available switch templates, see the 'Switch Templates' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
	Tags            []string `json:"tags,omitempty"`            // The list of tags of a device
}
type RequestDevicesBlinkDeviceLeds struct {
	Duration *int `json:"duration,omitempty"` // The duration in seconds. Must be between 5 and 120. Default is 20 seconds
	Duty     *int `json:"duty,omitempty"`     // The duty cycle as the percent active. Must be between 10 and 90. Default is 50.
	Period   *int `json:"period,omitempty"`   // The period in milliseconds. Must be between 100 and 1000. Default is 160 milliseconds
}
type RequestDevicesUpdateDeviceCellularSims struct {
	SimFailover *RequestDevicesUpdateDeviceCellularSimsSimFailover `json:"simFailover,omitempty"` // SIM Failover settings.
	Sims        *[]RequestDevicesUpdateDeviceCellularSimsSims      `json:"sims,omitempty"`        // List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.
}
type RequestDevicesUpdateDeviceCellularSimsSimFailover struct {
	Enabled *bool `json:"enabled,omitempty"` // Failover to secondary SIM (optional)
	Timeout *int  `json:"timeout,omitempty"` // Failover timeout in seconds (optional)
}
type RequestDevicesUpdateDeviceCellularSimsSims struct {
	Apns      *[]RequestDevicesUpdateDeviceCellularSimsSimsApns `json:"apns,omitempty"`      // APN configurations. If empty, the default APN will be used.
	IsPrimary *bool                                             `json:"isPrimary,omitempty"` // If true, this SIM is used for boot. Must be true on single-sim devices.
	Slot      string                                            `json:"slot,omitempty"`      // SIM slot being configured. Must be 'sim1' on single-sim devices.
}
type RequestDevicesUpdateDeviceCellularSimsSimsApns struct {
	AllowedIPTypes []string                                                      `json:"allowedIpTypes,omitempty"` // IP versions to support (permitted values include 'ipv4', 'ipv6').
	Authentication *RequestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication `json:"authentication,omitempty"` // APN authentication configurations.
	Name           string                                                        `json:"name,omitempty"`           // APN name.
}
type RequestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication struct {
	Password string `json:"password,omitempty"` // APN password, if type is set (if APN password is not supplied, the password is left unchanged).
	Type     string `json:"type,omitempty"`     // APN auth type.
	Username string `json:"username,omitempty"` // APN username, if type is set.
}
type RequestDevicesCreateDeviceLiveToolsArpTable struct {
	Callback *RequestDevicesCreateDeviceLiveToolsArpTableCallback `json:"callback,omitempty"` // Details for the callback. Please include either an httpServerId OR url and sharedSecret
}
type RequestDevicesCreateDeviceLiveToolsArpTableCallback struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsArpTableCallbackHTTPServer      `json:"httpServer,omitempty"`      // The webhook receiver used for the callback webhook.
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template of the webhook used for the callback
	SharedSecret    string                                                              `json:"sharedSecret,omitempty"`    // A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	URL             string                                                              `json:"url,omitempty"`             // The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
}
type RequestDevicesCreateDeviceLiveToolsArpTableCallbackHTTPServer struct {
	ID string `json:"id,omitempty"` // The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
}
type RequestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplate struct {
	ID string `json:"id,omitempty"` // The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.
}
type RequestDevicesCreateDeviceLiveToolsCableTest struct {
	Callback *RequestDevicesCreateDeviceLiveToolsCableTestCallback `json:"callback,omitempty"` // Details for the callback. Please include either an httpServerId OR url and sharedSecret
	Ports    []string                                              `json:"ports,omitempty"`    // A list of ports for which to perform the cable test.
}
type RequestDevicesCreateDeviceLiveToolsCableTestCallback struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsCableTestCallbackHTTPServer      `json:"httpServer,omitempty"`      // The webhook receiver used for the callback webhook.
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template of the webhook used for the callback
	SharedSecret    string                                                               `json:"sharedSecret,omitempty"`    // A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	URL             string                                                               `json:"url,omitempty"`             // The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
}
type RequestDevicesCreateDeviceLiveToolsCableTestCallbackHTTPServer struct {
	ID string `json:"id,omitempty"` // The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
}
type RequestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplate struct {
	ID string `json:"id,omitempty"` // The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.
}
type RequestDevicesCreateDeviceLiveToolsPing struct {
	Callback *RequestDevicesCreateDeviceLiveToolsPingCallback `json:"callback,omitempty"` // Details for the callback. Please include either an httpServerId OR url and sharedSecret
	Count    *int                                             `json:"count,omitempty"`    // Count parameter to pass to ping. [1..5], default 5
	Target   string                                           `json:"target,omitempty"`   // FQDN, IPv4 or IPv6 address
}
type RequestDevicesCreateDeviceLiveToolsPingCallback struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsPingCallbackHTTPServer      `json:"httpServer,omitempty"`      // The webhook receiver used for the callback webhook.
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsPingCallbackPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template of the webhook used for the callback
	SharedSecret    string                                                          `json:"sharedSecret,omitempty"`    // A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	URL             string                                                          `json:"url,omitempty"`             // The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
}
type RequestDevicesCreateDeviceLiveToolsPingCallbackHTTPServer struct {
	ID string `json:"id,omitempty"` // The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
}
type RequestDevicesCreateDeviceLiveToolsPingCallbackPayloadTemplate struct {
	ID string `json:"id,omitempty"` // The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.
}
type RequestDevicesCreateDeviceLiveToolsPingDevice struct {
	Callback *RequestDevicesCreateDeviceLiveToolsPingDeviceCallback `json:"callback,omitempty"` // Details for the callback. Please include either an httpServerId OR url and sharedSecret
	Count    *int                                                   `json:"count,omitempty"`    // Count parameter to pass to ping. [1..5], default 5
}
type RequestDevicesCreateDeviceLiveToolsPingDeviceCallback struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsPingDeviceCallbackHTTPServer      `json:"httpServer,omitempty"`      // The webhook receiver used for the callback webhook.
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsPingDeviceCallbackPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template of the webhook used for the callback
	SharedSecret    string                                                                `json:"sharedSecret,omitempty"`    // A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	URL             string                                                                `json:"url,omitempty"`             // The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
}
type RequestDevicesCreateDeviceLiveToolsPingDeviceCallbackHTTPServer struct {
	ID string `json:"id,omitempty"` // The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
}
type RequestDevicesCreateDeviceLiveToolsPingDeviceCallbackPayloadTemplate struct {
	ID string `json:"id,omitempty"` // The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.
}
type RequestDevicesCreateDeviceLiveToolsThroughputTest struct {
	Callback *RequestDevicesCreateDeviceLiveToolsThroughputTestCallback `json:"callback,omitempty"` // Details for the callback. Please include either an httpServerId OR url and sharedSecret
}
type RequestDevicesCreateDeviceLiveToolsThroughputTestCallback struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackHTTPServer      `json:"httpServer,omitempty"`      // The webhook receiver used for the callback webhook.
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template of the webhook used for the callback
	SharedSecret    string                                                                    `json:"sharedSecret,omitempty"`    // A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	URL             string                                                                    `json:"url,omitempty"`             // The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
}
type RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackHTTPServer struct {
	ID string `json:"id,omitempty"` // The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
}
type RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplate struct {
	ID string `json:"id,omitempty"` // The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.
}
type RequestDevicesCreateDeviceLiveToolsWakeOnLan struct {
	Callback *RequestDevicesCreateDeviceLiveToolsWakeOnLanCallback `json:"callback,omitempty"` // Details for the callback. Please include either an httpServerId OR url and sharedSecret
	Mac      string                                                `json:"mac,omitempty"`      // The target's MAC address
	VLANID   *int                                                  `json:"vlanId,omitempty"`   // The target's VLAN (1 to 4094)
}
type RequestDevicesCreateDeviceLiveToolsWakeOnLanCallback struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHTTPServer      `json:"httpServer,omitempty"`      // The webhook receiver used for the callback webhook.
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template of the webhook used for the callback
	SharedSecret    string                                                               `json:"sharedSecret,omitempty"`    // A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	URL             string                                                               `json:"url,omitempty"`             // The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
}
type RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHTTPServer struct {
	ID string `json:"id,omitempty"` // The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
}
type RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplate struct {
	ID string `json:"id,omitempty"` // The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.
}
type RequestDevicesUpdateDeviceManagementInterface struct {
	Wan1 *RequestDevicesUpdateDeviceManagementInterfaceWan1 `json:"wan1,omitempty"` // WAN 1 settings
	Wan2 *RequestDevicesUpdateDeviceManagementInterfaceWan2 `json:"wan2,omitempty"` // WAN 2 settings (only for MX devices)
}
type RequestDevicesUpdateDeviceManagementInterfaceWan1 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type RequestDevicesUpdateDeviceManagementInterfaceWan2 struct {
	StaticDNS        []string `json:"staticDns,omitempty"`        // Up to two DNS IPs.
	StaticGatewayIP  string   `json:"staticGatewayIp,omitempty"`  // The IP of the gateway on the WAN.
	StaticIP         string   `json:"staticIp,omitempty"`         // The IP the device should use on the WAN.
	StaticSubnetMask string   `json:"staticSubnetMask,omitempty"` // The subnet mask for the WAN.
	UsingStaticIP    *bool    `json:"usingStaticIp,omitempty"`    // Configure the interface to have static IP settings or use DHCP.
	VLAN             *int     `json:"vlan,omitempty"`             // The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	WanEnabled       string   `json:"wanEnabled,omitempty"`       // Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
}
type RequestDevicesClaimNetworkDevices struct {
	Serials []string `json:"serials,omitempty"` // A list of serials of devices to claim
}
type RequestDevicesVmxNetworkDevicesClaim struct {
	Size string `json:"size,omitempty"` // The size of the vMX you claim. It can be one of: small, medium, large, xlarge, 100
}
type RequestDevicesRemoveNetworkDevices struct {
	Serial string `json:"serial,omitempty"` // The serial of a device
}
type RequestDevicesCheckinNetworkSmDevices struct {
	IDs      []string `json:"ids,omitempty"`      // The ids of the devices to be checked-in.
	Scope    []string `json:"scope,omitempty"`    // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be checked-in.
	Serials  []string `json:"serials,omitempty"`  // The serials of the devices to be checked-in.
	WifiMacs []string `json:"wifiMacs,omitempty"` // The wifiMacs of the devices to be checked-in.
}
type RequestDevicesUpdateNetworkSmDevicesFields struct {
	DeviceFields *RequestDevicesUpdateNetworkSmDevicesFieldsDeviceFields `json:"deviceFields,omitempty"` // The new fields of the device. Each field of this object is optional.
	ID           string                                                  `json:"id,omitempty"`           // The id of the device to be modified.
	Serial       string                                                  `json:"serial,omitempty"`       // The serial of the device to be modified.
	WifiMac      string                                                  `json:"wifiMac,omitempty"`      // The wifiMac of the device to be modified.
}
type RequestDevicesUpdateNetworkSmDevicesFieldsDeviceFields struct {
	Name  string `json:"name,omitempty"`  // New name for the device
	Notes string `json:"notes,omitempty"` // New notes for the device
}
type RequestDevicesLockNetworkSmDevices struct {
	IDs      []string `json:"ids,omitempty"`      // The ids of the devices to be locked.
	Pin      *int     `json:"pin,omitempty"`      // The pin number for locking macOS devices (a six digit number). Required only for macOS devices.
	Scope    []string `json:"scope,omitempty"`    // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be locked.
	Serials  []string `json:"serials,omitempty"`  // The serials of the devices to be locked.
	WifiMacs []string `json:"wifiMacs,omitempty"` // The wifiMacs of the devices to be locked.
}
type RequestDevicesModifyNetworkSmDevicesTags struct {
	IDs          []string `json:"ids,omitempty"`          // The ids of the devices to be modified.
	Scope        []string `json:"scope,omitempty"`        // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified.
	Serials      []string `json:"serials,omitempty"`      // The serials of the devices to be modified.
	Tags         []string `json:"tags,omitempty"`         // The tags to be added, deleted, or updated.
	UpdateAction string   `json:"updateAction,omitempty"` // One of add, delete, or update. Only devices that have been modified will be returned.
	WifiMacs     []string `json:"wifiMacs,omitempty"`     // The wifiMacs of the devices to be modified.
}
type RequestDevicesMoveNetworkSmDevices struct {
	IDs        []string `json:"ids,omitempty"`        // The ids of the devices to be moved.
	NewNetwork string   `json:"newNetwork,omitempty"` // The new network to which the devices will be moved.
	Scope      []string `json:"scope,omitempty"`      // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be moved.
	Serials    []string `json:"serials,omitempty"`    // The serials of the devices to be moved.
	WifiMacs   []string `json:"wifiMacs,omitempty"`   // The wifiMacs of the devices to be moved.
}
type RequestDevicesRebootNetworkSmDevices struct {
	IDs                          []string `json:"ids,omitempty"`                          // The ids of the endpoints to be rebooted.
	KextPaths                    []string `json:"kextPaths,omitempty"`                    // The KextPaths of the endpoints to be rebooted. Available for macOS 11+
	NotifyUser                   *bool    `json:"notifyUser,omitempty"`                   // Whether or not to notify the user before rebooting the endpoint. Available for macOS 11.3+
	RebuildKernelCache           *bool    `json:"rebuildKernelCache,omitempty"`           // Whether or not to rebuild the kernel cache when rebooting the endpoint. Available for macOS 11+
	RequestRequiresNetworkTether *bool    `json:"requestRequiresNetworkTether,omitempty"` // Whether or not the request requires network tethering. Available for macOS and supervised iOS or tvOS
	Scope                        []string `json:"scope,omitempty"`                        // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the endpoints to be rebooted.
	Serials                      []string `json:"serials,omitempty"`                      // The serials of the endpoints to be rebooted.
	WifiMacs                     []string `json:"wifiMacs,omitempty"`                     // The wifiMacs of the endpoints to be rebooted.
}
type RequestDevicesShutdownNetworkSmDevices struct {
	IDs      []string `json:"ids,omitempty"`      // The ids of the endpoints to be shutdown.
	Scope    []string `json:"scope,omitempty"`    // The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the endpoints to be shutdown.
	Serials  []string `json:"serials,omitempty"`  // The serials of the endpoints to be shutdown.
	WifiMacs []string `json:"wifiMacs,omitempty"` // The wifiMacs of the endpoints to be shutdown.
}
type RequestDevicesWipeNetworkSmDevices struct {
	ID      string `json:"id,omitempty"`      // The id of the device to be wiped.
	Pin     *int   `json:"pin,omitempty"`     // The pin number (a six digit value) for wiping a macOS device. Required only for macOS devices.
	Serial  string `json:"serial,omitempty"`  // The serial of the device to be wiped.
	WifiMac string `json:"wifiMac,omitempty"` // The wifiMac of the device to be wiped.
}
type RequestDevicesInstallNetworkSmDeviceApps struct {
	AppIDs []string `json:"appIds,omitempty"` // ids of applications to be installed
	Force  *bool    `json:"force,omitempty"`  // By default, installation of an app which is believed to already be present on the device will be skipped. If you'd like to force the installation of the app, set this parameter to true.
}
type RequestDevicesUninstallNetworkSmDeviceApps struct {
	AppIDs []string `json:"appIds,omitempty"` // ids of applications to be uninstalled
}
type RequestOrganizationsCloneOrganizationSwitchDevices struct {
	SourceSerial  string   `json:"sourceSerial,omitempty"`  // Serial number of the source switch (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials,omitempty"` // Array of serial numbers of one or more target switches (must be on a network not bound to a template)
}

//GetDevice Return a single device
/* Return a single device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device
*/
func (s *DevicesService) GetDevice(serial string) (*ResponseDevicesGetDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDevice")
	}

	result := response.Result().(*ResponseDevicesGetDevice)
	return result, response, err

}

//GetDeviceCellularSims Return the SIM and APN configurations for a cellular device.
/* Return the SIM and APN configurations for a cellular device.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-cellular-sims
*/
func (s *DevicesService) GetDeviceCellularSims(serial string) (*ResponseDevicesGetDeviceCellularSims, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellular/sims"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceCellularSims{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCellularSims")
	}

	result := response.Result().(*ResponseDevicesGetDeviceCellularSims)
	return result, response, err

}

//GetDeviceClients List the clients of a device, up to a maximum of a month ago
/* List the clients of a device, up to a maximum of a month ago. The usage of each client is returned in kilobytes. If the device is a switch, the switchport is returned; otherwise the switchport field is null.

@param serial serial path parameter.
@param getDeviceClientsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-clients
*/
func (s *DevicesService) GetDeviceClients(serial string, getDeviceClientsQueryParams *GetDeviceClientsQueryParams) (*ResponseDevicesGetDeviceClients, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/clients"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceClientsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceClients{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceClients")
	}

	result := response.Result().(*ResponseDevicesGetDeviceClients)
	return result, response, err

}

//GetDeviceLiveToolsArpTable Return an ARP table live tool job.
/* Return an ARP table live tool job.

@param serial serial path parameter.
@param arpTableID arpTableId path parameter. Arp table ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-arp-table
*/
func (s *DevicesService) GetDeviceLiveToolsArpTable(serial string, arpTableID string) (*ResponseDevicesGetDeviceLiveToolsArpTable, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/arpTable/{arpTableId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{arpTableId}", fmt.Sprintf("%v", arpTableID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsArpTable{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsArpTable")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsArpTable)
	return result, response, err

}

//GetDeviceLiveToolsCableTest Return a cable test live tool job.
/* Return a cable test live tool job.

@param serial serial path parameter.
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-cable-test
*/
func (s *DevicesService) GetDeviceLiveToolsCableTest(serial string, id string) (*ResponseDevicesGetDeviceLiveToolsCableTest, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/cableTest/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsCableTest{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsCableTest")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsCableTest)
	return result, response, err

}

//GetDeviceLiveToolsPing Return a ping job
/* Return a ping job. Latency unit in response is in milliseconds. Size is in bytes.

@param serial serial path parameter.
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-ping
*/
func (s *DevicesService) GetDeviceLiveToolsPing(serial string, id string) (*ResponseDevicesGetDeviceLiveToolsPing, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/ping/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsPing{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsPing")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsPing)
	return result, response, err

}

//GetDeviceLiveToolsPingDevice Return a ping device job
/* Return a ping device job. Latency unit in response is in milliseconds. Size is in bytes.

@param serial serial path parameter.
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-ping-device
*/
func (s *DevicesService) GetDeviceLiveToolsPingDevice(serial string, id string) (*ResponseDevicesGetDeviceLiveToolsPingDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/pingDevice/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsPingDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsPingDevice")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsPingDevice)
	return result, response, err

}

//GetDeviceLiveToolsThroughputTest Return a throughput test job
/* Return a throughput test job

@param serial serial path parameter.
@param throughputTestID throughputTestId path parameter. Throughput test ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-throughput-test
*/
func (s *DevicesService) GetDeviceLiveToolsThroughputTest(serial string, throughputTestID string) (*ResponseDevicesGetDeviceLiveToolsThroughputTest, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/throughputTest/{throughputTestId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{throughputTestId}", fmt.Sprintf("%v", throughputTestID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsThroughputTest{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsThroughputTest")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsThroughputTest)
	return result, response, err

}

//GetDeviceLiveToolsWakeOnLan Return a Wake-on-LAN job
/* Return a Wake-on-LAN job

@param serial serial path parameter.
@param wakeOnLanID wakeOnLanId path parameter. Wake on lan ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-live-tools-wake-on-lan
*/
func (s *DevicesService) GetDeviceLiveToolsWakeOnLan(serial string, wakeOnLanID string) (*ResponseDevicesGetDeviceLiveToolsWakeOnLan, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/wakeOnLan/{wakeOnLanId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{wakeOnLanId}", fmt.Sprintf("%v", wakeOnLanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsWakeOnLan{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsWakeOnLan")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsWakeOnLan)
	return result, response, err

}

//GetDeviceLldpCdp List LLDP and CDP information for a device
/* List LLDP and CDP information for a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-lldp-cdp
*/
func (s *DevicesService) GetDeviceLldpCdp(serial string) (*ResponseDevicesGetDeviceLldpCdp, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/lldpCdp"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLldpCdp{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLldpCdp")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLldpCdp)
	return result, response, err

}

//GetDeviceLossAndLatencyHistory Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for MX, MG and Z devices.
/* Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for MX, MG and Z devices.

@param serial serial path parameter.
@param getDeviceLossAndLatencyHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-loss-and-latency-history
*/
func (s *DevicesService) GetDeviceLossAndLatencyHistory(serial string, getDeviceLossAndLatencyHistoryQueryParams *GetDeviceLossAndLatencyHistoryQueryParams) (*ResponseDevicesGetDeviceLossAndLatencyHistory, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/lossAndLatencyHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceLossAndLatencyHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceLossAndLatencyHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLossAndLatencyHistory")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLossAndLatencyHistory)
	return result, response, err

}

//GetDeviceManagementInterface Return the management interface settings for a device
/* Return the management interface settings for a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-management-interface
*/
func (s *DevicesService) GetDeviceManagementInterface(serial string) (*ResponseDevicesGetDeviceManagementInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/managementInterface"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceManagementInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceManagementInterface")
	}

	result := response.Result().(*ResponseDevicesGetDeviceManagementInterface)
	return result, response, err

}

//GetNetworkDevices List the devices in a network
/* List the devices in a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-devices
*/
func (s *DevicesService) GetNetworkDevices(networkID string) (*ResponseDevicesGetNetworkDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDevices)
	return result, response, err

}

//GetNetworkSmDeviceCellularUsageHistory Return the client's daily cellular data usage history
/* Return the client's daily cellular data usage history. Usage data is in kilobytes.

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-cellular-usage-history
*/
func (s *DevicesService) GetNetworkSmDeviceCellularUsageHistory(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceCellularUsageHistory, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceCellularUsageHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceCellularUsageHistory")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceCellularUsageHistory)
	return result, response, err

}

//GetNetworkSmDeviceCerts List the certs on a device
/* List the certs on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-certs
*/
func (s *DevicesService) GetNetworkSmDeviceCerts(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceCerts, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/certs"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceCerts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceCerts")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceCerts)
	return result, response, err

}

//GetNetworkSmDeviceDeviceProfiles Get the installed profiles associated with a device
/* Get the installed profiles associated with a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-device-profiles
*/
func (s *DevicesService) GetNetworkSmDeviceDeviceProfiles(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceDeviceProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/deviceProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceDeviceProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceDeviceProfiles")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceDeviceProfiles)
	return result, response, err

}

//GetNetworkSmDeviceNetworkAdapters List the network adapters of a device
/* List the network adapters of a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-network-adapters
*/
func (s *DevicesService) GetNetworkSmDeviceNetworkAdapters(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceNetworkAdapters, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/networkAdapters"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceNetworkAdapters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceNetworkAdapters")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceNetworkAdapters)
	return result, response, err

}

//GetNetworkSmDeviceRestrictions List the restrictions on a device
/* List the restrictions on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-restrictions
*/
func (s *DevicesService) GetNetworkSmDeviceRestrictions(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceRestrictions, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/restrictions"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceRestrictions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceRestrictions")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceRestrictions)
	return result, response, err

}

//GetNetworkSmDeviceSecurityCenters List the security centers on a device
/* List the security centers on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-security-centers
*/
func (s *DevicesService) GetNetworkSmDeviceSecurityCenters(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceSecurityCenters, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/securityCenters"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceSecurityCenters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceSecurityCenters")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceSecurityCenters)
	return result, response, err

}

//GetNetworkSmDeviceSoftwares Get a list of softwares associated with a device
/* Get a list of softwares associated with a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-softwares
*/
func (s *DevicesService) GetNetworkSmDeviceSoftwares(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceSoftwares, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/softwares"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceSoftwares{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceSoftwares")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceSoftwares)
	return result, response, err

}

//GetNetworkSmDeviceWLANLists List the saved SSID names on a device
/* List the saved SSID names on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sm-device-wlan-lists
*/
func (s *DevicesService) GetNetworkSmDeviceWLANLists(networkID string, deviceID string) (*ResponseDevicesGetNetworkSmDeviceWLANLists, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/wlanLists"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkSmDeviceWLANLists{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSmDeviceWlanLists")
	}

	result := response.Result().(*ResponseDevicesGetNetworkSmDeviceWLANLists)
	return result, response, err

}

//GetOrganizationDevicesAvailabilitiesChangeHistory List the availability history information for devices in an organization.
/* List the availability history information for devices in an organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-availabilities-change-history
*/
func (s *DevicesService) GetOrganizationDevicesAvailabilitiesChangeHistory(organizationID string, getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams *GetOrganizationDevicesAvailabilitiesChangeHistoryQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/availabilities/changeHistory"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationDevicesAvailabilitiesChangeHistory")
	}

	return response, err

}

//GetOrganizationDevicesBootsHistory Returns the history of device boots in reverse chronological order (most recent first)
/* Returns the history of device boots in reverse chronological order (most recent first). Currently supported for MS devices only.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesBootsHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-devices-boots-history
*/
func (s *DevicesService) GetOrganizationDevicesBootsHistory(organizationID string, getOrganizationDevicesBootsHistoryQueryParams *GetOrganizationDevicesBootsHistoryQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/boots/history"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesBootsHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationDevicesBootsHistory")
	}

	return response, err

}

//GetOrganizationInventoryDevicesSwapsBulk List of device swaps for a given request ID ({id}).
/* List of device swaps for a given request ID ({id}).

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-inventory-devices-swaps-bulk
*/
func (s *DevicesService) GetOrganizationInventoryDevicesSwapsBulk(organizationID string, id string) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices/swaps/bulk/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationInventoryDevicesSwapsBulk")
	}

	return response, err

}

//GetOrganizationInventoryDevice Return a single device from the inventory of an organization
/* Return a single device from the inventory of an organization

@param organizationID organizationId path parameter. Organization ID
@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-inventory-device
*/
func (s *DevicesService) GetOrganizationInventoryDevice(organizationID string, serial string) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices/{serial}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationInventoryDevice")
	}

	return response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationByDevice Get average channel utilization for all bands in a network, split by AP
/* Get average channel utilization for all bands in a network, split by AP

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-by-device
*/
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationByDevice(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams *GetOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationByDevice")
	}

	return response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationByNetwork Get average channel utilization across all bands for all networks in the organization
/* Get average channel utilization across all bands for all networks in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-by-network
*/
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationByNetwork(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams *GetOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationByNetwork")
	}

	return response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval Get a time-series of average channel utilization for all bands, segmented by device.
/* Get a time-series of average channel utilization for all bands, segmented by device.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-history-by-device-by-interval
*/
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval")
	}

	return response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval Get a time-series of average channel utilization for all bands
/* Get a time-series of average channel utilization for all bands

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-channel-utilization-history-by-network-by-interval
*/
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval")
	}

	return response, err

}

//GetOrganizationWirelessDevicesPacketLossByClient Get average packet loss for the given timespan for all clients in the organization.
/* Get average packet loss for the given timespan for all clients in the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByClientQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-packet-loss-by-client
*/
func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByClient(organizationID string, getOrganizationWirelessDevicesPacketLossByClientQueryParams *GetOrganizationWirelessDevicesPacketLossByClientQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byClient"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByClientQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByClient")
	}

	return response, err

}

//GetOrganizationWirelessDevicesPacketLossByDevice Get average packet loss for the given timespan for all devices in the organization
/* Get average packet loss for the given timespan for all devices in the organization. Does not include device's own traffic.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-packet-loss-by-device
*/
func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByDevice(organizationID string, getOrganizationWirelessDevicesPacketLossByDeviceQueryParams *GetOrganizationWirelessDevicesPacketLossByDeviceQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByDevice")
	}

	return response, err

}

//GetOrganizationWirelessDevicesPacketLossByNetwork Get average packet loss for the given timespan for all networks in the organization.
/* Get average packet loss for the given timespan for all networks in the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-wireless-devices-packet-loss-by-network
*/
func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByNetwork(organizationID string, getOrganizationWirelessDevicesPacketLossByNetworkQueryParams *GetOrganizationWirelessDevicesPacketLossByNetworkQueryParams) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byNetwork"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByNetwork")
	}

	return response, err

}

//BlinkDeviceLeds Blink the LEDs on a device
/* Blink the LEDs on a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!blink-device-leds
*/

func (s *DevicesService) BlinkDeviceLeds(serial string, requestDevicesBlinkDeviceLeds *RequestDevicesBlinkDeviceLeds) (*ResponseDevicesBlinkDeviceLeds, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/blinkLeds"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesBlinkDeviceLeds).
		SetResult(&ResponseDevicesBlinkDeviceLeds{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BlinkDeviceLeds")
	}

	result := response.Result().(*ResponseDevicesBlinkDeviceLeds)
	return result, response, err

}

//CreateDeviceLiveToolsArpTable Enqueue a job to perform a ARP table request for the device
/* Enqueue a job to perform a ARP table request for the device. This endpoint currently supports switches.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-arp-table
*/

func (s *DevicesService) CreateDeviceLiveToolsArpTable(serial string, requestDevicesCreateDeviceLiveToolsArpTable *RequestDevicesCreateDeviceLiveToolsArpTable) (*ResponseDevicesCreateDeviceLiveToolsArpTable, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/arpTable"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsArpTable).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsArpTable{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsArpTable")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsArpTable)
	return result, response, err

}

//CreateDeviceLiveToolsCableTest Enqueue a job to perform a cable test for the device on the specified ports.
/* Enqueue a job to perform a cable test for the device on the specified ports.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-cable-test
*/

func (s *DevicesService) CreateDeviceLiveToolsCableTest(serial string, requestDevicesCreateDeviceLiveToolsCableTest *RequestDevicesCreateDeviceLiveToolsCableTest) (*ResponseDevicesCreateDeviceLiveToolsCableTest, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/cableTest"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsCableTest).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsCableTest{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsCableTest")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsCableTest)
	return result, response, err

}

//CreateDeviceLiveToolsPing Enqueue a job to ping a target host from the device
/* Enqueue a job to ping a target host from the device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-ping
*/

func (s *DevicesService) CreateDeviceLiveToolsPing(serial string, requestDevicesCreateDeviceLiveToolsPing *RequestDevicesCreateDeviceLiveToolsPing) (*ResponseDevicesCreateDeviceLiveToolsPing, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/ping"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsPing).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsPing{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsPing")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsPing)
	return result, response, err

}

//CreateDeviceLiveToolsPingDevice Enqueue a job to check connectivity status to the device
/* Enqueue a job to check connectivity status to the device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-ping-device
*/

func (s *DevicesService) CreateDeviceLiveToolsPingDevice(serial string, requestDevicesCreateDeviceLiveToolsPingDevice *RequestDevicesCreateDeviceLiveToolsPingDevice) (*ResponseDevicesCreateDeviceLiveToolsPingDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/pingDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsPingDevice).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsPingDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsPingDevice")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsPingDevice)
	return result, response, err

}

//CreateDeviceLiveToolsThroughputTest Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput
/* Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-throughput-test
*/

func (s *DevicesService) CreateDeviceLiveToolsThroughputTest(serial string, requestDevicesCreateDeviceLiveToolsThroughputTest *RequestDevicesCreateDeviceLiveToolsThroughputTest) (*ResponseDevicesCreateDeviceLiveToolsThroughputTest, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/throughputTest"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsThroughputTest).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsThroughputTest{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsThroughputTest")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsThroughputTest)
	return result, response, err

}

//CreateDeviceLiveToolsWakeOnLan Enqueue a job to send a Wake-on-LAN packet from the device
/* Enqueue a job to send a Wake-on-LAN packet from the device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-live-tools-wake-on-lan
*/

func (s *DevicesService) CreateDeviceLiveToolsWakeOnLan(serial string, requestDevicesCreateDeviceLiveToolsWakeOnLan *RequestDevicesCreateDeviceLiveToolsWakeOnLan) (*ResponseDevicesCreateDeviceLiveToolsWakeOnLan, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/wakeOnLan"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsWakeOnLan).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsWakeOnLan{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsWakeOnLan")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsWakeOnLan)
	return result, response, err

}

//RebootDevice Reboot a device
/* Reboot a device

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!reboot-device
*/

func (s *DevicesService) RebootDevice(serial string) (*ResponseDevicesRebootDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/reboot"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesRebootDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RebootDevice")
	}

	result := response.Result().(*ResponseDevicesRebootDevice)
	return result, response, err

}

//ClaimNetworkDevices Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requests against that device to succeed)
/* Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requests against that device to succeed). This operation can be used up to ten times within a single five minute window.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!claim-network-devices
*/

func (s *DevicesService) ClaimNetworkDevices(networkID string, requestDevicesClaimNetworkDevices *RequestDevicesClaimNetworkDevices) (*ResponseDevicesClaimNetworkDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/claim"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesClaimNetworkDevices).
		SetResult(&ResponseDevicesClaimNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ClaimNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesClaimNetworkDevices)
	return result, response, err

}

//VmxNetworkDevicesClaim Claim a vMX into a network
/* Claim a vMX into a network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!vmx-network-devices-claim
*/

func (s *DevicesService) VmxNetworkDevicesClaim(networkID string, requestDevicesVmxNetworkDevicesClaim *RequestDevicesVmxNetworkDevicesClaim) (*ResponseDevicesVmxNetworkDevicesClaim, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/claim/vmx"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesVmxNetworkDevicesClaim).
		SetResult(&ResponseDevicesVmxNetworkDevicesClaim{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation VmxNetworkDevicesClaim")
	}

	result := response.Result().(*ResponseDevicesVmxNetworkDevicesClaim)
	return result, response, err

}

//RemoveNetworkDevices Remove a single device
/* Remove a single device

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-network-devices
*/

func (s *DevicesService) RemoveNetworkDevices(networkID string, requestDevicesRemoveNetworkDevices *RequestDevicesRemoveNetworkDevices) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/remove"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesRemoveNetworkDevices).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RemoveNetworkDevices")
	}

	return response, err

}

//CheckinNetworkSmDevices Force check-in a set of devices
/* Force check-in a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!checkin-network-sm-devices
*/

func (s *DevicesService) CheckinNetworkSmDevices(networkID string, requestDevicesCheckinNetworkSmDevices *RequestDevicesCheckinNetworkSmDevices) (*ResponseDevicesCheckinNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/checkin"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCheckinNetworkSmDevices).
		SetResult(&ResponseDevicesCheckinNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CheckinNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesCheckinNetworkSmDevices)
	return result, response, err

}

//LockNetworkSmDevices Lock a set of devices
/* Lock a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lock-network-sm-devices
*/

func (s *DevicesService) LockNetworkSmDevices(networkID string, requestDevicesLockNetworkSmDevices *RequestDevicesLockNetworkSmDevices) (*ResponseDevicesLockNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/lock"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesLockNetworkSmDevices).
		SetResult(&ResponseDevicesLockNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LockNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesLockNetworkSmDevices)
	return result, response, err

}

//ModifyNetworkSmDevicesTags Add, delete, or update the tags of a set of devices
/* Add, delete, or update the tags of a set of devices

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!modify-network-sm-devices-tags
*/

func (s *DevicesService) ModifyNetworkSmDevicesTags(networkID string, requestDevicesModifyNetworkSmDevicesTags *RequestDevicesModifyNetworkSmDevicesTags) (*ResponseDevicesModifyNetworkSmDevicesTags, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/modifyTags"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesModifyNetworkSmDevicesTags).
		SetResult(&ResponseDevicesModifyNetworkSmDevicesTags{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ModifyNetworkSmDevicesTags")
	}

	result := response.Result().(*ResponseDevicesModifyNetworkSmDevicesTags)
	return result, response, err

}

//MoveNetworkSmDevices Move a set of devices to a new network
/* Move a set of devices to a new network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!move-network-sm-devices
*/

func (s *DevicesService) MoveNetworkSmDevices(networkID string, requestDevicesMoveNetworkSmDevices *RequestDevicesMoveNetworkSmDevices) (*ResponseDevicesMoveNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/move"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesMoveNetworkSmDevices).
		SetResult(&ResponseDevicesMoveNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation MoveNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesMoveNetworkSmDevices)
	return result, response, err

}

//RebootNetworkSmDevices Reboot a set of endpoints
/* Reboot a set of endpoints

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!reboot-network-sm-devices
*/

func (s *DevicesService) RebootNetworkSmDevices(networkID string, requestDevicesRebootNetworkSmDevices *RequestDevicesRebootNetworkSmDevices) (*ResponseDevicesRebootNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/reboot"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesRebootNetworkSmDevices).
		SetResult(&ResponseDevicesRebootNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RebootNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesRebootNetworkSmDevices)
	return result, response, err

}

//ShutdownNetworkSmDevices Shutdown a set of endpoints
/* Shutdown a set of endpoints

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!shutdown-network-sm-devices
*/

func (s *DevicesService) ShutdownNetworkSmDevices(networkID string, requestDevicesShutdownNetworkSmDevices *RequestDevicesShutdownNetworkSmDevices) (*ResponseDevicesShutdownNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/shutdown"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesShutdownNetworkSmDevices).
		SetResult(&ResponseDevicesShutdownNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ShutdownNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesShutdownNetworkSmDevices)
	return result, response, err

}

//WipeNetworkSmDevices Wipe a device
/* Wipe a device

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!wipe-network-sm-devices
*/

func (s *DevicesService) WipeNetworkSmDevices(networkID string, requestDevicesWipeNetworkSmDevices *RequestDevicesWipeNetworkSmDevices) (*ResponseDevicesWipeNetworkSmDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/wipe"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesWipeNetworkSmDevices).
		SetResult(&ResponseDevicesWipeNetworkSmDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation WipeNetworkSmDevices")
	}

	result := response.Result().(*ResponseDevicesWipeNetworkSmDevices)
	return result, response, err

}

//InstallNetworkSmDeviceApps Install applications on a device
/* Install applications on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!install-network-sm-device-apps
*/

func (s *DevicesService) InstallNetworkSmDeviceApps(networkID string, deviceID string, requestDevicesInstallNetworkSmDeviceApps *RequestDevicesInstallNetworkSmDeviceApps) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/installApps"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesInstallNetworkSmDeviceApps).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation InstallNetworkSmDeviceApps")
	}

	return response, err

}

//RefreshNetworkSmDeviceDetails Refresh the details of a device
/* Refresh the details of a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!refresh-network-sm-device-details
*/

func (s *DevicesService) RefreshNetworkSmDeviceDetails(networkID string, deviceID string) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/refreshDetails"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation RefreshNetworkSmDeviceDetails")
	}

	return response, err

}

//UnenrollNetworkSmDevice Unenroll a device
/* Unenroll a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!unenroll-network-sm-device
*/

func (s *DevicesService) UnenrollNetworkSmDevice(networkID string, deviceID string) (*ResponseDevicesUnenrollNetworkSmDevice, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/unenroll"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesUnenrollNetworkSmDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UnenrollNetworkSmDevice")
	}

	result := response.Result().(*ResponseDevicesUnenrollNetworkSmDevice)
	return result, response, err

}

//UninstallNetworkSmDeviceApps Uninstall applications on a device
/* Uninstall applications on a device

@param networkID networkId path parameter. Network ID
@param deviceID deviceId path parameter. Device ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!uninstall-network-sm-device-apps
*/

func (s *DevicesService) UninstallNetworkSmDeviceApps(networkID string, deviceID string, requestDevicesUninstallNetworkSmDeviceApps *RequestDevicesUninstallNetworkSmDeviceApps) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/uninstallApps"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUninstallNetworkSmDeviceApps).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UninstallNetworkSmDeviceApps")
	}

	return response, err

}

//CreateOrganizationInventoryDevicesSwapsBulk Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.
/* Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-inventory-devices-swaps-bulk
*/

func (s *DevicesService) CreateOrganizationInventoryDevicesSwapsBulk(organizationID string) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices/swaps/bulk"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationInventoryDevicesSwapsBulk")
	}

	return response, err

}

//CloneOrganizationSwitchDevices Clone port-level and some switch-level configuration settings from a source switch to one or more target switches
/* Clone port-level and some switch-level configuration settings from a source switch to one or more target switches. Cloned settings include: Aggregation Groups, Power Settings, Multicast Settings, MTU Configuration, STP Bridge priority, Port Mirroring

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!clone-organization-switch-devices
*/

func (s *DevicesService) CloneOrganizationSwitchDevices(organizationID string) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/switch/devices/clone"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CloneOrganizationSwitchDevices")
	}

	return response, err

}

//UpdateDevice Update the attributes of a device
/* Update the attributes of a device

@param serial serial path parameter.
*/
func (s *DevicesService) UpdateDevice(serial string, requestDevicesUpdateDevice *RequestDevicesUpdateDevice) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDevice).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDevice")
	}

	return response, err

}

//UpdateDeviceCellularSims Updates the SIM and APN configurations for a cellular device.
/* Updates the SIM and APN configurations for a cellular device.

@param serial serial path parameter.
*/
func (s *DevicesService) UpdateDeviceCellularSims(serial string, requestDevicesUpdateDeviceCellularSims *RequestDevicesUpdateDeviceCellularSims) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellular/sims"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceCellularSims).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCellularSims")
	}

	return response, err

}

//UpdateDeviceManagementInterface Update the management interface settings for a device
/* Update the management interface settings for a device

@param serial serial path parameter.
*/
func (s *DevicesService) UpdateDeviceManagementInterface(serial string, requestDevicesUpdateDeviceManagementInterface *RequestDevicesUpdateDeviceManagementInterface) (*ResponseDevicesUpdateDeviceManagementInterface, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/managementInterface"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceManagementInterface).
		SetResult(&ResponseDevicesUpdateDeviceManagementInterface{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceManagementInterface")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceManagementInterface)
	return result, response, err

}

//UpdateNetworkSmDevicesFields Modify the fields of a device
/* Modify the fields of a device

@param networkID networkId path parameter. Network ID
*/
func (s *DevicesService) UpdateNetworkSmDevicesFields(networkID string, requestDevicesUpdateNetworkSmDevicesFields *RequestDevicesUpdateNetworkSmDevicesFields) (*ResponseDevicesUpdateNetworkSmDevicesFields, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/fields"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateNetworkSmDevicesFields).
		SetResult(&ResponseDevicesUpdateNetworkSmDevicesFields{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSmDevicesFields")
	}

	result := response.Result().(*ResponseDevicesUpdateNetworkSmDevicesFields)
	return result, response, err

}
