package main

// go build -gcflags '-m -l'
// -gcflags to view the results of the escape analysis mentioned above in output from the compiler.
// -m to print optimization decisions
// -l to omit inlining decisions

const numberOfItems = 20_000_000

type bigStruct struct {
	A, B, C int
	D, E, F string
	G, H, I bool
}

// moved to heap: element
// make([]*bigStruct, numberOfItems) escapes to heap
//
//go:noinline
func fillSliceWithAppend() []*bigStruct {
	data := make([]*bigStruct, 0, numberOfItems)

	element := bigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}

	for i := 0; i < numberOfItems; i++ {
		data = append(data, &element)
	}

	return data
}

// moved to heap: element
// make([]*bigStruct, 0, numberOfItems) escapes to heap
//
//go:noinline
func fillSliceWithIndex() []*bigStruct {
	data := make([]*bigStruct, numberOfItems)

	element := bigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}

	for i := 0; i < numberOfItems; i++ {
		data[i] = &element
	}

	return data
}
