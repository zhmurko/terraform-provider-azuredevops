package azuredevops

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProvider_HasChildResources(t *testing.T) {
	expectedResources := []string{
		"azuredevops_resource_authorization",
		"azuredevops_build_definition",
		"azuredevops_branch_policy_build_validation",
		"azuredevops_branch_policy_min_reviewers",
		"azuredevops_branch_policy_auto_reviewers",
		"azuredevops_branch_policy_work_item_linking",
		"azuredevops_branch_policy_comment_resolution",
		"azuredevops_project",
		"azuredevops_project_features",
		"azuredevops_serviceendpoint_github",
		"azuredevops_serviceendpoint_dockerregistry",
		"azuredevops_serviceendpoint_azurerm",
		"azuredevops_serviceendpoint_azurecr",
		"azuredevops_serviceendpoint_runpipeline",
		"azuredevops_serviceendpoint_bitbucket",
		"azuredevops_serviceendpoint_kubernetes",
		"azuredevops_serviceendpoint_aws",
		"azuredevops_variable_group",
		"azuredevops_git_repository",
		"azuredevops_user_entitlement",
		"azuredevops_group_membership",
		"azuredevops_group",
		"azuredevops_agent_pool",
		"azuredevops_agent_queue",
		"azuredevops_project_permissions",
		"azuredevops_git_permissions",
		"azuredevops_workitemquery_permissions",
		"azuredevops_area_permissions",
		"azuredevops_iteration_permissions",
	}

	resources := Provider().ResourcesMap
	require.Equal(t, len(expectedResources), len(resources), "There are an unexpected number of registered resources")

	for _, resource := range expectedResources {
		require.Contains(t, resources, resource, "An expected resource was not registered")
		require.NotNil(t, resources[resource], "A resource cannot have a nil schema")
	}
}

func TestProvider_HasChildDataSources(t *testing.T) {
	expectedDataSources := []string{
		"azuredevops_client_config",
		"azuredevops_group",
		"azuredevops_project",
		"azuredevops_projects",
		"azuredevops_git_repositories",
		"azuredevops_git_repository",
		"azuredevops_users",
		"azuredevops_agent_pool",
		"azuredevops_agent_pools",
		"azuredevops_agent_queue",
		"azuredevops_area",
		"azuredevops_iteration",
	}

	dataSources := Provider().DataSourcesMap
	require.Equal(t, len(expectedDataSources), len(dataSources), "There are an unexpected number of registered data sources")

	for _, resource := range expectedDataSources {
		require.Contains(t, dataSources, resource, "An expected data source was not registered")
		require.NotNil(t, dataSources[resource], "A data source cannot have a nil schema")
	}
}

func TestProvider_SchemaIsValid(t *testing.T) {
	type testParams struct {
		name          string
		required      bool
		defaultEnvVar string
		sensitive     bool
	}

	tests := []testParams{
		{"org_service_url", false, "AZDO_ORG_SERVICE_URL", false},
		{"personal_access_token", false, "AZDO_PERSONAL_ACCESS_TOKEN", true},
	}

	schema := Provider().Schema
	require.Equal(t, len(tests), len(schema), "There are an unexpected number of properties in the schema")

	for _, test := range tests {
		require.Contains(t, schema, test.name, "An expected property was not found in the schema")
		require.NotNil(t, schema[test.name], "A property in the schema cannot have a nil value")
		require.Equal(t, test.sensitive, schema[test.name].Sensitive, "A property in the schema has an incorrect sensitivity value")
		require.Equal(t, test.required, schema[test.name].Required, "A property in the schema has an incorrect required value")

		if test.defaultEnvVar != "" {
			expectedValue := os.Getenv(test.defaultEnvVar)

			actualValue, err := schema[test.name].DefaultFunc()
			if actualValue == nil {
				actualValue = ""
			}

			require.Nil(t, err, "An error occurred when getting the default value from the environment")
			require.Equal(t, expectedValue, actualValue, "The default value pulled from the environment has the wrong value")
		}
	}
}
