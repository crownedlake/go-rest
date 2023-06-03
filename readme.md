# what is this
I have created a go project to run an http server at `:8080`. Please note that this should be run after completing the steps under `../seed` dir. This project uses a running postgres server, you can find the db configuration in the project itself.

# how to run
1. Make sure you have latest version of go
2. go mod
3. go mod download
4. go mod verify
5. go get -d ./...
6. go build main.go
7. go run main.go

This will start the server on port 8080. You can use this api endpoint`http://localhost:8080/available-shifts?worker_id=1234&start_date=2023-02-01T05:00:00.005Z&end_date=2023-03-02T05:00:00.005Z` don't forget that start_date and end_date are timestamps.