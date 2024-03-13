package meraki

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
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
type GetOrganizationCameraOnboardingStatusesQueryParams struct {
	Serials    []string `url:"serials[],omitempty"`    //A list of serial numbers. The returned cameras will be filtered to only include these serials.
	NetworkIDs []string `url:"networkIds[],omitempty"` //A list of network IDs. The returned cameras will be filtered to only include these networks.
}

type ResponseCameraGetDeviceCameraAnalyticsLive struct {
	Ts    string                                           `json:"ts,omitempty"`    //
	Zones *ResponseCameraGetDeviceCameraAnalyticsLiveZones `json:"zones,omitempty"` //
}
type ResponseCameraGetDeviceCameraAnalyticsLiveZones struct {
	Status0 *ResponseCameraGetDeviceCameraAnalyticsLiveZones0 `json:"0,omitempty"` //
}
type ResponseCameraGetDeviceCameraAnalyticsLiveZones0 struct {
	Person *int `json:"person,omitempty"` //
}
type ResponseCameraGetDeviceCameraAnalyticsOverview []ResponseItemCameraGetDeviceCameraAnalyticsOverview // Array of ResponseCameraGetDeviceCameraAnalyticsOverview
type ResponseItemCameraGetDeviceCameraAnalyticsOverview struct {
	AverageCount *int   `json:"averageCount,omitempty"` //
	EndTs        string `json:"endTs,omitempty"`        //
	Entrances    *int   `json:"entrances,omitempty"`    //
	StartTs      string `json:"startTs,omitempty"`      //
	ZoneID       *int   `json:"zoneId,omitempty"`       //
}
type ResponseCameraGetDeviceCameraAnalyticsRecent []ResponseItemCameraGetDeviceCameraAnalyticsRecent // Array of ResponseCameraGetDeviceCameraAnalyticsRecent
type ResponseItemCameraGetDeviceCameraAnalyticsRecent struct {
	AverageCount *float64 `json:"averageCount,omitempty"` //
	EndTs        string   `json:"endTs,omitempty"`        //
	Entrances    *int     `json:"entrances,omitempty"`    //
	StartTs      string   `json:"startTs,omitempty"`      //
	ZoneID       *int     `json:"zoneId,omitempty"`       //
}
type ResponseCameraGetDeviceCameraAnalyticsZones []ResponseItemCameraGetDeviceCameraAnalyticsZones // Array of ResponseCameraGetDeviceCameraAnalyticsZones
type ResponseItemCameraGetDeviceCameraAnalyticsZones struct {
	ID               string                                                           `json:"id,omitempty"`               //
	Label            string                                                           `json:"label,omitempty"`            //
	RegionOfInterest *ResponseItemCameraGetDeviceCameraAnalyticsZonesRegionOfInterest `json:"regionOfInterest,omitempty"` //
	Type             string                                                           `json:"type,omitempty"`             //
}
type ResponseItemCameraGetDeviceCameraAnalyticsZonesRegionOfInterest struct {
	X0 string `json:"x0,omitempty"` //
	X1 string `json:"x1,omitempty"` //
	Y0 string `json:"y0,omitempty"` //
	Y1 string `json:"y1,omitempty"` //
}
type ResponseCameraGetDeviceCameraAnalyticsZoneHistory []ResponseItemCameraGetDeviceCameraAnalyticsZoneHistory // Array of ResponseCameraGetDeviceCameraAnalyticsZoneHistory
type ResponseItemCameraGetDeviceCameraAnalyticsZoneHistory struct {
	AverageCount *float64 `json:"averageCount,omitempty"` //
	EndTs        string   `json:"endTs,omitempty"`        //
	Entrances    *int     `json:"entrances,omitempty"`    //
	StartTs      string   `json:"startTs,omitempty"`      //
}
type ResponseCameraGetDeviceCameraCustomAnalytics struct {
	ArtifactID string                                                    `json:"artifactId,omitempty"` //
	Enabled    *bool                                                     `json:"enabled,omitempty"`    //
	Parameters *[]ResponseCameraGetDeviceCameraCustomAnalyticsParameters `json:"parameters,omitempty"` //
}
type ResponseCameraGetDeviceCameraCustomAnalyticsParameters struct {
	Name  string   `json:"name,omitempty"`  //
	Value *float64 `json:"value,omitempty"` //
}
type ResponseCameraUpdateDeviceCameraCustomAnalytics interface{}
type ResponseCameraGenerateDeviceCameraSnapshot interface{}
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
	ExternalRtspEnabled *bool  `json:"externalRtspEnabled,omitempty"` //
	RtspURL             string `json:"rtspUrl,omitempty"`             //
}
type ResponseCameraUpdateDeviceCameraVideoSettings interface{}
type ResponseCameraGetDeviceCameraVideoLink struct {
	URL string `json:"url,omitempty"` //
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
	AudioRecordingEnabled          *bool                                                                    `json:"audioRecordingEnabled,omitempty"`          //
	CloudArchiveEnabled            *bool                                                                    `json:"cloudArchiveEnabled,omitempty"`            //
	ID                             string                                                                   `json:"id,omitempty"`                             //
	MaxRetentionDays               *int                                                                     `json:"maxRetentionDays,omitempty"`               //
	MotionBasedRetentionEnabled    *bool                                                                    `json:"motionBasedRetentionEnabled,omitempty"`    //
	MotionDetectorVersion          *int                                                                     `json:"motionDetectorVersion,omitempty"`          //
	Name                           string                                                                   `json:"name,omitempty"`                           //
	NetworkID                      string                                                                   `json:"networkId,omitempty"`                      //
	RestrictedBandwidthModeEnabled *bool                                                                    `json:"restrictedBandwidthModeEnabled,omitempty"` //
	VideoSettings                  *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings `json:"videoSettings,omitempty"`                  //
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
	AudioRecordingEnabled          *bool                                                               `json:"audioRecordingEnabled,omitempty"`          //
	CloudArchiveEnabled            *bool                                                               `json:"cloudArchiveEnabled,omitempty"`            //
	ID                             string                                                              `json:"id,omitempty"`                             //
	MaxRetentionDays               *int                                                                `json:"maxRetentionDays,omitempty"`               //
	MotionBasedRetentionEnabled    *bool                                                               `json:"motionBasedRetentionEnabled,omitempty"`    //
	MotionDetectorVersion          *int                                                                `json:"motionDetectorVersion,omitempty"`          //
	Name                           string                                                              `json:"name,omitempty"`                           //
	NetworkID                      string                                                              `json:"networkId,omitempty"`                      //
	RestrictedBandwidthModeEnabled *bool                                                               `json:"restrictedBandwidthModeEnabled,omitempty"` //
	VideoSettings                  *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings `json:"videoSettings,omitempty"`                  //
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
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}
type ResponseCameraGetNetworkCameraWirelessProfiles []ResponseItemCameraGetNetworkCameraWirelessProfiles // Array of ResponseCameraGetNetworkCameraWirelessProfiles
type ResponseItemCameraGetNetworkCameraWirelessProfiles struct {
	AppliedDeviceCount *int                                                        `json:"appliedDeviceCount,omitempty"` //
	ID                 string                                                      `json:"id,omitempty"`                 //
	IDentity           *ResponseItemCameraGetNetworkCameraWirelessProfilesIDentity `json:"identity,omitempty"`           //
	Name               string                                                      `json:"name,omitempty"`               //
	SSID               *ResponseItemCameraGetNetworkCameraWirelessProfilesSSID     `json:"ssid,omitempty"`               //
}
type ResponseItemCameraGetNetworkCameraWirelessProfilesIDentity struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}
type ResponseItemCameraGetNetworkCameraWirelessProfilesSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       //
	EncryptionMode string `json:"encryptionMode,omitempty"` //
	Name           string `json:"name,omitempty"`           //
}
type ResponseCameraCreateNetworkCameraWirelessProfile interface{}
type ResponseCameraGetNetworkCameraWirelessProfile struct {
	AppliedDeviceCount *int                                                   `json:"appliedDeviceCount,omitempty"` //
	ID                 string                                                 `json:"id,omitempty"`                 //
	IDentity           *ResponseCameraGetNetworkCameraWirelessProfileIDentity `json:"identity,omitempty"`           //
	Name               string                                                 `json:"name,omitempty"`               //
	SSID               *ResponseCameraGetNetworkCameraWirelessProfileSSID     `json:"ssid,omitempty"`               //
}
type ResponseCameraGetNetworkCameraWirelessProfileIDentity struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}
type ResponseCameraGetNetworkCameraWirelessProfileSSID struct {
	AuthMode       string `json:"authMode,omitempty"`       //
	EncryptionMode string `json:"encryptionMode,omitempty"` //
	Name           string `json:"name,omitempty"`           //
}
type ResponseCameraUpdateNetworkCameraWirelessProfile interface{}
type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts []ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts // Array of ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts
type ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts struct {
	ArtifactID     string                                                                 `json:"artifactId,omitempty"`     //
	Name           string                                                                 `json:"name,omitempty"`           //
	OrganizationID string                                                                 `json:"organizationId,omitempty"` //
	Status         *ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus `json:"status,omitempty"`         //
}
type ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus struct {
	Type string `json:"type,omitempty"` //
}
type ResponseCameraCreateOrganizationCameraCustomAnalyticsArtifact interface{}
type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact struct {
	ArtifactID     string                                                            `json:"artifactId,omitempty"`     //
	Name           string                                                            `json:"name,omitempty"`           //
	OrganizationID string                                                            `json:"organizationId,omitempty"` //
	Status         *ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus `json:"status,omitempty"`         //
}
type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus struct {
	Type string `json:"type,omitempty"` //
}
type ResponseCameraGetOrganizationCameraOnboardingStatuses []ResponseItemCameraGetOrganizationCameraOnboardingStatuses // Array of ResponseCameraGetOrganizationCameraOnboardingStatuses
type ResponseItemCameraGetOrganizationCameraOnboardingStatuses struct {
	NetworkID string `json:"networkId,omitempty"` //
	Serial    string `json:"serial,omitempty"`    //
	Status    string `json:"status,omitempty"`    //
	UpdatedAt string `json:"updatedAt,omitempty"` //
}
type ResponseCameraUpdateOrganizationCameraOnboardingStatuses interface{}
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
	Quality                        string `json:"quality,omitempty"`                        // Quality of the camera. Can be one of 'Standard', 'High' or 'Enhanced'. Not all qualities are supported by every camera model.
	Resolution                     string `json:"resolution,omitempty"`                     // Resolution of the camera. Can be one of '1280x720', '1920x1080', '1080x1080', '2058x2058', '2112x2112', '2880x2880', '2688x1512' or '3840x2160'.Not all resolutions are supported by every camera model.
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
	AudioRecordingEnabled          *bool                                                                 `json:"audioRecordingEnabled,omitempty"`          // Whether or not to record audio. Can be either true or false. Defaults to false.
	CloudArchiveEnabled            *bool                                                                 `json:"cloudArchiveEnabled,omitempty"`            // Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.
	MaxRetentionDays               *int                                                                  `json:"maxRetentionDays,omitempty"`               // The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be one of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 14, 30, 60, 90] days.
	MotionBasedRetentionEnabled    *bool                                                                 `json:"motionBasedRetentionEnabled,omitempty"`    // Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	MotionDetectorVersion          *int                                                                  `json:"motionDetectorVersion,omitempty"`          // The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	Name                           string                                                                `json:"name,omitempty"`                           // The name of the new profile. Must be unique. This parameter is required.
	RestrictedBandwidthModeEnabled *bool                                                                 `json:"restrictedBandwidthModeEnabled,omitempty"` // Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	ScheduleID                     string                                                                `json:"scheduleId,omitempty"`                     // Schedule for which this camera will record video, or 'null' to always record.
	VideoSettings                  *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings `json:"videoSettings,omitempty"`                  // Video quality and resolution settings for all the camera models.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings struct {
	MV12MV22MV72 *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"` // Quality and resolution for MV12/MV22/MV72 camera models.
	MV12WE       *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE       `json:"MV12WE,omitempty"`         // Quality and resolution for MV12WE camera models.
	MV13         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13         `json:"MV13,omitempty"`           // Quality and resolution for MV13 camera models.
	MV21MV71     *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71     `json:"MV21/MV71,omitempty"`      // Quality and resolution for MV21/MV71 camera models.
	MV22XMV72X   *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X   `json:"MV22X/MV72X,omitempty"`    // Quality and resolution for MV22X/MV72X camera models.
	MV32         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32         `json:"MV32,omitempty"`           // Quality and resolution for MV32 camera models.
	MV33         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33         `json:"MV33,omitempty"`           // Quality and resolution for MV33 camera models.
	MV52         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52         `json:"MV52,omitempty"`           // Quality and resolution for MV52 camera models.
	MV63         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63         `json:"MV63,omitempty"`           // Quality and resolution for MV63 camera models.
	MV63X        *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X        `json:"MV63X,omitempty"`          // Quality and resolution for MV63X camera models.
	MV93         *RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93         `json:"MV93,omitempty"`           // Quality and resolution for MV93 camera models.
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
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2688x1512'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080' or '2688x1512'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2058x2058'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080' or '2688x1512'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
}
type RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfile struct {
	AudioRecordingEnabled          *bool                                                                 `json:"audioRecordingEnabled,omitempty"`          // Whether or not to record audio. Can be either true or false. Defaults to false.
	CloudArchiveEnabled            *bool                                                                 `json:"cloudArchiveEnabled,omitempty"`            // Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.
	MaxRetentionDays               *int                                                                  `json:"maxRetentionDays,omitempty"`               // The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be one of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 14, 30, 60, 90] days.
	MotionBasedRetentionEnabled    *bool                                                                 `json:"motionBasedRetentionEnabled,omitempty"`    // Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	MotionDetectorVersion          *int                                                                  `json:"motionDetectorVersion,omitempty"`          // The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	Name                           string                                                                `json:"name,omitempty"`                           // The name of the new profile. Must be unique.
	RestrictedBandwidthModeEnabled *bool                                                                 `json:"restrictedBandwidthModeEnabled,omitempty"` // Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	ScheduleID                     string                                                                `json:"scheduleId,omitempty"`                     // Schedule for which this camera will record video, or 'null' to always record.
	VideoSettings                  *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings `json:"videoSettings,omitempty"`                  // Video quality and resolution settings for all the camera models.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings struct {
	MV12MV22MV72 *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"` // Quality and resolution for MV12/MV22/MV72 camera models.
	MV12WE       *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE       `json:"MV12WE,omitempty"`         // Quality and resolution for MV12WE camera models.
	MV13         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13         `json:"MV13,omitempty"`           // Quality and resolution for MV13 camera models.
	MV21MV71     *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71     `json:"MV21/MV71,omitempty"`      // Quality and resolution for MV21/MV71 camera models.
	MV22XMV72X   *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X   `json:"MV22X/MV72X,omitempty"`    // Quality and resolution for MV22X/MV72X camera models.
	MV32         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32         `json:"MV32,omitempty"`           // Quality and resolution for MV32 camera models.
	MV33         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33         `json:"MV33,omitempty"`           // Quality and resolution for MV33 camera models.
	MV52         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52         `json:"MV52,omitempty"`           // Quality and resolution for MV52 camera models.
	MV63         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63         `json:"MV63,omitempty"`           // Quality and resolution for MV63 camera models.
	MV63X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X        `json:"MV63X,omitempty"`          // Quality and resolution for MV63X camera models.
	MV93         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93         `json:"MV93,omitempty"`           // Quality and resolution for MV93 camera models.
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
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2688x1512'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080' or '2688x1512'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2058x2058'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080' or '2688x1512'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
}
type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93 struct {
	Quality    string `json:"quality,omitempty"`    // Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Resolution string `json:"resolution,omitempty"` // Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
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

//GetDeviceCameraAnalyticsLive Returns live state from camera of analytics zones
/* Returns live state from camera of analytics zones

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-analytics-live
*/
func (s *CameraService) GetDeviceCameraAnalyticsLive(serial string) (*ResponseCameraGetDeviceCameraAnalyticsLive, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/live"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraAnalyticsLive{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraAnalyticsLive")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraAnalyticsLive)
	return result, response, err

}

