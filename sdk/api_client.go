package meraki

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"io/ioutil"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

var apiURL = "https://api.meraki.com/"

const MERAKI_BASE_URL = "MERAKI_BASE_URL"
const MERAKI_DASHBOARD_API_KEY = "MERAKI_DASHBOARD_API_KEY"
const MERAKI_DEBUG = "MERAKI_DEBUG"
const MERAKI_SSL_VERIFY = "MERAKI_SSL_VERIFY"

type FileDownload struct {
	FileName string
	FileData []byte
}

func (f *FileDownload) SaveDownload(path string) error {
	fpath := filepath.Join(path, f.FileName)
	return ioutil.WriteFile(fpath, f.FileData, 0664)
}

// Client manages communication with the Cisco DNA Center API
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
	client *resty.Client
}

// SetAuthToken defines the X-Auth-Token token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	accessToken = "Bearer " + accessToken
	s.common.client.SetHeader("Authorization", accessToken)
}

func (s *Client) SetUserAgent(userAgent string) {
	s.common.client.SetHeader("User-Agent", userAgent)
}

// Error indicates an error from the invocation of a Cisco DNA Center API.
var Error map[string]interface{}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient() (*Client, error) {
	var err error
	c, err := NewClientNoAuth()
	if err != nil {
		return nil, err
	}

	err = c.AuthClient()
	if err != nil {
		return c, err
	}

	c.SetUserAgent("MerakiGolang/2.0.2 Cisco")

	c.common.client.AddRetryCondition(
		// RetryConditionFunc type is for retry condition function
		// input: non-nil Response OR request execution error
		func(r *resty.Response, err error) bool {
			return r.StatusCode() == http.StatusTooManyRequests
		},
	)
	c.common.client.SetRetryCount(2)
	c.common.client.SetRetryWaitTime(5 * time.Second)
	return c, nil
}

// NewClientWithOptions is the client with options passed with parameters
func NewClientWithOptions(baseURL string, dashboardApiKey string, debug string, sslVerify string) (*Client, error) {
	var err error

	err = SetOptions(baseURL, dashboardApiKey, debug, sslVerify)
	if err != nil {
		return nil, err
	}

	return NewClient()
}

// SetOptions sets the environment variables
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

// NewClientNoAuth returns the client object without trying to authenticate
func NewClientNoAuth() (*Client, error) {
	var err error

	client := resty.New()
	c := &Client{}
	c.common.client = client

	if os.Getenv(MERAKI_DEBUG) == "true" {
		client.SetDebug(true)
	}

	if os.Getenv(MERAKI_SSL_VERIFY) == "false" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
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
func NewClientWithOptionsNoAuth(baseURL string, dashboardApiKey string, debug string, sslVerify string) (*Client, error) {
	var err error

	err = SetOptions(baseURL, dashboardApiKey, debug, sslVerify)
	if err != nil {
		return nil, err
	}

	return NewClientNoAuth()
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
