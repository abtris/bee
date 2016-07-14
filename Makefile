.PHONY: all clean fast release install

all: bee

bee:
	go build -ldflags "-X main.VERSION=`cat VERSION`" -o ./buid/bee ./main.go

fast:
	go build -i -ldflags "-X main.VERSION=`cat VERSION`-dev" -o ./build/bee ./main.go

clean:
	rm -fr ./build

release: bee
	git tag -a `cat VERSION`
	git push origin `cat VERSION`

install:
	go install -ldflags "-X main.VERSION=`cat VERSION`" ./main.go
