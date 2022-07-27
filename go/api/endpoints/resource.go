package endpoints

import (
	"api/data"
	"api/db"
	"api/security"
	"api/utils"
	"fmt"
	"lib/logger"
	"net/http"

	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////
// Structures and Variables

type CreateResourceStruct struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Capacity    *int    `json:"capacity,omitempty"`
	Picture     *string `json:"picture,omitempty"`
}

/////////////////////////////////////////////
// Endpoints

//ResourceHandlers
func ResourceHandlers(router *mux.Router) error {
	// Resource Identifier
	router.HandleFunc("/create", security.Validate(CreateIdentifierHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/batch-create", security.Validate(BatchCreateIdentifierHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/information", security.Validate(InformationIdentifiersHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/remove", security.Validate(DeleteIdentifierHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "IDENTIFIER")})).Methods("POST")

	// Resource room
	router.HandleFunc("/room/create", security.Validate(CreateRoomHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})).Methods("POST")
	router.HandleFunc("/room/information", security.Validate(InformationRoomsHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOM")})).Methods("POST")
	router.HandleFunc("/room/remove", security.Validate(DeleteRoomHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "ROOM")})).Methods("POST")

	// Resource room association
	router.HandleFunc("/room/association/create", security.Validate(CreateRoomAssociationHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOMASSOCIATION")})).Methods("POST")
	router.HandleFunc("/room/association/information", security.Validate(InformationRoomAssociationsHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOMASSOCIATION")})).Methods("POST")
	router.HandleFunc("/room/association/remove", security.Validate(DeleteRoomAssociationHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "ROOMASSOCIATION")})).Methods("POST")

	// Resource building
	router.HandleFunc("/building/create", security.Validate(CreateBuildingHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")})).Methods("POST")
	router.HandleFunc("/building/information", security.Validate(InformationBuildingsHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "BUILDING")})).Methods("POST")
	router.HandleFunc("/building/remove", security.Validate(DeleteBuildingHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "BUILDING")})).Methods("POST")

	return nil
}

/////////////////////////////////////////////
// Functions

////////////////
// Identifier

// CreateIdentifierHandler creates or updates a Resource Identifier
func CreateIdentifierHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Resource
	var identifier data.Resource
	err := utils.UnmarshalJSON(writer, request, &identifier)
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

	// TODO [KP]: Do more checks like if there exists a Identifier already etc

	da := data.NewResourceDA(access)
	err = da.StoreIdentifier(&identifier)
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
	logger.Access.Printf("%v created\n", identifier.Id) // TODO [KP]: Be more descriptive
	utils.Ok(writer, request)
}

func BatchCreateIdentifierHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Resource
	var identifiers []*data.Resource
	err := utils.UnmarshalJSON(writer, request, &identifiers)
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

	// TODO [KP]: Do more checks like if there exists a Identifier already etc

	da := data.NewResourceDA(access)
	err = da.BatchStoreIdentifier(identifiers)
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
	for i := 0; i < len(identifiers); i++ {
		identifier := identifiers[i]
		logger.Access.Printf("Resource created. ID: %v, Name: %v\n", identifier.Id, identifier.Name)
	}
	utils.Ok(writer, request)
}

// InformationIdentifiersHandler gets resources
func InformationIdentifiersHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal resource
	var identifier data.Resource
	err := utils.UnmarshalJSON(writer, request, &identifier)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// No check for permissions the database handles information permissions

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO [KP]: null checks etc.

	da := data.NewResourceDA(access)
	identifiers, err := da.FindIdentifier(&identifier, permissions)
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
	logger.Access.Printf("%v Identifier information requested\n", identifier.RoomId) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, identifiers)
}

// DeleteIdentifierHandler removes a resource
func DeleteIdentifierHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal resource
	var identifier data.Resource
	err := utils.UnmarshalJSON(writer, request, &identifier)
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

	// Get resource information if no resource id is specified
	da := data.NewResourceDA(access)
	if identifier.Id == nil {
		temp, err := da.FindIdentifier(&identifier, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "IDENTIFIER")})
		identifier = *temp.FindHead()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	identifierRemoved, err := da.DeleteIdentifier(&identifier)
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
	logger.Access.Printf("%v Identifier removed\n", identifier.RoomId) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, identifierRemoved)
}

////////////////
// Room

