# Summary

* Implement a prototype job worker service that provides an API to run arbitrary Linux processes.

# Functional requirements

* https://github.com/gravitational/careers/blob/main/challenges/systems/challenge-1.md#level-4

# Architecture diagram

![Design Diagram](job-worker-service.png)

# API 

* Refer to the [proto spec](../service/proto/jobworker.proto).

# Security (Authn/Authz)

# Out of scope
* State of job will not persist after restarts i.e no persistent storage such as log files or local sqllite database

# Non-functional requirements
* Security: Should pass vulncheck
* Race conditions: Should pass the race detector test