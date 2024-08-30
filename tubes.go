package main

import (
	"fmt"
)

// jenis kendaraan mobil/motor mobil/jam 5000 dan motor/jam 2000
// menambahkan variabel global untuk nilai penghasilan
// binRY SEARCHING SORTING

type petugas struct {
	namaPet, jenisKelaminPet string
	umurPet                  int
}

type pengunjung struct {
	platNomer, jenisKendaraan  string
	jamMasuk, jamKeluar, harga int
}

const NMAX int = 10

type petugasArr [NMAX]petugas
type pengunjungArr [NMAX]pengunjung

var totalPenghasilan int

func main() {
	var arrayPetugas petugasArr
	var arrayPengunjung pengunjungArr
	var jPetugas, jPengunjung int
	jPetugas = 0
	jPengunjung = 0
	login(arrayPetugas, arrayPengunjung, jPetugas, jPengunjung)
}

// login
func login(arrayPetugas petugasArr, arrayPengunjung pengunjungArr, jPetugas, jPengunjung int) {
	var a string
	for {
		fmt.Println("Selamat Datang silahkan memilih admin/pengunjung")
		fmt.Println("ketik logout jika ingin keluar")
		fmt.Scan(&a)
		if a == "admin" {
			admin(&arrayPetugas, &arrayPengunjung, &jPetugas, &jPengunjung)
		} else if a == "pengunjung" {
			pengguna(&arrayPetugas, &arrayPengunjung, &jPetugas, &jPengunjung)
		} else if a == "logout" {
			fmt.Println("program selesai")
			fmt.Println("goodbyee!!!")
			break
		} else {
			fmt.Println("input salah")
		}
	}
}

// admin
func admin(A *petugasArr, B *pengunjungArr, jPetugas, jPengunjung *int) {
	var a int
	for {
		fmt.Println("-----------------------------------")
		fmt.Println("1. Daftar Data pengunjung parkir")
		fmt.Println("2. Daftar Data petugas parkir")
		fmt.Println("3. Masukan data petugas parkir")
		fmt.Println("4. Total penghasilan")
		fmt.Println("5. Edit data petugas")
		fmt.Println("6. Cari petugas berdasarkan umur")
		fmt.Println("7. Log Out")
		fmt.Println("-----------------------------------")
		fmt.Scan(&a)
		if a == 1 {
			dpenp(*B, *jPengunjung)
		} else if a == 2 {
			dpetp(*A, *jPetugas)
		} else if a == 3 {
			jhp(A, jPetugas)
		} else if a == 4 {
			uhdpen(*B, *jPengunjung)
		} else if a == 5 {
			uhdpet(A, *jPetugas)
		} else if a == 6 {
			sort(A, *jPetugas)
			binary(*A, *jPetugas)
		} else if a == 7 {
			fmt.Println("Terimakasih")
			fmt.Println("kembali ke halaman login")
			return
		} else {
			fmt.Println("Masukan Salah")
		}
	}
}

func sort(A *petugasArr, n int) {
	var temp petugasArr
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[i].umurPet > A[j].umurPet {
				temp[i] = A[i]
				A[i] = A[j]
				A[j] = temp[i]
			}
		}
	}
}

func binary(A petugasArr, n int) {
	var x, s int
	fmt.Print("Masukkan umur : ")
	fmt.Scan(&x)
	fmt.Println("Petugas dengan umur ", x)
	fmt.Println("--------------------")
	s = binarySearch(A, x, n)
	if s == 0 {
		fmt.Println("data tidak ada")
	} else {
		for i := 0; i < n; i++ {
			if A[i].umurPet == x {
				fmt.Printf(A[i].namaPet)
			}
		}
	}
	fmt.Println("--------------------")
}

