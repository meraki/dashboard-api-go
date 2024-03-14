package meraki

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const (
	TEST_MERAKI_BASE_URL            = "https://api.meraki.com"
	TEST_MERAKI_DASHBOARD_API_KEY   = "0123456789"
	TEST_MERAKI_DEBUG               = "false"
	TEST_MERAKI_SSL_VERIFY          = "true"
	TEST_MERAKI_REQUESTS_PER_SECOND = 5
)

type result struct{}

// testClient returns a new Client for testing.
func testClient(t *testing.T) *Client {
	err := SetOptions(TEST_MERAKI_BASE_URL, TEST_MERAKI_DASHBOARD_API_KEY, TEST_MERAKI_DEBUG)
	assert.Equal(t, err, nil)
	client, err := NewClient()
	assert.Equal(t, err, nil)
	assert.NotEqual(t, client, nil)
	client.common.client.SetRetryCount(0)
	gock.InterceptClient(client.common.client.GetClient())
	return client
}

// TestSetOptions tests the Client.SetOptions method.
func TestSetOptions(t *testing.T) {
	err := SetOptions(TEST_MERAKI_BASE_URL, TEST_MERAKI_DASHBOARD_API_KEY, TEST_MERAKI_DEBUG)
	assert.Equal(t, err, nil)
	assert.Equal(t, os.Getenv(MERAKI_BASE_URL), TEST_MERAKI_BASE_URL)
	assert.Equal(t, os.Getenv(MERAKI_DASHBOARD_API_KEY), TEST_MERAKI_DASHBOARD_API_KEY)
	assert.Equal(t, os.Getenv(MERAKI_DEBUG), TEST_MERAKI_DEBUG)
}

// TestSetOptions tests the Client.SetOptionsWithRequests method.
func TestSetOptionsWithRequests(t *testing.T) {
	err := SetOptionsWithRequests(TEST_MERAKI_BASE_URL, TEST_MERAKI_DASHBOARD_API_KEY, TEST_MERAKI_DEBUG, TEST_MERAKI_REQUESTS_PER_SECOND)
	assert.Equal(t, err, nil)
	assert.Equal(t, os.Getenv(MERAKI_BASE_URL), TEST_MERAKI_BASE_URL)
	assert.Equal(t, os.Getenv(MERAKI_DASHBOARD_API_KEY), TEST_MERAKI_DASHBOARD_API_KEY)
	assert.Equal(t, os.Getenv(MERAKI_DEBUG), TEST_MERAKI_DEBUG)
	assert.Equal(t, os.Getenv(MERAKI_REQUESTS_PER_SECOND), strconv.Itoa(TEST_MERAKI_REQUESTS_PER_SECOND))
}

