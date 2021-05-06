package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:55044/?readPreference=primary&ssl=false"))
	if err != nil {
		panic(err)
	}
	col := mc.Database("coolcar").Collection("account")
	insertRows(c, col)

	findRows(c, col)

}
func findRows(c context.Context, col *mongo.Collection) {
	res := col.FindOne(c, bson.M{
		"open_id": "123",
	})
	fmt.Printf("%+v\n", res)
	cur, err := col.Find(c, bson.M{})
	if err != nil {
		panic(err)
	}
	var row struct {
		ID     primitive.ObjectID `bson:"_id"`
		OPENID string             `bson:"open_id"`
	}
	for cur.Next(c) {
		err = cur.Decode(&row)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", row)
	}

	err = res.Decode(&row)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", row)

}

func insertRows(c context.Context, col *mongo.Collection) {
	res, err := col.InsertMany(c, []interface{}{
		bson.M{
			"open_id": "123",
		},
		bson.M{
			"open_id": "456",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}