package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SensorService service

type GetNetworkSensorAlertsOverviewByMetricQueryParams struct {
	T0       string  `url:"t0,omitempty"`       //The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
	T1       string  `url:"t1,omitempty"`       //The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	Timespan float64 `url:"timespan,omitempty"` //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
	Interval int     `url:"interval,omitempty"` //The time interval in seconds for returned data. The valid intervals are: 86400, 604800. The default is 604800.
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
	Metrics       []string `url:"metrics[],omitempty"`     //Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
}
type GetOrganizationSensorReadingsLatestQueryParams struct {
	PerPage       int      `url:"perPage,omitempty"`       //The number of entries per page returned. Acceptable range is 3 - 100. Default is 100.
	StartingAfter string   `url:"startingAfter,omitempty"` //A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	EndingBefore  string   `url:"endingBefore,omitempty"`  //A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	NetworkIDs    []string `url:"networkIds[],omitempty"`  //Optional parameter to filter readings by network.
	Serials       []string `url:"serials[],omitempty"`     //Optional parameter to filter readings by sensor.
	Metrics       []string `url:"metrics[],omitempty"`     //Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
}

type ResponseSensorGetDeviceSensorRelationships []ResponseItemSensorGetDeviceSensorRelationships // Array of ResponseSensorGetDeviceSensorRelationships
type ResponseItemSensorGetDeviceSensorRelationships struct {
	Livestream *ResponseItemSensorGetDeviceSensorRelationshipsLivestream `json:"livestream,omitempty"` // A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
}
type ResponseItemSensorGetDeviceSensorRelationshipsLivestream struct {
	RelatedDevices *[]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices `json:"relatedDevices,omitempty"` // An array of the related devices for the role
}
type ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices struct {
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
	Door             *int                                                                    `json:"door,omitempty"`             // Number of sensors that are currently alerting due to an open door
	Humidity         *int                                                                    `json:"humidity,omitempty"`         // Number of sensors that are currently alerting due to humidity readings
	IndoorAirQuality *int                                                                    `json:"indoorAirQuality,omitempty"` // Number of sensors that are currently alerting due to indoor air quality readings
	Noise            *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise `json:"noise,omitempty"`            // Object containing the number of sensors that are currently alerting due to noise readings
	Pm25             *int                                                                    `json:"pm25,omitempty"`             // Number of sensors that are currently alerting due to PM2.5 readings
	Temperature      *int                                                                    `json:"temperature,omitempty"`      // Number of sensors that are currently alerting due to temperature readings
	Tvoc             *int                                                                    `json:"tvoc,omitempty"`             // Number of sensors that are currently alerting due to TVOC readings
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
	Door             *int                                                                 `json:"door,omitempty"`             // Number of sensor alerts that occurred due to an open door
	Humidity         *int                                                                 `json:"humidity,omitempty"`         // Number of sensor alerts that occurred due to humidity readings
	IndoorAirQuality *int                                                                 `json:"indoorAirQuality,omitempty"` // Number of sensor alerts that occurred due to indoor air quality readings
	Noise            *ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise `json:"noise,omitempty"`            // Object containing the number of sensor alerts that occurred due to noise readings
	Pm25             *int                                                                 `json:"pm25,omitempty"`             // Number of sensor alerts that occurred due to PM2.5 readings
	Temperature      *int                                                                 `json:"temperature,omitempty"`      // Number of sensor alerts that occurred due to temperature readings
	Tvoc             *int                                                                 `json:"tvoc,omitempty"`             // Number of sensor alerts that occurred due to TVOC readings
	Water            *int                                                                 `json:"water,omitempty"`            // Number of sensor alerts that occurred due to the presence of water
}
type ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise struct {
	Ambient *int `json:"ambient,omitempty"` // Number of sensor alerts that occurred due to ambient noise readings
}
type ResponseSensorGetNetworkSensorAlertsProfiles []ResponseItemSensorGetNetworkSensorAlertsProfiles // Array of ResponseSensorGetNetworkSensorAlertsProfiles
type ResponseItemSensorGetNetworkSensorAlertsProfiles struct {
	Conditions *[]ResponseItemSensorGetNetworkSensorAlertsProfilesConditions `json:"conditions,omitempty"` // List of conditions that will cause the profile to send an alert.
	Name       string                                                        `json:"name,omitempty"`       // Name of the sensor alert profile.
	ProfileID  string                                                        `json:"profileId,omitempty"`  // ID of the sensor alert profile.
	Recipients *ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients   `json:"recipients,omitempty"` // List of recipients that will recieve the alert.
	Schedule   *ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule     `json:"schedule,omitempty"`   // The sensor schedule to use with the alert profile.
	Serials    []string                                                      `json:"serials,omitempty"`    // List of device serials assigned to this sensor alert profile.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditions struct {
	Direction string                                                               `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.
	Duration  *int                                                                 `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0.
	Metric    string                                                               `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
	Threshold *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold struct {
	Door             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Humidity         *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Temperature      *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	Water            *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
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
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
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
	Conditions *[]ResponseSensorCreateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"` // List of conditions that will cause the profile to send an alert.
	Name       string                                                      `json:"name,omitempty"`       // Name of the sensor alert profile.
	ProfileID  string                                                      `json:"profileId,omitempty"`  // ID of the sensor alert profile.
	Recipients *ResponseSensorCreateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"` // List of recipients that will recieve the alert.
	Schedule   *ResponseSensorCreateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`   // The sensor schedule to use with the alert profile.
	Serials    []string                                                    `json:"serials,omitempty"`    // List of device serials assigned to this sensor alert profile.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                             `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.
	Duration  *int                                                               `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0.
	Metric    string                                                             `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
	Threshold *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThreshold struct {
	Door             *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Humidity         *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Temperature      *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	Water            *ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
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
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
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
	Conditions *[]ResponseSensorGetNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"` // List of conditions that will cause the profile to send an alert.
	Name       string                                                   `json:"name,omitempty"`       // Name of the sensor alert profile.
	ProfileID  string                                                   `json:"profileId,omitempty"`  // ID of the sensor alert profile.
	Recipients *ResponseSensorGetNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"` // List of recipients that will recieve the alert.
	Schedule   *ResponseSensorGetNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`   // The sensor schedule to use with the alert profile.
	Serials    []string                                                 `json:"serials,omitempty"`    // List of device serials assigned to this sensor alert profile.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditions struct {
	Direction string                                                          `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.
	Duration  *int                                                            `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0.
	Metric    string                                                          `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
	Threshold *ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold struct {
	Door             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Humidity         *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Temperature      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	Water            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
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
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
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
	Conditions *[]ResponseSensorUpdateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"` // List of conditions that will cause the profile to send an alert.
	Name       string                                                      `json:"name,omitempty"`       // Name of the sensor alert profile.
	ProfileID  string                                                      `json:"profileId,omitempty"`  // ID of the sensor alert profile.
	Recipients *ResponseSensorUpdateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"` // List of recipients that will recieve the alert.
	Schedule   *ResponseSensorUpdateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`   // The sensor schedule to use with the alert profile.
	Serials    []string                                                    `json:"serials,omitempty"`    // List of device serials assigned to this sensor alert profile.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                             `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.
	Duration  *int                                                               `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0.
	Metric    string                                                             `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
	Threshold *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThreshold struct {
	Door             *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Humidity         *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Temperature      *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	Water            *ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
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
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type ResponseSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
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
	Battery          *ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery          `json:"battery,omitempty"`          // Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'.
	Button           *ResponseItemSensorGetOrganizationSensorReadingsHistoryButton           `json:"button,omitempty"`           // Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'.
	Door             *ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor             `json:"door,omitempty"`             // Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'.
	Humidity         *ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity         `json:"humidity,omitempty"`         // Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'.
	IndoorAirQuality *ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'.
	Metric           string                                                                  `json:"metric,omitempty"`           // Type of sensor reading.
	Network          *ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork          `json:"network,omitempty"`          // Network to which the sensor belongs.
	Noise            *ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise            `json:"noise,omitempty"`            // Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'.
	Pm25             *ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25             `json:"pm25,omitempty"`             // Reading for the 'pm25' metric. This will only be present if the 'metric' property equals 'pm25'.
	Serial           string                                                                  `json:"serial,omitempty"`           // Serial number of the sensor that took the reading.
	Temperature      *ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature      `json:"temperature,omitempty"`      // Reading for the 'temperature' metric. This will only be present if the 'metric' property equals 'temperature'.
	Ts               string                                                                  `json:"ts,omitempty"`               // Time at which the reading occurred, in ISO8601 format.
	Tvoc             *ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc             `json:"tvoc,omitempty"`             // Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'.
	Water            *ResponseItemSensorGetOrganizationSensorReadingsHistoryWater            `json:"water,omitempty"`            // Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery struct {
	Percentage *int `json:"percentage,omitempty"` // Remaining battery life.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryButton struct {
	PressType string `json:"pressType,omitempty"` // Type of button press that occurred.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor struct {
	Open *bool `json:"open,omitempty"` // True if the door is open.
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
type ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Temperature reading in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Temperature reading in degrees Fahrenheit.
}
type ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc struct {
	Concentration *int `json:"concentration,omitempty"` // TVOC reading in micrograms per cubic meter.
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
	Battery          *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery          `json:"battery,omitempty"`          // Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'.
	Button           *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton           `json:"button,omitempty"`           // Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'.
	Door             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor             `json:"door,omitempty"`             // Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'.
	Humidity         *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity         `json:"humidity,omitempty"`         // Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'.
	IndoorAirQuality *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'.
	Metric           string                                                                         `json:"metric,omitempty"`           // Type of sensor reading.
	Noise            *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise            `json:"noise,omitempty"`            // Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'.
	Pm25             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25             `json:"pm25,omitempty"`             // Reading for the 'pm25' metric. This will only be present if the 'metric' property equals 'pm25'.
	Temperature      *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature      `json:"temperature,omitempty"`      // Reading for the 'temperature' metric. This will only be present if the 'metric' property equals 'temperature'.
	Ts               string                                                                         `json:"ts,omitempty"`               // Time at which the reading occurred, in ISO8601 format.
	Tvoc             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc             `json:"tvoc,omitempty"`             // Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'.
	Water            *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater            `json:"water,omitempty"`            // Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery struct {
	Percentage *int `json:"percentage,omitempty"` // Remaining battery life.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton struct {
	PressType string `json:"pressType,omitempty"` // Type of button press that occurred.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor struct {
	Open *bool `json:"open,omitempty"` // True if the door is open.
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
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Temperature reading in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Temperature reading in degrees Fahrenheit.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc struct {
	Concentration *int `json:"concentration,omitempty"` // TVOC reading in micrograms per cubic meter.
}
type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater struct {
	Present *bool `json:"present,omitempty"` // True if water is detected.
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
	Conditions *[]RequestSensorCreateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"` // List of conditions that will cause the profile to send an alert.
	Name       string                                                     `json:"name,omitempty"`       // Name of the sensor alert profile.
	Recipients *RequestSensorCreateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"` // List of recipients that will recieve the alert.
	Schedule   *RequestSensorCreateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`   // The sensor schedule to use with the alert profile.
	Serials    []string                                                   `json:"serials,omitempty"`    // List of device serials assigned to this sensor alert profile.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                            `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.
	Duration  *int                                                              `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0.
	Metric    string                                                            `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
	Threshold *RequestSensorCreateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThreshold struct {
	Door             *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Humidity         *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Temperature      *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	Water            *RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
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
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
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
	Conditions *[]RequestSensorUpdateNetworkSensorAlertsProfileConditions `json:"conditions,omitempty"` // List of conditions that will cause the profile to send an alert.
	Name       string                                                     `json:"name,omitempty"`       // Name of the sensor alert profile.
	Recipients *RequestSensorUpdateNetworkSensorAlertsProfileRecipients   `json:"recipients,omitempty"` // List of recipients that will recieve the alert.
	Schedule   *RequestSensorUpdateNetworkSensorAlertsProfileSchedule     `json:"schedule,omitempty"`   // The sensor schedule to use with the alert profile.
	Serials    []string                                                   `json:"serials,omitempty"`    // List of device serials assigned to this sensor alert profile.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditions struct {
	Direction string                                                            `json:"direction,omitempty"` // If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.
	Duration  *int                                                              `json:"duration,omitempty"`  // Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0.
	Metric    string                                                            `json:"metric,omitempty"`    // The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
	Threshold *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold `json:"threshold,omitempty"` // Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold struct {
	Door             *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor             `json:"door,omitempty"`             // Door open threshold. 'open' must be provided and set to true.
	Humidity         *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity         `json:"humidity,omitempty"`         // Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
	IndoorAirQuality *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"` // Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
	Noise            *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise            `json:"noise,omitempty"`            // Noise threshold. 'ambient' must be provided.
	Pm25             *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25             `json:"pm25,omitempty"`             // PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
	Temperature      *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature      `json:"temperature,omitempty"`      // Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
	Tvoc             *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc             `json:"tvoc,omitempty"`             // TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
	Water            *RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater            `json:"water,omitempty"`            // Water detection threshold. 'present' must be provided and set to true.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open *bool `json:"open,omitempty"` // Alerting threshold for a door open event. Must be set to true.
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
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    *float64 `json:"celsius,omitempty"`    // Alerting threshold in degrees Celsius.
	Fahrenheit *float64 `json:"fahrenheit,omitempty"` // Alerting threshold in degrees Fahrenheit.
	Quality    string   `json:"quality,omitempty"`    // Alerting threshold as a qualitative temperature level.
}
type RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration *int   `json:"concentration,omitempty"` // Alerting threshold as TVOC micrograms per cubic meter.
	Quality       string `json:"quality,omitempty"`       // Alerting threshold as a qualitative TVOC level.
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

//GetDeviceSensorRelationships List the sensor roles for a given sensor or camera device.
/* List the sensor roles for a given sensor or camera device.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-sensor-relationships
*/
func (s *SensorService) GetDeviceSensorRelationships(serial string) (*ResponseSensorGetDeviceSensorRelationships, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/sensor/relationships"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorGetDeviceSensorRelationships{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSensorRelationships")
	}

	result := response.Result().(*ResponseSensorGetDeviceSensorRelationships)
	return result, response, err

}

//GetNetworkSensorAlertsCurrentOverviewByMetric Return an overview of currently alerting sensors by metric
/* Return an overview of currently alerting sensors by metric

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sensor-alerts-current-overview-by-metric
*/
func (s *SensorService) GetNetworkSensorAlertsCurrentOverviewByMetric(networkID string) (*ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/current/overview/byMetric"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSensorAlertsCurrentOverviewByMetric")
	}

	result := response.Result().(*ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric)
	return result, response, err

}

//GetNetworkSensorAlertsOverviewByMetric Return an overview of alert occurrences over a timespan, by metric
/* Return an overview of alert occurrences over a timespan, by metric

@param networkID networkId path parameter. Network ID
@param getNetworkSensorAlertsOverviewByMetricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sensor-alerts-overview-by-metric
*/
func (s *SensorService) GetNetworkSensorAlertsOverviewByMetric(networkID string, getNetworkSensorAlertsOverviewByMetricQueryParams *GetNetworkSensorAlertsOverviewByMetricQueryParams) (*ResponseSensorGetNetworkSensorAlertsOverviewByMetric, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/overview/byMetric"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	queryString, _ := query.Values(getNetworkSensorAlertsOverviewByMetricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorGetNetworkSensorAlertsOverviewByMetric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSensorAlertsOverviewByMetric")
	}

	result := response.Result().(*ResponseSensorGetNetworkSensorAlertsOverviewByMetric)
	return result, response, err

}

