package data

import (
	"api/db"
	"database/sql"
)

// OverallStatistics
type OverallStatics struct {
	ResourceUtilisation []ResourceUtilisation
}

func mapOverallStatics(rows *sql.Rows) (interface{}, error) {
	var identifier OverallStatics
	err := rows.Scan(
		&identifier.ResourceUtilisation,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

type ResourceUtilisation struct {
	Percentage *string
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

//GetAllStatistics
// Go look at the FindIdentifier function in data/team.go
func (access *StatisticsDA) GetAllStatistics() (*OverallStatics, error) {

	return nil, nil
}
