example with goroutine

• instructions for running your code
    //for run testes
    go test 

    //for execute examples
    go run . .\example.stl
    go run . .\Moon.stl
    go run . .\big.stl

• an explanation of the design decisions you made
    a) I created the entities.go file with the Facet and Vertex structs to store the information;
    b) I created the parse.go file with the sole responsibility of reading the file and returning an array of Facets with the triangle information;
    c) I created the file calculate.go with the sole responsibility of calculating the total area of an array of triangles;
    d) I created the main.go file that validates the input parameter and calls the two main functions, parse and calculate, and finally displays the result;
    e) I created the stl_test.go file that runs tests to validate the number of triangles and surface the area of the example.stl and Moon.stl files;

