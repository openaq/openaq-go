# OpenAQ Golang API Client

[![Project Status: Suspended – Initial development has started, but there has not yet been a stable, usable release; work has been stopped for the time being but the author(s) intend on resuming work.](https://www.repostatus.org/badges/latest/suspended.svg)](https://www.repostatus.org/#suspended)

Develop on this library has been suspended. Usage is not recommended.

A low-level golang wrapper around the OpenAQ v3 REST API.


### Usage



#### Client

A client is initialized with a configuration struct which holds some global configuration for interacting with the API. 

```
client := NewClient(Config{
    APIKey: "my-openaq-api-key-123456-7890"
})
```

All methods take a `context.Context` as the first parameter to allow for context level cancellation and deadlines.
