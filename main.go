package main

import "fmt"

func main() {
	// =========== DEKLARASI SEMUA TIPE DATA ===========

	// Tipe data PRIMITIVE
	var namaPetualang string
	var nyawa, koin int
	var keberanian float64
	var punyaPeta, puntaPedangLegenda bool
	var simbol rune

	// Tipe data ARRAY (fixed size)
	var inventory [5]string

	// Tipe data SLICE (dynamic array)
	var musuhDikalahkan []string

	// Tipe data MAP
	var hargaItem map[string]int
	hargaItem = make(map[string]int)

	// Tipe data STRUCT
	type Lokasi struct {
		nama    string
		monster string
		harta   int
	}

	type Karakter struct {
		nama    string
		level   int
		exp     float64
		senjata string
	}

	// =========== INISIALISASI ===========

	// Inisialisasi variabel
	namaPetualang = ""
	nyawa = 100
	keberanian = 50.0
	koin = 10
	punyaPeta = false
	puntaPedangLegenda = false
	simbol = '⚔'

	// Inisialisasi array
	inventory[0] = "Obor"
	inventory[1] = "Ramuan Kecil"

	// Inisialisasi map
	hargaItem["Pedang"] = 100
	hargaItem["Perisai"] = 80
	hargaItem["Panah"] = 50
	hargaItem["Ramuan Besar"] = 150

	// Buat objek struct
	pemain := Karakter{nama: "", level: 1, exp: 0.0, senjata: "Tongkat"}

	var lokasiHutan = Lokasi{nama: "Hutan Angker", monster: "Goblin", harta: 50}
	var lokasiGua = Lokasi{nama: "Gua Gelap", monster: "Troll", harta: 100}
	var lokasiDanau = Lokasi{nama: "Danau Terkutuk", monster: "Naga", harta: 500}

	var lokasiSaatIni Lokasi
	var pilihanMenu int

	// =========== AWAL GAME ===========

	fmt.Println("===============================")
	fmt.Println("    PETUALANGAN DI HUTAN ANGKER")
	fmt.Println("===============================")

	fmt.Print("Masukkan nama petualang: ")
	fmt.Scanln(&namaPetualang)
	pemain.nama = namaPetualang

	fmt.Printf("\nSelamat datang, %s!\n", pemain.nama)
	fmt.Println("Kesehatan:", nyawa, "| Keberanian:", keberanian, "| Koin:", koin)
	fmt.Println("Senjata:", pemain.senjata)

	// =========== GAME LOOP ===========
	gameBerjalan := true
	ronde := 1

	for gameBerjalan {
		fmt.Printf("\n=== RONDE %d ===\n", ronde)
		fmt.Println("Pilih lokasi yang akan dijelajahi:")
		fmt.Println("1. Hutan Angker (Goblin)")
		fmt.Println("2. Gua Gelap (Troll)")
		fmt.Println("3. Danau Terkutuk (Naga)")
		fmt.Println("4. Kunjungi toko")
		fmt.Println("5. Cek status")
		fmt.Println("6. Lihat inventory")
		fmt.Println("7. Keluar dari game")
		fmt.Print("Pilihan (1-7): ")

		fmt.Scanln(&pilihanMenu)

		// PERCABANGAN menggunakan switch
		switch pilihanMenu {
		case 1:
			lokasiSaatIni = lokasiHutan
			fmt.Printf("\nAnda memasuki %s\n", lokasiSaatIni.nama)

		case 2:
			lokasiSaatIni = lokasiGua
			fmt.Printf("\nAnda memasuki %s\n", lokasiSaatIni.nama)

		case 3:
			lokasiSaatIni = lokasiDanau
			fmt.Printf("\nAnda memasuki %s\n", lokasiSaatIni.nama)

		case 4:
			// Toko
			fmt.Println("\n=== TOKO PETUALANG ===")
			fmt.Println("Koin Anda:", koin)
			fmt.Println("Barang yang tersedia:")

			// PERULANGAN for-range pada map
			for item, harga := range hargaItem {
				fmt.Printf("- %s: %d koin\n", item, harga)
			}

			fmt.Println("\n1. Beli Pedang (100 koin)")
			fmt.Println("2. Beli Perisai (80 koin)")
			fmt.Println("3. Beli Ramuan Besar (150 koin)")
			fmt.Println("4. Kembali")
			fmt.Print("Pilihan: ")

			var pilihanBeli int
			fmt.Scanln(&pilihanBeli)

			if pilihanBeli == 1 && koin >= 100 {
				koin -= 100
				pemain.senjata = "Pedang"
				fmt.Println("Anda membeli Pedang!")
			} else if pilihanBeli == 2 && koin >= 80 {
				koin -= 80
				// Cari slot kosong di inventory
				for i := 0; i < len(inventory); i++ {
					if inventory[i] == "" {
						inventory[i] = "Perisai"
						break
					}
				}
				fmt.Println("Anda membeli Perisai!")
			} else if pilihanBeli == 3 && koin >= 150 {
				koin -= 150
				nyawa += 50
				if nyawa > 100 {
					nyawa = 100
				}
				fmt.Println("Anda menggunakan Ramuan Besar! +50 HP")
			} else if pilihanBeli == 4 {
				// Kembali ke menu
			} else {
				fmt.Println("Koin tidak cukup atau pilihan tidak valid!")
			}
			continue

		case 5:
			fmt.Println("\n=== STATUS PEMAIN ===")
			fmt.Printf("Nama: %s\n", pemain.nama)
			fmt.Printf("Level: %d\n", pemain.level)
			fmt.Printf("EXP: %.1f/100\n", pemain.exp)
			fmt.Printf("Senjata: %s\n", pemain.senjata)
			fmt.Printf("Kesehatan: %d\n", nyawa)
			fmt.Printf("Keberanian: %.1f\n", keberanian)
			fmt.Printf("Koin: %d\n", koin)
			continue

		case 6:
			fmt.Println("\n=== INVENTORY ===")
			// PERULANGAN for dengan counter
			itemCount := 0
			for i := 0; i < len(inventory); i++ {
				if inventory[i] != "" {
					fmt.Printf("%d. %s\n", i+1, inventory[i])
					itemCount++
				}
			}
			if itemCount == 0 {
				fmt.Println("Inventory kosong!")
			}
			continue

		case 7:
			fmt.Println("Terima kasih telah bermain!")
			gameBerjalan = false
			continue

		default:
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		// =========== PERTEMPURAN ===========
		if pilihanMenu >= 1 && pilihanMenu <= 3 {
			fmt.Printf("Seekor %s muncul!\n", lokasiSaatIni.monster)
			fmt.Println("1. Serang dengan senjata")
			fmt.Println("2. Gunakan item dari inventory")
			fmt.Println("3. Kabur")
			fmt.Print("Pilihan (1-3): ")

			var pilihanTempur int
			fmt.Scanln(&pilihanTempur)

			// PERCABANGAN if-else
			if pilihanTempur == 1 {
				fmt.Printf("Anda menyerang %s dengan %s %c\n",
					lokasiSaatIni.monster, pemain.senjata, simbol)

				// Logika pertempuran sederhana
				if pemain.senjata == "Pedang" && keberanian > 40 {
					fmt.Printf("%s dikalahkan! +%d koin\n",
						lokasiSaatIni.monster, lokasiSaatIni.harta)
					koin += lokasiSaatIni.harta
					pemain.exp += 25.0
					musuhDikalahkan = append(musuhDikalahkan, lokasiSaatIni.monster)

					// Cek level up
					for pemain.exp >= 100 {
						pemain.level++
						pemain.exp -= 100
						nyawa += 20
						fmt.Printf("LEVEL UP! Sekarang level %d\n", pemain.level)
					}
				} else if pemain.senjata == "Tongkat" && keberanian > 60 {
					fmt.Printf("%s dikalahkan! +%d koin\n",
						lokasiSaatIni.monster, lokasiSaatIni.harta/2)
					koin += lokasiSaatIni.harta / 2
					pemain.exp += 15.0
				} else {
					fmt.Printf("%s terlalu kuat! Anda kalah -20 HP\n",
						lokasiSaatIni.monster)
					nyawa -= 20
				}

			} else if pilihanTempur == 2 {
				fmt.Println("Item dalam inventory:")
				// PERULANGAN while style
				i := 0
				for i < len(inventory) {
					if inventory[i] != "" {
						fmt.Printf("%d. %s\n", i+1, inventory[i])
					}
					i++
				}

				fmt.Print("Pilih item (0 untuk batal): ")
				var pilihItem int
				fmt.Scanln(&pilihItem)

				if pilihItem > 0 && pilihItem <= len(inventory) && inventory[pilihItem-1] != "" {
					item := inventory[pilihItem-1]
					fmt.Printf("Anda menggunakan %s\n", item)

					if item == "Ramuan Kecil" {
						nyawa += 20
						if nyawa > 100 {
							nyawa = 100
						}
						fmt.Println("+20 HP!")
					} else if item == "Obor" {
						fmt.Println("Obor menerangi jalan, monster kabur!")
						koin += 20
					}

					// Hapus item dari inventory
					inventory[pilihItem-1] = ""
				} else {
					fmt.Println("Item tidak valid!")
				}

			} else if pilihanTempur == 3 {
				fmt.Println("Anda kabur dengan selamat...")
				keberanian -= 10.0
				if keberanian < 0 {
					keberanian = 0
				}
			} else {
				fmt.Println("Pilihan tidak valid!")
			}

			// Cek apakah menemukan peta
			if pilihanMenu == 1 && ronde%2 == 0 {
				fmt.Println("Anda menemukan peta harta karun!")
				punyaPeta = true
			}

			// Cek apakah menemukan pedang legenda
			if pilihanMenu == 2 && koin >= 200 && !puntaPedangLegenda {
				fmt.Println("Anda menemukan Pedang Legenda di gua!")
				pemain.senjata = "Pedang Legenda"
				puntaPedangLegenda = true
			}

			// Cek game over
			if nyawa <= 0 {
				fmt.Println("\n💀 NYAWA ANDA HABIS! GAME OVER 💀")
				gameBerjalan = false
			}

			// Cek kemenangan
			if punyaPeta && puntaPedangLegenda && koin >= 1000 {
				fmt.Println("\n🎉 SELAMAT! ANDA MENJADI PETUALANG LEGENDARIS! 🎉")
				fmt.Println("Anda memiliki peta, pedang legenda, dan kaya raya!")
				gameBerjalan = false
			}

			ronde++
		}
	}

	// =========== AKHIR GAME ===========
	fmt.Println("\n===============================")
	fmt.Println("        HASIL AKHIR GAME")
	fmt.Println("===============================")

	fmt.Printf("Nama Petualang: %s\n", pemain.nama)
	fmt.Printf("Level Akhir: %d\n", pemain.level)
	fmt.Printf("Koin Akhir: %d\n", koin)
	fmt.Printf("Nyawa Akhir: %d\n", nyawa)
	fmt.Printf("Keberanian Akhir: %.1f\n", keberanian)

	if punyaPeta {
		fmt.Println("Status: Memiliki Peta Harta Karun ✓")
	}
	if puntaPedangLegenda {
		fmt.Println("Status: Memiliki Pedang Legenda ✓")
	}

	// Tampilkan musuh yang dikalahkan
	if len(musuhDikalahkan) > 0 {
		fmt.Println("\nMusuh yang dikalahkan:")
		// PERULANGAN for-range pada slice
		for index, musuh := range musuhDikalahkan {
			fmt.Printf("%d. %s\n", index+1, musuh)
		}
	} else {
		fmt.Println("\nTidak ada musuh yang dikalahkan.")
	}

	// Tampilkan inventory akhir
	fmt.Println("\nInventory akhir:")
	kosong := true
	for i := 0; i < len(inventory); i++ {
		if inventory[i] != "" {
			fmt.Printf("- %s\n", inventory[i])
			kosong = false
		}
	}
	if kosong {
		fmt.Println("Kosong")
	}

	fmt.Println("\nTerima kasih telah bermain!")
}
