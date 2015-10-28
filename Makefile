all: gresh

deps:
	go get gopkg.in/leyra/cli.v1
	go get gopkg.in/leyra/color.v1

gresh: deps gresh.go
	go fmt ./...
	go build -v -o gresh

run: gresh
	./gresh

install: gresh
	cp gresh /usr/local/bin/gresh

clean:
	rm gresh
