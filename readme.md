# what is this
This is a code demo of a simple rest API built with go and postgres without using any other library.

# how to run the go application?
Make sure you have started your postgres server before doing this.

1. Make sure you have latest version of go
2. go mod
3. go mod download
4. go mod verify
5. go get -d ./...
6. go build main.go
7. go run main.go

This will start the server on port 8080. Currrent working api endpoint is - `http://localhost:8080/available-shifts?worker_id=1234&start_date=2023-02-01T05:00:00.005Z&end_date=2023-03-02T05:00:00.005Z`, this will return in a http errorcode as your local db server will not understand this, you will have to change this for your needs as mentioned below.

# Please note
- you have to modify your db configuartion and api path for your use case
- you may wanna chane your server configuration also, by default the server will run at `:8080`
- I have tried to keep the api logic separate so that it is easy for you to replace it
