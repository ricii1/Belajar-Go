package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. If
	num := 10
	fmt.Println("Number is:", num)
	if num > 5 {
		fmt.Println("num is greater than 5")
	}

	// 2. If-Else
	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	// 3. If-Else If
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// 4. Switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// 5. Switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	// 6. For Loop
	fmt.Println("For Loop Example:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// 7. Looping through a collection
	fmt.Println("\nLooping through a collection:")
	rvariable := []string{"a", "b", "c", "d", "e"}

	for i, j := range rvariable {
		fmt.Println(i, j)
	}

	// 8. Defer
	// Defer bekerja dengan cara menunda eksekusi statement sampai fungsi selesai dieksekusi
	// Defer mempermudah dan meningkatkan keterbacaan kode serta memastikan bahwa proses clean-up dilakukan
	fmt.Println("\nDefer Example:")
	defer fmt.Println("This will be printed last (deferred)!")
	fmt.Println("This will be printed before the deferred statement")

	// 9. Stacking defers
	for i := 0; i < 10; i++ {
		defer fmt.Println("Defer:", i)
	}

	// Contoh defer pada fungsi. Defer tetap akan dilakukan meskipun ada error
	fmt.Println("\nUsing defer in a function:")
	fmt.Println("Result of calculation:", safeDivide(10, 2))
	fmt.Println("Attempt to divide by zero:")
	fmt.Println("Result of calculation:", safeDivide(10, 0), "\n")
}

// Contoh defer dalam fungsi
func safeDivide(a, b int) int {
	defer fmt.Println("safeDivide function execution completed")
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}
