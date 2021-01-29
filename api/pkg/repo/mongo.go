package repo

import (
	"context"
	"fmt"

	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/models"
	"github.com/PlanckProject/go-commons/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

func NewMongoDBDataProvider(lifecycle fx.Lifecycle, cfg *config.Configuration) IDataProvider {
	clientOptions := options.Client().ApplyURI(cfg.Mongo.ConnectionString)
	var client *mongo.Client
	var collection *mongo.Collection
	var err error

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			client, err = mongo.Connect(ctx, clientOptions)
			if err != nil {
				return err
			}
			db := client.Database(cfg.Mongo.Database)
			if db == nil {
				return fmt.Errorf("DB not found")
			}
			collection = db.Collection(cfg.Mongo.Collection)
			if collection == nil {
				return fmt.Errorf("Collection not found")
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return client.Disconnect(ctx)
		},
	})
	return &mongoProvider{c: client, collection: collection}
}

type mongoProvider struct {
	c          *mongo.Client
	collection *mongo.Collection
}

func (m *mongoProvider) GetNearCoordinates(ctx context.Context, lat float64, lng float64) ([]models.ServiceProvider, error) {
	logEntry := logger.NewEntry().WithContext(ctx).WithFields(logger.Fields{"latitude": lat, "longitude": lng})

	serviceProviders := make([]models.ServiceProvider, 0)
	cursor, err := m.collection.Find(ctx, bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometery": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat},
				},
			},
		}})
	if err != nil {
		return serviceProviders, err
	}
	for cursor.Next(ctx) {
		serviceProvider := models.ServiceProvider{}
		err := cursor.Decode(&serviceProvider)
		if err != nil {
			logEntry.WithField("error", err.Error()).Error("Error while decoding service provider")
		}
		serviceProviders = append(serviceProviders, serviceProvider)
	}
	return serviceProviders, err
}
