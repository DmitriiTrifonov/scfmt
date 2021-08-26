.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	mkdir -p bin
	go build -o ./bin/scfmt ./cmd/scfmt/main.go

.PHONY: test_run
test_run: build
	./bin/scfmt example/hello_world.c

.PHONY: test_run_stdout
test_run_stout: build
	./bin/scfmt -stdout example/hello_world.c