package main

import(
	"testing"
)

func TestHello(t *testing.T){
	got := Hello("Chris")
	expected := "Hello, Chris"
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}

}