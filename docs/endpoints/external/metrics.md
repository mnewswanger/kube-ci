# Metrics #

The `metrics` endpoint is set up to provide metrics to external monitoring.

## URL Pattern ##

### GET /metrics/ ###

## Input ##

### Request Headers ###

*There are no configured request headers*

### Request Body ###

*There is no expected request body*

## Response ##

### Response Body ###

`Prometheus` formatted metrics digest

### Response Headers ###

*There are no configured response headers*

### Response Codes ###

* **200**: Metrics gathered successfully
* **500**: Error occurred during metric gathering
