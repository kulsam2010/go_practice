package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Check repeat", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		if got != want {
			t.Errorf("Got %q , wanted %q", got, want)
		}

	})
}