// CreateRoomHandler creates or updates a room
func CreateRoomHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Room
	var room data.Room
	err := utils.UnmarshalJSON(writer, request, &room)
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

	// TODO [KP]: Do more checks like if there exists a Room already etc

	da := data.NewResourceDA(access)
	err = da.StoreRoomResource(&room)
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
	logger.Access.Printf("%v created\n", room.Id) // TODO [KP]: Be more descriptive
	utils.Ok(writer, request)
}

// InformationRoomsHandler gets rooms
func InformationRoomsHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Room
	var room data.Room
	err := utils.UnmarshalJSON(writer, request, &room)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// No check for permissions the database handles information permissions

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO [KP]: null checks etc.

	da := data.NewResourceDA(access)
	rooms, err := da.FindRoomResource(&room, permissions)
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
	logger.Access.Printf("%v Room information requested\n", room.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, rooms)
}

// DeleteRoomHandler removes a Room
func DeleteRoomHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Room
	var room data.Room
	err := utils.UnmarshalJSON(writer, request, &room)
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

	// Get resource information if no resource id is specified
	da := data.NewResourceDA(access)
	if room.Id == nil {
		temp, err := da.FindRoomResource(&room, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOM")})
		room = *temp.FindHead()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	roomRemoved, err := da.DeleteRoomResource(&room)
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
	logger.Access.Printf("%v Room removed\n", room.Id)
	utils.JSONResponse(writer, request, roomRemoved)
}

////////////////
// RoomAssociation

// CreateRoomAssociationHandler creates or updates a room association
func CreateRoomAssociationHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Room Association
	var roomAssociation data.RoomAssociation
	err := utils.UnmarshalJSON(writer, request, &roomAssociation)
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

	// TODO [KP]: Do more checks like if there exists a RoomAssociation already etc

	da := data.NewResourceDA(access)
	err = da.StoreRoomAssociationResource(&roomAssociation)
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
	logger.Access.Printf("%v created\n", roomAssociation.RoomId)
	utils.Ok(writer, request)
}

// InformationRoomAssociationsHandler gets room associations
func InformationRoomAssociationsHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Room Association
	var roomAssociation data.RoomAssociation
	err := utils.UnmarshalJSON(writer, request, &roomAssociation)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// No check for permissions the database handles information permissions

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO [KP]: null checks etc.

	da := data.NewResourceDA(access)
	roomAssociations, err := da.FindRoomAssociationResource(&roomAssociation, permissions)
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
	logger.Access.Printf("%v RoomAssociation information requested\n", roomAssociation.RoomId) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, roomAssociations)
}

// DeleteRoomAssociationHandler removes a room association
func DeleteRoomAssociationHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Room Association
	var roomAssociation data.RoomAssociation
	err := utils.UnmarshalJSON(writer, request, &roomAssociation)
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

	da := data.NewResourceDA(access)
	roomAssociationRemoved, err := da.DeleteRoomAssociationResource(&roomAssociation)
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
	logger.Access.Printf("%v RoomAssociation removed\n", roomAssociation.RoomId)
	utils.JSONResponse(writer, request, roomAssociationRemoved)
}

////////////////
// Building

// CreateBuildingHandler creates or updates a building
func CreateBuildingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Building
	var building data.Building
	err := utils.UnmarshalJSON(writer, request, &building)
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

	// TODO [KP]: Do more checks like if there exists a building already etc

	da := data.NewResourceDA(access)
	err = da.StoreBuildingResource(&building)
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
	logger.Access.Printf("%v created\n", building.Id)
	utils.Ok(writer, request)
}

// InformationBuildingsHandler gets buildings
func InformationBuildingsHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Building
	var building data.Building
	err := utils.UnmarshalJSON(writer, request, &building)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// No check for permissions the database handles information permissions

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewResourceDA(access)
	buildings, err := da.FindBuildingResource(&building, permissions)
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
	logger.Access.Printf("%v building information requested\n", building.Id)
	utils.JSONResponse(writer, request, buildings)
}

// DeleteBuildingHandler removes a building
func DeleteBuildingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Building
	var building data.Building
	err := utils.UnmarshalJSON(writer, request, &building)
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

	// Get building information if no building id is specified
	da := data.NewResourceDA(access)
	if building.Id == nil {
		temp, err := da.FindBuildingResource(&building, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "BUILDING")})
		building = *temp.FindHead()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	buildingRemoved, err := da.DeleteBuildingResource(&building)
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
	logger.Access.Printf("%v building removed\n", building.Id)
	utils.JSONResponse(writer, request, buildingRemoved)
}
