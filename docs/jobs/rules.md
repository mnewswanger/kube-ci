# Rules #

Rules are configured to determine whether or not a job should be triggered.

Rules are triggered based on `labels`.

## Applying Rules ##

Rules can be applied via the following methods:

* `match` - Provided string matches the label
* `regexMatch` - Provided regular expression matches the label
* `noMatch` - Provided string does not match the label
* `regexNoMatch` - Provided regular expression does not match the label

Mutliple rules can be combined in both AND / OR combinations.  Rules can be applied to `jobs`, `steps`, and `tasks`.

## Generated Labels ##

### Gitlab Webhooks ###

When consuming webhooks from Gitlab, the following labels will be available:

* `gitlab.project_name` - Name of the project
* `gitlab.branch` - Target branch of the update
* `gitlab.commit` - Commit ID
* `gitlab.project_name`- Project name
