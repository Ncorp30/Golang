package leetcode

import "testing"

func Test_two_sum(t *testing.T) {

	if result := TwoSum([]int{1, 3, 5}, 4); len(result) < 2 || (result[0] != 0 && result[1] != 1) {
		t.Errorf("Error")
	}
	if result := TwoSum([]int{3, 1, 5}, 6); len(result) < 2 || (result[0] != 1 && result[1] != 2) {
		t.Errorf("Error")
	}
}