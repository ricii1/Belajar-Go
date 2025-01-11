package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// 1. Struct
	fmt.Println("Struct Example:")
	VertexExample := Vertex{1, 2}
	fmt.Println("Vertex:", VertexExample)
	fmt.Println("Vertex X:", VertexExample.X)
	fmt.Println("Vertex Y:", VertexExample.Y)
	VertexExample.X = 4
	fmt.Println("Vertex:", VertexExample)

	// 2. Array
	fmt.Println("\nArray Example:")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Primes:", primes)

	// 3. Slice
	// Slice adalah potongan dari array yang memiliki ukuran dinamis
	fmt.Println("\nSlice Example:")
	var s []int = primes[1:4]
	fmt.Println("Slice from primes:", s)
	fmt.Println("First index of the slice: ", s[0])
	s = primes[0:4]
	fmt.Println("Slice from primes:", s)
	fmt.Println("First index of the slice: ", s[0])

	// Slice hanya referensi ke array
	fmt.Println("\nSlice Example 2:")
	s[0] = 100
	fmt.Println("Slice from primes:", s)
	fmt.Println("Primes:", primes)

	// Slice literal
	// Slice literal adalah cara membuat slice tanpa array
	fmt.Println("\nSlice Literal Example:")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	// Slice bounds
	fmt.Println("\nSlice Bounds Example:")
	arr := [6]int{2, 3, 5, 7, 11, 13}

	sl := arr[:]
	fmt.Println(sl)

	sl = arr[1:4]
	fmt.Println(sl)

	sl = arr[:2]
	fmt.Println(sl)

	sl = arr[1:]
	fmt.Println(sl)

	sl = arr[:len(arr)-2]
	fmt.Println(sl)

	// 4. Map
	// Map adalah tipe data yang memetakan key ke value
	fmt.Println("\nMap Example:")
	var m = map[string]Vertex{
		"Bell Labs": {40, -75},
		"Google":    {37, -122},
	}
	fmt.Println("Map:", m)
	fmt.Println("Map Bell Labs:", m["Bell Labs"])
	fmt.Println("Map Google:", m["Google"])

	// Insert or update an element in map
	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])

	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
