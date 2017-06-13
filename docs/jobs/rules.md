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

* `git.branch` - Branch being targeted by the operation
* `git.event` - Event type (`push`, `tag_push`)
* `git.project.avatar_url` - Project image
* `git.project.name` - Project name
* `git.project.namespace` - Project namespace
* `git.repository.url_http` - Repository path (HTTP)
* `git.repository.url_ssh` - Repository path (SSH)
* `git.start_commit` - Initial commit as a starting point to the action
* `git.tag` - Tag name that was pushed
* `git.target_commit` - Commit the action is targeting
* `git.user.avatar_url` - Author avatar URL
* `git.user.email` - Author email address
* `git.user.name` - Author username
