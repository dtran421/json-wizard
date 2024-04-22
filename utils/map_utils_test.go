package utils_test

import (
	"reflect"
	"testing"

	"github.com/dtran421/json-wizard/utils"
)

func TestSortedMap(t *testing.T) {
	cases := []struct {
		in       map[string]interface{}
		expected []utils.KeyValuePair
	}{
		{
			in: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			expected: []utils.KeyValuePair{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
				{Key: "key3", Value: "value3"},
			},
		},
		{
			in: map[string]interface{}{
				"key3": "value3",
				"key2": "value2",
				"key1": "value1",
			},
			expected: []utils.KeyValuePair{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
				{Key: "key3", Value: "value3"},
			},
		},
	}

	for _, c := range cases {
		got := utils.SortedMap(c.in)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("SortedMap(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}
