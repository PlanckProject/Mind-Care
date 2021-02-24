package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/models"
	"github.com/PlanckProject/go-commons/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
)

func NewMongoDBDataProvider(lifecycle fx.Lifecycle, cfg *config.Configuration) IServiceProvidersRepo {
	clientOptions := options.Client().ApplyURI(cfg.Mongo.ConnectionString)
	var err error
	mongoDataProvider := &mongoProvider{}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Debug("Attempting to connect to mongo")
			mongoDataProvider.c, err = mongo.Connect(ctx, clientOptions)
			if err != nil {
				return err
			}
			db := mongoDataProvider.c.Database(cfg.Mongo.Database)
			if db == nil {
				return fmt.Errorf("DB not found")
			}
			mongoDataProvider.collection = db.Collection(cfg.Mongo.Collection)
			if mongoDataProvider.collection == nil {
				return fmt.Errorf("Collection not found")
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return mongoDataProvider.c.Disconnect(ctx)
		},
	})
	return mongoDataProvider
}

type LocationQueryParams struct {
	Lat         float64
	Lon         float64
	MaxDistance float64
}

type mongoProvider struct {
	c          *mongo.Client
	collection *mongo.Collection
}

func (m *mongoProvider) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	return m.c.Ping(ctx, readpref.Primary())
}

func (m *mongoProvider) GetNearCoordinates(ctx context.Context, locationQueryParams *LocationQueryParams, skip int64, limit int64) ([]models.ServiceProvider, error) {
	logEntry := logger.NewEntry().WithContext(ctx).WithFields(logger.Fields{"latitude": locationQueryParams.Lat, "longitude": locationQueryParams.Lon, "max_distance": locationQueryParams.MaxDistance})

	opts := mongoOptions(skip, limit)

	serviceProviders := make([]models.ServiceProvider, 0)
	cursor, err := m.collection.Find(ctx, bson.M{
		"approved": true,
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{locationQueryParams.Lon, locationQueryParams.Lat},
				},
				"$maxDistance": locationQueryParams.MaxDistance,
			},
		},
	}, opts)

	if err != nil {
		logEntry.WithField("error", err).Error("Error while retrieving data from DB")
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

func (m *mongoProvider) Get(ctx context.Context, skip int64, limit int64) ([]models.ServiceProvider, error) {
	logEntry := logger.WithContext(ctx)

	opts := mongoOptions(skip, limit)

	serviceProviders := make([]models.ServiceProvider, 0)
	cursor, err := m.collection.Find(ctx, bson.M{"approved": true}, opts)

	if err != nil {
		logEntry.WithField("error", err).Error("Error while retrieving data from DB")
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

func (m *mongoProvider) GetByID(ctx context.Context, id primitive.ObjectID) (models.ServiceProvider, error) {
	serviceProvider := models.ServiceProvider{}
	result := m.collection.FindOne(ctx, bson.M{"approved": true, "_id": id})
	err := result.Decode(&serviceProvider)
	return serviceProvider, err
}

func (m *mongoProvider) Add(ctx context.Context, serviceProvider models.ServiceProvider) (string, error) {
	serviceProvider.ID = primitive.NewObjectID()
	serviceProvider.CreatedAt = time.Now()
	serviceProvider.UpdatedAt = time.Now()

	result, err := m.collection.InsertOne(ctx, serviceProvider)
	if err != nil {
		return "", err
	}

	serviceProvider.ID = result.InsertedID.(primitive.ObjectID)

	return serviceProvider.ID.Hex(), nil
}

func mongoOptions(skip int64, limit int64) *options.FindOptions {
	return &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
}
