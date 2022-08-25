package utils

import (
	"encoding/json"
	"lib/logger"
	"net/http"
	"strings"
)

/////////////////////////////////////////////
// Structures

type errorResponse struct {
	Error map[string]interface{} `json:"error"`
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

/////////////////////////////////////////////
// JSON

// UnmarshalJSON unmarshals an input interface from the reader
func UnmarshalJSON(writer http.ResponseWriter, request *http.Request, input interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)
	if err != nil {
		return err
	}
	return nil
}

/////////////////////////////////////////////
// Authorization

/////////////////////////////////////////////
// Responses

// jsonResponse
func jsonResponse(writer http.ResponseWriter, request *http.Request, status int, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(payload)
	if err != nil {
		logger.HTTP.Printf("INFO %v %v [500] Unable to Marshal value\n", request.RemoteAddr, request.RequestURI)
		return
	}
}

// JSONResponse the standard ok (200) response with JSON content
func JSONResponse(writer http.ResponseWriter, request *http.Request, value interface{}) {
	jsonResponse(writer, request, http.StatusOK, value)
	logger.HTTP.Printf("INFO %v %v [200]\n", request.RemoteAddr, request.RequestURI)
}

// BadRequest responds with bad request 400
func BadRequest(writer http.ResponseWriter, request *http.Request, message string) {
	payload := errorResponse{
		Error: map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": message,
		},
	}

	jsonResponse(writer, request, http.StatusBadRequest, payload)
	logger.HTTP.Printf("INFO %v %v [400] %v\n", request.RemoteAddr, request.RequestURI, message)
}

// InternalServerError responds with internal server errror 500
func InternalServerError(writer http.ResponseWriter, request *http.Request, err error) {
	payload := errorResponse{
		Error: map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err,
		},
	}

	jsonResponse(writer, request, http.StatusInternalServerError, payload)
	logger.HTTP.Printf("ERROR %v %v [500] internal_server_error %v\n", request.RemoteAddr, request.RequestURI, err)
}

// Ok responds with ok response 200
func Ok(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusOK)
	logger.HTTP.Printf("INFO %v %v [200]\n", request.RemoteAddr, request.RequestURI)
}

// AccessDenied provides the access denied (401) response
func AccessDenied(writer http.ResponseWriter, request *http.Request, accessError error) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusUnauthorized)
	logger.HTTP.Printf("INFO %v %v [401] access_denied %v\n", request.RemoteAddr, request.RequestURI, accessError)

	// Get address that the request originates from
	address := clientIP(request)

	// TODO get offending identifier
	identifier := ""
	if identifier != "" {
		logger.Access.Printf("%v %v [401] access_denied for %v %v\n", address, request.RequestURI, identifier, accessError)
		return
	}

	logger.Access.Printf("%v %v [401] access_denied %v\n", address, request.RequestURI, accessError)
}

// clientIP returns the ip of the requesting client
func clientIP(request *http.Request) string {
	ipAddress := request.Header.Get("X-Real-Ip")
	if ipAddress != "" {
		return ipAddress
	}

	ipAddress = request.Header.Get("X-Forwarded-For")
	if ipAddress != "" {
		if strings.Contains(ipAddress, ",") {
			ips := strings.Split(ipAddress, ",")
			if len(ips) > 0 {
				if ips[0] != "" {
					return strings.TrimSpace(ips[0])
				}
			}
		} else {
			return ipAddress
		}
	}

	return request.RemoteAddr
}
