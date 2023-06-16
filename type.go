package nizar

type MahasiswaTag struct {
	ID            int    `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa string `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM           string `bson:"nim" json:"nim"`
}
type TagihanRegistrasi struct {
	ID              int     `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa   string  `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM             string  `bson:"nim" json:"nim"`
	Semester        int     `bson:"semester" json:"semester"`
	BiayaRegistrasi float64 `bson:"biaya_registrasi" json:"biaya_registrasi"`
}
type TagihanSPP struct {
	ID            int     `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa string  `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM           string  `bson:"nim" json:"nim"`
	Semester      int     `bson:"semester" json:"semester"`
	BiayaSPP      float64 `bson:"biaya_spp" json:"biaya_spp"`
}
type TagihanMakanan struct {
	ID            int     `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa string  `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM           string  `bson:"nim" json:"nim"`
	Bulan         string  `bson:"bulan" json:"bulan"`
	Tahun         int     `bson:"tahun" json:"tahun"`
	BiayaMakanan  float64 `bson:"biaya_makanan" json:"biaya_makanan"`
}

type TagihanPerpustakaan struct {
	ID                int     `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa     string  `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM               string  `bson:"nim" json:"nim"`
	Tahun             int     `bson:"tahun" json:"tahun"`
	BiayaPerpustakaan float64 `bson:"biaya_perpustakaan" json:"biaya_perpustakaan"`
}
