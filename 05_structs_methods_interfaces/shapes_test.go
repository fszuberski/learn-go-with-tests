package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 62.83185307179586)
	})

	t.Run("triangles", func(t *testing.T) {
		triangle := Triangle{10.0, 6.0}
		checkPerimeter(t, triangle, -1)
	})
}

func TestArea(t *testing.T) {
	// table tests!
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v: got %g, want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
