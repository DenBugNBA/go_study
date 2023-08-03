package main

import (
	"fmt"
	"sort"
)

func UniqueSortedUserIDsTeacher(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	uniqueIds := map[int64]bool{}

	for _, id := range userIDs {
		uniqueIds[id] = true
	}

	userIDs = []int64{}

	for k, _ := range uniqueIds {
		userIDs = append(userIDs, k)
	}

	sort.Slice(userIDs, func(i int, j int) bool {
		return userIDs[i] < userIDs[j]
	})

	return userIDs
}

func main() {
	nums := []int{2, 1, 6, 5, 3, 4}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums) // [1 2 3 4 5 6]

	nums2 := []int{2, 1, 6, 5, 3, 4}
	sort.SliceStable(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	fmt.Println(nums2) // [1 2 3 4 5 6]
}
