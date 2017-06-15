# List Nofifiers #

The `list notifiers` endpoint shows notifiers currently loaded into the application.

## URL Pattern ##

### GET - /api/v1/notifiers ###

## Input ##

### Request Headers ###

### Request Body ###

## Response ##

### Response Body ###

```
[
    {
        #JOB OBJECT
    },
]
```

### Response Headers ###

* **x-node-id**: "Hostname of the container receiving the hook"

### Response Codes ###

* **200**: Job received successfully
* **500**: Error occurred during job receipt
