package main

import (
	"fmt"
	"strconv"
)

const MAX = 10

type users struct {
	users     string
	passwords string
	saldo     float64
}

var data = [MAX]users{
	{"Daffa", "123", 100.000},
	{"Paffu", "123", 100.000},
}

func main() {
	var username, valid int
	listAuth(&username, &valid)
	
}

// Prosedur

func register(username, valid *int) {
	var user, pass string
	var check int
	fmt.Println("Masukan Username:")
	fmt.Scan(&user)
	fmt.Println("Masukan Password:")
	fmt.Scan(&pass)

	// Shifting the data array to the right to insert new user at the beginning
	for i := MAX - 1; i > 0; i-- {
		data[i] = data[i-1]
	}
	// Adding new user to the beginning of the array
	data[0] = users{user, pass, 0.0}
	check = checkUser(user, pass, username)
	*valid = check
	if check >= 0 {
		fmt.Println("Registration successful")
	} else {
		fmt.Println("Username atau Password Salah")
		auth(username, valid)
	}
}

func auth(username, valid *int) {
	var user, pass string
	var check int
	fmt.Println("Masukan Username:")
	fmt.Scan(&user)
	fmt.Println("Masukan Password:")
	fmt.Scan(&pass)
	check = checkUser(user, pass, username)
	*valid = check
	fmt.Println(check)
	if check >= 0 {
		fmt.Println("Authentication successful")
	} else {
		fmt.Println("Username atau Password Salah")
		auth(username, valid)
	}

}

func listAuth(username, valid *int) {
	var pilih int
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Masukan Angka yang ingin dipilih :")

	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		register(username, valid)
	case 2:
		auth(username, valid)
	case 3:
		logout()
	default:
		fmt.Println("Invalid choice")
	}

	if *username >= 0 {
		list(*username)
	}
}

func list(i int) {
	// IS. masukan berupa index dari username
	// FS. keluaran berupa tampilan list dan input berupa pilihan dari list yang ada
	var pilih int
	fmt.Printf("%s Saldo : %.3f \n", data[i].users, data[i].saldo)
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Tambah Uang")
	fmt.Println("2. Bayar")
	fmt.Println("3. Transfer")
	fmt.Println("4. Lainnya")
	fmt.Println("5. Keluar")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Masukan Angka yang ingin dipilih :")

	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		addMoney(i)
	case 2:
		pay(i)
	case 3:
		transfer(i)
	case 4:
		listLainnya(i)
	case 5:
		logout()
	default:
		fmt.Println("Invalid choice")
		list(i)
	}
}

func listLainnya(i int) {
	// FS. keluaran berupa tampilan list dan input berupa pilihan dari list yang ada
	var pilih int
	fmt.Printf("%s Saldo : %.3f \n", data[i].users, data[i].saldo)
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Bayar BPJS")
	fmt.Println("2. Bayar Listrik")
	fmt.Println("3. Bayar Pulsa")
	fmt.Println("3. Kembali ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Masukan Angka yang ingin dipilih :")

	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		bayarBPJS(i)
	case 2:
		beliListrik(i)
	case 3:
		beliPulsa(i)
	case 4:
		list(i)
	default:
		fmt.Println("Invalid choice")
		list(i)
	}
}

func addMoney(i int) {
	// IS. Masukan Berupa Indeks dari authentications
	// FS. Keluaran berupa Input untuk menambahkan saldo

	// Input yang berupa string kemudian akan diubah menjadi float64
	var tambahUang string
	fmt.Println("Masukan Jumlah Uang :")
	fmt.Scan(&tambahUang)

	// digunakan untuk mengfilter jika user memasukan input berupa huruf kemudian
	// akan mengeluarkan input tidak falid jika user memasukan input berupa huruf
	tambah, err := strconv.ParseFloat(tambahUang, 64)
	if err != nil {
		fmt.Println("Input Tidak Valid")
		return
	}

	if tambah > 0 {
		data[i].saldo += tambah
	} else {
		fmt.Println("Input Tidak Valid")
		list(i)
	}

	fmt.Printf("Saldo %s sekarang: %.3f\n", data[i].users, data[i].saldo)

	list(i)
}

func pay(i int) {
	var bayar float64
	fmt.Println("Masukan Jumlah Uang Yang Ingin dibayarkan :")
	fmt.Scan(&bayar)
	if bayar < data[i].saldo {
		data[i].saldo -= bayar
		fmt.Println("Pembayaran Berhasil")
		list(i)
	} else {
		fmt.Println("Saldo Tidak Cukup")
		list(i)
	}
	
}

