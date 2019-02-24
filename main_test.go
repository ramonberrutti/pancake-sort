package main

import "testing"

func TestFlip(t *testing.T) {
	tables := []struct {
		arr    []int
		n      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 1, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{4, 3, 2, 1, 5, 6}},
	}

	for _, table := range tables {
		flip(table.arr, table.n)
		if len(table.arr) != len(table.expect) {
			t.Errorf("Invalid arr lenght, Expected: %d Get: %d", len(table.expect), len(table.arr))
		} else {
			for i := range table.arr {
				if table.arr[i] != table.expect[i] {
					t.Errorf("Invalid array value, Expected: %d Get: %d", table.expect[i], table.arr[i])
				}
			}
		}
	}
}

func TestPancakeSort(t *testing.T) {
	tables := []struct {
		arr    []int
		expect []int
	}{
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{4, 3, 7, 1, 6, 2, 5}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{50, 40, 25, 1, 30, 100}, []int{1, 25, 30, 40, 50, 100}},
		{[]int{3, 4, 1, 2, 8, 7, 9, 10, 50, 5, 500, 66, 43}, []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 43, 50, 66, 500}},
	}

	for _, table := range tables {
		pancakeSort(table.arr)
		if len(table.arr) != len(table.expect) {
			t.Errorf("Invalid arr lenght, Expected: %d Get: %d", len(table.expect), len(table.arr))
		} else {
			for i := range table.arr {
				if table.arr[i] != table.expect[i] {
					t.Errorf("Invalid array value, On index: %d expected: %d Get: %d", i, table.expect[i], table.arr[i])
				}
			}
		}
	}
}
