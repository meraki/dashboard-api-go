/*
Meraki Dashboard API

Testing FeaturesApiService

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

func Test_client_FeaturesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test FeaturesApiService CreateOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.FeaturesApi.CreateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeaturesApiService DeleteOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var optInId string

        resp, httpRes, err := apiClient.FeaturesApi.DeleteOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeaturesApiService GetOrganizationEarlyAccessFeatures", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.FeaturesApi.GetOrganizationEarlyAccessFeatures(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeaturesApiService GetOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var optInId string

        resp, httpRes, err := apiClient.FeaturesApi.GetOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeaturesApiService GetOrganizationEarlyAccessFeaturesOptIns", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.FeaturesApi.GetOrganizationEarlyAccessFeaturesOptIns(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeaturesApiService UpdateOrganizationEarlyAccessFeaturesOptIn", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var optInId string

        resp, httpRes, err := apiClient.FeaturesApi.UpdateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
