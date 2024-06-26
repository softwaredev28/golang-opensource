package main

import "fmt"

//*  cpu tidak suka dengan yang namanya function
//* function akan berguna ketika sudah menjadi sesuatu yang rutin
//* manusia tidak bisa membuatnya manual semua, jadi harus membutuhkan function
//* apabila tipe pada parameter function itu sama maka cukup deklarasikan tipenya di parameter yang terakhir pada function/tapi bisa juga dideklarasikan tipenya masing" untuk setiap parameternya
//* parameternya saat deklarasi (a int, b, int), saat memanggil itu adalah argument nama_fungsi(3,5)

// ? materi : pointer
// * pointer adalah isinya sebuah alamat/address dari suatu nilai/berperan sebagai petunjuk ke variable lain
// * kalau di codenya variable ini yang dikenali nilainya, tetapi kalau di balik layar yang dikenali dan disimpan adalah addressnya
// * jadi kita bisa mengakses nilai yang disimpan didalam pointer bisa menggunakan tanda * didepan nama variable pointernya (*p)
// * ketika deklarasi disetnya sebagai pointer, ketika diinisiasi sebuah nilai lain maka itu akan error


type myStr string

type myStruct struct {
	code int
	nama string
}

type product struct {
	code int
	product_name string
}

func thrid_meet() {	
	
// ! POINTER
// 	var i int
// 	var p *int
// 	i = 20
// 	p = &i
// 	q := &i
// 	// i := 10
// 	// p := &i       // menunjuk ke i
// 	// println(*p)     // baca i lewat pointer
// 	// *p = 20       // set i lewat pointer
// 	// println(i)      // lihat nilai terbaru dari i

// 	println(i)
// 	println(p)
// 	println(q)

// 	println("\n")

// 	println(i)
// 	println(*p)
// 	println(*q)
// // !

// ! ARRAY

	// var arr1 [3]int = [3]int{1,2,3}  // explicit
	// arr1 := [3]int{1,2,3} // implicit
	// var arr1 [3]int
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 4

	// var arr1 [3]bool
	// arr1[0] = true
	// arr1[1] = false
	// arr1[2] = true

	// var arr1 [4]int = [...]int{10,11,12,13}
	// arr1[1] = 20
	// fmt.Println(arr1)
	// fmt.Println(arr1[3])
	// println(len(arr1), cap(arr1))

	// println("\n")
	
	// * cap dan len akan terlihat berbeda ketika melakukan slicing
	// * panjang data dan kapasitas data berbeda
	
	// ! SLICING
	
	// * slicing ini hanya akan membuat pointer saja, dia tidak akan menyimpan sebuah memory baru. hanya merujuk saja ke array sebelumnya
	// * kalau kita punya array kemudian di slicing, ketika kita merubah data yang di dislicing, maka itu akan berdampak ke data array aslinya

	// slice1 := arr1[:]
	// fmt.Println(arr1)
	// fmt.Println(slice1)
	// println(len(slice1), cap(slice1))

	// println("\n")

	// slice1[1] = 18
	// fmt.Println(arr1)
	// fmt.Println(slice1)

	// println("\n")

	// * slice setelah di append akan berdampak ke array aslinya, ini menandakkan bahwa append() ini bertipe pointer juga
	// * append mempunyai 2 logic, jika penambahan membernya melebihi dari kapasitasnya maka dia akan membuat alamat/pointer baru/bukan rujukan dari array sebelumnya

	// slice1[0] = 22
	// slice1[1] = 12
	// slice1 = append(slice1, 99,8,29)
	// fmt.Println(arr1)
	// fmt.Println(slice1)

	// println("\n")

	// * kalau array harus diberikan capacitynya, kalau slicing tidak perlu ditambahkan capacitynya dia dinamis
	// * append() memiliki kemampuan untuk menambahkan data ketika melebihi capacity
	// var arr2 []int
	// ! panic: runtime error: index out of range [1] with length 0
	// arr2[1] = 100
	// arr2[3] = 99
	// fmt.Println(arr2)

	// arr2 = append(arr2, 12,232,234,22,65,24)
	// fmt.Println(arr2)
	
	// var arr3 []int = []int{10,2,30,4,20,5}
	// arr2 = append(arr2, arr3...)
	// * fungsi ... untuk mengubah slicing menjadi argumen. aslinya sebagai berikut
	// * arr2 = append(arr2, 10,2,30,4,20,5)
	// fmt.Println(arr2, arr3)

	// println("\n")

	// ! STRUCT
	// * array digunakan kalau kita ingin mengcontainer yang tipe datanya sama
	// * struct dipakai kalau kita ingin mengcontainer tipe data yang berbeda-beda
	var a myStr = "nama saya ketsar"
	println(a)

	var b myStruct = myStruct{code: 10, nama: "Ahmad"}
	fmt.Println(b)

	println("\n")

	// ! METHOD
	// * method yang diberikan receiver
	// * sebuah tipe data bisa diberi fungsi yang disebut sebagai method
	// * tipe yang bisa diberi receiver pada sebuah method adalah tipe lokal/tipe yang kita buat sendiri
	printOnly("Ketsar")
	a.cetakAja("Ketsar")

	// c := &a
	// c.code = 101
	// c.product_name = "Pensil"
	var c product = product{code: 10, product_name: "Pensil"}
	c.cetakLagi()
	
	d := &c
	d.code = 29
	d.product_name = "Pensil Hijau"
	d.cetakLagiDenganPointer()
	c.cetakLagiDenganPointer()
}

func printOnly(str string) {
	println("method printOnly =",str)
}

func (n myStr) cetakAja(str string) {
	println("method cetakAja =",str,n)
}

func (p product) cetakLagi() {
	println("method cetakLagi =",p.code, p.product_name)
}

func (p product) cetakLagiDenganPointer() {
	println("method cetakLagiDenganPointer =",p.code, p.product_name)
}