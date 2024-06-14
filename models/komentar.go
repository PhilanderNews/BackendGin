package models

type Komentar struct {
	ID        string `json:"id" bson:"id"`
	ID_berita string `json:"id_berita" bson:"id_berita"`
	Username  string `json:"username" bson:"username"`
	Tanggal   string `json:"tanggal" bson:"tanggal"`
	Komentar  string `json:"komentar" bson:"komentar"`
}
