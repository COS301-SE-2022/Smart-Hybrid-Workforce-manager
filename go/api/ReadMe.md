# General
 
The api will be compiled and run through Docker as specified in `DockerFile`

This image and its corresponding container need to be restarted in order to compile changes made to the api

# Structure

## Endpoints

All api endpoints are stored in `endpoints`. The endpoints are routed from main.go

## Database conenction

The functions managing database conenctions are stored in `db`

## Data access

All functions relating to data access from the postgres database is stored in `data`

## Testing

`go test` this will just run the tests normally
`go test -v` this will run the tests in verbose mode
`go test -cover` this will run the test and include code coverage
