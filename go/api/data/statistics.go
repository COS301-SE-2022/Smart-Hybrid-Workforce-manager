package data

import (
	"api/db"
	"database/sql"
)

// OverallStatistics
type OverallStatics struct {
	ResourceUtilisation []ResourceUtilisation
	AverageUtilisation	AverageUtilisation `json:"average"`
}

func mapOverallStatics(rows *sql.Rows) (interface{}, error) {
	var identifier OverallStatics
	err := rows.Scan(
		&identifier.ResourceUtilisation,
		&identifier.AverageUtilisation,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

type ResourceUtilisation struct {
	Id         *string `json:"id"`
	Percentage *string `json:"percentage"`
}

func mapUtilisation(rows *sql.Rows) (interface{}, error) {
	var identifier ResourceUtilisation
	err := rows.Scan(
		&identifier.Id,
		&identifier.Percentage,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

type AverageUtilisation struct {
	Value *float64 `json:"average"`
}

func mapAverageUtilisation(rows *sql.Rows) (interface{}, error) {
	var identifier AverageUtilisation
	err := rows.Scan(
		&identifier.Value,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

// StatisticsDA
type StatisticsDA struct {
	access *db.Access
}

// NewStatisticsDA creates a new data access from a underlying shared data access
func NewStatisticsDA(access *db.Access) *StatisticsDA {
	return &StatisticsDA{
		access: access,
	}
}

// Commit commits the current implicit transaction
func (access *StatisticsDA) Commit() error {
	return access.access.Commit()
}

// GetAllStatistics
func (access *StatisticsDA) GetAllStatistics() (*OverallStatics, error) {
	results, err := access.access.Query(
		`SELECT * FROM statistics.utilisation()`, mapUtilisation)
	if err != nil {
		return nil, err
	}
	tmp := make([]ResourceUtilisation, 0)
	for r, _ := range results {
		if value, ok := results[r].(ResourceUtilisation); ok {
			tmp = append(tmp, value)
		}
	}

	results2, err2 := access.access.Query(
		`SELECT * FROM statistics.average_utilisation()`, mapAverageUtilisation)
	if err2 != nil {
		return nil, err2
	}
	tmpRes := OverallStatics{ResourceUtilisation: tmp, AverageUtilisation: results2[0].(AverageUtilisation)}

	return &tmpRes, nil
}
