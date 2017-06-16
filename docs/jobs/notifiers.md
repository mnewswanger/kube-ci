# Notifiers #

## Notifier Types ##

### Webhook ###

Make a request out to a webhook (internal or external).  This process will be synchronous.  Request body and headers can be specified.

#### Properties ####

* `body`: HTTP Payload
* `headers`: Key / Value pairs for headers to submit with the request
* `method`: HTTP Method
* `mime_type`: Mime type for the request
* `url`: URL to target for the outgoing hook
* `verify_ssl`: Set to true to verify the target SSL key

## Event Triggers ##

Notifiers can be triggered by jobs, steps, and tasks.  Configuration options for each are available on their documentation pages.
