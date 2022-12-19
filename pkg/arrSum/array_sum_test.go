package arrSum

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("Sum positive array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}

		got := Sum(arr)
		want := 15

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 3}, []int{4, 1})
	want := []int{5, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
	t.Run("Make sums of small slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{9, 10, 11})
		want := []int{5, 21}

		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{9, 10, 11})
		want := []int{0, 21}

		checkSums(t, got, want)
	})
}
