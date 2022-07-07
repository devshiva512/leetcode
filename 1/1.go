package main

import (
	"fmt"
	"time"
	//"log"
)

// Problem: https://leetcode.com/problems/two-sum/
// Tags: Arrays

func main() {
	nums, target := []int{2,7,11,15}, 9

	fmt.Printf("Method: %s, Output: %s", "Brute Force", twoSum_BruteForce(nums, target))
	fmt.Printf("Method: %s, Output: %s", "Hash Map", twoSum_HashMap(nums, target))
}

// Method 1: Brute force    [Time Complexity - O(n^2), Space - O(1)]
func twoSum_BruteForce(nums []int, target int) []int {
	for i, val := range nums {
		for j:=i+1; j < len(nums); j++ {
			if (val + nums[j] == target){
				return []int{i,j}
			}
		}
	}

	return []int{}
}

// Method 2: Using HashMap  [Time Complexity - O(n), Space - O(n)]
func twoSum_HashMap(nums []int, target int) []int {
    searchMap := make(map[int]int)
    
    for i, val := range nums {
        diff := target - val 
        
        if indexOfDiff, ok := searchMap[diff]; ok {
            return []int{i, indexOfDiff}
        }
        
        searchMap[val] = i 
    }
    
    return []int{}
}