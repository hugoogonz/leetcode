# 1. Two Sum

`Easy`

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
 
**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

```
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
```

**Solutions:**

Big O notation: O(n):

```go
func twoSum(nums []int, target int) []int {

	newMap := make(map[int]int)

	for i, value := range nums {
		c := target - value
		if index, ok := newMap[c]; ok {
			return []int{index, i}
		}
		newMap[value] = i
	}
	return []int{}
}
```

Big O notation: O(n^2):

```go
func twoSum(nums []int, target int) []int {
	for i, value := range nums {
		for j := i + 1; j < len(nums); j++ {
			if value + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
```