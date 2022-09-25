package data

import (
	"api/db"
	"database/sql"
	"time"
)

// OverallStatistics
type OverallStatics struct {
	ResourceUtilisation []ResourceUtilisation
	AverageUtilisation	AverageUtilisation 	`json:"average"`
	AutomatedRatio		AutomatedRatio		`json:"automated"`
	CurrentOccupancy	CurrentOccupancy	`json:"occupancy"`
	YearlyUtilisation	[]YearlyUtilisation	`json:"yearly_utilisation"`
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

type AutomatedRatio struct {
	Automated	*int64 `json:"automated"`
	Manual		*int64 `json:"manual"`
}

func mapAutomatedRatio(rows *sql.Rows) (interface{}, error) {
	var identifier AutomatedRatio
	err := rows.Scan(
		&identifier.Automated,
		&identifier.Manual,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

type CurrentOccupancy struct {
	Occupied	*int64 `json:"occupied"`
	Total		*int64 `json:"total"`
}

func mapCurrentOccupancy(rows *sql.Rows) (interface{}, error) {
	var identifier CurrentOccupancy
	err := rows.Scan(
		&identifier.Occupied,
		&identifier.Total,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

type YearlyUtilisation struct {
	Month			*time.Time	`json:"date"`
	Percentage		*float64 	`json:"percentage"`
}

func mapYearlyUtilisation(rows *sql.Rows) (interface{}, error) {
	var identifier YearlyUtilisation
	err := rows.Scan(
		&identifier.Month,
		&identifier.Percentage,
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
	// Utilisation
	results, err := access.access.Query(
		`SELECT * FROM statistics.utilisation()`, mapUtilisation)
	if err != nil {
		return nil, err
	}
	tmp1 := make([]ResourceUtilisation, 0)
	for r, _ := range results {
		if value, ok := results[r].(ResourceUtilisation); ok {
			tmp1 = append(tmp1, value)
		}
	}

	// Average Utilisation
	results2, err2 := access.access.Query(
		`SELECT * FROM statistics.average_utilisation()`, mapAverageUtilisation)
	if err2 != nil {
		return nil, err2
	}

	// Automated Ratio
	results3, err3 := access.access.Query(
		`SELECT * FROM statistics.automated_ratio()`, mapAutomatedRatio)
	if err3 != nil {
		return nil, err3
	}

	// Current Occupancy
	results4, err4 := access.access.Query(
		`SELECT * FROM statistics.current_occupancy()`, mapCurrentOccupancy)
	if err4 != nil {
		return nil, err4
	}

	// Yearly Utilisation
	results5, err5 := access.access.Query(
		`SELECT * FROM statistics.yearly_utilisation()`, mapYearlyUtilisation)
	tmp5 := make([]YearlyUtilisation, 0)
	if err5 != nil {
		return nil, err5
	}
	for r, _ := range results5 {
		if value, ok := results5[r].(YearlyUtilisation); ok {
			tmp5 = append(tmp5, value)
		}
	}

	tmpRes := OverallStatics{ResourceUtilisation: tmp1, AverageUtilisation: results2[0].(AverageUtilisation), AutomatedRatio: results3[0].(AutomatedRatio), CurrentOccupancy: results4[0].(CurrentOccupancy), YearlyUtilisation: tmp5}
	return &tmpRes, nil
}
