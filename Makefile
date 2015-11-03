all: gresh

deps:
	go get gopkg.in/leyra/cli.v1
	go get gopkg.in/leyra/color.v1
	go get gopkg.in/leyra/go-github.v1/github
	go get golang.org/x/oauth2

gresh: deps gresh.go
	go fmt ./...
	go build -v -o gresh

run: gresh
	./gresh

install: gresh
	cp gresh /usr/local/bin/gresh

test: deps gresh.go
	go test -cover -v

clean:
	rm gresh
