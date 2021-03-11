package gatypes

import (
	"fmt"
	"net/url"
)

// StringList is used to represent repeated, string values in a query string.
type StringList []string

// EncodeValues converts the list of values to a key => value map.
func (l StringList) EncodeValues(key string, v *url.Values) error {
	for i, value := range l {
		k := fmt.Sprintf("%s%d", key, i+1)
		v.Set(k, value)
	}
	return nil
}

// Float64List is used to represent repeated, float values in a query string.
type Float64List []float64

// EncodeValues converts the list of values to a {key}/{index+1} => value map.
func (l Float64List) EncodeValues(key string, v *url.Values) error {
	for i, value := range l {
		k := fmt.Sprintf("%s%d", key, i+1)
		v.Set(k, fmt.Sprintf("%f", value))
	}
	return nil
}
