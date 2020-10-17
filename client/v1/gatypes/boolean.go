package gatypes

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type Boolean bool

func (b Boolean) EncodeValues(key string, v *url.Values) error {
	if b {
		v.Set(key, "1")
	} else {
		v.Set(key, "0")
	}
	return nil
}

var _ query.Encoder = Boolean(true)
