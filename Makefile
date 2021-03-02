.PHONY: all

all: get-deps # clean

get-deps: scripts/get-deps.sh
	@sh $<

clean:
	@echo "do clean"

run-mark: cmd/db-mark/main.go
	@go run $<