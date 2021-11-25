package dto

import uuid "github.com/satori/go.uuid"

type RumahUpdateDTO struct {
	ID                    uuid.UUID
	Kota                  int8
	Kecamatan             string
	Desa                  string
	Nagari                string
	Jorong                string
	Dusun                 string
	Rt                    string
	NomorRumah            string
	Lat                   string
	Long                  string
	JumlahKeluarga        int32
	JumlahPenghuni        int32
	NamaKepalaKeluarga    string
	NomorKK               string
	StatusKepemilikan     string
	KeteranganKepemilikan string
	LuasRumah             int64
	Kondisi               string
	Jenis                 string
	File                  string
	Keterangan            string
	Ekstensi              string
}

type RumahCreateDTO struct {
	Kota                  int8
	Kecamatan             string
	Desa                  string
	Nagari                string
	Jorong                string
	Dusun                 string
	Rt                    string
	NomorRumah            string
	Lat                   string
	Long                  string
	JumlahKeluarga        int32
	JumlahPenghuni        int32
	NamaKepalaKeluarga    string
	NomorKK               string
	StatusKepemilikan     string
	KeteranganKepemilikan string
	LuasRumah             int64
	Kondisi               string
	Jenis                 string
	File                  string
	Keterangan            string
	Ekstensi              string
}

type RumahFindIdDTO struct {
	Id string `json:"id" form:"id" binding:"required"`
}

type RumahDeleteDTO struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
}

type RumahDeleteMultiID struct {
	Ids []string `json:"ids" binding:"required"`
}
