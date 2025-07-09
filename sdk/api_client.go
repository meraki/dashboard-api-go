package meraki

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/juju/ratelimit"
)

const (
	MERAKI_BASE_URL             = "MERAKI_BASE_URL"
	MERAKI_USER_AGENT           = "MERAKI_USER_AGENT"
	MERAKI_DASHBOARD_API_KEY    = "MERAKI_DASHBOARD_API_KEY"
	MERAKI_DEBUG                = "MERAKI_DEBUG"
	MERAKI_REQUESTS_PER_SECOND  = "MERAKI_REQUESTS_PER_SECOND"
	MERAKI_RETRIES              = "MERAKI_RETRIES"
	MERAKI_RETRY_DELAY          = "MERAKI_RETRY_DELAY"
	MERAKI_RETRY_JITTER         = "MERAKI_RETRY_JITTER"
	MERAKI_USE_RETRY_HEADER     = "MERAKI_USE_RETRY_HEADER"
	DEFAULT_USER_AGENT          = "go-meraki/1.57.0"
	DEFAULT_REQUESTS_PER_SECOND = 10
	PAGINATION_PER_PAGE         = 1000
	DEFAULT_RETRIES             = 3
	DEFAULT_RETRY_DELAY         = 1000 * time.Millisecond
	DEFAULT_RETRY_JITTER        = 3000 * time.Millisecond
)

// Client manages communication with the Cisco Meraki API

type Backoff struct {
	MaxRetries     *int
	MaxRetryDelay  *time.Duration
	MaxRetryJitter *time.Duration
	UseRetryHeader *bool
}

func (backoff *Backoff) initBackoff(maxRetries *int, maxRetryDelay *time.Duration, maxRetryJitter *time.Duration, useRetryHeader *bool) {
	defaultMaxRetries := DEFAULT_RETRIES - 1
	defaultMaxRetryDelay := DEFAULT_RETRY_DELAY
	defaultMaxRetryJitter := DEFAULT_RETRY_JITTER
	defaultUseRetryHeader := false
	backoff.MaxRetries = func() *int {
		if maxRetries != nil {
			finalMaxRetries := *maxRetries
			if *maxRetries > 0 {
				finalMaxRetries = *maxRetries - 1
			}
			return &finalMaxRetries
		}
		return &defaultMaxRetries
	}()

	backoff.MaxRetryDelay = func() *time.Duration {
		if maxRetryDelay != nil {
			return maxRetryDelay
		}
		return &defaultMaxRetryDelay
	}()

	backoff.MaxRetryJitter = func() *time.Duration {
		if maxRetryJitter != nil {
			return maxRetryJitter
		}
		return &defaultMaxRetryJitter
	}()
	backoff.UseRetryHeader = func() *bool {
		if useRetryHeader != nil {
			return useRetryHeader
		}
		return &defaultUseRetryHeader
	}()
}

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
	backoff           *Backoff
}

func GET[Q any, H any](path string, client *resty.Client, params *Q, header *H) (*resty.Response, error) {
	queryString, _ := query.Values(params)

	if queryString.Get("perPage") == "-1" {
		queryString.Set("perPage", strconv.Itoa(PAGINATION_PER_PAGE))
	}

	return client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		Get(path)
}

func POST(path string, client *resty.Client, body any, params any) (*resty.Response, error) {
	queryString, _ := query.Values(params)
	response := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode())
	if body != nil {
		response.SetBody(body)
	}
	return response.Post(path)
}

func PUT(path string, client *resty.Client, body any) (*resty.Response, error) {
	response := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")
	if body != nil {
		response.SetBody(body)
	}
	return response.Put(path)
}

