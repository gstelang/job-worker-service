# Summary

* Implement a prototype job worker service that provides an API to run arbitrary Linux processes.

# Functional requirements

* https://github.com/gravitational/careers/blob/main/challenges/systems/challenge-1.md#level-4

# Job management 
* All the jobs details since the start of the server will be stored in memory.
```json
{
    "e17980c1-cbdc-46f9-a302-cf39b8bba501": {
        "status": "running|exited|signaled",
        "pid": 123,
        "exit_code": "EXIT_CODE_INT",
        "signal": "SIG_CODE_INT",
        "logs": [
            "foo happened at start",
            "bar thereafter after 1 sec",
            "...."
        ]
    }
}
```
* I will be utilizing Go's os/exec specifically [exec.Command](https://pkg.go.dev/os/exec#Command) to start and gather information on process and signal status. Exit code or signal info on job's end of execution will be captured and stored till server's lifetime.

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

# Authentication
* Communication between client and server will be with mutual TLS i.e both needs to verify their identity. 
* Use openssl to generate certificates. A certificate authority will sign all the clients and server cert.

# Authorization
* Role based authorization: Control access based on user roles.
    1. Only 2 roles - Admin and User.
    2. Admin can do all operations permitted  
        1. Start, stop, query and stream  
        2. Start the process with the resource limits - CPU, memory and disk only.
    3. User can only query and stream 
* Implementation:
    * 2 client certs with embedded OIDs (object identifiers) will be generated using openssl. This approach ties authentication (who you are) directly to authorization (what you're allowed to do).
    * The server can make authorization decisions based on the role OID in the client certificate.
![Authorization](authorization.png)

# TLS Setting
1. For the purpose of this project, allowing only clients with TLS1.3. As per docs [here](https://pkg.go.dev/crypto/tls@master) and [here](https://go-review.googlesource.com/c/go/+/314609), cipher suite selection with tls 1.3 is automatic. Here's the Go code server side, I plan to use.
```go
	creds := credentials.NewTLS(&tls.Config{
		MinVersion:   tls.VersionTLS13, // Only allow TLS 1.3
		MaxVersion:   tls.VersionTLS13, // Only allow TLS 1.3
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})
```

# Certs for mTLS
* All certs will be created with key size of 4096 bits using RSA.
* SHA-256 will be used as a hashing algorithm for signing certs.
* CA and server certs will be valid for 365 days whereas a shorter validity period (45 days) will be provided for client certs. This is to emphasize rotation and renewal. 

# Out of scope
* State of job will not persist after restarts i.e no persistent storage such as log files or local sqllite database.

# Non-functional requirements
* Security: Should pass vulncheck
* Race conditions: Should pass the race detector test