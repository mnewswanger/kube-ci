# Endpoints #

## External Endpoints ##

### / ###

Trigger a job.  Jobs are triggered based off of `labels` provided by the web request or computed by the inbound hooks with enabled providers.  See [jobs](/jobs/) section for details.

### /callback/<callbackToken\>/ ###

Provide an endpoint for callbacks from asynchronous jobs.  The `jobInstanceId` and `callbackToken` will are provided as variables to be provided to the webhook consumer.

### /documentation/ ###

View the documentation for the service.

### /metrics/ ###

Provide `prometheus` metrics endpoints.

## Internal Endpoints ##

### /execute-job/<jobId\>/ ###

Instruct a KubeCI server to run a specific job.  An inbound hook can trigger multiple jobs.  Each job gets sent in via a webhook endpoint and can be run against different servers.

### /notify/<notificationId\>/ ###

Fire a notification.  See [notifiers](/jobs/notifiers) for more details on configurations and callbacks.
