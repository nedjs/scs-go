
test:
	go test

all:
	go build -ldflags "-s -w" -trimpath -o ./bin/scs