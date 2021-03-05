package service

import (
	"context"

	"github.com/PlanckProject/Mind-Care/api/config"
	errorKeys "github.com/PlanckProject/Mind-Care/api/pkg/errors"
	"github.com/PlanckProject/Mind-Care/api/pkg/models"
	repository "github.com/PlanckProject/Mind-Care/api/pkg/repo"
	"github.com/PlanckProject/go-commons/errors"
	"github.com/PlanckProject/go-commons/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func New(repo repository.IServiceProvidersRepo, cfg *config.Configuration) IServiceProvidersService {
	return &serviceProvidersServiceImpl{repo: repo, cfg: cfg}
}

type IServiceProvidersService interface {
	GetByID(ctx context.Context, id string) (*models.ServiceProvider, error)
	Get(ctx context.Context, serviceProviderRequestParams *models.ServiceProviderRequestParams) ([]models.ServiceProvider, error)
	Add(ctx context.Context, serviceProvider models.ServiceProvider) (string, error)
}

type serviceProvidersServiceImpl struct {
	repo repository.IServiceProvidersRepo
	cfg  *config.Configuration
}

func (s *serviceProvidersServiceImpl) GetByID(ctx context.Context, id string) (*models.ServiceProvider, error) {
	mongoObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.
			WithFields(logger.Fields{"error": err, "id": id}).
			Error("Error parsing ID")
		return nil, errors.NewErrorWithMetadata().SetError(errorKeys.INVALID_ID.Error())
	}

	serviceProvider, err := s.repo.GetByID(ctx, mongoObjectId)
	if err == nil {
		return &serviceProvider, nil
	}

	return &serviceProvider, errors.NewErrorWithMetadata().SetError(err.Error())
}

func (s *serviceProvidersServiceImpl) Get(ctx context.Context, serviceProviderRequestParams *models.ServiceProviderRequestParams) ([]models.ServiceProvider, error) {
	var response []models.ServiceProvider
	var err error
	if serviceProviderRequestParams.Online {
		response, err = s.repo.GetOnline(ctx, serviceProviderRequestParams.Start, serviceProviderRequestParams.Limit)
	} else if serviceProviderRequestParams.Location {
		response, err = s.repo.GetNearCoordinates(ctx, &repository.LocationQueryParams{
			Lat:         serviceProviderRequestParams.LocationQuery.Geometery.Lat,
			Lon:         serviceProviderRequestParams.LocationQuery.Geometery.Lon,
			MaxDistance: serviceProviderRequestParams.LocationQuery.MaxDistance,
		},
			serviceProviderRequestParams.Start, serviceProviderRequestParams.Limit)
	} else {
		response, err = s.repo.Get(ctx, serviceProviderRequestParams.Start, serviceProviderRequestParams.Limit)
	}
	if err == nil {
		return response, err
	}

	return response, errors.NewErrorWithMetadata().SetError((err.Error()))
}

func (s *serviceProvidersServiceImpl) Add(ctx context.Context, serviceProvider models.ServiceProvider) (string, error) {
	var lat, lon float64
	var err error

	if len(serviceProvider.Services) == 0 {
		return "", errors.NewErrorWithMetadata().SetError(errorKeys.INVALID_SERVICES.Error())
	}

	if serviceProvider.Online {
		lat = 0
		lon = 0
		serviceProvider.Contact.Address.Coordinates = nil
	} else {
		lat, lon, err = getCoordinates(ctx, &serviceProvider.Contact.Address, &s.cfg.Maps)
		if err != nil {
			return "", errors.
				NewErrorWithMetadata().
				SetError(errorKeys.LOCATION_DATA_NOT_FOUND.Error()).
				SetMetadata("Unable to fetch coordinates. Please supply the coordinates to proceed")
		}
		serviceProvider.Contact.Address.Coordinates = []float64{lat, lon}
	}

	serviceProvider.Location.Type = "Point"
	serviceProvider.Location.Coordinates = []interface{}{lat, lon}
	serviceProvider.Location.Properties = make(map[string]interface{})

	id, err := s.repo.Add(ctx, serviceProvider)

	if err == nil {
		return id, nil
	}

	return id, errors.NewErrorWithMetadata().SetError(err.Error())
}
