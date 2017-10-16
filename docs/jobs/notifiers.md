# Notifiers #

## Notifier Types ##

### Webhook ###

Make a request out to a webhook (internal or external).  This process will be synchronous.  Request body and headers can be specified.

#### Properties ####

* **body** *string*: HTTP Payload - Supports Golang Templating
* **headers** *string*: Key / Value pairs for headers to submit with the request - Values support Golang Templating
* **method** *string*: HTTP Method
* **url** *string*: URL to target for the outgoing hook - Supports Golang Templating
* **disableSSLVerification** *bool*: Set to true to disable remote SSL verification

## Event Triggers ##

Notifiers can be triggered by jobs, steps, and tasks.  Configuration options for each are available on their documentation pages.