//GetDeviceCameraAnalyticsOverview Returns an overview of aggregate analytics data for a timespan
/* Returns an overview of aggregate analytics data for a timespan

@param serial serial path parameter.
@param getDeviceCameraAnalyticsOverviewQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-analytics-overview
*/
func (s *CameraService) GetDeviceCameraAnalyticsOverview(serial string, getDeviceCameraAnalyticsOverviewQueryParams *GetDeviceCameraAnalyticsOverviewQueryParams) (*ResponseCameraGetDeviceCameraAnalyticsOverview, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/overview"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceCameraAnalyticsOverviewQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCameraGetDeviceCameraAnalyticsOverview{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraAnalyticsOverview")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraAnalyticsOverview)
	return result, response, err

}

//GetDeviceCameraAnalyticsRecent Returns most recent record for analytics zones
/* Returns most recent record for analytics zones

@param serial serial path parameter.
@param getDeviceCameraAnalyticsRecentQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-analytics-recent
*/
func (s *CameraService) GetDeviceCameraAnalyticsRecent(serial string, getDeviceCameraAnalyticsRecentQueryParams *GetDeviceCameraAnalyticsRecentQueryParams) (*ResponseCameraGetDeviceCameraAnalyticsRecent, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/recent"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceCameraAnalyticsRecentQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCameraGetDeviceCameraAnalyticsRecent{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraAnalyticsRecent")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraAnalyticsRecent)
	return result, response, err

}

