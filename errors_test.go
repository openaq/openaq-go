package openaq

import (
	"net/http"
	"testing"
)

func TestErrorString(t *testing.T) {
	apiError := APIError{
		Code:    http.StatusNotFound,
		Message: `Not Found`,
	}
	expected := `OpenAQ API error (status code 404): "Not Found"`
	if got, want := apiError.Error(), expected; got != want {
		t.Fatalf("err.Error() = %q, want %q", got, want)
	}
	if got, want := apiError.String(), expected; got != want {
		t.Fatalf("err.String() = %q, want %q", got, want)
	}
	clientError := ClientError{
		ErrorType: MissingParamError,
		Message:   `Missing required parameter`,
	}
	expected = `OpenAQ client error MissingParamError: "Missing required parameter"`
	if got, want := clientError.Error(), expected; got != want {
		t.Fatalf("err.Error() = %q, want %q", got, want)
	}
	if got, want := clientError.String(), expected; got != want {
		t.Fatalf("err.String() = %q, want %q", got, want)
	}
}
