.EXPORT_ALL_VARIABLES:

DATABASE_URL=postgresql://postgres:postgres@localhost:5432/port

run: 
	go run cmd/cli/main.go parser --file=$(file) 


test :
	go clean -testcache
	go test ./... -race 