# Steps #

A job can have multiple steps assigned to it.  Steps are run serially (each must complete prior to starting the next job).

## Tasks ##

Each step is made up of one or more tasks.  Tasks, unlike steps, are run in parallel to each other.  A task follows rules to determine when to run.  Tasks can do multiple things:

### Call A Webhook ###

Make an outbound webhook call.  This can be internal or external to KubeCI.  When calling the user has the option to specify synchronous or asynchronous work.  When making a synchronous call, the KubeCI process waits for a reponse from the target.  When making an asynchronous call, the KubeCI process provides a callback to the target to notify completion to the system.  Both connection timeout and request timeout can be specified, at which point, a retry will be scheduled.

### Run A Kubernetes Job ###

Run a job on Kubernetes infrastructure.  This process will trigger a job and provide a notification endpoint to the process indicate job completion.  The user can also specify a timeout, and upon completion, the KubeCI process will aggregate the job's logs and clean the job from the cluser.

### Update A Kubernetes Deployment ###

Update properties of a Kubernetes deployment.  This will likely be an image update to update a code deployment, but it can also target scaling and environment updates.

## Events ##

The following events are associated to `steps`.  They can be applied to notifiers and event_handlers.

* `complete` - The step completed - *Only applicable to notifiers*
* `failed` - The step failed
* `started` - The step was started - *Only applicable to notifiers*
* `succeeded` - The step succeeded


The following events are associated to `tasks`.  They can be applied to notifiers and event_handlers.

* `attempt_failed` - The task attempt failed but may be retried - *Only applicable to notifiers*
* `complete` - The task completed - *Only applicable to notifiers*
* `failed` - The task failed and will not be retried
* `started` - The task was started - *Only applicable to notifiers*
* `succeeded` - The task succeeded

_Jobs and steps do not have an `attempt_failed` event similar to tasks since the job itself does not get retried; tasks do._

## Event Handler Actions ##

When configuring event handlers for `steps`, the following actions can be configured:

* `set_job_failure` - Mark the job as failed
* `set_job_success` - Mark the job as completed successfully

When configuring event handlers for `tasks`, the following actions can be configured:

* `set_step_failure` - Mark the step as failed
* `set_step_success` - Mark the step as completed successfully