func binarySearch(A petugasArr, x, n int) int {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if A[mid].umurPet == x {
			return mid
		} else if A[mid].umurPet < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func dpenp(B pengunjungArr, jPengunjung int) {
	fmt.Println("--------------------------------------------------------")
	fmt.Println("No Plat      Jam Masuk      Jam Keluar     Jenis Kendaraan    jumlah biaya")
	for i := 0; i < jPengunjung; i++ {
		fmt.Printf("%s        %d             %d             %s                 %d\n", B[i].platNomer, B[i].jamMasuk, B[i].jamKeluar, B[i].jenisKendaraan, B[i].harga)
	}
}

func dpetp(A petugasArr, jPetugas int) {
	fmt.Println("---------------------------")
	fmt.Println("nama  umur   jenis kelamin")
	for i := 0; i < jPetugas; i++ {
		fmt.Printf("%s    %d    %s\n", A[i].namaPet, A[i].umurPet, A[i].jenisKelaminPet)
	}
}

func jhp(A *petugasArr, j *int) {
	var a, c string
	var b int
	fmt.Println("masukan dengan format nama, umur, jenis kelamin")
	fmt.Println("jika sudah masukan 0 0 0")
	for {
		if *j >= NMAX {
			fmt.Println("array sudah penuh")
			return
		}
		fmt.Print("Masukan nama : ")
		fmt.Scan(&a)
		fmt.Print("Masukan umur : ")
		fmt.Scan(&b)
		fmt.Print("Masukan jenis kelamin (L/P) : ")
		fmt.Scan(&c)
		if a == "0" && b == 0 && c == "0" {
			fmt.Println("input selesai")
			return
		} else if a == "0" || b == 0 || c == "0" || (c != "L" && c != "P") {
			fmt.Println("input salah")
		} else {
			A[*j].namaPet = a
			A[*j].umurPet = b
			A[*j].jenisKelaminPet = c
			*j++
			fmt.Println("input berhasil,memasukan data kembali")
		}
	}
}

func uhdpen(B pengunjungArr, jPengunjung int) {
	var mb, mt int = 0, 0
	fmt.Println("Total Penghasilan")
	dpenp(B, jPengunjung)
	for i := 0; i < jPengunjung; i++ {
		if B[i].jenisKendaraan == "mobil" {
			mb = mb + B[i].harga
		} else if B[i].jenisKendaraan == "motor" {
			mt = mt + B[i].harga
		}
	}
	sortPengunjung(&B, jPengunjung)
	fmt.Printf("Penghasilan parkir mobil adalah : %d\n", mb)
	fmt.Printf("Penghasilan parkir motor adalah : %d\n", mt)
	fmt.Print("Plat nomor pengunjung dengan biaya termahal adalah : ", B[jPengunjung-1].platNomer)
	fmt.Println()
	fmt.Print("Plat nomor pengunjung dengan biaya termurah adalah : ", B[0].platNomer)
	fmt.Println()
}

func sortPengunjung(A *pengunjungArr, n int) {
	var temp pengunjungArr
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[i].harga > A[j].harga {
				temp[i] = A[i]
				A[i] = A[j]
				A[j] = temp[i]
			}
		}
	}
}

func uhdpet(A *petugasArr, j int) {
	var a int
	fmt.Println("----------------------")
	fmt.Println("1. Ubah data petugas")
	fmt.Println("2. Hapus data petugas")
	fmt.Println("3. Kembali ke menu")
	fmt.Scan(&a)
	if a == 1 {
		ubah(A, j)
	} else if a == 2 {
		hapus(A, &j)
	} else if a == 3 {
		fmt.Println("kembali ke menu...")
		return
	} else {
		fmt.Println("input salah")
	}
}

func ubah(A *petugasArr, j int) {
	var a, d, c string
	var b int
	fmt.Print("masukkan nama petugas yang ingin di edit: ")
	fmt.Scan(&a)
	for i := 0; i < j; i++ {
		if A[i].namaPet == a {
			fmt.Println("Data ditemukan!!")
			fmt.Print("Ubah nama : ")
			fmt.Scan(&d)
			A[i].namaPet = d
			fmt.Print("Ubah umur : ")
			fmt.Scan(&b)
			A[i].umurPet = b
			fmt.Print("Ubah jenis kelamin : ")
			fmt.Scan(&c)
			A[i].jenisKelaminPet = c
			fmt.Println("Data Berhasil diubah!!")
		}
	}
	fmt.Println("data tidak ditemukan")
}

