package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/go-uuid"

	"github.com/mjpitz/go-ga/client/v1"
	"github.com/mjpitz/go-ga/client/v1/gatypes"
)

// variables set during build using -X ldflag
var version string

func main() {
	if version == "" {
		version = "latest"
	}
	userAgent := "beacon/" + version

	trackingID := os.Getenv("TRACKING_ID")
	userID := os.Getenv("USER_ID")

	applicationName := os.Getenv("APPLICATION_NAME")
	applicationVersion := os.Getenv("APPLICATION_VERSION")

	pingIntervalEnv := os.Getenv("PING_INTERVAL")
	if pingIntervalEnv == "" {
		pingIntervalEnv = "5m"
	}

	pingInterval, err := time.ParseDuration(pingIntervalEnv)
	if err != nil {
		log.Fatalf("failed to parse PING_INTERVAL: %v\n", err)
	}

	clientID := ""
	if userID == "" {
		clientID, err = uuid.GenerateUUID()
		if err != nil {
			log.Fatalf("failed to generate clientID: %v\n", err)
		}
	}

	eventLabel := fmt.Sprintf("%s/%s", applicationName, applicationVersion)

	client := v1.NewClient(trackingID, userAgent)
	users := gatypes.Users{
		UserID:   userID,
		ClientID: clientID,
	}

	log.Printf("starting beacon")
	err = client.SendPost(&gatypes.Payload{
		DisableAdvertisingPersonalization: true,
		HitType:                           "event",
		Users:                             users,
		SessionControl: gatypes.SessionControl{
			SessionControl: "start",
		},
		Event: gatypes.Event{
			EventCategory: "beacon",
			EventAction:   "startup",
			EventLabel:    eventLabel,
		},
	})
	if err != nil {
		log.Println(err)
	}

	defer client.SendPost(&gatypes.Payload{
		DisableAdvertisingPersonalization: true,
		HitType:                           "event",
		Users:                             users,
		SessionControl: gatypes.SessionControl{
			SessionControl: "end",
		},
		Event: gatypes.Event{
			EventCategory: "beacon",
			EventAction:   "shutdown",
			EventLabel:    eventLabel,
		},
	})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	for {
		_ = client.SendPost(&gatypes.Payload{
			DisableAdvertisingPersonalization: true,
			HitType:                           "event",
			Users:                             users,
			Event: gatypes.Event{
				EventCategory: "beacon",
				EventAction:   "heartbeat",
				EventLabel:    eventLabel,
			},
		})

		select {
		case <-signals:
			log.Printf("received shutdown signal, stopping\n")
			return
		case <-time.Tick(pingInterval):
			continue
		}
	}
}
