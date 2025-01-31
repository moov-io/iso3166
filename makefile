.PHONY: all generate check
all: check

generate:
	@go run iso3166_gen.go

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./...
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=85.0 ./lint-project.sh
endif
