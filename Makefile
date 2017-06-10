all:
	goimports -w *.go
	go build
