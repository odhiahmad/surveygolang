package dto

type StatistikDTO struct {
	Rusak        int64 `json:"rusak,omitempty"`
	RumahKayu    int64 `json:"rumah_kayu,omitempty"`
	Baik         int64 `json:"baik,omitempty"`
	Jumlah       int64 `json:"jumlah,omitempty"`
	Permanen     int64 `json:"permanen,omitempty"`
	SemiPernamen int64 `json:"semi_permanen,omitempty"`
}
