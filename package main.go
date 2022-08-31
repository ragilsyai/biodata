package main

import (
	"fmt"
	"os"
	"strconv"
)

type List struct {
	absen  int
	nama   string
	age    string
	alamat string
	alasan string
}

func main() {
	Program := os.Args
	nomer, _ := strconv.Atoi(Program[1])
	getFindData(nomer)
}

var Data []List

func getFindData(nomer int) {

	var Data = []List{
		{absen: 1, nama: "ahamd", age: "24", alamat: "Indonesia", alasan: "Ingin"},
		{absen: 2, nama: "Koi", age: "31", alamat: "mars", alasan: "disuruh"},
	}
	for _, siswa := range Data {

		if siswa.absen == nomer {
			fmt.Println("Nama 			   : ", siswa.nama)
			fmt.Println("Alamat 			  : ", siswa.alamat)
			fmt.Println("Pekerjaan  	 	  : ", siswa.alamat)
			fmt.Println("Alasan memilih kelas Golang : ", siswa.alasan)
		}
	}
}
