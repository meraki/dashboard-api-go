package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type SensorService service

type GetDeviceSensorCommandsQueryParams struct {
	Operations    []string `url:"operations[],omitempty"`  //Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData.
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	SortOrder     string   `url:"sortOrder,omitempty"`     //Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 30 days from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 30 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days.
}
type GetNetworkSensorAlertsOverviewByMetricQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 731 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 366 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 366 days. The default is 7 days. If interval is provided, the timespan will be autocalculated.
	Interval int     `url:"interval,omitempty"` //The time interval in seconds for returned data. The valid intervals are: 900, 3600, 86400, 604800, 2592000. The default is 604800. Interval is calculated if time params are provided.
}
type GetOrganizationSensorReadingsHistoryQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	T0            string   `url:"t0,omitempty"`            //The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today.
	T1            string   `url:"t1,omitempty"`            //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan      float64  `url:"timespan,omitempty"`      //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Optional parameter to filter readings by network.
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter readings by sensor.
	Metrics       []string `url:"metrics[],omitempty"`     //Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved.
}
type GetOrganizationSensorReadingsLatestQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Optional parameter to filter readings by network.
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter readings by sensor.
	Metrics       []string `url:"metrics[],omitempty"`     //Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved.
}

type ResponseSensorGetDeviceSensorCommands []ResponseItemSensorGetDeviceSensorCommands // Array of ResponseSensorGetDeviceSensorCommands
type ResponseItemSensorGetDeviceSensorCommands struct {
	CommandID   string                                              `json:"commandId,omitempty"`   // ID to check the status of the command request
	CompletedAt string                                              `json:"completedAt,omitempty"` // Time when the command was completed
	CreatedAt   string                                              `json:"createdAt,omitempty"`   // Time when the command was triggered
	CreatedBy   *ResponseItemSensorGetDeviceSensorCommandsCreatedBy `json:"createdBy,omitempty"`   // Information about the admin who triggered the command
	Errors      []string                                            `json:"errors,omitempty"`      // Array of errors if failed
	Operation   string                                              `json:"operation,omitempty"`   // Operation run on the sensor
	Status      string                                              `json:"status,omitempty"`      // Status of the command request
}
type ResponseItemSensorGetDeviceSensorCommandsCreatedBy struct {
	AdminID string `json:"adminId,omitempty"` // ID of the admin
	Email   string `json:"email,omitempty"`   // Email of the admin
	Name    string `json:"name,omitempty"`    // Name of the admin
}
type ResponseSensorCreateDeviceSensorCommand struct {
	CommandID   string                                            `json:"commandId,omitempty"`   // ID to check the status of the command request
	CompletedAt string                                            `json:"completedAt,omitempty"` // Time when the command was completed
	CreatedAt   string                                            `json:"createdAt,omitempty"`   // Time when the command was triggered
	CreatedBy   *ResponseSensorCreateDeviceSensorCommandCreatedBy `json:"createdBy,omitempty"`   // Information about the admin who triggered the command
	Errors      []string                                          `json:"errors,omitempty"`      // Array of errors if failed
	Operation   string                                            `json:"operation,omitempty"`   // Operation run on the sensor
	Status      string                                            `json:"status,omitempty"`      // Status of the command request
}
type ResponseSensorCreateDeviceSensorCommandCreatedBy struct {
	AdminID string `json:"adminId,omitempty"` // ID of the admin
	Email   string `json:"email,omitempty"`   // Email of the admin
	Name    string `json:"name,omitempty"`    // Name of the admin
}
type ResponseSensorGetDeviceSensorCommand struct {
	CommandID   string                                         `json:"commandId,omitempty"`   // ID to check the status of the command request
	CompletedAt string                                         `json:"completedAt,omitempty"` // Time when the command was completed
	CreatedAt   string                                         `json:"createdAt,omitempty"`   // Time when the command was triggered
	CreatedBy   *ResponseSensorGetDeviceSensorCommandCreatedBy `json:"createdBy,omitempty"`   // Information about the admin who triggered the command
	Errors      []string                                       `json:"errors,omitempty"`      // Array of errors if failed
	Operation   string                                         `json:"operation,omitempty"`   // Operation run on the sensor
	Status      string                                         `json:"status,omitempty"`      // Status of the command request
}
type ResponseSensorGetDeviceSensorCommandCreatedBy struct {
	AdminID string `json:"adminId,omitempty"` // ID of the admin
	Email   string `json:"email,omitempty"`   // Email of the admin
	Name    string `json:"name,omitempty"`    // Name of the admin
}
type ResponseSensorGetDeviceSensorRelationships struct {
	Livestream *ResponseSensorGetDeviceSensorRelationshipsLivestream `json:"livestream,omitempty"` // A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
}
type ResponseSensorGetDeviceSensorRelationshipsLivestream struct {
	RelatedDevices *[]ResponseSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices `json:"relatedDevices,omitempty"` // An array of the related devices for the role
}
type ResponseSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices struct {
	ProductType string `json:"productType,omitempty"` // The product type of the related device
	Serial      string `json:"serial,omitempty"`      // The serial of the related device
}
type ResponseSensorUpdateDeviceSensorRelationships struct {
	Livestream *ResponseSensorUpdateDeviceSensorRelationshipsLivestream `json:"livestream,omitempty"` // A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
}
type ResponseSensorUpdateDeviceSensorRelationshipsLivestream struct {
	RelatedDevices *[]ResponseSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices `json:"relatedDevices,omitempty"` // An array of the related devices for the role
}
type ResponseSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices struct {
	ProductType string `json:"productType,omitempty"` // The product type of the related device
	Serial      string `json:"serial,omitempty"`      // The serial of the related device
}
type ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric struct {
	Counts           *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCounts `json:"counts,omitempty"`           // Counts of currently alerting sensors, aggregated by alerting metric
	SupportedMetrics []string                                                           `json:"supportedMetrics,omitempty"` // List of metrics that are supported for alerts, based on available sensor devices in the network
}
type ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCounts struct {
	ApparentPower    *int                                                                    `json:"apparentPower,omitempty"`    // Number of sensors that are currently alerting due to apparent power readings
	Co2              *int                                                                    `json:"co2,omitempty"`              // Number of sensors that are currently alerting due to CO2 readings
	Current          *int                                                                    `json:"current,omitempty"`          // Number of sensors that are currently alerting due to electrical current readings
	Door             *int                                                                    `json:"door,omitempty"`             // Number of sensors that are currently alerting due to an open door
	Frequency        *int                                                                    `json:"frequency,omitempty"`        // Number of sensors that are currently alerting due to frequency readings
	Humidity         *int                                                                    `json:"humidity,omitempty"`         // Number of sensors that are currently alerting due to humidity readings
	IndoorAirQuality *int                                                                    `json:"indoorAirQuality,omitempty"` // Number of sensors that are currently alerting due to indoor air quality readings
	Noise            *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise `json:"noise,omitempty"`            // Object containing the number of sensors that are currently alerting due to noise readings
	Pm25             *int                                                                    `json:"pm25,omitempty"`             // Number of sensors that are currently alerting due to PM2.5 readings
	PowerFactor      *int                                                                    `json:"powerFactor,omitempty"`      // Number of sensors that are currently alerting due to power factor readings
	RealPower        *int                                                                    `json:"realPower,omitempty"`        // Number of sensors that are currently alerting due to real power readings
	Temperature      *int                                                                    `json:"temperature,omitempty"`      // Number of sensors that are currently alerting due to temperature readings
	Tvoc             *int                                                                    `json:"tvoc,omitempty"`             // Number of sensors that are currently alerting due to TVOC readings
	UpstreamPower    *int                                                                    `json:"upstreamPower,omitempty"`    // Number of sensors that are currently alerting due to an upstream power outage
	Voltage          *int                                                                    `json:"voltage,omitempty"`          // Number of sensors that are currently alerting due to voltage readings
	Water            *int                                                                    `json:"water,omitempty"`            // Number of sensors that are currently alerting due to the presence of water
}
type ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise struct {
	Ambient *int `json:"ambient,omitempty"` // Number of sensors that are currently alerting due to ambient noise readings
}
type ResponseSensorGetNetworkSensorAlertsOverviewByMetric []ResponseItemSensorGetNetworkSensorAlertsOverviewByMetric // Array of ResponseSensorGetNetworkSensorAlertsOverviewByMetric
type ResponseItemSensorGetNetworkSensorAlertsOverviewByMetric struct {
	Counts  *ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCounts `json:"counts,omitempty"`  // Counts of sensor alerts over the timespan, by reading metric
	EndTs   string                                                          `json:"endTs,omitempty"`   // End of the timespan over which sensor alerts are counted
	StartTs string                                                          `json:"startTs,omitempty"` // Start of the timespan over which sensor alerts are counted
}
type ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCounts struct {
	ApparentPower    *int                                                                 `json:"apparentPower,omitempty"`    // Number of sensor alerts that occurred due to apparent power readings
	Co2              *int                                                                 `json:"co2,omitempty"`              // Number of sensors that are currently alerting due to CO2 readings
	Current          *int                                                                 `json:"current,omitempty"`          // Number of sensor alerts that occurred due to electrical current readings
	Door             *int                                                                 `json:"door,omitempty"`             // Number of sensor alerts that occurred due to an open door
	Frequency        *int                                                                 `json:"frequency,omitempty"`        // Number of sensor alerts that occurred due to frequency readings
	Humidity         *int                                                                 `json:"humidity,omitempty"`         // Number of sensor alerts that occurred due to humidity readings
	IndoorAirQuality *int                                                                 `json:"indoorAirQuality,omitempty"` // Number of sensor alerts that occurred due to indoor air quality readings
	Noise            *ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise `json:"noise,omitempty"`            // Object containing the number of sensor alerts that occurred due to noise readings
	Pm25             *int                                                                 `json:"pm25,omitempty"`             // Number of sensor alerts that occurred due to PM2.5 readings
	PowerFactor      *int                                                                 `json:"powerFactor,omitempty"`      // Number of sensor alerts that occurred due to power factor readings
	RealPower        *int                                                                 `json:"realPower,omitempty"`        // Number of sensor alerts that occurred due to real power readings
	Temperature      *int                                                                 `json:"temperature,omitempty"`      // Number of sensor alerts that occurred due to temperature readings
	Tvoc             *int                                                                 `json:"tvoc,omitempty"`             // Number of sensor alerts that occurred due to TVOC readings
	UpstreamPower    *int                                                                 `json:"upstreamPower,omitempty"`    // Number of sensor alerts that occurred due to upstream power outages
	Voltage          *int                                                                 `json:"voltage,omitempty"`          // Number of sensor alerts that occurred due to voltage readings
	Water            *int                                                                 `json:"water,omitempty"`            // Number of sensor alerts that occurred due to the presence of water
}
type ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise struct {
	Ambient *int `json:"ambient,omitempty"` // Number of sensor alerts that occurred due to ambient noise readings
}
type ResponseSensorGetNetworkSensorAlertsProfiles []ResponseItemSensorGetNetworkSensorAlertsProfiles // Array of ResponseSensorGetNetworkSensorAlertsProfiles
type ResponseItemSensorGetNetworkSensorAlertsProfiles struct {
	Conditions       *[]ResponseItemSensorGetNetworkSensorAlertsProfilesConditions `json:"conditions,omitempty"`       // List of conditions that will cause the profile to send an alert.
	IncludesensorURL *bool                                                         `json:"includeSensorUrl,omitempty"` // Include dashboard link to sensor in messages (default: true).
	Message          string                                                        `json:"message,omitempty"`          // A custom message that will appear in email and text message alerts.
	Name             string                                                        `json:"name,omitempty"`             // Name of the sensor alert profile.
	ProfileID        string                                                        `json:"profileId,omitempty"`        // ID of the sensor alert profile.
	Recipients       *ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients   `json:"recipients,omitempty"`       // List of recipients that will receive the alert.
	Schedule         *ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule     `json:"schedule,omitempty"`         // The sensor schedule to use with the alert profile.
	Serials          []string                                                      `json:"serials,omitempty"`          // List of device serials assigned to this sensor alert profile.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditions struct {
	Direction string                                                               `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
	Duration  *int                                                                 `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Metric    string                                                               `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes.
	Threshold *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold struct {
	ApparentPower    *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdApparentPower    `json:"apparentPower,omitempty"`    // Apparent power threshold. 'draw' must be provided.
	Co2              *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCo2              `json:"co2,omitempty"`              // CO2 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Current          *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCurrent          `json:"current,omitempty"`          // Electrical current threshold. 'level' must be provided.
	Door             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Frequency        *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdFrequency        `json:"frequency,omitempty"`        // Electrical frequency threshold. 'level' must be provided.
	Humidity         *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	PowerFactor      *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPowerFactor      `json:"powerFactor,omitempty"`      // Power factor threshold. 'percentage' must be provided.
	RealPower        *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdRealPower        `json:"realPower,omitempty"`        // Real power threshold. 'draw' must be provided.
	Temperature      *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	UpstreamPower    *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdUpstreamPower    `json:"upstreamPower,omitempty"`    // Upstream power threshold. 'outageDetected' must be provided and set to true.
	Voltage          *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdVoltage          `json:"voltage,omitempty"`          // Voltage threshold. 'level' must be provided.
	Water            *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in volt-amps. Must be between 0 and 3750.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCo2 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as CO2 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative CO2 level.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in amps. Must be between 0 and 15.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdFrequency struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in hertz. Must be between 0 and 60.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity struct {
	Quality            string `json:"quality,omitempty"`            // Alerting threshold as a qualitative humidity level.
	RelativePercentage *int   `json:"relativePercentage,omitempty"` // Alerting threshold in %RH.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality struct {
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative indoor air quality level.
	Score   *int   `json:"score,omitempty"`   // Alerting threshold as indoor air quality score.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise struct {
	Ambient *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoiseAmbient `json:"ambient,omitempty"` // Ambient noise threshold. One of 'level' or 'quality' must be provided.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoiseAmbient struct {
	Level   *int   `json:"level,omitempty"`   // Alerting threshold as adjusted decibels.
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative ambient noise level.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as PM2.5 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative PM2.5 level.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in watts. Must be between 0 and 3750.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdUpstreamPower struct {
	OutageDetected *bool `json:"outageDetected,omitempty"` // Alerting threshold for an upstream power event. Must be set to true.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdVoltage struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in volts. Must be between 0 and 250.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater struct {
	Present *bool `json:"present,omitempty"` // Alerting threshold for a water detection event. Must be set to true.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert.
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of webhook endpoint IDs that will receive information about the alert.
	SmsNumbers    []string `json:"smsNumbers,omitempty"`    // A list of SMS numbers that will receive information about the alert.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule struct {
	ID   string `json:"id,omitempty"`   // ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
	Name string `json:"name,omitempty"` // Name of the sensor schedule to use with the alert profile.
}
type ResponseSensorCreateNetworkSensorAlertsProfile struct {
	Conditions       *[]ResponseSensorCreateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"`       // List of conditions that will cause the profile to send an alert.
	IncludesensorURL *bool                                                       `json:"includeSensorUrl,omitempty"` // Include dashboard link to sensor in messages (default: true).
	Message          string                                                      `json:"message,omitempty"`          // A custom message that will appear in email and text message alerts.
	Name             string                                                      `json:"name,omitempty"`             // Name of the sensor alert profile.
	ProfileID        string                                                      `json:"profileId,omitempty"`        // ID of the sensor alert profile.
	Recipients       *ResponseSensorCreateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"`       // List of recipients that will receive the alert.
	Schedule         *ResponseSensorCreateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`         // The sensor schedule to use with the alert profile.
	Serials          []string                                                    `json:"serials,omitempty"`          // List of device serials assigned to this sensor alert profile.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                             `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
	Duration  *int                                                               `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Metric    string                                                             `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes.
	Threshold *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThreshold struct {
	ApparentPower    *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower    `json:"apparentPower,omitempty"`    // Apparent power threshold. 'draw' must be provided.
	Co2              *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdCo2              `json:"co2,omitempty"`              // CO2 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Current          *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent          `json:"current,omitempty"`          // Electrical current threshold. 'level' must be provided.
	Door             *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Frequency        *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency        `json:"frequency,omitempty"`        // Electrical frequency threshold. 'level' must be provided.
	Humidity         *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	PowerFactor      *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor      `json:"powerFactor,omitempty"`      // Power factor threshold. 'percentage' must be provided.
	RealPower        *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower        `json:"realPower,omitempty"`        // Real power threshold. 'draw' must be provided.
	Temperature      *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	UpstreamPower    *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower    `json:"upstreamPower,omitempty"`    // Upstream power threshold. 'outageDetected' must be provided and set to true.
	Voltage          *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage          `json:"voltage,omitempty"`          // Voltage threshold. 'level' must be provided.
	Water            *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in volt-amps. Must be between 0 and 3750.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdCo2 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as CO2 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative CO2 level.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in amps. Must be between 0 and 15.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in hertz. Must be between 0 and 60.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity struct {
	Quality            string `json:"quality,omitempty"`            // Alerting threshold as a qualitative humidity level.
	RelativePercentage *int   `json:"relativePercentage,omitempty"` // Alerting threshold in %RH.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality struct {
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative indoor air quality level.
	Score   *int   `json:"score,omitempty"`   // Alerting threshold as indoor air quality score.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise struct {
	Ambient *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient `json:"ambient,omitempty"` // Ambient noise threshold. One of 'level' or 'quality' must be provided.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient struct {
	Level   *int   `json:"level,omitempty"`   // Alerting threshold as adjusted decibels.
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative ambient noise level.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as PM2.5 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative PM2.5 level.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in watts. Must be between 0 and 3750.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower struct {
	OutageDetected *bool `json:"outageDetected,omitempty"` // Alerting threshold for an upstream power event. Must be set to true.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in volts. Must be between 0 and 250.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater struct {
	Present *bool `json:"present,omitempty"` // Alerting threshold for a water detection event. Must be set to true.
}
type ResponseSensorCreateNetworkSensorAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert.
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of webhook endpoint IDs that will receive information about the alert.
	SmsNumbers    []string `json:"smsNumbers,omitempty"`    // A list of SMS numbers that will receive information about the alert.
}
type ResponseSensorCreateNetworkSensorAlertsProfileSchedule struct {
	ID   string `json:"id,omitempty"`   // ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
	Name string `json:"name,omitempty"` // Name of the sensor schedule to use with the alert profile.
}
type ResponseSensorGetNetworkSensorAlertsProfile struct {
	Conditions       *[]ResponseSensorGetNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"`       // List of conditions that will cause the profile to send an alert.
	IncludesensorURL *bool                                                    `json:"includeSensorUrl,omitempty"` // Include dashboard link to sensor in messages (default: true).
	Message          string                                                   `json:"message,omitempty"`          // A custom message that will appear in email and text message alerts.
	Name             string                                                   `json:"name,omitempty"`             // Name of the sensor alert profile.
	ProfileID        string                                                   `json:"profileId,omitempty"`        // ID of the sensor alert profile.
	Recipients       *ResponseSensorGetNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"`       // List of recipients that will receive the alert.
	Schedule         *ResponseSensorGetNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`         // The sensor schedule to use with the alert profile.
	Serials          []string                                                 `json:"serials,omitempty"`          // List of device serials assigned to this sensor alert profile.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditions struct {
	Direction string                                                          `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
	Duration  *int                                                            `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Metric    string                                                          `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes.
	Threshold *ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold struct {
	ApparentPower    *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPower    `json:"apparentPower,omitempty"`    // Apparent power threshold. 'draw' must be provided.
	Co2              *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCo2              `json:"co2,omitempty"`              // CO2 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Current          *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrent          `json:"current,omitempty"`          // Electrical current threshold. 'level' must be provided.
	Door             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Frequency        *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequency        `json:"frequency,omitempty"`        // Electrical frequency threshold. 'level' must be provided.
	Humidity         *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	PowerFactor      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactor      `json:"powerFactor,omitempty"`      // Power factor threshold. 'percentage' must be provided.
	RealPower        *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPower        `json:"realPower,omitempty"`        // Real power threshold. 'draw' must be provided.
	Temperature      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	UpstreamPower    *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPower    `json:"upstreamPower,omitempty"`    // Upstream power threshold. 'outageDetected' must be provided and set to true.
	Voltage          *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltage          `json:"voltage,omitempty"`          // Voltage threshold. 'level' must be provided.
	Water            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in volt-amps. Must be between 0 and 3750.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCo2 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as CO2 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative CO2 level.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in amps. Must be between 0 and 15.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequency struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in hertz. Must be between 0 and 60.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity struct {
	Quality            string `json:"quality,omitempty"`            // Alerting threshold as a qualitative humidity level.
	RelativePercentage *int   `json:"relativePercentage,omitempty"` // Alerting threshold in %RH.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality struct {
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative indoor air quality level.
	Score   *int   `json:"score,omitempty"`   // Alerting threshold as indoor air quality score.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise struct {
	Ambient *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient `json:"ambient,omitempty"` // Ambient noise threshold. One of 'level' or 'quality' must be provided.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient struct {
	Level   *int   `json:"level,omitempty"`   // Alerting threshold as adjusted decibels.
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative ambient noise level.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as PM2.5 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative PM2.5 level.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in watts. Must be between 0 and 3750.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPower struct {
	OutageDetected *bool `json:"outageDetected,omitempty"` // Alerting threshold for an upstream power event. Must be set to true.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltage struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in volts. Must be between 0 and 250.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater struct {
	Present *bool `json:"present,omitempty"` // Alerting threshold for a water detection event. Must be set to true.
}
type ResponseSensorGetNetworkSensorAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert.
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of webhook endpoint IDs that will receive information about the alert.
	SmsNumbers    []string `json:"smsNumbers,omitempty"`    // A list of SMS numbers that will receive information about the alert.
}
type ResponseSensorGetNetworkSensorAlertsProfileSchedule struct {
	ID   string `json:"id,omitempty"`   // ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
	Name string `json:"name,omitempty"` // Name of the sensor schedule to use with the alert profile.
}
type ResponseSensorUpdateNetworkSensorAlertsProfile struct {
	Conditions       *[]ResponseSensorUpdateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"`       // List of conditions that will cause the profile to send an alert.
	IncludesensorURL *bool                                                       `json:"includeSensorUrl,omitempty"` // Include dashboard link to sensor in messages (default: true).
	Message          string                                                      `json:"message,omitempty"`          // A custom message that will appear in email and text message alerts.
	Name             string                                                      `json:"name,omitempty"`             // Name of the sensor alert profile.
	ProfileID        string                                                      `json:"profileId,omitempty"`        // ID of the sensor alert profile.
	Recipients       *ResponseSensorUpdateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"`       // List of recipients that will receive the alert.
	Schedule         *ResponseSensorUpdateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`         // The sensor schedule to use with the alert profile.
	Serials          []string                                                    `json:"serials,omitempty"`          // List of device serials assigned to this sensor alert profile.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                             `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
	Duration  *int                                                               `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Metric    string                                                             `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes.
	Threshold *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThreshold struct {
	ApparentPower    *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower    `json:"apparentPower,omitempty"`    // Apparent power threshold. 'draw' must be provided.
	Co2              *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCo2              `json:"co2,omitempty"`              // CO2 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Current          *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent          `json:"current,omitempty"`          // Electrical current threshold. 'level' must be provided.
	Door             *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Frequency        *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency        `json:"frequency,omitempty"`        // Electrical frequency threshold. 'level' must be provided.
	Humidity         *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	PowerFactor      *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor      `json:"powerFactor,omitempty"`      // Power factor threshold. 'percentage' must be provided.
	RealPower        *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower        `json:"realPower,omitempty"`        // Real power threshold. 'draw' must be provided.
	Temperature      *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	UpstreamPower    *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower    `json:"upstreamPower,omitempty"`    // Upstream power threshold. 'outageDetected' must be provided and set to true.
	Voltage          *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage          `json:"voltage,omitempty"`          // Voltage threshold. 'level' must be provided.
	Water            *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in volt-amps. Must be between 0 and 3750.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCo2 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as CO2 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative CO2 level.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in amps. Must be between 0 and 15.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in hertz. Must be between 0 and 60.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity struct {
	Quality            string `json:"quality,omitempty"`            // Alerting threshold as a qualitative humidity level.
	RelativePercentage *int   `json:"relativePercentage,omitempty"` // Alerting threshold in %RH.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality struct {
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative indoor air quality level.
	Score   *int   `json:"score,omitempty"`   // Alerting threshold as indoor air quality score.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise struct {
	Ambient *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient `json:"ambient,omitempty"` // Ambient noise threshold. One of 'level' or 'quality' must be provided.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient struct {
	Level   *int   `json:"level,omitempty"`   // Alerting threshold as adjusted decibels.
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative ambient noise level.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as PM2.5 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative PM2.5 level.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in watts. Must be between 0 and 3750.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower struct {
	OutageDetected *bool `json:"outageDetected,omitempty"` // Alerting threshold for an upstream power event. Must be set to true.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in volts. Must be between 0 and 250.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater struct {
	Present *bool `json:"present,omitempty"` // Alerting threshold for a water detection event. Must be set to true.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert.
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of webhook endpoint IDs that will receive information about the alert.
	SmsNumbers    []string `json:"smsNumbers,omitempty"`    // A list of SMS numbers that will receive information about the alert.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileSchedule struct {
	ID   string `json:"id,omitempty"`   // ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
	Name string `json:"name,omitempty"` // Name of the sensor schedule to use with the alert profile.
}
type ResponseSensorGetNetworkSensorMqttBrokers []ResponseItemSensorGetNetworkSensorMqttBrokers // Array of ResponseSensorGetNetworkSensorMqttBrokers
type ResponseItemSensorGetNetworkSensorMqttBrokers struct {
	Enabled      *bool  `json:"enabled,omitempty"`      // Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.
	MqttBrokerID string `json:"mqttBrokerId,omitempty"` // ID of the MQTT Broker.
}
type ResponseSensorGetNetworkSensorMqttBroker struct {
	Enabled      *bool  `json:"enabled,omitempty"`      // Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.
	MqttBrokerID string `json:"mqttBrokerId,omitempty"` // ID of the MQTT Broker.
}
type ResponseSensorUpdateNetworkSensorMqttBroker struct {
	Enabled      *bool  `json:"enabled,omitempty"`      // Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.
	MqttBrokerID string `json:"mqttBrokerId,omitempty"` // ID of the MQTT Broker.
}
type ResponseSensorGetNetworkSensorRelationships []ResponseItemSensorGetNetworkSensorRelationships // Array of ResponseSensorGetNetworkSensorRelationships
type ResponseItemSensorGetNetworkSensorRelationships struct {
	Device        *ResponseItemSensorGetNetworkSensorRelationshipsDevice        `json:"device,omitempty"`        // A sensor or gateway device in the network
	Relationships *ResponseItemSensorGetNetworkSensorRelationshipsRelationships `json:"relationships,omitempty"` // An object describing the relationships defined between the device and other devices
}
type ResponseItemSensorGetNetworkSensorRelationshipsDevice struct {
	Name        string `json:"name,omitempty"`        // The name of the device
	ProductType string `json:"productType,omitempty"` // The product type of the device
	Serial      string `json:"serial,omitempty"`      // The serial of the device
}
type ResponseItemSensorGetNetworkSensorRelationshipsRelationships struct {
	Livestream *ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestream `json:"livestream,omitempty"` // A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
}
type ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestream struct {
	RelatedDevices *[]ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestreamRelatedDevices `json:"relatedDevices,omitempty"` // An array of the related devices for the role
}
type ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestreamRelatedDevices struct {
	ProductType string `json:"productType,omitempty"` // The product type of the related device
	Serial      string `json:"serial,omitempty"`      // The serial of the related device
}
type ResponseSensorGetOrganizationSensorReadingsHistory []ResponseItemSensorGetOrganizationSensorReadingsHistory // Array of ResponseSensorGetOrganizationSensorReadingsHistory
type ResponseItemSensorGetOrganizationSensorReadingsHistory struct {
	ApparentPower       *ResponseItemSensorGetOrganizationSensorReadingsHistoryApparentPower       `json:"apparentPower,omitempty"`       // Reading for the 'apparentPower' metric. This will only be present if the 'metric' property equals 'apparentPower'.
	Battery             *ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery             `json:"battery,omitempty"`             // Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'.
	Button              *ResponseItemSensorGetOrganizationSensorReadingsHistoryButton              `json:"button,omitempty"`              // Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'.
	Co2                 *ResponseItemSensorGetOrganizationSensorReadingsHistoryCo2                 `json:"co2,omitempty"`                 // Reading for the 'co2' metric. This will only be present if the 'metric' property equals 'co2'.
	Current             *ResponseItemSensorGetOrganizationSensorReadingsHistoryCurrent             `json:"current,omitempty"`             // Reading for the 'current' metric. This will only be present if the 'metric' property equals 'current'.
	Door                *ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor                `json:"door,omitempty"`                // Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'.
	DownstreamPower     *ResponseItemSensorGetOrganizationSensorReadingsHistoryDownstreamPower     `json:"downstreamPower,omitempty"`     // Reading for the 'downstreamPower' metric. This will only be present if the 'metric' property equals 'downstreamPower'.
	Frequency           *ResponseItemSensorGetOrganizationSensorReadingsHistoryFrequency           `json:"frequency,omitempty"`           // Reading for the 'frequency' metric. This will only be present if the 'metric' property equals 'frequency'.
	Humidity            *ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity            `json:"humidity,omitempty"`            // Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'.
	IndoorAirQuality    *ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality    `json:"indoorAirQuality,omitempty"`    // Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'.
	Metric              string                                                                     `json:"metric,omitempty"`              // Type of sensor reading.
	Network             *ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork             `json:"network,omitempty"`             // Network to which the sensor belongs.
	Noise               *ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise               `json:"noise,omitempty"`               // Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'.
	Pm25                *ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25                `json:"pm25,omitempty"`                // Reading for the 'pm25' metric. This will only be present if the 'metric' property equals 'pm25'.
	PowerFactor         *ResponseItemSensorGetOrganizationSensorReadingsHistoryPowerFactor         `json:"powerFactor,omitempty"`         // Reading for the 'powerFactor' metric. This will only be present if the 'metric' property equals 'powerFactor'.
	RealPower           *ResponseItemSensorGetOrganizationSensorReadingsHistoryRealPower           `json:"realPower,omitempty"`           // Reading for the 'realPower' metric. This will only be present if the 'metric' property equals 'realPower'.
	RemoteLockoutSwitch *ResponseItemSensorGetOrganizationSensorReadingsHistoryRemoteLockoutSwitch `json:"remoteLockoutSwitch,omitempty"` // Reading for the 'remoteLockoutSwitch' metric. This will only be present if the 'metric' property equals 'remoteLockoutSwitch'.
	Serial              string                                                                     `json:"serial,omitempty"`              // Serial number of the sensor that took the reading.
	Temperature         *ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature         `json:"temperature,omitempty"`         // Reading for the 'temperature' metric. This will only be present if the 'metric' property equals 'temperature'.
	Ts                  string                                                                     `json:"ts,omitempty"`                  // Time at which the reading occurred, in ISO8601 format.
	Tvoc                *ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc                `json:"tvoc,omitempty"`                // Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'.
	Voltage             *ResponseItemSensorGetOrganizationSensorReadingsHistoryVoltage             `json:"voltage,omitempty"`             // Reading for the 'voltage' metric. This will only be present if the 'metric' property equals 'voltage'.
	Water               *ResponseItemSensorGetOrganizationSensorReadingsHistoryWater               `json:"water,omitempty"`               // Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Apparent power reading in volt-amperes.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery struct {
	Percentage *int `json:"percentage,omitempty"` // Remaining battery life.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryButton struct {
	PressType string `json:"pressType,omitempty"` // Type of button press that occurred.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryCo2 struct {
	Concentration *int `json:"concentration,omitempty"` // CO2 reading in parts per million.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Electrical current reading in amperes.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor struct {
	Open *bool `json:"open,omitempty"` // True if the door is open.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryDownstreamPower struct {
	Enabled *bool `json:"enabled,omitempty"` // True if power is turned on to the device that is connected downstream of the MT40 power monitor.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryFrequency struct {
	Level *float64 `json:"level,omitempty"` // Electrical current frequency reading in hertz.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity struct {
	RelativePercentage *int `json:"relativePercentage,omitempty"` // Humidity reading in %RH.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality struct {
	Score *int `json:"score,omitempty"` // Indoor air quality score between 0 and 100.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork struct {
	ID   string `json:"id,omitempty"`   // ID of the network.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise struct {
	Ambient *ResponseItemSensorGetOrganizationSensorReadingsHistoryNoiseAmbient `json:"ambient,omitempty"` // Ambient noise reading.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryNoiseAmbient struct {
	Level *int `json:"level,omitempty"` // Ambient noise reading in adjusted decibels.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25 struct {
	Concentration *int `json:"concentration,omitempty"` // PM2.5 reading in micrograms per cubic meter.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Power factor reading as a percentage.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Real power reading in watts.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryRemoteLockoutSwitch struct {
	Locked *bool `json:"locked,omitempty"` // True if power controls are disabled via the MT40's physical remote lockout switch.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Temperature reading in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Temperature reading in degrees Fahrenheit.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc struct {
	Concentration *int `json:"concentration,omitempty"` // TVOC reading in micrograms per cubic meter.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryVoltage struct {
	Level *float64 `json:"level,omitempty"` // Voltage reading in volts.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryWater struct {
	Present *bool `json:"present,omitempty"` // True if water is detected.
}
type ResponseSensorGetOrganizationSensorReadingsLatest []ResponseItemSensorGetOrganizationSensorReadingsLatest // Array of ResponseSensorGetOrganizationSensorReadingsLatest
type ResponseItemSensorGetOrganizationSensorReadingsLatest struct {
	Network  *ResponseItemSensorGetOrganizationSensorReadingsLatestNetwork    `json:"network,omitempty"`  // Network to which the sensor belongs.
	Readings *[]ResponseItemSensorGetOrganizationSensorReadingsLatestReadings `json:"readings,omitempty"` // Array of latest readings from the sensor. Each object represents a single reading for a single metric.
	Serial   string                                                           `json:"serial,omitempty"`   // Serial number of the sensor that took the readings.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestNetwork struct {
	ID   string `json:"id,omitempty"`   // ID of the network.
	Name string `json:"name,omitempty"` // Name of the network.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadings struct {
	ApparentPower       *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsApparentPower       `json:"apparentPower,omitempty"`       // Reading for the 'apparentPower' metric. This will only be present if the 'metric' property equals 'apparentPower'.
	Battery             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery             `json:"battery,omitempty"`             // Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'.
	Button              *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton              `json:"button,omitempty"`              // Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'.
	Co2                 *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCo2                 `json:"co2,omitempty"`                 // Reading for the 'co2' metric. This will only be present if the 'metric' property equals 'co2'.
	Current             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCurrent             `json:"current,omitempty"`             // Reading for the 'current' metric. This will only be present if the 'metric' property equals 'current'.
	Door                *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor                `json:"door,omitempty"`                // Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'.
	DownstreamPower     *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDownstreamPower     `json:"downstreamPower,omitempty"`     // Reading for the 'downstreamPower' metric. This will only be present if the 'metric' property equals 'downstreamPower'.
	Frequency           *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsFrequency           `json:"frequency,omitempty"`           // Reading for the 'frequency' metric. This will only be present if the 'metric' property equals 'frequency'.
	Humidity            *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity            `json:"humidity,omitempty"`            // Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'.
	IndoorAirQuality    *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality    `json:"indoorAirQuality,omitempty"`    // Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'.
	Metric              string                                                                            `json:"metric,omitempty"`              // Type of sensor reading.
	Noise               *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise               `json:"noise,omitempty"`               // Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'.
	Pm25                *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25                `json:"pm25,omitempty"`                // Reading for the 'pm25' metric. This will only be present if the 'metric' property equals 'pm25'.
	PowerFactor         *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPowerFactor         `json:"powerFactor,omitempty"`         // Reading for the 'powerFactor' metric. This will only be present if the 'metric' property equals 'powerFactor'.
	RealPower           *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRealPower           `json:"realPower,omitempty"`           // Reading for the 'realPower' metric. This will only be present if the 'metric' property equals 'realPower'.
	RemoteLockoutSwitch *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRemoteLockoutSwitch `json:"remoteLockoutSwitch,omitempty"` // Reading for the 'remoteLockoutSwitch' metric. This will only be present if the 'metric' property equals 'remoteLockoutSwitch'.
	Temperature         *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature         `json:"temperature,omitempty"`         // Reading for the 'temperature' metric. This will only be present if the 'metric' property equals 'temperature'.
	Ts                  string                                                                            `json:"ts,omitempty"`                  // Time at which the reading occurred, in ISO8601 format.
	Tvoc                *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc                `json:"tvoc,omitempty"`                // Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'.
	Voltage             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsVoltage             `json:"voltage,omitempty"`             // Reading for the 'voltage' metric. This will only be present if the 'metric' property equals 'voltage'.
	Water               *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater               `json:"water,omitempty"`               // Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Apparent power reading in volt-amperes.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery struct {
	Percentage *int `json:"percentage,omitempty"` // Remaining battery life.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton struct {
	PressType string `json:"pressType,omitempty"` // Type of button press that occurred.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCo2 struct {
	Concentration *int `json:"concentration,omitempty"` // CO2 reading in parts per million.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Electrical current reading in amperes.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor struct {
	Open *bool `json:"open,omitempty"` // True if the door is open.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDownstreamPower struct {
	Enabled *bool `json:"enabled,omitempty"` // True if power is turned on to the device that is connected downstream of the MT40 power monitor.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsFrequency struct {
	Level *float64 `json:"level,omitempty"` // Electrical current frequency reading in hertz.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity struct {
	RelativePercentage *int `json:"relativePercentage,omitempty"` // Humidity reading in %RH.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality struct {
	Score *int `json:"score,omitempty"` // Indoor air quality score between 0 and 100.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise struct {
	Ambient *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoiseAmbient `json:"ambient,omitempty"` // Ambient noise reading.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoiseAmbient struct {
	Level *int `json:"level,omitempty"` // Ambient noise reading in adjusted decibels.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25 struct {
	Concentration *int `json:"concentration,omitempty"` // PM2.5 reading in micrograms per cubic meter.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Power factor reading as a percentage.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Real power reading in watts.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRemoteLockoutSwitch struct {
	Locked *bool `json:"locked,omitempty"` // True if power controls are disabled via the MT40's physical remote lockout switch.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Temperature reading in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Temperature reading in degrees Fahrenheit.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc struct {
	Concentration *int `json:"concentration,omitempty"` // TVOC reading in micrograms per cubic meter.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsVoltage struct {
	Level *float64 `json:"level,omitempty"` // Voltage reading in volts.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater struct {
	Present *bool `json:"present,omitempty"` // True if water is detected.
}
type RequestSensorCreateDeviceSensorCommand struct {
	Operation string `json:"operation,omitempty"` // Operation to run on the sensor. 'enableDownstreamPower', 'disableDownstreamPower', and 'cycleDownstreamPower' turn power on/off to the device that is connected downstream of an MT40 power monitor. 'refreshData' causes an MT15 or MT40 device to upload its latest readings so that they are immediately available in the Dashboard API.
}
type RequestSensorUpdateDeviceSensorRelationships struct {
	Livestream *RequestSensorUpdateDeviceSensorRelationshipsLivestream `json:"livestream,omitempty"` // A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
}
type RequestSensorUpdateDeviceSensorRelationshipsLivestream struct {
	RelatedDevices *[]RequestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices `json:"relatedDevices,omitempty"` // An array of the related devices for the role
}
type RequestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices struct {
	Serial string `json:"serial,omitempty"` // The serial of the related device
}
type RequestSensorCreateNetworkSensorAlertsProfile struct {
	Conditions       *[]RequestSensorCreateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"`       // List of conditions that will cause the profile to send an alert.
	IncludesensorURL *bool                                                      `json:"includeSensorUrl,omitempty"` // Include dashboard link to sensor in messages (default: true).
	Message          string                                                     `json:"message,omitempty"`          // A custom message that will appear in email and text message alerts.
	Name             string                                                     `json:"name,omitempty"`             // Name of the sensor alert profile.
	Recipients       *RequestSensorCreateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"`       // List of recipients that will receive the alert.
	Schedule         *RequestSensorCreateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`         // The sensor schedule to use with the alert profile.
	Serials          []string                                                   `json:"serials,omitempty"`          // List of device serials assigned to this sensor alert profile.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                            `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
	Duration  *int                                                              `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Metric    string                                                            `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes.
	Threshold *RequestSensorCreateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThreshold struct {
	ApparentPower    *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower    `json:"apparentPower,omitempty"`    // Apparent power threshold. 'draw' must be provided.
	Co2              *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCo2              `json:"co2,omitempty"`              // CO2 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Current          *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent          `json:"current,omitempty"`          // Electrical current threshold. 'level' must be provided.
	Door             *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Frequency        *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency        `json:"frequency,omitempty"`        // Electrical frequency threshold. 'level' must be provided.
	Humidity         *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	PowerFactor      *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor      `json:"powerFactor,omitempty"`      // Power factor threshold. 'percentage' must be provided.
	RealPower        *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower        `json:"realPower,omitempty"`        // Real power threshold. 'draw' must be provided.
	Temperature      *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	UpstreamPower    *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower    `json:"upstreamPower,omitempty"`    // Upstream power threshold. 'outageDetected' must be provided and set to true.
	Voltage          *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage          `json:"voltage,omitempty"`          // Voltage threshold. 'level' must be provided.
	Water            *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in volt-amps. Must be between 0 and 3750.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCo2 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as CO2 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative CO2 level.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in amps. Must be between 0 and 15.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in hertz. Must be between 0 and 60.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity struct {
	Quality            string `json:"quality,omitempty"`            // Alerting threshold as a qualitative humidity level.
	RelativePercentage *int   `json:"relativePercentage,omitempty"` // Alerting threshold in %RH.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality struct {
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative indoor air quality level.
	Score   *int   `json:"score,omitempty"`   // Alerting threshold as indoor air quality score.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise struct {
	Ambient *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient `json:"ambient,omitempty"` // Ambient noise threshold. One of 'level' or 'quality' must be provided.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient struct {
	Level   *int   `json:"level,omitempty"`   // Alerting threshold as adjusted decibels.
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative ambient noise level.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as PM2.5 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative PM2.5 level.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in watts. Must be between 0 and 3750.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower struct {
	OutageDetected *bool `json:"outageDetected,omitempty"` // Alerting threshold for an upstream power event. Must be set to true.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in volts. Must be between 0 and 250.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater struct {
	Present *bool `json:"present,omitempty"` // Alerting threshold for a water detection event. Must be set to true.
}
type RequestSensorCreateNetworkSensorAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert.
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of webhook endpoint IDs that will receive information about the alert.
	SmsNumbers    []string `json:"smsNumbers,omitempty"`    // A list of SMS numbers that will receive information about the alert.
}
type RequestSensorCreateNetworkSensorAlertsProfileSchedule struct {
	ID string `json:"id,omitempty"` // ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
}
type RequestSensorUpdateNetworkSensorAlertsProfile struct {
	Conditions       *[]RequestSensorUpdateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"`       // List of conditions that will cause the profile to send an alert.
	IncludesensorURL *bool                                                      `json:"includeSensorUrl,omitempty"` // Include dashboard link to sensor in messages (default: true).
	Message          string                                                     `json:"message,omitempty"`          // A custom message that will appear in email and text message alerts.
	Name             string                                                     `json:"name,omitempty"`             // Name of the sensor alert profile.
	Recipients       *RequestSensorUpdateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"`       // List of recipients that will receive the alert.
	Schedule         *RequestSensorUpdateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`         // The sensor schedule to use with the alert profile.
	Serials          []string                                                   `json:"serials,omitempty"`          // List of device serials assigned to this sensor alert profile.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                            `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
	Duration  *int                                                              `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Metric    string                                                            `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes.
	Threshold *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold struct {
	ApparentPower    *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower    `json:"apparentPower,omitempty"`    // Apparent power threshold. 'draw' must be provided.
	Co2              *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCo2              `json:"co2,omitempty"`              // CO2 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Current          *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent          `json:"current,omitempty"`          // Electrical current threshold. 'level' must be provided.
	Door             *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Frequency        *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency        `json:"frequency,omitempty"`        // Electrical frequency threshold. 'level' must be provided.
	Humidity         *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	PowerFactor      *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor      `json:"powerFactor,omitempty"`      // Power factor threshold. 'percentage' must be provided.
	RealPower        *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower        `json:"realPower,omitempty"`        // Real power threshold. 'draw' must be provided.
	Temperature      *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	UpstreamPower    *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower    `json:"upstreamPower,omitempty"`    // Upstream power threshold. 'outageDetected' must be provided and set to true.
	Voltage          *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage          `json:"voltage,omitempty"`          // Voltage threshold. 'level' must be provided.
	Water            *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in volt-amps. Must be between 0 and 3750.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCo2 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as CO2 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative CO2 level.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in amps. Must be between 0 and 15.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in hertz. Must be between 0 and 60.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity struct {
	Quality            string `json:"quality,omitempty"`            // Alerting threshold as a qualitative humidity level.
	RelativePercentage *int   `json:"relativePercentage,omitempty"` // Alerting threshold in %RH.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality struct {
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative indoor air quality level.
	Score   *int   `json:"score,omitempty"`   // Alerting threshold as indoor air quality score.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise struct {
	Ambient *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient `json:"ambient,omitempty"` // Ambient noise threshold. One of 'level' or 'quality' must be provided.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient struct {
	Level   *int   `json:"level,omitempty"`   // Alerting threshold as adjusted decibels.
	Quality string `json:"quality,omitempty"` // Alerting threshold as a qualitative ambient noise level.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25 struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as PM2.5 parts per million.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative PM2.5 level.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor struct {
	Percentage *int `json:"percentage,omitempty"` // Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower struct {
	Draw *float64 `json:"draw,omitempty"` // Alerting threshold in watts. Must be between 0 and 3750.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower struct {
	OutageDetected *bool `json:"outageDetected,omitempty"` // Alerting threshold for an upstream power event. Must be set to true.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage struct {
	Level *float64 `json:"level,omitempty"` // Alerting threshold in volts. Must be between 0 and 250.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater struct {
	Present *bool `json:"present,omitempty"` // Alerting threshold for a water detection event. Must be set to true.
}
type RequestSensorUpdateNetworkSensorAlertsProfileRecipients struct {
	Emails        []string `json:"emails,omitempty"`        // A list of emails that will receive information about the alert.
	HTTPServerIDs []string `json:"httpServerIds,omitempty"` // A list of webhook endpoint IDs that will receive information about the alert.
	SmsNumbers    []string `json:"smsNumbers,omitempty"`    // A list of SMS numbers that will receive information about the alert.
}
type RequestSensorUpdateNetworkSensorAlertsProfileSchedule struct {
	ID string `json:"id,omitempty"` // ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
}
type RequestSensorUpdateNetworkSensorMqttBroker struct {
	Enabled *bool `json:"enabled,omitempty"` // Set to true to enable MQTT broker for sensor network
}

//GetDeviceSensorCommands Returns a historical log of all commands
/* Returns a historical log of all commands

@param serial serial path parameter.
@param getDeviceSensorCommandsQueryParams Filtering parameter


*/

func (s *SensorService) GetDeviceSensorCommands(serial string, getDeviceSensorCommandsQueryParams *GetDeviceSensorCommandsQueryParams) (*ResponseSensorGetDeviceSensorCommands, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/sensor/commands"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetDeviceSensorCommands](
		func() (*resty.Response, error) {
			return GET(path, s.client, getDeviceSensorCommandsQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		func(dst, src ResponseSensorGetDeviceSensorCommands) ResponseSensorGetDeviceSensorCommands {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getDeviceSensorCommandsQueryParams != nil {
				return getDeviceSensorCommandsQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetDeviceSensorCommand Returns information about the command's execution, including the status
/* Returns information about the command's execution, including the status

@param serial serial path parameter.
@param commandID commandId path parameter. Command ID


*/

func (s *SensorService) GetDeviceSensorCommand(serial string, commandID string) (*ResponseSensorGetDeviceSensorCommand, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/sensor/commands/{commandId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{commandId}", fmt.Sprintf("%v", commandID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetDeviceSensorCommand](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceSensorRelationships List the sensor roles for a given sensor or camera device.
/* List the sensor roles for a given sensor or camera device.

@param serial serial path parameter.


*/

func (s *SensorService) GetDeviceSensorRelationships(serial string) (*ResponseSensorGetDeviceSensorRelationships, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/sensor/relationships"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetDeviceSensorRelationships](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkSensorAlertsCurrentOverviewByMetric Return an overview of currently alerting sensors by metric
/* Return an overview of currently alerting sensors by metric

@param networkID networkId path parameter. Network ID


*/

func (s *SensorService) GetNetworkSensorAlertsCurrentOverviewByMetric(networkID string) (*ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/current/overview/byMetric"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkSensorAlertsOverviewByMetric Return an overview of alert occurrences over a timespan, by metric
/* Return an overview of alert occurrences over a timespan, by metric

@param networkID networkId path parameter. Network ID
@param getNetworkSensorAlertsOverviewByMetricQueryParams Filtering parameter


*/

func (s *SensorService) GetNetworkSensorAlertsOverviewByMetric(networkID string, getNetworkSensorAlertsOverviewByMetricQueryParams *GetNetworkSensorAlertsOverviewByMetricQueryParams) (*ResponseSensorGetNetworkSensorAlertsOverviewByMetric, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/overview/byMetric"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetNetworkSensorAlertsOverviewByMetric](
		func() (*resty.Response, error) {
			return GET(path, s.client, getNetworkSensorAlertsOverviewByMetricQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkSensorAlertsProfiles Lists all sensor alert profiles for a network.
/* Lists all sensor alert profiles for a network.

@param networkID networkId path parameter. Network ID


*/

func (s *SensorService) GetNetworkSensorAlertsProfiles(networkID string) (*ResponseSensorGetNetworkSensorAlertsProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetNetworkSensorAlertsProfiles](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkSensorAlertsProfile Show details of a sensor alert profile for a network.
/* Show details of a sensor alert profile for a network.

@param networkID networkId path parameter. Network ID
@param id id path parameter.


*/

func (s *SensorService) GetNetworkSensorAlertsProfile(networkID string, id string) (*ResponseSensorGetNetworkSensorAlertsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetNetworkSensorAlertsProfile](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkSensorMqttBrokers List the sensor settings of all MQTT brokers for this network
/* List the sensor settings of all MQTT brokers for this network. To get the brokers themselves, use /networks/{networkId}/mqttBrokers.

@param networkID networkId path parameter. Network ID


*/

func (s *SensorService) GetNetworkSensorMqttBrokers(networkID string) (*ResponseSensorGetNetworkSensorMqttBrokers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/mqttBrokers"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetNetworkSensorMqttBrokers](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkSensorMqttBroker Return the sensor settings of an MQTT broker
/* Return the sensor settings of an MQTT broker. To get the broker itself, use /networks/{networkId}/mqttBrokers/{mqttBrokerId}.

@param networkID networkId path parameter. Network ID
@param mqttBrokerID mqttBrokerId path parameter. Mqtt broker ID


*/

func (s *SensorService) GetNetworkSensorMqttBroker(networkID string, mqttBrokerID string) (*ResponseSensorGetNetworkSensorMqttBroker, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{mqttBrokerId}", fmt.Sprintf("%v", mqttBrokerID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetNetworkSensorMqttBroker](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkSensorRelationships List the sensor roles for devices in a given network
/* List the sensor roles for devices in a given network

@param networkID networkId path parameter. Network ID


*/

func (s *SensorService) GetNetworkSensorRelationships(networkID string) (*ResponseSensorGetNetworkSensorRelationships, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/relationships"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetNetworkSensorRelationships](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationSensorReadingsHistory Return all reported readings from sensors in a given timespan, sorted by timestamp
/* Return all reported readings from sensors in a given timespan, sorted by timestamp

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSensorReadingsHistoryQueryParams Filtering parameter


*/

func (s *SensorService) GetOrganizationSensorReadingsHistory(organizationID string, getOrganizationSensorReadingsHistoryQueryParams *GetOrganizationSensorReadingsHistoryQueryParams) (*ResponseSensorGetOrganizationSensorReadingsHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/sensor/readings/history"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetOrganizationSensorReadingsHistory](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationSensorReadingsHistoryQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		func(dst, src ResponseSensorGetOrganizationSensorReadingsHistory) ResponseSensorGetOrganizationSensorReadingsHistory {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getOrganizationSensorReadingsHistoryQueryParams != nil {
				return getOrganizationSensorReadingsHistoryQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//GetOrganizationSensorReadingsLatest Return the latest available reading for each metric from each sensor, sorted by sensor serial
/* Return the latest available reading for each metric from each sensor, sorted by sensor serial

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSensorReadingsLatestQueryParams Filtering parameter


*/

func (s *SensorService) GetOrganizationSensorReadingsLatest(organizationID string, getOrganizationSensorReadingsLatestQueryParams *GetOrganizationSensorReadingsLatestQueryParams) (*ResponseSensorGetOrganizationSensorReadingsLatest, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/sensor/readings/latest"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorGetOrganizationSensorReadingsLatest](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationSensorReadingsLatestQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		func(dst, src ResponseSensorGetOrganizationSensorReadingsLatest) ResponseSensorGetOrganizationSensorReadingsLatest {
			dst = append(dst, src...)
			return dst
		},
		func() bool {
			if getOrganizationSensorReadingsLatestQueryParams != nil {
				return getOrganizationSensorReadingsLatestQueryParams.PerPage == -1
			}
			return false
		}(),
	)

}

//CreateDeviceSensorCommand Sends a command to a sensor
/* Sends a command to a sensor

@param serial serial path parameter.


*/

func (s *SensorService) CreateDeviceSensorCommand(serial string, requestSensorCreateDeviceSensorCommand *RequestSensorCreateDeviceSensorCommand) (*ResponseSensorCreateDeviceSensorCommand, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/sensor/commands"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Past way

	return doWithRetriesAndResult[ResponseSensorCreateDeviceSensorCommand](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestSensorCreateDeviceSensorCommand, nil)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//CreateNetworkSensorAlertsProfile Creates a sensor alert profile for a network.
/* Creates a sensor alert profile for a network.

@param networkID networkId path parameter. Network ID


*/

func (s *SensorService) CreateNetworkSensorAlertsProfile(networkID string, requestSensorCreateNetworkSensorAlertsProfile *RequestSensorCreateNetworkSensorAlertsProfile) (*ResponseSensorCreateNetworkSensorAlertsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseSensorCreateNetworkSensorAlertsProfile](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestSensorCreateNetworkSensorAlertsProfile, nil)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//UpdateDeviceSensorRelationships Assign one or more sensor roles to a given sensor or camera device.
/* Assign one or more sensor roles to a given sensor or camera device.

@param serial serial path parameter.
*/
func (s *SensorService) UpdateDeviceSensorRelationships(serial string, requestSensorUpdateDeviceSensorRelationships *RequestSensorUpdateDeviceSensorRelationships) (*ResponseSensorUpdateDeviceSensorRelationships, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/sensor/relationships"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorUpdateDeviceSensorRelationships](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestSensorUpdateDeviceSensorRelationships)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//UpdateNetworkSensorAlertsProfile Updates a sensor alert profile for a network.
/* Updates a sensor alert profile for a network.

@param networkID networkId path parameter. Network ID
@param id id path parameter.
*/
func (s *SensorService) UpdateNetworkSensorAlertsProfile(networkID string, id string, requestSensorUpdateNetworkSensorAlertsProfile *RequestSensorUpdateNetworkSensorAlertsProfile) (*ResponseSensorUpdateNetworkSensorAlertsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorUpdateNetworkSensorAlertsProfile](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestSensorUpdateNetworkSensorAlertsProfile)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//UpdateNetworkSensorMqttBroker Update the sensor settings of an MQTT broker
/* Update the sensor settings of an MQTT broker. To update the broker itself, use /networks/{networkId}/mqttBrokers/{mqttBrokerId}.

@param networkID networkId path parameter. Network ID
@param mqttBrokerID mqttBrokerId path parameter. Mqtt broker ID
*/
func (s *SensorService) UpdateNetworkSensorMqttBroker(networkID string, mqttBrokerID string, requestSensorUpdateNetworkSensorMqttBroker *RequestSensorUpdateNetworkSensorMqttBroker) (*ResponseSensorUpdateNetworkSensorMqttBroker, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{mqttBrokerId}", fmt.Sprintf("%v", mqttBrokerID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseSensorUpdateNetworkSensorMqttBroker](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestSensorUpdateNetworkSensorMqttBroker)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//DeleteNetworkSensorAlertsProfile Deletes a sensor alert profile from a network.
/* Deletes a sensor alert profile from a network.

@param networkID networkId path parameter. Network ID
@param id id path parameter.


*/
func (s *SensorService) DeleteNetworkSensorAlertsProfile(networkID string, id string) (*resty.Response, error) {
	//networkID string,id string
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
		s.backoff,
	)
}
