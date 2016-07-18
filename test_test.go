package main

import "testing"

func TestSaysomthing(t *testing.T) {
	if saysomething() != "hello1" {
		t.Fail()
	}
}
