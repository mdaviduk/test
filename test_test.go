package main

import "testing"

func TestSaysomthing(t *testing.T) {
	if saysomething() != "hello2" {
		t.Fail()
	}
}
