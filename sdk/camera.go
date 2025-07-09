package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type CameraService service

type GetDeviceCameraAnalyticsOverviewQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour.
	ObjectType string  `url:"objectType,omitempty"` //[optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
}
type GetDeviceCameraAnalyticsRecentQueryParams struct {
	ObjectType string `url:"objectType,omitempty"` //[optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
}
type GetDeviceCameraAnalyticsZoneHistoryQueryParams struct {
	T0         string  `url:"t0,omitempty"`         //The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
	T1         string  `url:"t1,omitempty"`         //The end of the timespan for the data. t1 can be a maximum of 14 hours after t0.
	Timespan   float64 `url:"timespan,omitempty"`   //The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour.
	Resolution int     `url:"resolution,omitempty"` //The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60.
	ObjectType string  `url:"objectType,omitempty"` //[optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
}
type GetDeviceCameraVideoLinkQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` //[optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time.
}
type GetOrganizationCameraBoundariesAreasByDeviceQueryParams struct {
	Serials []string `url:"serials[],omitempty"` //A list of serial numbers. The returned cameras will be filtered to only include these serials.
}
type GetOrganizationCameraBoundariesLinesByDeviceQueryParams struct {
	Serials []string `url:"serials[],omitempty"` //A list of serial numbers. The returned cameras will be filtered to only include these serials.
}
type GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalQueryParams struct {
	BoundaryIDs   []string `url:"boundaryIds[],omitempty"`   //A list of boundary ids. The returned cameras will be filtered to only include these ids.
	Ranges        []string `url:"ranges[],omitempty"`        //A list of time ranges with intervals
	Duration      int      `url:"duration,omitempty"`        //The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60.
	PerPage       int      `url:"perPage,omitempty"`         //The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000.
	BoundaryTypes []string `url:"boundaryTypes[],omitempty"` //The detection types. Defaults to 'person'.
}
type GetOrganizationCameraOnboardingStatusesQueryParams struct {
	Serials    []string `url:"serials[],omitempty"`    //A list of serial numbers. The returned cameras will be filtered to only include these serials.
	NetworkIDs []string `url:"networkIds[],omitempty"` //A list of network IDs. The returned cameras will be filtered to only include these networks.
}

type ResponseCameraGetDeviceCameraAnalyticsLive struct {
	Ts    string                                           `json:"ts,omitempty"`    // The current time
	Zones *ResponseCameraGetDeviceCameraAnalyticsLiveZones `json:"zones,omitempty"` // The zones state
}
type ResponseCameraGetDeviceCameraAnalyticsLiveZones struct {
	ZoneID *ResponseCameraGetDeviceCameraAnalyticsLiveZonesZoneID `json:"zoneId,omitempty"` // The zone state, dynamic
}
type ResponseCameraGetDeviceCameraAnalyticsLiveZonesZoneID struct {
	Person *int `json:"person,omitempty"` // The count per type, dynamic
}
type ResponseCameraGetDeviceCameraAnalyticsOverview []ResponseItemCameraGetDeviceCameraAnalyticsOverview // Array of ResponseCameraGetDeviceCameraAnalyticsOverview
type ResponseItemCameraGetDeviceCameraAnalyticsOverview struct {
	AverageCount *float64 `json:"averageCount,omitempty"` // The average count
	EndTs        string   `json:"endTs,omitempty"`        // The end time
	Entrances    *int     `json:"entrances,omitempty"`    // The number of sentrances
	StartTs      string   `json:"startTs,omitempty"`      // The start time
	ZoneID       *int     `json:"zoneId,omitempty"`       // The zone id
}
type ResponseCameraGetDeviceCameraAnalyticsRecent []ResponseItemCameraGetDeviceCameraAnalyticsRecent // Array of ResponseCameraGetDeviceCameraAnalyticsRecent
type ResponseItemCameraGetDeviceCameraAnalyticsRecent struct {
	AverageCount *float64 `json:"averageCount,omitempty"` // The average count
	EndTs        string   `json:"endTs,omitempty"`        // The end time
	Entrances    *int     `json:"entrances,omitempty"`    // The number of entrances
	StartTs      string   `json:"startTs,omitempty"`      // The start time
	ZoneID       *int     `json:"zoneId,omitempty"`       // The zone id
}
type ResponseCameraGetDeviceCameraAnalyticsZones []ResponseItemCameraGetDeviceCameraAnalyticsZones // Array of ResponseCameraGetDeviceCameraAnalyticsZones
type ResponseItemCameraGetDeviceCameraAnalyticsZones struct {
	ID               string                                                           `json:"id,omitempty"`               // The zone ID
	Label            string                                                           `json:"label,omitempty"`            // The zone label
	RegionOfInterest *ResponseItemCameraGetDeviceCameraAnalyticsZonesRegionOfInterest `json:"regionOfInterest,omitempty"` // The region of interest
	Type             string                                                           `json:"type,omitempty"`             // The zone type
}
type ResponseItemCameraGetDeviceCameraAnalyticsZonesRegionOfInterest struct {
	X0 string `json:"x0,omitempty"` // The x0 coordinate
	X1 string `json:"x1,omitempty"` // The x1 coordinate
	Y0 string `json:"y0,omitempty"` // The y0 coordinate
	Y1 string `json:"y1,omitempty"` // The y1 coordinate
}
type ResponseCameraGetDeviceCameraAnalyticsZoneHistory []ResponseItemCameraGetDeviceCameraAnalyticsZoneHistory // Array of ResponseCameraGetDeviceCameraAnalyticsZoneHistory
type ResponseItemCameraGetDeviceCameraAnalyticsZoneHistory struct {
	AverageCount *float64 `json:"averageCount,omitempty"` // The average count
	EndTs        string   `json:"endTs,omitempty"`        // The end time
	Entrances    *int     `json:"entrances,omitempty"`    // The number of entrances
	StartTs      string   `json:"startTs,omitempty"`      // The start time
}
type ResponseCameraGetDeviceCameraCustomAnalytics struct {
	ArtifactID string                                                    `json:"artifactId,omitempty"` // Custom analytics artifact ID
	Enabled    *bool                                                     `json:"enabled,omitempty"`    // Whether custom analytics is enabled
	Parameters *[]ResponseCameraGetDeviceCameraCustomAnalyticsParameters `json:"parameters,omitempty"` // Parameters for the custom analytics workload
}
type ResponseCameraGetDeviceCameraCustomAnalyticsParameters struct {
	Name  string   `json:"name,omitempty"`  // Name of the parameter
	Value *float64 `json:"value,omitempty"` // Value of the parameter
}
type ResponseCameraUpdateDeviceCameraCustomAnalytics struct {
	ArtifactID string                                                       `json:"artifactId,omitempty"` // Custom analytics artifact ID
	Enabled    *bool                                                        `json:"enabled,omitempty"`    // Whether custom analytics is enabled
	Parameters *[]ResponseCameraUpdateDeviceCameraCustomAnalyticsParameters `json:"parameters,omitempty"` // Parameters for the custom analytics workload
}
type ResponseCameraUpdateDeviceCameraCustomAnalyticsParameters struct {
	Name  string   `json:"name,omitempty"`  // Name of the parameter
	Value *float64 `json:"value,omitempty"` // Value of the parameter
}
type ResponseCameraGenerateDeviceCameraSnapshot struct {
	Expiry string `json:"expiry,omitempty"` // Expiration details for snapshot image access
	URL    string `json:"url,omitempty"`    // Url for the snapshot
}
type ResponseCameraGetDeviceCameraQualityAndRetention struct {
	AudioRecordingEnabled          *bool  `json:"audioRecordingEnabled,omitempty"`          //
	MotionBasedRetentionEnabled    *bool  `json:"motionBasedRetentionEnabled,omitempty"`    //
	MotionDetectorVersion          *int   `json:"motionDetectorVersion,omitempty"`          //
	ProfileID                      string `json:"profileId,omitempty"`                      //
	Quality                        string `json:"quality,omitempty"`                        //
	Resolution                     string `json:"resolution,omitempty"`                     //
	RestrictedBandwidthModeEnabled *bool  `json:"restrictedBandwidthModeEnabled,omitempty"` //
}
type ResponseCameraUpdateDeviceCameraQualityAndRetention interface{}
type ResponseCameraGetDeviceCameraSense struct {
	AudioDetection *ResponseCameraGetDeviceCameraSenseAudioDetection `json:"audioDetection,omitempty"` //
	MqttBrokerID   string                                            `json:"mqttBrokerId,omitempty"`   //
	MqttTopics     []string                                          `json:"mqttTopics,omitempty"`     //
	SenseEnabled   *bool                                             `json:"senseEnabled,omitempty"`   //
}
type ResponseCameraGetDeviceCameraSenseAudioDetection struct {
	Enabled *bool `json:"enabled,omitempty"` //
}
type ResponseCameraUpdateDeviceCameraSense interface{}
type ResponseCameraGetDeviceCameraSenseObjectDetectionModels []ResponseItemCameraGetDeviceCameraSenseObjectDetectionModels // Array of ResponseCameraGetDeviceCameraSenseObjectDetectionModels
type ResponseItemCameraGetDeviceCameraSenseObjectDetectionModels struct {
	Description string `json:"description,omitempty"` //
	ID          string `json:"id,omitempty"`          //
}
type ResponseCameraGetDeviceCameraVideoSettings struct {
	ExternalRtspEnabled *bool  `json:"externalRtspEnabled,omitempty"` // Boolean indicating if external rtsp stream is exposed
	RtspURL             string `json:"rtspUrl,omitempty"`             // External rstp url. Will only be returned if external rtsp stream is exposed
}
type ResponseCameraUpdateDeviceCameraVideoSettings struct {
	ExternalRtspEnabled *bool  `json:"externalRtspEnabled,omitempty"` // Boolean indicating if external rtsp stream is exposed
	RtspURL             string `json:"rtspUrl,omitempty"`             // External rstp url. Will only be returned if external rtsp stream is exposed
}
type ResponseCameraGetDeviceCameraVideoLink struct {
	URL       string `json:"url,omitempty"`       //
	VisionURL string `json:"visionUrl,omitempty"` //
}
type ResponseCameraGetDeviceCameraWirelessProfiles struct {
	IDs *ResponseCameraGetDeviceCameraWirelessProfilesIDs `json:"ids,omitempty"` //
}
type ResponseCameraGetDeviceCameraWirelessProfilesIDs struct {
	Backup    string `json:"backup,omitempty"`    //
	Primary   string `json:"primary,omitempty"`   //
	Secondary string `json:"secondary,omitempty"` //
}
type ResponseCameraUpdateDeviceCameraWirelessProfiles interface{}
type ResponseCameraGetNetworkCameraQualityRetentionProfiles []ResponseItemCameraGetNetworkCameraQualityRetentionProfiles // Array of ResponseCameraGetNetworkCameraQualityRetentionProfiles
type ResponseItemCameraGetNetworkCameraQualityRetentionProfiles struct {
	AudioRecordingEnabled          *bool                                                                     `json:"audioRecordingEnabled,omitempty"`          //
	CloudArchiveEnabled            *bool                                                                     `json:"cloudArchiveEnabled,omitempty"`            //
	ID                             string                                                                    `json:"id,omitempty"`                             //
	MaxRetentionDays               *int                                                                      `json:"maxRetentionDays,omitempty"`               //
	MotionBasedRetentionEnabled    *bool                                                                     `json:"motionBasedRetentionEnabled,omitempty"`    //
	MotionDetectorVersion          *int                                                                      `json:"motionDetectorVersion,omitempty"`          //
	Name                           string                                                                    `json:"name,omitempty"`                           //
	NetworkID                      string                                                                    `json:"networkId,omitempty"`                      //
	RestrictedBandwidthModeEnabled *bool                                                                     `json:"restrictedBandwidthModeEnabled,omitempty"` //
	ScheduleID                     string                                                                    `json:"scheduleId,omitempty"`                     //
	SmartRetention                 *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesSmartRetention `json:"smartRetention,omitempty"`                 //
	VideoSettings                  *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings  `json:"videoSettings,omitempty"`                  //
}
type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesSmartRetention struct {
	Enabled *bool `json:"enabled,omitempty"` //
}
type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings struct {
	MV12MV22MV72 *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"` //
	MV12WE       *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12WE       `json:"MV12WE,omitempty"`         //
	MV21MV71     *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV21MV71     `json:"MV21/MV71,omitempty"`      //
	MV32         *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV32         `json:"MV32,omitempty"`           //
}
type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12WE struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV21MV71 struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV32 struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseCameraCreateNetworkCameraQualityRetentionProfile interface{}
type ResponseCameraGetNetworkCameraQualityRetentionProfile struct {
	AudioRecordingEnabled          *bool                                                                `json:"audioRecordingEnabled,omitempty"`          //
	CloudArchiveEnabled            *bool                                                                `json:"cloudArchiveEnabled,omitempty"`            //
	ID                             string                                                               `json:"id,omitempty"`                             //
	MaxRetentionDays               *int                                                                 `json:"maxRetentionDays,omitempty"`               //
	MotionBasedRetentionEnabled    *bool                                                                `json:"motionBasedRetentionEnabled,omitempty"`    //
	MotionDetectorVersion          *int                                                                 `json:"motionDetectorVersion,omitempty"`          //
	Name                           string                                                               `json:"name,omitempty"`                           //
	NetworkID                      string                                                               `json:"networkId,omitempty"`                      //
	RestrictedBandwidthModeEnabled *bool                                                                `json:"restrictedBandwidthModeEnabled,omitempty"` //
	ScheduleID                     string                                                               `json:"scheduleId,omitempty"`                     //
	SmartRetention                 *ResponseCameraGetNetworkCameraQualityRetentionProfileSmartRetention `json:"smartRetention,omitempty"`                 //
	VideoSettings                  *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings  `json:"videoSettings,omitempty"`                  //
}
type ResponseCameraGetNetworkCameraQualityRetentionProfileSmartRetention struct {
	Enabled *bool `json:"enabled,omitempty"` //
}
type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings struct {
	MV12MV22MV72 *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"` //
	MV12WE       *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WE       `json:"MV12WE,omitempty"`         //
	MV21MV71     *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71     `json:"MV21/MV71,omitempty"`      //
	MV32         *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32         `json:"MV32,omitempty"`           //
}
type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WE struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32 struct {
	Quality    string `json:"quality,omitempty"`    //
	Resolution string `json:"resolution,omitempty"` //
}
type ResponseCameraUpdateNetworkCameraQualityRetentionProfile interface{}
type ResponseCameraGetNetworkCameraSchedules []ResponseItemCameraGetNetworkCameraSchedules // Array of ResponseCameraGetNetworkCameraSchedules
type ResponseItemCameraGetNetworkCameraSchedules struct {
	ID   string `json:"id,omitempty"`   // Schedule id
	Name string `json:"name,omitempty"` // Schedule name
}
type ResponseCameraGetNetworkCameraWirelessProfiles []ResponseItemCameraGetNetworkCameraWirelessProfiles // Array of ResponseCameraGetNetworkCameraWirelessProfiles
type ResponseItemCameraGetNetworkCameraWirelessProfiles struct {
	AppliedDeviceCount *int                                                        `json:"appliedDeviceCount,omitempty"` // The count of the applied devices.
	ID                 string                                                      `json:"id,omitempty"`                 // The ID of the camera wireless profile.
	IDentity           *ResponseItemCameraGetNetworkCameraWirelessProfilesIDentity `json:"identity,omitempty"`           // The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
	Name               string                                                      `json:"name,omitempty"`               // The name of the camera wireless profile.
	SSID               *ResponseItemCameraGetNetworkCameraWirelessProfilesSSID     `json:"ssid,omitempty"`               // The details of the SSID config.
}
type ResponseItemCameraGetNetworkCameraWirelessProfilesIDentity struct {
	Password string `json:"password,omitempty"` // The password of the identity.
	Username string `json:"username,omitempty"` // The username of the identity.
}
type ResponseItemCameraGetNetworkCameraWirelessProfilesSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       // The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	EncryptionMode string `json:"encryptionMode,omitempty"` // The encryption mode of the SSID.
	Name           string `json:"name,omitempty"`           // The name of the SSID.
	Psk            string `json:"psk,omitempty"`            // The pre-shared key of the SSID, if mode is PSK
}
type ResponseCameraCreateNetworkCameraWirelessProfile struct {
	AppliedDeviceCount *int                                                      `json:"appliedDeviceCount,omitempty"` // The count of the applied devices.
	ID                 string                                                    `json:"id,omitempty"`                 // The ID of the camera wireless profile.
	IDentity           *ResponseCameraCreateNetworkCameraWirelessProfileIDentity `json:"identity,omitempty"`           // The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
	Name               string                                                    `json:"name,omitempty"`               // The name of the camera wireless profile.
	SSID               *ResponseCameraCreateNetworkCameraWirelessProfileSSID     `json:"ssid,omitempty"`               // The details of the SSID config.
}
type ResponseCameraCreateNetworkCameraWirelessProfileIDentity struct {
	Password string `json:"password,omitempty"` // The password of the identity.
	Username string `json:"username,omitempty"` // The username of the identity.
}
type ResponseCameraCreateNetworkCameraWirelessProfileSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       // The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	EncryptionMode string `json:"encryptionMode,omitempty"` // The encryption mode of the SSID.
	Name           string `json:"name,omitempty"`           // The name of the SSID.
	Psk            string `json:"psk,omitempty"`            // The pre-shared key of the SSID, if mode is PSK
}
type ResponseCameraGetNetworkCameraWirelessProfile struct {
	AppliedDeviceCount *int                                                   `json:"appliedDeviceCount,omitempty"` // The count of the applied devices.
	ID                 string                                                 `json:"id,omitempty"`                 // The ID of the camera wireless profile.
	IDentity           *ResponseCameraGetNetworkCameraWirelessProfileIDentity `json:"identity,omitempty"`           // The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
	Name               string                                                 `json:"name,omitempty"`               // The name of the camera wireless profile.
	SSID               *ResponseCameraGetNetworkCameraWirelessProfileSSID     `json:"ssid,omitempty"`               // The details of the SSID config.
}
type ResponseCameraGetNetworkCameraWirelessProfileIDentity struct {
	Password string `json:"password,omitempty"` // The password of the identity.
	Username string `json:"username,omitempty"` // The username of the identity.
}
type ResponseCameraGetNetworkCameraWirelessProfileSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       // The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	EncryptionMode string `json:"encryptionMode,omitempty"` // The encryption mode of the SSID.
	Name           string `json:"name,omitempty"`           // The name of the SSID.
	Psk            string `json:"psk,omitempty"`            // The pre-shared key of the SSID, if mode is PSK
}
type ResponseCameraUpdateNetworkCameraWirelessProfile struct {
	AppliedDeviceCount *int                                                      `json:"appliedDeviceCount,omitempty"` // The count of the applied devices.
	ID                 string                                                    `json:"id,omitempty"`                 // The ID of the camera wireless profile.
	IDentity           *ResponseCameraUpdateNetworkCameraWirelessProfileIDentity `json:"identity,omitempty"`           // The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
	Name               string                                                    `json:"name,omitempty"`               // The name of the camera wireless profile.
	SSID               *ResponseCameraUpdateNetworkCameraWirelessProfileSSID     `json:"ssid,omitempty"`               // The details of the SSID config.
}
type ResponseCameraUpdateNetworkCameraWirelessProfileIDentity struct {
	Password string `json:"password,omitempty"` // The password of the identity.
	Username string `json:"username,omitempty"` // The username of the identity.
}
type ResponseCameraUpdateNetworkCameraWirelessProfileSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       // The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	EncryptionMode string `json:"encryptionMode,omitempty"` // The encryption mode of the SSID.
	Name           string `json:"name,omitempty"`           // The name of the SSID.
	Psk            string `json:"psk,omitempty"`            // The pre-shared key of the SSID, if mode is PSK
}
type ResponseCameraGetOrganizationCameraBoundariesAreasByDevice []ResponseItemCameraGetOrganizationCameraBoundariesAreasByDevice // Array of ResponseCameraGetOrganizationCameraBoundariesAreasByDevice
type ResponseItemCameraGetOrganizationCameraBoundariesAreasByDevice struct {
	Boundaries *ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundaries `json:"boundaries,omitempty"` // Configured area boundaries of the camera
	NetworkID  string                                                                    `json:"networkId,omitempty"`  // The network id of the camera
	Serial     string                                                                    `json:"serial,omitempty"`     // The serial number of the camera
}
type ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundaries struct {
	ID       string                                                                              `json:"id,omitempty"`       // The area boundary id
	Name     string                                                                              `json:"name,omitempty"`     // The area boundary name
	Type     string                                                                              `json:"type,omitempty"`     // The area boundary type
	Vertices *[]ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices `json:"vertices,omitempty"` // The area boundary vertices
}
type ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices struct {
	X *float64 `json:"x,omitempty"` // The vertex x coordinate
	Y *float64 `json:"y,omitempty"` // The vertex y coordinate
}
type ResponseCameraGetOrganizationCameraBoundariesLinesByDevice []ResponseItemCameraGetOrganizationCameraBoundariesLinesByDevice // Array of ResponseCameraGetOrganizationCameraBoundariesLinesByDevice
type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDevice struct {
	Boundaries *ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundaries `json:"boundaries,omitempty"` // Configured line boundaries of the camera
	NetworkID  string                                                                    `json:"networkId,omitempty"`  // The network id of the camera
	Serial     string                                                                    `json:"serial,omitempty"`     // The serial number of the camera
}
type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundaries struct {
	DirectionVertex *ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesDirectionVertex `json:"directionVertex,omitempty"` // The line boundary crossing direction vertex
	ID              string                                                                                   `json:"id,omitempty"`              // The line boundary id
	Name            string                                                                                   `json:"name,omitempty"`            // The line boundary name
	Type            string                                                                                   `json:"type,omitempty"`            // The line boundary type
	Vertices        *[]ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesVertices      `json:"vertices,omitempty"`        // The line boundary vertices
}
type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesDirectionVertex struct {
	X *float64 `json:"x,omitempty"` // The vertex x coordinate
	Y *float64 `json:"y,omitempty"` // The vertex y coordinate
}
type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesVertices struct {
	X *float64 `json:"x,omitempty"` // The vertex x coordinate
	Y *float64 `json:"y,omitempty"` // The vertex y coordinate
}
type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts []ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts // Array of ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts
type ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts struct {
	ArtifactID     string                                                                 `json:"artifactId,omitempty"`     // Custom analytics artifact ID
	Name           string                                                                 `json:"name,omitempty"`           // Custom analytics artifact name
	OrganizationID string                                                                 `json:"organizationId,omitempty"` // Organization ID
	Status         *ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus `json:"status,omitempty"`         // Custom analytics artifact status
}
type ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus struct {
	Message string `json:"message,omitempty"` // Status message
	Type    string `json:"type,omitempty"`    // Status type
}
type ResponseCameraCreateOrganizationCameraCustomAnalyticsArtifact struct {
	ArtifactID      string                                                               `json:"artifactId,omitempty"`      // Custom analytics artifact ID
	Name            string                                                               `json:"name,omitempty"`            // Custom analytics artifact name
	OrganizationID  string                                                               `json:"organizationId,omitempty"`  // Organization ID
	Status          *ResponseCameraCreateOrganizationCameraCustomAnalyticsArtifactStatus `json:"status,omitempty"`          // Custom analytics artifact status
	UploadID        string                                                               `json:"uploadId,omitempty"`        // Upload ID
	UploadURL       string                                                               `json:"uploadUrl,omitempty"`       // Upload URL
	UploadURLExpiry string                                                               `json:"uploadUrlExpiry,omitempty"` // Upload URL expiry time
}
type ResponseCameraCreateOrganizationCameraCustomAnalyticsArtifactStatus struct {
	Message string `json:"message,omitempty"` // Status message
	Type    string `json:"type,omitempty"`    // Status type
}
type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact struct {
	ArtifactID     string                                                            `json:"artifactId,omitempty"`     // Custom analytics artifact ID
	Name           string                                                            `json:"name,omitempty"`           // Custom analytics artifact name
	OrganizationID string                                                            `json:"organizationId,omitempty"` // Organization ID
	Status         *ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus `json:"status,omitempty"`         // Custom analytics artifact status
}
type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus struct {
	Message string `json:"message,omitempty"` // Status message
	Type    string `json:"type,omitempty"`    // Status type
}
type ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval []ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval // Array of ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval
type ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval struct {
	BoundaryID string                                                                               `json:"boundaryId,omitempty"` // The boundary id
	Results    *ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalResults `json:"results,omitempty"`    // The analytics data
	Type       string                                                                               `json:"type,omitempty"`       // The boundary type
}
type ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalResults struct {
	EndTime    string `json:"endTime,omitempty"`    // The period end time
	In         *int   `json:"in,omitempty"`         // The number of detections entered
	ObjectType string `json:"objectType,omitempty"` // The detection type
	Out        *int   `json:"out,omitempty"`        // The number of detections exited
	StartTime  string `json:"startTime,omitempty"`  // The period start time
}
type ResponseCameraGetOrganizationCameraOnboardingStatuses []ResponseItemCameraGetOrganizationCameraOnboardingStatuses // Array of ResponseCameraGetOrganizationCameraOnboardingStatuses
type ResponseItemCameraGetOrganizationCameraOnboardingStatuses struct {
	NetworkID string `json:"networkId,omitempty"` //
	Serial    string `json:"serial,omitempty"`    //
	Status    string `json:"status,omitempty"`    //
	UpdatedAt string `json:"updatedAt,omitempty"` //
}
type ResponseCameraUpdateOrganizationCameraOnboardingStatuses interface{}
type ResponseCameraGetOrganizationCameraPermissions []ResponseItemCameraGetOrganizationCameraPermissions // Array of ResponseCameraGetOrganizationCameraPermissions
type ResponseItemCameraGetOrganizationCameraPermissions struct {
	ID    string `json:"id,omitempty"`    // Permission scope id
	Level string `json:"level,omitempty"` // Permission scope level
	Name  string `json:"name,omitempty"`  // Name of permission scope
}
type ResponseCameraGetOrganizationCameraPermission struct {
	ID    string `json:"id,omitempty"`    // Permission scope id
	Level string `json:"level,omitempty"` // Permission scope level
	Name  string `json:"name,omitempty"`  // Name of permission scope
}
type ResponseCameraGetOrganizationCameraRoles []ResponseItemCameraGetOrganizationCameraRoles // Array of ResponseCameraGetOrganizationCameraRoles
type ResponseItemCameraGetOrganizationCameraRoles struct {
	ID                string                                                           `json:"id,omitempty"`                //
	AppliedOnDevices  *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOnDevices  `json:"appliedOnDevices,omitempty"`  //
	AppliedOnNetworks *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOnNetworks `json:"appliedOnNetworks,omitempty"` //
	AppliedOrgWide    *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOrgWide    `json:"appliedOrgWide,omitempty"`    //
	Name              string                                                           `json:"name,omitempty"`              //
}
type ResponseItemCameraGetOrganizationCameraRolesAppliedOnDevices struct {
	ID                string `json:"id,omitempty"`                //
	PermissionLevel   string `json:"permissionLevel,omitempty"`   //
	PermissionScope   string `json:"permissionScope,omitempty"`   //
	PermissionScopeID string `json:"permissionScopeId,omitempty"` //
	Tag               string `json:"tag,omitempty"`               //
}
type ResponseItemCameraGetOrganizationCameraRolesAppliedOnNetworks struct {
	ID                string `json:"id,omitempty"`                //
	PermissionLevel   string `json:"permissionLevel,omitempty"`   //
	PermissionScope   string `json:"permissionScope,omitempty"`   //
	PermissionScopeID string `json:"permissionScopeId,omitempty"` //
	Tag               string `json:"tag,omitempty"`               //
}
type ResponseItemCameraGetOrganizationCameraRolesAppliedOrgWide struct {
	PermissionLevel   string `json:"permissionLevel,omitempty"`   //
	PermissionScope   string `json:"permissionScope,omitempty"`   //
	PermissionScopeID string `json:"permissionScopeId,omitempty"` //
	Tag               string `json:"tag,omitempty"`               //
}
type ResponseCameraCreateOrganizationCameraRole interface{}
type ResponseCameraGetOrganizationCameraRole struct {
	AppliedOnDevices  *[]ResponseCameraGetOrganizationCameraRoleAppliedOnDevices  `json:"appliedOnDevices,omitempty"`  //
	AppliedOnNetworks *[]ResponseCameraGetOrganizationCameraRoleAppliedOnNetworks `json:"appliedOnNetworks,omitempty"` //
	AppliedOrgWide    *[]ResponseCameraGetOrganizationCameraRoleAppliedOrgWide    `json:"appliedOrgWide,omitempty"`    //
	Name              string                                                      `json:"name,omitempty"`              //
	ID                string                                                      `json:"id,omitempty"`                //
}
type ResponseCameraGetOrganizationCameraRoleAppliedOnDevices struct {
	ID                string `json:"id,omitempty"`                //
	PermissionLevel   string `json:"permissionLevel,omitempty"`   //
	PermissionScope   string `json:"permissionScope,omitempty"`   //
	PermissionScopeID string `json:"permissionScopeId,omitempty"` //
	Tag               string `json:"tag,omitempty"`               //
}
type ResponseCameraGetOrganizationCameraRoleAppliedOnNetworks struct {
	ID                string `json:"id,omitempty"`                //
	PermissionLevel   string `json:"permissionLevel,omitempty"`   //
	PermissionScope   string `json:"permissionScope,omitempty"`   //
	PermissionScopeID string `json:"permissionScopeId,omitempty"` //
	Tag               string `json:"tag,omitempty"`               //
}
type ResponseCameraGetOrganizationCameraRoleAppliedOrgWide struct {
	PermissionLevel   string `json:"permissionLevel,omitempty"`   //
	PermissionScope   string `json:"permissionScope,omitempty"`   //
	PermissionScopeID string `json:"permissionScopeId,omitempty"` //
	Tag               string `json:"tag,omitempty"`               //
}
type ResponseCameraUpdateOrganizationCameraRole interface{}
type RequestCameraUpdateDeviceCameraCustomAnalytics struct {
	ArtifactID string                                                      `json:"artifactId,omitempty"` // The ID of the custom analytics artifact
	Enabled    *bool                                                       `json:"enabled,omitempty"`    // Enable custom analytics
	Parameters *[]RequestCameraUpdateDeviceCameraCustomAnalyticsParameters `json:"parameters,omitempty"` // Parameters for the custom analytics workload
}
type RequestCameraUpdateDeviceCameraCustomAnalyticsParameters struct {
	Name  string `json:"name,omitempty"`  // Name of the parameter
	Value string `json:"value,omitempty"` // Value of the parameter
}
type RequestCameraGenerateDeviceCameraSnapshot struct {
	Fullframe *bool  `json:"fullframe,omitempty"` // [optional] If set to "true" the snapshot will be taken at full sensor resolution. This will error if used with timestamp.
	Timestamp string `json:"timestamp,omitempty"` // [optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time.
}
type RequestCameraUpdateDeviceCameraQualityAndRetention struct {
	AudioRecordingEnabled          *bool  `json:"audioRecordingEnabled,omitempty"`          // Boolean indicating if audio recording is enabled(true) or disabled(false) on the camera
	MotionBasedRetentionEnabled    *bool  `json:"motionBasedRetentionEnabled,omitempty"`    // Boolean indicating if motion-based retention is enabled(true) or disabled(false) on the camera.
	MotionDetectorVersion          *int   `json:"motionDetectorVersion,omitempty"`          // The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	ProfileID                      string `json:"profileId,omitempty"`                      // The ID of a quality and retention profile to assign to the camera. The profile's settings will override all of the per-camera quality and retention settings. If the value of this parameter is null, any existing profile will be unassigned from the camera.
	Quality                        string `json:"quality,omitempty"`                        // Quality of the camera. Can be one of 'Standard', 'High', 'Enhanced' or 'Ultra'. Not all qualities are supported by every camera model.
	Resolution                     string `json:"resolution,omitempty"`                     // Resolution of the camera. Can be one of '1280x720', '1920x1080', '1080x1080', '2112x2112', '2880x2880', '2688x1512' or '3840x2160'.Not all resolutions are supported by every camera model.
	RestrictedBandwidthModeEnabled *bool  `json:"restrictedBandwidthModeEnabled,omitempty"` // Boolean indicating if restricted bandwidth is enabled(true) or disabled(false) on the camera. This setting does not apply to MV2 cameras.
}
type RequestCameraUpdateDeviceCameraSense struct {
	AudioDetection   *RequestCameraUpdateDeviceCameraSenseAudioDetection `json:"audioDetection,omitempty"`   // The details of the audio detection config.
	DetectionModelID string                                              `json:"detectionModelId,omitempty"` // The ID of the object detection model
	MqttBrokerID     string                                              `json:"mqttBrokerId,omitempty"`     // The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera
	SenseEnabled     *bool                                               `json:"senseEnabled,omitempty"`     // Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera
}
type RequestCameraUpdateDeviceCameraSenseAudioDetection struct {
	Enabled *bool `json:"enabled,omitempty"` // Boolean indicating if audio detection is enabled(true) or disabled(false) on the camera
}
type RequestCameraUpdateDeviceCameraVideoSettings struct {
	ExternalRtspEnabled *bool `json:"externalRtspEnabled,omitempty"` // Boolean indicating if external rtsp stream is exposed
}
type RequestCameraUpdateDeviceCameraWirelessProfiles struct {
	IDs *RequestCameraUpdateDeviceCameraWirelessProfilesIDs `json:"ids,omitempty"` // The ids of the wireless profile to assign to the given camera
}
type RequestCameraUpdateDeviceCameraWirelessProfilesIDs struct {
	Backup    string `json:"backup,omitempty"`    // The id of the backup wireless profile
	Primary   string `json:"primary,omitempty"`   // The id of the primary wireless profile
	Secondary string `json:"secondary,omitempty"` // The id of the secondary wireless profile
}
type RequestCameraCreateNetworkCameraQualityRetentionProfile struct {
	AudioRecordingEnabled          *bool                                                                  `json:"audioRecordingEnabled,omitempty"`          // Whether or not to record audio. Can be either true or false. Defaults to false.
	CloudArchiveEnabled            *bool                                                                  `json:"cloudArchiveEnabled,omitempty"`            // Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.
	MaxRetentionDays               *int                                                                   `json:"maxRetentionDays,omitempty"`               // The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be in the range of one to ninety days.
	MotionBasedRetentionEnabled    *bool                                                                  `json:"motionBasedRetentionEnabled,omitempty"`    // Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	MotionDetectorVersion          *int                                                                   `json:"motionDetectorVersion,omitempty"`          // The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	Name                           string                                                                 `json:"name,omitempty"`                           // The name of the new profile. Must be unique. This parameter is required.
	RestrictedBandwidthModeEnabled *bool                                                                  `json:"restrictedBandwidthModeEnabled,omitempty"` // Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	ScheduleID                     string                                                                 `json:"scheduleId,omitempty"`                     // Schedule for which this camera will record video, or 'null' to always record.
	SmartRetention                 *RequestCameraCreateNetworkCameraQualityRetentionProfileSmartRetention `json:"smartRetention,omitempty"`                 // Smart Retention records footage in two qualities and intelligently retains higher quality when motion, people or vehicles are detected.
	VideoSettings                  *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings  `json:"videoSettings,omitempty"`                  // Video quality and resolution settings for all the camera models.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileSmartRetention struct {
	Enabled *bool `json:"enabled,omitempty"` // Boolean indicating if Smart Retention is enabled(true) or disabled(false).
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings struct {
	MV12MV22MV72 *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"` // Quality and resolution for MV12/MV22/MV72 camera models.
	MV12WE       *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE       `json:"MV12WE,omitempty"`         // Quality and resolution for MV12WE camera models.
	MV13         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13         `json:"MV13,omitempty"`           // Quality and resolution for MV13 camera models.
	MV13M        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13M        `json:"MV13M,omitempty"`          // Quality and resolution for MV13M camera models.
	MV21MV71     *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71     `json:"MV21/MV71,omitempty"`      // Quality and resolution for MV21/MV71 camera models.
	MV22XMV72X   *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X   `json:"MV22X/MV72X,omitempty"`    // Quality and resolution for MV22X/MV72X camera models.
	MV23         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23         `json:"MV23,omitempty"`           // Quality and resolution for MV23 camera models.
	MV23M        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23M        `json:"MV23M,omitempty"`          // Quality and resolution for MV23M camera models.
	MV23X        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23X        `json:"MV23X,omitempty"`          // Quality and resolution for MV23X camera models.
	MV32         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32         `json:"MV32,omitempty"`           // Quality and resolution for MV32 camera models.
	MV33         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33         `json:"MV33,omitempty"`           // Quality and resolution for MV33 camera models.
	MV33M        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33M        `json:"MV33M,omitempty"`          // Quality and resolution for MV33M camera models.
	MV52         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52         `json:"MV52,omitempty"`           // Quality and resolution for MV52 camera models.
	MV53X        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV53X        `json:"MV53X,omitempty"`          // Quality and resolution for MV53X camera models.
	MV63         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63         `json:"MV63,omitempty"`           // Quality and resolution for MV63 camera models.
	MV63M        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63M        `json:"MV63M,omitempty"`          // Quality and resolution for MV63M camera models.
	MV63X        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X        `json:"MV63X,omitempty"`          // Quality and resolution for MV63X camera models.
	MV73         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73         `json:"MV73,omitempty"`           // Quality and resolution for MV73 camera models.
	MV73M        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73M        `json:"MV73M,omitempty"`          // Quality and resolution for MV73M camera models.
	MV73X        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73X        `json:"MV73X,omitempty"`          // Quality and resolution for MV73X camera models.
	MV84X        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV84X        `json:"MV84X,omitempty"`          // Quality and resolution for MV84X camera models.
	MV93         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93         `json:"MV93,omitempty"`           // Quality and resolution for MV93 camera models.
	MV93M        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93M        `json:"MV93M,omitempty"`          // Quality and resolution for MV93M camera models.
	MV93X        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X        `json:"MV93X,omitempty"`          // Quality and resolution for MV93X camera models.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080' or '2688x1512'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV53X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV84X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard' or 'Enhanced'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1440x1080' or '2560x1920'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfile struct {
	AudioRecordingEnabled          *bool                                                                  `json:"audioRecordingEnabled,omitempty"`          // Whether or not to record audio. Can be either true or false. Defaults to false.
	CloudArchiveEnabled            *bool                                                                  `json:"cloudArchiveEnabled,omitempty"`            // Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.
	MaxRetentionDays               *int                                                                   `json:"maxRetentionDays,omitempty"`               // The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be in the range of one to ninety days.
	MotionBasedRetentionEnabled    *bool                                                                  `json:"motionBasedRetentionEnabled,omitempty"`    // Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	MotionDetectorVersion          *int                                                                   `json:"motionDetectorVersion,omitempty"`          // The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	Name                           string                                                                 `json:"name,omitempty"`                           // The name of the new profile. Must be unique.
	RestrictedBandwidthModeEnabled *bool                                                                  `json:"restrictedBandwidthModeEnabled,omitempty"` // Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	ScheduleID                     string                                                                 `json:"scheduleId,omitempty"`                     // Schedule for which this camera will record video, or 'null' to always record.
	SmartRetention                 *RequestCameraUpdateNetworkCameraQualityRetentionProfileSmartRetention `json:"smartRetention,omitempty"`                 // Smart Retention records footage in two qualities and intelligently retains higher quality when motion, people or vehicles are detected.
	VideoSettings                  *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings  `json:"videoSettings,omitempty"`                  // Video quality and resolution settings for all the camera models.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileSmartRetention struct {
	Enabled *bool `json:"enabled,omitempty"` // Boolean indicating if Smart Retention is enabled(true) or disabled(false).
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings struct {
	MV12MV22MV72 *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"` // Quality and resolution for MV12/MV22/MV72 camera models.
	MV12WE       *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE       `json:"MV12WE,omitempty"`         // Quality and resolution for MV12WE camera models.
	MV13         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13         `json:"MV13,omitempty"`           // Quality and resolution for MV13 camera models.
	MV13M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13M        `json:"MV13M,omitempty"`          // Quality and resolution for MV13M camera models.
	MV21MV71     *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71     `json:"MV21/MV71,omitempty"`      // Quality and resolution for MV21/MV71 camera models.
	MV22XMV72X   *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X   `json:"MV22X/MV72X,omitempty"`    // Quality and resolution for MV22X/MV72X camera models.
	MV23         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23         `json:"MV23,omitempty"`           // Quality and resolution for MV23 camera models.
	MV23M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23M        `json:"MV23M,omitempty"`          // Quality and resolution for MV23M camera models.
	MV23X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23X        `json:"MV23X,omitempty"`          // Quality and resolution for MV23X camera models.
	MV32         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32         `json:"MV32,omitempty"`           // Quality and resolution for MV32 camera models.
	MV33         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33         `json:"MV33,omitempty"`           // Quality and resolution for MV33 camera models.
	MV33M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33M        `json:"MV33M,omitempty"`          // Quality and resolution for MV33M camera models.
	MV52         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52         `json:"MV52,omitempty"`           // Quality and resolution for MV52 camera models.
	MV53X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53X        `json:"MV53X,omitempty"`          // Quality and resolution for MV53X camera models.
	MV63         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63         `json:"MV63,omitempty"`           // Quality and resolution for MV63 camera models.
	MV63M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63M        `json:"MV63M,omitempty"`          // Quality and resolution for MV63M camera models.
	MV63X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X        `json:"MV63X,omitempty"`          // Quality and resolution for MV63X camera models.
	MV73         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73         `json:"MV73,omitempty"`           // Quality and resolution for MV73 camera models.
	MV73M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73M        `json:"MV73M,omitempty"`          // Quality and resolution for MV73M camera models.
	MV73X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73X        `json:"MV73X,omitempty"`          // Quality and resolution for MV73X camera models.
	MV84X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84X        `json:"MV84X,omitempty"`          // Quality and resolution for MV84X camera models.
	MV93         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93         `json:"MV93,omitempty"`           // Quality and resolution for MV93 camera models.
	MV93M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93M        `json:"MV93M,omitempty"`          // Quality and resolution for MV93M camera models.
	MV93X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X        `json:"MV93X,omitempty"`          // Quality and resolution for MV93X camera models.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080' or '2688x1512'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard' or 'Enhanced'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1440x1080' or '2560x1920'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93M struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraCreateNetworkCameraWirelessProfile struct {
	IDentity *RequestCameraCreateNetworkCameraWirelessProfileIDentity `json:"identity,omitempty"` // The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
	Name     string                                                   `json:"name,omitempty"`     // The name of the camera wireless profile. This parameter is required.
	SSID     *RequestCameraCreateNetworkCameraWirelessProfileSSID     `json:"ssid,omitempty"`     // The details of the SSID config.
}
type RequestCameraCreateNetworkCameraWirelessProfileIDentity struct {
	Password string `json:"password,omitempty"` // The password of the identity.
	Username string `json:"username,omitempty"` // The username of the identity.
}
type RequestCameraCreateNetworkCameraWirelessProfileSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       // The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	EncryptionMode string `json:"encryptionMode,omitempty"` // The encryption mode of the SSID. It can be set to ('wpa', 'wpa-eap'). With 'wpa' mode, the authMode should be 'psk' and with 'wpa-eap' the authMode should be '8021x-radius'
	Name           string `json:"name,omitempty"`           // The name of the SSID.
	Psk            string `json:"psk,omitempty"`            // The pre-shared key of the SSID.
}
type RequestCameraUpdateNetworkCameraWirelessProfile struct {
	IDentity *RequestCameraUpdateNetworkCameraWirelessProfileIDentity `json:"identity,omitempty"` // The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
	Name     string                                                   `json:"name,omitempty"`     // The name of the camera wireless profile.
	SSID     *RequestCameraUpdateNetworkCameraWirelessProfileSSID     `json:"ssid,omitempty"`     // The details of the SSID config.
}
type RequestCameraUpdateNetworkCameraWirelessProfileIDentity struct {
	Password string `json:"password,omitempty"` // The password of the identity.
	Username string `json:"username,omitempty"` // The username of the identity.
}
type RequestCameraUpdateNetworkCameraWirelessProfileSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       // The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	EncryptionMode string `json:"encryptionMode,omitempty"` // The encryption mode of the SSID. It can be set to ('wpa', 'wpa-eap'). With 'wpa' mode, the authMode should be 'psk' and with 'wpa-eap' the authMode should be '8021x-radius'
	Name           string `json:"name,omitempty"`           // The name of the SSID.
	Psk            string `json:"psk,omitempty"`            // The pre-shared key of the SSID.
}
type RequestCameraCreateOrganizationCameraCustomAnalyticsArtifact struct {
	Name string `json:"name,omitempty"` // Unique name of the artifact
}
type RequestCameraUpdateOrganizationCameraOnboardingStatuses struct {
	Serial                  string `json:"serial,omitempty"`                  // Serial of camera
	WirelessCredentialsSent *bool  `json:"wirelessCredentialsSent,omitempty"` // Note whether credentials were sent successfully
}
type RequestCameraCreateOrganizationCameraRole struct {
	AppliedOnDevices  *[]RequestCameraCreateOrganizationCameraRoleAppliedOnDevices  `json:"appliedOnDevices,omitempty"`  // Device tag on which this specified permission is applied.
	AppliedOnNetworks *[]RequestCameraCreateOrganizationCameraRoleAppliedOnNetworks `json:"appliedOnNetworks,omitempty"` // Network tag on which this specified permission is applied.
	AppliedOrgWide    *[]RequestCameraCreateOrganizationCameraRoleAppliedOrgWide    `json:"appliedOrgWide,omitempty"`    // Permissions to be applied org wide.
	Name              string                                                        `json:"name,omitempty"`              // The name of the new role. Must be unique. This parameter is required.
}
type RequestCameraCreateOrganizationCameraRoleAppliedOnDevices struct {
	ID                string `json:"id,omitempty"`                // Device id.
	InNetworksWithID  string `json:"inNetworksWithId,omitempty"`  // Network id scope
	InNetworksWithTag string `json:"inNetworksWithTag,omitempty"` // Network tag scope
	PermissionScopeID string `json:"permissionScopeId,omitempty"` // Permission scope id
	Tag               string `json:"tag,omitempty"`               // Device tag.
}
type RequestCameraCreateOrganizationCameraRoleAppliedOnNetworks struct {
	ID                string `json:"id,omitempty"`                // Network id.
	PermissionScopeID string `json:"permissionScopeId,omitempty"` // Permission scope id
	Tag               string `json:"tag,omitempty"`               // Network tag
}
type RequestCameraCreateOrganizationCameraRoleAppliedOrgWide struct {
	PermissionScopeID string `json:"permissionScopeId,omitempty"` // Permission scope id
}
type RequestCameraUpdateOrganizationCameraRole struct {
	AppliedOnDevices  *[]RequestCameraUpdateOrganizationCameraRoleAppliedOnDevices  `json:"appliedOnDevices,omitempty"`  // Device tag on which this specified permission is applied.
	AppliedOnNetworks *[]RequestCameraUpdateOrganizationCameraRoleAppliedOnNetworks `json:"appliedOnNetworks,omitempty"` // Network tag on which this specified permission is applied.
	AppliedOrgWide    *[]RequestCameraUpdateOrganizationCameraRoleAppliedOrgWide    `json:"appliedOrgWide,omitempty"`    // Permissions to be applied org wide.
	Name              string                                                        `json:"name,omitempty"`              // The name of the new role. Must be unique.
}
type RequestCameraUpdateOrganizationCameraRoleAppliedOnDevices struct {
	ID                string `json:"id,omitempty"`                // Device id.
	InNetworksWithID  string `json:"inNetworksWithId,omitempty"`  // Network id scope
	InNetworksWithTag string `json:"inNetworksWithTag,omitempty"` // Network tag scope
	PermissionScopeID string `json:"permissionScopeId,omitempty"` // Permission scope id
	Tag               string `json:"tag,omitempty"`               // Device tag.
}
type RequestCameraUpdateOrganizationCameraRoleAppliedOnNetworks struct {
	ID                string `json:"id,omitempty"`                // Network id
	PermissionScopeID string `json:"permissionScopeId,omitempty"` // Permission scope id
	Tag               string `json:"tag,omitempty"`               // Network tag
}
type RequestCameraUpdateOrganizationCameraRoleAppliedOrgWide struct {
	PermissionScopeID string `json:"permissionScopeId,omitempty"` // Permission scope id
}

