package v1

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Decoder provides an interface for decoding values.
type Decoder interface {
	DecodeValues(key string, v *url.Values) error
}

var decoderType = reflect.TypeOf(new(Decoder)).Elem()

// Decode converts a URL's values back into the target structure.
func Decode(values url.Values, v interface{}) error {
	if v == nil {
		return nil
	}

	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("v1: Decode() expects struct input. Got %v", val.Kind())
	}

	return decode(&values, val, "")
}

func decode(values *url.Values, val reflect.Value, key string) error {
	v := reflect.New(val.Type())
	if v.Type().Implements(decoderType) {
		err := v.Interface().(Decoder).DecodeValues(key, values)
		if err != nil {
			return err
		}
		val.Set(v.Elem())
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		key := fieldType.Tag.Get("url")
		parts := strings.Split(key, ",")

		if err := set(values, field, parts[0]); err != nil {
			return err
		}
	}

	return nil
}

func set(values *url.Values, val reflect.Value, key string) error {
	value := values.Get(key)

	switch val.Kind() {
	case reflect.Bool:
		if value != "" {
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			val.SetBool(b)
		}
	case reflect.String:
		if value != "" {
			val.SetString(value)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value != "" {
			i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			val.SetInt(i)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if value != "" {
			i, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			val.SetUint(i)
		}
	case reflect.Float32, reflect.Float64:
		if value != "" {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			val.SetFloat(f)
		}
	default:
		return decode(values, val, key)
	}

	return nil
}
