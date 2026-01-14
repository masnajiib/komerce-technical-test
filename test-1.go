package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// 1. Membaca Input
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Masukkan satu baris kata (S): ")
	
	if scanner.Scan() {
		input := scanner.Text()

		var vowels string
		var consonants string

		// 2. Proses Logika
		for _, char := range strings.ToLower(input) {
			
			// Abaikan spasi
			if char == ' ' {
				continue
			}

			// Validasi: Pastikan karakter adalah huruf
			// (Opsional: hapus blok if ini jika simbol ingin dianggap konsonan)
			if !unicode.IsLetter(char) {
				continue
			}

			// Cek apakah karakter adalah vokal
			switch char {
			case 'a', 'e', 'i', 'o', 'u':
				vowels += string(char)
			default:
				consonants += string(char)
			}
		}

		fmt.Println("Karakter Vokal :")
		fmt.Println(vowels)
		fmt.Println("Karakter Konsonan :")
		fmt.Println(consonants)
	}
}