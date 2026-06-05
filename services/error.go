package services

import "fmt"

type ServiceInitError struct {
	ServiceName string
}

func (err *ServiceInitError) Error() string {
	return fmt.Sprintf("Failed to initialise %q", err.ServiceName)
}
