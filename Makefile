
GOPATH=$(PWD)/.go

test:
	cd src && go test

deps:
	go get -u github.com/kpetku/sam3

build:
	cd src/main && go build -o "$(PWD)/bin/samrtc"

clean:
	rm -f "$(PWD)/bin/samrtc"
