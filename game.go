package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Inisialisasi semua tipe data dasar
	var namaPetualang string      // string
	var nyawa int = 100           // int
	var keberanian float64 = 50.0 // float64
	var punyaPeta bool = false    // bool
	var simbolPedang rune = '⚔'   // rune
	var lokasi complex64 = 0 + 0i // complex64 (koordinat hutan)

	// Input nama pemain
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama petualang: ")
	namaPetualang, _ = reader.ReadString('\n')
	namaPetualang = strings.TrimSpace(namaPetualang)

	// Cerita pembuka
	fmt.Printf("\n%s, kau berdiri di tepi Hutan Angker yang legendaris.\n", namaPetualang)
	fmt.Printf("Koordinat masuk: (%.0f, %.0f)\n", real(lokasi), imag(lokasi))
	fmt.Println("Dengar bisikan angin: \"Hanya yang berhati murni yang menemukan harta karun Naga Tua\"")

	// Game loop - 3 area berbeda
	for area := 1; area <= 3 && nyawa > 0; area++ {
		fmt.Printf("\n===== AREA %d =====\n", area)
		fmt.Printf("Status: Nyawa [%d] | Keberanian [%.0f%%] | Pedang: %c\n",
			nyawa, keberanian, simbolPedang)

		// Update lokasi kompleks (tanpa perhitungan matematika rumit)
		switch area {
		case 1:
			lokasi = 3 + 4i
		case 2:
			lokasi = -2 + 7i
		case 3:
			lokasi = 5 - 3i
		}
		fmt.Printf("Koordinat saat ini: (%.0f, %.0f)\n", real(lokasi), imag(lokasi))

		// Konten berbeda per area (switch-case)
		switch area {
		case 1:
			fmt.Println("\n.Seekor rubah terluka menghalangi jalan.")
			fmt.Println("1. Bantu rubah dengan ramuan")
			fmt.Println("2. Lewati begitu saja")
			pilihan := getInput()

			if pilihan == "1" {
				fmt.Println("Rubah itu memberimu peta rahasia! (+20 keberanian)")
				punyaPeta = true
				keberanian += 20.0
				if keberanian > 100.0 {
					keberanian = 100.0
				}
			} else {
				fmt.Println("Rubah itu menggigit kakimu! (-15 nyawa)")
				nyawa -= 15
			}

		case 2:
			fmt.Println("\n.Goa gelap menganga di depanmu. Suara gemuruh terdengar.")
			fmt.Println("1. Nyalakan obor dan masuk")
			fmt.Println("2. Pasang perangkap di mulut gua")
			pilihan := getInput()

			if pilihan == "1" {
				if keberanian > 60.0 {
					fmt.Println("Kau menemukan pedang legendaris di dalam gua! (+30 nyawa)")
					nyawa += 30
					if nyawa > 100 {
						nyawa = 100
					}
				} else {
					fmt.Println("Keberanianmu tidak cukup! Terjebak dalam kegelapan (-25 nyawa)")
					nyawa -= 25
				}
			} else if pilihan == "2" {
				fmt.Println("Seekor beruang terperangkap! Kau lolos dengan selamat.")
				keberanian += 10.0
			}

		case 3:
			fmt.Println("\n.Tanah bergetar! Naga penjaga harta muncul!")
			fmt.Println("1. Tantang naga dengan pedang")
			fmt.Println("2. Perlihatkan peta pemberian rubah")
			pilihan := getInput()

			if pilihan == "1" {
				if nyawa > 70 && keberanian > 80.0 {
					fmt.Println("Kau mengalahkan naga! Harta karun jadi milikmu!")
				} else {
					fmt.Println("Naga membakarmu! Nyawa -40")
					nyawa -= 40
				}
			} else if pilihan == "2" {
				if punyaPeta {
					fmt.Println("Naga mengenali peta rubah keramat. Ia memberimu harta karun!")
				} else {
					fmt.Println("Kau tak punya peta! Naga marah dan menyerang (-35 nyawa)")
					nyawa -= 35
				}
			}
		}
	}

	// Akhir game
	fmt.Println("\n===== PETUALANGAN BERAKHIR =====")
	if nyawa <= 0 {
		fmt.Printf("%s gugur di Hutan Angker...\n", namaPetualang)
	} else if punyaPeta {
		fmt.Printf("SELAMAT %s!\n", namaPetualang)
		fmt.Println("Kau membawa pulang harta karun dengan selamat!")
		fmt.Printf("Skor akhir: Nyawa [%d] | Keberanian [%.0f%%]\n", nyawa, keberanian)
	} else {
		fmt.Printf("%s selamat tapi pulang dengan tangan hampa.\n", namaPetualang)
		fmt.Println("Lain kali jangan abaikan makhluk hutan...")
	}
}

// Fungsi input aman
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pilihan (1/2): ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
