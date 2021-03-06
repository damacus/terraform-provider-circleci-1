package circleci

import (
	"github.com/stretchr/testify/require"
	"testing"
)
func TestProvider_HasChildResources(t *testing.T) {
	expectedResources := []string{
		"circleci_project",
	}

	resources := Provider().ResourcesMap
	require.Equal(t, len(expectedResources), len(resources), "There are an unexpected number of registered resources")

	for _, resource := range expectedResources {
		require.Contains(t, resources, resource, "An expected resource was not registered")
		require.NotNil(t, resources[resource], "A resource cannot have a nil schema")
	}
}

func TestProvider_SchemaIsValid(t *testing.T) {
	type testParams struct {
		token          string
		organization   string
	}

	tests := []testParams{
		{"myTestToken", "terraform-providers-circleci"},
	}

	schema := Provider().Schema
	require.Equal(t, len(tests), len(schema), "There are an unexpected number of properties in the schema")

	for _, test := range tests {
		require.NotNil(t, schema[test.organization], "A property in the schema cannot have a nil value")
		require.NotNil(t, schema[test.organization], "A property in the schema cannot have a nil value")
	}
}