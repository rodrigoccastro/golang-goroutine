package main

type Vertex struct {
	X, Y, Z float64
}

type Facet struct {
	NI, NJ, NK float64
	V1, V2, V3 Vertex
}
