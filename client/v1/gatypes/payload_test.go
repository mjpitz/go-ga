package gatypes_test

import (
	"encoding/json"
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
		CustomDimensions: []string{
			"dimension-1",
			"dimension-2",
		},
		CustomMetrics: []float64{
			1.1,
			2.2,
			3.3,
		},
	}

	{
		v, err := query.Values(p)
		require.Nil(t, err)

		expected := "aip=1&cd1=dimension-1&cd2=dimension-2&cm1=1.100000&cm2=2.200000&cm3=3.300000&ds=test&ea=test&ec=test&ni=1&npa=1&qt=100&t=event&tid=UA-1234567-1&v=1&z=123456"
		require.Equal(t, expected, v.Encode())
	}

	{
		data, err := json.Marshal(p)
		require.Nil(t, err)

		expected := `{"version":"1","tracking_id":"UA-1234567-1","anonymize_ip":true,"disable_ad_personalization":true,"data_source":"test","queue_time":100,"cache_buster":"123456","hit_type":"event","non_interaction_hit":true,"custom_dimensions":["dimension-1","dimension-2"],"custom_metrics":[1.1,2.2,3.3],"event_category":"test","event_action":"test"}`
		require.Equal(t, expected, string(data))
	}
}
