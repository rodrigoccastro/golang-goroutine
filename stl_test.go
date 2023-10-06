package main

import "testing"

func TestStl_Integration(t *testing.T) {

	t.Run("Validate file example", func(t *testing.T) {
		arr, _ := parseFileStl("example.stl")
		validateTriangles(t,"example.stl", 2, len(arr))

		surfaceTriangles := calculateSurfaceFacets(arr)
		validateSurface(t,"example.stl", 1.4142, surfaceTriangles)
		
		surfaceTriangles = calculateSurfaceFacetsWithGoroutine(arr)
		validateSurface(t,"example.stl", 1.4142, surfaceTriangles)
	})

	t.Run("Validate file Moon", func(t *testing.T) {
		arr, _ := parseFileStl("Moon.stl")
		validateTriangles(t,"Moon.stl", 116, len(arr))

		surfaceTriangles := calculateSurfaceFacets(arr)
		validateSurface(t,"Moon.stl", 7.7726, surfaceTriangles)
		
		surfaceTriangles = calculateSurfaceFacetsWithGoroutine(arr)
		validateSurface(t,"Moon.stl", 7.7726, surfaceTriangles)
	})

}

func validateTriangles(t *testing.T, name string, want int, got int) {
	t.Run("Validate number triangles of file "+name, func(t *testing.T) {
		if want != got {
			t.Fatalf("Number of triangles is %d , but expected %d", got, want)
		}
	})
}

func validateSurface(t *testing.T, name string, want float64, got float64) {
	t.Run("Validate calculate surface of file "+name, func(t *testing.T) {
		if want != got {
			t.Fatalf("Number of triangles is %f , but expected %f", got, want)
		}
	})
}
