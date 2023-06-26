/*
Meraki Dashboard API

Testing FirewalledServicesApiService

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

func Test_client_FirewalledServicesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test FirewalledServicesApiService GetNetworkApplianceFirewallFirewalledService", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var service string

        resp, httpRes, err := apiClient.FirewalledServicesApi.GetNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FirewalledServicesApiService GetNetworkApplianceFirewallFirewalledServices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.FirewalledServicesApi.GetNetworkApplianceFirewallFirewalledServices(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FirewalledServicesApiService UpdateNetworkApplianceFirewallFirewalledService", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var service string

        resp, httpRes, err := apiClient.FirewalledServicesApi.UpdateNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
