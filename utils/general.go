package utils

// MakeMap creates map to be used in the json response
func MakeMap(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{key: value}
}
