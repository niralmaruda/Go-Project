package main

import (
	"testing"
)

func TestParimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Parimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %0.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.hasArea
		if got != want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
