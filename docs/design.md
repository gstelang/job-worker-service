# Summary

* Implement a prototype job worker service that provides an API to run arbitrary Linux processes.

# Functional requirements

* https://github.com/gravitational/careers/blob/main/challenges/systems/challenge-1.md#level-4

# Architecture diagram

![Design Diagram](job-worker-service.png)

# API 

* Refer to the [proto spec](../service/proto/jobworker.proto).

# CLI interface

1. Start
```
./worker-service start -command "sort" "largefile.txt -o sortedfile.txt"
example output: job with id 761db04c-0150-4f0b-a6fd-5cab9b9a48bf started.
```
2. Stop
```
./worker-sevice stop 761db04c-0150-4f0b-a6fd-5cab9b9a48bf
Output: Success! job with id 761db04c-0150-4f0b-a6fd-5cab9b9a48bf stopped.
```
3. Query
```
./worker-service query 761db04c-0150-4f0b-a6fd-5cab9b9a48bf
Output: job with id 761db04c-0150-4f0b-a6fd-5cab9b9a48bf [is running|has stopped].
```
4. Stream
```
./worker-service stream 761db04c-0150-4f0b-a6fd-5cab9b9a48bf
Output: 
2024-08-26T14:35:21Z [INFO] Server started on port 8080
2024-08-26T14:35:22Z [DEBUG] foo bar
2024-08-26T14:35:23Z [ERROR] bar foo
2024-08-26T14:35:24Z [WARN] High memory usage with foo as bar
2024-08-26T14:35:25Z [INFO] bar is foo now.
......
......
```
5. Start with resource limits
```
./worker-service start -command "sort" "largefile.txt -o sortedfile.txt" --cpu-limit 512 --memory-limit 500M --disk-limit 500
example output: job with id 761db04c-0150-4f0b-a6fd-5cab9b9a48bf started with resource limits.
```

# Security (Authn/Authz)
* Authentication: Use openssl to generate certificates to implement mTLS i.e verify identity of both client and server.
* Role based authorization: Control access based on user roles.
    1. Only 2 roles - Admin and User.
    2. Admin can do all operations permitted  
        1. Start, stop, query and stream  
        2. Start the process with the resource limits - CPU, memory and disk only.
    3. User can only query and stream 

# Out of scope
* State of job will not persist after restarts i.e no persistent storage such as log files or local sqllite database.

# Non-functional requirements
* Security: Should pass vulncheck
* Race conditions: Should pass the race detector test