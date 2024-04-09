package main

import (
	"helloworld/controllers"
	"testing"
)

func TestNewString(t *testing.T) {
	if controllers.NewString("Hello, World! ") != "Hello, World! 12345"{
		t.Error("Not the expected test output")
	}
}