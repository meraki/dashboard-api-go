/*
Meraki Dashboard API

Testing QualityAndRetentionApiService

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

func Test_client_QualityAndRetentionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test QualityAndRetentionApiService GetDeviceCameraQualityAndRetention", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.QualityAndRetentionApi.GetDeviceCameraQualityAndRetention(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test QualityAndRetentionApiService UpdateDeviceCameraQualityAndRetention", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.QualityAndRetentionApi.UpdateDeviceCameraQualityAndRetention(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}