# Documentation ##

https://docs.zeus.fyi

# Recent Articles ##

## Show Me the Stats
Optimal adaptive load balancing in stochastic environments.
Recommended reading for scientists, engineers, data driven individuals

![Screenshot 2023-09-14 at 11 12 23 PM](https://github.com/zeus-fyi/zeus/assets/17446735/025d3201-9236-40e9-8723-3a7d2d7a3e0a)

https://medium.com/zeusfyi/show-me-the-stats-6740f8d6d0b7

## Adaptive RPC Load Balancer on QuickNode Marketplace
Accurate, Reliable, Performant Node Traffic at Web3 Scale

![Screenshot 2023-09-14 at 11 11 30 PM](https://github.com/zeus-fyi/zeus/assets/17446735/802b7670-6b30-4e65-9348-e45e2a0cfcac)

https://medium.com/zeusfyi/adaptive-rpc-load-balancer-on-quicknode-marketplace-e68bb7c9d8ac

![Screenshot 2023-09-14 at 11 13 55 PM](https://github.com/zeus-fyi/zeus/assets/17446735/1d2a263e-5aa7-418c-a0f0-1f497cab0353)

Beta Testing Sign Up (free to use): 
https://marketplace.quicknode.com/add-on/zeusfyi-4

## zK8s == Kubernetes + Zeus
Here we overview the core concepts needed to understand how you can build, deploy, configure K8s apps using Zeus, with a full walkthrough example of how we created an Ethereum beacon.

https://medium.com/@zeusfyi/zeus-k8s-infra-as-code-concepts-47e690c6e3c5

## Authenticated API in 5 Minutes
Step by step tutorial using our UI

https://medium.com/@zeusfyi/zeus-ui-no-code-kubernetes-authenticated-api-tutorial-c468d5ef0446

## zK8s Apps & Clients ##
zK8s is an expressive language for cloud infrastructure, used for building, assembling, and keeping them running over their entire lifecycle. Enabling cost efficient, effortless large scale infra automation, coordination, customization, and control.

#### ```cluster_config_drivers ```
#### ```system_config_drivers ```
#### ```workload_config_drivers ```

Workflow & Proxy Programmable Automation (Rolling releases coming through end of year)

#### ```artemis_workflows ```
#### ```iris_programmable_proxy ```

QuickNode MarketPlace users can find Load Balancing documentation in the iris programmable proxy directory
+ Adaptive Load Balancer Documentation
  
#### API Endpoints 

Documentation and code examples are found here
[API_README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/API_README.md)

How to use the test suite to setup your own api calls
[README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/README.md)

The test directory contains useful mocks and tools for interacting with the API. It also contains a useful
config-sample.yaml, convert this to config.yaml and set your bearer token here, which then allows you to
use the demo code to create your first api request in seconds

## Overview

1. Automates translation of kubernetes yaml configurations into representative SQL models
2. Users upload these infrastructure configurations via API where they are stored in the DB
3. Users can then query the contents of these infrastructure components, deploy, mutate or destroy them on demand
   
### Currently Supported Workload Types

1. Deployments
2. StatefulSets
3. Services
4. ConfigMaps
5. Ingresses
6. ServiceMonitors
7. Secrets

+ Node Tainting Automation & Infra Provisioning (Servers & SSDs) Deployable & Scalable on Demand.

### Pods Controller 

1. GetPods
2. GetPodLogs
3. PortforwardReqToPods
4. DeletePods

Not every possible field type is supported, but the most common ones are, and even a decent amount of the uncommon ones. If you find a field you need isn't supported please send us an email at support@zeus.fyi

### Hades Library

Hades is used to interact with Kubernetes workloads via API, and can apply saved Zeus workloads & cookbooks onto your
own in house infrastructure.

### Hera Client

This client uses the OpenAI API to generate code with AI. This service is available at OpenAI cost, so just pay for the
token cost, otherwise it is free to use.

## Cookbooks ##

Contains common web2 & web3 recipes for deploying applications, and managing infrastructure using zK8s or vanilla
Kubernetes.

### Zeus UI Highlights

![Screenshot 2023-04-11 at 1 58 00 PM](https://user-images.githubusercontent.com/17446735/231288647-e3bf8db7-67c9-4e0c-af38-c0f1091da726.png)

![Screenshot 2023-04-11 at 1 57 46 PM](https://user-images.githubusercontent.com/17446735/231288668-f8c147fe-d06b-49e1-8313-a7453cbe6f19.png)

![Screenshot 2023-04-07 at 6 46 08 PM](https://user-images.githubusercontent.com/17446735/231288683-a350f36b-d103-428f-88b3-eac80742a9c4.png)

![Screenshot 2023-04-05 at 2 11 33 PM](https://user-images.githubusercontent.com/17446735/231288689-f970cd81-76b3-4b85-9241-3f30ad7c80b9.png)
