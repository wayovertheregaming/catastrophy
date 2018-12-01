echo "Getting dependencies"
go get -v
go get -v github.com/kevinburke/go-bindata

go install github.com/kevinburke/go-bindata

echo "Building assets"
go-bindata -o assets/binassets.go -pkg assets -debug assets/...

echo "Running tests"
go test ./... -race

echo "Running application"
go run --tags=debug main.go
