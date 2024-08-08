package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

func NewBaseModel() *BaseModel {
	now := time.Now()
	return &BaseModel{
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (m *BaseModel) InsertOne(ctx context.Context, collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(ctx, document)
}

func (m *BaseModel) FindOne(ctx context.Context, collection *mongo.Collection, filter interface{}) *mongo.SingleResult {
	return collection.FindOne(ctx, filter)
}

func (m *BaseModel) FindMany(ctx context.Context, collection *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {
	return collection.Find(ctx, filter)
}

func (m *BaseModel) UpdateOne(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return collection.UpdateOne(ctx, filter, update)
}

func (m *BaseModel) DeleteOne(ctx context.Context, collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
	return collection.DeleteOne(ctx, filter)
}
