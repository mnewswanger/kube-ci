# Jobs #

```
---
name: Git Tag Push Example
namespace: examples
labels:
  label1: value1
  label2: value2
steps:
  - name: Do something
    tasks:
    - name: Task 1
      action: webhook
      arguments:
        url: https://www.google.com
      retries: 0
      timeout: 5
      verify_ssl: true
  - name: Do something else
    tasks:
    - name: Task 2
      action: webhook
      arguments:
        url: https://www.google.com
      retries: 0
      timeout: 5
      verify_ssl: true
    - name: Task 3
      action: webhook
      arguments:
        url: https://www.google.com
      retries: 0
      timeout: 5
      verify_ssl: true
notifiers:
  complete:
    - name: Send Slack notification
      notifier: examples.generic-webhook
      arguments:
        key_1: value
        key_2: value
      events: complete
rules:
  mode: all
  name: git_tag_push
  rules:
    - invert_match: false
      label_name: git.event
      label_value: tag_push
      match_mode: exact
      name: Tag push event
...
```

```
---
name: Git Commit Push Example
namespace: examples
labels:
  label1: value1
  label2: value2
steps:
  - name: Do something
    tasks:
    - name: Task 1
      action: webhook
      arguments:
        url: https://www.google.com
      retries: 0
      timeout: 5
  - name: Do something else
    tasks:
    - name: Concurrent Task 1
      action: kubernetes_job
      arguments:
        job_tempate: 1
      event_handlers:
        failure: continue
        success: mark_step_complete
      retries: 1
      timeout: 120
    - name: Concurrent Task 2
      action: kubernetes_deployment
      arguments:
        deployment_name: 1
      retries: 3
      timeout: 3600
notifiers:
  success:
  - name: Send RocketChat notification
    notifier: generic.rocketchat
    arguments:
      channel: "site-reliability-engineering"
    events: success
rules:
  mode: all
  name: push_to_dev_or_master_on_any_project
  rules:
    - invert_match: false
      label_name: git.event
      label_value: push
      match_mode: exact
      name: git_push
    - invert_match: true
      label_name: git.project.name
      label_value: ""
      match_mode: exact
      name: git_request
    - invert_match: true
      label_name: git.target_commit
      label_value: ""
      match_mode: exact
      name: gitlab_project_foo
    - mode: any
      name: Branch master or develop
      rules:
        - invert_match: false
          label_name: git.branch
          label_value: develop
          match_mode: exact
          name: gitlab_branch_develop
        - invert_match: false
          label_name: git.branch
          label_value: master
          match_mode: exact
          name: gitlab_branch_master
...
```