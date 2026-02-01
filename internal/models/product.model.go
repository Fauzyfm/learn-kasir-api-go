package models

type Product struct {
    ID          int    `json:"id"`
    NamaBarang  string `json:"nama_barang" validate:"required"`
    HargaBarang int    `json:"harga_barang" validate:"required,min=1"`
    Stok        int    `json:"stok" validate:"required,min=0"`
    CategoryID  int    `json:"category_id" validate:"required,min=1"`
}


type ProductWithCategory struct {
    ID           int       `json:"id"`
    NamaBarang   string    `json:"nama_barang"`
    HargaBarang  int       `json:"harga_barang"`
    Stok         int       `json:"stok"`
    CategoryID   int       `json:"category_id"`
    Category     Categories `json:"category"`
}
