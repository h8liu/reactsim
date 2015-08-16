.PHONY: all fmt tags doc

all:
	go install -v ./...

rall:
	touch `find . -name "*.go"`
	go install -v ./...

fmt:
	gofmt -s -w -l .

tags:
	gotags -R . > tags

test:
	go test ./...

testv:
	go test -v ./...

lc:
	wc -l `find . -name "*.go"`

doc:
	godoc -http=:8000

lint:
	golint ./...
