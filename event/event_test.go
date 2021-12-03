package event_test

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/event"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {

	event.RegisterHook(event.Start, func() {
		fmt.Println("register start")
	})
	event.RegisterHook(event.LocalConfigComplete, func() {
		fmt.Println("register LocalConfigComplete 1")
	})
	go func() {
		event.RegisterHook(event.ConfigComplete, func() {
			fmt.Println("register ConfigComplete")
		})
	}()
	go func() {
		event.TriggerEvent(event.LocalConfigComplete)
	}()
	time.Sleep(1 * time.Second)
	event.RegisterHook(event.LocalConfigComplete, func() {
		fmt.Println("register LocalConfigComplete 2")
	})

	time.Sleep(1 * time.Second)

}
