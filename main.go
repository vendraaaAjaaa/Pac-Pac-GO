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
		level   int
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
	koin = 50
	punyaPeta = false
	puntaPedangLegenda = false
	simbol = '⚔'

	// Inisialisasi array
	inventory[0] = "Obor"
	inventory[1] = "Ramuan Kecil"
	inventory[2] = "Pelindung Dada"

	// Inisialisasi map
	hargaItem["Pedang"] = 100
	hargaItem["Perisai"] = 80
	hargaItem["Panah"] = 50
	hargaItem["Ramuan Besar"] = 150

	// Buat objek struct
	pemain := Karakter{nama: "", level: 1, exp: 0.0, senjata: "Tongkat"}

	var lokasiHutan = Lokasi{nama: "Hutan Angker", monster: "Goblin", harta: 50, level: 1}
	var lokasiGua = Lokasi{nama: "Gua Gelap", monster: "Troll", harta: 100, level: 2}
	var lokasiDanau = Lokasi{nama: "Danau Terkutuk", monster: "Naga", harta: 500, level: 3}

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
	fmt.Println("TIP: Kunjungi toko dulu untuk beli Pedang (100 koin)!")

	// =========== GAME LOOP ===========
	gameBerjalan := true
	ronde := 1

	for gameBerjalan {
		fmt.Printf("\n=== RONDE %d ===\n", ronde)
		fmt.Println("Pilih lokasi yang akan dijelajahi:")
		fmt.Println("1. Hutan Angker (Goblin) - Level 1")
		fmt.Println("2. Gua Gelap (Troll) - Level 2")
		fmt.Println("3. Danau Terkutuk (Naga) - Level 3")
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
			fmt.Println("4. Jual Ramuan Kecil (20 koin)")
			fmt.Println("5. Kembali")
			fmt.Print("Pilihan: ")

			var pilihanBeli int
			fmt.Scanln(&pilihanBeli)

			if pilihanBeli == 1 && koin >= 100 {
				koin -= 100
				pemain.senjata = "Pedang"
				fmt.Println("Anda membeli Pedang! (+20% keberanian)")
				keberanian += 20.0
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
				// Jual Ramuan Kecil
				for i := 0; i < len(inventory); i++ {
					if inventory[i] == "Ramuan Kecil" {
						inventory[i] = ""
						koin += 20
						fmt.Println("Ramuan Kecil terjual! +20 koin")
						break
					}
				}
			} else if pilihanBeli == 5 {
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
			pertempuranSelesai := false

			for !pertempuranSelesai {
				fmt.Printf("\nSeekor %s (Level %d) muncul!\n", lokasiSaatIni.monster, lokasiSaatIni.level)
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

					// LOGIKA PERTEMPURAN YANG DIPERBAIKI
					if pemain.senjata == "Pedang" && keberanian > 30 {
						fmt.Printf("%s dikalahkan! +%d koin\n",
							lokasiSaatIni.monster, lokasiSaatIni.harta)
						koin += lokasiSaatIni.harta
						pemain.exp += 25.0
						musuhDikalahkan = append(musuhDikalahkan, lokasiSaatIni.monster)
						keberanian += 10.0

						// Cek level up
						for pemain.exp >= 100 {
							pemain.level++
							pemain.exp -= 100
							nyawa += 20
							keberanian += 15.0
							fmt.Printf("LEVEL UP! Sekarang level %d (+20 HP, +15 keberanian)\n", pemain.level)
						}
						pertempuranSelesai = true
					} else if pemain.senjata == "Tongkat" && keberanian > 40 {
						fmt.Printf("%s dikalahkan dengan susah payah! +%d koin\n",
							lokasiSaatIni.monster, lokasiSaatIni.harta/2)
						koin += lokasiSaatIni.harta / 2
						pemain.exp += 15.0
						keberanian += 5.0
						pertempuranSelesai = true
					} else if pemain.senjata == "Pedang Legenda" {
						fmt.Printf("%s dikalahkan dengan mudah! +%d koin\n",
							lokasiSaatIni.monster, lokasiSaatIni.harta*2)
						koin += lokasiSaatIni.harta * 2
						pemain.exp += 40.0
						musuhDikalahkan = append(musuhDikalahkan, lokasiSaatIni.monster)
						pertempuranSelesai = true
					} else {
						fmt.Printf("%s terlalu kuat! Anda kalah -15 HP\n",
							lokasiSaatIni.monster)
						nyawa -= 15
						keberanian -= 5.0
						pertempuranSelesai = true
					}

				} else if pilihanTempur == 2 {
					itemDigunakan := false

					for !itemDigunakan {
						fmt.Println("\nItem dalam inventory:")

						// PERULANGAN while style untuk inventory
						i := 0
						for i < len(inventory) {
							if inventory[i] != "" {
								fmt.Printf("%d. %s\n", i+1, inventory[i])
							}
							i++
						}

						fmt.Println("0. Kembali ke menu pertempuran")
						fmt.Print("Pilih item (0 untuk batal): ")

						var pilihItem int
						fmt.Scanln(&pilihItem)

						// Jika memilih 0, kembali ke menu pertempuran
						if pilihItem == 0 {
							fmt.Println("Batal menggunakan item.")
							itemDigunakan = true // Keluar dari loop item
							// TIDAK mengatur pertempuranSelesai = true, sehingga kembali ke menu pertempuran
						} else if pilihItem > 0 && pilihItem <= len(inventory) && inventory[pilihItem-1] != "" {
							item := inventory[pilihItem-1]
							fmt.Printf("Anda menggunakan %s\n", item)

							if item == "Ramuan Kecil" {
								nyawa += 20
								if nyawa > 100 {
									nyawa = 100
								}
								fmt.Println("+20 HP!")
								// Item healing tidak mengakhiri pertempuran
								inventory[pilihItem-1] = ""
								itemDigunakan = true
								// Pertempuran belum selesai, bisa memilih aksi lagi
								fmt.Println("Pertempuran berlanjut...")

							} else if item == "Obor" {
								fmt.Println("Obor menerangi jalan, monster kabur!")
								koin += 30
								keberanian += 5.0
								inventory[pilihItem-1] = ""
								itemDigunakan = true
								pertempuranSelesai = true // Monster kabur, pertempuran selesai

							} else if item == "Perisai" || item == "Pelindung Dada" {
								fmt.Println("Anda bertahan dengan perisai! Monster mundur.")
								koin += 25
								inventory[pilihItem-1] = ""
								itemDigunakan = true
								pertempuranSelesai = true // Monster mundur, pertempuran selesai

							} else {
								fmt.Println("Item tidak dikenal!")
								// Tetap dalam loop item
							}
						} else {
							fmt.Println("Item tidak valid!")
							// Tetap dalam loop item
						}
					}

				} else if pilihanTempur == 3 {
					fmt.Println("Anda kabur dengan selamat...")
					keberanian -= 10.0
					if keberanian < 0 {
						keberanian = 0
					}
					pertempuranSelesai = true

				} else {
					fmt.Println("Pilihan tidak valid! Silakan pilih 1-3")
					// Loop akan berlanjut, meminta input lagi
				}

				// Cek jika pemain mati selama pertempuran
				if nyawa <= 0 {
					fmt.Println("\n💀 NYAWA ANDA HABIS SELAMA PERTEMPURAN!")
					pertempuranSelesai = true
					gameBerjalan = false
				}
			}

			// Cek apakah menemukan peta (setelah pertempuran selesai)
			if pilihanMenu == 1 && ronde%3 == 0 && nyawa > 0 {
				fmt.Println("Anda menemukan peta harta karun di rerumputan!")
				punyaPeta = true
			}

			// Cek apakah menemukan pedang legenda (setelah pertempuran selesai)
			if pilihanMenu == 2 && koin >= 200 && !puntaPedangLegenda && pemain.level >= 3 && nyawa > 0 {
				fmt.Println("Anda menemukan Pedang Legenda tersembunyi di dalam gua!")
				pemain.senjata = "Pedang Legenda"
				puntaPedangLegenda = true
				keberanian += 30.0
			}

			// =========== KONDISI KEMENANGAN DIPERBAIKI ===========
			// CARA 1: Menjadi legendaris
			if punyaPeta && puntaPedangLegenda && koin >= 1000 && nyawa > 0 {
				fmt.Println("\n🎉 SELAMAT! ANDA MENJADI PETUALANG LEGENDARIS! 🎉")
				fmt.Println("Anda memiliki peta, pedang legenda, dan kaya raya!")
				gameBerjalan = false
			}
			// CARA 2: Mengalahkan Naga (alternatif)
			if lokasiSaatIni.monster == "Naga" && (pemain.senjata == "Pedang" || pemain.senjata == "Pedang Legenda") && keberanian > 40 && nyawa > 0 {
				fmt.Println("\n🐉 EPIC VICTORY! ANDA MENGALAHKAN NAGA! 🐉")
				fmt.Println("Anda menyelamatkan kerajaan dan menjadi pahlawan!")
				koin += 1000
				gameBerjalan = false
			}
			// CARA 3: Level maksimal
			if pemain.level >= 10 && nyawa > 0 {
				fmt.Println("\n👑 LEGENDARY STATUS! Level maksimal tercapai!")
				fmt.Println("Anda diakui sebagai petualang terhebat sepanjang masa!")
				gameBerjalan = false
			}

			// Cek game over
			if nyawa <= 0 {
				fmt.Println("\n💀 NYAWA ANDA HABIS! GAME OVER 💀")
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
		fmt.Printf("Total: %d monster dikalahkan!\n", len(musuhDikalahkan))
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
