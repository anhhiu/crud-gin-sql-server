package models

type Student struct {
	Id            int     `json:"id"`
	Ten           string  `json:"ten"`
	Tuoi          int     `json:"tuoi"`
	Lop           string  `json:"lop"`
	DiemTrungBinh float64 `json:"diem_trung_binh"`
}
