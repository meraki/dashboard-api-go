package meraki

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/juju/ratelimit"
)

var apiURL = "https://api.meraki.com/"

const MERAKI_BASE_URL = "MERAKI_BASE_URL"
const MERAKI_DASHBOARD_API_KEY = "MERAKI_DASHBOARD_API_KEY"
const MERAKI_DEBUG = "MERAKI_DEBUG"
const MERAKI_SSL_VERIFY = "MERAKI_SSL_VERIFY"
const MERAKI_USER_AGENT = "MERAKI_USER_AGENT"
const DEFAULT_USER_AGENT = "golang-meraki/1.33.0"

// Client manages communication with the Dashboard API
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	Administered    *AdministeredService
	Appliance       *ApplianceService
	Camera          *CameraService
	CellularGateway *CellularGatewayService
	Devices         *DevicesService
	Insight         *InsightService
	Licensing       *LicensingService
	Networks        *NetworksService
	Organizations   *OrganizationsService
	Sensor          *SensorService
	Sm              *SmService
	Switch          *SwitchService
	Wireless        *WirelessService
	CustomCall      *CustomCallService
}

type service struct {
	client            *resty.Client
	rateLimiterBucket *ratelimit.Bucket
}

// SetAuthToken defines the X-Auth-Token token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	accessToken = "Bearer " + accessToken
	s.common.client.SetHeader("Authorization", accessToken)
}

func (s *Client) SetUserAgent(userAgent string) {
	s.common.client.SetHeader("User-Agent", userAgent)
}

// Error indicates an error from the invocation of a Dashboard Meraki API.
var Error map[string]interface{}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient(request_per_second *int) (*Client, error) {
	var err error
	c, err := NewClientNoAuth(request_per_second)
	if err != nil {
		return nil, err
	}

	err = c.AuthClient()
	if err != nil {
		return c, err
	}
	var userAgent string
	if os.Getenv(MERAKI_USER_AGENT) != "" {
		userAgent = os.Getenv(MERAKI_USER_AGENT)
	} else {
		userAgent = DEFAULT_USER_AGENT
	}
	c.SetUserAgent(userAgent)
	c.common.client.SetLogger(&CustomLogger{})

	c.common.client.AddRetryCondition(
		// RetryConditionFunc type is for retry condition function
		// input: non-nil Response OR request execution error
		func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() == http.StatusTooManyRequests
		},
	)
	c.common.client.SetRetryCount(3)
	c.common.client.SetRetryWaitTime(time.Second)
	c.common.client.SetRetryMaxWaitTime(5 * time.Second)
	c.common.client.SetRetryAfter(func(cl *resty.Client, r *resty.Response) (time.Duration, error) {
		retryAfter, err := strconv.Atoi(r.Header().Get("Retry-After"))
		if err != nil {
			retryAfter = 0
		}
		return time.Duration(retryAfter), nil
	})

	return c, nil
}

// NewClientWithOptions is the client with options passed with parameters
func NewClientWithOptions(baseURL string, dashboardApiKey string, debug string, customUserAgent *string, request_per_second *int) (*Client, error) {
	var err error

	err = SetOptions(baseURL, dashboardApiKey, debug, customUserAgent)
	if err != nil {
		return nil, err
	}

	return NewClient(request_per_second)
}

// SetOptions sets the environment variables
func SetOptions(baseURL string, dashboardApiKey string, debug string, customUserAgent *string) error {
	var err error
	err = os.Setenv(MERAKI_BASE_URL, baseURL)
	if err != nil {
		return err
	}
	err = os.Setenv(MERAKI_DEBUG, debug)
	if err != nil {
		return err
	}
	err = os.Setenv(MERAKI_DASHBOARD_API_KEY, dashboardApiKey)
	if err != nil {
		return err
	}

	if customUserAgent == nil {
		custom := DEFAULT_USER_AGENT
		customUserAgent = &custom
	}

	err = os.Setenv(MERAKI_USER_AGENT, *customUserAgent)
	if err != nil {
		return err
	}

	return nil
}

// NewClientNoAuth returns the client object without trying to authenticate
func NewClientNoAuth(request_per_second *int) (*Client, error) {
	var err error

	client := resty.New()
	c := &Client{}
	c.common.client = client
	var request_per_second_no_pointer int
	if request_per_second != nil {
		request_per_second_no_pointer = *request_per_second
	} else {
		request_per_second_no_pointer = 10
	}
	c.common.rateLimiterBucket = ratelimit.NewBucketWithQuantum(time.Second, int64(request_per_second_no_pointer), int64(request_per_second_no_pointer))
	if os.Getenv(MERAKI_DEBUG) == "true" {
		client.SetDebug(true)
	}

	if os.Getenv(MERAKI_BASE_URL) != "" {
		client.SetBaseURL(os.Getenv(MERAKI_BASE_URL))
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", MERAKI_BASE_URL)
	}

	if err != nil {
		return nil, err
	}

	c.Administered = (*AdministeredService)(&c.common)
	c.Appliance = (*ApplianceService)(&c.common)
	c.Camera = (*CameraService)(&c.common)
	c.CellularGateway = (*CellularGatewayService)(&c.common)
	c.Devices = (*DevicesService)(&c.common)
	c.Insight = (*InsightService)(&c.common)
	c.Licensing = (*LicensingService)(&c.common)
	c.Networks = (*NetworksService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.Sensor = (*SensorService)(&c.common)
	c.Sm = (*SmService)(&c.common)
	c.Switch = (*SwitchService)(&c.common)
	c.Wireless = (*WirelessService)(&c.common)
	c.CustomCall = (*CustomCallService)(&c.common)

	return c, nil
}

// NewClientWithOptionsNoAuth returns the client object without trying to authenticate and sets environment variables
func NewClientWithOptionsNoAuth(baseURL string, dashboardApiKey string, debug string, customUserAgent *string, requestPerSecond *int) (*Client, error) {
	var err error

	err = SetOptions(baseURL, dashboardApiKey, debug, customUserAgent)
	if err != nil {
		return nil, err
	}

	return NewClientNoAuth(requestPerSecond)
}

func (s *Client) AuthClient() error {
	var meraki_key string
	if os.Getenv(MERAKI_DASHBOARD_API_KEY) != "" {
		meraki_key = os.Getenv(MERAKI_DASHBOARD_API_KEY)
	} else {
		return fmt.Errorf("enviroment variable %s was not defined", MERAKI_DASHBOARD_API_KEY)
	}

	s.SetAuthToken(meraki_key)

	return nil
}

// RestyClient returns the resty.Client used by the sdk
func (s *Client) RestyClient() *resty.Client {
	return s.common.client
}

func convertToString(i interface{}) string {
	switch v := i.(type) {
	case float64:
		return strconv.Itoa(int(v))
	case string:
		return v
	default:
		return ""
	}
}