//GetDeviceCameraAnalyticsZones Returns all configured analytic zones for this camera
/* Returns all configured analytic zones for this camera

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-analytics-zones
*/
func (s *CameraService) GetDeviceCameraAnalyticsZones(serial string) (*ResponseCameraGetDeviceCameraAnalyticsZones, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/zones"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraAnalyticsZones{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraAnalyticsZones")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraAnalyticsZones)
	return result, response, err

}

//GetDeviceCameraAnalyticsZoneHistory Return historical records for analytic zones
/* Return historical records for analytic zones

@param serial serial path parameter.
@param zoneID zoneId path parameter. Zone ID
@param getDeviceCameraAnalyticsZoneHistoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-analytics-zone-history
*/
func (s *CameraService) GetDeviceCameraAnalyticsZoneHistory(serial string, zoneID string, getDeviceCameraAnalyticsZoneHistoryQueryParams *GetDeviceCameraAnalyticsZoneHistoryQueryParams) (*ResponseCameraGetDeviceCameraAnalyticsZoneHistory, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/analytics/zones/{zoneId}/history"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)
	path = strings.Replace(path, "{zoneId}", fmt.Sprintf("%v", zoneID), -1)

	queryString, _ := query.Values(getDeviceCameraAnalyticsZoneHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCameraGetDeviceCameraAnalyticsZoneHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraAnalyticsZoneHistory")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraAnalyticsZoneHistory)
	return result, response, err

}

//GetDeviceCameraCustomAnalytics Return custom analytics settings for a camera
/* Return custom analytics settings for a camera

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-custom-analytics
*/
func (s *CameraService) GetDeviceCameraCustomAnalytics(serial string) (*ResponseCameraGetDeviceCameraCustomAnalytics, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/customAnalytics"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraCustomAnalytics{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraCustomAnalytics")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraCustomAnalytics)
	return result, response, err

}

