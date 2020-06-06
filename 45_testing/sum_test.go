package main

import (
	"fmt"
	"testing"
)

func TestSum1(t *testing.T) {
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	if total != 45 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 45)
	} else {
		t.Logf("(1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9), got: %d, want: %d.", total, 45)
	}
}

func TestSum2(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		} else {
			t.Logf("(%d + %d) , got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestSum3(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		t.Run(fmt.Sprintf("Sum of %d, %d", table.x, table.y), func(t *testing.T) {
			total := sum(table.x, table.y)
			if total != table.n {
				t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
				t.Fatal("test case failed.")
			} else {
				t.Logf("(%d + %d) , got: %d, want: %d.", table.x, table.y, total, table.n)
			}
		})
	}
}

func TestSum4(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		table := table // capture range variable
		t.Run(fmt.Sprintf("Sum of %d, %d", table.x, table.y), func(t *testing.T) {
			total := sum(table.x, table.y)
			t.Parallel()
			if total != table.n {
				t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
				t.Fatal("test case failed.")
			} else {
				t.Logf("(%d + %d) , got: %d, want: %d.", table.x, table.y, total, table.n)
			}
		})
	}
}
