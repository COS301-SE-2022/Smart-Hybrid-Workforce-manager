package endpoints

import (
	"api/data"
	"api/db"
	"fmt"
	"lib/logger"
	"lib/utils"
	"net/http"

	"github.com/gorilla/mux"
)

//StatisticsHandlers registers the user
func StatisticsHandlers(router *mux.Router) error {
	// router.HandleFunc("/resource_utilisation", ResourceUtilisation).Methods("POST")
	router.HandleFunc("/all", AllHandler).Methods("POST")
	return nil
}

func AllHandler(writer http.ResponseWriter, request *http.Request) {
	var statistics data.OverallStatics
	err := utils.UnmarshalJSON(writer, request, &statistics)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewStatisticsDA(access)
	all_statistics, err := da.GetAllStatistics()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Print("All statistics information requested\n")
	utils.JSONResponse(writer, request, all_statistics)
}
