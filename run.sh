echo "Getting dependencies"
go get -v
go get -v -u github.com/kevinburke/go-bindata/...

go install github.com/kevinburke/go-bindata

echo "Building assets"
rm -f assets/binassets.go
go-bindata -o assets/binassets.go -pkg assets -debug -ignore=.*\\.md -ignore=.*\\.xcf assets/...

echo "Running tests"
go test ./... -race

echo "Running application"
go run --tags=debug main.go
