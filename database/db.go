package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MyDataBase struct {
	client *mongo.Client
}

func (d *MyDataBase) DataBaseINIT() (client *mongo.Client, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading envirement files: ", err)
		return nil, fmt.Errorf("Error in loading envirement files: %s", err)
	}

	dbURL := os.Getenv("DATABASE")
	fmt.Println(dbURL)
	if dbURL != "" {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(dbURL).SetServerAPIOptions(serverAPI)
		client, err = mongo.Connect(context.TODO(), opts)

		if err != nil {
			log.Fatal("Error in connecting Database: ", err)
			return nil, fmt.Errorf("Error in connecting Database: %s", err)
		}

		err := client.Ping(context.Background(), nil)
		if err != nil {
			return nil, fmt.Errorf("error in connection with database %s", err)
		}

		d.client = client
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
		return d.client, nil
	} else {
		return nil, fmt.Errorf("database not found")
	}

}

func (d *MyDataBase) DbDisconnect() {
	d.client.Disconnect(context.TODO())
}