//GetNetworkSensorAlertsProfiles Lists all sensor alert profiles for a network.
/* Lists all sensor alert profiles for a network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sensor-alerts-profiles
*/
func (s *SensorService) GetNetworkSensorAlertsProfiles(networkID string) (*ResponseSensorGetNetworkSensorAlertsProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorGetNetworkSensorAlertsProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSensorAlertsProfiles")
	}

	result := response.Result().(*ResponseSensorGetNetworkSensorAlertsProfiles)
	return result, response, err

}

//GetNetworkSensorAlertsProfile Show details of a sensor alert profile for a network.
/* Show details of a sensor alert profile for a network.

@param networkID networkId path parameter. Network ID
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sensor-alerts-profile
*/
func (s *SensorService) GetNetworkSensorAlertsProfile(networkID string, id string) (*ResponseSensorGetNetworkSensorAlertsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorGetNetworkSensorAlertsProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSensorAlertsProfile")
	}

	result := response.Result().(*ResponseSensorGetNetworkSensorAlertsProfile)
	return result, response, err

}

//GetNetworkSensorMqttBrokers List the sensor settings of all MQTT brokers for this network
/* List the sensor settings of all MQTT brokers for this network. To get the brokers themselves, use /networks/{networkId}/mqttBrokers.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sensor-mqtt-brokers
*/
func (s *SensorService) GetNetworkSensorMqttBrokers(networkID string) (*ResponseSensorGetNetworkSensorMqttBrokers, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/mqttBrokers"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorGetNetworkSensorMqttBrokers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSensorMqttBrokers")
	}

	result := response.Result().(*ResponseSensorGetNetworkSensorMqttBrokers)
	return result, response, err

}

