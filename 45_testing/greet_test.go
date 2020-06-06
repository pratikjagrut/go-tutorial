package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("greetNil", func(t *testing.T) {
		greetNil(t)
	})

	t.Run("greetName", func(t *testing.T) {
		greetName(t)
	})
}

func greetNil(t *testing.T) {
	say := greet("")

	if say != "Hello, World!" {
		t.Errorf("greet(nil) failed, expected: Hello, World!, Got: %s", say)
	} else {
		t.Logf("greet(nil) succesfull, expected: Hello, World!, Got: %s", say)
	}
}

func greetName(t *testing.T) {
	say := greet("Johny")

	if say != "Hello, Johny" {
		t.Errorf("greet(name) failed, expected: Hello, Johny, Got: %s", say)
	} else {
		t.Logf("greet(nil) succesfull, expected: Hello, Johny, Got: %s", say)
	}
}
