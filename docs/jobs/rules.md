# Rules #

Rules are configured to determine whether or not a job should be triggered.

Rules are triggered based on `labels`.

## Applying Rules ##

Rules can be applied via the following methods:

* `exact` - Provided string matches the label
* `exists` - Passes if the provided label exists
* `regex` - Provided regular expression matches the label

Each rule also supports inversion.

Multiple rules can be combined in both AND / OR combinations.  Rules can be applied to `jobs`, `steps`, and `tasks`.

## Generated Labels ##

### Gitlab Webhooks ###

When consuming webhooks from Gitlab, the following labels will be available:

* `gitlab.project_name` - Name of the project
* `gitlab.action` - Action being applied to the repository
* `gitlab.branch` - Target branch of the update
* `gitlab.commit` - Commit ID
* `gitlab.url`- Project URL
