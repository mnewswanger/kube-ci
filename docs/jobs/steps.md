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

Events can be applied to `jobs`, `steps`, and `tasks`.

### Job Level Events ###

The job does not have a `jobTryFailed` similar to the properties below since the job itself does not get retried; steps do.

* `jobFailed` - The job failed
* `jobSucceeded` - The job succeeded

### Step Level Events ###

* `stepTryFailed` - The step failed
* `stepFailed` - The step failed, and no retries remain
* `stepSucceeded` - The step succeeded

### Task Level Events ###

* `taskTryFailed` - The task failed
* `taskFailed` - The task failed, and no retries remain
* `taskSucceeded` - The task succeeded
