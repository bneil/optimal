.PHONY: build clean test

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/app cmd/app/main.go

darwin:
	env GOOS=darwin go build -o bin/d_app cmd/app/main.go

watch:
	env ENV=local air -c .air.toml

test:
	go -v ./...

clean:
	rm -rf ./bin

