# job-worker-service

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
    1. Able to stream to multiple client and output to should be similar to `docker -f`
    2. Test logs that outputs non-UTF-8 encoded content.
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