//GetNetworkSensorMqttBroker Return the sensor settings of an MQTT broker
/* Return the sensor settings of an MQTT broker. To get the broker itself, use /networks/{networkId}/mqttBrokers/{mqttBrokerId}.

@param networkID networkId path parameter. Network ID
@param mqttBrokerID mqttBrokerId path parameter. Mqtt broker ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sensor-mqtt-broker
*/
func (s *SensorService) GetNetworkSensorMqttBroker(networkID string, mqttBrokerID string) (*ResponseSensorGetNetworkSensorMqttBroker, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{mqttBrokerId}", fmt.Sprintf("%v", mqttBrokerID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorGetNetworkSensorMqttBroker{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSensorMqttBroker")
	}

	result := response.Result().(*ResponseSensorGetNetworkSensorMqttBroker)
	return result, response, err

}

//GetNetworkSensorRelationships List the sensor roles for devices in a given network
/* List the sensor roles for devices in a given network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-sensor-relationships
*/
func (s *SensorService) GetNetworkSensorRelationships(networkID string) (*ResponseSensorGetNetworkSensorRelationships, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/relationships"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorGetNetworkSensorRelationships{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkSensorRelationships")
	}

	result := response.Result().(*ResponseSensorGetNetworkSensorRelationships)
	return result, response, err

}