//GetDeviceCameraQualityAndRetention Returns quality and retention settings for the given camera
/* Returns quality and retention settings for the given camera

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-quality-and-retention
*/
func (s *CameraService) GetDeviceCameraQualityAndRetention(serial string) (*ResponseCameraGetDeviceCameraQualityAndRetention, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/qualityAndRetention"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraQualityAndRetention{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraQualityAndRetention")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraQualityAndRetention)
	return result, response, err

}

//GetDeviceCameraSense Returns sense settings for a given camera
/* Returns sense settings for a given camera

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-sense
*/
func (s *CameraService) GetDeviceCameraSense(serial string) (*ResponseCameraGetDeviceCameraSense, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/sense"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraSense{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraSense")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraSense)
	return result, response, err

}

//GetDeviceCameraSenseObjectDetectionModels Returns the MV Sense object detection model list for the given camera
/* Returns the MV Sense object detection model list for the given camera

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-sense-object-detection-models
*/
func (s *CameraService) GetDeviceCameraSenseObjectDetectionModels(serial string) (*ResponseCameraGetDeviceCameraSenseObjectDetectionModels, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/sense/objectDetectionModels"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraSenseObjectDetectionModels{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraSenseObjectDetectionModels")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraSenseObjectDetectionModels)
	return result, response, err

}

