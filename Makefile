run-dev:
	export APP_ENV="dev" && \
 	go run -race cmd/main.go

run-prod:
	export APP_ENV="prod" && \
    go run -race cmd/main.go

build:
	docker build --tag abahernest/file-upload-api:$(version) --platform=linux/amd64 .

test:
	go test -failfast -race -v ./...

lint:
	golangci-lint run ./...

test-all-pkg:
	cd ./pkg && go test -failfast -race -v ./...

# test a single package
test-pkg:
	cd $(pkg) && go test -failfast -race -v ./...
	# usage: make test-pkg pkg=./pkg/logger

# test single function within a package
test-pkg-fxn:
	cd $(pkg) && go test -failfast -race -v -run $(fxn) ./...
    # usage: make test-pkg-fxn pkg=./pkg/logger fxn=TestInitLogger