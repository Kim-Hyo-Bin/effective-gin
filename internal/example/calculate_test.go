package example

import (
	"math"
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	testCases := []struct {
		name     string
		rect     Rectangle // Use exported type name
		expected float64
	}{
		{"standard rectangle", Rectangle{Length: 10, Width: 5}, 50.0}, // Use exported field names
		{"square", Rectangle{Length: 7, Width: 7}, 49.0},
		{"zero width", Rectangle{Length: 10, Width: 0}, 0.0},
		{"zero length", Rectangle{Length: 0, Width: 5}, 0.0},
		{"zero dimensions", Rectangle{Length: 0, Width: 0}, 0.0},
		{"decimal dimensions", Rectangle{Length: 2.5, Width: 4}, 10.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the exported method (starts with uppercase)
			actual := tc.rect.Area() // Corrected: area() -> Area()
			if actual != tc.expected {
				// Use exported field names in error message if needed
				t.Errorf("Area() for Rectangle{Length: %f, Width: %f} = %f; want %f", tc.rect.Length, tc.rect.Width, actual, tc.expected)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	testCases := []struct {
		name     string
		rect     Rectangle // Use exported type name
		expected float64
	}{
		{"standard rectangle", Rectangle{Length: 10, Width: 5}, 30.0}, // Use exported field names
		{"square", Rectangle{Length: 7, Width: 7}, 28.0},
		{"zero width", Rectangle{Length: 10, Width: 0}, 20.0},
		{"zero length", Rectangle{Length: 0, Width: 5}, 10.0},
		{"zero dimensions", Rectangle{Length: 0, Width: 0}, 0.0},
		{"decimal dimensions", Rectangle{Length: 2.5, Width: 4}, 13.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the exported method (starts with uppercase)
			actual := tc.rect.Perimeter() // Corrected: perimeter() -> Perimeter()
			if actual != tc.expected {
				// Use exported field names in error message if needed
				t.Errorf("Perimeter() for Rectangle{Length: %f, Width: %f} = %f; want %f", tc.rect.Length, tc.rect.Width, actual, tc.expected)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	testCases := []struct {
		name     string
		circ     Circle // Use exported type name
		expected float64
	}{
		{"standard circle", Circle{Radius: 10}, math.Pi * 100}, // Use exported field name
		{"small circle", Circle{Radius: 1}, math.Pi},
		{"zero radius", Circle{Radius: 0}, 0.0},
		{"decimal radius", Circle{Radius: 2.5}, math.Pi * 6.25},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the exported method (starts with uppercase)
			actual := tc.circ.Area() // Corrected: area() -> Area()
			if math.Abs(actual-tc.expected) > 1e-9 {
				// Use exported field name in error message if needed
				t.Errorf("Area() for Circle{Radius: %f} = %f; want %f", tc.circ.Radius, actual, tc.expected)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	testCases := []struct {
		name     string
		circ     Circle // Use exported type name
		expected float64
	}{
		{"standard circle", Circle{Radius: 10}, 2 * math.Pi * 10}, // Use exported field name
		{"small circle", Circle{Radius: 1}, 2 * math.Pi},
		{"zero radius", Circle{Radius: 0}, 0.0},
		{"decimal radius", Circle{Radius: 2.5}, 2 * math.Pi * 2.5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the exported method (starts with uppercase)
			actual := tc.circ.Perimeter() // Corrected: perimeter() -> Perimeter()
			if math.Abs(actual-tc.expected) > 1e-9 {
				// Use exported field name in error message if needed
				t.Errorf("Perimeter() for Circle{Radius: %f} = %f; want %f", tc.circ.Radius, actual, tc.expected)
			}
		})
	}
}

// Optional: Test using the interface
func TestShapeCalculations(t *testing.T) {
	shapes := []struct {
		name          string
		s             shape // Interface type is fine
		expectedArea  float64
		expectedPerim float64
	}{
		// Use exported types and fields when creating concrete instances
		{"rectangle", Rectangle{Length: 4, Width: 6}, 24.0, 20.0},
		{"circle", Circle{Radius: 5}, math.Pi * 25, 2 * math.Pi * 5},
	}

	for _, tc := range shapes {
		t.Run(tc.name+" area", func(t *testing.T) {
			// Call methods via the interface (still uses the underlying exported methods)
			actualArea := tc.s.Area() // Corrected: area() -> Area()
			if math.Abs(actualArea-tc.expectedArea) > 1e-9 {
				t.Errorf("Area() = %f; want %f", actualArea, tc.expectedArea)
			}
		})
		t.Run(tc.name+" perimeter", func(t *testing.T) {
			// Call methods via the interface (still uses the underlying exported methods)
			actualPerim := tc.s.Perimeter() // Corrected: perimeter() -> Perimeter()
			if math.Abs(actualPerim-tc.expectedPerim) > 1e-9 {
				t.Errorf("Perimeter() = %f; want %f", actualPerim, tc.expectedPerim)
			}
		})
	}
}
