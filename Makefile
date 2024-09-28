run: build
	@./bin/app

build:
	@templ generate
	@CGO_ENABLED=0 go build -o bin/app .
