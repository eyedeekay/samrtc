
GOPATH=$(PWD)/.go

deps:
	go get -u github.com/kpetku/sam3
	go get -u github.com/eyedeekay/sam-forwarder
	go get -u github.com/eyedeekay/sam-forwarder/config

build:
	mkdir -p bin
	cd src/main && go build -o "$(PWD)/bin/samrtc"

noopts: build

fmt:
	gofmt -w src/*.go

lint:
	golint src/*.go

vet:
	go vet src/*.go

test: fmt lint vet
	cd src && go test

clean:
	rm -f "$(PWD)/bin/samrtc"
	find . -name *.i2pkeys -exec rm {} \;
