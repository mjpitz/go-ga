# go-ga

Send data to Google Analytics from Go.

```go
package main

import (
    "log"
    "time"

	"github.com/mjpitz/go-ga/client/v1"
    "github.com/mjpitz/go-ga/client/v1/gatypes"
)

func main() {
    client := v1.NewClient("UA-XXXXXX-1")
    ping := &gatypes.Payload{
        HitType: "event",
        NonInteractionHit: true,
        DisableAdvertisingPersonalization: true,
        Event: &gatypes.Event{
            EventCategory: "beacon",
            EventAction: "heartbeat",
        },
    }
    
    for {
        if err := client.SendPost(ping); err != nil {
            log.Fatal(err)
        }
        time.Sleep(time.Minute)
    }
}
```

Or generate a URL you can embed in an email:

```go
package main

import (
    "github.com/mjpitz/go-ga/client/v1"
    "github.com/mjpitz/go-ga/client/v1/gatypes"
)

func main() {
    client := v1.NewClient("UA-XXXXXX-1")
    ping := &gatypes.Payload{
        HitType: "event",
        NonInteractionHit: true,
        DisableAdvertisingPersonalization: true,
        Event: &gatypes.Event{
            EventCategory: "email",
            EventAction: "open",
        },
    }

    url, err := client.GetURL(ping)
    // do something with url
}
```
