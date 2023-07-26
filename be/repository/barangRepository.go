package repositories

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BarangRepository struct {
	Db *gorm.DB
}

func NewBarangRepository() *BarangRepository {
	db, err := initializers.GetDBInstance()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &BarangRepository{Db: db}
}

func (r *BarangRepository) GetBarang(id string) (*model.Barang, error) {
	var barang model.Barang
	if err := r.Db.Where("ID = ?", id).First(&barang).Error; err != nil {
		return nil, err
	}

	return &barang, nil
}

func (r *BarangRepository) GetAllBarangs(q string, perusahaan string) (*[]model.Barang, error) {
	var barangs []model.Barang
	query := r.Db.Model(&model.Barang{}).Select("id, kode_barang, nama_barang, harga_barang, stok_barang, perusahaanpembuat")

	if q != "" {
		query = query.Where("nama_barang LIKE ? OR kode_barang LIKE ?", "%"+q+"%", "%"+q+"%")
	}

	if perusahaan != "" {
		query = query.Where("PerusahaanPembuat = ?", perusahaan)
	}

	if err := query.Find(&barangs).Error; err != nil {
		return nil, err
	}

	return &barangs, nil
}

func (r *BarangRepository) DeleteBarang(id string) (*model.Barang, error) {
	var barang model.Barang
	if err := r.Db.Where("ID = ?", id).First(&barang).Error; err != nil {
		return nil, err
	}

	r.Db.Delete(&barang)
	return &barang, nil
}

type BarangPostRequest struct {
	NamaBarang   string `json:"nama"`
	HargaBarang  int    `json:"harga"`
	StokBarang   int    `json:"stok"`
	PerusahaanID string `json:"perusahaan_id"`
	KodeBarang   string `json:"kode"`
}

func (r *BarangRepository) CreateBarang(request *BarangPostRequest) (*model.Barang, error) {
	existingBarang := &model.Barang{}
	result := r.Db.Where("kode_barang = ?", request.KodeBarang).First(existingBarang)
	if result.Error == nil {
		return nil, errors.New("Barang with the same KodeBarang already exists")
	}

	barang := &model.Barang{
		ID:                uuid.New().String(),
		KodeBarang:        request.KodeBarang,
		NamaBarang:        request.NamaBarang,
		HargaBarang:       request.HargaBarang,
		StokBarang:        request.StokBarang,
		PerusahaanPembuat: request.PerusahaanID,
	}
	result = r.Db.Create(barang)
	if result.Error != nil {
		return nil, result.Error
	}

	return barang, nil
}
type BarangUpdateRequest struct {
	NamaBarang        string `json:"nama"`
	HargaBarang       int    `json:"harga"`
	StokBarang        int    `json:"stok"`
	PerusahaanPembuat string `json:"perusahaan_id"`
	KodeBarang        string `json:"kode"`
}

func (r *BarangRepository) UpdateBarang(id string, requestData *BarangUpdateRequest) (*model.Barang, error) {
	var barang model.Barang
	if err := r.Db.Where("ID = ?", id).First(&barang).Error; err != nil {
		return nil, err
	}

	barang.NamaBarang = requestData.NamaBarang
	barang.HargaBarang = requestData.HargaBarang
	barang.StokBarang = requestData.StokBarang
	barang.PerusahaanPembuat = requestData.PerusahaanPembuat
	barang.KodeBarang = requestData.KodeBarang

	if err := r.Db.Save(&barang).Error; err != nil {
		return nil, err
	}

	return &barang, nil
}

func (r *BarangRepository) UpdateStokBarang(id string, stokBarang int) (*model.Barang, error) {
	var barang model.Barang
	if err := r.Db.Where("ID = ?", id).First(&barang).Error; err != nil {
		return nil, err
	}

	barang.StokBarang = stokBarang
	if err := r.Db.Save(&barang).Error; err != nil {
		return nil, err
	}

	return &barang, nil
}
