package openaq

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/hashicorp/go-cleanhttp"
)

const (
	defaultBaseURLScheme = "https"
	defaultBaseURLHost   = "api.openaq.org"
	defaultBasePath      = "/v3/"
	defaultUserAgent     = "openaq-go"
)

type Client struct {
	baseURL     *url.URL
	apiKey      string
	userAgent   string
	httpHeaders map[string]string
	client      *http.Client
}

// Config contains client configuration.
type Config struct {
	// BaseURLScheme is the url scheme to use defaults to https
	baseURLScheme string
	// BaseURLHost is the base url to use defualts to api.openaq.org
	baseURLHost string
	// APIKey is an optional API key.
	apiKey string
	// userAgent is an optional HTTP header.
	userAgent string
	// HTTPHeaders are additional optional HTTP headers.
	httpHeaders map[string]string
	// Client provides an optional HTTP client, otherwise a default will be used.
	client *http.Client
}

// New creates a new OpenAQ client.
func NewClient(config Config) (*Client, error) {
	client := config.client
	if client == nil {
		client = cleanhttp.DefaultClient()
	}
	var baseURLScheme string
	if config.baseURLScheme == "" {
		baseURLScheme = defaultBaseURLScheme
	} else {
		baseURLScheme = config.baseURLScheme
	}
	var baseURLHost string
	if config.baseURLHost == "" {
		baseURLHost = defaultBaseURLHost
	} else {
		baseURLHost = config.baseURLHost

	}
	var userAgent string
	if config.userAgent == "" {
		userAgent = defaultUserAgent
	} else {
		userAgent = config.userAgent
	}

	return &Client{
		baseURL: &url.URL{
			Scheme: baseURLScheme,
			Host:   baseURLHost,
			Path:   defaultBasePath,
		},
		apiKey:      config.apiKey,
		userAgent:   userAgent,
		httpHeaders: config.httpHeaders,
		client:      client,
	}, nil
}

func (c *Client) request(ctx context.Context, method, requestPath string, query url.Values, responseStruct interface{}) error {
	var (
		req          *http.Request
		resp         *http.Response
		err          error
		bodyContents []byte
	)

	req, err = c.newRequest(requestPath, query)

	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	resp, err = c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyContents, err = io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("status: %d, body: %v", resp.StatusCode, string(bodyContents))
	}

	if responseStruct == nil {
		return nil
	}

	err = json.Unmarshal(bodyContents, responseStruct)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) newRequest(requestPath string, query url.Values) (*http.Request, error) {
	url := c.baseURL
	url.Path = path.Join(url.Path, requestPath)
	url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return req, err
	}

	if c.apiKey != "" {
		req.Header.Add("X-API-key", c.apiKey)
	}
	if c.httpHeaders != nil {
		for k, v := range c.httpHeaders {
			req.Header.Add(k, v)
		}
	}

	req.Header.Add("Content-Type", "application/json")
	return req, err
}
