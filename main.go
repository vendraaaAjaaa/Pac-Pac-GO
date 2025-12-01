package main

import "fmt"

func main() {
	var a, b int
	var op string
	var jawab string = "y"

	for jawab == "y" {
		fmt.Print("Masukkan Operasi (+, -, *, /): ")
		fmt.Scanln(&op)
		if op == "+" {
			fmt.Print("Masukkan angka pertama: ")
			fmt.Scanln(&a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Scanln(&b)
			fmt.Printf("Hasil: %d\n", a+b)
		}
		if op == "-" {
			fmt.Print("Masukkan angka pertama: ")
			fmt.Scanln(&a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Scanln(&b)
			fmt.Printf("Hasil: %d\n", a-b)
		}
		if op == "*" {
			fmt.Print("Masukkan angka pertama: ")
			fmt.Scanln(&a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Scanln(&b)
			fmt.Printf("Hasil: %d\n", a*b)
		}
		if op == "/" {
			fmt.Print("Masukkan angka pertama: ")
			fmt.Scanln(&a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Scanln(&b)
			if b != 0 {
				fmt.Printf("Hasil: %d\n", a/b)
			} else {
				fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
			}
		} else if op != "+" && op != "-" && op != "*" && op != "/" {
			fmt.Println("Operasi tidak valid.")
		}

		fmt.Print("Apakah Anda ingin melanjutkan penggunaan kalkulator ini? (y/n): ")
		fmt.Scanln(&jawab)

		if jawab != "y" {
			fmt.Println("Terima kasih telah menggunakan kalkulator sederhana ini.")
		}
	}
}
