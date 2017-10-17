# List Jobs #

The `list jobs` endpoint shows notifiers currently loaded into the application.

## URL Pattern ##

### GET - /api/v1/jobs ###

## Input ##

### Request Headers ###

### Request Body ###

## Response ##

### Response Body ###

```json
{
    "generic.Git Commit Push Notification": {
        "name": "Git Commit Push Notification",
        "namespace": "generic",
        "notifiers": [
            {
                "name": "Send Slack Notification",
                "properties": {
                    "key": "value"
                },
                "events": ["success"],
                "notifier": "slack.notifications"
            }
        ],
        "rules": {
            "name": "Push Event",
            "mode": "all",
            "rules": [
                {
                    "name": "git_push",
                    "invert_match": false,
                    "label_name": "git.event",
                    "label_value": "push",
                    "match_mode": "exact"
                }
            ]
        },
        "steps": null
    }
}
```

### Response Headers ###

* **x-node-id**: "Hostname of the container receiving the hook"

### Response Codes ###

* **200**: Job received successfully
* **500**: Error occurred during job receipt
