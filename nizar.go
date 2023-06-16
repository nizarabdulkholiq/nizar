package nizar

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertTagihanSPP(db string, tagihan TagihanSPP) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("tagihan").InsertOne(context.TODO(), tagihan)
	if err != nil {
		fmt.Printf("Inserttagihan: %v\n", err)
	}
	return insertResult.InsertedID
}
func GetTagihanSPP(NamaMahasiswa string) (data []TagihanSPP) {
	user := MongoConnect("tagihan").Collection("NamaMahasiswa")
	filter := bson.M{"NamaMahasiswa": NamaMahasiswa}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetNamaNamaMahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func InsertMahaTag(db *mongo.Database, collect string, NamaMahasiswa string, NIM string) (InsertedID interface{}) {
	var srt MahasiswaTag
	srt.NamaMahasiswa = NamaMahasiswa
	srt.NIM = NIM
	return InsertOneDoc(db, collect, srt)
}
