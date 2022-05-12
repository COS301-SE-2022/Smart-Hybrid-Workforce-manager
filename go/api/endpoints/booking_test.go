package endpoints

import (
	"fmt"
	"lib/testutils"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCreateBookingHandler(t *testing.T) {
	dbDsnEnv := os.Getenv("DATABASE_DSN")
	dbMaxIdleEnv := os.Getenv("DATABASE_MAX_IDLE_CONNECTIONS")
	dbMaxOpenEnv := os.Getenv("DATABASE_MAX_OPENCONNECTIONS")

	// Check if all environment variables could be read
	if dbDsnEnv == "" {
		// DB must be running locally in order for this test to work
		// Set environment variables
		dbDsnEnv = "host=localhost port=5432 user=admin dbname=arche sslmode=disable"
		fmt.Println(testutils.Scolourf(testutils.CYAN, "DATABASE_DSN environmet var could not be read, using %s", dbDsnEnv))
		t.Setenv("DATABASE_DSN", dbDsnEnv)
	}

	if dbMaxIdleEnv == "" {
		// DB must be running locally in order for this test to work
		dbMaxIdleEnv = "5"
		fmt.Println(testutils.Scolourf(testutils.CYAN, "DATABASE_MAX_IDLE_CONNECTIONS environmet var could not be read, using %s", dbMaxIdleEnv))
		t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", "5")
	}

	if dbMaxOpenEnv == "" {
		dbMaxOpenEnv = "5"
		fmt.Println(testutils.Scolourf(testutils.CYAN, "DATABASE_MAX_OPEN_CONNECTIONS environmet var could not be read, using %s", dbMaxOpenEnv))
		t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", "5")
	}

	writer1 := httptest.NewRecorder()
	reader1 := httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`{}`))
	InformationBookingHandler(writer1, reader1)

}
