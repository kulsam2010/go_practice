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

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
