package repository

import (
	"github.com/odhiahmad/apiuser/dto"
	"github.com/odhiahmad/apiuser/entity"
	"gorm.io/gorm"
)

type RumahRepository interface {
	InsertRumah(rumah entity.Rumah) entity.Rumah
	UpdateRumah(rumah entity.Rumah) entity.Rumah
	FindAll(pagination dto.Pagination) (*dto.Pagination, error)
	FindAllByKota(pagination dto.Pagination) (*dto.Pagination, error)
	FindById(id string) entity.Rumah
	FindByKota(id string) []entity.Rumah
	Delete(id uint64) entity.Rumah
	Statistik(statistik dto.StatistikDTO) (*dto.StatistikDTO, error)
	DeleteByIds(id *[]string) entity.Rumah
	IsDuplicateNomorKK(nomorKK string) (tx *gorm.DB)
}

type rumahConnection struct {
	connection *gorm.DB
}

func NewRumahRepository(db *gorm.DB) RumahRepository {
	return &rumahConnection{
		connection: db,
	}
}

func (db *rumahConnection) InsertRumah(rumah entity.Rumah) entity.Rumah {
	db.connection.Save(&rumah)

	return rumah
}

func (db *rumahConnection) UpdateRumah(rumah entity.Rumah) entity.Rumah {

	db.connection.Where("id = ?", rumah.ID).Save(&rumah)

	return rumah
}

func (db *rumahConnection) FindById(id string) entity.Rumah {
	var rumah entity.Rumah
	db.connection.Where("id = ?", id).Take(&rumah)

	return rumah
}

func (db *rumahConnection) FindByKota(id string) []entity.Rumah {
	var rumah []entity.Rumah
	db.connection.Where("kota = ?", id).Find(&rumah)

	return rumah
}

func (db *rumahConnection) Delete(id uint64) entity.Rumah {
	var rumah entity.Rumah
	db.connection.Where("id = ?", id).Delete(&rumah)

	return rumah
}

func (db *rumahConnection) DeleteByIds(ids *[]string) entity.Rumah {

	var rumah entity.Rumah
	db.connection.Where("ID IN (?)", *ids).Delete(&rumah)

	return rumah
}

func (db *rumahConnection) FindAll(pagination dto.Pagination) (*dto.Pagination, error) {
	var rumah []entity.Rumah

	db.connection.Where("nomor_kk like ? or nama_kepala_keluarga like ?", pagination.Cari+"%", pagination.Cari+"%").Scopes(Paginate(rumah, &pagination, db.connection)).Find(&rumah)
	pagination.Rows = rumah

	return &pagination, nil
}

func (db *rumahConnection) FindAllByKota(pagination dto.Pagination) (*dto.Pagination, error) {
	var rumah []entity.Rumah

	db.connection.Where("kota = ?", pagination.Kota).Where("jorong like ? or desa like ? or nagari like ? or nomor_kk like ? or nama_kepala_keluarga like ?", pagination.Cari+"%", pagination.Cari+"%", pagination.Cari+"%", pagination.Cari+"%", pagination.Cari+"%").Scopes(Paginate(rumah, &pagination, db.connection)).Find(&rumah)
	pagination.Rows = rumah

	return &pagination, nil
}

func (db *rumahConnection) Statistik(statistik dto.StatistikDTO) (*dto.StatistikDTO, error) {

	var rumah []entity.Rumah
	var hitungPermanen int64 = 0
	var hitungSemi int64 = 0
	var hitungKayu int64 = 0
	var hitungBaik int64 = 0
	var hitungRusak int64 = 0
	var hitungJumlah int64 = 0

	db.connection.Where("jenis = 'Permanen'").Find(&rumah).Count(&hitungPermanen)
	db.connection.Where("jenis = 'Semi Permanen'").Find(&rumah).Count(&hitungSemi)
	db.connection.Where("jenis = 'Rumah Kayu'").Find(&rumah).Count(&hitungKayu)
	db.connection.Where("kondisi = 'Baik'").Find(&rumah).Count(&hitungBaik)
	db.connection.Where("kondisi = 'Rusak'").Find(&rumah).Count(&hitungRusak)
	db.connection.Find(&rumah).Count(&hitungJumlah)

	if hitungSemi == 0 {
		statistik.SemiPernamen = -1
	} else {
		statistik.SemiPernamen = hitungSemi
	}

	if hitungPermanen == 0 {
		statistik.Permanen = -1
	} else {
		statistik.Permanen = hitungPermanen
	}

	if hitungBaik == 0 {
		statistik.Baik = -1
	} else {
		statistik.Baik = hitungBaik
	}

	if hitungRusak == 0 {
		statistik.Rusak = -1
	} else {
		statistik.Rusak = hitungRusak
	}

	if hitungJumlah == 0 {
		statistik.Jumlah = -1
	} else {
		statistik.Jumlah = hitungJumlah
	}

	if hitungKayu == 0 {
		statistik.RumahKayu = -1
	} else {
		statistik.RumahKayu = hitungKayu
	}

	return &statistik, nil
}

func (db *rumahConnection) IsDuplicateNomorKK(nomorKK string) (tx *gorm.DB) {
	var rumah entity.Rumah
	return db.connection.Where("nomor_kk = ?", nomorKK).Take(&rumah)
}
