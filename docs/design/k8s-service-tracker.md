## Objective
To make connection tracing easier for applications running on kubernetes. The wider goal is to make network tracing easier, but to reduce scope, need to start with something small and give immediate results, i.e. connection tracking. 

## Goals:
- provide an interface for admins to see failed connection attempts of applications deployed on kubernetes
- provide an interface to track network outages for the application

## Non-goals
- TODO

## Design

- take the input that recognises what service to track connections for.
  1. in the form of kubernetes service name
  2. in the form of node name and port on the node
  3. in the form of node selector and port on the node
- for each pod serving traffic to the selected service, create a subresource:
    this is required to avoid racing when updating the same resource. This 
    feels like a cleaner architecture.
- update the subresource for each pod serving traffic with the following details:
  1. failure source IP and port 
  2. closed connection counter.
  3. if failure source IP is from a pod in cluster then display the podname
  3. if any, open connections counter from this source
  4. if any failures observed, a list of: 
     1. timestamp of first failure observed
     2. timestamp of last failure observed, blank if no successful connection from this source IP
        after the first failure observed.
- pick the failures for each serving pod from the subresource and aggregated list in main resource.
## Alternatives
