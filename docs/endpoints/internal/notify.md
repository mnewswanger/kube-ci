# Notify #

The `notify` endpoint tells the receiving node to execute the given `notifier`.  The job it relates to is passed as a second URL chunk.

## URL Pattern ##

### POST - /api/v1/notify/<notifierId\>/<jobId\>/ ###

## Input ##

### Request Headers ###

*There are no expected request headers*

### Request Body ###

*There is no expected request body*

## Response ##

### Response Body ###

```json
{
    "error": "Descriptive error message if something went wrong"
}
```

### Response Headers ###

* **x-node-id**: "Hostname of the container receiving the hook"

### Response Codes ###

* **200**: Job received successfully
* **500**: Error occurred during job receipt
