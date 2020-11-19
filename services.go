package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Services is a string array storing
// the services that are to be waited for
type Services []string

// Set is used to append a string
// to the service, to implement
// the interface flag.Value
func (s *Services) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// String returns a string
// representation of the flag,
// to implement the interface
// flag.Value
func (s *Services) String() string {
	return strings.Join(*s, ", ")
}