func hapus(A *petugasArr, j *int) {
	var a string
	fmt.Print("masukkan nama petugas yang ingin dihapus: ")
	fmt.Scan(&a)
	for i := 0; i < *j; i++ {
		if A[i].namaPet == a {
			fmt.Println("Data ditemukan")
			delete(A, j, i)
			fmt.Println("Data berhasil dihapus")
			return
		}
	}
	fmt.Println("data tidak ditemukan")
}

func delete(A *petugasArr, j *int, i int) {
	for n := i; n < *j; n++ {
		A[n] = A[n+1]
		if i == *j {
			A[i].namaPet = "0"
			A[i].jenisKelaminPet = "0"
			A[i].umurPet = 0
		}
	}
	*j--

	fmt.Println("Data berhasil dihapus")
}

func pengguna(arrayPetugas *petugasArr, arrayPengunjung *pengunjungArr, jPetugas, jPengunjung *int) {
	var a int
	for {
		fmt.Println("-----------------------------------")
		fmt.Println("1. Daftar Parkir")
		fmt.Println("2. Bayar transaksi")
		fmt.Println("3. Log Out")
		fmt.Println("-----------------------------------")
		fmt.Scan(&a)
		if a == 1 {
			daftarParkir(arrayPengunjung, jPengunjung)
		} else if a == 2 {
			bayarTransaksi(*arrayPengunjung, *jPengunjung)
		} else if a == 3 {
			fmt.Println("kembali kehalaman login")
			return
		} else {
			fmt.Println("Masukan Salah")
		}
	}

}

// Untuk mengetahui biaya parkir
func bayarTransaksi(A pengunjungArr, j int) {
	var plat string
	fmt.Print("Masukkan nomor plat kendaraan yang ingin dibayar:")
	fmt.Scan(&plat)
	for i := 0; i < j; i++ {
		if A[i].platNomer == plat {
			fmt.Printf("Biaya parkir untuk plat %s adalah %d\n", plat, A[i].harga)
			return
		}
	}
	fmt.Println("Kendaraan tidak ditemukan")
}

// Mendaftar untuk parkir
func daftarParkir(A *pengunjungArr, J *int) {
	var plat, jked string
	var jm, jk, i int
	fmt.Println("Masukan dengan format nomor plat, jam masuk, jam keluar, dan jenis kendaraan")
	fmt.Println("jika sudah, silakan masukan 0 0 0 0")
	fmt.Println("untuk memasukkan waktu masuk jika jam 05.00 ditulis 500")
	fmt.Println("dan untuk jam 16.00 ditulis 1600")
	fmt.Println("menggunakan waktu 24 jam")
	i = *J
	for i < NMAX {
		fmt.Print("Nomor plat: ")
		fmt.Scan(&plat)
		fmt.Print("Jam masuk: ")
		fmt.Scan(&jm)
		fmt.Print("Jam keluar: ")
		fmt.Scan(&jk)
		fmt.Print("Jenis kendaraan: ")
		fmt.Scan(&jked)
		if plat == "0" && jm == 0 && jk == 0 && jked == "0" {
			fmt.Println("Input selesai")
			*J = i
			return
		} else if plat == "0" || jm == 0 || jk == 0 || jked == "0" || (jked != "mobil" && jked != "motor") {
			fmt.Println("Input yang anda masukan salah, silakan ulangi kembali")
		} else {
			A[i].platNomer = plat
			A[i].jamMasuk = jm
			A[i].jamKeluar = jk
			A[i].jenisKendaraan = jked
			if jk < jm {
				if jked == "mobil" {
					A[i].harga = 5000 * ((jk+2400)/100 - jm/100)
				} else if jked == "motor" {
					A[i].harga = 2000 * ((jk+2400)/100 - jm/100)
				}
			} else if (jk-jm)/100 > 0 {
				if jked == "mobil" {
					A[i].harga = 5000 * (jk/100 - jm/100)
				} else if jked == "motor" {
					A[i].harga = 2000 * (jk/100 - jm/100)
				}
			} else {
				if jked == "mobil" {
					A[i].harga = 5000
				} else if jked == "motor" {
					A[i].harga = 2000
				}
			}
			i++
			fmt.Println("Berhasil mendaftar parkir")
		}
	}
}
