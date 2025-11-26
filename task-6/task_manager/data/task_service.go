package data

import (
	"context"
	"errors"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

func ConnectDB() *mongo.Client {
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return Client
}
func GetAllTasks() []models.Task {
	Client := ConnectDB()
	collection := Client.Database("taskManager").Collection("tasks")
	data, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var Tasks []models.Task
	if err = data.All(context.TODO(), &Tasks); err != nil {
		log.Fatal(err)

	}
	return Tasks
}

func GetTaskById(id string) ([]models.Task, error) {
	Client := ConnectDB()
	collection := Client.Database("taskManager").Collection("tasks")
	data, err := collection.Find(context.TODO(), bson.M{"ID": id})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var Tasks []models.Task
	if err = data.All(context.TODO(), &Tasks); err != nil {
		log.Fatal(err)
		return nil, err
	}
	if len(Tasks) == 0 {
		return nil, errors.New("task not found")
	}
	return Tasks, nil
}

func AddNewTask(task models.Task) []models.Task {

	Client := ConnectDB()
	collection := Client.Database("taskManager").Collection("tasks")
	_, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Fatal(err)
	}
	data, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var Tasks []models.Task
	if err = data.All(context.TODO(), &Tasks); err != nil {
		log.Fatal(err)
	}
	return Tasks
}

func UpdateTask(id string, updated models.Task) ([]models.Task, error) {
	Client := ConnectDB()
	collection := Client.Database("taskManager").Collection("tasks")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"ID": id}, bson.M{"$set": bson.M{"Title": updated.Title, "Description": updated.Description, "Deadline": updated.Deadline}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("task not found")
	}
	data, err := collection.Find(context.TODO(), bson.M{"ID": id})
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("task not found")
	}
	var Tasks []models.Task
	if err = data.All(context.TODO(), &Tasks); err != nil {
		log.Fatal(err)

	}

	return Tasks, nil
}

func DeleteTask(id string) error {
	Client := ConnectDB()
	collection := Client.Database("taskManager").Collection("tasks")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"ID": id})
	if err != nil {
		log.Fatal(err)
		return errors.New("task not found")
	}
	return nil
}
