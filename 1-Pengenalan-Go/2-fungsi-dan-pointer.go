package main

import "fmt"

// fungsi pada go dibuat dengan cara seperti di bawah ini
// func nama_fungsi(parameter1 tipe_data, parameter2 tipe_data, ...) tipe_data_return
func add(x int, y int) int {
	return x + y
}

// penulisan seperti di bawah ini dapat menyingkat penulisan (jika tipe data sama)
func substract(x, y int) int {
	return x - y
}

// fungsi dengan pointer
func div(x, y int, z *float32) {
	*z = float32(x) / float32(y)
}

// fungsi dengan multiple return
func addSubstract(x, y int) (int, int) {
	return x + y, x - y
}

func adder() func(int) int {
	sum := 0
	calculation := func(x int) int {
		sum += x
		return sum
	}
	return calculation
}

func main() {
	// Contoh fungsi
	fmt.Println("Function example:")
	fmt.Println("add(10,5):", add(10, 5))
	// Contoh fungsi dengan multiple return
	fmt.Println("\nMultiple return example:")
	addition, substraction := addSubstract(10, 5)
	fmt.Println("addSubtract(10,5):", addition, substraction)
	// Contoh fungsi dengan pointer
	fmt.Println("\nFunction with Pointer example:")
	x := 10
	y := 5
	var z float32
	div(x, y, &z)
	fmt.Println("div(x,y,&z): ", z)

	// Closure Function
	// Closure adalah fungsi yang mengembalikan fungsi lain
	// Fungsi yang dikembalikan ini dapat mengakses variabel yang dideklarasikan di dalam fungsi yang mengembalikan
	fmt.Println("\nClosure Function example:")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
