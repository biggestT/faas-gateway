## FaaS gateway

Discover services and maps incoming requests to appropriate services if available. Balances traffic different hosts available in a round-robin fashion.

## Service discovery

The gateway discovers appearances and disappearances of containers within the same docker network tagged with these labels:

- `faas.app=true`
- `faas.name=SERVICE_NAME`
- `faas.port=PORT`

## Environment variables

- `POLL_FREQ`: frequency of service discovery poll
- `DOCKER_API_VERSION`: should be set to the docker API version of the hosts docker server (`docker version`)
- `CORS_ORIGIN`: what origin should be allowed to make CORS request
