run:
	@-go run .

linux:
	@-env GOOS=linux GOARCH=amd64 go build

clean:
	@-rm blog
