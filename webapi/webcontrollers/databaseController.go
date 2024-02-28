package webcontrollers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"bridge/webapi/webentities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseController struct {
	DB *mongo.Database
}

func NewDatabaseController(db *mongo.Database) *DatabaseController {
	fmt.Println("Created Database Controller")
	return &DatabaseController{
		DB: db,
	}
}

func CreateDBInstance() (*mongo.Database, error) {
	log.Println("Connecting to MongoDB instance...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connectionString := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(connectionString)

	var err error
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		return nil, fmt.Errorf("error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	dbName := os.Getenv("DB_NAME")
	// dbPlayer := os.Getenv("COLLECTION_PLAYER")
	// dbTable := os.Getenv("COLLECTION_TABLE")

	db := client.Database(dbName)

	fmt.Println("collection instance created")

	return db, nil
}

func (db *DatabaseController) GetPlayerByChatID(chatID int64, player *webentities.Player) error {
	collectionName := os.Getenv("COLLECTION_PLAYER")
	collection := db.DB.Collection(collectionName)
	filter := bson.M{"chat_id": chatID}
	err := collection.FindOne(context.Background(), filter).Decode(player)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("player with chat ID %d not found", chatID)
		}
		return fmt.Errorf("failed to get collection %s", collectionName)
	}
	return nil
}

func (db *DatabaseController) DeletePlayersByTableID(tableID uint32) error {
	collectionName := os.Getenv("COLLECTION_PLAYER")
	collection := db.DB.Collection(collectionName)
	filter := bson.M{"table_id": tableID}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v document(s)\n", result.DeletedCount)
	return nil
}

// func (db *DatabaseController) DeletePlayerByChatID(tableID uint32) error {
// 	collectionName := os.Getenv("COLLECTION_TABLE")
// 	collection := db.DB.Collection(collectionName)
// 	filter := bson.M{"table_id": tableID}
// 	result, err := collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("Deleted %v document(s)\n", result.DeletedCount)
// 	return nil
// }

func (db *DatabaseController) GetTableByTableID(tableID uint32, table *webentities.Table) error {
	collectionName := os.Getenv("COLLECTION_TABLE")
	collection := db.DB.Collection(collectionName)
	filter := bson.M{"table_id": tableID}
	err := collection.FindOne(context.Background(), filter).Decode(table)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("table with table ID %d not found", tableID)
		}
		return fmt.Errorf("failed to get collection %s", collectionName)
	}
	return nil
}

func (db *DatabaseController) InsertPlayer(player *webentities.Player) error {
	collectionName := os.Getenv("COLLECTION_PLAYER")
	collection := db.DB.Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		return err
	}
	return nil
}

func (db *DatabaseController) InsertTable(table *webentities.Table) error {
	collectionName := os.Getenv("COLLECTION_TABLE")
	collection := db.DB.Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), table)
	if err != nil {
		return err
	}
	return nil
}

func (db *DatabaseController) UpdateTable(table *webentities.Table) {
	collectionName := os.Getenv("COLLECTION_TABLE")
	collection := db.DB.Collection(collectionName)
	filter := bson.M{"table_id": table.TableID}
	count := table.Count + 1
	update := bson.M{"$set": bson.M{"count": count}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated %v document(s)\n", result.ModifiedCount)

}
