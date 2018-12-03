echo "Getting dependencies"
go get -v
go get -v -u github.com/kevinburke/go-bindata/...

go install github.com/kevinburke/go-bindata

echo "Building assets"
rm -f assets/binassets.go
go-bindata -o assets/binassets.go -pkg assets -ignore=.*\\.md -ignore=.*\\.xcf assets/...

echo "Running tests"
go test ./... -race

echo "Building application"
mkdir -p dest
go build -ldflags "-H=windowsgui -s -w" -o dest/Catastrophy_windows.exe
