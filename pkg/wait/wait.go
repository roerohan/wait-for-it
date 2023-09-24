package wait

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Services is a string array storing the services that are to be waited for
type Services []Service

// Service is a struct meant to denote a service hostname:port with a wait condition upon start up
type Service struct {
	hostname string
	port     int
}

var (
	// ErrServiceMaxTimeout is the error message to use in case a max service startup timeout is exceeded.
	ErrServiceMaxTimeout = fmt.Errorf("max service startup timeout duration exceeded waiting for service dependencies")
)

// String prints out the human-readable Service hostname:port string.
func (s *Service) String() string {
	return fmt.Sprintf("%s:%d", s.hostname, s.port)
}

// NewService creates a wait.Service type.
func NewService(hostname string, port int) Service {
	return Service{
		hostname: hostname,
		port:     port,
	}
}

// Set is used to append a Service to the slice of Services,
// to implement the interface flag.Value
func (s *Services) Set(value string) error {
	const separator = ":"

	// Note: serviceInfo[0] = hostname, serviceInfo[1] = port
	serviceInfo := strings.Split(value, separator)
	port, err := strconv.Atoi(serviceInfo[1])
	if err != nil {
		return err
	}

	service := NewService(serviceInfo[0], port)
	*s = append(*s, service)
	return nil
}

// String returns a string representation of the flag,
// to implement the interface flag.Value
func (s *Services) String() string {
	var sb strings.Builder
	const formatter string = ", "
	for _, service := range *s {
		sb.WriteString(service.String())
		sb.WriteString(formatter)
	}

	// trim the last comma that was added for last service
	return strings.TrimSuffix(sb.String(), formatter)
}

// ForDependencies allows the service to wait for its dependencies to be up and ready for a configurable amount of time.
// If the service dependency request timeout is reached and the dependent services are not yet available,
// then the timeout wait interval will continue until the dependencies are up for a maximum wait time of maxTimeout.
func ForDependencies(waitServices Services, serviceRequestTimeout, maxTimeout time.Duration) error {
	if len(waitServices) == 0 {
		return nil
	}

	success := make(chan bool, 1)
	ok := wait(waitServices, serviceRequestTimeout)
	if ok {
		success <- true
	}

	// return err if service wait time exceeds ServiceMaxTimeout time
	select {
	case <-success:
		return nil
	case <-time.After(maxTimeout * time.Second):
		return ErrServiceMaxTimeout
	}
}

func wait(waitServices Services, waitTimeOut time.Duration) bool {
	now := time.Now()

	var wg sync.WaitGroup
	wg.Add(len(waitServices))

	success := make(chan bool, 1)

	go func() {
		for _, service := range waitServices {
			go waitOne(service, &wg, now)
		}
		wg.Wait()
		success <- true
	}()

	select {
	case <-success:
		return true
	case <-time.After(waitTimeOut * time.Second):
		return false
	}

}

func waitOne(service Service, wg *sync.WaitGroup, start time.Time) {
	defer wg.Done()
	for {
		_, err := net.Dial("tcp", service.String())
		if err == nil {
			Log(fmt.Sprintf("%s is available after %s", service.String(), time.Since(start)))
			break
		}
		opErr, ok := err.(*net.OpError)
		if ok && errors.Is(err, opErr) {
			Log(fmt.Sprintf("failed to dial service %s with err: %s", service.String(), opErr.Error()))
			break
		}
		time.Sleep(time.Second)
	}
}
