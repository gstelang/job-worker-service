# Testing
1. Basic start/stop/query/stream
```
./worker-service start -command "sort" "largefile.txt -o sortedfile.txt"
./worker-sevice stop <jobid>
./worker-service query <jobid>
./worker-service stream <jobid>
```
2. Test start with 
    1. Basic commands
    2. Running process (simulate by infinite do while loop)
3. Stream to multiple clients (concurrency + real time updates)
    1. Able to stream to single client and output to should be similar to `docker -f`
    2. Test logs that outputs non-UTF-8 encoded content.
    3. Should be able to stream to multiple clients.
    4. Simulate a slow client (add a delay in the client's processing loop). There should be no degradation to other clients.
    5. Simulate the failure of the Linux process whose logs are being streamed. Verify that the server handles this gracefully, such as stopping the log stream and notifying the clients of the exit code.
    6. Test with large log messages.
4. cgroups testing
    1. Make sure resource limits are enforced based on time of the day. Check 
    ```
    cat /sys/fs/cgroup/cpu/my_cgroup/cpuacct.usage
    ```
    2. Make sure cgroups are cleaned up after usage.
5. Coverage 
    1. Basic unit tests
    2. Test with `-race` flag
    3. govulncheck
