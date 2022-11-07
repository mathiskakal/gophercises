package main

import "testing"

func TestArea(t *testing.T) {

	// CheckArea Helper Function
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
	})
	/*
		t.Run("rectangles", func(t *testing.T) {

			rectangle := Rectangle{12.0, 6.0}

			got := rectangle.Area()
			want := 72.0

			if got != want {
				t.Errorf("got %g want %g", got, want)
			}
		})
		t.Run("circles", func(t *testing.T) {

			circle := Circle{10}

			got := circle.Area()
			want := 314.1592653589793

			if got != want {
				t.Errorf("got %g want %g", got, want)
			}
		})*/
}

// Write Here

/*
func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// the %.2f stands for two decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
*/
