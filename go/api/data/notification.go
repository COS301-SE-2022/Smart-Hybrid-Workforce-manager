package data

//////////////////////////////////////////////////
// Structures and Variables

// Notification structure
type Notification struct {
	To        *string `json:"to,omitempty"`
	StartDate *string `json:"sDate,omitempty"`
	StartTime *string `json:"sTime,omitempty"`
	EndDate   *string `json:"eDate,omitempty"`
	EndTime   *string `json:"eTime,omitempty"`
}
