# Testing memory: Example 1

1. Set memory.max to 100 MB 
2. Pass the command such as
```json
	{
		"command_args": [
		   "--vm",
		   "2",
		   "--vm-bytes",
		   "90M",
		   "--timeout",
		   "120"
		],
		"command_name": "stress"
	}
```
Note: seems like the "stress" voluntarily exits before the OOM killer intervenes. It returns exit code as 2 for an internal error rather than SIGKILL.
3.  Check if job_<uuid> cgroup is generated and cleanedup
```
cd /sys/fs/cgroup
watch -n 1 "ls -ltr | grep 'job'"
```
4. Note: need to do `sudo swapoff -a` so that OOM happens

# Testing memory: Example 2

1. Set memory.max to 100 MB 
2. Pass the command such as
```
{
    "command_args": [
        "/home/fedora/job-worker-service/memory_hog.py"
    ],
    "command_name": "python"
}
```
3. Run the following python program
```py
memory_hog = []
while True:
    memory_hog.append(' ' * 1024 * 1024 * 10)  # Allocate 10MB blocks
```
4. In this case, you should get signal 9


# Testing CPU
*

# Testing IO