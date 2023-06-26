/*
Meraki Dashboard API

Testing FloorPlansApiService

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

func Test_client_FloorPlansApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test FloorPlansApiService CreateNetworkFloorPlan", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.FloorPlansApi.CreateNetworkFloorPlan(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FloorPlansApiService DeleteNetworkFloorPlan", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var floorPlanId string

        resp, httpRes, err := apiClient.FloorPlansApi.DeleteNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FloorPlansApiService GetNetworkFloorPlan", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var floorPlanId string

        resp, httpRes, err := apiClient.FloorPlansApi.GetNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FloorPlansApiService GetNetworkFloorPlans", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.FloorPlansApi.GetNetworkFloorPlans(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FloorPlansApiService UpdateNetworkFloorPlan", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var floorPlanId string

        resp, httpRes, err := apiClient.FloorPlansApi.UpdateNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
