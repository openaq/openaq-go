package openaq

import "fmt"

type ErrorType int

const (
	MissingParamError ErrorType = iota
)

func (e ErrorType) String() string {
	switch e {
	case MissingParamError:
		return "MissingParamError"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type ClientError struct {
	errorType ErrorType
	message   string
}

func (c *ClientError) Error() string {
	return fmt.Sprintf("%s: %v", c.errorType, c.message)
}
