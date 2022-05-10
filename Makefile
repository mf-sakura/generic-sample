.PHONY: build
build:
	go build -o ./bin/sample .

.PHONY: run
run:
	./bin/sample
