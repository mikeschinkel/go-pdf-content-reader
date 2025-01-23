# Run the example
run: build
	./bin/go-pdf-example

# Build the example binary
build: clean
	go mod tidy \
	&& go build -o "bin/go-pdf-example" cmd/main.go

# Clean up the build artifact
clean:
	rm -f bin/go-pdf-example

.PHONY: clean build run
