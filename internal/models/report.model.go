package models

type Report struct {
	TotalRevenue   int                 `json:"total_revenue"`
	TotalTransaksi int                 `json:"total_transaksi"`
	ProdukTerlaris ProdukTerlaris      `json:"produk_terlaris"`
	Transactions   []ReportTransaction `json:"transactions"`
}

type ReportTransaction struct {
	ID          int                  `json:"id"`
	TotalAmount int                  `json:"total_amount"`
	Details     []TransactionDetails `json:"details"`
}

type ProdukTerlaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
