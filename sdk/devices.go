package meraki

import (
	"encoding/json"
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

type ResponseDevicesGetDevice struct {
	Address        string                                  `json:"address,omitempty"`        // Physical address of the device
	BeaconIDParams *ResponseDevicesGetDeviceBeaconIDParams `json:"beaconIdParams,omitempty"` // Beacon Id parameters with an identifier and major and minor versions
	Details        *[]ResponseDevicesGetDeviceDetails      `json:"details,omitempty"`        // Additional device information
	Firmware       string                                  `json:"firmware,omitempty"`       // Firmware version of the device
	FloorPlanID    string                                  `json:"floorPlanId,omitempty"`    // The floor plan to associate to this device. null disassociates the device from the floorplan.
	LanIP          string                                  `json:"lanIp,omitempty"`          // LAN IP address of the device
	Lat            *float64                                `json:"lat,omitempty"`            // Latitude of the device
	Lng            *float64                                `json:"lng,omitempty"`            // Longitude of the device
	Mac            string                                  `json:"mac,omitempty"`            // MAC address of the device
	Model          string                                  `json:"model,omitempty"`          // Model of the device
	Name           string                                  `json:"name,omitempty"`           // Name of the device
	NetworkID      string                                  `json:"networkId,omitempty"`      // ID of the network the device belongs to
	Notes          string                                  `json:"notes,omitempty"`          // Notes for the device, limited to 255 characters
	Serial         string                                  `json:"serial,omitempty"`         // Serial number of the device
	Tags           []string                                `json:"tags,omitempty"`           // List of tags assigned to the device
}
type ResponseDevicesGetDeviceBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` // The major number to be used in the beacon identifier
	Minor *int   `json:"minor,omitempty"` // The minor number to be used in the beacon identifier
	UUID  string `json:"uuid,omitempty"`  // The UUID to be used in the beacon identifier
}
type ResponseDevicesGetDeviceDetails struct {
	Name  string `json:"name,omitempty"`  // Additional property name
	Value string `json:"value,omitempty"` // Additional property value
}
type ResponseDevicesUpdateDevice struct {
	Address        string                                     `json:"address,omitempty"`        // Physical address of the device
	BeaconIDParams *ResponseDevicesUpdateDeviceBeaconIDParams `json:"beaconIdParams,omitempty"` // Beacon Id parameters with an identifier and major and minor versions
	Details        *[]ResponseDevicesUpdateDeviceDetails      `json:"details,omitempty"`        // Additional device information
	Firmware       string                                     `json:"firmware,omitempty"`       // Firmware version of the device
	FloorPlanID    string                                     `json:"floorPlanId,omitempty"`    // The floor plan to associate to this device. null disassociates the device from the floorplan.
	LanIP          string                                     `json:"lanIp,omitempty"`          // LAN IP address of the device
	Lat            *float64                                   `json:"lat,omitempty"`            // Latitude of the device
	Lng            *float64                                   `json:"lng,omitempty"`            // Longitude of the device
	Mac            string                                     `json:"mac,omitempty"`            // MAC address of the device
	Model          string                                     `json:"model,omitempty"`          // Model of the device
	Name           string                                     `json:"name,omitempty"`           // Name of the device
	NetworkID      string                                     `json:"networkId,omitempty"`      // ID of the network the device belongs to
	Notes          string                                     `json:"notes,omitempty"`          // Notes for the device, limited to 255 characters
	Serial         string                                     `json:"serial,omitempty"`         // Serial number of the device
	Tags           []string                                   `json:"tags,omitempty"`           // List of tags assigned to the device
}
type ResponseDevicesUpdateDeviceBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` // The major number to be used in the beacon identifier
	Minor *int   `json:"minor,omitempty"` // The minor number to be used in the beacon identifier
	UUID  string `json:"uuid,omitempty"`  // The UUID to be used in the beacon identifier
}
type ResponseDevicesUpdateDeviceDetails struct {
	Name  string `json:"name,omitempty"`  // Additional property name
	Value string `json:"value,omitempty"` // Additional property value
}
type ResponseDevicesBlinkDeviceLeds struct {
	Duration *int `json:"duration,omitempty"` // The duration in seconds. Will be between 5 and 120. Default is 20 seconds
	Duty     *int `json:"duty,omitempty"`     // The duty cycle as the percent active. Will be between 10 and 90. Default is 50
	Period   *int `json:"period,omitempty"`   // The period in milliseconds. Will be between 100 and 1000. Default is 160 milliseconds
}
type ResponseDevicesGetDeviceCellularSims struct {
	SimFailover *ResponseDevicesGetDeviceCellularSimsSimFailover `json:"simFailover,omitempty"` // SIM Failover settings.
	SimOrdering []string                                         `json:"simOrdering,omitempty"` // Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It's required for devices with 3 or more SIMs and can be used in place of 'isPrimary' for dual-SIM devices. To indicate eSIM, use 'sim3'. Sim failover will occur only between primary and secondary sim slots.
	Sims        *[]ResponseDevicesGetDeviceCellularSimsSims      `json:"sims,omitempty"`        // List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.
}
type ResponseDevicesGetDeviceCellularSimsSimFailover struct {
	Enabled *bool `json:"enabled,omitempty"` // Failover to secondary SIM
	Timeout *int  `json:"timeout,omitempty"` // Failover timeout in seconds
}
type ResponseDevicesGetDeviceCellularSimsSims struct {
	Apns      *[]ResponseDevicesGetDeviceCellularSimsSimsApns `json:"apns,omitempty"`      // APN configurations. If empty, the default APN will be used.
	IsPrimary *bool                                           `json:"isPrimary,omitempty"` // If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using 'simOrdering'.
	Slot      string                                          `json:"slot,omitempty"`      // SIM slot being configured. Must be 'sim1' on single-sim devices. Use 'sim3' for eSIM.
}
type ResponseDevicesGetDeviceCellularSimsSimsApns struct {
	AllowedIPTypes []string                                                    `json:"allowedIpTypes,omitempty"` // IP versions to support (permitted values include 'ipv4', 'ipv6').
	Authentication *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication `json:"authentication,omitempty"` // APN authentication configurations.
	Name           string                                                      `json:"name,omitempty"`           // APN name.
}
type ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication struct {
	Password string `json:"password,omitempty"` // APN password, if type is set (if APN password is not supplied, the password is left unchanged).
	Type     string `json:"type,omitempty"`     // APN auth type.
	Username string `json:"username,omitempty"` // APN username, if type is set.
}
type ResponseDevicesUpdateDeviceCellularSims struct {
	SimFailover *ResponseDevicesUpdateDeviceCellularSimsSimFailover `json:"simFailover,omitempty"` // SIM Failover settings.
	SimOrdering []string                                            `json:"simOrdering,omitempty"` // Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It's required for devices with 3 or more SIMs and can be used in place of 'isPrimary' for dual-SIM devices. To indicate eSIM, use 'sim3'. Sim failover will occur only between primary and secondary sim slots.
	Sims        *[]ResponseDevicesUpdateDeviceCellularSimsSims      `json:"sims,omitempty"`        // List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.
}
type ResponseDevicesUpdateDeviceCellularSimsSimFailover struct {
	Enabled *bool `json:"enabled,omitempty"` // Failover to secondary SIM
	Timeout *int  `json:"timeout,omitempty"` // Failover timeout in seconds
}
type ResponseDevicesUpdateDeviceCellularSimsSims struct {
	Apns      *[]ResponseDevicesUpdateDeviceCellularSimsSimsApns `json:"apns,omitempty"`      // APN configurations. If empty, the default APN will be used.
	IsPrimary *bool                                              `json:"isPrimary,omitempty"` // If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using 'simOrdering'.
	Slot      string                                             `json:"slot,omitempty"`      // SIM slot being configured. Must be 'sim1' on single-sim devices. Use 'sim3' for eSIM.
}
type ResponseDevicesUpdateDeviceCellularSimsSimsApns struct {
	AllowedIPTypes []string                                                       `json:"allowedIpTypes,omitempty"` // IP versions to support (permitted values include 'ipv4', 'ipv6').
	Authentication *ResponseDevicesUpdateDeviceCellularSimsSimsApnsAuthentication `json:"authentication,omitempty"` // APN authentication configurations.
	Name           string                                                         `json:"name,omitempty"`           // APN name.
}
type ResponseDevicesUpdateDeviceCellularSimsSimsApnsAuthentication struct {
	Password string `json:"password,omitempty"` // APN password, if type is set (if APN password is not supplied, the password is left unchanged).
	Type     string `json:"type,omitempty"`     // APN auth type.
	Username string `json:"username,omitempty"` // APN username, if type is set.
}
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
	VLAN                *int                                      `json:"vlan,omitempty"`                // The client-assigned name of the VLAN the client is connected to
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
type ResponseDevicesCreateDeviceLiveToolsLedsBlink struct {
	Callback    *ResponseDevicesCreateDeviceLiveToolsLedsBlinkCallback `json:"callback,omitempty"`    // Information for callback used to send back results
	Error       string                                                 `json:"error,omitempty"`       // An error message for a failed Blink LEDs execution, if present
	LedsBlinkID string                                                 `json:"ledsBlinkId,omitempty"` // ID of led blink job
	Request     *ResponseDevicesCreateDeviceLiveToolsLedsBlinkRequest  `json:"request,omitempty"`     // The parameters of the leds blink request
	Status      string                                                 `json:"status,omitempty"`      // Status of the leds blink request
	URL         string                                                 `json:"url,omitempty"`         // GET this url to check the status of your leds blink request
}
type ResponseDevicesCreateDeviceLiveToolsLedsBlinkCallback struct {
	ID     string `json:"id,omitempty"`     // The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Status string `json:"status,omitempty"` // The status of the callback
	URL    string `json:"url,omitempty"`    // The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
}
type ResponseDevicesCreateDeviceLiveToolsLedsBlinkRequest struct {
	Duration *int   `json:"duration,omitempty"` // The duration to blink leds in seconds
	Serial   string `json:"serial,omitempty"`   // Device serial number
}
type ResponseDevicesGetDeviceLiveToolsLedsBlink struct {
	Error       string                                             `json:"error,omitempty"`       // An error message for a failed Blink LEDs execution, if present
	LedsBlinkID string                                             `json:"ledsBlinkId,omitempty"` // ID of led blink job
	Request     *ResponseDevicesGetDeviceLiveToolsLedsBlinkRequest `json:"request,omitempty"`     // The parameters of the leds blink request
	Status      string                                             `json:"status,omitempty"`      // Status of the leds blink request
	URL         string                                             `json:"url,omitempty"`         // GET this url to check the status of your leds blink request
}
type ResponseDevicesGetDeviceLiveToolsLedsBlinkRequest struct {
	Duration *int   `json:"duration,omitempty"` // The duration to blink leds in seconds
	Serial   string `json:"serial,omitempty"`   // Device serial number
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
	Count  *int   `json:"count,omitempty"`  // Number of pings to send. [1..5], default 5
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
	Count  *int   `json:"count,omitempty"`  // Number of pings to send. [1..5], default 5
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
	Count  *int   `json:"count,omitempty"`  // Number of pings to send. [1..5], default 5
	Serial string `json:"serial,omitempty"` // Device serial number
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
	Count  *int   `json:"count,omitempty"`  // Number of pings to send. [1..5], default 5
	Serial string `json:"serial,omitempty"` // Device serial number
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
	Downstream *float64 `json:"downstream,omitempty"` // Shows the download speed from shard (Mbps)
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
	Downstream *float64 `json:"downstream,omitempty"` // Shows the download speed from shard (Mbps)
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
	Address        string                                              `json:"address,omitempty"`        // Physical address of the device
	BeaconIDParams *ResponseItemDevicesGetNetworkDevicesBeaconIDParams `json:"beaconIdParams,omitempty"` // Beacon Id parameters with an identifier and major and minor versions
	Details        *[]ResponseItemDevicesGetNetworkDevicesDetails      `json:"details,omitempty"`        // Additional device information
	Firmware       string                                              `json:"firmware,omitempty"`       // Firmware version of the device
	FloorPlanID    string                                              `json:"floorPlanId,omitempty"`    // The floor plan to associate to this device. null disassociates the device from the floorplan.
	LanIP          string                                              `json:"lanIp,omitempty"`          // LAN IP address of the device
	Lat            *float64                                            `json:"lat,omitempty"`            // Latitude of the device
	Lng            *float64                                            `json:"lng,omitempty"`            // Longitude of the device
	Mac            string                                              `json:"mac,omitempty"`            // MAC address of the device
	Model          string                                              `json:"model,omitempty"`          // Model of the device
	Name           string                                              `json:"name,omitempty"`           // Name of the device
	NetworkID      string                                              `json:"networkId,omitempty"`      // ID of the network the device belongs to
	Notes          string                                              `json:"notes,omitempty"`          // Notes for the device, limited to 255 characters
	Serial         string                                              `json:"serial,omitempty"`         // Serial number of the device
	Tags           []string                                            `json:"tags,omitempty"`           // List of tags assigned to the device
}
type ResponseItemDevicesGetNetworkDevicesBeaconIDParams struct {
	Major *int   `json:"major,omitempty"` // The major number to be used in the beacon identifier
	Minor *int   `json:"minor,omitempty"` // The minor number to be used in the beacon identifier
	UUID  string `json:"uuid,omitempty"`  // The UUID to be used in the beacon identifier
}
type ResponseItemDevicesGetNetworkDevicesDetails struct {
	Name  string `json:"name,omitempty"`  // Additional property name
	Value string `json:"value,omitempty"` // Additional property value
}
type ResponseDevicesClaimNetworkDevices struct {
	Errors  *[]ResponseDevicesClaimNetworkDevicesErrors `json:"errors,omitempty"`  // Errors for devices that were not added
	Serials []string                                    `json:"serials,omitempty"` // The serials of the devices
}
type ResponseDevicesClaimNetworkDevicesErrors struct {
	Errors []string `json:"errors,omitempty"` // The errors for the device
	Serial string   `json:"serial,omitempty"` // The serial of the device
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
type ResponseDevicesBatchNetworkFloorPlansDevicesUpdate struct {
	Success *bool `json:"success,omitempty"` // Status of attempt to update device floorplan assignments
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
type ResponseDevicesInstallNetworkSmDeviceApps interface{}
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
type ResponseDevicesRefreshNetworkSmDeviceDetails interface{}
type ResponseDevicesGetNetworkSmDeviceRestrictions struct {
	Restrictions *[]ResponseDevicesGetNetworkSmDeviceRestrictionsRestrictions `json:"restrictions,omitempty"` // The list of restrictions for the device
}
type ResponseDevicesGetNetworkSmDeviceRestrictionsRestrictions struct {
	Profile      string                                                                 `json:"profile,omitempty"`      // The profile identifier.
	Restrictions *ResponseDevicesGetNetworkSmDeviceRestrictionsRestrictionsRestrictions `json:"restrictions,omitempty"` // The restrictions for the profile.
}
type ResponseDevicesGetNetworkSmDeviceRestrictionsRestrictionsRestrictions interface{}
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
type ResponseDevicesUninstallNetworkSmDeviceApps interface{}
type ResponseDevicesGetNetworkSmDeviceWLANLists []ResponseItemDevicesGetNetworkSmDeviceWLANLists // Array of ResponseDevicesGetNetworkSmDeviceWlanLists
type ResponseItemDevicesGetNetworkSmDeviceWLANLists struct {
	CreatedAt string `json:"createdAt,omitempty"` // When the Meraki record for the wlanList was created.
	ID        string `json:"id,omitempty"`        // The Meraki managed Id of the wlanList record.
	Xml       string `json:"xml,omitempty"`       // An XML string containing the WLAN List for the device.
}
type ResponseDevicesGetOrganizationDevicesAvailabilitiesChangeHistory []ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistory // Array of ResponseDevicesGetOrganizationDevicesAvailabilitiesChangeHistory
type ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistory struct {
	Details *ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDetails `json:"details,omitempty"` // Details about the status changes
	Device  *ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDevice  `json:"device,omitempty"`  // Device information
	Network *ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryNetwork `json:"network,omitempty"` // Network information
	Ts      string                                                                       `json:"ts,omitempty"`      // Timestamp, in iso8601 format, at which the event happened
}
type ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDetails struct {
	New *[]ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew `json:"new,omitempty"` // Details about the new status
	Old *[]ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld `json:"old,omitempty"` // Details about the old status
}
type ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew struct {
	Name  string `json:"name,omitempty"`  // Name of the detail
	Value string `json:"value,omitempty"` // Value of the detail
}
type ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld struct {
	Name  string `json:"name,omitempty"`  // Name of the detail
	Value string `json:"value,omitempty"` // Value of the detail
}
type ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryDevice struct {
	Model       string `json:"model,omitempty"`       // Device model
	Name        string `json:"name,omitempty"`        // Device name
	ProductType string `json:"productType,omitempty"` // Device product type.
	Serial      string `json:"serial,omitempty"`      // Device serial number
}
type ResponseItemDevicesGetOrganizationDevicesAvailabilitiesChangeHistoryNetwork struct {
	ID   string   `json:"id,omitempty"`   // Network id
	Name string   `json:"name,omitempty"` // Network name
	Tags []string `json:"tags,omitempty"` // Network tags
	URL  string   `json:"url,omitempty"`  // Network dashboard url
}
type ResponseDevicesBulkUpdateOrganizationDevicesDetails struct {
	Serials []string `json:"serials,omitempty"` // A list of serials of devices updated
}
type ResponseDevicesGetOrganizationDevicesOverviewByModel struct {
	Counts *[]ResponseDevicesGetOrganizationDevicesOverviewByModelCounts `json:"counts,omitempty"` // Counts of devices per model
}
type ResponseDevicesGetOrganizationDevicesOverviewByModelCounts struct {
	Model string `json:"model,omitempty"` // Device model
	Total *int   `json:"total,omitempty"` // Total number of devices for the model
}
type ResponseDevicesGetOrganizationFloorPlansAutoLocateDevices []ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevices // Array of ResponseDevicesGetOrganizationFloorPlansAutoLocateDevices
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevices struct {
	Items *[]ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItems `json:"items,omitempty"` // Items in the paginated dataset
	Meta  *ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItems struct {
	AutoLocate *ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItemsAutoLocate `json:"autoLocate,omitempty"` // The auto locate position for this device
	FloorPlan  *ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItemsFloorPlan  `json:"floorPlan,omitempty"`  // The assigned floor plan for this device
	IsAnchor   *bool                                                                         `json:"isAnchor,omitempty"`   // Whether or not this auto locate position is an anchor
	Lat        *float64                                                                      `json:"lat,omitempty"`        // Latitude
	Lng        *float64                                                                      `json:"lng,omitempty"`        // Longitude
	Mac        string                                                                        `json:"mac,omitempty"`        // MAC Address
	Model      string                                                                        `json:"model,omitempty"`      // Model
	Name       string                                                                        `json:"name,omitempty"`       // Device Name
	Network    *ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItemsNetwork    `json:"network,omitempty"`    // Network info
	Serial     string                                                                        `json:"serial,omitempty"`     // Device Serial Number
	Status     string                                                                        `json:"status,omitempty"`     // Device Status
	Tags       []string                                                                      `json:"tags,omitempty"`       // Tags
	Type       string                                                                        `json:"type,omitempty"`       // The type of auto locate position. Possible values: 'user', 'gnss', and 'calculated'
}
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItemsAutoLocate struct {
	Lat *float64 `json:"lat,omitempty"` // Latitude
	Lng *float64 `json:"lng,omitempty"` // Longitude
}
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItemsFloorPlan struct {
	ID     string `json:"id,omitempty"`     // Floor plan ID
	Status string `json:"status,omitempty"` // Floor plan name
}
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesItemsNetwork struct {
	ID string `json:"id,omitempty"` // ID for the network containing this device
}
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesMeta struct {
	Counts *ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesMetaCounts struct {
	Items *ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseItemDevicesGetOrganizationFloorPlansAutoLocateDevicesMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulk struct {
	JobID string                                                             `json:"jobId,omitempty"` // The ID of the job that was used to create all of the device swaps.
	Swaps *[]ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwaps `json:"swaps,omitempty"` // An array of recent swap requests and their statuses.
}
type ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwaps struct {
	AfterAction string                                                                  `json:"afterAction,omitempty"` // An action to perform on the devices.old object after swap is complete.
	CompletedAt string                                                                  `json:"completedAt,omitempty"` // An iso8601 timestamp for when the swap completed or failed.
	CreatedAt   string                                                                  `json:"createdAt,omitempty"`   // An iso8601 timestamp for the creation of the swap request.
	Devices     *ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevices `json:"devices,omitempty"`     // The devices involved in the swap
	Errors      []string                                                                `json:"errors,omitempty"`      // A list of error messages for why a swap failed.
	ID          string                                                                  `json:"id,omitempty"`          // Swap Request ID
	Status      string                                                                  `json:"status,omitempty"`      // The current status of the swap job.
}
type ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevices struct {
	New *ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevicesNew `json:"new,omitempty"` // The device that will have settings cloned to
	Old *ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevicesOld `json:"old,omitempty"` // The device that will be cloned
}
type ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevicesNew struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model name for device
	Name   string `json:"name,omitempty"`   // Customized name for device, or MAC address
	Serial string `json:"serial,omitempty"` // Serial number of device
}
type ResponseDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevicesOld struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model name for device
	Name   string `json:"name,omitempty"`   // Customized name for device, or MAC address
	Serial string `json:"serial,omitempty"` // Serial number of device
}
type ResponseDevicesGetOrganizationInventoryDevicesSwapsBulk struct {
	JobID string                                                          `json:"jobId,omitempty"` // The ID of the job that was used to create all of the device swaps.
	Swaps *[]ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwaps `json:"swaps,omitempty"` // An array of recent swap requests and their statuses.
}
type ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwaps struct {
	AfterAction string                                                               `json:"afterAction,omitempty"` // An action to perform on the devices.old object after swap is complete.
	CompletedAt string                                                               `json:"completedAt,omitempty"` // An iso8601 timestamp for when the swap completed or failed.
	CreatedAt   string                                                               `json:"createdAt,omitempty"`   // An iso8601 timestamp for the creation of the swap request.
	Devices     *ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwapsDevices `json:"devices,omitempty"`     // The devices involved in the swap
	Errors      []string                                                             `json:"errors,omitempty"`      // A list of error messages for why a swap failed.
	ID          string                                                               `json:"id,omitempty"`          // Swap Request ID
	Status      string                                                               `json:"status,omitempty"`      // The current status of the swap job.
}
type ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwapsDevices struct {
	New *ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwapsDevicesNew `json:"new,omitempty"` // The device that will have settings cloned to
	Old *ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwapsDevicesOld `json:"old,omitempty"` // The device that will be cloned
}
type ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwapsDevicesNew struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model name for device
	Name   string `json:"name,omitempty"`   // Customized name for device, or MAC address
	Serial string `json:"serial,omitempty"` // Serial number of device
}
type ResponseDevicesGetOrganizationInventoryDevicesSwapsBulkSwapsDevicesOld struct {
	Mac    string `json:"mac,omitempty"`    // MAC address of the device
	Model  string `json:"model,omitempty"`  // Model name for device
	Name   string `json:"name,omitempty"`   // Customized name for device, or MAC address
	Serial string `json:"serial,omitempty"` // Serial number of device
}
type ResponseDevicesGetOrganizationInventoryDevice struct {
	ClaimedAt             string                                                  `json:"claimedAt,omitempty"`             // Claimed time of the device
	CountryCode           string                                                  `json:"countryCode,omitempty"`           // Country/region code from device, network, or store order
	Details               *[]ResponseDevicesGetOrganizationInventoryDeviceDetails `json:"details,omitempty"`               // Additional device information
	LicenseExpirationDate string                                                  `json:"licenseExpirationDate,omitempty"` // License expiration date of the device
	Mac                   string                                                  `json:"mac,omitempty"`                   // MAC address of the device
	Model                 string                                                  `json:"model,omitempty"`                 // Model type of the device
	Name                  string                                                  `json:"name,omitempty"`                  // Name of the device
	NetworkID             string                                                  `json:"networkId,omitempty"`             // Network Id of the device
	OrderNumber           string                                                  `json:"orderNumber,omitempty"`           // Order number of the device
	ProductType           string                                                  `json:"productType,omitempty"`           // Product type of the device
	Serial                string                                                  `json:"serial,omitempty"`                // Serial number of the device
	Tags                  []string                                                `json:"tags,omitempty"`                  // Device tags
}
type ResponseDevicesGetOrganizationInventoryDeviceDetails struct {
	Name  string `json:"name,omitempty"`  // Additional property name
	Value string `json:"value,omitempty"` // Additional property value
}
type ResponseDevicesCloneOrganizationSwitchDevices struct {
	SourceSerial  string   `json:"sourceSerial,omitempty"`  // Serial number of the source switch (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials,omitempty"` // Array of serial numbers of one or more target switches (must be on a network not bound to a template)
}
type ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice []ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice // Array of ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice struct {
	ByBand  *[]ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	Mac     string                                                                               `json:"mac,omitempty"`     // The MAC address of the device.
	Network *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	Serial  string                                                                               `json:"serial,omitempty"`  // The serial number for the device.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand struct {
	Band    string                                                                                    `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork []ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork // Array of ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork struct {
	ByBand  *[]ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	Network *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand struct {
	Band    string                                                                                     `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval []ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval // Array of ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval struct {
	ByBand  *[]ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	EndTs   string                                                                                                `json:"endTs,omitempty"`   // The end time of the channel utilization interval.
	Mac     string                                                                                                `json:"mac,omitempty"`     // The MAC address of the device.
	Network *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	Serial  string                                                                                                `json:"serial,omitempty"`  // The serial number for the device.
	StartTs string                                                                                                `json:"startTs,omitempty"` // The start time of the channel utilization interval.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBand struct {
	Band    string                                                                                                     `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval []ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval // Array of ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval struct {
	ByBand  *[]ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBand `json:"byBand,omitempty"`  // Channel utilization broken down by band.
	EndTs   string                                                                                                 `json:"endTs,omitempty"`   // The end time of the channel utilization interval.
	Network *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalNetwork  `json:"network,omitempty"` // Network for the given utilization metrics.
	StartTs string                                                                                                 `json:"startTs,omitempty"` // The start time of the channel utilization interval.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBand struct {
	Band    string                                                                                                      `json:"band,omitempty"`    // The band for the given metrics.
	NonWifi *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandNonWifi `json:"nonWifi,omitempty"` // An object containing non-wifi utilization.
	Total   *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandTotal   `json:"total,omitempty"`   // An object containing total channel utilization.
	Wifi    *ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandWifi    `json:"wifi,omitempty"`    // An object containing wifi utilization.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandNonWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of non-wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandTotal struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of total channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalByBandWifi struct {
	Percentage *float64 `json:"percentage,omitempty"` // Percentage of wifi channel utiliation for the given band.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalNetwork struct {
	ID string `json:"id,omitempty"` // Network ID of the given utilization metrics.
}
type ResponseDevicesGetOrganizationWirelessDevicesPacketLossByClient []ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClient // Array of ResponseDevicesGetOrganizationWirelessDevicesPacketLossByClient
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClient struct {
	Client     *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientClient     `json:"client,omitempty"`     // Client.
	Downstream *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientClient struct {
	ID  string `json:"id,omitempty"`  // Client ID.
	Mac string `json:"mac,omitempty"` // MAC address.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByClientUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type ResponseDevicesGetOrganizationWirelessDevicesPacketLossByDevice []ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDevice // Array of ResponseDevicesGetOrganizationWirelessDevicesPacketLossByDevice
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDevice struct {
	Device     *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceDevice     `json:"device,omitempty"`     // Device.
	Downstream *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceDevice struct {
	Mac    string `json:"mac,omitempty"`    // MAC address
	Name   string `json:"name,omitempty"`   // Name
	Serial string `json:"serial,omitempty"` // Serial Number
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByDeviceUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type ResponseDevicesGetOrganizationWirelessDevicesPacketLossByNetwork []ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetwork // Array of ResponseDevicesGetOrganizationWirelessDevicesPacketLossByNetwork
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetwork struct {
	Downstream *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetworkDownstream `json:"downstream,omitempty"` // Packets sent from an AP to a client.
	Network    *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetworkNetwork    `json:"network,omitempty"`    // Network.
	Upstream   *ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetworkUpstream   `json:"upstream,omitempty"`   // Packets sent from a client to an AP.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetworkDownstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by an AP that did not reach the client.
	Total          *int     `json:"total,omitempty"`          // Total packets received by a client.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetworkNetwork struct {
	ID   string `json:"id,omitempty"`   // Network ID.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemDevicesGetOrganizationWirelessDevicesPacketLossByNetworkUpstream struct {
	LossPercentage *float64 `json:"lossPercentage,omitempty"` // Percentage of lost packets.
	Lost           *int     `json:"lost,omitempty"`           // Total packets sent by a client and did not reach the AP.
	Total          *int     `json:"total,omitempty"`          // Total packets sent by a client to an AP.
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDevice struct {
	Items *[]ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItems `json:"items,omitempty"` // List of Catalyst access points information
	Meta  *ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItems struct {
	Controller  *ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsController `json:"controller,omitempty"`  // Associated wireless controller
	CountryCode string                                                                                   `json:"countryCode,omitempty"` // Country code (2 characters)
	Details     *[]ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsDetails  `json:"details,omitempty"`     // Catalyst access point details
	JoinedAt    string                                                                                   `json:"joinedAt,omitempty"`    // The time when AP joins wireless controller
	Mode        string                                                                                   `json:"mode,omitempty"`        // AP mode (local, flex, etc.)
	Model       string                                                                                   `json:"model,omitempty"`       // AP model
	Network     *ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsNetwork    `json:"network,omitempty"`     // Catalyst access point network
	Serial      string                                                                                   `json:"serial,omitempty"`      // AP cloud ID
	Tags        *[]ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsTags     `json:"tags,omitempty"`        // The tags of the catalyst access point
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsController struct {
	Serial string `json:"serial,omitempty"` // Associated wireless controller cloud ID
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsDetails struct {
	Name  string `json:"name,omitempty"`  // Item name
	Value string `json:"value,omitempty"` // Item value
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsNetwork struct {
	ID string `json:"id,omitempty"` // Catalyst access point network ID
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsTags struct {
	Policy string `json:"policy,omitempty"` // Policy tag
	Rf     string `json:"rf,omitempty"`     // RF tag
	Site   string `json:"site,omitempty"`   // Site tag
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems `json:"items,omitempty"` // Wireless LAN controller L2 interfaces
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems struct {
	Interfaces *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces `json:"interfaces,omitempty"` // Layer 2 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                        `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces struct {
	ChannelGroup     *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup `json:"channelGroup,omitempty"`     // The channel group of this wireless LAN controller interface
	Description      string                                                                                                  `json:"description,omitempty"`      // The description of the wireless LAN controller interface
	Enabled          *bool                                                                                                   `json:"enabled,omitempty"`          // The status of the wireless LAN controller interface
	IsRedundancyPort *bool                                                                                                   `json:"isRedundancyPort,omitempty"` // Indicate whether the interface is a redundancy port used to perform HA role negotiation
	IsUplink         *bool                                                                                                   `json:"isUplink,omitempty"`         // Indicate whether the interface is uplink
	LinkNegotiation  string                                                                                                  `json:"linkNegotiation,omitempty"`  // The interface negotiation mode
	Mac              string                                                                                                  `json:"mac,omitempty"`              // The MAC address of the wireless LAN controller interface
	Module           *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule       `json:"module,omitempty"`           // The module of this wireless LAN controller interface
	Name             string                                                                                                  `json:"name,omitempty"`             // The name of the wireless LAN controller interface
	Speed            string                                                                                                  `json:"speed,omitempty"`            // The current data transfer rate which the interface is operating at. enum = [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps]
	Status           string                                                                                                  `json:"status,omitempty"`           // The status of the wireless LAN controller interface
	VLAN             *int                                                                                                    `json:"vlan,omitempty"`             // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup struct {
	Number *int `json:"number,omitempty"` // The interface channel group number
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule struct {
	Model string `json:"model,omitempty"` // The module type of this wireless LAN controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems `json:"items,omitempty"` // Wireless LAN controller layer 2 interfaces historical status
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems struct {
	Interfaces *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces `json:"interfaces,omitempty"` // layer 2 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                             `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces struct {
	Changes *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges `json:"changes,omitempty"` // The statuses of layer 2 interfaces of the wireless LAN controller
	Mac     string                                                                                                                    `json:"mac,omitempty"`     // The MAC address of the wireless LAN controller interface
	Name    string                                                                                                                    `json:"name,omitempty"`    // The name of the wireless LAN controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges struct {
	Errors   []string `json:"errors,omitempty"`   // All errors present on the port
	Status   string   `json:"status,omitempty"`   // The status of the interface
	Ts       string   `json:"ts,omitempty"`       // The timestamp of current status of the interface
	Warnings []string `json:"warnings,omitempty"` // All warnings present on the port
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller layer 2 interfaces usage
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems struct {
	Readings *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings `json:"readings,omitempty"` // The usages of layer 2 interfaces of the wireless LAN controller. Usage is in bytes
	Serial   string                                                                                                    `json:"serial,omitempty"`   // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings struct {
	Mac  string `json:"mac,omitempty"`  // The MAC address of the wireless controller interface
	Name string `json:"name,omitempty"` // The name of the wireless LAN controller interface
	Recv *int   `json:"recv,omitempty"` // The volume of data, in bytes/sec, received by wireless controller interface
	Send *int   `json:"send,omitempty"` // The volume of data, in bytes/sec, transmitted by wireless controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItems `json:"items,omitempty"` // Wireless LAN controller L3 interfaces
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItems struct {
	Interfaces *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfaces `json:"interfaces,omitempty"` // Layer 3 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                        `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfaces struct {
	Addresses       *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesAddresses  `json:"addresses,omitempty"`       // Available addresses for the interface.
	ChannelGroup    *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesChannelGroup `json:"channelGroup,omitempty"`    // The channel group of this wireless LAN controller interface
	Description     string                                                                                                  `json:"description,omitempty"`     // The description of the wireless LAN controller interface
	IsUplink        *bool                                                                                                   `json:"isUplink,omitempty"`        // Indicate whether the interface is uplink
	LinkNegotiation string                                                                                                  `json:"linkNegotiation,omitempty"` // The interface negotiation mode
	Mac             string                                                                                                  `json:"mac,omitempty"`             // The MAC address of the wireless LAN controller interface
	Module          *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesModule       `json:"module,omitempty"`          // The module of this wireless LAN controller interface
	Name            string                                                                                                  `json:"name,omitempty"`            // The name of the wireless LAN controller interface
	Speed           string                                                                                                  `json:"speed,omitempty"`           // The current data transfer rate which the interface is operating at. enum = [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps]
	Status          string                                                                                                  `json:"status,omitempty"`          // The status of the wireless LAN controller interface
	VLAN            *int                                                                                                    `json:"vlan,omitempty"`            // The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	Vrf             *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesVrf          `json:"vrf,omitempty"`             // The virtual routing and forwarding (VRF) for the wireless LAN controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesAddresses struct {
	Address  string `json:"address,omitempty"`  // The address of the wireless LAN controller interface
	Protocol string `json:"protocol,omitempty"` // Type of address for the device uplink. Available options are: ipv4, ipv6. enum = [ipv4, ipv6]
	Subnet   string `json:"subnet,omitempty"`   // The address of the wireless LAN controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesChannelGroup struct {
	Number *int `json:"number,omitempty"` // The interface channel group number
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesModule struct {
	Model string `json:"model,omitempty"` // The module type of this wireless LAN controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceItemsInterfacesVrf struct {
	Name string `json:"name,omitempty"` // The virtual routing and forwarding (VRF) name
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems `json:"items,omitempty"` // Wireless LAN controller layer 3 interfaces historical status
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems struct {
	Interfaces *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces `json:"interfaces,omitempty"` // layer 3 interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                             `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces struct {
	Changes *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges `json:"changes,omitempty"` // The statuses of layer 3 interfaces of the wireless LAN controller
	Mac     string                                                                                                                    `json:"mac,omitempty"`     // The MAC address of the wireless LAN controller interface
	Name    string                                                                                                                    `json:"name,omitempty"`    // The name of the wireless LAN controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges struct {
	Errors   []string `json:"errors,omitempty"`   // All errors present on the port
	Status   string   `json:"status,omitempty"`   // The status of the interface
	Ts       string   `json:"ts,omitempty"`       // The timestamp of current status of the interface
	Warnings []string `json:"warnings,omitempty"` // All warnings present on the port
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller layer 3 interfaces usage
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItems struct {
	Readings *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItemsReadings `json:"readings,omitempty"` // The usages of layer 3 interfaces of the wireless LAN controller. Usage is in bytes
	Serial   string                                                                                                    `json:"serial,omitempty"`   // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalItemsReadings struct {
	Mac  string `json:"mac,omitempty"`  // The MAC address of the wireless controller interface
	Name string `json:"name,omitempty"` // The name of the wireless LAN controller interface
	Recv *int   `json:"recv,omitempty"` // The volume of data, in bytes/sec, received by wireless controller interface
	Send *int   `json:"send,omitempty"` // The volume of data, in bytes/sec, transmitted by wireless controller interface
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems `json:"items,omitempty"` // Wireless LAN controller interfaces packets statuses
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems struct {
	Interfaces *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces `json:"interfaces,omitempty"` // Interfaces belongs to the wireless LAN controller
	Serial     string                                                                                                     `json:"serial,omitempty"`     // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces struct {
	Name     string                                                                                                             `json:"name,omitempty"`     // The name of the wireless LAN controller interface
	Readings *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings `json:"readings,omitempty"` // The status of packets counter on the interfaces of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings struct {
	Name  string                                                                                                               `json:"name,omitempty"`  // The type of packets being counted
	Rate  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate `json:"rate,omitempty"`  // The interface packet rates measured in packets per second
	Recv  *int                                                                                                                 `json:"recv,omitempty"`  // The total count of packets received by the interface during the timespan
	Send  *int                                                                                                                 `json:"send,omitempty"`  // The total count of packets sent by the interface during the timespan
	Total *int                                                                                                                 `json:"total,omitempty"` // The total count of sent and received packets during the timespan
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate struct {
	Recv  *int `json:"recv,omitempty"`  // The rate of packets received during the timespan
	Send  *int `json:"send,omitempty"`  // The rate of packets sent during the timespan
	Total *int `json:"total,omitempty"` // The rate of all packets sent and received during the timespan
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller interfaces usage data
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems struct {
	Intervals *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals `json:"intervals,omitempty"` // Time interval snapshots of interfaces usage data of the wireless LAN controller
	Serial    string                                                                                                   `json:"serial,omitempty"`    // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals struct {
	ByInterface *[]ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface `json:"byInterface,omitempty"` // The usage data on the interfaces of the wireless LAN controller
	EndTs       string                                                                                                              `json:"endTs,omitempty"`       // The end time interval snapshots of the query range with iso8601 format
	Overall     *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall       `json:"overall,omitempty"`     // The overall usage of all queried interfaces of the wireless LAN controller
	StartTs     string                                                                                                              `json:"startTs,omitempty"`     // The start time interval snapshots of the query range with iso8601 format
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface struct {
	Name  string                                                                                                                 `json:"name,omitempty"`  // The name of the wireless LAN controller interface
	Usage *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage `json:"usage,omitempty"` // The usage on the interfaces of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage struct {
	Recv  *int `json:"recv,omitempty"`  // The received usage on the interface during the interval, unit is bit/sec
	Send  *int `json:"send,omitempty"`  // The sent usage on the interface during the interval, unit is bit/sec
	Total *int `json:"total,omitempty"` // The total usage on the interface during the interval, unit is bit/sec
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall struct {
	Recv  *int `json:"recv,omitempty"`  // The received usage of all queried interfaces during the interval, unit is bit/sec
	Send  *int `json:"send,omitempty"`  // The sent usage of all queried interfaces during the interval, unit is bit/sec
	Total *int `json:"total,omitempty"` // The total usage of all queried interfaces during the interval, unit is bit/sec
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory []ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory // Array of ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory struct {
	Items *[]ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems `json:"items,omitempty"` // Wireless LAN controller HA failover events
	Meta  *ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems struct {
	Active *ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive `json:"active,omitempty"` // Details about the active unit
	Failed *ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed `json:"failed,omitempty"` // Details about the failed unit
	Reason string                                                                                           `json:"reason,omitempty"` // Failover reason
	Serial string                                                                                           `json:"serial,omitempty"` // Wireless LAN controller cloud ID
	Ts     string                                                                                           `json:"ts,omitempty"`     // Failover time
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive struct {
	Chassis *ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis `json:"chassis,omitempty"` // Details about the active unit chassis
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis struct {
	Name string `json:"name,omitempty"` // The name of the active chassis unit
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed struct {
	Chassis *ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis `json:"chassis,omitempty"` // Details about the failed unit chassis
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis struct {
	Name string `json:"name,omitempty"` // The name of the failed chassis unit
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta struct {
	Counts *ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts struct {
	Items *ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseItemDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatuses struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItems `json:"items,omitempty"` // Wireless LAN controller redundancy statuses
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItems struct {
	Enabled     *bool                                                                                   `json:"enabled,omitempty"`     // Wireless LAN controller redundancy enablement
	Failover    *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover `json:"failover,omitempty"`    // Wireless LAN controller failover information
	MobilityMac string                                                                                  `json:"mobilityMac,omitempty"` // Wireless LAN controller redundancy mobility mac
	Mode        string                                                                                  `json:"mode,omitempty"`        // Wireless LAN controller redundancy SSO (stateful switchover)
	Serial      string                                                                                  `json:"serial,omitempty"`      // Wireless LAN controller cloud ID
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts `json:"counts,omitempty"` // Wireless LAN controller switchover counts
	Last   *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast   `json:"last,omitempty"`   // Wireless LAN controller last failover information
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts struct {
	Total *int `json:"total,omitempty"` // Total number of failovers
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast struct {
	Reason string `json:"reason,omitempty"` // Wireless LAN controller last redundancy switchover reason
	Ts     string `json:"ts,omitempty"`     // Wireless LAN controller last redundancy switchover time
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval struct {
	Items *[]ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems `json:"items,omitempty"` // Wireless LAN controller CPU usage data
	Meta  *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta    `json:"meta,omitempty"`  // Metadata relevant to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems struct {
	Intervals *[]ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals `json:"intervals,omitempty"` // Time interval snapshots of CPU usage data of the wireless LAN controller
	Serial    string                                                                                                     `json:"serial,omitempty"`    // The cloud ID of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals struct {
	ByCore  *[]ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore `json:"byCore,omitempty"`  // The CPU usage per core of the wireless LAN controller
	EndTs   string                                                                                                           `json:"endTs,omitempty"`   // The end time of the query range  with iso8601 format
	Overall *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall  `json:"overall,omitempty"` // The overall CPU usage of the wireless LAN controller
	StartTs string                                                                                                           `json:"startTs,omitempty"` // The start time of the query range with iso8601 format
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore struct {
	Name  string                                                                                                              `json:"name,omitempty"`  // The CPU core name
	Usage *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage `json:"usage,omitempty"` // The specific core CPU usage of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage struct {
	Average *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage `json:"average,omitempty"` // The specific core average CPU usage of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage struct {
	Percentage *float64 `json:"percentage,omitempty"` // The specific core CPU usage percentage of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall struct {
	Usage *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage `json:"usage,omitempty"` // The CPU usage of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage struct {
	Average *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage `json:"average,omitempty"` // The average CPU usage of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage struct {
	Percentage *float64 `json:"percentage,omitempty"` // The CPU usage percentage of the wireless LAN controller
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta struct {
	Counts *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts `json:"counts,omitempty"` // Counts relating to the paginated dataset
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts struct {
	Items *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems `json:"items,omitempty"` // Counts relating to the paginated items
}
type ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems struct {
	Remaining *int `json:"remaining,omitempty"` // The number of items in the dataset that are available on subsequent pages
	Total     *int `json:"total,omitempty"`     // The total number of items in the dataset
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
	SimOrdering []string                                           `json:"simOrdering,omitempty"` // Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It's required for devices with 3 or more SIMs and can be used in place of 'isPrimary' for dual-SIM devices. To indicate eSIM, use 'sim3'. Sim failover will occur only between primary and secondary sim slots.
	Sims        *[]RequestDevicesUpdateDeviceCellularSimsSims      `json:"sims,omitempty"`        // List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.
}
type RequestDevicesUpdateDeviceCellularSimsSimFailover struct {
	Enabled *bool `json:"enabled,omitempty"` // Failover to secondary SIM (optional)
	Timeout *int  `json:"timeout,omitempty"` // Failover timeout in seconds (optional)
}
type RequestDevicesUpdateDeviceCellularSimsSims struct {
	Apns      *[]RequestDevicesUpdateDeviceCellularSimsSimsApns `json:"apns,omitempty"`      // APN configurations. If empty, the default APN will be used.
	IsPrimary *bool                                             `json:"isPrimary,omitempty"` // If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using 'simOrdering'.
	SimOrder  *int                                              `json:"simOrder,omitempty"`  // Priority of SIM slot being configured. Use a value between 1 and total number of SIMs available. The value must be unique for each SIM.
	Slot      string                                            `json:"slot,omitempty"`      // SIM slot being configured. Must be 'sim1' on single-sim devices. Use 'sim3' for eSIM.
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
	Ports    []string                                              `json:"ports,omitempty"`    // A list of ports for which to perform the cable test.  For Catalyst switches, IOS interface names are also supported, such as "GigabitEthernet1/0/8", "Gi1/0/8", or even "1/0/8".
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
type RequestDevicesCreateDeviceLiveToolsLedsBlink struct {
	Callback *RequestDevicesCreateDeviceLiveToolsLedsBlinkCallback `json:"callback,omitempty"` // Details for the callback. Please include either an httpServerId OR url and sharedSecret
	Duration *int                                                  `json:"duration,omitempty"` // The duration in seconds to blink LEDs.
}
type RequestDevicesCreateDeviceLiveToolsLedsBlinkCallback struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHTTPServer      `json:"httpServer,omitempty"`      // The webhook receiver used for the callback webhook.
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplate `json:"payloadTemplate,omitempty"` // The payload template of the webhook used for the callback
	SharedSecret    string                                                               `json:"sharedSecret,omitempty"`    // A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	URL             string                                                               `json:"url,omitempty"`             // The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
}
type RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHTTPServer struct {
	ID string `json:"id,omitempty"` // The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
}
type RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplate struct {
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
type RequestDevicesBatchNetworkFloorPlansDevicesUpdate struct {
	Assignments *[]RequestDevicesBatchNetworkFloorPlansDevicesUpdateAssignments `json:"assignments,omitempty"` // List of floorplan assignments to update. Up to 100 floor plan assignments can be provided in a request.
}
type RequestDevicesBatchNetworkFloorPlansDevicesUpdateAssignments struct {
	FloorPlan *RequestDevicesBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlan `json:"floorPlan,omitempty"` // Floorplan to be assigned or unassigned
	Serial    string                                                                 `json:"serial,omitempty"`    // Serial of the device to change the floor plan assignment for
}
type RequestDevicesBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlan struct {
	ID string `json:"id,omitempty"` // The ID of the floor plan to assign the device to, or null to unassign the device from its floor plan
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
type RequestDevicesBulkUpdateOrganizationDevicesDetails struct {
	Details *[]RequestDevicesBulkUpdateOrganizationDevicesDetailsDetails `json:"details,omitempty"` // An array of details
	Serials []string                                                     `json:"serials,omitempty"` // A list of serials of devices to update
}
type RequestDevicesBulkUpdateOrganizationDevicesDetailsDetails struct {
	Name  string `json:"name,omitempty"`  // Name of device detail
	Value string `json:"value,omitempty"` // Value of device detail
}
type RequestDevicesCreateOrganizationInventoryDevicesSwapsBulk struct {
	Swaps *[]RequestDevicesCreateOrganizationInventoryDevicesSwapsBulkSwaps `json:"swaps,omitempty"` // List of replacments to perform
}
type RequestDevicesCreateOrganizationInventoryDevicesSwapsBulkSwaps struct {
	AfterAction string                                                                 `json:"afterAction,omitempty"` // What action to perform on devices.old after the device cloning is complete. 'remove from network' will return the device to inventory, while 'release from organization inventory' will free up the license attached to the device.
	Devices     *RequestDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevices `json:"devices,omitempty"`     // The devices involved in the swap.
}
type RequestDevicesCreateOrganizationInventoryDevicesSwapsBulkSwapsDevices struct {
	New string `json:"new,omitempty"` // The serial of the device that the old device's settings will be cloned to.
	Old string `json:"old,omitempty"` // The serial of the device to be cloned.
}
type RequestDevicesCloneOrganizationSwitchDevices struct {
	SourceSerial  string   `json:"sourceSerial,omitempty"`  // Serial number of the source switch (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials,omitempty"` // Array of serial numbers of one or more target switches (must be on a network not bound to a template)
}

//GetDevice Return a single device
/* Return a single device

@param serial serial path parameter.


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

//GetDeviceLiveToolsLedsBlink Return a blink LEDs job
/* Return a blink LEDs job

@param serial serial path parameter.
@param ledsBlinkID ledsBlinkId path parameter. Leds blink ID


*/

func (s *DevicesService) GetDeviceLiveToolsLedsBlink(serial string, ledsBlinkID string) (*ResponseDevicesGetDeviceLiveToolsLedsBlink, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/leds/blink/{ledsBlinkId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{ledsBlinkId}", fmt.Sprintf("%v", ledsBlinkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceLiveToolsLedsBlink{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceLiveToolsLedsBlink")
	}

	result := response.Result().(*ResponseDevicesGetDeviceLiveToolsLedsBlink)
	return result, response, err

}

//GetDeviceLiveToolsPing Return a ping job
/* Return a ping job. Latency unit in response is in milliseconds. Size is in bytes.

@param serial serial path parameter.
@param id id path parameter.


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


*/

func (s *DevicesService) GetOrganizationDevicesAvailabilitiesChangeHistory(organizationID string, getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams *GetOrganizationDevicesAvailabilitiesChangeHistoryQueryParams) (*ResponseDevicesGetOrganizationDevicesAvailabilitiesChangeHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/availabilities/changeHistory"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams != nil && getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationDevicesAvailabilitiesChangeHistory
		println("Paginate")
		getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationDevicesAvailabilitiesChangeHistoryPaginate, organizationID, "", getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationDevicesAvailabilitiesChangeHistory
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationDevicesAvailabilitiesChangeHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesAvailabilitiesChangeHistory")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationDevicesAvailabilitiesChangeHistory)
	return result, response, err

}
func (s *DevicesService) GetOrganizationDevicesAvailabilitiesChangeHistoryPaginate(organizationID string, getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams any) (any, *resty.Response, error) {
	getOrganizationDevicesAvailabilitiesChangeHistoryQueryParamsConverted := getOrganizationDevicesAvailabilitiesChangeHistoryQueryParams.(*GetOrganizationDevicesAvailabilitiesChangeHistoryQueryParams)

	return s.GetOrganizationDevicesAvailabilitiesChangeHistory(organizationID, getOrganizationDevicesAvailabilitiesChangeHistoryQueryParamsConverted)
}

//GetOrganizationDevicesOverviewByModel Lists the count for each device model
/* Lists the count for each device model

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationDevicesOverviewByModelQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationDevicesOverviewByModel(organizationID string, getOrganizationDevicesOverviewByModelQueryParams *GetOrganizationDevicesOverviewByModelQueryParams) (*ResponseDevicesGetOrganizationDevicesOverviewByModel, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/overview/byModel"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationDevicesOverviewByModelQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationDevicesOverviewByModel{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationDevicesOverviewByModel")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationDevicesOverviewByModel)
	return result, response, err

}

//GetOrganizationFloorPlansAutoLocateDevices List auto locate details for each device in your organization
/* List auto locate details for each device in your organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationFloorPlansAutoLocateDevicesQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationFloorPlansAutoLocateDevices(organizationID string, getOrganizationFloorPlansAutoLocateDevicesQueryParams *GetOrganizationFloorPlansAutoLocateDevicesQueryParams) (*ResponseDevicesGetOrganizationFloorPlansAutoLocateDevices, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/floorPlans/autoLocate/devices"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationFloorPlansAutoLocateDevicesQueryParams != nil && getOrganizationFloorPlansAutoLocateDevicesQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationFloorPlansAutoLocateDevices
		println("Paginate")
		getOrganizationFloorPlansAutoLocateDevicesQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationFloorPlansAutoLocateDevicesPaginate, organizationID, "", getOrganizationFloorPlansAutoLocateDevicesQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationFloorPlansAutoLocateDevices
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationFloorPlansAutoLocateDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationFloorPlansAutoLocateDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationFloorPlansAutoLocateDevices")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationFloorPlansAutoLocateDevices)
	return result, response, err

}
func (s *DevicesService) GetOrganizationFloorPlansAutoLocateDevicesPaginate(organizationID string, getOrganizationFloorPlansAutoLocateDevicesQueryParams any) (any, *resty.Response, error) {
	getOrganizationFloorPlansAutoLocateDevicesQueryParamsConverted := getOrganizationFloorPlansAutoLocateDevicesQueryParams.(*GetOrganizationFloorPlansAutoLocateDevicesQueryParams)

	return s.GetOrganizationFloorPlansAutoLocateDevices(organizationID, getOrganizationFloorPlansAutoLocateDevicesQueryParamsConverted)
}

//GetOrganizationInventoryDevicesSwapsBulk List of device swaps for a given request ID ({id}).
/* List of device swaps for a given request ID ({id}).

@param organizationID organizationId path parameter. Organization ID
@param id id path parameter.


*/

func (s *DevicesService) GetOrganizationInventoryDevicesSwapsBulk(organizationID string, id string) (*ResponseDevicesGetOrganizationInventoryDevicesSwapsBulk, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices/swaps/bulk/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOrganizationInventoryDevicesSwapsBulk{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationInventoryDevicesSwapsBulk")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationInventoryDevicesSwapsBulk)
	return result, response, err

}

//GetOrganizationInventoryDevice Return a single device from the inventory of an organization
/* Return a single device from the inventory of an organization

@param organizationID organizationId path parameter. Organization ID
@param serial serial path parameter.


*/

func (s *DevicesService) GetOrganizationInventoryDevice(organizationID string, serial string) (*ResponseDevicesGetOrganizationInventoryDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/inventory/devices/{serial}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOrganizationInventoryDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationInventoryDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationInventoryDevice)
	return result, response, err

}

//GetOrganizationWirelessDevicesChannelUtilizationByDevice Get average channel utilization for all bands in a network, split by AP
/* Get average channel utilization for all bands in a network, split by AP

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationByDevice(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams *GetOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams != nil && getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice
		println("Paginate")
		getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesChannelUtilizationByDevicePaginate, organizationID, "", getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationByDevicePaginate(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParamsConverted := getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams.(*GetOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams)

	return s.GetOrganizationWirelessDevicesChannelUtilizationByDevice(organizationID, getOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessDevicesChannelUtilizationByNetwork Get average channel utilization across all bands for all networks in the organization
/* Get average channel utilization across all bands for all networks in the organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationByNetwork(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams *GetOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams != nil && getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork
		println("Paginate")
		getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesChannelUtilizationByNetworkPaginate, organizationID, "", getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationByNetwork")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationByNetwork)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationByNetworkPaginate(organizationID string, getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParamsConverted := getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams.(*GetOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams)

	return s.GetOrganizationWirelessDevicesChannelUtilizationByNetwork(organizationID, getOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParamsConverted)
}

//GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval Get a time-series of average channel utilization for all bands, segmented by device.
/* Get a time-series of average channel utilization for all bands, segmented by device.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams *GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams != nil && getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval
		println("Paginate")
		getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalPaginate, organizationID, "", getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalPaginate(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParamsConverted := getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams.(*GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParams)

	return s.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(organizationID, getOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval Get a time-series of average channel utilization for all bands
/* Get a time-series of average channel utilization for all bands

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams != nil && getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval
		println("Paginate")
		getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalPaginate, organizationID, "", getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalPaginate(organizationID string, getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParamsConverted := getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams.(*GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParams)

	return s.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(organizationID, getOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessDevicesPacketLossByClient Get average packet loss for the given timespan for all clients in the organization.
/* Get average packet loss for the given timespan for all clients in the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByClientQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByClient(organizationID string, getOrganizationWirelessDevicesPacketLossByClientQueryParams *GetOrganizationWirelessDevicesPacketLossByClientQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesPacketLossByClient, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byClient"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesPacketLossByClientQueryParams != nil && getOrganizationWirelessDevicesPacketLossByClientQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesPacketLossByClient
		println("Paginate")
		getOrganizationWirelessDevicesPacketLossByClientQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesPacketLossByClientPaginate, organizationID, "", getOrganizationWirelessDevicesPacketLossByClientQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesPacketLossByClient
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByClientQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesPacketLossByClient{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByClient")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesPacketLossByClient)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByClientPaginate(organizationID string, getOrganizationWirelessDevicesPacketLossByClientQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesPacketLossByClientQueryParamsConverted := getOrganizationWirelessDevicesPacketLossByClientQueryParams.(*GetOrganizationWirelessDevicesPacketLossByClientQueryParams)

	return s.GetOrganizationWirelessDevicesPacketLossByClient(organizationID, getOrganizationWirelessDevicesPacketLossByClientQueryParamsConverted)
}

//GetOrganizationWirelessDevicesPacketLossByDevice Get average packet loss for the given timespan for all devices in the organization
/* Get average packet loss for the given timespan for all devices in the organization. Does not include device's own traffic.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByDevice(organizationID string, getOrganizationWirelessDevicesPacketLossByDeviceQueryParams *GetOrganizationWirelessDevicesPacketLossByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesPacketLossByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesPacketLossByDeviceQueryParams != nil && getOrganizationWirelessDevicesPacketLossByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesPacketLossByDevice
		println("Paginate")
		getOrganizationWirelessDevicesPacketLossByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesPacketLossByDevicePaginate, organizationID, "", getOrganizationWirelessDevicesPacketLossByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesPacketLossByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesPacketLossByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesPacketLossByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByDevicePaginate(organizationID string, getOrganizationWirelessDevicesPacketLossByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesPacketLossByDeviceQueryParamsConverted := getOrganizationWirelessDevicesPacketLossByDeviceQueryParams.(*GetOrganizationWirelessDevicesPacketLossByDeviceQueryParams)

	return s.GetOrganizationWirelessDevicesPacketLossByDevice(organizationID, getOrganizationWirelessDevicesPacketLossByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessDevicesPacketLossByNetwork Get average packet loss for the given timespan for all networks in the organization.
/* Get average packet loss for the given timespan for all networks in the organization.

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesPacketLossByNetworkQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByNetwork(organizationID string, getOrganizationWirelessDevicesPacketLossByNetworkQueryParams *GetOrganizationWirelessDevicesPacketLossByNetworkQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesPacketLossByNetwork, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/packetLoss/byNetwork"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesPacketLossByNetworkQueryParams != nil && getOrganizationWirelessDevicesPacketLossByNetworkQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesPacketLossByNetwork
		println("Paginate")
		getOrganizationWirelessDevicesPacketLossByNetworkQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesPacketLossByNetworkPaginate, organizationID, "", getOrganizationWirelessDevicesPacketLossByNetworkQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesPacketLossByNetwork
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesPacketLossByNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesPacketLossByNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesPacketLossByNetwork")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesPacketLossByNetwork)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesPacketLossByNetworkPaginate(organizationID string, getOrganizationWirelessDevicesPacketLossByNetworkQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesPacketLossByNetworkQueryParamsConverted := getOrganizationWirelessDevicesPacketLossByNetworkQueryParams.(*GetOrganizationWirelessDevicesPacketLossByNetworkQueryParams)

	return s.GetOrganizationWirelessDevicesPacketLossByNetwork(organizationID, getOrganizationWirelessDevicesPacketLossByNetworkQueryParamsConverted)
}

//GetOrganizationWirelessDevicesWirelessControllersByDevice List of Catalyst access points information
/* List of Catalyst access points information

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessDevicesWirelessControllersByDevice(organizationID string, getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams *GetOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wireless/devices/wirelessControllers/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams != nil && getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDevice
		println("Paginate")
		getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessDevicesWirelessControllersByDevicePaginate, organizationID, "", getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessDevicesWirelessControllersByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessDevicesWirelessControllersByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessDevicesWirelessControllersByDevicePaginate(organizationID string, getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParamsConverted := getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams.(*GetOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams)

	return s.GetOrganizationWirelessDevicesWirelessControllersByDevice(organizationID, getOrganizationWirelessDevicesWirelessControllersByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice List wireless LAN controller layer 2 interfaces in an organization
/* List wireless LAN controller layer 2 interfaces in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l2/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL2ByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice List wireless LAN controller layer 2 interfaces history status in an organization
/* List wireless LAN controller layer 2 interfaces history status in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l2/statuses/changeHistory/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval List wireless LAN controller layer 2 interfaces history usage in an organization
/* List wireless LAN controller layer 2 interfaces history usage in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l2/usage/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice List wireless LAN controller layer 3 interfaces in an organization
/* List wireless LAN controller layer 3 interfaces in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l3/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL3ByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3ByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL3ByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL3ByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice List wireless LAN controller layer 3 interfaces history status in an organization
/* List wireless LAN controller layer 3 interfaces history status in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l3/statuses/changeHistory/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval List wireless LAN controller layer 3 interfaces history usage in an organization
/* List wireless LAN controller layer 3 interfaces history usage in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/l3/usage/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice Retrieve the packet counters for the interfaces of a Wireless LAN controller
/* Retrieve the packet counters for the interfaces of a Wireless LAN controller

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(organizationID string, getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams *GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/packets/overview/byDevice"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevicePaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevicePaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(organizationID, getOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval Retrieve the traffic for the interfaces of a Wireless LAN controller
/* Retrieve the traffic for the interfaces of a Wireless LAN controller

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/interfaces/usage/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory List the failover events of wireless LAN controllers in an organization
/* List the failover events of wireless LAN controllers in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory(organizationID string, getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams *GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/redundancy/failover/history"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams != nil && getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory
		println("Paginate")
		getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryPaginate, organizationID, "", getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result = append(*result, *resultTmp...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryPaginate(organizationID string, getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParamsConverted := getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams.(*GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams)

	return s.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory(organizationID, getOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesRedundancyStatuses List redundancy details of wireless LAN controllers in an organization
/* List redundancy details of wireless LAN controllers in an organization. The failover count refers to the total failovers system happens from the moment of this device onboarding to Dashboard

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesRedundancyStatuses(organizationID string, getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams *GetOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/redundancy/statuses"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams != nil && getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatuses
		println("Paginate")
		getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesRedundancyStatusesPaginate, organizationID, "", getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatuses
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesRedundancyStatuses")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesRedundancyStatuses)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesRedundancyStatusesPaginate(organizationID string, getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParamsConverted := getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams.(*GetOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams)

	return s.GetOrganizationWirelessControllerDevicesRedundancyStatuses(organizationID, getOrganizationWirelessControllerDevicesRedundancyStatusesQueryParamsConverted)
}

//GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval List cpu utilization data of wireless LAN controllers in an organization
/* List cpu utilization data of wireless LAN controllers in an organization

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams Filtering parameter


*/

func (s *DevicesService) GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(organizationID string, getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams *GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams) (*ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/wirelessController/devices/system/utilization/history/byInterval"
	s.rateLimiterBucket.Wait(1)

	if getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams != nil && getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams.PerPage == -1 {
		var result *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval
		println("Paginate")
		getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams.PerPage = PAGINATION_PER_PAGE
		result2, response, err := Paginate(s.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalPaginate, organizationID, "", getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams)
		if err != nil {
			return nil, nil, err
		}
		jsonResult, err := json.Marshal(result2)
		// Verficar el error
		if err != nil {
			return nil, nil, err
		}
		var paginatedResponse []any
		err = json.Unmarshal(jsonResult, &paginatedResponse)
		// for para recorrer "paginatedResponse"
		for i := 0; i < len(paginatedResponse); i++ {
			var resultTmp *ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval
			jsonResult2, _ := json.Marshal(paginatedResponse[i])
			err = json.Unmarshal(jsonResult2, &resultTmp)
			// Verificar si result es nil, si lo es inicialiarlo
			if result == nil {
				result = resultTmp
			} else {
				*result.Items = append(*result.Items, *resultTmp.Items...)
			}
		}
		return result, response, err
	}
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval)
	return result, response, err

}
func (s *DevicesService) GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalPaginate(organizationID string, getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams any) (any, *resty.Response, error) {
	getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParamsConverted := getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams.(*GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams)

	return s.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(organizationID, getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParamsConverted)
}

//BlinkDeviceLeds Blink the LEDs on a device
/* Blink the LEDs on a device.  This endpoint is deprecrated in favor of "/devices/{serial}/liveTools/leds/blink".

@param serial serial path parameter.


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
/* Enqueue a job to perform a ARP table request for the device. This endpoint currently supports switches. This endpoint has a sustained rate limit of one request every five seconds per device, with an allowed burst of five requests.

@param serial serial path parameter.


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

//CreateDeviceLiveToolsCableTest Enqueue a job to perform a cable test for the device on the specified ports
/* Enqueue a job to perform a cable test for the device on the specified ports. This endpoint has a sustained rate limit of one request every five seconds per device, with an allowed burst of five requests.

@param serial serial path parameter.


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

//CreateDeviceLiveToolsLedsBlink Enqueue a job to blink LEDs on a device
/* Enqueue a job to blink LEDs on a device. This endpoint has a rate limit of one request every 10 seconds.

@param serial serial path parameter.


*/

func (s *DevicesService) CreateDeviceLiveToolsLedsBlink(serial string, requestDevicesCreateDeviceLiveToolsLedsBlink *RequestDevicesCreateDeviceLiveToolsLedsBlink) (*ResponseDevicesCreateDeviceLiveToolsLedsBlink, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/liveTools/leds/blink"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateDeviceLiveToolsLedsBlink).
		SetResult(&ResponseDevicesCreateDeviceLiveToolsLedsBlink{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateDeviceLiveToolsLedsBlink")
	}

	result := response.Result().(*ResponseDevicesCreateDeviceLiveToolsLedsBlink)
	return result, response, err

}

//CreateDeviceLiveToolsPing Enqueue a job to ping a target host from the device
/* Enqueue a job to ping a target host from the device. This endpoint has a sustained rate limit of one request every five seconds per device, with an allowed burst of five requests.

@param serial serial path parameter.


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
/* Enqueue a job to check connectivity status to the device. This endpoint has a sustained rate limit of one request every five seconds per device, with an allowed burst of five requests.

@param serial serial path parameter.


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
/* Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput. This endpoint has a rate limit of one request every five seconds per device.

@param serial serial path parameter.


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
/* Enqueue a job to send a Wake-on-LAN packet from the device. This endpoint has a sustained rate limit of one request every five seconds per device, with an allowed burst of five requests.

@param serial serial path parameter.


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
/* Reboot a device. This endpoint has a sustained rate limit of one request every 60 seconds.

@param serial serial path parameter.


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
@param claimNetworkDevicesQueryParams Filtering parameter


*/

func (s *DevicesService) ClaimNetworkDevices(networkID string, requestDevicesClaimNetworkDevices *RequestDevicesClaimNetworkDevices, claimNetworkDevicesQueryParams *ClaimNetworkDevicesQueryParams) (*ResponseDevicesClaimNetworkDevices, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/devices/claim"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(claimNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesClaimNetworkDevices).
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

//BatchNetworkFloorPlansDevicesUpdate Update floorplan assignments for a batch of devices
/* Update floorplan assignments for a batch of devices

@param networkID networkId path parameter. Network ID


*/

func (s *DevicesService) BatchNetworkFloorPlansDevicesUpdate(networkID string, requestDevicesBatchNetworkFloorPlansDevicesUpdate *RequestDevicesBatchNetworkFloorPlansDevicesUpdate) (*ResponseDevicesBatchNetworkFloorPlansDevicesUpdate, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/floorPlans/devices/batchUpdate"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesBatchNetworkFloorPlansDevicesUpdate).
		SetResult(&ResponseDevicesBatchNetworkFloorPlansDevicesUpdate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation BatchNetworkFloorPlansDevicesUpdate")
	}

	result := response.Result().(*ResponseDevicesBatchNetworkFloorPlansDevicesUpdate)
	return result, response, err

}

//CheckinNetworkSmDevices Force check-in a set of devices
/* Force check-in a set of devices

@param networkID networkId path parameter. Network ID


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
		// SetResult(&ResponseDevicesInstallNetworkSmDeviceApps{}).
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


*/

func (s *DevicesService) RefreshNetworkSmDeviceDetails(networkID string, deviceID string) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sm/devices/{deviceId}/refreshDetails"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseDevicesRefreshNetworkSmDeviceDetails{}).
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
		// SetResult(&ResponseDevicesUninstallNetworkSmDeviceApps{}).
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

//BulkUpdateOrganizationDevicesDetails Updating device details (currently only used for Catalyst devices)
/* Updating device details (currently only used for Catalyst devices)

@param organizationID organizationId path parameter. Organization ID


*/

func (s *DevicesService) BulkUpdateOrganizationDevicesDetails(organizationID string) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/devices/details/bulkUpdate"
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
		return response, fmt.Errorf("error with operation BulkUpdateOrganizationDevicesDetails")
	}

	return response, err

}

//CreateOrganizationInventoryDevicesSwapsBulk Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.
/* Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.

@param organizationID organizationId path parameter. Organization ID


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
func (s *DevicesService) UpdateDevice(serial string, requestDevicesUpdateDevice *RequestDevicesUpdateDevice) (*ResponseDevicesUpdateDevice, *resty.Response, error) {
	path := "/api/v1/devices/{serial}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDevice).
		SetResult(&ResponseDevicesUpdateDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDevice")
	}

	result := response.Result().(*ResponseDevicesUpdateDevice)
	return result, response, err

}

//UpdateDeviceCellularSims Updates the SIM and APN configurations for a cellular device.
/* Updates the SIM and APN configurations for a cellular device.

@param serial serial path parameter.
*/
func (s *DevicesService) UpdateDeviceCellularSims(serial string, requestDevicesUpdateDeviceCellularSims *RequestDevicesUpdateDeviceCellularSims) (*ResponseDevicesUpdateDeviceCellularSims, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/cellular/sims"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceCellularSims).
		SetResult(&ResponseDevicesUpdateDeviceCellularSims{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceCellularSims")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceCellularSims)
	return result, response, err

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
