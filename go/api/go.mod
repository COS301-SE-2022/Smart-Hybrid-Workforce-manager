module api

go 1.16

replace lib => ../lib

require (
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.5
	lib v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/ory/dockertest/v3 v3.8.1
)
