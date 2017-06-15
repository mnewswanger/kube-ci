# Healthz #

The `healtz` endpoint is set up to provide a health check endpoint for the service.

## URL Pattern ##

### GET /healthz/ ###

## Input ##

### Request Headers ###

*There are no configured request headers*

### Request Body ###


*There is no expected request body*

## Response ##

### Response Body ###

'KubeCI is Up'

### Response Headers ###

*There are no configured response headers*

### Response Codes ###

* **200**: Service is ready
* **500**: Service is in a failed state
* **503**: Service is not yet ready
