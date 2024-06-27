package main

/*
? === APA ITU gRPC ? ===
* gRPC adalah framework RPC yang modern
* opensource
* performace tinggi
* mempunyai banyak fitur
* awalnya dibangun oleh google
* sekarang menjadi sebagian dari CNCF (sebuah organisasi yang berfokus pada ekosistem cloud - bagian dari linux foundation)
* gRPC : g mempunyai arti yang berbeda-beda pada setiap rilisnya, seperti : good, green, etc
* gRPC = RPC adalah Remote Procedure Call
* banyak di terapkan untuk komunikasi antar service

? PROTOKOL API
* SOAP
* RPC (Remote Procedure Call)
* REST API

? === APA ITU RPC ===
* protokol yang mengizinkan program untuk mengeksekusi sebuah prosedur yang ada diprogram lain(service dan server lain)
* implementasi di kode, client tinggal memanggil fungsi di server lain secara langsung
* kode di client dan di server bisa menggunakan bahasa pemrograman yang berbeda
* singkatnya :  memanggil fungsi ditempat lain secara remote

? === BEDANYA gRPC vs REST ===
* protokol =>  gRPC : HTTP/2 & HTTP/3 | REST : HTTP/1.1 & HTTP/2
* paylod => gRPC : Protobuf (binary) | REST : JSON (text)
* kontrak => gRPC : Strict(.proto) | REST : Loose (OpenAPI)
* Generator Kode => gRPC : Built-in(protoc) | REST : Thrid-party (swagger)
* Keamanan => gRPC : TLS/SSL | REST : TLS/SSL
* Streaming 2 Arah => gRPC : YA | REST : TIDAK
* Browser Support => gRPC : TIDAK | REST : YA

* secara performance gRPC lebih powerful daripada REST, oleh karena itu ketemu yang namanya arsitektur microservices
* kekurangan microservices : ada 2 service yang berkomunikasi, untuk service tersebut dapat berkomunikasi maka dibutuhkanlah sebuah API , API musuhnya adalah latensi, kalau diaplikasi yang sama tidak ada latensi

* membuat file .proto dikenalkan dua hal :
* message & kontrak
*/
// func grpc() {

// }