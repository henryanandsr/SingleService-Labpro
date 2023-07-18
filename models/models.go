package models

type Barang struct {
	ID                string `gorm:"primary_key;auto_increment" json:"id"`
	KodeBarang        string `gorm:"not null;unique" json:"kodeBarang"`
	NamaBarang        string `gorm:"size:255;not null" json:"namaBarang"`
	HargaBarang       int    `gorm:"not null;check:harga_barang > 0" json:"hargaBarang"`
	StokBarang        int    `gorm:"not null;check:stok_barang >= 0" json:"stokBarang"`
	PerusahaanPembuat string `gorm:"column:perusahaanpembuat;not null" json:"perusahaanPembuat"`
}

type Company struct {
	ID        string   `gorm:"primary_key;auto_increment" json:"id"`
	Nama      string   `gorm:"size:255;not null" json:"nama"`
	Alamat    string   `gorm:"size:255;not null" json:"alamat"`
	NoTelepon string   `gorm:"size:20;not null" json:"noTelepon"`
	KodePajak string   `gorm:"size:3;not null;check:kode_pajak ~ '^[A-Z]{3}$'" json:"kodePajak"`
	Barangs   []Barang `gorm:"foreignKey:PerusahaanPembuat;constraint:OnDelete:CASCADE"`
}

type User struct {
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
