package gatypes

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// Boolean defines a Google Analytics boolean type.
type Boolean bool

// EncodeValues sets a 1 or 0 depending on the value of the boolean.
func (b Boolean) EncodeValues(key string, v *url.Values) error {
	if b {
		v.Set(key, "1")
	} else {
		v.Set(key, "0")
	}
	return nil
}

// DecodeValues will pull boolean data out of an existing values structure.
func (b *Boolean) DecodeValues(key string, v *url.Values) error {
	*b = v.Get(key) == "1"
	return nil
}

var _ query.Encoder = Boolean(true)
