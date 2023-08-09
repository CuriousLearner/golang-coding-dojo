package main

import (
	"testing"
)

func TestNoNegativeValues(t *testing.T) {
	msg, err := FizzBuzz(-1)
	want := "Negative values not allowed"
	if err != want {
		t.Fatalf(`FizzBuzz(-1) = %q, %v, want match for "", %#q`, msg, err, want)
	}
}

func TestNoValueGreaterThan15(t *testing.T) {
	msg, err := FizzBuzz(16)
	want := "Values greater than 15 are not allowed"
	if err != want {
		t.Fatalf(`FizzBuzz(16) = %q, %v, want match for "", %#q`, msg, err, want)
	}
}

func TestFizzForMultiplesOfThree(t *testing.T) {
	msg, err := FizzBuzz(3)
	want := "1\n2\nFizz\n"
	if msg != want {
		t.Fatalf(`FizzBuzz(3) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestFizzForMultiplesOfFive(t *testing.T) {
	msg, err := FizzBuzz(5)
	want := "1\n2\nFizz\n4\nBuzz\n"
	if msg != want {
		t.Fatalf(`FizzBuzz(5) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestFizzForMultiplesOfThreeAndFive(t *testing.T) {
	msg, err := FizzBuzz(15)
	want := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n"
	if msg != want {
		t.Fatalf(`FizzBuzz(15) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
