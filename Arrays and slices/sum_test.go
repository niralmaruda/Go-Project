package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})
	t.Run("Sum of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d and want %d", got, want)
	}
}
