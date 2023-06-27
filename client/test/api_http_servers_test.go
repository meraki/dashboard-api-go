/*
Meraki Dashboard API

Testing HttpServersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func Test_client_HttpServersApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test HttpServersApiService CreateNetworkWebhooksHttpServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.HttpServersApi.CreateNetworkWebhooksHttpServer(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test HttpServersApiService DeleteNetworkWebhooksHttpServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var httpServerId string

        resp, httpRes, err := apiClient.HttpServersApi.DeleteNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test HttpServersApiService GetNetworkWebhooksHttpServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var httpServerId string

        resp, httpRes, err := apiClient.HttpServersApi.GetNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test HttpServersApiService GetNetworkWebhooksHttpServers", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.HttpServersApi.GetNetworkWebhooksHttpServers(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test HttpServersApiService UpdateNetworkWebhooksHttpServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var httpServerId string

        resp, httpRes, err := apiClient.HttpServersApi.UpdateNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