//GetOrganizationSensorReadingsHistory Return all reported readings from sensors in a given timespan, sorted by timestamp
/* Return all reported readings from sensors in a given timespan, sorted by timestamp

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSensorReadingsHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-sensor-readings-history
*/
func (s *SensorService) GetOrganizationSensorReadingsHistory(organizationID string, getOrganizationSensorReadingsHistoryQueryParams *GetOrganizationSensorReadingsHistoryQueryParams) (*ResponseSensorGetOrganizationSensorReadingsHistory, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/sensor/readings/history"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSensorReadingsHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorGetOrganizationSensorReadingsHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSensorReadingsHistory")
	}

	result := response.Result().(*ResponseSensorGetOrganizationSensorReadingsHistory)
	return result, response, err

}

//GetOrganizationSensorReadingsLatest Return the latest available reading for each metric from each sensor, sorted by sensor serial
/* Return the latest available reading for each metric from each sensor, sorted by sensor serial

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationSensorReadingsLatestQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-sensor-readings-latest
*/
func (s *SensorService) GetOrganizationSensorReadingsLatest(organizationID string, getOrganizationSensorReadingsLatestQueryParams *GetOrganizationSensorReadingsLatestQueryParams) (*ResponseSensorGetOrganizationSensorReadingsLatest, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/sensor/readings/latest"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationSensorReadingsLatestQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorGetOrganizationSensorReadingsLatest{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationSensorReadingsLatest")
	}

	result := response.Result().(*ResponseSensorGetOrganizationSensorReadingsLatest)
	return result, response, err

}

