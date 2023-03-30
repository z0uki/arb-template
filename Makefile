PROJECT=arb

.PHONY: build
build:
	go build -o build/${PROJECT} ./main.go
	@echo "Build successfully"

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/${PROJECT} ./main.go
	@echo "Build successfully"

.PHONY: run
run:
	go run ./main.go

.PHONY: dev
dev:
	go build -o build/${PROJECT} ./main.go && cd build && ./${PROJECT}

