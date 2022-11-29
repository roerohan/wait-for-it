package wait

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Services is a string array storing
// the services that are to be waited for
type Services []Service

// Service is a string meant to denote a service with a wait condition upon start up
type Service string

func (s *Service) String() string {
	return string(*s)
}

// Set is used to append a Service
// to the slice of Services, to implement
// the interface flag.Value
func (s *Services) Set(value string) error {
	service := interface{}(value).(Service)
	*s = append(*s, service)
	return nil
}

// String returns a string
// representation of the flag,
// to implement the interface
// flag.Value
func (s *Services) String() string {
	var sb strings.Builder
	const formatter string = ", "
	for _, service := range *s {
		sb.WriteString(service.String())
		sb.WriteString(formatter)
	}
	return sb.String()
}

// Wait waits for all services
func (s *Services) Wait(tSeconds int) bool {
	t := time.Duration(tSeconds) * time.Second
	now := time.Now()

	var wg sync.WaitGroup
	wg.Add(len(*s))

	success := make(chan bool, 1)

	go func() {
		for _, service := range *s {
			go waitOne(service, &wg, now)
		}
		wg.Wait()
		success <- true
	}()

	select {
	case <-success:
		return true
	case <-time.After(t):
		return false
	}

}

func waitOne(service Service, wg *sync.WaitGroup, start time.Time) {
	defer wg.Done()
	for {
		_, err := net.Dial("tcp", service.String())
		if err == nil {
			Log(fmt.Sprintf("%s is available after %s", service, time.Since(start)))
		}
		break
	}
	time.Sleep(time.Second)
}
