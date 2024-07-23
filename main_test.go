package main

import (
	"reflect"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	x, y := 1, 2
	expected := 3

	if calculateValues(x, y) != expected {
		t.Errorf("Expected %d but got %d", expected, calculateValues(x, y))
	}
}

func TestEqualPlayers(t *testing.T) {
	expected := Player{name: "Mark", hp: 100}
	actual := Player{name: "Mark", hp: 100}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v but got %+v", expected, actual)
	}
}
