package models

type Berita struct {
	ID       string   `json:"id" bson:"id"`
	Kategori string   `json:"kategori" bson:"kategori"`
	Judul    string   `json:"judul" bson:"judul"`
	Preview  string   `json:"preview" bson:"preview"`
	Konten   Paragraf `json:"konten" bson:"konten"`
	Penulis  string   `json:"penulis" bson:"penulis"`
	Sumber   string   `json:"sumber" bson:"sumber"`
	Image    string   `json:"image" bson:"image"`
	Waktu    string   `json:"waktu" bson:"waktu"`
}

type Paragraf struct {
	Paragraf1  string `json:"paragraf1" bson:"paragraf1"`
	Paragraf2  string `json:"paragraf2" bson:"paragraf2"`
	Paragraf3  string `json:"paragraf3" bson:"paragraf3"`
	Paragraf4  string `json:"paragraf4" bson:"paragraf4"`
	Paragraf5  string `json:"paragraf5" bson:"paragraf5"`
	Paragraf6  string `json:"paragraf6" bson:"paragraf6"`
	Paragraf7  string `json:"paragraf7" bson:"paragraf7"`
	Paragraf8  string `json:"paragraf8" bson:"paragraf8"`
	Paragraf9  string `json:"paragraf9" bson:"paragraf9"`
	Paragraf10 string `json:"paragraf10" bson:"paragraf10"`
}
