package struct_function_interface

import "testing"

func TestPerimeter(t *testing.T) {
	r := &Rectangle{
		Width:  10.0,
		Height: 15.0,
	}
	got := r.Perimeter()
	want := 50.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		r := &Rectangle{
			Width:  10.0,
			Height: 15.0,
		}
		want := 150.0
		checkArea(t, r, want)
	})

	t.Run("circle", func(t *testing.T) {
		c := &Circle{
			Radius: 10.0,
		}
		want := 314.1592653589793

		checkArea(t, c, want)
	})
}
