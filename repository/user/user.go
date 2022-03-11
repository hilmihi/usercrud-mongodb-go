package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-mongodb/entities"
)

type UserRepository struct {
	db *mongo.Client
}

func New(db *mongo.Client) *UserRepository {
	return &UserRepository{db: db}
}

func (tr *UserRepository) Get() ([]entities.User, error) {
	var users []entities.User
	usersCollection := tr.db.Database("myFirstDatabase").Collection("users")
	// retrieve all the documents that match the filter
	filter := bson.D{}
	cursor, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var user entities.User
		err = cursor.Decode(&user)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}
	fmt.Println(users)
	return users, nil
}

func (tr *UserRepository) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	usersCollection := tr.db.Database("myFirstDatabase").Collection("users")
	// retrieve all the documents that match the filter
	filter := bson.M{"email": email}
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		panic(err)
	}

	fmt.Println(user)
	return user, nil
}

func (tr *UserRepository) Create(user entities.User) (entities.User, error) {
	usersCollection := tr.db.Database("myFirstDatabase").Collection("users")
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
	return user, nil
}

func (tr *UserRepository) Update(user entities.User, id primitive.ObjectID) error {
	usersCollection := tr.db.Database("myFirstDatabase").Collection("users")
	// create the update query for the client
	update := bson.D{
		{"$set",
			bson.D{
				{"name", user.Name},
			},
		},
		{"$set",
			bson.D{
				{"email", user.Email},
			},
		},
		{"$set",
			bson.D{
				{"password", user.Password},
			},
		},
	}

	// execute the UpdateByID() function with the filter and update query
	result, err := usersCollection.UpdateByID(context.TODO(), id, update)
	// check for errors in the updating
	if err != nil {
		panic(err)
	}
	// display the number of documents updated
	fmt.Println("Number of documents updated:", result.ModifiedCount)

	return nil
}

func (tr *UserRepository) Delete(id primitive.ObjectID) error {
	usersCollection := tr.db.Database("myFirstDatabase").Collection("users")
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"_id", id},
				},
			},
		},
	}
	// delete the first document that match the filter
	result, err := usersCollection.DeleteOne(context.TODO(), filter)
	// check for errors in the deleting
	if err != nil {
		panic(err)
	}
	// display the number of documents deleted
	fmt.Println("deleting the first result from the search filter")
	fmt.Println("Number of documents deleted:", result.DeletedCount)
	return nil
}
