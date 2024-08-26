# job-worker-service

# Testing
1. Basic start/stop/query/stream
```
./worker-service start -command "wc -l abc.txt"
./worker-sevice stop <jobid>
./worker-service query <jobid>
./worker-service stream <jobid>
```
2. Test start with 
    1. Basic commands
    2. Running process (simulate by infinite do while loop)
3. Stream to multiple clients (concurrency + real time updates)
    1. Able to stream to multiple client and output to should be similar to `docker -f`
4. cgroups testing commands (TBD)

# Coverage
1. Basic unit tests
2. Test with `-race` flag
3. govulncheck