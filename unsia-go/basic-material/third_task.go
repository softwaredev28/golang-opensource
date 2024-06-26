package main

import "fmt"


func main() {
	var arr = [5]int{10,20,30,40,50}
	// Membuat array
	fmt.Println("Array:", arr)

	var slice = arr[1:4]
	// Membuat slice dari array
	fmt.Println("Slice:", slice)

	// Menggunakan fungsi variadik untuk menjumlahkan elemen-elemen dalam slice
	total := sum(slice...)
	fmt.Println("Sum of slice elements:", total)

	// Menggunakan pointer untuk mengubah nilai dalam array
	// Pointer ke elemen ketiga dalam array
	var p = &arr[2]
	*p = 100     // Mengubah nilai elemen ketiga
	fmt.Println("Array setelah diubah melalui pointer:", arr)
}

func sum(x ...int)int {
	var total int
	for _, data := range x {
		total += data
	}
	return total
}

//* output :
//* Array: [10 20 30 40 50]
//* Slice: [20 30 40]
//* Sum of slice elements: 90
//* Array setelah diubah melalui pointer: [10 20 100 40 50]