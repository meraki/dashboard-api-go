package meraki

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/juju/ratelimit"
)

const (
	MERAKI_BASE_URL             = "MERAKI_BASE_URL"
	MERAKI_USER_AGENT           = "MERAKI_USER_AGENT"
	MERAKI_DASHBOARD_API_KEY    = "MERAKI_DASHBOARD_API_KEY"
	MERAKI_DEBUG                = "MERAKI_DEBUG"
	MERAKI_REQUESTS_PER_SECOND  = "MERAKI_REQUESTS_PER_SECOND"
	DEFAULT_USER_AGENT          = "go-meraki/1.57.0"
	DEFAULT_REQUESTS_PER_SECOND = 10
	PAGINATION_PER_PAGE         = 1000
)

// Client manages communication with the Cisco Meraki API
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	Administered       *AdministeredService
	Appliance          *ApplianceService
	Camera             *CameraService
	CellularGateway    *CellularGatewayService
	Devices            *DevicesService
	Insight            *InsightService
	Licensing          *LicensingService
	Networks           *NetworksService
	Organizations      *OrganizationsService
	Sensor             *SensorService
	Sm                 *SmService
	Switch             *SwitchService
	Wireless           *WirelessService
	WirelessController *WirelessControllerService
	CustomCall         *CustomCallService
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
func (c *Client) SetUserAgent() {
	userAgent := fmt.Sprintf("%s %s", DEFAULT_USER_AGENT, os.Getenv("MERAKI_USER_AGENT"))
	c.common.client.SetHeader("User-Agent", userAgent)
}

// SetRequestsPerSecond sets the maximum number of requests per second
func (c *Client) SetRequestsPerSecond(requestsPerSecond int) {
	c.common.rateLimiterBucket = ratelimit.NewBucketWithQuantum(time.Second, int64(requestsPerSecond), int64(requestsPerSecond))
}

// Error indicates an error from the invocation of a Cisco Meraki API.
var Error map[string]interface{}

