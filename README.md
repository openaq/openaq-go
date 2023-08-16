# OpenAQ Golang API Client

This library is in early development. **DO NOT USE UNTIL 1.0.0 RELEASE**

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