func DELETE[Q any](path string, client *resty.Client, params *Q) (*resty.Response, error) {
	queryString, _ := query.Values(params)

	return client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		Delete(path)
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

var HeaderDefault interface{} = nil
var QueryParamsDefault interface{} = nil

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
	// client.SetRetryCount(2)
	// client.SetRetryWaitTime(time.Second)
	// client.SetRetryMaxWaitTime(1 * time.Minute)
	// // Respect "Retry-After" header in response
	// client.SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
	// 	retryHeader := resp.Header().Get("Retry-After")
	// 	if _, err := strconv.Atoi(retryHeader); err == nil {
	// 		t, _ := time.ParseDuration(retryHeader + "s")
	// 		return t, nil
	// 	}
	// 	return 0, nil
	// })
	// // Retry on 429
	// client.AddRetryCondition(
	// 	func(r *resty.Response, err error) bool {
	// 		return err != nil || r.StatusCode() == http.StatusTooManyRequests
	// 	},
	// )

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

func (c *Client) SetBackoff(maxRetries *int, maxRetryDelay *time.Duration, maxRetryJitter *time.Duration, useRetryHeader *bool) error {
	backoff := Backoff{}
	backoff.initBackoff(maxRetries, maxRetryDelay, maxRetryJitter, useRetryHeader)
	c.common.backoff = &backoff
	return nil
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

	// if !validateUserAgent(userAgent) {
	// 	return errors.New("user-agent bad format, expected: `AplicationName VendorName Client`")
	// }

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

// func validateUserAgent(ua string) bool {
// 	regex := regexp.MustCompile(`^\S+ \S+ \S+$`)
// 	return regex.MatchString(ua)
// }

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

// Helper function to parse Link header and get next URL, this is used in cases where the responses are not a list and we need to paginate
// the results using the Link header in the response
func getNextPageURL(linkHeader string) string {
	if linkHeader == "" {
		return ""
	}

	links := strings.Split(linkHeader, ",")
	for _, link := range links {
		parts := strings.Split(strings.TrimSpace(link), ";")
		if len(parts) < 2 {
			continue
		}

		urlPart := strings.TrimSpace(parts[0])
		if !strings.HasPrefix(urlPart, "<") || !strings.HasSuffix(urlPart, ">") {
			continue
		}

		url := urlPart[1 : len(urlPart)-1]

		relPart := strings.TrimSpace(parts[1])
		if strings.Contains(relPart, "rel=\"next\"") || strings.Contains(relPart, "rel=next") {
			return url
		}
	}

	return ""
}

func doWithRetriesAndResult[T any](
	operation func() (*resty.Response, error),
	client *resty.Client,
	backoff *Backoff,
	merge func(dst, src T) T,
	paginate ...bool,
) (*T, *resty.Response, error) {
	var result *T
	var resp *resty.Response
	var err error
	maxRetries, maxRetryDelay, maxRetryJitter, useRetryHeader := getBackoffValues(backoff)

	for attempt := 0; attempt <= maxRetries; attempt++ {
		fmt.Println("MAX_RETRIES: ", maxRetries+1)
		resp, err = operation()

		if err != nil && resp.StatusCode() != http.StatusTooManyRequests {

			log.Printf("Error 1: %v", err)
			log.Printf("Response: %v", resp.StatusCode())
			return nil, resp, err
		}
		if resp.IsError() && resp.StatusCode() != http.StatusTooManyRequests {
			log.Printf("Error 2: %v", resp)
			log.Printf("Status code: %v", resp.StatusCode())
			return nil, resp, fmt.Errorf("error with operation: %s Error:\n %s", resp.Request.URL, resp)
		}
		if resp != nil && resp.StatusCode() != http.StatusTooManyRequests {
			json.Unmarshal(resp.Body(), &result)
			pages := false
			if len(paginate) > 0 {
				pages = paginate[0]
			} else {
				perPage, _ := strconv.Atoi(resp.Request.QueryParam.Get("perPage"))
				if perPage == -1 {
					pages = true // si el perPage es -1, entonces se debe paginar
				}
			}
			if result != nil && !pages {
				return result, resp, nil
			}

			// Pagination
			if result != nil && pages {
				link := getNextPageURL(resp.Header().Get("Link"))
				var result2 *T
				for link != "" {
					// Add jitter backoff attempt
					for attempt := 0; attempt <= maxRetries; attempt++ {
						resp, err = GET(link, client, &QueryParamsDefault, &HeaderDefault)
						if err == nil {
							break
						}
						if resp.IsError() {
							log.Printf("Error 3: %v", resp)
							if resp.StatusCode() != http.StatusTooManyRequests {
								return result, resp, fmt.Errorf("error with get operation: %s", link)
							}
						}
						delay := maxRetryDelay * time.Duration(1<<attempt)
						jitter := time.Duration(rand.Int63n(int64(maxRetryJitter)))
						wait := delay + jitter
						log.Printf("[retry] attempt %d: received 429, waiting %v", attempt, wait)
						time.Sleep(wait)
					}
					json.Unmarshal(resp.Body(), &result2)

					*result = merge(*result, *result2)

					link = getNextPageURL(resp.Header().Get("Link"))
				}
				if resp.IsError() {
					return nil, resp, fmt.Errorf("error with get operation: %s", link)
				}
				return result, resp, nil
			}

			return nil, resp, nil
		}

		// Exponential backoff with jitter
		// Calculate exponential backoff delay by multiplying baseDelay by 2^attempt
		// For example: if baseDelay is 1s and attempt is 2, delay will be 1s * 2^2 = 4s
		delay := maxRetryDelay * time.Duration(1<<attempt)
		// Generate a random duration between 0 and jitterMax to add randomness to retry delays
		jitter := time.Duration(rand.Int63n(int64(maxRetryJitter)))
		wait := delay + jitter

		if resp != nil && useRetryHeader {
			if retryAfter := resp.Header().Get("Retry-After"); retryAfter != "" {
				if secs, err := strconv.Atoi(retryAfter); err == nil {
					wait = time.Duration(secs) * time.Second
				}
			}
		}

		log.Printf("[retry] attempt %d: received 429, waiting %v", attempt+1, wait)
		time.Sleep(wait)
	}

	// If result is nil, return zero value
	log.Printf("error 4: %v", err)
	if result != nil {

		return result, resp, fmt.Errorf("failed after %d retries", maxRetries+1)
	}

	return nil, resp, fmt.Errorf("failed after %d retries", maxRetries+1)
}

func getBackoffValues(backoff *Backoff) (int, time.Duration, time.Duration, bool) {
	maxRetries, maxRetryDelay, maxRetryJitter, useRetryHeader := DEFAULT_RETRIES, DEFAULT_RETRY_DELAY, DEFAULT_RETRY_JITTER, false
	if backoff == nil {
		return maxRetries, maxRetryDelay, maxRetryJitter, useRetryHeader
	}

	if backoff.MaxRetries != nil {
		maxRetries = *backoff.MaxRetries
	}
	if backoff.MaxRetryDelay != nil {
		maxRetryDelay = *backoff.MaxRetryDelay
	}
	if backoff.MaxRetryJitter != nil {
		maxRetryJitter = *backoff.MaxRetryJitter
	}
	if backoff.UseRetryHeader != nil {
		useRetryHeader = *backoff.UseRetryHeader
	}
	return maxRetries, maxRetryDelay, maxRetryJitter, useRetryHeader
}

func doWithRetriesAndNotResult(
	operation func() (*resty.Response, error),
	backoff *Backoff,
) (*resty.Response, error) {
	var resp *resty.Response
	var err error
	maxRetries, maxRetryDelay, maxRetryJitter, useRetryHeader := getBackoffValues(backoff)
	fmt.Println("MAX_RETRIES: ", maxRetries+1)
	for attempt := 0; attempt <= maxRetries; attempt++ {
		resp, err = operation()

		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				// continue to retry
			} else {
				return resp, err
			}
		}

		if resp != nil && resp.StatusCode() != http.StatusTooManyRequests {
			return resp, nil
		}

		delay := maxRetryDelay * time.Duration(1<<attempt)
		jitter := time.Duration(rand.Int63n(int64(maxRetryJitter)))
		wait := delay + jitter

		if resp != nil && useRetryHeader {
			if retryAfter := resp.Header().Get("Retry-After"); retryAfter != "" {
				if secs, err := strconv.Atoi(retryAfter); err == nil {
					wait = time.Duration(secs) * time.Second
				}
			}
		}

		log.Printf("[retry] attempt %d: received 429, waiting %v", attempt+1, wait)
		time.Sleep(wait)
	}

	return resp, fmt.Errorf("failed after %d retries", maxRetries+1)
}
