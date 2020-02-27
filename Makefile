GOFILES = $(shell find . -name '*.go')

default: build

build:
	mkdir -p build

build: build/protometry

build-native: $(GOFILES)
	go build -o build/native-protometry .

build/protometry: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/protometry .