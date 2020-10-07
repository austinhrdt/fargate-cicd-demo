# Fargate via GitHub Actions

![Build](https://github.com/austinhrdt/fargate-cicd-demo/workflows/Build/badge.svg)

## Introduction

Using GitHub Actions to build and deploy a small, containerized, Go API on AWS Fargate. The container is built using [go-build-template](https://github.com/austinhrdt/go-build-template) and pushed to AWS Elastic Container Registry (ECR). Using the newly built container image, a task definition is rendered and deployed.

[View Slides >>](https://austinhrdt.github.io/fargate-cicd-demo/)