//GetDeviceCameraVideoSettings Returns video settings for the given camera
/* Returns video settings for the given camera

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-video-settings
*/
func (s *CameraService) GetDeviceCameraVideoSettings(serial string) (*ResponseCameraGetDeviceCameraVideoSettings, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/video/settings"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraVideoSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraVideoSettings")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraVideoSettings)
	return result, response, err

}

//GetDeviceCameraVideoLink Returns video link to the specified camera
/* Returns video link to the specified camera. If a timestamp is supplied, it links to that timestamp.

@param serial serial path parameter.
@param getDeviceCameraVideoLinkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-video-link
*/
func (s *CameraService) GetDeviceCameraVideoLink(serial string, getDeviceCameraVideoLinkQueryParams *GetDeviceCameraVideoLinkQueryParams) (*ResponseCameraGetDeviceCameraVideoLink, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/videoLink"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	queryString, _ := query.Values(getDeviceCameraVideoLinkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCameraGetDeviceCameraVideoLink{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraVideoLink")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraVideoLink)
	return result, response, err

}

//GetDeviceCameraWirelessProfiles Returns wireless profile assigned to the given camera
/* Returns wireless profile assigned to the given camera

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-camera-wireless-profiles
*/
func (s *CameraService) GetDeviceCameraWirelessProfiles(serial string) (*ResponseCameraGetDeviceCameraWirelessProfiles, *resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/wirelessProfiles"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetDeviceCameraWirelessProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCameraWirelessProfiles")
	}

	result := response.Result().(*ResponseCameraGetDeviceCameraWirelessProfiles)
	return result, response, err

}

//GetNetworkCameraQualityRetentionProfiles List the quality retention profiles for this network
/* List the quality retention profiles for this network

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-camera-quality-retention-profiles
*/
func (s *CameraService) GetNetworkCameraQualityRetentionProfiles(networkID string) (*ResponseCameraGetNetworkCameraQualityRetentionProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetNetworkCameraQualityRetentionProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCameraQualityRetentionProfiles")
	}

	result := response.Result().(*ResponseCameraGetNetworkCameraQualityRetentionProfiles)
	return result, response, err

}

