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
4. cgroups testing
```
./worker-service start -command "sort" "largefile.txt -o sortedfile.txt" --cpu-limit 512 --memory-limit 500M --disk-limit 500

Notes on cgroup:
// Default value of CPU limit (`cpu.shares`) is 1024. 512 would mean half of that CPU share.
// Caps the memory limit (`memory.limit_in_bytes`) to 500MB. Exceeding will cause OOM.
// Disk limit (`blkio.weight`): Between 10-1000. Determines IO priority.
```
5. Coverage 
    1. Basic unit tests
    2. Test with `-race` flag
    3. govulncheck