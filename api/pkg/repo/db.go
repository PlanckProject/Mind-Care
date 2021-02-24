package repo

import (
	"context"

	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewMongoDBDataProvider))

type IServiceProvidersRepo interface {
	Ping(ctx context.Context) error
	GetNearCoordinates(ctx context.Context, locationQueryParams *LocationQueryParams, skip int64, limit int64) ([]models.ServiceProvider, error)
	Get(ctx context.Context, skip int64, limit int64) ([]models.ServiceProvider, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (models.ServiceProvider, error)
	Add(ctx context.Context, serviceProvider models.ServiceProvider) (string, error)
}
