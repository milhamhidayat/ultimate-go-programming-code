package main

import "fmt"

func main() {
	// deklarasi variabel
	var firstName string

	// set firstName value
	firstName = "bob"
	fmt.Printf("var firstName string \t %T [%v]\n", firstName, firstName) // var firstName string 	 string [bob]

	// deklarasi variabel beserta value
	var lastName string = "johnson"
	fmt.Printf("var lastName string \t %T [%v]\n", lastName, lastName) // var lastName string 	 string [johnson]

	// deklarasi variabel tanpa menggunakan tipe data
	address := "jakarta selatan"
	fmt.Printf("address := \"jakarta selatan\" \t %T [%v]\n", address, address) // firstName := "Bob" 	 string [bob]

	var job string
	var telephone int

	fmt.Printf("var job string \t %T [%v]\n", job, job)                   // var job string 	 string []
	fmt.Printf("var telephone string \t %T [%v]\n", telephone, telephone) // var telephone string 	 int [0]

	// Boolean false
	// Integer 0
	// Floating Point 0
	// Complex 0i
	// String "" (empty string)
	// Pointer nil
}
