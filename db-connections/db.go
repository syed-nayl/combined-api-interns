package db

import (
	models "combined-api/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitP() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	dbcon, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	dbcon.AutoMigrate(&models.Book{})

	return dbcon
}

func InitConnectionMongo() *mongo.Collection {
	//client option

	const connectionString = "mongodb+srv://umer:umer@cluster0.ofuslri.mongodb.net/?retryWrites=true&w=majority"
	const dbName = "netflix"
	const colName = "watchlist"

	// connect with monogoDB

	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection := client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")

	return collection
}
