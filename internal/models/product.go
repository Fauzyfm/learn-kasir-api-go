package models

type Product struct {
	ID          int    `json:"id"`
	NamaBarang  string `json:"nama_barang"`
	HargaBarang int    `json:"harga_barang"`
	Stok        int    `json:"stok"`
}
