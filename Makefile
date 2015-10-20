all: gresh

gresh: gresh.go
	go build -v -o gresh

clean:
	rm gresh
