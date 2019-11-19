package main

import "fmt"

func main() {
	// -----------
	// Built in types
	// -----------
	//
	// Golang memiliki tipe data dengan tujuan
	// 1. Memberikan default memory untuk variabel
	// 2. Memberi informasi tipe data variabel
	//
	// Tipe data bisa spesifik:
	// Cth:  int8, int16, int32
	//
	// Ketika tidak menspesifikasikan tipe data secara detail,
	// akan dipetakan sesuai arsitektur os
	// Cth: 64 bit -> int menjadi int64
	//      32 bit -> int menjadi int32
	//
	// alias dari byte : uint8
	//       rune : int32
	//
	// ------------
	// Zero value
	// ------------
	//
	// Setiap tipe data memiliki default 0 value
	// Cth:
	// int = 0
	// string = ""
	// slice = nil
	//
	// Untuk set variable dengan 0 value, pakai `var`
	// saat deklarasi variabel

	var a int
	var b string

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)

	// String merupakan kumpulan dari uint8
	// String punya 2 struktur data dalam memory, pointer dan panjang kata di string
	// String = immutable
	// https://research.swtch.com/godata
	//

	// ------------
	// Inisialisasi
	// ------------
	// Untuk deklarasi dan inisialiasi, bisa menggunakan shortcut :=
	aa := 10
	bb := "hello" // 1st word points to a array of characters, 2nd word is 5 bytes
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// ------------
	// Type Conversion
	// ------------
	// Di golang, tidak ada type casting. Hanya ada type conversion
	// Kita harus mengalokasikan memori sesuai tipe data yang diinginkan

	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}
