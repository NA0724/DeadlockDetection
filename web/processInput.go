package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Process the input given from user and returns the following vectors
existing: existing maximum resources
totalAllocated: total allocated resources
pid_allocated: resources allocated to each process
available: available resources
pid_needed: resources needed by a process to complete
pReq: slice of requesting processes
req: slice of resources requested by each process
*/
func processInput(input []string) ([]int, []int, []int, map[string][]int, map[string][]int, []string, [][]int) {

	existing := convert(input[0])
	r := len(existing)
	available := make([]int, r)
	totalAllocated := make([]int, r)
	pid_allocated := make(map[string][]int)
	pid_needed := make(map[string][]int)
	var pReq []string
	var req [][]int

	for _, str := range input[1:] {
		str = strings.TrimSpace(str)
		s := strings.Split(str, " ")
		pid := s[0]
		flag := false
		//find the index of assigned, need and request words in each line
		var assignedIdx, needIdx, reqIdx int
		for i, s1 := range s {
			if strings.Contains(s1, "assigned") {
				assignedIdx = i
				flag = true
			} else if strings.Contains(s1, "need") {
				needIdx = i
			} else if strings.Contains(s1, "request") {
				reqIdx = i
			}
		}
		var assigned, needed, requested []int
		if flag {
			for _, s1 := range s[assignedIdx+1 : needIdx] {
				if tmp, err := strconv.Atoi(s1); err == nil {
					assigned = append(assigned, tmp)
				}
			}
			for _, s1 := range s[needIdx+1:] {
				if tmp, err := strconv.Atoi(s1); err == nil {
					needed = append(needed, tmp)
				}

			}
			pid_allocated[pid] = assigned
			pid_needed[pid] = needed
			totalAllocated = addSlices(totalAllocated, assigned)
		} else {
			for _, s1 := range s[reqIdx+1:] {
				if s1 != "" {
					if tmp, err := strconv.Atoi(s1); err == nil {
						requested = append(requested, tmp)
					}
				}
			}
			pReq = append(pReq, pid)
			req = append(req, requested)
		}
	}
	available = subSlices(existing, totalAllocated)
	return existing, totalAllocated, available, pid_allocated, pid_needed, pReq, req
}

// function to add multiple slices
func addSlices(slices ...[]int) []int {
	var result []int
	for i := 0; ; i++ {
		var sum int
		for _, slice := range slices {
			if i >= len(slice) {
				return result
			}
			sum += slice[i]
		}
		result = append(result, sum)
	}
}

// function to subtract two slices
func subSlices(slice1, slice2 []int) []int {
	if len(slice1) != len(slice2) {
		return nil
	}
	result := make([]int, len(slice1))
	for i := range slice1 {
		result[i] = slice1[i] - slice2[i]
	}
	return result
}

// convert string to int slice
func convert(line1 string) []int {
	getLine1 := strings.Split(line1, " ")
	result := make([]int, 0, len(getLine1))
	for _, str := range getLine1 {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Print("Error in parsing string to int: ", err)
			fmt.Println()
			return nil
		}
		result = append(result, num)
	}
	return result
}
