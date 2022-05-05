package endpoints

import (
	"api/data"
	"api/db"
	"api/utils"
	"lib/logger"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////
// Structures and Variables

var emailRegex = regexp.MustCompile(`^(?:[^@\t\n ])+@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]*$`)

type RegisterUserStruct struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     string  `json:"email"`
	Picture   *string `json:"picture,omitempty"`
	Password  *string `json:"password"`
}

/////////////////////////////////////////////
// Endpoints

//RegisterUserHandlers registers the user
func RegisterUserHandlers(router *mux.Router) error {
	router.HandleFunc("/register", RegisterUserHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// RegisterUserHandler registers a new user
func RegisterUserHandler(writer http.ResponseWriter, request *http.Request) {
	var registerUserStruct RegisterUserStruct

	err := utils.UnmarshalJSON(writer, request, &registerUserStruct)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	if !emailRegex.MatchString(registerUserStruct.Email) {
		utils.BadRequest(writer, request, "invalid_email")
		return
	}

	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewUserDA(access)

	user, err := da.FindIdentifier(registerUserStruct.Email)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	if user != nil {
		utils.BadRequest(writer, request, "user_already_exists")
		return
	}

	user = &data.User{
		Identifier: registerUserStruct.Email,
		Email:      &registerUserStruct.Email,
		FirstName:  registerUserStruct.FirstName,
		LastName:   registerUserStruct.LastName,
		Picture:    registerUserStruct.Picture,
	}

	err = da.StoreIdentifier(user)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	clientID := "local." + user.Identifier
	err = da.StoreCredential(clientID, registerUserStruct.Password, user.Identifier)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v registered\n", user.Identifier)

	utils.Ok(writer, request)
}
