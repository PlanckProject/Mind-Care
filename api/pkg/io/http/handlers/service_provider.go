package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/core/utils"
	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/models"
	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/service"
	"github.com/PlanckProject/go-commons/errors"
	"github.com/PlanckProject/go-commons/logger"
	"github.com/gin-gonic/gin"
)

func registerServiceProviderHandler(g *gin.Engine, cfg *config.Configuration, svc service.IServiceProvidersService) {
	{
		g.POST("/service_provider", addServiceProvider(svc))
		g.GET("/service_providers", getServiceProvidersByLocation(svc, cfg))
	}
}

func addServiceProvider(svc service.IServiceProvidersService) func(*gin.Context) {
	return func(c *gin.Context) {
		serviceProvider := &models.ServiceProvider{}
		if err := c.ShouldBindJSON(serviceProvider); err != nil {
			logger.WithField("error", err).Error("Cannot deserialize the request body")
			utils.Respond(c, http.StatusBadRequest, nil, fmt.Errorf("Cannot deserialize the request body"))
			return
		}

		id, err := svc.Add(c.Request.Context(), *serviceProvider)
		if err == nil {
			utils.Respond(c, http.StatusCreated, id, nil)
			return
		}

		merr := err.(errors.ErrorWithMetadata)
		switch merr.ErrorValue() {
		default:
			logger.WithField("error", err).Error("Error while adding data")
			utils.Respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Internal server error"))
		}
	}
}

func getServiceProvidersByLocation(svc service.IServiceProvidersService, cfg *config.Configuration) func(*gin.Context) {
	return func(c *gin.Context) {
		st, li := parseStartAndLimitQueries(c.Query("st"), c.Query("li"), cfg.App.MaxQueryLimit)
		if li == 0 {
			utils.Respond(c, http.StatusBadRequest, nil, fmt.Errorf("Invalid limit query"))
			return
		}

		var serviceProviders []models.ServiceProvider
		var err error
		serviceProviderRequestParams := &models.ServiceProviderRequestParams{Start: st, Limit: li}

		if c.Query("loc") == "true" {
			lat, lon, maxDistance, parseErr := parseLatAndLon(c.Query("lat"), c.Query("lon"), c.Query("dist"), cfg.App.MaxDistanceDefault)
			if parseErr != nil {
				utils.Respond(c, http.StatusBadRequest, nil, err)
				return
			}
			serviceProviderRequestParams.Location = true
			serviceProviderRequestParams.LocationQuery.Geometery.Lat = lat
			serviceProviderRequestParams.LocationQuery.Geometery.Lon = lon
			serviceProviderRequestParams.LocationQuery.MaxDistance = maxDistance
		}

		serviceProviders, err = svc.Get(c.Request.Context(), serviceProviderRequestParams)

		if err == nil {
			utils.Respond(c, http.StatusOK, serviceProviders, err)
			return
		}

		merr := err.(errors.ErrorWithMetadata)
		switch merr.ErrorValue() {
		default:
			logger.WithField("error", err).Error("Error while retrieving data")
			utils.Respond(c, http.StatusInternalServerError, nil, fmt.Errorf("Internal server error"))
		}
	}
}

func parseStartAndLimitQueries(start, limit string, maxQueryLimit int64) (st, li int64) {
	var err error

	st, err = strconv.ParseInt(start, 10, 64)
	if err != nil {
		logger.Warn("Unable to parse start query param, using default")
	}

	if st < 0 {
		logger.Warn("Invalid start query, using default")
		st = 0
	}

	li, err = strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logger.Warn("Unable to parse limit query param, using default")
	}

	if li > maxQueryLimit {
		li = maxQueryLimit
		logger.Warnf("Queried more items that allowed. Defaulting to %d items", li)
	}

	if li <= 0 {
		logger.Warn("Invalid limit query, dropping the request")
		li = 0
	}

	return
}

func parseLatAndLon(latString, lonString, maxDistanceString string, defaultMaxDistanceFromConfig float64) (lat, lon, maxDistance float64, err error) {
	lat, err = strconv.ParseFloat(latString, 10)
	if err != nil {
		logger.WithField("error", err).Warn("Unable to parse latitude query param, aborting")
		err = fmt.Errorf("Invalid latitude value")
		return
	}

	lon, err = strconv.ParseFloat(lonString, 10)
	if err != nil {
		logger.WithField("error", err).Warn("Unable to parse longitude query param, aborting")
		err = fmt.Errorf("Invalid longitude value")
		return
	}

	maxDistance, err = strconv.ParseFloat(maxDistanceString, 10)
	if err != nil {
		logger.WithField("error", err).Warn("Unable to parse distance query param, setting default")
		maxDistance = defaultMaxDistanceFromConfig
		err = nil
	}

	return
}
