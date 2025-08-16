package lib

import (
	"context"
	"errors"
	"log"
	model "pkg/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Action   string
	Database string
	Table    string
	Payload  any
	Filter   *bson.M
}

func (r *DBOpts) Init() (ctx context.Context, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ContextDeadline*time.Second)
	defer cancel()

	return ctx, nil
}

func (r *DBOpts) Connect(ctx context.Context) (client *mongo.Client, err error) {
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(r.ConnectionString))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *MongoDB) Execute() (res any, err error) {
	collection := r.Client.Database(r.Database).Collection(r.Table)

	switch r.Action {
	case "list":
		findOpts := options.Find().SetLimit(50)

		cur, err := collection.Find(context.Background(), bson.D{}, findOpts)
		if err != nil {
			log.Printf("‼️ Failed Listing Records %v", err.Error())
			return nil, err
		}
		defer cur.Close(context.Background())

		var recordList model.DBMessageRecordList
		var ctr int64
		for cur.Next(context.Background()) {
			// To decode into a struct, use cursor.Decode()
			result := model.DBMessageRecord{}

			if err := cur.Decode(&result); err != nil {
				log.Fatal(err)
				log.Printf("‼️ Failed Decoding Records %v", err.Error())
				return nil, err
			}

			recordList.MessageList = append(recordList.MessageList, result)
			ctr += 1
		}

		recordList.Count = ctr

		if err := cur.Err(); err != nil {
			return nil, err
		}

		return recordList, err

	case "find":
		res, err := collection.FindOne(context.Background(), r.Filter).Raw()
		if err != nil {
			log.Printf("‼️ Failed Fetching Record %v", err.Error())
			return nil, err
		}

		var record model.DBMessageRecord
		if err = bson.Unmarshal(res, &record); err != nil {
			log.Printf("‼️ Failed Parsing Record %v", err.Error())
			return nil, err
		}

		return record, nil

	case "insert":
		resData, err := collection.InsertOne(context.Background(), r.Payload)
		if err != nil {
			log.Printf("‼️ Failed Inserting Data %v", err.Error())
			return nil, err
		}

		// Get the inserted ID as primitive.ObjectID
		oid, ok := resData.InsertedID.(primitive.ObjectID)
		if !ok {
			log.Printf("‼️ Failed Fetching Record - insert status (%v) ", ok)
			var err error = errors.New("failed fetching record")
			return nil, err
		}

		log.Printf("💡 Inserted ID (hex): %v", oid.Hex())

		return oid.Hex(), nil

	case "delete":
		resData, err := collection.DeleteOne(context.Background(), r.Filter)
		if err != nil {
			log.Printf("‼️ Failed Fetching Record %v", err.Error())
			return nil, err
		}

		var record model.DBMessageDeleted
		record.DeletedCount = resData.DeletedCount
		return record, nil

	default:
		var err error = errors.New("unable to process request")
		return nil, err
	}
}
