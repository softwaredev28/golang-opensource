package main

type customeType func(int) bool
type customeType2 func(int, int) bool

func second_meet() {
	// ! bilangan prima
	// ! 1 & dirinya sendiri

	// num := 11
	// if num == 1 {
	// 	return
	// }
	// if num == 2 {
	// 	print(2)
	// 	return
	// }

	// for i := 2; i < num; i++ {
	// 	if num%i == 0 {
	// 		return
	// 	}
	// }

	// println(num)
	// return

	// ! melakukan optimasi bilangan prima
	// num := 13
	// if num <= 1 {
	// 	fmt.Println(num, "bukan bilangan prima")
	// 	return
	// }

	// if num == 2 {
	// 	fmt.Println(num, "merupakan bilangan prima")
	// 	return
	// }

	// if num % 2 ==0 {
	// 	fmt.Println(num, "bukan bilangan prima")
	// 	return
	// }
	// for j := 3; j <= num/2; j++ {
	// 	if num%j == 0 {
	// 		break
	// 	}

	// }
	// fmt.Println(num, "merupakan bilangan prima")
	// return

	// ! dengan function
	// num := 13
	// println("berikut merupakan bilangan prima : ")
	// for i := 1; i <= num; i++ {
	// 	if isPrime(i) {
	// 		println(i)
	// 	}
	// }

	// ! anonymous function
	// num := 4
	// for i := 1; i <= num; i++ {
	// 	if func(num int) bool {
	// 		if num <= 1 {
	// 			return false
	// 		}

	// 		if num == 2 {
	// 			return true
	// 		}

	// 		if num%2 == 0 {
	// 			return false
	// 		}
	// 		for j := 3; j <= num/2; j++ {
	// 			if num%j == 0 {
	// 				return false
	// 			}

	// 		}
	// 		return true
	// 	}(num) {
	// 		println(num, "merupakan bilangan prima")
	// 		break
	// 	}
	// }

	// ! closure
	// isPrime := func(num int) bool {
	// 	if num <= 1 {
	// 		return false
	// 	}

	// 	if num == 2 {
	// 		return true
	// 	}

	// 	if num%2 == 0 {
	// 		return false
	// 	}
	// 	for j := 3; j <= num/2; j++ {
	// 		if num%j == 0 {
	// 			return false
	// 		}

	// 	}
	// 	return true
	// }

	// println("berikut angka yang merupakan bilangan prima : ")
	// for i := 1; i < 13; i++ {
	// 	if isPrime(i) {
	// 		print(i, " ")
	// 	}
	// }

	// ! function typr
	var isPrime customeType = func(num int) bool {
		if num <= 1 {
			return false
		}

		if num == 2 {
			return true
		}

		if num%2 == 0 {
			return false
		}
		for j := 3; j <= num/2; j++ {
			if num%j == 0 {
				return false
			}

		}
		return true
	}

	println("berikut angka yang merupakan bilangan prima : ")
	for i := 1; i < 13; i++ {
		if isPrime(i) {
			print(i, " ")
		}
	}
}

// func isPrime(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}

// 	if num == 2 {
// 		return true
// 	}

// 	if num%2 == 0 {
// 		return false
// 	}
// 	for j := 3; j <= num/2; j++ {
// 		if num%j == 0 {
// 			return false
// 		}

// 	}
// 	return true
// }