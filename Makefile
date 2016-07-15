package = github.com/abtris/bee

.PHONY: all clean deps fast release push install

all: bee

bee: deps
	go build -ldflags "-X main.VERSION=`cat VERSION`" -o ./buid/bee ./main.go

fast:
	go build -i -ldflags "-X main.VERSION=`cat VERSION`-dev" -o ./build/bee ./main.go

deps: glide
	./glide install

glide:
	curl https://glide.sh/get | sh

clean:
	rm -fr ./build

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/bee-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/bee-linux-386 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/bee-darwin-amd64 $(package)
	GOOS=windows GOARCH=amd64 go build -o release/bee-windows-amd64 $(package)


push:
	git tag -a `cat VERSION`
	git push origin `cat VERSION`

install: deps
	go install -ldflags "-X main.VERSION=`cat VERSION`" ./main.go
