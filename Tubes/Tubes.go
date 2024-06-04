package main

import (
	"fmt"
)

const NMAX = 100

type Partai struct {
	Name       string
	DataCalon  [NMAX]Calon
	CountCalon int
}

type Calon struct {
	Name        string
	Umur        int
	NVoter      int
	PartaiName  string
	DataPemilih [NMAX]Voter
}

type Voter struct {
	Name, Password, Username string
}

type TabPartai [NMAX]Partai
type TabCalon [NMAX]Calon
type TabVoter [NMAX]Voter

var threshold int
var nPartai, nCalon, nVoter int
var dataPartai TabPartai
var dataCalon TabCalon
var dataVoter TabVoter

func main() {
	mainPage()
}

func mainPage() {
	var choice int
	for {
		fmt.Println("================== PEMILU ===================")
		fmt.Println("1. Pegawai")
		fmt.Println("2. Umum")
		fmt.Println("3. Exit")
		fmt.Println("=============================================")
		fmt.Print("Saya masuk sebagai (1/2/3): ")
		fmt.Scan(&choice)

		if choice == 1 {
			if loginAdmin() {
				fmt.Println("Saya masuk sebagai ADMIN")
				adminPage()
			} else {
				fmt.Println("Password atau username salah")
			}
		} else if choice == 2 {
			voterPage()
		} else if choice == 3 {
			fmt.Println("Exit Program!")
			return
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func loginAdmin() bool {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	return username == "admin" && password == "admin123"
}

func adminPage() {
	var choice int

	for {
		fmt.Println("================== ADMIN PAGE ===================")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Tampilkan Data Terurut")
		fmt.Println("5. Set Threshold")
		fmt.Println("6. Pencarian Data")
		fmt.Println("7. Kembali")
		fmt.Println("=============================================")
		fmt.Print("Masukan pilihan (1/2/3/4/5/6/7): ")
		fmt.Scan(&choice)

		if choice == 1 {
			addData()
		} else if choice == 2 {
			editData()
		} else if choice == 3 {
			deleteData()
		} else if choice == 4 {
			displaySortedData()
		} else if choice == 5 {
			setThreshold()
		} else if choice == 6 {
			searchData()
		} else if choice == 7 {
			return // keluar dari adminPage
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func voterPage() {
	var choice int
	fmt.Println("================== VOTER PAGE ===================")
	fmt.Println("1. Lihat daftar calon")
	fmt.Println("2. Voting")
	fmt.Println("3. Kembali")
	fmt.Println("=============================================")
	fmt.Print("Masukan pilihan (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		displayCalon()
	} else if choice == 2 {
		vote()
	} else if choice == 3 {
		return // keluar dari voterPage
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func loginVoter() bool {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < nVoter; i++ {
		if dataVoter[i].Username == username && dataVoter[i].Password == password {
			return true
		}
	}
	return false
}

func validateVotingTime() bool {
	// Simplified version without using time package
	// Assume the voting end time is 12:00
	var hour, minute int
	fmt.Print("Masukkan jam saat ini (24 jam format, HH MM): ")
	fmt.Scan(&hour, &minute)
	return hour < 12 || (hour == 12 && minute == 0)
}

func addData() {
	var choice int
	fmt.Println("1. Tambah Partai")
	fmt.Println("2. Tambah Calon")
	fmt.Println("3. Tambah Pemilih")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		if nPartai < NMAX {
			fmt.Print("Masukkan nama partai: ")
			fmt.Scan(&dataPartai[nPartai].Name)
			dataPartai[nPartai].CountCalon = 0
			nPartai++
		} else {
			fmt.Println("Kuota partai penuh")
		}
	} else if choice == 2 {
		if nCalon < NMAX {
			var partaiName string
			fmt.Print("Masukkan nama calon: ")
			fmt.Scan(&dataCalon[nCalon].Name)
			fmt.Print("Masukkan umur calon: ")
			fmt.Scan(&dataCalon[nCalon].Umur)
			fmt.Print("Masukkan nama partai calon: ")
			fmt.Scan(&partaiName)
			if AdadiPartai(partaiName) {
				dataCalon[nCalon].PartaiName = partaiName
				nCalon++
			} else {
				fmt.Println("Partai tidak ditemukan")
			}
		} else {
			fmt.Println("Kuota calon penuh")
		}
	} else if choice == 3 {
		if nVoter < NMAX {
			fmt.Print("Masukkan nama pemilih: ")
			fmt.Scan(&dataVoter[nVoter].Name)
			fmt.Print("Masukkan username pemilih: ")
			fmt.Scan(&dataVoter[nVoter].Username)
			fmt.Print("Masukkan password pemilih: ")
			fmt.Scan(&dataVoter[nVoter].Password)
			nVoter++
		} else {
			fmt.Println("Kuota pemilih penuh")
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func editData() {
	var choice, idx int
	fmt.Println("1. Edit Partai")
	fmt.Println("2. Edit Calon")
	fmt.Println("3. Edit Pemilih")
	fmt.Print("Masukan pilihan (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Masukkan indeks partai yang ingin diubah: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < nPartai {
			fmt.Print("Masukkan nama partai baru: ")
			fmt.Scan(&dataPartai[idx].Name)
		} else {
			fmt.Println("Indeks tidak valid")
		}
	} else if choice == 2 {
		fmt.Print("Masukkan indeks calon yang ingin diubah: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < nCalon {
			fmt.Print("Masukkan nama calon baru: ")
			fmt.Scan(&dataCalon[idx].Name)
			fmt.Print("Masukkan umur calon baru: ")
			fmt.Scan(&dataCalon[idx].Umur)
			fmt.Print("Masukkan nama partai calon baru: ")
			fmt.Scan(&dataCalon[idx].PartaiName)
		} else {
			fmt.Println("Indeks tidak valid")
		}
	} else if choice == 3 {
		fmt.Print("Masukkan indeks pemilih yang ingin diubah: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < nVoter {
			fmt.Print("Masukkan nama pemilih baru: ")
			fmt.Scan(&dataVoter[idx].Name)
			fmt.Print("Masukkan username pemilih baru: ")
			fmt.Scan(&dataVoter[idx].Username)
			fmt.Print("Masukkan password pemilih baru: ")
			fmt.Scan(&dataVoter[idx].Password)
		} else {
			fmt.Println("Indeks tidak valid")
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func deleteData() {
	var choice, idx int
	fmt.Println("1. Hapus Partai")
	fmt.Println("2. Hapus Calon")
	fmt.Println("3. Hapus Pemilih")
	fmt.Print("Masukan pilihan (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Masukkan indeks partai yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < nPartai {
			for i := idx; i < nPartai-1; i++ {
				dataPartai[i] = dataPartai[i+1]
			}
			nPartai--
		} else {
			fmt.Println("Indeks tidak valid")
		}
	} else if choice == 2 {
		fmt.Print("Masukkan indeks calon yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < nCalon {
			for i := idx; i < nCalon-1; i++ {
				dataCalon[i] = dataCalon[i+1]
			}
			nCalon--
		} else {
			fmt.Println("Indeks tidak valid")
		}
	} else if choice == 3 {
		fmt.Print("Masukkan indeks pemilih yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < nVoter {
			for i := idx; i < nVoter-1; i++ {
				dataVoter[i] = dataVoter[i+1]
			}
			nVoter--
		} else {
			fmt.Println("Indeks tidak valid")
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func displaySortedData() {
	var choice, sortChoice int
	fmt.Println("1. Tampilkan Partai")
	fmt.Println("2. Tampilkan Calon")
	fmt.Println("3. Tampilkan Pemilih")
	fmt.Print("Masukan pilihan (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Pilih metode sorting:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Masukan pilihan (1/2): ")
		fmt.Scan(&sortChoice)
		if sortChoice == 1 {
			selectionSortPartai()
		} else if sortChoice == 2 {
			insertionSortPartai()
		} else {
			fmt.Println("Pilihan tidak valid")
			return
		}
		for i := 0; i < nPartai; i++ {
			fmt.Printf("Partai %d: %s\n", i, dataPartai[i].Name)
		}
	} else if choice == 2 {
		fmt.Println("Pilih metode sorting:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Masukan pilihan (1/2): ")
		fmt.Scan(&sortChoice)
		if sortChoice == 1 {
			selectionSortCalon()
		} else if sortChoice == 2 {
			insertionSortCalon()
		} else {
			fmt.Println("Pilihan tidak valid")
			return
		}
		for i := 0; i < nCalon; i++ {
			fmt.Printf("Calon %d: Nama: %s, Umur: %d, Partai: %s\n", i, dataCalon[i].Name, dataCalon[i].Umur, dataCalon[i].PartaiName)
		}
	} else if choice == 3 {
		fmt.Println("Pilih metode sorting:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Masukan pilihan (1/2): ")
		fmt.Scan(&sortChoice)
		if sortChoice == 1 {
			selectionSortVoter()
		} else if sortChoice == 2 {
			insertionSortVoter()
		} else {
			fmt.Println("Pilihan tidak valid")
			return
		}
		for i := 0; i < nVoter; i++ {
			fmt.Printf("Pemilih %d: Nama: %s, Username: %s\n", i, dataVoter[i].Name, dataVoter[i].Username)
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func setThreshold() {
	fmt.Print("Masukkan threshold: ")
	fmt.Scan(&threshold)
}

func searchData() {
	var choice int
	fmt.Println("1. Cari Partai")
	fmt.Println("2. Cari Calon")
	fmt.Println("3. Cari Pemilih")
	fmt.Print("Masukan pilihan (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		var name string
		fmt.Print("Masukkan nama partai: ")
		fmt.Scan(&name)
		idx := sequentialSearchPartai(name)
		if idx != -1 {
			fmt.Printf("Partai ditemukan: %s\n", dataPartai[idx].Name)
		} else {
			fmt.Println("Partai tidak ditemukan")
		}
	} else if choice == 2 {
		var name string
		fmt.Print("Masukkan nama calon: ")
		fmt.Scan(&name)
		idx := binarySearchCalon(name)
		if idx != -1 {
			fmt.Printf("Calon ditemukan: Nama: %s, Umur: %d, Partai: %s\n", dataCalon[idx].Name, dataCalon[idx].Umur, dataCalon[idx].PartaiName)
		} else {
			fmt.Println("Calon tidak ditemukan")
		}
	} else if choice == 3 {
		var name string
		fmt.Print("Masukkan nama pemilih: ")
		fmt.Scan(&name)
		idx := sequentialSearchVoter(name)
		if idx != -1 {
			fmt.Printf("Pemilih ditemukan: Nama: %s, Username: %s\n", dataVoter[idx].Name, dataVoter[idx].Username)
		} else {
			fmt.Println("Pemilih tidak ditemukan")
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func displayCalon() {
	for i := 0; i < nCalon; i++ {
		fmt.Printf("Calon %d: Nama: %s, Umur: %d, Partai: %s\n", i, dataCalon[i].Name, dataCalon[i].Umur, dataCalon[i].PartaiName)
	}
}

func vote() {
	if !loginVoter() {
		fmt.Println("Username atau password salah")
		return
	}
	if !validateVotingTime() {
		fmt.Println("Waktu voting sudah habis")
		return
	}

	var idx int
	displayCalon()
	fmt.Print("Masukkan indeks calon yang ingin dipilih: ")
	fmt.Scan(&idx)

	if idx >= 0 && idx < nCalon {
		dataCalon[idx].NVoter++
		fmt.Println("Terima kasih sudah memilih!")
	} else {
		fmt.Println("Indeks tidak valid")
	}
}

func AdadiPartai(partaiName string) bool {
	for i := 0; i < nPartai; i++ {
		if dataPartai[i].Name == partaiName {
			return true
		}
	}
	return false
}

func sequentialSearchPartai(name string) int {
	for i := 0; i < nPartai; i++ {
		if dataPartai[i].Name == name {
			return i
		}
	}
	return -1
}

func sequentialSearchVoter(name string) int {
	for i := 0; i < nVoter; i++ {
		if dataVoter[i].Name == name {
			return i
		}
	}
	return -1
}

func binarySearchCalon(name string) int {
	low, high := 0, nCalon-1
	for low <= high {
		mid := (low + high) / 2
		if dataCalon[mid].Name == name {
			return mid
		} else if dataCalon[mid].Name < name {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func selectionSortPartai() {
	for i := 0; i < nPartai-1; i++ {
		minIdx := i
		for j := i + 1; j < nPartai; j++ {
			if dataPartai[j].Name < dataPartai[minIdx].Name {
				minIdx = j
			}
		}
		dataPartai[i], dataPartai[minIdx] = dataPartai[minIdx], dataPartai[i]
	}
}

func selectionSortCalon() {
	for i := 0; i < nCalon-1; i++ {
		minIdx := i
		for j := i + 1; j < nCalon; j++ {
			if dataCalon[j].Name < dataCalon[minIdx].Name {
				minIdx = j
			}
		}
		dataCalon[i], dataCalon[minIdx] = dataCalon[minIdx], dataCalon[i]
	}
}

func selectionSortVoter() {
	for i := 0; i < nVoter-1; i++ {
		minIdx := i
		for j := i + 1; j < nVoter; j++ {
			if dataVoter[j].Name < dataVoter[minIdx].Name {
				minIdx = j
			}
		}
		dataVoter[i], dataVoter[minIdx] = dataVoter[minIdx], dataVoter[i]
	}
}

func insertionSortPartai() {
	for i := 1; i < nPartai; i++ {
		key := dataPartai[i]
		j := i - 1
		for j >= 0 && dataPartai[j].Name > key.Name {
			dataPartai[j+1] = dataPartai[j]
			j--
		}
		dataPartai[j+1] = key
	}
}

func insertionSortCalon() {
	for i := 1; i < nCalon; i++ {
		key := dataCalon[i]
		j := i - 1
		for j >= 0 && dataCalon[j].Name > key.Name {
			dataCalon[j+1] = dataCalon[j]
			j--
		}
		dataCalon[j+1] = key
	}
}

func insertionSortVoter() {
	for i := 1; i < nVoter; i++ {
		key := dataVoter[i]
		j := i - 1
		for j >= 0 && dataVoter[j].Name > key.Name {
			dataVoter[j+1] = dataVoter[j]
			j--
		}
		dataVoter[j+1] = key
	}
}
