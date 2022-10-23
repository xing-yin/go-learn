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

func Test_Area(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "rectangle",
			shape: &Rectangle{Width: 10.0, Height: 15.0},
			want:  150.0,
		},
		{
			name:  "circle",
			shape: &Circle{Radius: 10.0},
			want:  314.1592653589793,
		},
		{
			name: "triangle",
			shape: &Triangle{
				Base:   12.0,
				Height: 6.0,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.shape.Area(); got != tt.want {
				t.Errorf("Area() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}
