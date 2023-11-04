package customerrors

import "errors"

type RouteError struct {
	Err        error
	StatusCode int
}

func (err RouteError) Error() string {
	return err.Err.Error()
}

func NewRouteError(statusCode int, message string) *RouteError {
	return &RouteError{
		Err:        errors.New(message),
		StatusCode: statusCode,
	}
}