// TestNewClient tests the Client.NewClient method.
func TestNewClient(t *testing.T) {
	client := testClient(t)
	assert.Equal(t, client.common.client.HostURL, TEST_MERAKI_BASE_URL)
	debug, _ := strconv.ParseBool(TEST_MERAKI_DEBUG)
	assert.Equal(t, client.common.client.Debug, debug)
	assert.Equal(t, client.common.client.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", TEST_MERAKI_DASHBOARD_API_KEY))
	assert.Equal(t, client.common.client.Header.Get("User-Agent"), DEFAULT_USER_AGENT)
}

// TestNewClientWithOptions tests the Client.NewClientWithOptions method.
func TestNewClientWithOptions(t *testing.T) {
	url := "url"
	key := "key"
	debug := false
	client, err := NewClientWithOptions(url, key, strconv.FormatBool(debug))
	assert.Equal(t, err, nil)
	assert.NotEqual(t, client, nil)
	assert.Equal(t, client.common.client.HostURL, url)
	assert.Equal(t, client.common.client.Debug, debug)
	assert.Equal(t, client.common.client.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", key))
}

// TestNewClientWithOptions tests the Client.NewClientWithOptions method.
func TestNewClientWithOptionsAndRequests(t *testing.T) {
	url := "url"
	key := "key"
	debug := false
	requests := 6
	client, err := NewClientWithOptionsAndRequests(url, key, strconv.FormatBool(debug), requests)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, client, nil)
	assert.Equal(t, client.common.client.HostURL, url)
	assert.Equal(t, client.common.client.Debug, debug)
	assert.Equal(t, client.common.client.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", key))
	assert.Equal(t, client.common.rateLimiterBucket.Capacity(), int64(requests))
	assert.Equal(t, client.common.rateLimiterBucket.Rate(), float64(requests))
}

// TestSetAuthToken tests the Client.SetAuthToken method.
func TestSetAuthToken(t *testing.T) {
	newToken := "new"
	client := testClient(t)
	client.SetAuthToken(newToken)
	assert.Equal(t, client.common.client.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", newToken))
}

// TestSetUserAgent tests the Client.SetUserAgent method.
func TestSetUserAgent(t *testing.T) {
	newUserAgent := "user2"
	client := testClient(t)
	client.SetUserAgent(newUserAgent)
	assert.Equal(t, client.common.client.Header.Get("User-Agent"), newUserAgent)
}

// TestSetRequestsPerSecond tests the Client.SetRequestsPerSecond method.
func TestSetRequestsPerSecond(t *testing.T) {
	requestsPerSecond := 10
	client := testClient(t)
	client.SetRequestsPerSecond(requestsPerSecond)
	assert.Equal(t, client.common.rateLimiterBucket.Capacity(), int64(requestsPerSecond))
	assert.Equal(t, client.common.rateLimiterBucket.Rate(), float64(requestsPerSecond))
}

// TestRestyClient tests the Client.RestyClient method.
func TestRestyClient(t *testing.T) {
	client := testClient(t)
	assert.Equal(t, client.RestyClient(), client.common.client)
	assert.IsType(t, *(client.RestyClient()), resty.Client{})
}

// TestGet tests the Client.Get method.
func TestGet(t *testing.T) {
	defer gock.Off()
	client := testClient(t)

	// Success
	gock.New(TEST_MERAKI_BASE_URL).Get("/url").Reply(200)
	_, response, err := Get[result](client.common.client, client.common.rateLimiterBucket, "/url", &result{})
	assert.NoError(t, err)
	assert.True(t, response.IsSuccess())

	// HTTP error
	gock.New(TEST_MERAKI_BASE_URL).Get("/url").ReplyError(errors.New("fail"))
	_, _, err = Get[result](client.common.client, client.common.rateLimiterBucket, "/url", &result{})
	assert.Error(t, err)

	// HTTP status code
	gock.New(TEST_MERAKI_BASE_URL).Get("/url").Reply(405)
	_, response, _ = Get[result](client.common.client, client.common.rateLimiterBucket, "/url", &result{})
	assert.Equal(t, response.StatusCode(), 405)
}

// TestPost tests the Client.Post method.
func TestPost(t *testing.T) {
	defer gock.Off()
	client := testClient(t)
	type body struct{ b string }

	// Success
	gock.New(TEST_MERAKI_BASE_URL).Post("/url").Reply(200)
	_, response, err := Post[body, result](client.common.client, client.common.rateLimiterBucket, "/url", &body{b: "body"}, &result{})
	assert.NoError(t, err)
	assert.True(t, response.IsSuccess())

	// HTTP error
	gock.New(TEST_MERAKI_BASE_URL).Post("/url").ReplyError(errors.New("fail"))
	_, _, err = Post[body, result](client.common.client, client.common.rateLimiterBucket, "/url", &body{b: "body"}, &result{})
	assert.Error(t, err)

	// HTTP status code
	gock.New(TEST_MERAKI_BASE_URL).Post("/url").Reply(405)
	_, response, _ = Post[body, result](client.common.client, client.common.rateLimiterBucket, "/url", &body{b: "body"}, &result{})
	assert.Equal(t, response.StatusCode(), 405)
}

// TestPut tests the Client.Put method.
func TestPut(t *testing.T) {
	defer gock.Off()
	client := testClient(t)
	type body struct{ b string }

	// Success
	gock.New(TEST_MERAKI_BASE_URL).Put("/url").Reply(200)
	_, response, err := Put[body, result](client.common.client, client.common.rateLimiterBucket, "/url", &body{b: "body"}, &result{})
	assert.NoError(t, err)
	assert.True(t, response.IsSuccess())

	// HTTP error
	gock.New(TEST_MERAKI_BASE_URL).Put("/url").ReplyError(errors.New("fail"))
	_, _, err = Put[body, result](client.common.client, client.common.rateLimiterBucket, "/url", &body{b: "body"}, &result{})
	assert.Error(t, err)

	// HTTP status code
	gock.New(TEST_MERAKI_BASE_URL).Put("/url").Reply(405)
	_, response, _ = Put[body, result](client.common.client, client.common.rateLimiterBucket, "/url", &body{b: "body"}, &result{})
	assert.Equal(t, response.StatusCode(), 405)
}

// TestDelete tests the Client.Delete method.
func TestDelete(t *testing.T) {
	defer gock.Off()
	client := testClient(t)

	// Success
	gock.New(TEST_MERAKI_BASE_URL).Delete("/url").Reply(200)
	response, err := Delete(client.common.client, client.common.rateLimiterBucket, "/url")
	assert.NoError(t, err)
	assert.True(t, response.IsSuccess())

	// HTTP error
	gock.New(TEST_MERAKI_BASE_URL).Delete("/url").ReplyError(errors.New("fail"))
	_, err = Delete(client.common.client, client.common.rateLimiterBucket, "/url")
	assert.Error(t, err)

	// HTTP status code
	gock.New(TEST_MERAKI_BASE_URL).Delete("/url").Reply(405)
	response, _ = Delete(client.common.client, client.common.rateLimiterBucket, "/url")
	assert.Equal(t, response.StatusCode(), 405)
}

// TestRateLimit tests the rate limiting of the Client.Get method.
func TestRateLimit(t *testing.T) {
	defer gock.Off()
	client := testClient(t)

	gock.New(TEST_MERAKI_BASE_URL).Get("/url").Reply(200)
	start := time.Now()
	for i := 0; i < 11; i++ {
		Get[result](client.common.client, client.common.rateLimiterBucket, "/url", &result{})
	}
	assert.LessOrEqual(t, time.Second, time.Since(start))
}

// TestRetry tests a retry of the Client.Get method.
func TestRetry(t *testing.T) {
	defer gock.Off()
	client := testClient(t)
	client.common.client.SetRetryCount(1)
	client.common.client.SetRetryWaitTime(100 * time.Millisecond)

	gock.New(TEST_MERAKI_BASE_URL).Get("/url").Reply(429)
	gock.New(TEST_MERAKI_BASE_URL).Get("/url").Reply(200)
	start := time.Now()
	Get[result](client.common.client, client.common.rateLimiterBucket, "/url", &result{})
	assert.LessOrEqual(t, 100*time.Millisecond, time.Since(start))
}

// TestRetryAfter tests a retry of the Client.Get method with a Retry-After response header.
func TestRetryAfter(t *testing.T) {
	defer gock.Off()
	client := testClient(t)
	client.common.client.SetRetryCount(1)

	gock.New(TEST_MERAKI_BASE_URL).Get("/url").Reply(429).AddHeader("Retry-After", "1")
	gock.New(TEST_MERAKI_BASE_URL).Get("/url").Reply(200)
	start := time.Now()
	Get[result](client.common.client, client.common.rateLimiterBucket, "/url", &result{})
	assert.LessOrEqual(t, time.Second, time.Since(start))
}
