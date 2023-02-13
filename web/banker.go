package main

import (
	"fmt"
)

// implement bankers algorithm for deadlock detection and check if current state is safe or not
func bankers(available []int, pid_allocated map[string][]int, pid_needed map[string][]int) bool {

	executed := false
	safe := false
	finished := make(map[string]bool)
	for key := range pid_allocated {
		finished[key] = false
	}
	p := len(pid_allocated)
	r := len(available)

	for i := p; i > 0; i-- {
		for key := range pid_allocated {
			if !finished[key] { //if process is not finished
				executed = true
				curr := pid_allocated[key]
				need := pid_needed[key]
				for j := 0; j < r; j++ {
					if need[j] > available[j] {
						executed = false
						break
					}
				}
				if executed {
					finished[key] = true // mark the process as finished if executed is true
					safe = true
					for j := 0; j < r; j++ {
						available[j] += curr[j]
					}
					break
				} else {
					safe = false
				}
			}
		}
		if !safe {
			break
		}
	}
	return safe
}

/*
fetch the current state and verify if it is safe or not and then
if the first request is safe (and granted if safe, denied if unsafe), and then the second request, and so on
*/
func doOperations(input []string) {
	existing, totalAllocated, available, pid_allocated, pid_needed, pReq, req := processInput(input)

	isSafe := bankers(available, pid_allocated, pid_needed)
	if isSafe {
		fmt.Println("Current State is safe")
	} else {
		fmt.Println("Unsafe state")
		return
	}

	for i, key := range pReq {
		// granting the request for process id = key
		pid_allocated[key] = addSlices(pid_allocated[key], req[i])
		pid_needed[key] = addSlices(pid_needed[key], req[i])
		totalAllocated = addSlices(totalAllocated, req[i])
		available = subSlices(existing, totalAllocated)
		//check if safe after granting
		safe := bankers(available, pid_allocated, pid_needed)
		if safe {
			fmt.Println(key, " is granted")
		} else {
			fmt.Println(key, " is denied")
			// revert the granted requested back to original form
			pid_allocated[key] = subSlices(pid_allocated[key], req[i])
			pid_needed[key] = subSlices(pid_needed[key], req[i])
			totalAllocated = subSlices(totalAllocated, req[i])
			available = addSlices(existing, totalAllocated)
			return
		}
	}

}
