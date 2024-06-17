package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/ferhateroglu/go-fiber/databases"
	"github.com/ferhateroglu/go-fiber/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TodoRepository interface {
	Create(todo *models.Todo) error
	GetAll() ([]models.Todo, error)
	GetById(id string) (*models.Todo, error)
	Update(id string, todo *models.Todo) error
	Delete(id string) error
}

type MongoTodoRepository struct {
	collection *mongo.Collection
}

func NewMongoTodoRepository(mongoDatabase *databases.MongoDatabase) TodoRepository {
	return &MongoTodoRepository{
		collection: mongoDatabase.GetDatabase().Collection("todos"),
	}
}
func (r *MongoTodoRepository) Create(todo *models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.CreatedAt = time.Now().Unix()
	result, err := r.collection.InsertOne(ctx, todo)

	if err != nil || result.InsertedID == nil {
		return errors.New("failed to create todo")
	}

	return nil
}

func (r *MongoTodoRepository) GetAll() ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []models.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *MongoTodoRepository) GetById(id string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var todo models.Todo
	err = r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}

	return &todo, nil
}

func (r *MongoTodoRepository) Update(id string, todo *models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":   todo.Title,
			"content": todo.Content,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectId}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("todo not found")
	}

	return nil
}

func (r *MongoTodoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("todo not found")
	}

	return nil
}
