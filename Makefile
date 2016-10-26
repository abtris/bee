package = github.com/abtris/bee

.PHONY: all clean deps fast release push install

all: bee

bee: deps
	go build -ldflags "-X main.VERSION=`cat VERSION`" -o ./buid/bee ./main.go

fast:
	go build -i -ldflags "-X main.VERSION=`cat VERSION`-dev" -o ./build/bee ./main.go

test:
	go test ./common

coverage:
	go test -coverprofile=c.out ./common
# go get golang.org/x/tools/cmd/cover
coverage-report:
	go tool cover -html=c.out -o coverage.html

deps: glide
	./glide install

glide:
	curl https://glide.sh/get | sh

clean:
	rm -fr ./build

build:
	GOOS=$(goos) GOARCH=$(goarch) go build -o release/bee-$(goos)-$(goarch) $(package)

release:
	mkdir -p release
	goos=linux goarch=amd64 make build
	goos=linux goarch=386 make build
	goos=darwin goarch=amd64 make build
	goos=windows goarch=amd64 make build


push:
	git tag -a `cat VERSION`
	git push origin `cat VERSION`

install: deps
	go install -ldflags "-X main.VERSION=`cat VERSION`" ./main.go
