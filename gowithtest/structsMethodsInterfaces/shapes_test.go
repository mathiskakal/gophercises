package main

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

	// CheckArea Helper Function
	/*
		checkArea := func(t testing.TB, shape Shape, want float64) {
			t.Helper()
			got := shape.Area()
			if got != want {
				t.Errorf("got %g, want %g", got, want)
			}
		}

		t.Run("rectangles", func(t *testing.T) {
			rectangle := Rectangle{12, 6}
			checkArea(t, rectangle, 72.0)
		})

		t.Run("circles", func(t *testing.T) {
			circle := Circle{10}
			checkArea(t, circle, 314.1592653589793)
		})*/
}

// Write Here
