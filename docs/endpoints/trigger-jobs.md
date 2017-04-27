# Trigger Jobs #

The `trigger jobs` endpoint is set up as the base entrypoint for job interaction with KubeCI.

## URL Pattern ##

### POST - / ###

## Input ##

### Request Headers ###

### Request Body ###

KubeCI supports the webhook format for the following external services:

* `Gitlab`: v9

For custom calls, the following format should be used:


```
{
    "labels": [
        "namespaced.label-1": "Label Value"
        "namespaced.label-2": "Another Value"
    ]
}
```

## Response ##

### Response Body ###

```
{
    "error": "Descriptive error message if something went wrong"
}
```

### Response Headers ###

* **x-node-id**: "Hostname of the container receiving the hook"

### Response Codes ###

* **200**: Job received successfully
* **500**: Error occurred during job receipt
