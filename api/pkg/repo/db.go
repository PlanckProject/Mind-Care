package repo

import (
	"context"

	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/models"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewMongoDBDataProvider))

type IDataProvider interface {
	GetNearCoordinates(ctx context.Context, lat float64, lng float64) ([]models.ServiceProvider, error)
}
