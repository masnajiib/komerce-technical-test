package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 1. Input Jumlah Keluarga (n)
	fmt.Print("Masukkan jumlah keluarga: ")
	if !scanner.Scan() {
		return
	}
	nStr := scanner.Text()
	n, err := strconv.Atoi(strings.TrimSpace(nStr))
	if err != nil {
		fmt.Println("Error: Input harus berupa angka.")
		return
	}

	// 2. Input Anggota Keluarga
	fmt.Print("Masukkan jumlah anggota setiap keluarga (pisahkan dengan spasi): ")
	if !scanner.Scan() {
		return
	}
	membersStr := scanner.Text()
	
	// Memecah input string menjadi slice (array) berdasarkan spasi
	memberFields := strings.Fields(membersStr)

	// 3. Validasi: Apakah jumlah input sesuai dengan n?
	if len(memberFields) != n {
		fmt.Println("Input harus sama dengan jumlah keluarga")
		return
	}

	// Konversi string ke integer untuk diolah
	var families []int
	for _, f := range memberFields {
		val, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("Error: Data anggota keluarga harus berupa angka.")
			return
		}
		families = append(families, val)
	}

	// 4. Algoritma Utama (Greedy / Two Pointers)
	// Urutkan dari kecil ke besar agar mudah dipasangkan
	sort.Ints(families)

	buses := 0
	left := 0              // Pointer ke keluarga paling kecil
	right := len(families) - 1 // Pointer ke keluarga paling besar

	for left <= right {
		// Jika pointer left dan right bertemu di satu keluarga yang sama,
		// berarti tinggal 1 keluarga tersisa. Dia butuh 1 bus sendirian.
		if left == right {
			buses++
			break
		}

		// Menggabungkan keluarga terbesar (right) dengan keluarga terkecil (left)
		// Syarat: Total anggota <= 4 (Kapasitas Bus)
		if families[right] + families[left] <= 4 {
			// Jika muat, mereka satu bus (maksimal 2 keluarga per bus terpenuhi)
			left++  // Keluarga kecil sudah terangkut
			right-- // Keluarga besar sudah terangkut
			buses++ // Tambah 1 bus
		} else {
			// Jika tidak muat digabung, keluarga terbesar (right) harus naik bus sendiri
			// Keluarga terkecil (left) menunggu giliran berikutnya untuk dipasangkan
			right-- 
			buses++ // Tambah 1 bus
		}
	}

	// 5. Output Hasil
	fmt.Printf("Minimum bus yang diperlukan adalah: %d\n", buses)
}