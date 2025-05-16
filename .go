package main

import (
	"fmt"
	"os"
)

const filename = "data.txt"

func main() {
	// data lama
	fmt.Println("Data yang sudah tersimpan:")
	readDataFromFile()

	// contoh inputannya
	var nama, umur, alamat string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan umur: ")
	fmt.Scanln(&umur)

	fmt.Print("Masukkan alamat: ")
	fmt.Scanln(&alamat)

	// format penyimpanan datanya
	data := fmt.Sprintf("Nama: %s, Umur: %s, Alamat: %s\n", nama, umur, alamat)

	err := appendDataToFile(data)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
		return
	}

	fmt.Println("Data berhasil disimpan.")
}

func readDataFromFile() {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Belum ada data.")
			return
		}
		fmt.Println("Gagal membaca file:", err)
		return
	}

	fmt.Println(string(data))
}

func appendDataToFile(data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}