//CreateNetworkSensorAlertsProfile Creates a sensor alert profile for a network.
/* Creates a sensor alert profile for a network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-sensor-alerts-profile
*/

func (s *SensorService) CreateNetworkSensorAlertsProfile(networkID string, requestSensorCreateNetworkSensorAlertsProfile *RequestSensorCreateNetworkSensorAlertsProfile) (*ResponseSensorCreateNetworkSensorAlertsProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorCreateNetworkSensorAlertsProfile).
		SetResult(&ResponseSensorCreateNetworkSensorAlertsProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetworkSensorAlertsProfile")
	}

	result := response.Result().(*ResponseSensorCreateNetworkSensorAlertsProfile)
	return result, response, err

}

//UpdateDeviceSensorRelationships Assign one or more sensor roles to a given sensor or camera device.
/* Assign one or more sensor roles to a given sensor or camera device.

@param serial serial path parameter.
*/
func (s *SensorService) UpdateDeviceSensorRelationships(serial string, requestSensorUpdateDeviceSensorRelationships *RequestSensorUpdateDeviceSensorRelationships) (*ResponseSensorUpdateDeviceSensorRelationships, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/sensor/relationships"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorUpdateDeviceSensorRelationships).
		SetResult(&ResponseSensorUpdateDeviceSensorRelationships{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceSensorRelationships")
	}

	result := response.Result().(*ResponseSensorUpdateDeviceSensorRelationships)
	return result, response, err

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

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorUpdateNetworkSensorAlertsProfile).
		SetResult(&ResponseSensorUpdateNetworkSensorAlertsProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSensorAlertsProfile")
	}

	result := response.Result().(*ResponseSensorUpdateNetworkSensorAlertsProfile)
	return result, response, err

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

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorUpdateNetworkSensorMqttBroker).
		SetResult(&ResponseSensorUpdateNetworkSensorMqttBroker{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetworkSensorMqttBroker")
	}

	result := response.Result().(*ResponseSensorUpdateNetworkSensorMqttBroker)
	return result, response, err

}

//DeleteNetworkSensorAlertsProfile Deletes a sensor alert profile from a network.
/* Deletes a sensor alert profile from a network.

@param networkID networkId path parameter. Network ID
@param id id path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-sensor-alerts-profile
*/
func (s *SensorService) DeleteNetworkSensorAlertsProfile(networkID string, id string) (*resty.Response, error) {
	//networkID string,id string
	path := "/api/v1/networks/{networkId}/sensor/alerts/profiles/{id}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkSensorAlertsProfile")
	}

	return response, err

}