// NewClient creates a new API client.
func NewClient() (*Client, error) {
	client := resty.New()
	c := &Client{}
	c.common.client = client

	if os.Getenv(MERAKI_DEBUG) == "true" {
		client.SetDebug(true)
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

	if os.Getenv(MERAKI_REQUESTS_PER_SECOND) != "" {
		v, err := strconv.Atoi(os.Getenv(MERAKI_REQUESTS_PER_SECOND))
		if err != nil || v < 0 {
			return nil, fmt.Errorf("enviroment variable %s is not a valid integer greater or equal to zero", MERAKI_REQUESTS_PER_SECOND)
		}
		c.SetRequestsPerSecond(v)
	} else {
		c.SetRequestsPerSecond(DEFAULT_REQUESTS_PER_SECOND)
	}

	c.SetUserAgent()
	client.SetLogger(&CustomLogger{})
	client.SetRetryCount(2)
	client.SetRetryWaitTime(time.Second)
	client.SetRetryMaxWaitTime(1 * time.Minute)
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
	c.WirelessController = (*WirelessControllerService)(&c.common)
	c.CustomCall = (*CustomCallService)(&c.common)

	return c, nil
}

// NewClientWithOptions creates a new API client with options passed with parameters
func NewClientWithOptions(baseURL string, dashboardApiKey string, debug string, userAgent string) (*Client, error) {
	err := SetOptions(baseURL, dashboardApiKey, debug, userAgent)
	if err != nil {
		return nil, err
	}

	return NewClient()
}

// NewClientWithOptionsAndRequests creates a new API client with options passed with parameters including the requests per second
func NewClientWithOptionsAndRequests(baseURL string, dashboardApiKey string, debug string, userAgent string, requestsPerSecond int) (*Client, error) {
	err := SetOptionsWithRequests(baseURL, dashboardApiKey, debug, userAgent, requestsPerSecond)
	if err != nil {
		return nil, err
	}

	return NewClient()
}

// SetOptions sets the required environment variables
func SetOptions(baseURL string, dashboardApiKey string, debug string, userAgent string) error {
	var err error

	if !validateUserAgent(userAgent) {
		return errors.New("user-agent bad format, expected: `AplicationName VendorName Client`")
	}

	err = os.Setenv(MERAKI_USER_AGENT, userAgent)
	if err != nil {
		return err
	}

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
	return nil
}

// SetOptionsWithRequests sets the required environment variables including the requests per second
func SetOptionsWithRequests(baseURL string, dashboardApiKey string, debug string, userAgent string, requestsPerSecond int) error {
	err := SetOptions(baseURL, dashboardApiKey, debug, userAgent)
	if err != nil {
		return err
	}
	err = os.Setenv(MERAKI_REQUESTS_PER_SECOND, strconv.Itoa(requestsPerSecond))
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

// Put performs a PUT request
func Put[B, R any](client *resty.Client, rateLimiterBucket *ratelimit.Bucket, path string, body *B, result *R) (*R, *resty.Response, error) {
	rateLimiterBucket.Wait(1)
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(result).
		SetError(&Error).
		SetBody(body).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with put operation: %s", path)
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

// func convertToString(i interface{}) string {
// 	switch v := i.(type) {
// 	case float64:
// 		return strconv.Itoa(int(v))
// 	case string:
// 		return v
// 	default:
// 		return ""
// 	}
// }

func validateUserAgent(ua string) bool {
	regex := regexp.MustCompile(`^\S+ \S+ \S+$`)
	return regex.MatchString(ua)
}

func Paginate(fn any, id string, secondId string, params interface{}) (interface{}, *resty.Response, error) {
	var result []interface{}
	var result2 interface{}
	var response *resty.Response
	var err error

	if id != "" && secondId == "" {
		if f, ok := fn.(func(string, interface{}) (interface{}, *resty.Response, error)); ok {
			a := params
			for {
				result2, response, err = f(id, a)
				if err != nil {
					return nil, response, err
				}
				result = append(result, result2)
				startingAfter, err := getStartingAfter(*response)
				if err == nil {
					a = changeParams(a, startingAfter)
				}

				if err != nil {
					break
				}
			}
		}
	} else if id != "" && secondId != "" {

		if f, ok := fn.(func(string, string, interface{}) (interface{}, *resty.Response, error)); ok {
			a := params
			for {
				result2, response, err = f(id, secondId, a)
				if err != nil {
					return nil, response, err
				}
				result = append(result, result2)
				startingAfter, err := getStartingAfter(*response)
				if err == nil {
					a = changeParams(a, startingAfter)
				}

				if err != nil {
					break
				}
			}
		}
	} else if id == "" && secondId == "" {
		if f, ok := fn.(func(interface{}) (interface{}, *resty.Response, error)); ok {
			a := params
			for {
				result2, response, err = f(a)
				if err != nil {
					return nil, response, err
				}
				result = append(result, result2)
				startingAfter, err := getStartingAfter(*response)
				if err == nil {
					a = changeParams(a, startingAfter)
				}

				if err != nil {
					break
				}
			}
		}
	} else {
		return nil, nil, fmt.Errorf("invalid function type")
	}
	return result, response, err
}

func changeParams(params interface{}, newValue string) interface{} {
	valueA := reflect.ValueOf(params)
	var perPage int64
	perPage = PAGINATION_PER_PAGE
	valueA = valueA.Elem()
	newParams := reflect.New(valueA.Type()).Elem()
	newParams.FieldByName("PerPage").SetInt(perPage)
	newParams.FieldByName(("StartingAfter")).SetString(newValue)
	// copiar los demas valores en newParams
	for i := 0; i < valueA.NumField(); i++ {
		if valueA.Type().Field(i).Name != "PerPage" && valueA.Type().Field(i).Name != "StartingAfter" {
			newParams.Field(i).Set(valueA.Field(i))
		}
	}
	//

	fmt.Println("New Params: ", newParams)
	// Devolver el nuevo objeto como interface{}
	return newParams.Addr().Interface()
}

func getStartingAfter(r2 resty.Response) (string, error) {
	links := strings.Split(r2.Header().Get("Link"), ",")
	var link string

	for linkIndex := range links {
		if linkIndex != 0 {
			if strings.Contains(links[linkIndex], "startingAfter") {
				link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[linkIndex], ";")[0], ">", ""), "<", ""), MERAKI_BASE_URL, "")
				break
			}
		}
	}
	// link =

	link = strings.TrimSpace(link)

	myUrl, _ := url.Parse(link)
	params, _ := url.ParseQuery(myUrl.RawQuery)
	if params["startingAfter"] == nil {
		return "", fmt.Errorf("StartingAfter absent")
	}
	return params["startingAfter"][0], nil
}
