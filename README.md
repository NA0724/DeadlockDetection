# Deadlock Detection using Banker's Algorithm for multiple processes and resources

What is Banker's Algorithm?
The Banker's Algorithm is a deadlock avoidance algorithm that is used in computer operating systems to prevent deadlocks from occurring. Deadlocks occur when two or more processes are blocked and waiting for each other to release resources, resulting in a system-wide halt.

The Banker's Algorithm ensures that the system never enters an unsafe state where deadlocks can occur by determining whether a resource allocation request from a process will leave the system in a safe state or not.

The algorithm works by maintaining information about the resources available, the resources currently allocated to processes, and the resources needed by processes. It uses this information to calculate whether a new resource request can be granted without causing a deadlock.



Go version - 1.19
