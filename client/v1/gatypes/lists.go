package gatypes

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
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

// DecodeValues will pull array data out of an existing values structure.
func (l *StringList) DecodeValues(key string, v *url.Values) error {
	values := make([]string, len(*v))
	c := 0

	for k, value := range *v {
		if strings.HasPrefix(k, key) {
			i, err := strconv.Atoi(strings.TrimPrefix(k, key))
			if err == nil {
				values[i-1] = value[0]
				c++
			}
		}
	}

	*l = values[:c]
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

// DecodeValues will pull array data out of an existing values structure.
func (l *Float64List) DecodeValues(key string, v *url.Values) error {
	values := make([]float64, len(*v))
	c := 0

	for k, value := range *v {
		if strings.HasPrefix(k, key) {
			i, err := strconv.Atoi(strings.TrimPrefix(k, key))
			if err == nil {
				values[i-1], _ = strconv.ParseFloat(value[0], 64)
				c++
			}
		}
	}

	*l = values[:c]
	return nil
}
