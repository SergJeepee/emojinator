build:
	GOOS=darwin GOARCH=arm64 go build -o bin/poopinator-arm cmd/poopinator/main/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/poopinator-x64 cmd/poopinator/main/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/emojinator-arm cmd/emojinator/main/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/emojinator-x64 cmd/emojinator/main/main.go
