package storage

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/balugcath/mongo-test/pkg/types"
)

// MongoClient ...
type MongoClient struct {
	*mongo.Client
}

// NewMongoClient ...
func NewMongoClient(conn string) (*MongoClient, error) {
	var (
		err error
		s   = &MongoClient{}
	)

	s.Client, err = mongo.NewClient(options.Client().ApplyURI(conn))
	if err != nil {
		return nil, fmt.Errorf("%w, %s", types.ErrStorageMongoDB, err)
	}

	if err = s.Client.Connect(context.TODO()); err != nil {
		return nil, fmt.Errorf("%w, %s", types.ErrStorageMongoDB, err)
	}

	if err = s.Client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("%w, %s", types.ErrStorageMongoDB, err)
	}

	return s, nil
}

// Close ...
func (s MongoClient) Close() error {
	if err := s.Client.Disconnect(context.TODO()); err != nil {
		return fmt.Errorf("%w, %s", types.ErrStorageMongoDB, err)
	}
	return nil
}

// MongoCollection ...
type MongoCollection struct {
	client     *MongoClient
	collection *mongo.Collection
}

// NewMongoCollection ...
func NewMongoCollection(client *MongoClient, database, collection string) *MongoCollection {
	s := &MongoCollection{client: client}
	s.collection = s.client.Database(database).Collection(collection)
	return s
}

// StoreProductPrice ...
func (s *MongoCollection) StoreProductPrice(data []types.ProductPrice) error {

	opts := options.Update().SetUpsert(true)

	for _, v := range data {

		update := bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: types.ProductNameMongoTag, Value: v.ProductName},
					{Key: types.PriceMogoTag, Value: v.Price},
				},
			},
			{
				Key: "$currentDate", Value: bson.D{
					{Key: types.LastUpdateMongoTag, Value: true},
				},
			},
			{
				Key: "$inc", Value: bson.D{
					{Key: types.CountUpdateMongoTag, Value: 1},
				},
			},
		}

		filter := bson.D{
			{Key: types.ProductNameMongoTag, Value: v.ProductName},
		}

		if _, err := s.collection.UpdateOne(context.TODO(), filter, update, opts); err != nil {
			return fmt.Errorf("%w, %s", types.ErrStorageMongoDB, err)
		}
	}
	return nil
}

// FetchProductPrice ...
func (s *MongoCollection) FetchProductPrice(pageParams types.ProductPriceListParams) ([]types.ProductPrice, error) {
	if pageParams.LenPage == 0 {
		return nil, types.ErrLenPage
	}
	if pageParams.SortField == "" {
		return nil, types.ErrSortField
	}
	if pageParams.SortOrder == 0 {
		return nil, types.ErrSortOrder
	}

	opts := options.Find().
		SetSort(bson.D{{Key: pageParams.SortField, Value: pageParams.SortOrder}}).
		SetSkip(pageParams.Page * pageParams.LenPage).
		SetLimit(pageParams.LenPage)

	cursor, err := s.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("%w, %s", types.ErrStorageMongoDB, err)
	}
	defer cursor.Close(context.TODO())

	var (
		res []types.ProductPrice
		p   types.ProductPrice
	)

	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&p); err != nil {
			log.Println(err)
		}
		res = append(res, p)
	}

	return res, nil
}