//GetDeviceCameraAnalyticsLive Returns live state from camera analytics zones
/* Returns live state from camera analytics zones

@param serial serial path parameter.


*/
// {'deprecated': True, 'description': 'Returns live state from camera analytics zones', 'operationId': 'getDeviceCameraAnalyticsLive', 'parameters': [{'description': 'Serial', 'in': 'path', 'name': 'serial', 'required': True, 'schema': {'type': 'string'}}], 'responses': [{'code': '200', 'content_type': 'application/json', 'schema': None, 'data': {'properties': {'ts': {'description': 'The current time', 'type': 'string'}, 'zones': {'description': 'The zones state', 'properties': {'zoneId': {'description': 'The zone state, dynamic', 'properties': {'person': {'description': 'The count per type, dynamic', 'type': 'integer'}}, 'type': 'object'}}, 'type': 'object'}}, 'type': 'object'}}], 'summary': 'Returns live state from camera analytics zones', 'tags': ['camera', 'monitor', 'analytics', 'live'], 'content_types': ['application/json'], 'response_type': {}, 'method': 'GET', 'request_types': [], 'path': '/api/v1/devices/{serial}/camera/analytics/live', 'originalURL': '/api/v1/devices/{serial}/camera/analytics/live', 'headers': {}, 'path_params': {'serial': {'type': 'string', 'description': 'serial path parameter.', 'required': True}}, 'params': {}, 'response_json_schema': {'properties': {'ts': {'description': 'The current time', 'type': 'string'}, 'zones': {'description': 'The zones state', 'properties': {'zoneId': {'description': 'The zone state, dynamic', 'properties': {'person': {'description': 'The count per type, dynamic', 'type': 'integer'}}, 'type': 'object'}}, 'type': 'object'}}, 'type': 'object', '$schema': 'http://json-schema.org/draft-04/schema#'}, 'response': {'ts': 'string', 'zones': {'zoneId': {'person': 0}}}, 'data': {}, 'id': 'aca5cc3c-e030-585d-8fe0-b6f9d6b87a42', 'alt_name': 'getDeviceCameraAnalyticsLive', 'name': 'getDeviceCameraAnalyticsLive', 'has_rename': False, 'kwargs': ''}
func (s *CameraService) GetDeviceCameraAnalyticsLive(serial string) (*ResponseCameraGetDeviceCameraAnalyticsLive, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/live"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraAnalyticsLive](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraAnalyticsOverview Returns an overview of aggregate analytics data for a timespan
/* Returns an overview of aggregate analytics data for a timespan

@param serial serial path parameter.
@param getDeviceCameraAnalyticsOverviewQueryParams Filtering parameter


*/
// {'deprecated': True, 'description': 'Returns an overview of aggregate analytics data for a timespan', 'operationId': 'getDeviceCameraAnalyticsOverview', 'parameters': [{'description': 'Serial', 'in': 'path', 'name': 'serial', 'required': True, 'schema': {'type': 'string'}}, {'description': 'The beginning of the timespan for the data. The maximum lookback period is 365 days from today.', 'in': 'query', 'name': 't0', 'schema': {'type': 'string'}}, {'description': 'The end of the timespan for the data. t1 can be a maximum of 7 days after t0.', 'in': 'query', 'name': 't1', 'schema': {'type': 'string'}}, {'description': 'The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour.', 'in': 'query', 'name': 'timespan', 'schema': {'format': 'float', 'maximum': 604800, 'type': 'number'}}, {'description': '[optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].', 'in': 'query', 'name': 'objectType', 'schema': {'enum': ['person', 'vehicle'], 'type': 'string'}}], 'responses': [{'code': '200', 'content_type': 'application/json', 'schema': None, 'data': {'items': {'properties': {'averageCount': {'description': 'The average count', 'format': 'float', 'type': 'number'}, 'endTs': {'description': 'The end time', 'format': 'date-time', 'type': 'string'}, 'entrances': {'description': 'The number of sentrances', 'type': 'integer'}, 'startTs': {'description': 'The start time', 'format': 'date-time', 'type': 'string'}, 'zoneId': {'description': 'The zone id', 'type': 'integer'}}, 'type': 'object'}, 'type': 'array'}}], 'summary': 'Returns an overview of aggregate analytics data for a timespan', 'tags': ['camera', 'monitor', 'analytics', 'overview'], 'content_types': ['application/json'], 'response_type': [], 'method': 'GET', 'request_types': [], 'path': '/api/v1/devices/{serial}/camera/analytics/overview', 'originalURL': '/api/v1/devices/{serial}/camera/analytics/overview', 'headers': {}, 'path_params': {'serial': {'type': 'string', 'description': 'serial path parameter.', 'required': True}}, 'params': {'t0': {'type': 'string', 'description': 't0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 365 days from today.', 'required': False}, 't1': {'type': 'string', 'description': 't1 query parameter. The end of the timespan for the data. t1 can be a maximum of 7 days after t0.', 'required': False}, 'timespan': {'type': 'number', 'description': 'timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour.', 'required': False}, 'objectType': {'type': 'string', 'description': 'objectType query parameter. [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].', 'required': False}}, 'response_json_schema': {'items': {'properties': {'averageCount': {'description': 'The average count', 'format': 'float', 'type': 'number'}, 'endTs': {'description': 'The end time', 'format': 'date-time', 'type': 'string'}, 'entrances': {'description': 'The number of sentrances', 'type': 'integer'}, 'startTs': {'description': 'The start time', 'format': 'date-time', 'type': 'string'}, 'zoneId': {'description': 'The zone id', 'type': 'integer'}}, 'type': 'object'}, 'type': 'array', '$schema': 'http://json-schema.org/draft-04/schema#'}, 'response': [{'averageCount': 0, 'endTs': 'string', 'entrances': 0, 'startTs': 'string', 'zoneId': 0}], 'data': {}, 'id': '6f282ea8-e896-5ca2-af45-6ed584b34d29', 'alt_name': 'getDeviceCameraAnalyticsOverview', 'name': 'getDeviceCameraAnalyticsOverview', 'has_rename': False, 'kwargs': ''}
func (s *CameraService) GetDeviceCameraAnalyticsOverview(serial string, getDeviceCameraAnalyticsOverviewQueryParams *GetDeviceCameraAnalyticsOverviewQueryParams) (*ResponseCameraGetDeviceCameraAnalyticsOverview, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/overview"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraAnalyticsOverview](
		func() (*resty.Response, error) {
			return GET(path, s.client, getDeviceCameraAnalyticsOverviewQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraAnalyticsRecent Returns most recent record for analytics zones
/* Returns most recent record for analytics zones

@param serial serial path parameter.
@param getDeviceCameraAnalyticsRecentQueryParams Filtering parameter


*/
// {'deprecated': True, 'description': 'Returns most recent record for analytics zones', 'operationId': 'getDeviceCameraAnalyticsRecent', 'parameters': [{'description': 'Serial', 'in': 'path', 'name': 'serial', 'required': True, 'schema': {'type': 'string'}}, {'description': '[optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].', 'in': 'query', 'name': 'objectType', 'schema': {'enum': ['person', 'vehicle'], 'type': 'string'}}], 'responses': [{'code': '200', 'content_type': 'application/json', 'schema': None, 'data': {'items': {'properties': {'averageCount': {'description': 'The average count', 'format': 'float', 'type': 'number'}, 'endTs': {'description': 'The end time', 'format': 'date-time', 'type': 'string'}, 'entrances': {'description': 'The number of entrances', 'type': 'integer'}, 'startTs': {'description': 'The start time', 'format': 'date-time', 'type': 'string'}, 'zoneId': {'description': 'The zone id', 'type': 'integer'}}, 'type': 'object'}, 'type': 'array'}}], 'summary': 'Returns most recent record for analytics zones', 'tags': ['camera', 'monitor', 'analytics', 'recent'], 'content_types': ['application/json'], 'response_type': [], 'method': 'GET', 'request_types': [], 'path': '/api/v1/devices/{serial}/camera/analytics/recent', 'originalURL': '/api/v1/devices/{serial}/camera/analytics/recent', 'headers': {}, 'path_params': {'serial': {'type': 'string', 'description': 'serial path parameter.', 'required': True}}, 'params': {'objectType': {'type': 'string', 'description': 'objectType query parameter. [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].', 'required': False}}, 'response_json_schema': {'items': {'properties': {'averageCount': {'description': 'The average count', 'format': 'float', 'type': 'number'}, 'endTs': {'description': 'The end time', 'format': 'date-time', 'type': 'string'}, 'entrances': {'description': 'The number of entrances', 'type': 'integer'}, 'startTs': {'description': 'The start time', 'format': 'date-time', 'type': 'string'}, 'zoneId': {'description': 'The zone id', 'type': 'integer'}}, 'type': 'object'}, 'type': 'array', '$schema': 'http://json-schema.org/draft-04/schema#'}, 'response': [{'averageCount': 0, 'endTs': 'string', 'entrances': 0, 'startTs': 'string', 'zoneId': 0}], 'data': {}, 'id': 'c4784d56-951c-5974-9bda-395c6aaad916', 'alt_name': 'getDeviceCameraAnalyticsRecent', 'name': 'getDeviceCameraAnalyticsRecent', 'has_rename': False, 'kwargs': ''}
func (s *CameraService) GetDeviceCameraAnalyticsRecent(serial string, getDeviceCameraAnalyticsRecentQueryParams *GetDeviceCameraAnalyticsRecentQueryParams) (*ResponseCameraGetDeviceCameraAnalyticsRecent, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/recent"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraAnalyticsRecent](
		func() (*resty.Response, error) {
			return GET(path, s.client, getDeviceCameraAnalyticsRecentQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraAnalyticsZones Returns all configured analytic zones for this camera
/* Returns all configured analytic zones for this camera

@param serial serial path parameter.


*/
// {'deprecated': True, 'description': 'Returns all configured analytic zones for this camera', 'operationId': 'getDeviceCameraAnalyticsZones', 'parameters': [{'description': 'Serial', 'in': 'path', 'name': 'serial', 'required': True, 'schema': {'type': 'string'}}], 'responses': [{'code': '200', 'content_type': 'application/json', 'schema': None, 'data': {'items': {'properties': {'id': {'description': 'The zone ID', 'type': 'string'}, 'label': {'description': 'The zone label', 'type': 'string'}, 'regionOfInterest': {'description': 'The region of interest', 'properties': {'x0': {'description': 'The x0 coordinate', 'type': 'string'}, 'x1': {'description': 'The x1 coordinate', 'type': 'string'}, 'y0': {'description': 'The y0 coordinate', 'type': 'string'}, 'y1': {'description': 'The y1 coordinate', 'type': 'string'}}, 'type': 'object'}, 'type': {'description': 'The zone type', 'type': 'string'}}, 'type': 'object'}, 'type': 'array'}}], 'summary': 'Returns all configured analytic zones for this camera', 'tags': ['camera', 'monitor', 'analytics', 'zones'], 'content_types': ['application/json'], 'response_type': [], 'method': 'GET', 'request_types': [], 'path': '/api/v1/devices/{serial}/camera/analytics/zones', 'originalURL': '/api/v1/devices/{serial}/camera/analytics/zones', 'headers': {}, 'path_params': {'serial': {'type': 'string', 'description': 'serial path parameter.', 'required': True}}, 'params': {}, 'response_json_schema': {'items': {'properties': {'id': {'description': 'The zone ID', 'type': 'string'}, 'label': {'description': 'The zone label', 'type': 'string'}, 'regionOfInterest': {'description': 'The region of interest', 'properties': {'x0': {'description': 'The x0 coordinate', 'type': 'string'}, 'x1': {'description': 'The x1 coordinate', 'type': 'string'}, 'y0': {'description': 'The y0 coordinate', 'type': 'string'}, 'y1': {'description': 'The y1 coordinate', 'type': 'string'}}, 'type': 'object'}, 'type': {'description': 'The zone type', 'type': 'string'}}, 'type': 'object'}, 'type': 'array', '$schema': 'http://json-schema.org/draft-04/schema#'}, 'response': [{'id': 'string', 'label': 'string', 'regionOfInterest': {'x0': 'string', 'x1': 'string', 'y0': 'string', 'y1': 'string'}, 'type': 'string'}], 'data': {}, 'id': '63cb69f1-baa9-5a2d-b847-a75ded2dfb00', 'alt_name': 'getDeviceCameraAnalyticsZones', 'name': 'getDeviceCameraAnalyticsZones', 'has_rename': False, 'kwargs': ''}
func (s *CameraService) GetDeviceCameraAnalyticsZones(serial string) (*ResponseCameraGetDeviceCameraAnalyticsZones, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/zones"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraAnalyticsZones](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraAnalyticsZoneHistory Return historical records for analytic zones
/* Return historical records for analytic zones

@param serial serial path parameter.
@param zoneID zoneId path parameter. Zone ID
@param getDeviceCameraAnalyticsZoneHistoryQueryParams Filtering parameter


*/
// {'deprecated': True, 'description': 'Return historical records for analytic zones', 'operationId': 'getDeviceCameraAnalyticsZoneHistory', 'parameters': [{'description': 'Serial', 'in': 'path', 'name': 'serial', 'required': True, 'schema': {'type': 'string'}}, {'description': 'Zone ID', 'in': 'path', 'name': 'zoneId', 'required': True, 'schema': {'type': 'string'}}, {'description': 'The beginning of the timespan for the data. The maximum lookback period is 365 days from today.', 'in': 'query', 'name': 't0', 'schema': {'type': 'string'}}, {'description': 'The end of the timespan for the data. t1 can be a maximum of 14 hours after t0.', 'in': 'query', 'name': 't1', 'schema': {'type': 'string'}}, {'description': 'The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour.', 'in': 'query', 'name': 'timespan', 'schema': {'format': 'float', 'maximum': 50400, 'type': 'number'}}, {'description': 'The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60.', 'in': 'query', 'name': 'resolution', 'schema': {'type': 'integer'}}, {'description': '[optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].', 'in': 'query', 'name': 'objectType', 'schema': {'enum': ['person', 'vehicle'], 'type': 'string'}}], 'responses': [{'code': '200', 'content_type': 'application/json', 'schema': None, 'data': {'items': {'properties': {'averageCount': {'description': 'The average count', 'format': 'float', 'type': 'number'}, 'endTs': {'description': 'The end time', 'format': 'date-time', 'type': 'string'}, 'entrances': {'description': 'The number of entrances', 'type': 'integer'}, 'startTs': {'description': 'The start time', 'format': 'date-time', 'type': 'string'}}, 'type': 'object'}, 'type': 'array'}}], 'summary': 'Return historical records for analytic zones', 'tags': ['camera', 'monitor', 'analytics', 'zones', 'history'], 'content_types': ['application/json'], 'response_type': [], 'method': 'GET', 'request_types': [], 'path': '/api/v1/devices/{serial}/camera/analytics/zones/{zoneId}/history', 'originalURL': '/api/v1/devices/{serial}/camera/analytics/zones/{zoneId}/history', 'headers': {}, 'path_params': {'serial': {'type': 'string', 'description': 'serial path parameter.', 'required': True}, 'zoneId': {'type': 'string', 'description': 'zoneId path parameter. Zone ID', 'required': True}}, 'params': {'t0': {'type': 'string', 'description': 't0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 365 days from today.', 'required': False}, 't1': {'type': 'string', 'description': 't1 query parameter. The end of the timespan for the data. t1 can be a maximum of 14 hours after t0.', 'required': False}, 'timespan': {'type': 'number', 'description': 'timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour.', 'required': False}, 'resolution': {'type': 'integer', 'description': 'resolution query parameter. The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60.', 'required': False}, 'objectType': {'type': 'string', 'description': 'objectType query parameter. [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].', 'required': False}}, 'response_json_schema': {'items': {'properties': {'averageCount': {'description': 'The average count', 'format': 'float', 'type': 'number'}, 'endTs': {'description': 'The end time', 'format': 'date-time', 'type': 'string'}, 'entrances': {'description': 'The number of entrances', 'type': 'integer'}, 'startTs': {'description': 'The start time', 'format': 'date-time', 'type': 'string'}}, 'type': 'object'}, 'type': 'array', '$schema': 'http://json-schema.org/draft-04/schema#'}, 'response': [{'averageCount': 0, 'endTs': 'string', 'entrances': 0, 'startTs': 'string'}], 'data': {}, 'id': 'b787212f-cf12-54cb-ac69-c0e157ab7a9d', 'alt_name': 'getDeviceCameraAnalyticsZoneHistory', 'name': 'getDeviceCameraAnalyticsZoneHistory', 'has_rename': False, 'kwargs': ''}
func (s *CameraService) GetDeviceCameraAnalyticsZoneHistory(serial string, zoneID string, getDeviceCameraAnalyticsZoneHistoryQueryParams *GetDeviceCameraAnalyticsZoneHistoryQueryParams) (*ResponseCameraGetDeviceCameraAnalyticsZoneHistory, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/zones/{zoneId}/history"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{zoneId}", fmt.Sprintf("%v", zoneID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraAnalyticsZoneHistory](
		func() (*resty.Response, error) {
			return GET(path, s.client, getDeviceCameraAnalyticsZoneHistoryQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraCustomAnalytics Return custom analytics settings for a camera
/* Return custom analytics settings for a camera

@param serial serial path parameter.


*/

func (s *CameraService) GetDeviceCameraCustomAnalytics(serial string) (*ResponseCameraGetDeviceCameraCustomAnalytics, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/customAnalytics"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraCustomAnalytics](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraQualityAndRetention Returns quality and retention settings for the given camera
/* Returns quality and retention settings for the given camera

@param serial serial path parameter.


*/

func (s *CameraService) GetDeviceCameraQualityAndRetention(serial string) (*ResponseCameraGetDeviceCameraQualityAndRetention, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/qualityAndRetention"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraQualityAndRetention](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraSense Returns sense settings for a given camera
/* Returns sense settings for a given camera

@param serial serial path parameter.


*/

func (s *CameraService) GetDeviceCameraSense(serial string) (*ResponseCameraGetDeviceCameraSense, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/sense"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraSense](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraSenseObjectDetectionModels Returns the MV Sense object detection model list for the given camera
/* Returns the MV Sense object detection model list for the given camera

@param serial serial path parameter.


*/

func (s *CameraService) GetDeviceCameraSenseObjectDetectionModels(serial string) (*ResponseCameraGetDeviceCameraSenseObjectDetectionModels, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/sense/objectDetectionModels"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraSenseObjectDetectionModels](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraVideoSettings Returns video settings for the given camera
/* Returns video settings for the given camera

@param serial serial path parameter.


*/

func (s *CameraService) GetDeviceCameraVideoSettings(serial string) (*ResponseCameraGetDeviceCameraVideoSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/video/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraVideoSettings](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraVideoLink Returns video link to the specified camera
/* Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp.

@param serial serial path parameter.
@param getDeviceCameraVideoLinkQueryParams Filtering parameter


*/

func (s *CameraService) GetDeviceCameraVideoLink(serial string, getDeviceCameraVideoLinkQueryParams *GetDeviceCameraVideoLinkQueryParams) (*ResponseCameraGetDeviceCameraVideoLink, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/videoLink"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraVideoLink](
		func() (*resty.Response, error) {
			return GET(path, s.client, getDeviceCameraVideoLinkQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetDeviceCameraWirelessProfiles Returns wireless profile assigned to the given camera
/* Returns wireless profile assigned to the given camera

@param serial serial path parameter.


*/

func (s *CameraService) GetDeviceCameraWirelessProfiles(serial string) (*ResponseCameraGetDeviceCameraWirelessProfiles, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/wirelessProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetDeviceCameraWirelessProfiles](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkCameraQualityRetentionProfiles List the quality retention profiles for this network
/* List the quality retention profiles for this network

@param networkID networkId path parameter. Network ID


*/

func (s *CameraService) GetNetworkCameraQualityRetentionProfiles(networkID string) (*ResponseCameraGetNetworkCameraQualityRetentionProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetNetworkCameraQualityRetentionProfiles](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkCameraQualityRetentionProfile Retrieve a single quality retention profile
/* Retrieve a single quality retention profile

@param networkID networkId path parameter. Network ID
@param qualityRetentionProfileID qualityRetentionProfileId path parameter. Quality retention profile ID


*/

func (s *CameraService) GetNetworkCameraQualityRetentionProfile(networkID string, qualityRetentionProfileID string) (*ResponseCameraGetNetworkCameraQualityRetentionProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qualityRetentionProfileId}", fmt.Sprintf("%v", qualityRetentionProfileID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetNetworkCameraQualityRetentionProfile](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkCameraSchedules Returns a list of all camera recording schedules.
/* Returns a list of all camera recording schedules.

@param networkID networkId path parameter. Network ID


*/

func (s *CameraService) GetNetworkCameraSchedules(networkID string) (*ResponseCameraGetNetworkCameraSchedules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/schedules"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetNetworkCameraSchedules](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkCameraWirelessProfiles List the camera wireless profiles for this network.
/* List the camera wireless profiles for this network.

@param networkID networkId path parameter. Network ID


*/

func (s *CameraService) GetNetworkCameraWirelessProfiles(networkID string) (*ResponseCameraGetNetworkCameraWirelessProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetNetworkCameraWirelessProfiles](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetNetworkCameraWirelessProfile Retrieve a single camera wireless profile.
/* Retrieve a single camera wireless profile.

@param networkID networkId path parameter. Network ID
@param wirelessProfileID wirelessProfileId path parameter. Wireless profile ID


*/

func (s *CameraService) GetNetworkCameraWirelessProfile(networkID string, wirelessProfileID string) (*ResponseCameraGetNetworkCameraWirelessProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{wirelessProfileId}", fmt.Sprintf("%v", wirelessProfileID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetNetworkCameraWirelessProfile](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraBoundariesAreasByDevice Returns all configured area boundaries of cameras
/* Returns all configured area boundaries of cameras

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCameraBoundariesAreasByDeviceQueryParams Filtering parameter


*/

func (s *CameraService) GetOrganizationCameraBoundariesAreasByDevice(organizationID string, getOrganizationCameraBoundariesAreasByDeviceQueryParams *GetOrganizationCameraBoundariesAreasByDeviceQueryParams) (*ResponseCameraGetOrganizationCameraBoundariesAreasByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/boundaries/areas/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraBoundariesAreasByDevice](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationCameraBoundariesAreasByDeviceQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraBoundariesLinesByDevice Returns all configured crossingline boundaries of cameras
/* Returns all configured crossingline boundaries of cameras

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCameraBoundariesLinesByDeviceQueryParams Filtering parameter


*/

func (s *CameraService) GetOrganizationCameraBoundariesLinesByDevice(organizationID string, getOrganizationCameraBoundariesLinesByDeviceQueryParams *GetOrganizationCameraBoundariesLinesByDeviceQueryParams) (*ResponseCameraGetOrganizationCameraBoundariesLinesByDevice, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/boundaries/lines/byDevice"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraBoundariesLinesByDevice](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationCameraBoundariesLinesByDeviceQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraCustomAnalyticsArtifacts List Custom Analytics Artifacts
/* List Custom Analytics Artifacts

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CameraService) GetOrganizationCameraCustomAnalyticsArtifacts(organizationID string) (*ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraCustomAnalyticsArtifact Get Custom Analytics Artifact
/* Get Custom Analytics Artifact

@param organizationID organizationId path parameter. Organization ID
@param artifactID artifactId path parameter. Artifact ID


*/

func (s *CameraService) GetOrganizationCameraCustomAnalyticsArtifact(organizationID string, artifactID string) (*ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{artifactId}", fmt.Sprintf("%v", artifactID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraDetectionsHistoryByBoundaryByInterval Returns analytics data for timespans
/* Returns analytics data for timespans

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCameraDetectionsHistoryByBoundaryByIntervalQueryParams Filtering parameter


*/

func (s *CameraService) GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(organizationID string, getOrganizationCameraDetectionsHistoryByBoundaryByIntervalQueryParams *GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalQueryParams) (*ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/detections/history/byBoundary/byInterval"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationCameraDetectionsHistoryByBoundaryByIntervalQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraOnboardingStatuses Fetch onboarding status of cameras
/* Fetch onboarding status of cameras

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCameraOnboardingStatusesQueryParams Filtering parameter


*/

func (s *CameraService) GetOrganizationCameraOnboardingStatuses(organizationID string, getOrganizationCameraOnboardingStatusesQueryParams *GetOrganizationCameraOnboardingStatusesQueryParams) (*ResponseCameraGetOrganizationCameraOnboardingStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/onboarding/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraOnboardingStatuses](
		func() (*resty.Response, error) {
			return GET(path, s.client, getOrganizationCameraOnboardingStatusesQueryParams, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraPermissions List the permissions scopes for this organization
/* List the permissions scopes for this organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CameraService) GetOrganizationCameraPermissions(organizationID string) (*ResponseCameraGetOrganizationCameraPermissions, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/permissions"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraPermissions](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraPermission Retrieve a single permission scope
/* Retrieve a single permission scope

@param organizationID organizationId path parameter. Organization ID
@param permissionScopeID permissionScopeId path parameter. Permission scope ID


*/

func (s *CameraService) GetOrganizationCameraPermission(organizationID string, permissionScopeID string) (*ResponseCameraGetOrganizationCameraPermission, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/permissions/{permissionScopeId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{permissionScopeId}", fmt.Sprintf("%v", permissionScopeID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraPermission](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraRoles List all the roles in this organization
/* List all the roles in this organization

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CameraService) GetOrganizationCameraRoles(organizationID string) (*ResponseCameraGetOrganizationCameraRoles, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/roles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraRoles](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GetOrganizationCameraRole Retrieve a single role.
/* Retrieve a single role.

@param organizationID organizationId path parameter. Organization ID
@param roleID roleId path parameter. Role ID


*/

func (s *CameraService) GetOrganizationCameraRole(organizationID string, roleID string) (*ResponseCameraGetOrganizationCameraRole, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/roles/{roleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraGetOrganizationCameraRole](
		func() (*resty.Response, error) {
			return GET(path, s.client, &QueryParamsDefault, &HeaderDefault)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//GenerateDeviceCameraSnapshot Generate a snapshot of what the camera sees at the specified time and return a link to that image.
/* Generate a snapshot of what the camera sees at the specified time and return a link to that image.

@param serial serial path parameter.


*/

func (s *CameraService) GenerateDeviceCameraSnapshot(serial string, requestCameraGenerateDeviceCameraSnapshot *RequestCameraGenerateDeviceCameraSnapshot) (*ResponseCameraGenerateDeviceCameraSnapshot, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/generateSnapshot"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Past way

	return doWithRetriesAndResult[ResponseCameraGenerateDeviceCameraSnapshot](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestCameraGenerateDeviceCameraSnapshot, nil)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//CreateNetworkCameraQualityRetentionProfile Creates new quality retention profile for this network.
/* Creates new quality retention profile for this network.

@param networkID networkId path parameter. Network ID


*/

func (s *CameraService) CreateNetworkCameraQualityRetentionProfile(networkID string, requestCameraCreateNetworkCameraQualityRetentionProfile *RequestCameraCreateNetworkCameraQualityRetentionProfile) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return POST(path, s.client, requestCameraCreateNetworkCameraQualityRetentionProfile, nil)
		},
		s.backoff,
	)

}

//CreateNetworkCameraWirelessProfile Creates a new camera wireless profile for this network.
/* Creates a new camera wireless profile for this network.

@param networkID networkId path parameter. Network ID


*/

func (s *CameraService) CreateNetworkCameraWirelessProfile(networkID string, requestCameraCreateNetworkCameraWirelessProfile *RequestCameraCreateNetworkCameraWirelessProfile) (*ResponseCameraCreateNetworkCameraWirelessProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseCameraCreateNetworkCameraWirelessProfile](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestCameraCreateNetworkCameraWirelessProfile, nil)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//CreateOrganizationCameraCustomAnalyticsArtifact Create custom analytics artifact
/* Create custom analytics artifact. Returns an artifact upload URL with expiry time. Upload the artifact file with a put request to the returned upload URL before its expiry.

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CameraService) CreateOrganizationCameraCustomAnalyticsArtifact(organizationID string, requestCameraCreateOrganizationCameraCustomAnalyticsArtifact *RequestCameraCreateOrganizationCameraCustomAnalyticsArtifact) (*ResponseCameraCreateOrganizationCameraCustomAnalyticsArtifact, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndResult[ResponseCameraCreateOrganizationCameraCustomAnalyticsArtifact](
		func() (*resty.Response, error) {
			return POST(path, s.client, requestCameraCreateOrganizationCameraCustomAnalyticsArtifact, nil)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//CreateOrganizationCameraRole Creates new role for this organization.
/* Creates new role for this organization.

@param organizationID organizationId path parameter. Organization ID


*/

func (s *CameraService) CreateOrganizationCameraRole(organizationID string, requestCameraCreateOrganizationCameraRole *RequestCameraCreateOrganizationCameraRole) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/roles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Past way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return POST(path, s.client, requestCameraCreateOrganizationCameraRole, nil)
		},
		s.backoff,
	)

}

//UpdateDeviceCameraCustomAnalytics Update custom analytics settings for a camera
/* Update custom analytics settings for a camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraCustomAnalytics(serial string, requestCameraUpdateDeviceCameraCustomAnalytics *RequestCameraUpdateDeviceCameraCustomAnalytics) (*ResponseCameraUpdateDeviceCameraCustomAnalytics, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/customAnalytics"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraUpdateDeviceCameraCustomAnalytics](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateDeviceCameraCustomAnalytics)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//UpdateDeviceCameraQualityAndRetention Update quality and retention settings for the given camera
/* Update quality and retention settings for the given camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraQualityAndRetention(serial string, requestCameraUpdateDeviceCameraQualityAndRetention *RequestCameraUpdateDeviceCameraQualityAndRetention) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/qualityAndRetention"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateDeviceCameraQualityAndRetention)
		},
		s.backoff,
	)

}

//UpdateDeviceCameraSense Update sense settings for the given camera
/* Update sense settings for the given camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraSense(serial string, requestCameraUpdateDeviceCameraSense *RequestCameraUpdateDeviceCameraSense) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/sense"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateDeviceCameraSense)
		},
		s.backoff,
	)

}

//UpdateDeviceCameraVideoSettings Update video settings for the given camera
/* Update video settings for the given camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraVideoSettings(serial string, requestCameraUpdateDeviceCameraVideoSettings *RequestCameraUpdateDeviceCameraVideoSettings) (*ResponseCameraUpdateDeviceCameraVideoSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/video/settings"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraUpdateDeviceCameraVideoSettings](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateDeviceCameraVideoSettings)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//UpdateDeviceCameraWirelessProfiles Assign wireless profiles to the given camera
/* Assign wireless profiles to the given camera. Incremental updates are not supported, all profile assignment need to be supplied at once.

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraWirelessProfiles(serial string, requestCameraUpdateDeviceCameraWirelessProfiles *RequestCameraUpdateDeviceCameraWirelessProfiles) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/wirelessProfiles"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateDeviceCameraWirelessProfiles)
		},
		s.backoff,
	)

}

//UpdateNetworkCameraQualityRetentionProfile Update an existing quality retention profile for this network.
/* Update an existing quality retention profile for this network.

@param networkID networkId path parameter. Network ID
@param qualityRetentionProfileID qualityRetentionProfileId path parameter. Quality retention profile ID
*/
func (s *CameraService) UpdateNetworkCameraQualityRetentionProfile(networkID string, qualityRetentionProfileID string, requestCameraUpdateNetworkCameraQualityRetentionProfile *RequestCameraUpdateNetworkCameraQualityRetentionProfile) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qualityRetentionProfileId}", fmt.Sprintf("%v", qualityRetentionProfileID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateNetworkCameraQualityRetentionProfile)
		},
		s.backoff,
	)

}

//UpdateNetworkCameraWirelessProfile Update an existing camera wireless profile in this network.
/* Update an existing camera wireless profile in this network.

@param networkID networkId path parameter. Network ID
@param wirelessProfileID wirelessProfileId path parameter. Wireless profile ID
*/
func (s *CameraService) UpdateNetworkCameraWirelessProfile(networkID string, wirelessProfileID string, requestCameraUpdateNetworkCameraWirelessProfile *RequestCameraUpdateNetworkCameraWirelessProfile) (*ResponseCameraUpdateNetworkCameraWirelessProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{wirelessProfileId}", fmt.Sprintf("%v", wirelessProfileID), -1)

	// Other way

	return doWithRetriesAndResult[ResponseCameraUpdateNetworkCameraWirelessProfile](
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateNetworkCameraWirelessProfile)
		},
		s.client,
		s.backoff,
		nil,
	)

}

//UpdateOrganizationCameraOnboardingStatuses Notify that credential handoff to camera has completed
/* Notify that credential handoff to camera has completed

@param organizationID organizationId path parameter. Organization ID
*/
func (s *CameraService) UpdateOrganizationCameraOnboardingStatuses(organizationID string, requestCameraUpdateOrganizationCameraOnboardingStatuses *RequestCameraUpdateOrganizationCameraOnboardingStatuses) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/onboarding/statuses"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateOrganizationCameraOnboardingStatuses)
		},
		s.backoff,
	)

}

//UpdateOrganizationCameraRole Update an existing role in this organization.
/* Update an existing role in this organization.

@param organizationID organizationId path parameter. Organization ID
@param roleID roleId path parameter. Role ID
*/
func (s *CameraService) UpdateOrganizationCameraRole(organizationID string, roleID string, requestCameraUpdateOrganizationCameraRole *RequestCameraUpdateOrganizationCameraRole) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/roles/{roleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleID), -1)

	// Other way

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return PUT(path, s.client, requestCameraUpdateOrganizationCameraRole)
		},
		s.backoff,
	)

}

//DeleteNetworkCameraQualityRetentionProfile Delete an existing quality retention profile for this network.
/* Delete an existing quality retention profile for this network.

@param networkID networkId path parameter. Network ID
@param qualityRetentionProfileID qualityRetentionProfileId path parameter. Quality retention profile ID


*/
func (s *CameraService) DeleteNetworkCameraQualityRetentionProfile(networkID string, qualityRetentionProfileID string) (*resty.Response, error) {
	//networkID string,qualityRetentionProfileID string
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qualityRetentionProfileId}", fmt.Sprintf("%v", qualityRetentionProfileID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
		s.backoff,
	)
}

//DeleteNetworkCameraWirelessProfile Delete an existing camera wireless profile for this network.
/* Delete an existing camera wireless profile for this network.

@param networkID networkId path parameter. Network ID
@param wirelessProfileID wirelessProfileId path parameter. Wireless profile ID


*/
func (s *CameraService) DeleteNetworkCameraWirelessProfile(networkID string, wirelessProfileID string) (*resty.Response, error) {
	//networkID string,wirelessProfileID string
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{wirelessProfileId}", fmt.Sprintf("%v", wirelessProfileID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
		s.backoff,
	)
}

//DeleteOrganizationCameraCustomAnalyticsArtifact Delete Custom Analytics Artifact
/* Delete Custom Analytics Artifact

@param organizationID organizationId path parameter. Organization ID
@param artifactID artifactId path parameter. Artifact ID


*/
func (s *CameraService) DeleteOrganizationCameraCustomAnalyticsArtifact(organizationID string, artifactID string) (*resty.Response, error) {
	//organizationID string,artifactID string
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{artifactId}", fmt.Sprintf("%v", artifactID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
		s.backoff,
	)
}

//DeleteOrganizationCameraRole Delete an existing role for this organization.
/* Delete an existing role for this organization.

@param organizationID organizationId path parameter. Organization ID
@param roleID roleId path parameter. Role ID


*/
func (s *CameraService) DeleteOrganizationCameraRole(organizationID string, roleID string) (*resty.Response, error) {
	//organizationID string,roleID string
	path := "/api/v1/organizations/{organizationId}/camera/roles/{roleId}"
	s.rateLimiterBucket.Wait(1)
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleID), -1)

	return doWithRetriesAndNotResult(
		func() (*resty.Response, error) {
			return DELETE(path, s.client, &QueryParamsDefault)
		},
		s.backoff,
	)
}
