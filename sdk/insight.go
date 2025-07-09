package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type InsightService service

type GetNetworkInsightApplicationHealthByTimeQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 7 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.
	Resolution int     `url:"resolution,omitempty"` //The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 3600, 86400. The default is 300.
}

type ResponseInsightGetNetworkInsightApplicationHealthByTime []ResponseItemInsightGetNetworkInsightApplicationHealthByTime // Array of ResponseInsightGetNetworkInsightApplicationHealthByTime
type ResponseItemInsightGetNetworkInsightApplicationHealthByTime struct {
	EndTs            string   `json:"endTs,omitempty"`            // The end time of the query range
	LanGoodput       *int     `json:"lanGoodput,omitempty"`       // LAN goodput (Number of useful information bits delivered over a LAN per unit of time)
	LanLatencyMs     *float64 `json:"lanLatencyMs,omitempty"`     // LAN latency in milliseconds
	LanLossPercent   *float64 `json:"lanLossPercent,omitempty"`   // LAN loss percentage
	NumClients       *int     `json:"numClients,omitempty"`       // Number of clients
	Recv             *int     `json:"recv,omitempty"`             // Received kilobytes-per-second
	ResponseDuration *int     `json:"responseDuration,omitempty"` // Duration of the response, in milliseconds
	Sent             *int     `json:"sent,omitempty"`             // Sent kilobytes-per-second
	StartTs          string   `json:"startTs,omitempty"`          // The start time of the query range
	WanGoodput       *int     `json:"wanGoodput,omitempty"`       // WAN goodput (Number of useful information bits delivered over a WAN per unit of time)
	WanLatencyMs     *float64 `json:"wanLatencyMs,omitempty"`     // WAN latency in milliseconds
	WanLossPercent   *float64 `json:"wanLossPercent,omitempty"`   // WAN loss percentage
}
type ResponseInsightGetOrganizationInsightApplications []ResponseItemInsightGetOrganizationInsightApplications // Array of ResponseInsightGetOrganizationInsightApplications
type ResponseItemInsightGetOrganizationInsightApplications struct {
	ApplicationID string                                                           `json:"applicationId,omitempty"` // Application identifier
	Name          string                                                           `json:"name,omitempty"`          // Application name
	Thresholds    *ResponseItemInsightGetOrganizationInsightApplicationsThresholds `json:"thresholds,omitempty"`    // Thresholds defined by a user or Meraki models for each application
}
type ResponseItemInsightGetOrganizationInsightApplicationsThresholds struct {
	ByNetwork *[]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork `json:"byNetwork,omitempty"` // Threshold for each network
	Type      string                                                                      `json:"type,omitempty"`      // Threshold type (static or smart)
}
type ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork struct {
	Goodput          *int   `json:"goodput,omitempty"`          // Number of useful information bits delivered over a network per unit of time
	NetworkID        string `json:"networkId,omitempty"`        // Network identifier
	ResponseDuration *int   `json:"responseDuration,omitempty"` // Duration of the response, in milliseconds
}
type ResponseInsightGetOrganizationInsightMonitoredMediaServers []ResponseItemInsightGetOrganizationInsightMonitoredMediaServers // Array of ResponseInsightGetOrganizationInsightMonitoredMediaServers
type ResponseItemInsightGetOrganizationInsightMonitoredMediaServers struct {
	Address                     string `json:"address,omitempty"`                     // The IP address (IPv4 only) or hostname of the media server to monitor
	BestEffortMonitoringEnabled *bool  `json:"bestEffortMonitoringEnabled,omitempty"` // Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead
	ID                          string `json:"id,omitempty"`                          // Monitored media server id
	Name                        string `json:"name,omitempty"`                        // The name of the VoIP provider
}
type ResponseInsightCreateOrganizationInsightMonitoredMediaServer struct {
	Address                     string `json:"address,omitempty"`                     // The IP address (IPv4 only) or hostname of the media server to monitor
	BestEffortMonitoringEnabled *bool  `json:"bestEffortMonitoringEnabled,omitempty"` // Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead
	ID                          string `json:"id,omitempty"`                          // Monitored media server id
	Name                        string `json:"name,omitempty"`                        // The name of the VoIP provider
}
type ResponseInsightGetOrganizationInsightMonitoredMediaServer struct {
	Address                     string `json:"address,omitempty"`                     // The IP address (IPv4 only) or hostname of the media server to monitor
	BestEffortMonitoringEnabled *bool  `json:"bestEffortMonitoringEnabled,omitempty"` // Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead
	ID                          string `json:"id,omitempty"`                          // Monitored media server id
	Name                        string `json:"name,omitempty"`                        // The name of the VoIP provider
}
type ResponseInsightUpdateOrganizationInsightMonitoredMediaServer struct {
	Address                     string `json:"address,omitempty"`                     // The IP address (IPv4 only) or hostname of the media server to monitor
	BestEffortMonitoringEnabled *bool  `json:"bestEffortMonitoringEnabled,omitempty"` // Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead
	ID                          string `json:"id,omitempty"`                          // Monitored media server id
	Name                        string `json:"name,omitempty"`                        // The name of the VoIP provider
}
type RequestInsightCreateOrganizationInsightMonitoredMediaServer struct {
	Address                     string `json:"address,omitempty"`                     // The IP address (IPv4 only) or hostname of the media server to monitor
	BestEffortMonitoringEnabled *bool  `json:"bestEffortMonitoringEnabled,omitempty"` // Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead.
	Name                        string `json:"name,omitempty"`                        // The name of the VoIP provider
}
type RequestInsightUpdateOrganizationInsightMonitoredMediaServer struct {
	Address                     string `json:"address,omitempty"`                     // The IP address (IPv4 only) or hostname of the media server to monitor
	BestEffortMonitoringEnabled *bool  `json:"bestEffortMonitoringEnabled,omitempty"` // Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead.
	Name                        string `json:"name,omitempty"`                        // The name of the VoIP provider
}

