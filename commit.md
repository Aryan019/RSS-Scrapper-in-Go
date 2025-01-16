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