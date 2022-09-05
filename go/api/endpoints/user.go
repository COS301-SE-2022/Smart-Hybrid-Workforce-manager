package endpoints

import (
	"api/data"
	"api/db"
	"api/redis"
	"lib/utils"
	"errors"

	// "errors"
	"fmt"
	"lib/logger"
	"net/http"
	"regexp"
	"time"

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

//UserHandlers registers the user
func UserHandlers(router *mux.Router) error {
	router.HandleFunc("/register", RegisterUserHandler).Methods("POST")
	router.HandleFunc("/information", InformationUserHandler).Methods("POST")
	router.HandleFunc("/update", UpdateUserHandler).Methods("POST")
	router.HandleFunc("/remove", RemoveUserHandler).Methods("POST")
	router.HandleFunc("/login", LoginUserHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

func TempUserHandlerfunc(writer http.ResponseWriter, request *http.Request) {
	utils.Ok(writer, request)
}

// RegisterUserHandler registers a new user
func RegisterUserHandler(writer http.ResponseWriter, request *http.Request) {
	// Unmarshall register user
	var registerUserStruct RegisterUserStruct
	err := utils.UnmarshalJSON(writer, request, &registerUserStruct)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Validate email
	if !emailRegex.MatchString(registerUserStruct.Email) {
		utils.BadRequest(writer, request, "invalid_email")
		return
	}

	// Create database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// Check if user already exists
	da := data.NewUserDA(access)
	users, err := da.FindIdentifier(&data.User{Email: &registerUserStruct.Email})
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	user := users.FindHead()
	if user != nil {
		utils.BadRequest(writer, request, "user_already_exists")
		return
	}

	// Create user
	defaultWorkFromHome := false
	defaultParking := "STANDARD"
	defaultOfficeDays := 0
	defaultPreferredStartTime, timeErr := time.Parse("15:04", "09:00")
	if timeErr != nil {
		fmt.Println(timeErr)
		return
	}

	defaultPreferredEndTime, timeErr := time.Parse("15:04", "17:00")
	if timeErr != nil {
		fmt.Println(timeErr)
		return
	}
	user = &data.User{
		Identifier:         &registerUserStruct.Email,
		Email:              &registerUserStruct.Email,
		FirstName:          registerUserStruct.FirstName,
		LastName:           registerUserStruct.LastName,
		Picture:            registerUserStruct.Picture,
		WorkFromHome:       &defaultWorkFromHome,
		Parking:            &defaultParking,
		OfficeDays:         &defaultOfficeDays,
		PreferredStartTime: &defaultPreferredStartTime,
		PreferredEndTime:   &defaultPreferredEndTime,
	}

	id, err := da.StoreIdentifier(user)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	clientID := "local." + *user.Identifier
	err = da.StoreCredential(clientID, registerUserStruct.Password, *user.Identifier)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Add default user permissions
	err = addDefaultPermissions(id, access)
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
	logger.Access.Printf("%v registered\n", *user.Identifier)
	utils.JSONResponse(writer, request, user.Id)
}

func addDefaultPermissions(user string, access *db.Access) error {
	dp := data.NewPermissionDA(access)
	err := dp.StorePermission(data.CreatePermission(user, "USER", "CREATE", "BOOKING", "USER", user))
	if err != nil {
		return err
	}
	err = dp.StorePermission(data.CreatePermission(user, "USER", "VIEW", "BOOKING", "USER", user))
	if err != nil {
		return err
	}
	err = dp.StorePermission(data.CreatePermission(user, "USER", "DELETE", "BOOKING", "USER", user))
	if err != nil {
		return err
	}
	err = dp.StorePermission(data.CreatePermission(user, "USER", "VIEW", "ROLE", "USER", user))
	if err != nil {
		return err
	}
	err = dp.StorePermission(data.CreatePermission(user, "USER", "VIEW", "PERMISSION", "USER", user))
	if err != nil {
		return err
	}
	err = dp.StorePermission(data.CreatePermission(user, "USER", "CREATE", "BOOKING", "USER", ""))
	if err != nil {
		return err
	}
	err = dp.StorePermission(data.CreatePermission(user, "USER", "CREATE", "BOOKING", "TEAM", ""))
	if err != nil {
		return err
	}
	err = dp.StorePermission(data.CreatePermission(user, "USER", "CREATE", "BOOKING", "ROLE", ""))
	if err != nil {
		return err
	}
	return nil
}

func InformationUserHandler(writer http.ResponseWriter, request *http.Request) {
	var user data.User

	err := utils.UnmarshalJSON(writer, request, &user)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	logger.Access.Printf("\nlogin user\n%v\n", user)

	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewUserDA(access)

	users, err := da.FindIdentifier(&user)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v user information requested\n", user.Id)

	utils.JSONResponse(writer, request, users)
}

func LoginUserHandler(writer http.ResponseWriter, request *http.Request) { // TODO [EN]: Fix this funciton if broken
	var userCred data.Credential

	err := utils.UnmarshalJSON(writer, request, &userCred)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewUserDA(access)

	credentials, err := da.FindCredential(&userCred)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		logger.Error.Fatalf("\nerror\n%v\n", err)
		return
	}
	if(len(credentials) == 0){
		utils.AccessDenied(writer, request, errors.New("incorrect email password combination"))
		return
	}
	
	users, err := da.FindIdentifier(&data.User{Identifier: credentials[0].Identifier})
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	user := users.FindHead()
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	authData, err := redis.UserLogin(*user.Id,*credentials[0].Identifier,*user.FirstName,*user.LastName)
	if err != nil {
		logger.Error.Println("Error user login endpoint")
		utils.InternalServerError(writer, request, err)
	}
	logger.Access.Printf("User Login:%v\n", *user.Id)
	utils.JSONResponse(writer, request, authData)
}

func UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
	// Unmarshall user
	var identifier data.User
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

	da := data.NewUserDA(access)
	identifierUpdated, err := da.StoreIdentifier(&identifier)
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
	logger.Access.Printf("%v updated\n", identifier.Id)
	utils.JSONResponse(writer, request, identifierUpdated)
}

func RemoveUserHandler(writer http.ResponseWriter, request *http.Request) {
	// Unmarshal resource
	var identifier data.User
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

	// Get user information if no user id is specified
	da := data.NewUserDA(access)
	if identifier.Id == nil {
		temp, err := da.FindIdentifier(&identifier)
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
	logger.Access.Printf("%v User removed\n", identifier.Id)
	utils.JSONResponse(writer, request, identifierRemoved)
}
