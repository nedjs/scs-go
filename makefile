
test:
	go test

build:
	go build -ldflags "-s -w" -trimpath -o ./bin/scs

all:
	env GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -trimpath -o ./bin/scs-linux64
	env GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -trimpath -o ./bin/scs-macos64
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./bin/scs-win64