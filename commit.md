Initial commit in here
Running in the project with - 
go build && ./RSS-Scrapper

log.fatal leaves in the program immediately with error code 1 
dont want to manually set in the env file want to add in a functionality by which it automatically retrieves in the port 

That is done by using go get github.com/joho/godotenv

go mod vendor -> gives you an local copy of the modules and dependencies

This helps in pulling in the env - 
godotenv.Load()

run in go mod tidy 
and go mod vendor to download in packages locally 

github.com/go-chi/chi -> Used to include in the router for the server
go get github.com/go-chi/cors

Boiler plate for http server completed till now 

go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

Installing in the goose here 
go install github.com/pressly/goose/v3/cmd/goose@latest


anything inside --goose up is a up statement 
and anything downside --goose down is a down statement 