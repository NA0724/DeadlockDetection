# Deadlock Detection using Banker's Algorithm 

##What is Banker's Algorithm?

The Banker's Algorithm is a deadlock avoidance algorithm that is used in computer operating systems to prevent deadlocks from occurring. Deadlocks occur when two or more processes are blocked and waiting for each other to release resources, resulting in a system-wide halt.

The Banker's Algorithm ensures that the system never enters an unsafe state where deadlocks can occur by determining whether a resource allocation request from a process will leave the system in a safe state or not.

The algorithm works by maintaining information about the resources available, the resources currently allocated to processes, and the resources needed by processes. It uses this information to calculate whether a new resource request can be granted without causing a deadlock.

##Where is Banker's Algorithm used?

The Banker's Algorithm is commonly used in multi-process and multi-threaded systems, such as operating systems, database management systems, and distributed systems, where multiple processes or threads may need to access shared resources.

In particular, the Banker's Algorithm can be used in systems where resources are limited, such as in banking applications, where multiple customers may be accessing their accounts simultaneously and requesting transactions that require the allocation of shared resources, such as funds or credit.


###Tech Stack

Go version - 1.19