//GetNetworkInsightApplicationHealthByTime Get application health by time
/* Get application health by time

@param networkID networkId path parameter. Network ID
@param applicationID applicationId path parameter. Application ID
@param getNetworkInsightApplicationHealthByTimeQueryParams Filtering parameter


*/

func (s *InsightService) GetNetworkInsightApplicationHealthByTime(networkID string, applicationID string, getNetworkInsightApplicationHealthByTimeQueryParams *GetNetworkInsightApplicationHealthByTimeQueryParams) (*ResponseInsightGetNetworkInsightApplicationHealthByTime, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/insight/applications/{applicationId}/healthByTime"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{applicationId}", fmt.Sprintf("%v", applicationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseInsightGetNetworkInsightApplicationHealthByTime](
		func() (*resty.Response, error) {
			return GET(path, s.client, getNetworkInsightApplicationHealthByTimeQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationInsightApplications List all Insight tracked applications
/* List all Insight tracked applications

@param organizationID organizationId path parameter. Organization ID


*/

func (s *InsightService) GetOrganizationInsightApplications(organizationID string) (*ResponseInsightGetOrganizationInsightApplications, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/insight/applications"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseInsightGetOrganizationInsightApplications](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationInsightMonitoredMediaServers List the monitored media servers for this organization
/* List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.

@param organizationID organizationId path parameter. Organization ID


*/

func (s *InsightService) GetOrganizationInsightMonitoredMediaServers(organizationID string) (*ResponseInsightGetOrganizationInsightMonitoredMediaServers, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/insight/monitoredMediaServers"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseInsightGetOrganizationInsightMonitoredMediaServers](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationInsightMonitoredMediaServer Return a monitored media server for this organization
/* Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.

@param organizationID organizationId path parameter. Organization ID
@param monitoredMediaServerID monitoredMediaServerId path parameter. Monitored media server ID


*/

func (s *InsightService) GetOrganizationInsightMonitoredMediaServer(organizationID string, monitoredMediaServerID string) (*ResponseInsightGetOrganizationInsightMonitoredMediaServer, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{monitoredMediaServerId}", fmt.Sprintf("%v", monitoredMediaServerID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseInsightGetOrganizationInsightMonitoredMediaServer](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//CreateOrganizationInsightMonitoredMediaServer Add a media server to be monitored for this organization
/* Add a media server to be monitored for this organization. Only valid for organizations with Meraki Insight.

@param organizationID organizationId path parameter. Organization ID


*/

func (s *InsightService) CreateOrganizationInsightMonitoredMediaServer(organizationID string, requestInsightCreateOrganizationInsightMonitoredMediaServer *RequestInsightCreateOrganizationInsightMonitoredMediaServer) (*ResponseInsightCreateOrganizationInsightMonitoredMediaServer, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/insight/monitoredMediaServers"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseInsightCreateOrganizationInsightMonitoredMediaServer](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestInsightCreateOrganizationInsightMonitoredMediaServer, nil)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//UpdateOrganizationInsightMonitoredMediaServer Update a monitored media server for this organization
/* Update a monitored media server for this organization. Only valid for organizations with Meraki Insight.

@param organizationID organizationId path parameter. Organization ID
@param monitoredMediaServerID monitoredMediaServerId path parameter. Monitored media server ID
*/
func (s *InsightService) UpdateOrganizationInsightMonitoredMediaServer(organizationID string, monitoredMediaServerID string, requestInsightUpdateOrganizationInsightMonitoredMediaServer *RequestInsightUpdateOrganizationInsightMonitoredMediaServer) (*ResponseInsightUpdateOrganizationInsightMonitoredMediaServer, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{monitoredMediaServerId}", fmt.Sprintf("%v", monitoredMediaServerID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseInsightUpdateOrganizationInsightMonitoredMediaServer](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestInsightUpdateOrganizationInsightMonitoredMediaServer)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//DeleteOrganizationInsightMonitoredMediaServer Delete a monitored media server from this organization
/* Delete a monitored media server from this organization. Only valid for organizations with Meraki Insight.

@param organizationID organizationId path parameter. Organization ID
@param monitoredMediaServerID monitoredMediaServerId path parameter. Monitored media server ID


*/
func (s *InsightService) DeleteOrganizationInsightMonitoredMediaServer(organizationID string, monitoredMediaServerID string) (*resty.Response, error) {
	//organizationID string,monitoredMediaServerID string
	path := "/api/v1/organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{monitoredMediaServerId}", fmt.Sprintf("%v", monitoredMediaServerID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
		s.backoff,
	)
}
