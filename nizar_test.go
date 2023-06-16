package nizar

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertUser(t *testing.T) {
	dbname := "tagihan"
	user := TagihanSPP{
		ID:            primitive.NewObjectID(),
		NamaMahasiswa: "Nizar",
		NIM:           "nizar@gmail.com",
		Semester:      "dua",
		BiayaSPP:      "10000000",
	}
	insertedID := InsertTagihanSPP(dbname, user)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestGetTagihanUser(t *testing.T) {
	NamaMahasiswa := "Nizar"
	hiya := GetTagihanSPP(NamaMahasiswa)
	fmt.Println(hiya)
}
