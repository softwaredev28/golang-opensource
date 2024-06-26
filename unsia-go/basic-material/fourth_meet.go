package main

/*
* kalau di challange dengan performa hindari penggunaan interface
* kalau di challange dengan fleksibiltas gunakan interface
* semua ada risk dan ada cost-nya, jadi harus di pertimbangkan karena ada banyak algoritma yang bisa digunakan


* variable local : variable yang hidupnya diblock function tertentu/hanya dapat di akses diblock scope tersebut
* variable global : variable yang siklus hidupnya diluar function/dapat diakses discope manapun

* karakteristik variable global : 1. ramah terhadap runtime, 2. tidak ramah dengan memory karena tidak dihapus variable tersebut
* variable global cost-nya akan membebani memory (tidak direkomendasikan), sebaiknya selalu membuat variable local karena hemat terhadap memory

* interface digunakan ketika kita ingin membuat fungsi" yang bisa di wrapping/dibungkus dengan tipe lain.
* interface kosong digunakan ketika kita tidak bisa mengontrol tipe datanya. kalau bisa dikontrol langsung saja strict/tuliskan tipe datanya
 */

// * =========== INTERFACE ================

//  * interface
type i interface {
	SomeMethod()
	OtherMethod()
}

// * nil interface
type nilInterface interface {

}


// * =========== IMPLEMENT ================
type myStr string

// anonymous function
func (data myStr) SomeMethod() {
	println(data, " - some method!")
}
func (data myStr) OtherMethod() {
	println(data, " - other method!")
}

func fourth_meet(){

/*
* == INTERFACE ==
? interface berisi kumpulan method yang abstract
? tidak punya real code/implementasi pada method di interface tersebut

* == NIL INTERFACE (INTERFACE KOSONG) ==
? dia sama seperti tipe any
? apa aja yang di assign ke nil interface itu dianggap sah
? karena dia tidak mempunyai method abstract 
? untuk mencetak nil interface harus di tambahkan .(tipedata)
? karena secara default yang dikembalikan hanya kode hexadecimal
? jadi harus di type aserting ke tipe data aslinya untuk mendapatkan datanya
*/ 

	var someInterface i
	str := myStr("Hello World!")
	// someInterface = &str // ! error : str tidak mengimplementasikan method abstract, jadi tidak dapat di gunakan. supaya bisa digunakan jadi harus dibuatkan implementasi terlebih dahulu
	someInterface = &str
	someInterface.SomeMethod() 
	someInterface.OtherMethod() 

	var someData nilInterface
	someData = false
	println(someData.(bool))

	someData = "Ketsar Ali!"
	println(someData.(string))

	someData = 1231
	println(someData.(int))


}