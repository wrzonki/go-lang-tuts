package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(3, 4)
	want := 7
	if got != want {
		t.Error("Expected", got, "wanted", want)
	}
}

func TestAddNamedVariables(t *testing.T) {
	got1, got2 := addNamedVariables(3, 4)
	want1, want2 := 7, -1
	t.Run("addNamedVariables_got1", func(t *testing.T) {
		if got1 != want1 {
			t.Error("Expected", got1, "wanted", want1)
		}
	})
	t.Run("addNamedVariables_got2", func(t *testing.T) {
		if got2 != want2 {
			t.Error("Expected", got2, "wanted", want2)
		}
	})
}
