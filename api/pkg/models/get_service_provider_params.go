package models

type ServiceProviderRequestParams struct {
	Online        bool
	Start         int64
	Limit         int64
	Location      bool
	LocationQuery struct {
		Geometery struct {
			Lat float64
			Lon float64
		}
		MaxDistance float64
	}
}