//GetNetworkCameraQualityRetentionProfile Retrieve a single quality retention profile
/* Retrieve a single quality retention profile

@param networkID networkId path parameter. Network ID
@param qualityRetentionProfileID qualityRetentionProfileId path parameter. Quality retention profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-camera-quality-retention-profile
*/
func (s *CameraService) GetNetworkCameraQualityRetentionProfile(networkID string, qualityRetentionProfileID string) (*ResponseCameraGetNetworkCameraQualityRetentionProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qualityRetentionProfileId}", fmt.Sprintf("%v", qualityRetentionProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetNetworkCameraQualityRetentionProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCameraQualityRetentionProfile")
	}

	result := response.Result().(*ResponseCameraGetNetworkCameraQualityRetentionProfile)
	return result, response, err

}

//GetNetworkCameraSchedules Returns a list of all camera recording schedules.
/* Returns a list of all camera recording schedules.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-camera-schedules
*/
func (s *CameraService) GetNetworkCameraSchedules(networkID string) (*ResponseCameraGetNetworkCameraSchedules, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/schedules"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetNetworkCameraSchedules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCameraSchedules")
	}

	result := response.Result().(*ResponseCameraGetNetworkCameraSchedules)
	return result, response, err

}

//GetNetworkCameraWirelessProfiles List the camera wireless profiles for this network.
/* List the camera wireless profiles for this network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-camera-wireless-profiles
*/
func (s *CameraService) GetNetworkCameraWirelessProfiles(networkID string) (*ResponseCameraGetNetworkCameraWirelessProfiles, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetNetworkCameraWirelessProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCameraWirelessProfiles")
	}

	result := response.Result().(*ResponseCameraGetNetworkCameraWirelessProfiles)
	return result, response, err

}

//GetNetworkCameraWirelessProfile Retrieve a single camera wireless profile.
/* Retrieve a single camera wireless profile.

@param networkID networkId path parameter. Network ID
@param wirelessProfileID wirelessProfileId path parameter. Wireless profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-camera-wireless-profile
*/
func (s *CameraService) GetNetworkCameraWirelessProfile(networkID string, wirelessProfileID string) (*ResponseCameraGetNetworkCameraWirelessProfile, *resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{wirelessProfileId}", fmt.Sprintf("%v", wirelessProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetNetworkCameraWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkCameraWirelessProfile")
	}

	result := response.Result().(*ResponseCameraGetNetworkCameraWirelessProfile)
	return result, response, err

}

//GetOrganizationCameraCustomAnalyticsArtifacts List Custom Analytics Artifacts
/* List Custom Analytics Artifacts

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-camera-custom-analytics-artifacts
*/
func (s *CameraService) GetOrganizationCameraCustomAnalyticsArtifacts(organizationID string) (*ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCameraCustomAnalyticsArtifacts")
	}

	result := response.Result().(*ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts)
	return result, response, err

}

//GetOrganizationCameraCustomAnalyticsArtifact Get Custom Analytics Artifact
/* Get Custom Analytics Artifact

@param organizationID organizationId path parameter. Organization ID
@param artifactID artifactId path parameter. Artifact ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-camera-custom-analytics-artifact
*/
func (s *CameraService) GetOrganizationCameraCustomAnalyticsArtifact(organizationID string, artifactID string) (*ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{artifactId}", fmt.Sprintf("%v", artifactID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCameraCustomAnalyticsArtifact")
	}

	result := response.Result().(*ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact)
	return result, response, err

}

//GetOrganizationCameraOnboardingStatuses Fetch onboarding status of cameras
/* Fetch onboarding status of cameras

@param organizationID organizationId path parameter. Organization ID
@param getOrganizationCameraOnboardingStatusesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-camera-onboarding-statuses
*/
func (s *CameraService) GetOrganizationCameraOnboardingStatuses(organizationID string, getOrganizationCameraOnboardingStatusesQueryParams *GetOrganizationCameraOnboardingStatusesQueryParams) (*ResponseCameraGetOrganizationCameraOnboardingStatuses, *resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/onboarding/statuses"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	queryString, _ := query.Values(getOrganizationCameraOnboardingStatusesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseCameraGetOrganizationCameraOnboardingStatuses{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationCameraOnboardingStatuses")
	}

	result := response.Result().(*ResponseCameraGetOrganizationCameraOnboardingStatuses)
	return result, response, err

}

//GenerateDeviceCameraSnapshot Generate a snapshot of what the camera sees at the specified time and return a link to that image.
/* Generate a snapshot of what the camera sees at the specified time and return a link to that image.

@param serial serial path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!generate-device-camera-snapshot
*/

