echo "Getting dependencies"
go get -v

echo "Running tests"
go test ./... -race

echo "Running application"
go run --tags=debug main.go
