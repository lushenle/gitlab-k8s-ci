gitlab-k8s-ci
---
[![Build Status](https://travis-ci.com/lushenle/gitlab-k8s-ci.svg?branch=v1.0.0)](https://travis-ci.com/lushenle/gitlab-k8s-ci)
[![GitHub issues](https://img.shields.io/github/issues/lushenle/gitlab-k8s-ci.svg)](https://github.com/lushenle/gitlab-k8s-ci/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/lushenle/gitlab-k8s-ci)](https://goreportcard.com/report/github.com/lushenle/gitlab-k8s-ci)
[![Coverage Status](https://coveralls.io/repos/github/lushenle/gitlab-k8s-ci/badge.svg?branch=v1.0.0)](https://coveralls.io/github/lushenle/gitlab-k8s-ci?branch=master)
[![GoDoc](https://godoc.org/github.com/lushenle/gitlab-k8s-ci?status.svg)](https://godoc.org/github.com/lushenle/gitlab-k8s-ci)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://blog.abelotech.com/mit-license/)

These are the example files for presentation about GitLab + Kubernetes for Continuous Integration and Delivery. They are also partly used in my GitLab CI posts.

## Table of Contents
* [Requirements](#requirements)
* [Features](#features)
* [Using this repository](#using-this-repository)
* [GitLab Docs References](#gitlab-docs-references)
* [File Structure](#file-structure)
    * [Example Application](#example-application)
    * [Kubernetes Base GitLab CI Manifests](#kubernetes-base-gitlab-ci-manifests)
    * [Build Process](#build-process)
    * [Deployment Manifests](#deployment-manifests)

---

## Requirements

The following points are required for this repository to work correctly:
* GitLab (`>= 11.3`) with the following features configured:
    * [Container Registry](https://docs.gitlab.com/ce/user/project/container_registry.html)
    * [GitLab CI](https://about.gitlab.com/features/gitlab-ci-cd/) (with working [GitLab CI Runners](https://docs.gitlab.com/ce/ci/runners/), at least version `>= 11.3`)
* [Kubernetes](https://kubernetes.io/) cluster
    * You need to be "bound" to the `admin` (`cluster-admin`) ClusterRole, see [Kubernetes.io Using RBAC Authorization - User-facing Roles](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles).
    * An Ingress controller should already been deployed, see [Kubernetes.io Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/).
* `kubectl` installed locally.
* Editor of your choice.

## Features

This repository shows off/uses the following GitLab CI features:
* [GitLab CI](https://docs.gitlab.com/ce/ci/README.html)
    * [Manual CI Steps](https://docs.gitlab.com/ce/ci/yaml/#when-manual)
    * [Artifacts](https://docs.gitlab.com/ce/user/project/pipelines/job_artifacts.html)
    * [App review](https://docs.gitlab.com/ce/ci/review_apps/index.html)
* [GitLab Container Registry](https://docs.gitlab.com/ce/user/project/container_registry.html)
* [GitLab CI Kubernetes Cluster Integration](https://docs.gitlab.com/ce/user/project/clusters/index.html)

Other features also shown are:
* [coreos/prometheus-operator ServiceMonitor]() - for automatic monitoring of deployed applications.

## Using this repository

You have to replace the following addresses in all files:

* `git.moqi.ai` with your GitLab address (e.g. `gitlab.example.com`).
* `ops.internal.moqi.ai` (in the Ingress manifest) with your domain name.
    * You probably also want to change the subdomain name while you are at it.
* `presentation-gitlab-k8s` with the Namespace name of your choice.

If you are using [coreos/prometheus-operator](https://github.com/coreos/prometheus-operator), then you also need to replace
`monitoring` with the Namespace your Prometheus instance is running in,
in this file [`/gitlab-ci/monitoring/service-monitor.yaml`](/gitlab-ci/monitoring/service-monitor.yaml).
You then also want to `kubectl` create/apply the file to your Kubernetes cluster during creation/apply process for the manifests in [`gitlab-ci/`](/gitlab-ci/).

You also need to create a "Docker Login" Secret which contains your GitLab Registry access data (e.g. Username and Access token with registry access) named whatever your want in the Namespace `presentation-gitlab-k8s`.
A guide for that can be found here: [Kubernetes.io - Pull an Image from a Private Registry](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/).
Instead of using the `imagePullSecrets`, we'll be using the `default` `ServiceAccount` in the  Namespace to automatically use the created Docker login `Secret`, see [Kubernetes - Configure Service Accounts for Pods - Add ImagePullSecrets to a service account](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account).

The Namespace manifest is in the [`gitlab-ci/`](/gitlab-ci/) directory.

Then you can just import the repository into your GitLab instance and are ready to go.

For information on how to use these files and setup GitLab Kubernetes cluster/integration, see the above blog post and in specific this post [GitLab + Kubernetes: Perfect Match for Continuous Delivery with Container](https://edenmal.moe/post/2017/GitLab-Kubernetes-Perfect-Match-for-Continuous-Delivery-with-Container/).

## GitLab Docs References

* GitLab Kubernetes Integration Docs: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html
* GitLab Kubernetes Integration Docs Environment variables: https://docs.gitlab.com/ce/user/project/integrations/kubernetes.html#deployment-variables

As of GitLab `10.3` the Kubernetes Integration is marked as deprecated and with `10.4` it is now disabled, the following docs show the new feature called Clusters:
* GitLab 10.3 release - Kubernetes integration service: https://about.gitlab.com/2017/12/22/gitlab-10-3-released/#kubernetes-integration-service
* GitLab Clusters Feature Docs: https://docs.gitlab.com/ce/user/project/clusters/index.html

## File Structure

### Example Application

* [`main.go`](/main.go) - The Golang example application code.
* [`go.mod`](/go.mod) and [`go.sum`](/go.sum) - [Golang modules files](https://github.com/golang/go/wiki/Modules).

### Kubernetes Base GitLab CI Manifests

* [`gitlab-ci/`](/gitlab-ci/)
    * [`monitoring/`](/gitlab-ci/monitoring/)
        * [`service-monitor.yaml`](/gitlab-ci/monitoring/service-monitor.yaml) - Contains a coreos/prometheus-operator ServiceMonitor manifest to automatically monitor the application(s).
    * [`namespace.yaml`](/gitlab-ci/namespace.yaml) - Namespace in which the GitLab CI will deploy the application.
    * [`rbac.yaml`](/gitlab-ci/rbac.yaml) - Contains GitLab CI RBAC Role, RoleBinding and ServiceAccount.
    * [`secret.yaml`](/gitlab-ci/secret.yaml) - Contains a TLS wildcard certificate for the application Ingress.

### Build Process

* [`Dockerfile`](/Dockerfile) - Contains the Docker image build instructions.
* [`.gitlab-ci.yml`](/.gitlab-ci.yml) - Contains the GitLab CI instructions.

### Deployment Manifests

* [`manifests/`](/manifests/) - Kubernetes manifests used to deploy the Docker image built in the CI pipeline.
    * [`deployment.yaml`](/manifests/deployment.yaml) - Deployment for the Docker image.
    * [`ingress.yaml`](/manifests/ingress.yaml) - Ingress for the application.
    * [`service.yaml`](/manifests/service.yaml) - Service for the application.
