# Squeeze Go Client

This project implements an HTTP client for the Squeeze API.

**Maturity**: Experimental. We have already created multiple HTTP clients internally but never a fully featured
client which supported all API endpoints of Squeeze. We will try to generate an API client based on the OpenAPI
specification once the [OpenAPI generator](https://openapi-generator.tech/) is able to generate a fully working
client.

**Support**: Currently, we will not officially support this client. Still, if you want to use the client and have
problems or questions, feel free to open an issue. We would love to encourage usage of this client and the HTTP
api of our products in general.

## Usage

Add client as dependency:

```go
go get github.com/dexpro-solutions-gmbh/squeeze-go-client
```

## Tests

To run tests on this project, use these ENV variables to configure the Squeeze API to be used when testing:

- SQZ_BASE_PATH: The base path of the Squeeze API. Example: `https://your.squeeze.one/api/v2`
- SQZ_KEY: API key to authenticate with
