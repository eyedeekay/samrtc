
GOPATH=$(PWD)/.go

deps:
	go get -u github.com/kpetku/sam3

build:
	cd src/main && go build -o "$(PWD)/bin/samrtc"

noopts: build

test:
	cd src && go test

clean:
	rm -f "$(PWD)/bin/samrtc"
