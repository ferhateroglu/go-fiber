package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/ferhateroglu/go-fiber/internal/models"
	"github.com/ferhateroglu/go-fiber/internal/repositories"
	"github.com/ferhateroglu/go-fiber/pkg/databases"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mockMongoDatabase struct {
	DB *mongo.Database
}

func (m *mockMongoDatabase) GetDatabase() *mongo.Database {
	return m.DB
}

func (m *mockMongoDatabase) Close(ctx context.Context) error {
	return m.DB.Client().Disconnect(ctx)
}

func setupInMemoryMongoDB(t *testing.T) (databases.Database, func()) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017/?ephemeralForTest=true"))
	if err != nil {
		t.Fatalf("Failed to connect to in-memory MongoDB: %v", err)
	}

	db := client.Database("test_db")
	mockDB := &mockMongoDatabase{DB: db}

	return mockDB, func() {
		if err = mockDB.Close(context.Background()); err != nil {
			t.Fatalf("Failed to disconnect from in-memory MongoDB: %v", err)
		}
	}
}
func TestTodoRepository(t *testing.T) {
	mockDB, cleanup := setupInMemoryMongoDB(t)
	defer cleanup()

	repo := repositories.NewTodoRepository(mockDB)

	t.Run("CreateAndGetById", func(t *testing.T) {
		todo := &models.Todo{Title: "Test Todo", Content: "Test Content"}
		todo.Id = bson.NewObjectID()
		todo.CreatedAt = time.Now().Unix()

		err := repo.Create(todo)
		assert.NoError(t, err)
		assert.NotEmpty(t, todo.Id)

		fetchedTodo, err := repo.GetById(todo.Id.Hex())
		assert.NoError(t, err)
		assert.Equal(t, todo.Title, fetchedTodo.Title)
		assert.Equal(t, todo.Content, fetchedTodo.Content)
	})

	t.Run("GetAll", func(t *testing.T) {
		mockDB.GetDatabase().Collection("todos").DeleteMany(context.Background(), bson.M{})

		todo1 := &models.Todo{Title: "Todo 1", Content: "Content 1"}
		todo2 := &models.Todo{Title: "Todo 2", Content: "Content 2"}
		repo.Create(todo1)
		repo.Create(todo2)

		todos, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, todos, 2)
	})

	t.Run("Update", func(t *testing.T) {
		todo := &models.Todo{Title: "Original Title", Content: "Original Content"}
		repo.Create(todo)

		updatedTodo := &models.Todo{Title: "Updated Title", Content: "Updated Content"}
		err := repo.Update(todo.Id.Hex(), updatedTodo)
		assert.NoError(t, err)

		fetchedTodo, err := repo.GetById(todo.Id.Hex())
		assert.NoError(t, err)
		assert.Equal(t, "Updated Title", fetchedTodo.Title)
		assert.Equal(t, "Updated Content", fetchedTodo.Content)
	})

	t.Run("Delete", func(t *testing.T) {
		todo := &models.Todo{Title: "To be deleted", Content: "Content"}
		repo.Create(todo)

		err := repo.Delete(todo.Id.Hex())
		assert.NoError(t, err)

		_, err = repo.GetById(todo.Id.Hex())
		assert.Error(t, err)
	})
}
