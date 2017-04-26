# Kubernetes Continuous Integration #

## About the Project ##

KubeCI is designed to handle Continuous Integration and Continuous Delivery workloads via Jobs.

## Features ##

* Interacts with `kubernetes` `jobs` and `deployments`
* Handles webhooks in both synchronous and asynchronous modes
* Enable distributed deployment with multi-master capacity for large-scale and highly available deployments
* Offer management both via Web Interface and CLI

## Project Goals ##

### Interact Well With Kubernetes ###

KubeCI provides the ability to interact with `kubernetes` to run jobs and manipulate deployments.

### Easy to Deploy ###

KubeCI can be fully deployed with no interactive steps required from the user.

### Easy to Manage at Scale ###

Because the service can be managed via CLI using JSON / YAML payloads, KubeCI can be easily deployed via configuration management tools.

### Provide A Scalable Solution ###

Provide the ability to run across mutliple master servers with fault tolerance through the entire pipeline process.

### Provide Visibility to Users ###

KubeCI provides `prometheus` endpoints for statistic tracking for jobs.  It logs in formats easily parsable by common log aggregation projects.  It also provides easy log viewing via the the Web UI.
