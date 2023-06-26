/*
Meraki Dashboard API

Testing NetworkAdaptersApiService

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

func Test_client_NetworkAdaptersApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test NetworkAdaptersApiService GetNetworkSmDeviceNetworkAdapters", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var deviceId string

        resp, httpRes, err := apiClient.NetworkAdaptersApi.GetNetworkSmDeviceNetworkAdapters(context.Background(), networkId, deviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
