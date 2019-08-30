package poker

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_HOST       = "localhost"
	MONGODB_PORT       = "27017"
	MONGODB_USER       = ""
	MONGODB_PASS       = ""
	MONGODB_DB         = "poker"
	MONGODB_COLLECTION = "players"
)

type MongoPlayerStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoPlayerStore(env string) (*MongoPlayerStore, error) {

	// set client options
	conn := fmt.Sprintf("mongodb://%s:%s", MONGODB_HOST, MONGODB_PORT)
	clientOptions := options.Client().ApplyURI(conn)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	log.Printf("Connected to MongoDB on %s:%s", MONGODB_HOST, MONGODB_PORT)

	suffix := ""
	if env != "" {
		suffix = "_test"
	}

	collection := client.Database(MONGODB_DB).Collection(MONGODB_COLLECTION + suffix)
	return &MongoPlayerStore{client, collection}, nil
}

func (m *MongoPlayerStore) GetLeague() League {
	league := League{}

	filter := bson.D{{}}
	options := options.Find()
	options.SetSort(bson.D{{"wins", -1}}) // sort based on field "wins" descending

	cur, err := m.collection.Find(context.TODO(), filter, options)

	if err != nil {
		log.Printf("Error finding documents in collection: %v", err)
		return nil
	}

	for cur.Next(context.TODO()) {
		var elem Player
		err := cur.Decode(&elem)

		if err != nil {
			log.Printf("Error decoding a document")
		} else {
			league = append(league, elem)
		}

	}
	return league
}

func (m *MongoPlayerStore) GetPlayerScore(player string) (int, error) {

	filter := bson.D{{"name", player}}
	var result Player

	err := m.collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Printf("Error searching for a document: %v", err)

		switch err.Error() {
		case "mongo: no documents in result":
			return 0, ERRPLAYERNOTFOUND
		default:
			return 0, ERRINTERNALERROR
		}
	}

	log.Printf("Found a single document: %+v\n", result)
	return result.Wins, nil
}

func (m *MongoPlayerStore) RecordWin(player string) error {

	var updateResult *mongo.UpdateResult
	var err error
	_, err = m.GetPlayerScore(player)

	switch err {

	case ERRPLAYERNOTFOUND:
		doc := Player{player, 1}
		_, err = m.collection.InsertOne(context.TODO(), doc)

	case nil:
		filter := bson.D{{"name", player}}

		update := bson.D{
			{"$inc", bson.D{
				{"wins", 1},
			}},
		}
		updateResult, err = m.collection.UpdateOne(context.TODO(), filter, update)
		log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	default:
		log.Printf("Error checking the player score: %v", err)
		return ERRINTERNALERROR
	}

	if err != nil {
		log.Printf("Error recording player %s win: %v", player, err)
		return ERRINTERNALERROR
	}

	return nil

}

func (m *MongoPlayerStore) Teardown() {
	err := m.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatalf("Error disconnecting from DB %v", err)
	}

	log.Print("Connection to MongoDB closed...")
}

func (m *MongoPlayerStore) deleteCollection() error {
	deleteResult, err := m.collection.DeleteMany(context.TODO(), bson.D{{}})

	if err != nil {
		log.Printf("Error deleting documents in collection %v: %v", MONGODB_COLLECTION, err)
		return err
	}

	log.Printf("Deleted %v documents in the %v collection\n", MONGODB_COLLECTION, deleteResult.DeletedCount)
	return nil
}
