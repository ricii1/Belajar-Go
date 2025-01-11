package main

// Cara import single package
// import "fmt"
// import "reflect"

// Cara import multiple packages
import (
	"fmt"
	"reflect"
)

func main() {
	// Deklarasi variabel menggunakan 'var'
	var a int = 10
	var b string = "Hello, Go!"

	// Menggunakan ':=' untuk deklarasi dan inisialisasi secara otomatis
	c := 20
	d := "Welcome to Go"
	// Bisa juga seperti
	// c, d := 20, "Welcome to Go"

	// Deklarasi konstanta menggunakan 'const'
	const e = 100
	const f = "This is a constant"

	// Menampilkan variabel
	fmt.Println("Using 'var' for declaration and initialization:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Menampilkan variabel dengan ':='
	fmt.Println("\nUsing ':=' for declaration and initialization:")
	fmt.Println("c:", c, "Type:", reflect.TypeOf(c))
	fmt.Println("d:", d, "Type:", reflect.TypeOf(d))

	// Menampilkan nilai konstanta
	fmt.Println("\nUsing 'const' for declaring constants:")
	fmt.Println("e:", e)
	fmt.Println("f:", f)

	// Perbedaan penggunaan ':=' dan '='
	// ':=' digunakan saat pertama kali mendeklarasikan dan menginisialisasi variabel,
	// sedangkan '=' digunakan untuk assignment pada variabel yang sudah ada.

	// Assignment menggunakan '=' pada variabel yang sudah ada
	a = 15
	b = "Updated value"

	// Menampilkan hasil assignment
	fmt.Println("\nAfter using '=' to modify existing variables:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// String concatenation
	fullString := "Go " + "is " + "awesome!"
	fmt.Println("\nString Concatenation Example:")
	fmt.Println(fullString)

	// String formatting menggunakan fmt.Printf
	// Menggunakan placeholder %d untuk integer dan %s untuk string
	fmt.Println("\nString Formatting Example:")
	fmt.Printf("Value of a: %d, Value of b: %s\n", a, b)
	fmt.Printf("Concatenated string: %s\n", fullString)
	/*
		Jalankan dengan perintah 'go run 1.go' atau 'go build 1.go' lalu jalankan file executable yang dihasilkan.
	*/
}
