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

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return user, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}
func (c *UserClient) Update(id string, update interface{}) (models.UserUpdate, error) {
	result := models.UserUpdate{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	user, err := c.Get(id)
	if err != nil {
		return result, err
	}
	var exist map[string]interface{}
	b, err := json.Marshal(user)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &exist)

	change := update.(map[string]interface{})
	for k := range change {
		if change[k] == exist[k] {
			delete(change, k)
		}
	}

	if len(change) == 0 {
		return result, nil
	}

	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
	if err != nil {
		return result, err
	}

	newUser, err := c.Get(id)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = res.ModifiedCount
	result.Result = newUser
	return result, nil
}
func (c *UserClient) Delete(id string) (models.UserDelete, error) {
	result := models.UserDelete{
		DeletedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}
	result.DeletedCount = res.DeletedCount
	return result, nil
}
func (c *UserClient) Get(id string) (models.User, error) {
	user := models.User{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (c *UserClient) Search(filter interface{}) ([]models.User, error) {
	users := []models.User{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil {
		return users, err
	}

	for cursor.Next(c.Ctx) {
		row := models.User{}
		cursor.Decode(&row)
		users = append(users, row)
	}

	return users, nil
}
