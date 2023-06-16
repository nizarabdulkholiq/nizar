package nizar

import (
	"testing"
)

func TestInsertDataTagihan(t *testing.T) {
	// Koneksi ke database MongoDB
	db := "your-database-name"

	// Data tagihan yang akan disisipkan
	datatagihan := nizar.DataTagihan{
		ID:         "1",
		Nama:       "John Doe",
		Tagihan:    500000,
		SudahBayar: false,
	}

	// Menyisipkan data tagihan
	insertedID := nizar.InsertDataTagihan(db, datatagihan)

	// Verifikasi apakah data berhasil disisipkan
	if insertedID == nil {
		t.Errorf("Failed to insert data tagihan")
	}
}

func TestInsertDataSudah(t *testing.T) {
	// Koneksi ke database MongoDB
	db := "your-database-name"

	// Data tagihan yang sudah dibayarkan
	datasudah := nizar.DataSudah{
		ID:           "1",
		Nama:         "John Doe",
		Tagihan:      500000,
		TanggalBayar: "2023-06-16",
	}

	// Menyisipkan data tagihan yang sudah dibayarkan
	insertedID := nizar.InsertDataSudah(db, datasudah)

	// Verifikasi apakah data berhasil disisipkan
	if insertedID == nil {
		t.Errorf("Failed to insert data tagihan yang sudah dibayarkan")
	}
}

func TestInsertDataBelum(t *testing.T) {
	// Koneksi ke database MongoDB
	db := "your-database-name"

	// Data tagihan yang belum dibayarkan
	databelum := nizar.DataBelum{
		ID:           "1",
		Nama:         "John Doe",
		Tagihan:      500000,
		TenggatBayar: "2023-06-30",
	}

	// Menyisipkan data tagihan yang belum dibayarkan
	insertedID := nizar.InsertDataBelum(db, databelum)

	// Verifikasi apakah data berhasil disisipkan
	if insertedID == nil {
		t.Errorf("Failed to insert data tagihan yang belum dibayarkan")
	}
}
