package main

import (
	"fmt"
	"os"
	"strconv"
	"submission1/helper"
)

var TemanList = map[int]helper.Teman{
	1: {"Tamir", "Jalan Teguh Permai 1 no 6 A", "Data Engineering", "Tertarik belajar Data Engineering"},
	2: {"Gading", "Jalan Kertapati", "Data Scientist", "Ingin memanfaatkan data untuk mengambil keputusan dalam bisnis"},
	3: {"Retno", "Jalan Sokarno", "Full Stack Web Development", "Ingin memahami end-to-end web development"},
}

func GetDataByAbsen(absen int) (helper.Teman, error) {
	teman, found := TemanList[absen]
	if !found {
		return helper.Teman{}, fmt.Errorf("Data teman dengan absen %d tidak ditemukan", absen)
	}
	return teman, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <nomor_absen>")
		return
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat.")
		return
	}

	teman, err := GetDataByAbsen(absen)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Data Teman dengan Absen %d:\n", absen)
	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", teman.Nama, teman.Alamat, teman.Pekerjaan, teman.Alasan)
}
