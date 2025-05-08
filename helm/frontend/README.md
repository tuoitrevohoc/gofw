# Frontend Helm Chart

This Helm chart is designed to deploy the frontend application of the project. It provides a simple way to manage the deployment, service, and ingress resources required for the application.

## Prerequisites

- Kubernetes cluster
- Helm installed on your local machine

## Verify Helm Chart before installing

```shell
cd gofw/helm
helm dependencies build ./frontend
helm lint ./frontend/ --set environmentName=dev 
```

To debug the template

```shell
helm install reactjs-app ./frontend/  --create-namespace --namespace frontend --dry-run
helm template reactjs-app ./frontend  --create-namespace --namespace frontend --debug > ./frontend/template-result.yaml
```

## Installation

To install the chart, use the following command:

```shell
kubectl create namespace frontend
helm install reactjs-app ./frontend  --create-namespace --namespace frontend
```

To update

```shell
helm upgrade -i reactjs-app ./frontend/
```

## Configuration

You can customize the deployment by modifying the `values.yaml` file. This file contains default configuration values that can be overridden during installation.

## Usage

After installation, you can access the frontend application through the service or ingress defined in the chart. Make sure to configure the ingress resource with the appropriate host and paths to route traffic correctly.

### Install Kubernetes Dashboard

https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/

## Uninstallation

To uninstall the chart, use the following command:

```shell
helm uninstall reactjs-app
```