func (s *CameraService) GenerateDeviceCameraSnapshot(serial string, requestCameraGenerateDeviceCameraSnapshot *RequestCameraGenerateDeviceCameraSnapshot) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/generateSnapshot"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraGenerateDeviceCameraSnapshot).
		// SetResult(&ResponseCameraGenerateDeviceCameraSnapshot{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GenerateDeviceCameraSnapshot")
	}

	return response, err

}

//CreateNetworkCameraQualityRetentionProfile Creates new quality retention profile for this network.
/* Creates new quality retention profile for this network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-camera-quality-retention-profile
*/

func (s *CameraService) CreateNetworkCameraQualityRetentionProfile(networkID string, requestCameraCreateNetworkCameraQualityRetentionProfile *RequestCameraCreateNetworkCameraQualityRetentionProfile) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraCreateNetworkCameraQualityRetentionProfile).
		// SetResult(&ResponseCameraCreateNetworkCameraQualityRetentionProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkCameraQualityRetentionProfile")
	}

	return response, err

}

//CreateNetworkCameraWirelessProfile Creates a new camera wireless profile for this network.
/* Creates a new camera wireless profile for this network.

@param networkID networkId path parameter. Network ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-camera-wireless-profile
*/

func (s *CameraService) CreateNetworkCameraWirelessProfile(networkID string, requestCameraCreateNetworkCameraWirelessProfile *RequestCameraCreateNetworkCameraWirelessProfile) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraCreateNetworkCameraWirelessProfile).
		// SetResult(&ResponseCameraCreateNetworkCameraWirelessProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateNetworkCameraWirelessProfile")
	}

	return response, err

}

//CreateOrganizationCameraCustomAnalyticsArtifact Create custom analytics artifact
/* Create custom analytics artifact. Returns an artifact upload URL with expiry time. Upload the artifact file with a put request to the returned upload URL before its expiry.

@param organizationID organizationId path parameter. Organization ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-organization-camera-custom-analytics-artifact
*/

func (s *CameraService) CreateOrganizationCameraCustomAnalyticsArtifact(organizationID string, requestCameraCreateOrganizationCameraCustomAnalyticsArtifact *RequestCameraCreateOrganizationCameraCustomAnalyticsArtifact) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraCreateOrganizationCameraCustomAnalyticsArtifact).
		// SetResult(&ResponseCameraCreateOrganizationCameraCustomAnalyticsArtifact{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation CreateOrganizationCameraCustomAnalyticsArtifact")
	}

	return response, err

}

//UpdateDeviceCameraCustomAnalytics Update custom analytics settings for a camera
/* Update custom analytics settings for a camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraCustomAnalytics(serial string, requestCameraUpdateDeviceCameraCustomAnalytics *RequestCameraUpdateDeviceCameraCustomAnalytics) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/customAnalytics"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateDeviceCameraCustomAnalytics).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCameraCustomAnalytics")
	}

	return response, err

}

//UpdateDeviceCameraQualityAndRetention Update quality and retention settings for the given camera
/* Update quality and retention settings for the given camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraQualityAndRetention(serial string, requestCameraUpdateDeviceCameraQualityAndRetention *RequestCameraUpdateDeviceCameraQualityAndRetention) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/qualityAndRetention"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateDeviceCameraQualityAndRetention).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCameraQualityAndRetention")
	}

	return response, err

}

//UpdateDeviceCameraSense Update sense settings for the given camera
/* Update sense settings for the given camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraSense(serial string, requestCameraUpdateDeviceCameraSense *RequestCameraUpdateDeviceCameraSense) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/sense"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateDeviceCameraSense).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCameraSense")
	}

	return response, err

}

//UpdateDeviceCameraVideoSettings Update video settings for the given camera
/* Update video settings for the given camera

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraVideoSettings(serial string, requestCameraUpdateDeviceCameraVideoSettings *RequestCameraUpdateDeviceCameraVideoSettings) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/video/settings"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateDeviceCameraVideoSettings).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCameraVideoSettings")
	}

	return response, err

}

//UpdateDeviceCameraWirelessProfiles Assign wireless profiles to the given camera
/* Assign wireless profiles to the given camera. Incremental updates are not supported, all profile assignment need to be supplied at once.

@param serial serial path parameter.
*/
func (s *CameraService) UpdateDeviceCameraWirelessProfiles(serial string, requestCameraUpdateDeviceCameraWirelessProfiles *RequestCameraUpdateDeviceCameraWirelessProfiles) (*resty.Response, error) {
	path := "/api/v1/devices/{serial}/camera/wirelessProfiles"
	path = strings.Replace(path, "{serial}", fmt.Sprintf("%v", serial), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateDeviceCameraWirelessProfiles).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateDeviceCameraWirelessProfiles")
	}

	return response, err

}

