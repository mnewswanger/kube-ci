# Notifiers #

## Examples ##

### Slack ###

```yaml
---
namespace: slack
name: general
type: webhook
properties:
  url: "<YOUR_SLACK_WEBHOOK_URL>"
  body: '{"text":"{{.Job.Namespace}} - {{.Job.Name}}","attachments":[{"title":"Image!","title_link":"https://github.com/mnewswanger/kube-ci","text":":+1: The job worked!","image_url":"https://avatars2.githubusercontent.com/u/12140008?s=460&v=4","color":"#764FA5"}]}'
  method: "POST"
  headers:
    "Content-Type": "application/json"
...
```
