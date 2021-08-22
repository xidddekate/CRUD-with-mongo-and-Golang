package database

import (
	"context"
	"encoding/json"
	"go-users/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInterface interface {
	Insert(models.User) (models.User, error)
	Update(string, interface{}) (models.UserUpdate, error)
	Delete(string) (models.UserDelete, error)
	Get(string) (models.User, error)
	Search(interface{}) ([]models.User, error)
}

type UserClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *UserClient) Insert(docs models.User) (models.User, error) {
	user := models.User{}

	// Performing INSERT operation
	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return user, err
	}
	// Getting back _id of inserted object
	id := res.InsertedID.(primitive.ObjectID).Hex()
	// Retrieving the inserted object and returning it
	return c.Get(id)
}
func (c *UserClient) Update(id string, update interface{}) (models.UserUpdate, error) {
	result := models.UserUpdate{
		ModifiedCount: 0,
	}
	// Getting object _id (mongoDb format) from hex
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	// Retrieving the object from mongoDB
	user, err := c.Get(id)
	if err != nil {
		return result, err
	}
	// converting to JSON
	var exist map[string]interface{}
	b, err := json.Marshal(user)
	if err != nil {
		return result, err
	}
	// converting to Golang object type
	json.Unmarshal(b, &exist)

	// checking if object retrieved from mongoDB is similar to one that is to be updated (i.e param of this function)
	change := update.(map[string]interface{})
	for k := range change {
		if change[k] == exist[k] {
			delete(change, k)
		}
	}

	// if there is a match then return
	if len(change) == 0 {
		return result, nil
	}

	// else update the entry in mongoDB
	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
	if err != nil {
		return result, err
	}

	// Retrieving the updated object
	newUser, err := c.Get(id)
	if err != nil {
		return result, err
	}
	// Getting count of no. of objects updated
	result.ModifiedCount = res.ModifiedCount
	result.Result = newUser
	return result, nil
}
func (c *UserClient) Delete(id string) (models.UserDelete, error) {
	result := models.UserDelete{
		DeletedCount: 0,
	}
	// Getting object _id (mongoDb format) from hex
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	// Performing a delete operation
	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}

	// Getting count of no. of objects deleted
	result.DeletedCount = res.DeletedCount
	return result, nil
}
func (c *UserClient) Get(id string) (models.User, error) {
	user := models.User{}

	// Getting object _id (mongoDb format) from hex
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	// Performing a find operation on DB
	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (c *UserClient) Search(filter interface{}) ([]models.User, error) {
	users := []models.User{}

	// Loading the required filters in BSON format
	if filter == nil {
		filter = bson.M{}
	}

	// Obtaining cursor to iterate overs documents in mongo
	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil {
		return users, err
	}

	// Finding documents using cursor and appending them
	for cursor.Next(c.Ctx) {
		row := models.User{}
		cursor.Decode(&row)
		users = append(users, row)
	}

	return users, nil
}
