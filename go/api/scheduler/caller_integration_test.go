package scheduler

import (
	"api/db"
	"net/http"
	"testing"
	"time"
)

func Test_Integration_call(t *testing.T) {
	t.SkipNow() // Skip till properly implemented
	HTTPClient = &http.Client{
		Timeout: timeout,
	}

	t.Setenv("DATABASE_DSN", "host=127.0.0.1 port=5432 user=admin dbname=arche sslmode=disable")
	t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", "5")
	t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", "25")
	t.Setenv("DATABASE_URL", "postgresql://admin:admin@localhost:5432/db?schema=public")

	endpointURL = "http://127.0.0.1:8101/echo" // Mock endpoint to test if communication is possible

	err := db.RegisterAccess()
	if err != nil {
		t.Fatal(err)
	}

	from := time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC)
	to := time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)
	schedulerData, err := GetSchedulerData(from, to)
	if err != nil {
		t.Fatalf("Err while calling GetSchedulerData(), error = %v", err)
	}

	err = Call(schedulerData)
	if err != nil {
		t.Fatalf("Err while calling Call(), error = %v", err)
	}

}