//UpdateNetworkCameraQualityRetentionProfile Update an existing quality retention profile for this network.
/* Update an existing quality retention profile for this network.

@param networkID networkId path parameter. Network ID
@param qualityRetentionProfileID qualityRetentionProfileId path parameter. Quality retention profile ID
*/
func (s *CameraService) UpdateNetworkCameraQualityRetentionProfile(networkID string, qualityRetentionProfileID string, requestCameraUpdateNetworkCameraQualityRetentionProfile *RequestCameraUpdateNetworkCameraQualityRetentionProfile) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qualityRetentionProfileId}", fmt.Sprintf("%v", qualityRetentionProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateNetworkCameraQualityRetentionProfile).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkCameraQualityRetentionProfile")
	}

	return response, err

}

//UpdateNetworkCameraWirelessProfile Update an existing camera wireless profile in this network.
/* Update an existing camera wireless profile in this network.

@param networkID networkId path parameter. Network ID
@param wirelessProfileID wirelessProfileId path parameter. Wireless profile ID
*/
func (s *CameraService) UpdateNetworkCameraWirelessProfile(networkID string, wirelessProfileID string, requestCameraUpdateNetworkCameraWirelessProfile *RequestCameraUpdateNetworkCameraWirelessProfile) (*resty.Response, error) {
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{wirelessProfileId}", fmt.Sprintf("%v", wirelessProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateNetworkCameraWirelessProfile).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateNetworkCameraWirelessProfile")
	}

	return response, err

}

//UpdateOrganizationCameraOnboardingStatuses Notify that credential handoff to camera has completed
/* Notify that credential handoff to camera has completed

@param organizationID organizationId path parameter. Organization ID
*/
func (s *CameraService) UpdateOrganizationCameraOnboardingStatuses(organizationID string, requestCameraUpdateOrganizationCameraOnboardingStatuses *RequestCameraUpdateOrganizationCameraOnboardingStatuses) (*resty.Response, error) {
	path := "/api/v1/organizations/{organizationId}/camera/onboarding/statuses"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCameraUpdateOrganizationCameraOnboardingStatuses).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation UpdateOrganizationCameraOnboardingStatuses")
	}

	return response, err

}

//DeleteNetworkCameraQualityRetentionProfile Delete an existing quality retention profile for this network.
/* Delete an existing quality retention profile for this network.

@param networkID networkId path parameter. Network ID
@param qualityRetentionProfileID qualityRetentionProfileId path parameter. Quality retention profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-camera-quality-retention-profile
*/
func (s *CameraService) DeleteNetworkCameraQualityRetentionProfile(networkID string, qualityRetentionProfileID string) (*resty.Response, error) {
	//networkID string,qualityRetentionProfileID string
	path := "/api/v1/networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{qualityRetentionProfileId}", fmt.Sprintf("%v", qualityRetentionProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkCameraQualityRetentionProfile")
	}

	return response, err

}

//DeleteNetworkCameraWirelessProfile Delete an existing camera wireless profile for this network.
/* Delete an existing camera wireless profile for this network.

@param networkID networkId path parameter. Network ID
@param wirelessProfileID wirelessProfileId path parameter. Wireless profile ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-camera-wireless-profile
*/
func (s *CameraService) DeleteNetworkCameraWirelessProfile(networkID string, wirelessProfileID string) (*resty.Response, error) {
	//networkID string,wirelessProfileID string
	path := "/api/v1/networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}"
	path = strings.Replace(path, "{networkId}", fmt.Sprintf("%v", networkID), -1)
	path = strings.Replace(path, "{wirelessProfileId}", fmt.Sprintf("%v", wirelessProfileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteNetworkCameraWirelessProfile")
	}

	return response, err

}

//DeleteOrganizationCameraCustomAnalyticsArtifact Delete Custom Analytics Artifact
/* Delete Custom Analytics Artifact

@param organizationID organizationId path parameter. Organization ID
@param artifactID artifactId path parameter. Artifact ID

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-organization-camera-custom-analytics-artifact
*/
func (s *CameraService) DeleteOrganizationCameraCustomAnalyticsArtifact(organizationID string, artifactID string) (*resty.Response, error) {
	//organizationID string,artifactID string
	path := "/api/v1/organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId}"
	path = strings.Replace(path, "{organizationId}", fmt.Sprintf("%v", organizationID), -1)
	path = strings.Replace(path, "{artifactId}", fmt.Sprintf("%v", artifactID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteOrganizationCameraCustomAnalyticsArtifact")
	}

	return response, err

}
