package main

import (
	"math"
	"sync"
)

func calculateSurfaceFacets(arr []Facet) float64 {
    var tot float64

    for i := 0; i < len(arr); i++ {
        tot = tot + calculateSurfaceVertex(arr[i].V1,arr[i].V2,arr[i].V3)
	}

    //format for 4 decimais
    return format(tot)
}


// number of max goroutines in parallel
const MaxParallel = 100

// for resolve problem of memory
var mut sync.Mutex

func calculateSurfaceFacetsWithGoroutine(arr []Facet) float64 {
	var tot float64

    // control the max number of goroutines in parallel
	var chnMaxParallel = make(chan struct{}, MaxParallel)

    // for wait all goroutines finish
	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		// add more one item in group
        wg.Add(1)

		go func(i int) {
			chnMaxParallel <- struct{}{} // acquire token
			//execute when finished method
            defer wg.Done()
			defer func() { <-chnMaxParallel }() //release token

			v := calculateSurfaceVertex(arr[i].V1, arr[i].V2, arr[i].V3)
			mut.Lock() // lock for avoid problem of memory
			tot = tot + v
			mut.Unlock()
		}(i)
	}

    // wait until all goroutines is finished
	wg.Wait()

	//format for 4 decimais
	return format(tot)
}

func format(value float64) float64 {
   return math.Round(value*10000) / 10000
}

func calculateSurfaceVertex(v1, v2, v3 Vertex) float64 {

    // Calculate the lengths of the triangle's sides
    a := math.Sqrt(math.Pow(v2.X-v1.X, 2) + math.Pow(v2.Y-v1.Y, 2) + math.Pow(v2.Z-v1.Z, 2))
    b := math.Sqrt(math.Pow(v3.X-v2.X, 2) + math.Pow(v3.Y-v2.Y, 2) + math.Pow(v3.Z-v2.Z, 2))
    c := math.Sqrt(math.Pow(v1.X-v3.X, 2) + math.Pow(v1.Y-v3.Y, 2) + math.Pow(v1.Z-v3.Z, 2))

    // Calculate the semiperimeter
    s := (a + b + c) / 2

    // Calculate the area using Heron's formula
    return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}