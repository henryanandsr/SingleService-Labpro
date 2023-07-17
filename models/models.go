package models

type Barang struct {
	ID                string `gorm:"primary_key;auto_increment" json:"id"`
	KodeBarang        string `gorm:"not null" json:"kodeBarang"`
	NamaBarang        string `gorm:"size:255;not null" json:"namaBarang"`
	HargaBarang       int    `gorm:"not null" json:"hargaBarang"`
	StokBarang        int    `gorm:"not null" json:"stokBarang"`
	PerusahaanPembuat string `gorm:"column:perusahaanpembuat;not null" json:"perusahaanPembuat"`
}

type Company struct {
	ID        string `gorm:"primary_key;auto_increment" json:"id"`
	Nama      string `gorm:"size:255;not null" json:"nama"`
	Alamat    string `gorm:"size:255;not null" json:"alamat"`
	NoTelepon string `gorm:"size:20;not null" json:"noTelepon"`
	KodePajak string `gorm:"size:3;not null" json:"kodePajak"`
}
