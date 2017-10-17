# List Notifiers #

The `list notifiers` endpoint shows notifiers currently loaded into the application.

## URL Pattern ##

### GET - /api/v1/notifiers ###

## Input ##

### Request Headers ###

### Request Body ###

## Response ##

### Response Body ###

```
{
    "slack.debug": {
        "name": "debug",
        "namespace": "slack",
        "properties": {
            "body": "{\"text\":\"{{.Job.Event}}: {{.Job.Namespace}} - {{.Job.Name}}\"}",
            "headers": {
                "Content-Type": "application/json"
            },
            "method": "POST",
            "url": "<Slack URL>"
        },
        "retries": 0,
        "type": "webhook"
    }
}
```

### Response Headers ###

* **x-node-id**: "Hostname of the container receiving the hook"

### Response Codes ###

* **200**: Job received successfully
* **500**: Error occurred during job receipt
