# Jobs #

## Overview ##

The core functionality of `Kube CI` is driven by `Jobs`.

## Job Composition ##

A job consists of three main components: `Rules`, `Steps`, and `Notifications`.  These allow for separation of components and distribution of tasks.

### Rules ###

Rules define when a job should be run.  These are applied by labels - either passed in via webhooks or computed from external system hooks.  Rules can be applied at the job level, step level, or task level.

See the `Rules` documentation for details.

[Rules](rules)

### Steps ###

Steps define what a job should do when it runs.  Each job contains one or more steps.

See the `Steps` documentation for details.

[Steps](steps)

### Notifiers ###

Notifiers can be triggered by events at a job, step, or task level, and each can be triggered over completion, success, or error conditions.

See the `Notifiers` documentation for details.

[Notifiers](notifiers)

## Events ##

The following events are associated to `jobs`.  They can be applied to notifiers.

* `complete` - The job completed
* `failed` - The job failed
* `started` - The job was started
* `succeeded` - The job succeeded

_Jobs and steps do not have an `attempt_failed` event similar to tasks since the job itself does not get retried; tasks do._
