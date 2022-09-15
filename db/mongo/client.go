package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ahsanfayaz523/graphql-demo/graph/model"
)

const (
	collectionName = "animals"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		errors.New("error occurred while create connection")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

// Save - creates a new animal in database
func (db *DB) Save(animal *model.NewAnimal) *model.Animals {
	collection := db.client.Database("vetpartners").Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, animal)
	if err != nil {
		log().Errorf("error: %s", err)
	}

	return &model.Animals{
		ID:   res.InsertedID.(primitive.ObjectID).Hex(),
		Name: animal.Name,
		Age:  animal.Age,
	}
}

// FetchAnimal - get animal using ID from database
func (db *DB) FetchAnimal(id string) *model.Animals {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log().Errorf("error: %s", err)
	}
	collection := db.client.Database("vetpartners").Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	animal := model.Animals{}
	res.Decode(&animal)
	return &animal
}

// FetchAll - retrieve list of all animals
func (db *DB) FetchAll() []*model.Animals {
	collection := db.client.Database("vetpartners").Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log().Errorf("error: %s", err)
	}
	var animals []*model.Animals
	for cur.Next(ctx) {
		var animal *model.Animals
		err := cur.Decode(&animal)
		if err != nil {
			log().Errorf("error: %s", err)
		}
		animals = append(animals, animal)
	}
	return animals
}
