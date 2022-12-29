package withmongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gogen-transaction/domain_belajartransaction/model/entity"
	"gogen-transaction/shared/gogen"
	"gogen-transaction/shared/infrastructure/config"
	"gogen-transaction/shared/infrastructure/database"
	"gogen-transaction/shared/infrastructure/logger"
)

type gateway struct {
	*database.MongoWithTransaction
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	//client  *mongo.Client
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	uri := "mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=rs0&readPreference=primary&ssl=false"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &gateway{
		MongoWithTransaction: database.NewMongoDBWithTransaction(client, log),
		log:                  log,
		appData:              appData,
		config:               cfg,
		//client:               client,
	}
}

func (r *gateway) SaveProduct(ctx context.Context, obj *entity.Product) error {
	r.log.Info(ctx, "called Product")

	coll := r.MongoClient.Database("belajartrx_db").Collection("product")

	filter := bson.D{{"_id", obj.ID}}
	update := bson.D{{"$set", obj}}
	opts := options.Update().SetUpsert(true)

	result, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	r.log.Info(ctx, "inserted with id %v", result.ModifiedCount)

	return nil
}

func (r *gateway) SaveOrder(ctx context.Context, obj *entity.Order) error {
	r.log.Info(ctx, "called SaveOrder")

	coll := r.MongoClient.Database("belajartrx_db").Collection("order")

	filter := bson.D{{"_id", obj.ID}}
	update := bson.D{{"$set", obj}}
	opts := options.Update().SetUpsert(true)

	result, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	r.log.Info(ctx, "inserted with id %v", result.ModifiedCount)

	return nil
}
