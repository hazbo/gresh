all: gresh

deps:
	go get gopkg.in/leyra/cli.v1

gresh: deps gresh.go
	go build -v -o gresh

run: gresh
	./gresh

clean:
	rm gresh
