package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    // validate arg of file path 
    if len(os.Args[1:])==0 {
        fmt.Println("Not found argument of path of file!")
        return
    }

    // read file and return array of facets
    arr, err := parseFileStl(os.Args[1:][0])
    if err != nil {
        fmt.Println("Error when opening STL file:", err)
        return
    }
    numberTriangles := len(arr)

    //calculate surface of all triangles in array
    start := time.Now().UnixNano() / int64(time.Millisecond)
    surfaceTriangles := calculateSurfaceFacets(arr)
    end := time.Now().UnixNano() / int64(time.Millisecond)
    diff := end - start
    
    // display result
    fmt.Println("\n*** without goroutines ***")
    fmt.Println("number of triangles in the file: ", numberTriangles)
    fmt.Println("surface area of the 3D part: ", surfaceTriangles)
    fmt.Println("miliseconds: ", diff)

    //calculate surface of all triangles in array
    start = time.Now().UnixNano() / int64(time.Millisecond)
    surfaceTriangles = calculateSurfaceFacetsWithGoroutine(arr)
    end = time.Now().UnixNano() / int64(time.Millisecond)
    diff = end - start
    
    // display result
    fmt.Println("\n*** with goroutines ***")
    fmt.Println("number of triangles in the file: ", numberTriangles)
    fmt.Println("surface area of the 3D part: ", surfaceTriangles)
    fmt.Println("miliseconds: ", diff)

}