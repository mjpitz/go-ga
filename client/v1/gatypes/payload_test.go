package gatypes_test

import (
	"testing"

	"github.com/google/go-querystring/query"

	"github.com/mjpitz/go-ga/client/v1/gatypes"

	"github.com/stretchr/testify/require"
)

func Test_PayloadFull(t *testing.T) {
	p := &gatypes.Payload{
		Version:                           "1",
		TrackingID:                        "UA-1234567-1",
		AnonymizeIP:                       true,
		DisableAdvertisingPersonalization: true,
		DataSource:                        "test",
		QueueTime:                         100,
		CacheBuster:                       "123456",
		HitType:                           "event",
		NonInteractionHit:                 true,
		Event: gatypes.Event{
			EventCategory: "test",
			EventAction:   "test",
		},
	}

	v, err := query.Values(p)
	require.Nil(t, err)
	require.Equal(t, "aip=1&ds=test&ea=test&ec=test&ni=1&npa=1&qt=100&t=event&tid=UA-1234567-1&v=1&z=123456", v.Encode())
}
