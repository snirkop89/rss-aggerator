build:
	go build -o bin/rssagg .

run: build
	./bin/rssagg

tidy:
	@go mod tidy
	@go mod vendor