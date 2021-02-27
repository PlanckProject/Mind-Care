package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/config"
	"github.com/PlanckProject/Mental-Wellbeing-Resources/api/pkg/models"
	"github.com/PlanckProject/go-commons/http/request"
	"github.com/PlanckProject/go-commons/logger"
)

type osmResponseItem struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}

type bingMapsResponse struct {
	ResourcesSet []struct {
		Resources []struct {
			Point struct {
				Coordinates []float64 `json:"coordinates"`
			} `json:"point"`
		} `json:"resources"`
	} `json:"resourceSets"`
}

type googleMapsResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

// Since this in an internal module, this is not exported.
func getCoordinates(ctx context.Context, address *models.Address, mapsConfig *config.MapsConfig) (lat float64, lon float64, err error) {
	if len(address.Coordinates) == 2 {
		if address.Coordinates[0] < -90 || address.Coordinates[0] > 90 {
			return 0, 0, fmt.Errorf("Invalid latitude")
		} else if address.Coordinates[1] < -180 || address.Coordinates[1] > 180 {
			return 0, 0, fmt.Errorf("Invalid longitude")
		}
		return address.Coordinates[0], address.Coordinates[1], nil
	}

	// Generate address string
	addressString := getAddressAsString(address)
	logEntry := logger.WithField("address_string", addressString)

	// Check Google Maps
	if mapsConfig.GoogleMaps.Enabled {
		logEntry.Info("Using Google maps to fetch coordinates")
		lat, lon, err = getCoordinatesFromGoogleMaps(ctx, addressString, mapsConfig)
		if err != nil {
			logEntry.Warn("Could not fetch coordinates from Google Maps")
		} else {
			return
		}
	}

	// Check Bing Maps
	if mapsConfig.BingMaps.Enabled {
		logEntry.Info("Using Bing maps to fetch coordinates")
		lat, lon, err = getCoordinatesFromBingMaps(ctx, addressString, mapsConfig)
		if err != nil {
			logEntry.Warn("Could not fetch coordinates from Bing Maps")
		} else {
			return
		}
	}

	// Check OSM
	if mapsConfig.OSM.Enabled {
		logEntry.Info("Using OSM to fetch coordinates")
		lat, lon, err = getCoordinatesFromOSM(ctx, addressString, mapsConfig)
		if err != nil {
			logEntry.Warn("Could not fetch coordinates from OSM")
		} else {
			return
		}
	}

	return 0, 0, fmt.Errorf("Could not locate the coordinates of the given address")
}

func getCoordinatesFromOSM(ctx context.Context, addressStr string, mapsConfig *config.MapsConfig) (float64, float64, error) {
	queryUrl := strings.Replace(mapsConfig.OSM.URL, mapsConfig.LocationQueryPlaceholder, addressStr, 1)

	responseBody, err := makeHTTPRequest(ctx, queryUrl)
	if err != nil {
		return 0, 0, err
	}

	osmResponse := make([]osmResponseItem, 0)
	_ = json.Unmarshal(responseBody, &osmResponse)
	if len(osmResponse) == 0 {
		return 0, 0, fmt.Errorf("No responses found from OSM")
	}

	lat, _ := strconv.ParseFloat(osmResponse[0].Latitude, 64)
	lon, _ := strconv.ParseFloat(osmResponse[0].Longitude, 64)

	return lat, lon, nil
}

func getCoordinatesFromBingMaps(ctx context.Context, addressStr string, mapsConfig *config.MapsConfig) (float64, float64, error) {
	queryUrl := strings.Replace(mapsConfig.BingMaps.URL, mapsConfig.LocationQueryPlaceholder, addressStr, 1)
	queryUrl = strings.Replace(queryUrl, mapsConfig.ProviderAPIKeyPlaceholder, mapsConfig.BingMaps.Key, 1)

	parsedURL, _ := url.Parse(queryUrl)

	requestURL := fmt.Sprintf("%s://%s%s?%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path, parsedURL.Query().Encode())
	responseBody, err := makeHTTPRequest(ctx, requestURL)
	if err != nil {
		return 0, 0, err
	}

	bingMapsResp := bingMapsResponse{}
	_ = json.Unmarshal(responseBody, &bingMapsResp)
	if len(bingMapsResp.ResourcesSet) == 0 {
		return 0, 0, fmt.Errorf("No responses found from Bing Maps")
	}
	if len(bingMapsResp.ResourcesSet[0].Resources) == 0 {
		return 0, 0, fmt.Errorf("An empty resource set was found")
	}
	return bingMapsResp.ResourcesSet[0].Resources[0].Point.Coordinates[0], bingMapsResp.ResourcesSet[0].Resources[0].Point.Coordinates[1], nil
}

func getCoordinatesFromGoogleMaps(ctx context.Context, addressStr string, mapsConfig *config.MapsConfig) (float64, float64, error) {
	queryUrl := strings.Replace(mapsConfig.GoogleMaps.URL, mapsConfig.LocationQueryPlaceholder, addressStr, 1)
	queryUrl = strings.Replace(queryUrl, mapsConfig.ProviderAPIKeyPlaceholder, mapsConfig.GoogleMaps.Key, 1)

	parsedURL, _ := url.Parse(queryUrl)

	requestURL := fmt.Sprintf("%s://%s%s?%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path, parsedURL.Query().Encode())
	responseBody, err := makeHTTPRequest(ctx, requestURL)
	if err != nil {
		return 0, 0, err
	}

	googleMapsResp := googleMapsResponse{}
	_ = json.Unmarshal(responseBody, &googleMapsResp)
	if len(googleMapsResp.Results) == 0 {
		return 0, 0, fmt.Errorf("No responses found from Google Maps")
	}
	return googleMapsResp.Results[0].Geometry.Location.Lat, googleMapsResp.Results[0].Geometry.Location.Lon, nil
}

func getAddressAsString(address *models.Address) string {
	result := ""

	if len(address.StreetAddress1) != 0 {
		result += address.StreetAddress1
	}

	if len(address.StreetAddress2) != 0 {
		result += address.StreetAddress2
	}

	if len(address.City) != 0 {
		if len(result) > 0 {
			result += ","
		}
		result += address.City
	}

	if len(address.State) != 0 {
		if len(result) > 0 {
			result += ","
		}
		result += address.State
	}

	if len(address.Country) != 0 {
		if len(result) > 0 {
			result += ","
		}
		result += address.Country
	}

	if len(address.ZipCode) != 0 {
		if len(result) > 0 {
			result += ","
		}
		result += address.ZipCode
	}

	return result
}

func makeHTTPRequest(ctx context.Context, uri string) (json.RawMessage, error) {
	resp, err := request.New().SetContext(ctx).SetMethod(http.MethodGet).SetURI(uri).Do()
	if err != nil {
		logger.WithFields(logger.Fields{"request.uri": uri,
			"response.code": resp.StatusCode,
		}).Error("Http request failed")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Warn("HTTP request is complete but did not return a success code")
	}

	var rawBody json.RawMessage
	var respBodyReadBodyError error

	if resp.Body != nil {
		rawBody, respBodyReadBodyError = ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		logger.WithFields(logger.Fields{"request.uri": uri,
			"response.code":            resp.StatusCode,
			"response.body":            string(rawBody),
			"response.body.read_error": respBodyReadBodyError,
		}).Info("HTTP Response received")
	} else {
		return nil, fmt.Errorf("Response contains a nil body reader")
	}

	return rawBody, nil
}
