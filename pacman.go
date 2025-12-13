package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maze = [][]string{
	{"#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#"},
	{"#", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", "#"},
	{"#", " ", ".", " ", "#", " ", "#", " ", ".", " ", "#", " ", "#", " ", ".", " ", "#", " ", "#", " ", ".", " ", "#"},
	{"#", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", "#"},
	{"#", " ", ".", " ", "#", " ", "#", " ", ".", " ", "#", " ", "#", " ", ".", " ", "#", " ", "#", " ", ".", " ", "#"},
	{"#", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", ".", " ", "#"},
	{"#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#", " ", "#"},
}

var (
	pacmanX = 1
	pacmanY = 1
	score   = 0
	lives   = 3
)

func main() {

	maze[pacmanY][pacmanX] = " "

	reader := bufio.NewReader(os.Stdin)

	for {

		clearScreen()
		fmt.Println("\n" + strings.Repeat("=", 40))
		for y := 0; y < len(maze); y++ {
			for x := 0; x < len(maze[y]); x++ {
				if y == pacmanY && x == pacmanX {
					fmt.Print("P")
				} else {
					fmt.Print(maze[y][x])
				}
			}
			fmt.Println()
		}
		fmt.Printf("\nScore: %d | Lives: %d\n", score, lives)
		fmt.Println("Gerak: w (atas), a (kiri), s (bawah), d (kanan)")
		fmt.Println("Ketik 'q' untuk keluar. Tekan Enter setelah input.")

		if !dotsLeft() {
			fmt.Println("\n YOU WINN!!")
			return
		}

		if lives <= 0 {
			fmt.Println("\n💀 Game Over!")
			return
		}

		fmt.Print("→ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "q" {
			fmt.Println("Keluar dari game.")
			return
		}

		newX, newY := pacmanX, pacmanY
		switch input {
		case "w":
			newY--
		case "s":
			newY++
		case "a":
			newX--
		case "d":
			newX++
		default:
			fmt.Println("Input tidak valid!")
			continue
		}

		if newY >= 0 && newY < len(maze) && newX >= 0 && newX < len(maze[0]) {
			cell := maze[newY][newX]
			if cell != "#" { // Bukan dinding
				pacmanX, pacmanY = newX, newY

				if cell == "." {
					score += 10
					maze[newY][newX] = " "
				}
			}
		}
	}
}

func dotsLeft() bool {
	for _, row := range maze {
		for _, cell := range row {
			if cell == "." {
				return true
			}
		}
	}
	return false
}

func clearScreen() {
	for i := 0; i < 5; i++ {
		fmt.Println()
	}
}
