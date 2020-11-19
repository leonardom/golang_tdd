package structs

import (
	"testing"
)

var checkArea = func(t *testing.T, shape Shape, expected float64) {
	t.Helper()
	got := shape.Area()

	if got != expected {
		t.Errorf("%#v got %g expected %g", shape, got, expected)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTable(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, expected: 72.0},
		{name: "Cicle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.expected)
		})
	}
}
