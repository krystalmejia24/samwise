name ?= samwise
cmd = cmd/http/main.go

withColors = sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

run:
	go run $(cmd)

run_docker:
	docker-compose up

build_docker:
	docker-compose build

build:
	go build $(cmd)

bench:
	go test -bench=. ./...

test:
	go test -v -race ./... | $(withColors)

cover:
	go test -cover ./... -covermode=atomic

test_cover:
	go test -v ./... -tags test -race -coverprofile=coverage.txt -covermode=atomic | $(withColors)

.PHONY: run build test cover test_cover