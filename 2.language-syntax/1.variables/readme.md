# Language Syntax : Variables

**Definisi**
- tempat untuk menyimpan suatu value
- menyediakan kemampuan untuk read & write value ke memory

**Deklarasi Variabel**
- Di Golang, deklarasi variabel menggunakan _var_ beserta tipe datanya

    Cth:

    ```go
    // deklarasi variabel
    var firstName string

    // set firstName value
    firstName = "bob"
    fmt.Printf("var firstName string \t %T [%v]\n",firstName, firstName) // var firstName string 	 string [bob]

    // deklarasi variabel beserta value
    var lastName string = "johnson"
    fmt.Printf("var lastName string \t %T [%v]\n",lastName, lastName) // var lastName string 	 string [johnson]
    ```

- Untuk deklarasi tanpa menggunakan tipe data, menggunakan short variable declaration ```go :=```

    Cth:

    ```go
    // deklarasi variabel tanpa menggunakan tipe data
    address := "jakarta selatan"
    fmt.Printf("address := \"jakarta selatan\" \t %T [%v]\n", address, address) // firstName := "Bob" 	 string [bob]
    ```

- Secara _**default**_ ketika menggunakan var, nilainya adalah zero value (sesuai tipe data)

    Cth:

    ```go
    var job string
    var telephone int

    fmt.Printf("var job string \t %T [%v]\n",job, job) // var job string 	 string []
	fmt.Printf("var telephone string \t %T [%v]\n",telephone, telephone) // var telephone string 	 int [0]

    // Boolean false
	// Integer 0
	// Floating Point 0
	// Complex 0i
	// String "" (empty string)
	// Pointer nil
    ```

**NOTE**

- Tujuan semua program adalah untuk transformasi data dari satu bentuk ke bentuk lain
- Jika tidak mengerti data, maka tidak akan mengerti masalah yang dihadapi
- Kita mengerti masalah lebih baik dengan memahami data

[Sample Code](https://github.com/milhamhidayat/ultimate-go-programming-code/blob/master/2.language-syntax/1.variables/example1.go)