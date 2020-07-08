# API Platform CS monitoring

## Overview

OCI Cloud Native services come with seemingly integrated Monitoring service. But what if a bit older services, like API Platform CS, are used? There are a few approaches that can be utilized, depending on the metrics we are interesting in.

This repo showcases one of the approaches, based on the REST APIs exposed by the API Platform service. It's by no means the complete solution, just a demo to get a feel of a concept.

## Functionality

1. Application checks the last polling time for the list of configured gateways
2. If defined threshold is not met for any of them - email notification is sent to defined receipients

## Architecture

![Architecture](./img/architecture.png)


## Deployment Architecture

![Architecture](./img/deployment_architecture.png)

## Installation
