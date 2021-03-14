package repositoryimpl

import (
	"context"
	"fmt"
	"log"
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/domain/repository"
	"ms-dna/shared/customerror"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type dnaMongoRepository struct {
	db *mongo.Client
}

// New  es una implementación del repositorio basado en Mongo
func New(db *mongo.Client) repository.DnaRepository {
	return dnaMongoRepository{db: db}
}

// SaveDna saves a dna record
func (repo dnaMongoRepository) SaveDna(entity *entity.Dna) error {
	collection := repo.db.Database("test").Collection("DNAs")
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), entity)
	if err != nil {
		fmt.Println("Error insertando", err)
		log.Fatal(err)
	}
	fmt.Println("insertResult", insertResult.InsertedID)
	return nil
}

// Fetch return all records saved in storage
func (repo dnaMongoRepository) GetDnas() ([]*entity.Dna, error) {
	collection := repo.db.Database("test").Collection("DNAs")
	var results []*entity.Dna
	cursorR, error := collection.Find(context.TODO(), bson.D{})
	if error != nil {
		return nil, customerror.ErrMongo
	}
	for cursorR.Next(context.TODO()) {
		var elem entity.Dna
		err := cursorR.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cursorR.Err(); err != nil {
		log.Fatal(err)
	}

	// Cerrar el cursor
	cursorR.Close(context.TODO())
	return results, nil
}
