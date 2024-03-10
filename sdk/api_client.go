package meraki

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/juju/ratelimit"
)

const (
	MERAKI_BASE_URL          = "MERAKI_BASE_URL"
	MERAKI_DASHBOARD_API_KEY = "MERAKI_DASHBOARD_API_KEY"
	MERAKI_DEBUG             = "MERAKI_DEBUG"
	MERAKI_SSL_VERIFY        = "MERAKI_SSL_VERIFY"
	DEFAULT_USER_AGENT       = "MerakiGolang/2.0.4 Cisco"
)

// Client manages communication with the Cisco Meraki API
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

// SetAuthToken sets the Authorization bearer token sent in the request
func (c *Client) SetAuthToken(accessToken string) {
	accessToken = "Bearer " + accessToken
	c.common.client.SetHeader("Authorization", accessToken)
}

// SetUserAgent sets the User-Agent header in the request
func (c *Client) SetUserAgent(userAgent string) {
	c.common.client.SetHeader("User-Agent", userAgent)
}

// Error indicates an error from the invocation of a Cisco Meraki API.
var Error map[string]interface{}

// NewClient creates a new API client.
func NewClient() (*Client, error) {
	client := resty.New()
	c := &Client{}
	c.common.client = client
	c.common.rateLimiterBucket = ratelimit.NewBucketWithQuantum(time.Second, 10, 10)

	if os.Getenv(MERAKI_DEBUG) == "true" {
		client.SetDebug(true)
	}

	if os.Getenv(MERAKI_SSL_VERIFY) == "false" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if os.Getenv(MERAKI_BASE_URL) != "" {
		client.SetBaseURL(os.Getenv(MERAKI_BASE_URL))
	} else {
		return nil, fmt.Errorf("enviroment variable %s was not defined", MERAKI_BASE_URL)
	}

	if os.Getenv(MERAKI_DASHBOARD_API_KEY) != "" {
		c.SetAuthToken(os.Getenv(MERAKI_DASHBOARD_API_KEY))
	} else {
		return nil, fmt.Errorf("enviroment variable %s was not defined", MERAKI_DASHBOARD_API_KEY)
	}

	c.SetUserAgent(DEFAULT_USER_AGENT)
	client.SetLogger(&CustomLogger{})
	client.SetRetryCount(2)
	client.SetRetryWaitTime(time.Second)
	client.SetRetryMaxWaitTime(5 * time.Second)
	// Respect "Retry-After" header in response
	client.SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
		retryHeader := resp.Header().Get("Retry-After")
		if _, err := strconv.Atoi(retryHeader); err == nil {
			t, _ := time.ParseDuration(retryHeader + "s")
			return t, nil
		}
		return 0, nil
	})
	// Retry on 429
	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() == http.StatusTooManyRequests
		},
	)

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

// NewClientWithOptions creates a new API client with options passed with parameters
func NewClientWithOptions(baseURL string, dashboardApiKey string, debug string, sslVerify string) (*Client, error) {
	err := SetOptions(baseURL, dashboardApiKey, debug, sslVerify)
	if err != nil {
		return nil, err
	}

	return NewClient()
}

// SetOptions sets the required environment variables
func SetOptions(baseURL string, dashboardApiKey string, debug string, sslVerify string) error {
	var err error
	err = os.Setenv(MERAKI_BASE_URL, baseURL)
	if err != nil {
		return err
	}
	err = os.Setenv(MERAKI_DEBUG, debug)
	if err != nil {
		return err
	}
	err = os.Setenv(MERAKI_SSL_VERIFY, sslVerify)
	if err != nil {
		return err
	}
	err = os.Setenv(MERAKI_DASHBOARD_API_KEY, dashboardApiKey)
	if err != nil {
		return err
	}
	return nil
}

// RestyClient returns the resty.Client used by the sdk
func (c *Client) RestyClient() *resty.Client {
	return c.common.client
}

// Get performs a GET request
func Get[R any](client *resty.Client, rateLimiterBucket *ratelimit.Bucket, path string, result *R) (*R, *resty.Response, error) {
	rateLimiterBucket.Wait(1)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(result).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with get operation: %s", path)
	}

	return response.Result().(*R), response, err
}

// Post performs a POST request
func Post[B, R any](client *resty.Client, rateLimiterBucket *ratelimit.Bucket, path string, body *B, result *R) (*R, *resty.Response, error) {
	rateLimiterBucket.Wait(1)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(result).
		SetError(&Error).
		SetBody(body).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with post operation: %s", path)
	}

	return response.Result().(*R), response, err
}

// Delete performs a DELETE request
func Delete(client *resty.Client, rateLimiterBucket *ratelimit.Bucket, path string) (*resty.Response, error) {
	rateLimiterBucket.Wait(1)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return response, fmt.Errorf("error with delete operation: %s", path)
	}

	return response, err
}
