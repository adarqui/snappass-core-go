all: deps test
	go build

deps:
	go get github.com/garyburd/redigo/redis
	go get code.google.com/p/go-uuid/uuid

test:
	go test
