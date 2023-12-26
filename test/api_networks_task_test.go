/*
VirtyAPI

Testing NetworksTaskAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package virty-go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_virty-go_NetworksTaskAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworksTaskAPIService CreateNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworksTaskAPI.CreateNetwork(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworksTaskAPIService CreateNetworkOvs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.NetworksTaskAPI.CreateNetworkOvs(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworksTaskAPIService DeleteNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.NetworksTaskAPI.DeleteNetwork(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworksTaskAPIService DeleteNetworkOvs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string
		var name string

		resp, httpRes, err := apiClient.NetworksTaskAPI.DeleteNetworkOvs(context.Background(), uuid, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworksTaskAPIService PostUuidOvsApiTasksNetworksProvidersPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworksTaskAPI.PostUuidOvsApiTasksNetworksProvidersPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworksTaskAPIService RefreshNetworks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworksTaskAPI.RefreshNetworks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}