package openaq

import (
	"fmt"
	"strings"
)

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
	ErrorType ErrorType
	Message   string
}

func (err ClientError) Error() string {
	var sb strings.Builder
	sb.Grow(128)

	fmt.Fprint(&sb, "OpenAQ client error ", err.ErrorType, ": ")

	fmt.Fprintf(&sb, "%q", err.Message)

	return sb.String()
}

func (err ClientError) String() string {
	return err.Error()
}

type APIError struct {
	// Error message.
	Message string `json:"message"`
	// HTTP code.
	Code int
}

func (err APIError) Error() string {
	var sb strings.Builder
	sb.Grow(128)

	fmt.Fprint(&sb, "OpenAQ API error (status code ", err.Code, "): ")

	fmt.Fprintf(&sb, "%q", err.Message)

	return sb.String()
}

func (err APIError) String() string {
	return err.Error()
}
