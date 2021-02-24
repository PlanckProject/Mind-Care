package service

import (
	"context"
	"fmt"

	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/models"
	repository "github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/repo"
)

func New(repo repository.IServiceProvidersRepo, cfg *config.Configuration) IServiceProvidersService {
	return &serviceProvidersServiceImpl{repo: repo, cfg: cfg}
}

type IServiceProvidersService interface {
	Get(ctx context.Context, serviceProviderRequestParams *models.ServiceProviderRequestParams) ([]models.ServiceProvider, error)
	Add(ctx context.Context, serviceProvider models.ServiceProvider) (string, error)
}

type serviceProvidersServiceImpl struct {
	repo repository.IServiceProvidersRepo
	cfg  *config.Configuration
}

func (s *serviceProvidersServiceImpl) Get(ctx context.Context, serviceProviderRequestParams *models.ServiceProviderRequestParams) ([]models.ServiceProvider, error) {
	if serviceProviderRequestParams.Location {
		return s.repo.GetNearCoordinates(ctx, &repository.LocationQueryParams{
			Lat:         serviceProviderRequestParams.LocationQuery.Geometery.Lat,
			Lon:         serviceProviderRequestParams.LocationQuery.Geometery.Lon,
			MaxDistance: serviceProviderRequestParams.LocationQuery.MaxDistance,
		},
			serviceProviderRequestParams.Start, serviceProviderRequestParams.Limit)
	} else {
		return s.repo.Get(ctx, serviceProviderRequestParams.Start, serviceProviderRequestParams.Limit)
	}
}

func (s *serviceProvidersServiceImpl) Add(ctx context.Context, serviceProvider models.ServiceProvider) (string, error) {
	lat, lon, err := getCoordinates(ctx, &serviceProvider.Contact.Address, &s.cfg.Maps)
	if err != nil {

		// TODO: Convert this into error with metadata
		return "", fmt.Errorf("Unable to fetch coordinates. Please supply the coordinates to proceed")
	}

	serviceProvider.Location.Type = "Point"
	serviceProvider.Location.Coordinates = []interface{}{lon, lat}
	serviceProvider.Location.Properties = make(map[string]interface{})

	id, err := s.repo.Add(ctx, serviceProvider)

	// TODO: Convert this into error with metadata
	return id, err
}
