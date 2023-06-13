package openaq

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestNewClient(t *testing.T) {
	apiKey := "123456789abcdefghijklmno"

	config := &Config{
		apiKey:        apiKey,
		baseURLScheme: "http",
		baseURLHost:   "staging.openaq.org",
		userAgent:     "test",
	}

	c, err := NewClient(*config)
	if err != nil {
		fmt.Println("uh oh")
	}

	if got, want := c.apiKey, apiKey; got != want {
		t.Fatalf("c.apiKey = %q, want %q", got, want)
	}

	if got, want := c.baseURL.Scheme, "http"; got != want {
		t.Fatalf("c.baseURL.Scheme = %q, want %q", got, want)
	}

	if got, want := c.userAgent, "test"; got != want {
		t.Fatalf("c.userAgent = %q, want %q", got, want)
	}

}
