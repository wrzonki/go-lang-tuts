package main

import (
	"testing"
)

func Test_runMain(t *testing.T) {
	tests = Testing{
		A: 7,
		B: 7,
		C: -1,
		D: 3,
		E: 1,
	}
	got1 := runMain(tests)
	want1 := 0

	if got1 != want1 {
		t.Error("Expected", got1, "wanted", want1)
	}
}

func Test_add(t *testing.T) {
	got := add(3, 4)
	want := 7
	if got != want {
		t.Error("Expected", got, "wanted", want)
	}
}

func Test_addNamedVariables(t *testing.T) {
	got1, got2 := addNamedVariables(3, 4)
	want1, want2 := 7, -1
	if got1 != want1 {
		t.Error("Expected", got1, "wanted", want1)
	}
	if got2 != want2 {
		t.Error("Expected", got2, "wanted", want2)
	}
}

func Test_divide(t *testing.T) {
	got1, got2 := divide(9, 3)
	want1, want2 := 3, 0
	t.Run("divide_got1", func(t *testing.T) {
		if got1 != want1 {
			t.Error("Expected", got1, "wanted", want1)
		}
	})
	t.Run("divide_got2", func(t *testing.T) {
		if got2 != want2 {
			t.Error("Expected", got2, "wanted", want2)
		}
	})
}