func transfer(i int) {
	var saldo float64
	var nama string
	fmt.Println("Masukan Nama Akun Yang Ingin Dituju :")
	fmt.Scan(&nama)
	fmt.Println("Masukan Nominal Uang Yang Ingin Ditransfer :")
	fmt.Scan(&saldo)


	if saldo < data[i].saldo {
		for x := 0; x < MAX; x++ {
			if nama == data[x].users {
				data[x].saldo += saldo 
				data[i].saldo -= saldo 
				list(i)
			} 
		}
		list(i)
	} else {
		fmt.Println("Saldo Kurang")
		list(i)
	}

	
}

// function

func checkUser(user, pass string, username *int) int {
	// IS. Masukan Berupa Username dan Password
	// FS. Keluaran Berupa Indeks dari Username yang terdapat pada array Jika tidak terdapat pada array maka akan mengeluarkan -1

	// Perulangan untuk mengecek apakah user terdapat pada array
	for i := 0; i < MAX; i++ {
		if user == data[i].users && pass == data[i].passwords {
			*username = i

			return i
		}

	}

	return -1
}

func logout() {
	var username, valid int
	username, valid = -1, -1
	fmt.Println("Akun Berhasil Keluar")
	listAuth(&username, &valid)
	
}

func bayarBPJS(i int) {
	// IS. Masukan Berupa id pengguna dan sampai kapan ingin membayar
	var pilih, id int
	var total, harga, bulan float64

	harga = 150.000
	fmt.Println("Masukan ID BPJS :")
	fmt.Scan(&id)
	fmt.Println("Bayar berapa bulan? :")
	fmt.Scan(&bulan)

	total = harga * bulan

	fmt.Printf("Total yang harus dibayar : %d", total)
	fmt.Println("Lanjutkan pembayaran?")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("-----------------------------------------------")

	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		if total < data[i].saldo {
			data[i].saldo -= total
			fmt.Println("Pembayaran Berhasil")
		} else {
			fmt.Println("Saldo Tidak Cukup")
		}
		list(i)
	case 2:
		list(i)
	default:
		fmt.Println("Invalid choice")
		bayarBPJS(i)
	}
}

func beliPulsa(i int) {
	// IS. Masukan Berupa nomor tele[pn]

	var pulsa, pilih, noTelp int
	var admin, total float64

	admin = 2.500

	fmt.Println("Masukan Nomor Telepon :")
	fmt.Scan(&noTelp)

	fmt.Printf("Mau beli pulsa berapa?")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. 10.000")
	fmt.Println("2. 25.000")
	fmt.Println("3. 50.000")
	fmt.Println("4. 100.000")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Masukan Angka yang ingin dipilih :")

	fmt.Scan(&pulsa)
	switch pulsa {
	case 1:
		total = 10.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliPulsa(i)
		}
	case 2:
		total = 25.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliPulsa(i)
		}
	case 3:
		total = 50.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliPulsa(i)
		}
	case 4:
		total = 100.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliPulsa(i)
		}
	}

}

func beliListrik(i int) {
	// IS. Masukan Berupa nomor tele[pn]

	var listrik, pilih, noMeter int
	var admin, total float64

	admin = 2.500

	fmt.Println("Masukan Nomor Meteran :")
	fmt.Scan(&noMeter)

	fmt.Printf("Mau beli listrik berapa?")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. 20.000")
	fmt.Println("2. 50.000")
	fmt.Println("3. 100.000")
	fmt.Println("4. 200.000")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Masukan Angka yang ingin dipilih :")

	fmt.Scan(&listrik)
	switch listrik {
	case 1:
		total = 20.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
				fmt.Println("Nomor Meter :", noMeter)
				fmt.Println("Token : 123456020")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliListrik(i)
		}
	case 2:
		total = 50.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
				fmt.Println("Nomor Meter :", noMeter)
				fmt.Println("Token : 123456050")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliListrik(i)
		}
	case 3:
		total = 100.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
				fmt.Println("Nomor Meter :", noMeter)
				fmt.Println("Token : 123456100")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliListrik(i)
		}
	case 4:
		total = 200.000 + admin

		fmt.Printf("Total yang harus dibayar : %d", total)

		fmt.Println("Lanjutkan pembayaran?")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Println("-----------------------------------------------")

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if total < data[i].saldo {
				data[i].saldo -= total
				fmt.Println("Pembayaran Berhasil")
				fmt.Println("Nomor Meter :", noMeter)
				fmt.Println("Token : 123456200")
			} else {
				fmt.Println("Saldo Tidak Cukup")
			}
			list(i)
		case 2:
			list(i)
		default:
			fmt.Println("Invalid choice")
			beliListrik(i)
		}
	}

}
