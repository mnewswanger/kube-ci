# Job Callback #

The `job callback` endpoint is set up as a receive hook to be triggered when the client completes an action requested by KubeCI.

## URL Pattern ##

### POST /api/v1/callback/<callbackToken\>/ ###

## Input ##

### Request Headers ###

*There are no configured request headers*

### Request Body ###

```json
{
    "error": "String with error message if the job failed"
}
```

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

* **200**: Callback processed successfully
* **404**: Callback received but could not be located
* **500**: Error occurred during callback receipt
