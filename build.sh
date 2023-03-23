GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64.exe .
GOOS=darwin GOARCH=amd64 go build -o bin/macos-amd64 .
GOOS=darwin GOARCH=arm64 go build -o bin/macos-arm64 .
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64 .