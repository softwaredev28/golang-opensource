package main

func square(n int) int {
	return n * n
}

// func process(num1, num2, num3, num4 int, cb func(int) int) {
// 	println(cb(num1))
// 	println(cb(num2))
// 	println(cb(num3))
// 	println(cb(num4))
// }

// * bisa juga menggunkan pariadic parameter. ini cara yang lebih efisien dan dinamis bisa dilempar argumen berapapun jumlahnya, tetapi harus disimpan diposisi terakhir pada parameter sebuah function. contohnya seperti dibawah ini
// * pariadic argument bisa diisi dengan sebuah slicing pada argumen, yang nantinya akan diterima oleh pariadic parameter
// * kalau pariadic parameter ini titik tiganya didepan nama parameternya, kalau pariadic argument, titik tiganya dibelakang parameter

func process(cb func(int) int, nums ...int) {
	for _, data := range nums {
		println(cb(data))
	}
}

func second_task() {
	// num1 := 1
	// num2 := 2
	// num3 := 3
	// num4 := 4
	// process(num1, num2, num3, num4, square) // non-pariadic
	// process(square, num1, num2, num3, num4) // pariadic function

	slice := []int{1, 2, 3, 4}
	process(square, slice...) // variadic argument
	// output
	// 1
	// 4
	// 9
	// 16
}
