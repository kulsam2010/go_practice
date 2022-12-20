package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Pass to buffer ", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet("Sameer", &buffer)

		got := buffer.String()
		want := "Hello Sameer"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

}
