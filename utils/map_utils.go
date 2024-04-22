package utils

import "sort"

type KeyValuePair struct {
	Key   string
	Value interface{}
}

func SortedMap(m map[string]interface{}) []KeyValuePair {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	keyValuePairs := make([]KeyValuePair, 0, len(keys))
	for _, key := range keys {
		keyValuePairs = append(keyValuePairs, KeyValuePair{Key: key, Value: m[key]})
	}

	return keyValuePairs
}
