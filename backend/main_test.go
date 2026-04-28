package main

import (
	"testing"
)

func TestBasicMath(t *testing.T) {
	result := 2 + 2
	if result != 4 {
		t.Errorf("Expected 2+2 to be 4, but got %d", result)
	}
}

func TestHelloMessage(t *testing.T) {
	message := "Hello"
	if message != "Hello" {
		t.Errorf("Expected message to be 'Hello', but got %s", message)
	}
}

func TestMultiplication(t *testing.T) {
	result := 5 * 5
	if result != 25 {
		t.Errorf("Expected 5*5 to be 25, but got %d", result)
	}
}

func TestDivision(t *testing.T) {
	result := 10 / 2
	if result != 5 {
		t.Errorf("Expected 10/2 to be 5, but got %d", result)
	}
}

func TestStringConcatenation(t *testing.T) {
	s1 := "Hello"
	s2 := "World"
	result := s1 + " " + s2
	if result != "Hello World" {
		t.Errorf("Expected 'Hello World', but got %s", result)
	}
}

func TestSliceLength(t *testing.T) {
	items := []string{"apple", "banana", "cherry"}
	if len(items) != 3 {
		t.Errorf("Expected length 3, but got %d", len(items))
	}
}
