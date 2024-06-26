package main

import (
	"fmt"
	"unsia/lib"
)

func first_meet() {
	println("hello world")

	println("-----------------------------------")

	fmt.Println(lib.GetArray())

	println("-----------------------------------")

	var a int32 = 128
	var b float64 = 32.2
	var c string = "ketsar"
	var d rune = 'K'
	var e bool = true
	var f = "mantap"
	var g int32 = 10

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)

	println("-----------------------------------")

	if g != 10 {
		println("ini benar g bukan 10")
	} else {
		println("g bernilai 10")
	}

	println("-----------------------------------")

	if e && g < a {
		println("kondisi terpenuhi")
	} else {
		println("kondisi tidak terpenuhi")
	}

	println("-----------------------------------")

	fmt.Println("Hasil")

	if true {
		fmt.Println("condition is true")
	}

	println("-----------------------------------")

	if k := 12; k > 15 {
		println("betul k lebih kecil dari pada 15")
	} else if k > 5 && k < 13 {
		println("yahhh.. k lebih besar dari pada 5 dan k lebih kecil dari pada 13")
	} else {
		println("saya tidak paham maksud anda")
	}

	println("-----------------------------------")

	for idx := 1; idx <= 5; idx++ {
		println("1). data ke ", idx)
	}

	println("-----------------------------------")

	for idx := 1; idx <= 5; {
		println("2). data ke ", idx)
		idx++ 
	}

	println("-----------------------------------")

	idx := 1
	for ; idx <= 5; {
		println("3). data ke ", idx)
		idx++ 
	}
}