package main

import "fmt"

func main() {

	// Deklarasi Variabel
	var nama_customer, nama_barang string
	var harga float32 = 25000
	var quantity int32
	var hasil_discount, total_harga float32

	// Input Nama Customer
	fmt.Print("Input Nama Customer: ")
	fmt.Scanf("%s\n", &nama_customer)

	// Input Nama Barang
	fmt.Print("Input Nama Barang: ")
	fmt.Scanf("%s\n", &nama_barang)

	// Input Quantity Barang
	fmt.Print("Input Quantity Barang: ")
	fmt.Scanf("%d\n", &quantity)

	// Kondisi Diskon, kalau kebih dari 5 dapat 10%, selain itu 2%
	switch {
		case quantity > 5:
		hasil_discount = 0.1 // 10%
	default:
		hasil_discount = 0.02 // 2%
	}

	// Proses
	sub_total := float32(quantity) * harga
	total_harga = sub_total - (hasil_discount * sub_total)

	// Tampilkan HasilHarga
	fmt.Println("hasil_discount: ", hasil_discount)
	fmt.Println("quantity: ", quantity)
	fmt.Println("harga: ", harga)
	fmt.Println("Total Harga adalah: ", total_harga)

}
