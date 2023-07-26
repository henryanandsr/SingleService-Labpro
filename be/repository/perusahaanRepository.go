package repositories

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PerusahaanRepository struct {
	Db *gorm.DB
}

func NewPerusahaanRepository() *PerusahaanRepository {
	db, err := initializers.GetDBInstance()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &PerusahaanRepository{Db: db}
}

func (r *PerusahaanRepository) GetPerusahaan(id string) (*model.Company, error) {
	var perusahaan model.Company
	if err := r.Db.Where("ID = ?", id).First(&perusahaan).Error; err != nil {
		return nil, err
	}

	return &perusahaan, nil
}

func (r *PerusahaanRepository) GetAllPerusahaans(q string) (*[]model.Company, error) {
	var companies []model.Company
	query := r.Db.Model(&model.Company{}).Select("id, nama, kode_pajak, alamat, no_telepon")

	if q != "" {
		query = query.Where("Nama LIKE ? OR kode_pajak LIKE ?", "%"+q+"%", "%"+q+"%")
	}

	if err := query.Find(&companies).Error; err != nil {
		return nil, err
	}

	return &companies, nil
}
func (r *PerusahaanRepository) DeletePerusahaan(id string) (*model.Company, error) {
	tx := r.Db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var company model.Company
	if err := tx.Where("ID = ?", id).First(&company).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var barangs []model.Barang
	if err := tx.Where("PerusahaanPembuat = ?", company.ID).Find(&barangs).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, barang := range barangs {
		if err := tx.Delete(&barang).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Delete(&company).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &company, nil
}

type PerusahaanPostRequest struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	NoTelp string `json:"no_telp"`
	Kode   string `json:"kode"`
}

func (r *PerusahaanRepository) CreatePerusahaan(request *PerusahaanPostRequest) (*model.Company, error) {
	company := &model.Company{
		ID:        uuid.New().String(),
		Nama:      request.Nama,
		Alamat:    request.Alamat,
		NoTelepon: request.NoTelp,
		KodePajak: request.Kode,
	}

	result := r.Db.Create(company)
	if result.Error != nil {
		return nil, result.Error
	}
	return company, nil
}

type PerusahaanUpdateRequest struct {
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	NoTelepon string `json:"no_telp"`
	KodePajak string `json:"kode"`
}

func (r *PerusahaanRepository) UpdatePerusahaan(id string, requestData *PerusahaanUpdateRequest) (*model.Company, error) {
	var company model.Company
	if err := r.Db.Where("ID = ?", id).First(&company).Error; err != nil {
		return nil, err
	}

	company.Nama = requestData.Nama
	company.Alamat = requestData.Alamat
	company.NoTelepon = requestData.NoTelepon
	company.KodePajak = requestData.KodePajak

	if err := r.Db.Save(&company).Error; err != nil {
		return nil, err
	}

	return &company, nil
}
