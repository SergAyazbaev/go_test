// package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "go_test/db"

	"database/sql"
	_ "github.com/lib/pq"



	// "github.com/urfave/cli/v2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// func init() {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection = client.Database("tasker").Collection("tasks")
// }

// type Trainer struct {
// 	Name string
// 	Age  int
// 	City string
// }
